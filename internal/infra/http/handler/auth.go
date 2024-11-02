package handler

import (
	"fmt"
	"net/http"

	model "github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store"
	"github.com/pttrulez/activitypeople/internal/infra/view/auth"

	"golang.org/x/crypto/bcrypt"
)

func (c *AuthController) LoginPage(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, auth.LoginPage())
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) error {
	credentials := model.LoginUserDto{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	valid, errs := Validate(credentials)
	if !valid {
		return render(r, w, auth.LoginForm(credentials, errs))
	}
	user, err := c.userRepo.GetByEmail(r.Context(), credentials.Email)
	if err != nil || user == nil {
		return render(r, w, auth.LoginForm(credentials, map[string]string{
			"Credentials": "invalid credentials",
		}))
	}
	fmt.Printf("(c *AuthController) Login user: %v\n", user)
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword),
		[]byte(credentials.Password))
	if err != nil {
		return render(r, w, auth.LoginForm(credentials, map[string]string{
			"Credentials": "invalid credentials",
		}))
	}
	err = c.sessionStore.SetUserIntoSession(w, r, store.UserSession{
		Id:    user.Id,
		Email: user.Email,
	})
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
	dto := model.RegisterUserDto{
		ConfirmPassword: r.FormValue("confirmPassword"),
		Email:           r.FormValue("email"),
		Name:            r.FormValue("name"),
		Password:        r.FormValue("password"),
		Role:            model.Scoof,
	}

	valid, errs := Validate(dto)
	if !valid {
		return render(r, w, auth.RegisterForm(dto, errs))
	}
	existingUser, err := c.userRepo.GetByEmail(r.Context(), dto.Email)
	if existingUser != nil {
		return render(r, w, auth.RegisterForm(dto, errs))
	} else if err != nil {
		return err
	}

	encpw, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := model.User{
		Email:          dto.Email,
		Name:           dto.Name,
		Role:           model.Scoof,
		HashedPassword: string(encpw),
	}

	id, err := c.userRepo.Insert(r.Context(), &user)
	fmt.Printf("err: %v\n", err)
	if err != nil {
		return render(r, w, auth.RegisterForm(dto, map[string]string{"Credentials": "Что-то пошло не так"}))
	}

	c.sessionStore.SetUserIntoSession(w, r, store.UserSession{Id: id, Email: user.Email})

	return HtmxRedirect(w, r, "/")
}

type AuthController struct {
	sessionStore store.SessionStore
	userRepo     store.UserStore
}

func NewAuthController(repo store.UserStore, sessionStore store.SessionStore) *AuthController {
	return &AuthController{
		sessionStore: sessionStore,
		userRepo:     repo,
	}
}
