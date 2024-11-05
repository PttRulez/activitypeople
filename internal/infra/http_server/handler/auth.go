package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http_server/contracts"
	"github.com/pttrulez/activitypeople/internal/infra/view/auth"
	"github.com/pttrulez/activitypeople/internal/service"
)

func (c *AuthController) LoginPage(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, auth.LoginPage())
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) error {
	credentials := contracts.LoginUserRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	valid, errs := Validate(credentials)
	if !valid {
		return render(r, w, auth.LoginForm(credentials, errs))
	}

	user, err := c.authService.Login(r.Context(), credentials.Email, credentials.Password)
	if err != nil {
		return render(r, w, auth.LoginForm(credentials, map[string]string{
			"Credentials": "invalid credentials",
		}))
	}

	err = c.sessionStore.SetUserIntoSession(w, r, user)
	if err != nil {
		return render(r, w, auth.LoginForm(credentials, map[string]string{
			"Credentials": "invalid credentials",
		}))
	}
	return HtmxRedirect(w, r, "/")
}

func (c *AuthController) RegisterPage(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, auth.RegisterPage())
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) error {
	req := contracts.RegisterUserRequest{
		ConfirmPassword: r.FormValue("confirmPassword"),
		Email:           r.FormValue("email"),
		Name:            r.FormValue("name"),
		Password:        r.FormValue("password"),
		Role:            domain.Scoof,
	}

	valid, errs := Validate(req)
	if !valid {
		return render(r, w, auth.RegisterForm(req, errs))
	}

	newUser, err := c.authService.Register(r.Context(), req.Email, req.Password, req.Name)
	if errors.Is(err, service.ErrAlreadyExists) {
		return render(r, w, auth.RegisterForm(req, map[string]string{
			"Email": "email already exists",
		}))
	}
	if err != nil {
		return render(r, w, auth.RegisterForm(req, map[string]string{
			"Email": err.Error(),
		}))
	}

	c.sessionStore.SetUserIntoSession(w, r, newUser)

	return HtmxRedirect(w, r, "/")
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (domain.User, error)
	Register(ctx context.Context, email, password, name string) (domain.User, error)
}

type AuthController struct {
	authService  AuthService
	sessionStore SessionStore
}

func NewAuthController(authService AuthService, sessionStore SessionStore) *AuthController {
	return &AuthController{
		authService:  authService,
		sessionStore: sessionStore,
	}
}
