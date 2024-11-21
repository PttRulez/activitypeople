package handler

// import (
// 	"context"
// 	"fmt"
// 	"net/http"

// 	"github.com/pttrulez/activitypeople/internal/domain"
// 	"github.com/pttrulez/activitypeople/internal/infra/view/pages/home"
// )

// func (c *StravaController) StravaOAuthCallback(w http.ResponseWriter, r *http.Request) error {
// 	code := r.URL.Query().Get("code")
// 	if code == "" {
// 		return HtmxRedirect(w, r, "/")
// 	}

// 	user := GetUserFromRequest(r)
// 	aToken, rToken, err := c.stravaService.OAuthStrava(r.Context(), code, user.Id)
// 	if err != nil {
// 		fmt.Printf("StravaOAuthCallback error: %s\n", err)
// 		return HtmxRedirect(w, r, "/")
// 	}

// 	user.Strava.AccessToken = &aToken
// 	user.Strava.RefreshToken = &rToken

// 	err = c.sessionStore.SetUserIntoSession(w, r, user)
// 	if err != nil {
// 		return render(r, w, home.Index("", user))
// 	}

// 	return HtmxRedirect(w, r, "/")
// }

// func (c *StravaController) SyncStrava(w http.ResponseWriter, r *http.Request) error {
// 	user := GetUserFromRequest(r)
// 	err := c.stravaService.SyncActivities(r.Context(), user)
// 	if err != nil {
// 		fmt.Printf("SyncStrava error: %s\n", err)
// 	}

// 	return HtmxRedirect(w, r, "/")
// }

// type StravaController struct {
// 	sessionStore  SessionStore
// 	stravaService StravaService
// }

// type StravaService interface {
// 	OAuthStrava(ctx context.Context, userCode string, userID int) (string, string, error)
// 	SyncActivities(ctx context.Context, user domain.User) error
// }

// func NewStravaController(
// 	sessionStore SessionStore,
// 	stravaService StravaService,
// ) *StravaController {
// 	return &StravaController{
// 		sessionStore:  sessionStore,
// 		stravaService: stravaService,
// 	}
// }
