package app

import (
	"antiscoof/internal/config"
	"antiscoof/internal/handler"
	"antiscoof/internal/store"
	"antiscoof/internal/store/pgstore"
	"antiscoof/internal/store/session"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func StartServer() {
	cfg := config.MustLoadConfig()

	router := chi.NewMux()
	var sessionStore store.SessionStore = session.NewGorillaCookiesSessionsStore(
		[]byte(cfg.SessionSecret), cfg.UserSessionKey)
	pgConn := pgstore.CreatePGConnection(cfg.Postgres)
	// stravaStore := pgstore.NewStravaPostgres(pgConn)
	// stravaApi := stravaclient.NewStravaApi(cfg)
	stravaStore := pgstore.NewStravaPostgres(pgConn)
	userStore := pgstore.NewUserPostgres(pgConn)

	router.Use(handler.AddUserToContext(sessionStore, userStore))

	router.Handle("/public/*", http.StripPrefix("/public",
		http.FileServer(http.Dir("./public"))))

	authController := handler.NewAuthController(userStore, sessionStore)
	homeCocntroller := handler.NewHomeController(cfg.Strava.OAuthLink)
	activitiesController := handler.NewActivitiesController(cfg, stravaStore)
	// stravaCocntroller := handler.NewStravaController(stravaApi, stravaStore)

	// Handlers
	router.Get("/", handler.Make(homeCocntroller.HandlerHomeIndex))
	router.Get("/register", handler.Make(authController.RegisterPage))
	router.Post("/register", handler.Make(authController.Register))
	router.Get("/login", handler.Make(authController.LoginPage))
	router.Post("/login", handler.Make(authController.Login))

	router.Group(func(authorized chi.Router) {
		authorized.Use(handler.OnlyAuthenticated)
		authorized.Get("/activities", handler.Make(activitiesController.GetActivitiesPage))
		// authorized.Get("/strava-oauth-callback", handler.Make(stravaCocntroller.StravaOAuthCallback))
	})

	fmt.Printf("Listening on port %s\n", cfg.HttpListenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpListenPort), router))
}
