package app

import (
	"fmt"
	"net/http"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pttrulez/activitypeople/internal/config"
	httpserver "github.com/pttrulez/activitypeople/internal/infra/http-server"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/handler"
	"github.com/pttrulez/activitypeople/internal/infra/store/pgstore"
	"github.com/pttrulez/activitypeople/internal/service/auth"
)

func StartServer() {
	cfg := config.MustLoadConfig()
	// validator := validator.New()

	// stores
	pgConn := pgstore.CreatePGConnection(cfg.Postgres)
	// activitiesStore := pgstore.NewActivitiesPostgres(pgConn)
	// stravaStore := pgstore.NewStravaPostgres(pgConn)
	userStore := pgstore.NewUserPostgres(pgConn)
	// foodStore := pgstore.NewFoodPostgres(pgConn)

	// clients
	// strava := strava.NewStrava(cfg.Strava.ClientID, cfg.Strava.ClientSecret)

	// services
	authService := auth.NewService(userStore)
	// activitiesService := activities.NewService(activitiesStore, strava, stravaStore)
	// foodService := food.NewFoodService(foodStore)

	// Routing
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "localhost"},
		AllowCredentials: true,
	}))

	e.HTTPErrorHandler = httpserver.HTTPErrorHandler

	// static files
	e.Static("/public", "public")

	// controllers
	authController := handler.NewAuthController(authService, cfg.JwtSecret)
	// homeCocntroller := handler.NewHomeController(cfg.Strava.OAuthLink)
	// activitiesController := handler.NewActivitiesController(activitiesService)
	// foodController := handler.NewFoodController(foodService, validator)
	// stravaController := handler.NewStravaController(sessionStore, activitiesService)

	// handlers
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)
	// e.POST("/logout", authController.Logout)

	authorized := e.Group("")
	authorized.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(cfg.JwtSecret),
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "cannot go mista")
		},
	}))

	ro := authorized.GET("/", authController.Login)
	fmt.Println("hui", ro.Path)

	// // Activities
	// authorized.GET("/activities", handler.Make(activitiesController.GetActivitiesPage))

	// // Diary
	// authorized.GET("/diary", handler.Make(activitiesController.GetActivitiesPage))

	// // Food
	// authorized.GET("/food/search", handler.Make(foodController.Search))
	// authorized.POST("/food", handler.Make(foodController.CreateFood))
	// authorized.DELETE("/food/{id}", handler.Make(activitiesController.GetActivitiesPage))

	// // Strava
	// authorized.GET("/sync-strava", handler.Make(stravaController.SyncStrava))
	// authorized.GET("/strava-oauth-callback", handler.Make(stravaController.StravaOAuthCallback))

	fmt.Printf("Listening on port %s\n", cfg.HttpListenPort)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.HttpListenPort)))
}
