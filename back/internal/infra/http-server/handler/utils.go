package handler

// import (
// 	"context"
// 	"fmt"
// 	"net/http"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/pttrulez/activitypeople/internal/domain"
// )

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
