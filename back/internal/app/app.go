package app

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pttrulez/activitypeople/internal/config"
	"github.com/pttrulez/activitypeople/internal/domain"
	httpserver "github.com/pttrulez/activitypeople/internal/infra/http-server"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/handler"
	"github.com/pttrulez/activitypeople/internal/infra/store/pgstore"
	"github.com/pttrulez/activitypeople/internal/infra/strava"
	"github.com/pttrulez/activitypeople/internal/service/activities"
	"github.com/pttrulez/activitypeople/internal/service/auth"
	"github.com/pttrulez/activitypeople/internal/service/food"
)

func StartServer() {
	cfg := config.MustLoadConfig()
	// validator := validator.New()

	// stores
	pgConn := pgstore.CreatePGConnection(cfg.Postgres)

	activitiesStore := pgstore.NewActivitiesPostgres(pgConn)
	foodStore := pgstore.NewFoodPostgres(pgConn)
	mealStore := pgstore.NewMealPostgres(pgConn)
	stravaStore := pgstore.NewStravaPostgres(pgConn)
	userStore := pgstore.NewUserPostgres(pgConn)

	// clients
	strava := strava.NewStrava(cfg.Strava.ClientID, cfg.Strava.ClientSecret)

	// services
	authService := auth.NewService(userStore)
	activitiesService := activities.NewService(activitiesStore, strava, stravaStore)
	foodService := food.NewFoodService(foodStore, mealStore)

	// Routing
	echo.NotFoundHandler = func(c echo.Context) error {
		// render your 404 page
		return c.String(http.StatusNotFound, "not found page")
	}
	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "localhost"},
		AllowCredentials: true,
	}))

	e.HTTPErrorHandler = httpserver.HTTPErrorHandler

	// static files
	e.Static("/public", "public")

	// controllers
	authController := handler.NewAuthController(authService, cfg.JwtSecret)
	activitiesController := handler.NewActivitiesController(activitiesService)
	foodController := handler.NewFoodController(foodService)
	mealController := handler.NewMealController(foodService)

	// handlers
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)
	// e.POST("/logout", authController.Logout)

	authorized := e.Group("")
	authorized.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(cfg.JwtSecret),
		ErrorHandler: func(c echo.Context, err error) error {
			c.Response().WriteHeader(http.StatusUnauthorized)
			return echo.NewHTTPError(http.StatusUnauthorized, "cannot go mista", err)
		},
		SuccessHandler: func(c echo.Context) {
			token, _ := c.Get("user").(*jwt.Token)
			claims, ok := token.Claims.(*handler.JwtClaims)
			if !ok {
				fmt.Println("failed to cast Token to JwtClaims")
			}
			user := domain.User{
				Email: claims.Email,
				Id:    claims.Id,
				Name:  claims.Name,
				Role:  claims.Role,
			}

			c.Set("u", user)
		},
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &handler.JwtClaims{}
		},
	}))

	// // Activities
	authorized.GET("/activities", activitiesController.GetActivities)

	// // Diary
	// authorized.GET("/diary", handler.Make(activitiesController.GetActivitiesPage))

	// Food
	authorized.GET("/food/search", foodController.Search)
	authorized.POST("/food", foodController.CreateFood)
	authorized.DELETE("/food/:id", foodController.DeleteFood)

	// Meal
	authorized.GET("/meal", mealController.GetMeals)
	authorized.POST("/meal", mealController.CreateMeal)

	// Strava
	authorized.GET("/strava-oauth", activitiesController.OAuthStrava)
	authorized.GET("/sync-strava", activitiesController.SyncStrava)

	fmt.Printf("Listening on port %s\n", cfg.HttpListenPort)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.HttpListenPort)))
}
