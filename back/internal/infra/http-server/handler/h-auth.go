package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
)

type JwtClaims struct {
	BMR   int
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

	claims := &JwtClaims{
		BMR:   user.BMR,
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	fmt.Println("CLIAM",  claims)
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
		"user":        user.JSON(),
	})
}

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

	claims := &JwtClaims{
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

	// token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{})
	// refreshToken, err := token.SignedString([]byte("moysecret"))
	// if err != nil {
	// 	return err
	// }

	return e.JSON(http.StatusOK, echo.Map{
		"accessToken": accessToken,
		"user":        user.JSON(),
	})
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
