package handler

// import (
// 	"context"

// 	"github.com/labstack/echo/v4"
// 	"github.com/pttrulez/activitypeople/internal/domain"
// )

// // import (
// // 	"context"
// // 	"fmt"
// // 	"net/http"

// // 	"github.com/pttrulez/activitypeople/internal/domain"
// // 	"github.com/pttrulez/activitypeople/internal/infra/view/pages/home"
// // )





// type StravaController struct {
// 	stravaService StravaService
// }

// type StravaService interface {
// 	OAuthStrava(ctx context.Context, userCode string, userID int) (string, string, error)
// 	SyncActivities(ctx context.Context, user domain.User) error
// }

// func NewStravaController(
// 	stravaService StravaService,
// ) *StravaController {
// 	return &StravaController{
// 		stravaService: stravaService,
// 	}
// }
