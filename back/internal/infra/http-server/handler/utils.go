package handler

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/pttrulez/activitypeople/internal/domain"
)

// import (
// 	"context"
// 	"fmt"
// 	"net/http"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/pttrulez/activitypeople/internal/domain"
// )

func GetUserFromContext(c echo.Context) (domain.User, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return domain.User{}, errors.New("Failed to get Token from echo context")
	}

	claims, ok := token.Claims.(JwtClaims)
	if !ok {
		return domain.User{}, errors.New("Failed to cast claims to jwtClaims")
	}

	user := domain.User{
		Email: claims.Email,
		Id:    claims.Id,
		Name:  claims.Name,
		Role:  claims.Role,
	}

	return user, nil
}

// func GetUserFromRequest(r *http.Request) domain.User {
// 	return GetUserFromContext(r.Context())
// }

// func GetUserFromContext(ctx context.Context) domain.User {
// 	user, ok := ctx.Value(userKey).(domain.User)
// 	if !ok {
// 		return domain.User{}
// 	}
// 	return user
// }

// func getValErrMessage(err validator.FieldError) string {
// 	switch err.ActualTag() {
// 	case "required":
// 		return fmt.Sprintf("Поле %s обязательно для заполнения", err.Field())
// 	case "email":
// 		return fmt.Sprintf("Поле %s должно быть валидным email'ом", err.Field())
// 	default:
// 		return err.Error()
// 		// return fmt.Sprintf("Поле %s должно неверно заполнено", err.Field())
// 	}
// }
