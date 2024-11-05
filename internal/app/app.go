package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pttrulez/activitypeople/internal/config"
	"github.com/pttrulez/activitypeople/internal/infra/http_server/handler"
	"github.com/pttrulez/activitypeople/internal/infra/store/pgstore"
	"github.com/pttrulez/activitypeople/internal/infra/store/session"
	"github.com/pttrulez/activitypeople/internal/infra/strava"
	"github.com/pttrulez/activitypeople/internal/service/activities"
	"github.com/pttrulez/activitypeople/internal/service/auth"

	"github.com/go-chi/chi/v5"
)

func StartServer() {
	cfg := config.MustLoadConfig()

	// stores
	sessionStore := session.NewGorillaCookiesSessionsStore(
		[]byte(cfg.SessionSecret))
	pgConn := pgstore.CreatePGConnection(cfg.Postgres)
	stravaStore := pgstore.NewStravaPostgres(pgConn)
	userStore := pgstore.NewUserPostgres(pgConn)

	// clients
	strava := strava.NewStrava(cfg.Strava.ClientID, cfg.Strava.ClientSecret)

	// services
	authService := auth.NewService(userStore)
	activitiesService := activities.NewService(strava, stravaStore)

	// Routing
	router := chi.NewMux()
	router.Use(handler.AddUserToContextMiddleware(sessionStore))

	// static files
	router.Handle("/public/*", http.StripPrefix("/public",
		http.FileServer(http.Dir("./public"))))

	// controllers
	authController := handler.NewAuthController(authService, sessionStore)
	homeCocntroller := handler.NewHomeController(cfg.Strava.OAuthLink)
	activitiesController := handler.NewActivitiesController(activitiesService)
	stravaController := handler.NewStravaController(activitiesService)

	// handlers
	router.Get("/", handler.Make(homeCocntroller.HandlerHomeIndex))
	router.Get("/register", handler.Make(authController.RegisterPage))
	router.Post("/register", handler.Make(authController.Register))
	router.Get("/login", handler.Make(authController.LoginPage))
	router.Post("/login", handler.Make(authController.Login))

	router.Group(func(authorized chi.Router) {
		authorized.Use(handler.OnlyAuthenticatedMiddleware)
		authorized.Get("/activities", handler.Make(activitiesController.GetActivitiesPage))
		authorized.Get("/strava-oauth-callback", handler.Make(stravaController.StravaOAuthCallback))
	})

	fmt.Printf("Listening on port %s\n", cfg.HttpListenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpListenPort), router))
}
