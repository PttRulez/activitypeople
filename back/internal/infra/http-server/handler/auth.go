package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
)

type jwtClaims struct {
	Id    int
	Name  string
	Email string
	Role  domain.Role
	jwt.RegisteredClaims
}

func (c *AuthController) Login(e echo.Context) error {
	var req contracts.LoginUserRequest
	err := e.Bind(&req)
	if err != nil {
		return err
	}
	err = ValidateStruct(req)
	if err != nil {
		return err
	}

	user, err := c.authService.Login(e.Request().Context(), req.Email, req.Password)
	if err != nil {
		return err
	}

	claims := &jwtClaims{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(c.jwtSecret))
	if err != nil {
		return err
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{})
	refreshToken, err := token.SignedString([]byte(c.jwtSecret))
	if err != nil {
		return err
	}

	e.SetCookie(&http.Cookie{
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   24 * 60 * 60,
		Name:     "activity-refresh",
		Secure:   true,
		Value:    refreshToken,
	})

	return e.JSON(http.StatusOK, echo.Map{
		"accessToken": accessToken,
		// "refreshToken": refreshToken,
	})
}

// func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) error {
// 	err := c.sessionStore.ClearUserSession(w, r)
// 	if err != nil {
// 		return err
// 	}
// 	return HtmxRedirect(w, r, "/")
// }

func (c *AuthController) Register(e echo.Context) error {
	var r contracts.RegisterUserRequest
	err := e.Bind(&r)
	if err != nil {
		return err
	}
	err = ValidateStruct(r)
	if err != nil {
		return err
	}

	user, err := c.authService.Register(e.Request().Context(),
		r.Email, r.Password, r.Name)
	if err != nil {
		return err
	}

	claims := &jwtClaims{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte("moysecret"))
	if err != nil {
		return err
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{})
	refreshToken, err := token.SignedString([]byte("moysecret"))
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, echo.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

// func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) error {
// 	req := contracts.RegisterUserRequest{
// 		ConfirmPassword: r.FormValue("confirmPassword"),
// 		Email:           r.FormValue("email"),
// 		Name:            r.FormValue("name"),
// 		Password:        r.FormValue("password"),
// 		Role:            domain.Scoof,
// 	}
// 	valid, errs := Validate(req)
// 	if !valid {
// 		return render(r, w, auth.RegisterForm(req, errs))
// 	}
// 	newUser, err := c.authService.Register(r.Context(), req.Email, req.Password, req.Name)
// 	if errors.Is(err, service.ErrAlreadyExists) {
// 		return render(r, w, auth.RegisterForm(req, map[string]string{
// 			"Email": "email already exists",
// 		}))
// 	}
// 	if err != nil {
// 		return render(r, w, auth.RegisterForm(req, map[string]string{
// 			"Email": err.Error(),
// 		}))
// 	}
// 	c.sessionStore.SetUserIntoSession(w, r, newUser)
// 	return HtmxRedirect(w, r, "/")
// }

type AuthService interface {
	Login(ctx context.Context, email, password string) (domain.User, error)
	Register(ctx context.Context, email, password, name string) (domain.User, error)
}

type AuthController struct {
	authService AuthService
	jwtSecret   string
}

func NewAuthController(authService AuthService, jwtSecret string) *AuthController {
	return &AuthController{
		authService: authService,
		jwtSecret:   jwtSecret,
	}
}