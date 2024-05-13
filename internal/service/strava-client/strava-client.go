package stravaclient

import (
	"antiscoof/internal/model"
	models "antiscoof/internal/service/strava-client/stravamodels"
	"antiscoof/internal/store"
	"antiscoof/internal/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	strava "github.com/obalunenko/strava-api/client"
)

func PrintJSON(v any) {
	indent, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("PrintJSON err:\n", err)
	}
	fmt.Println(string(indent))
}

func (api *StravaApi) ObaGetAthleteActivities() {
	activities, err := api.obalunenko.Activities.GetLoggedInAthleteActivities(context.TODO())
	if err != nil {
		fmt.Println("ObaGetAthleteActivities err:\n", err)
	}
	fmt.Println("ObaGetAthleteActivities printJSON:")
	PrintJSON(activities)
}

func (api *StravaApi) GetAthleteActivities() ([]models.ActivityInfo, error) {
	req, _ := http.NewRequest(http.MethodGet, api.baseURl+"athlete/activities", nil)
	req.Header.Add("Authorization", "Bearer "+api.userAccessToken)
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetAthleteActivities httpClient.Do(req): \n%s", err)
	}
	var activityData []models.ActivityInfo
	err = json.NewDecoder(resp.Body).Decode(&activityData)
	if err != nil {
		return nil, fmt.Errorf("GetAthleteActivities json.NewDecoder: \n%s", err)
	}
	return activityData, nil
}

func (api *StravaApi) OAuth(userCode string) (*models.OAuthResponse, error) {
	// This funcion is used after user was redirected to Strava and authorized our app, we got userCode from callback
	url := fmt.Sprintf(`https://www.strava.com/oauth/token?
		client_id=%s
		&client_secret=%s
		&code=%s
		&grant_type=authorization_code
		`,
		api.appClientId, api.appClientSecret, userCode,
	)

	b := bytes.NewBuffer([]byte(""))
	r, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		return nil, fmt.Errorf("stravaApi OAuth() http.Get(url): \n%s", err)
	}
	defer r.Body.Close()

	// Почему-то не декодится в StravaOAuthResponse
	responseData := &models.OAuthResponse{}
	m := make(map[string]any)
	err = json.NewDecoder(r.Body).Decode(&m)
	// err = json.NewDecoder(r.Body).Decode(responseData)
	if err != nil {
		return nil, fmt.Errorf("stravaApi OAuth() json.NewDecoder: \n%s", err)
	}

	responseData.AccessToken = m["access_token"].(string)
	responseData.RefreshToken = m["refresh_token"].(string)
	return responseData, nil
}

func (m *StravaApi) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := m.transport.RoundTrip(req)
	fmt.Println("RoundTrip first req URL:", req.URL.String())
	fmt.Println("RoundTrip first response code", resp.StatusCode)

	if resp.StatusCode == http.StatusUnauthorized {
		refreshTokenRequest, _ := http.NewRequest(http.MethodPost, "https://www.strava.com/oauth/token", nil)
		params := url.Values{}
		params.Add("client_id", m.appClientId)
		params.Add("client_secret", m.appClientSecret)
		params.Add("refresh_token", m.userRefreshToken)
		params.Add("grant_type", "refresh_token")
		refreshTokenRequest.URL.RawQuery = params.Encode()

		resp, err = m.transport.RoundTrip(refreshTokenRequest)

		refreshData := &models.OAuthRefreshTokenResponse{}
		json.NewDecoder(resp.Body).Decode(refreshData)
		if err != nil {
			return nil, fmt.Errorf("StravaApi RoundTrip: \n%s", err)
		}
		fmt.Printf("RoundTrip refreshData %+v\n", refreshData)

		// Store tokens
		m.storeTokens(refreshData.AccessToken, refreshData.RefreshToken)
		m.userAccessToken = refreshData.AccessToken
		m.userRefreshToken = refreshData.RefreshToken

		//Repeat initial Request
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.userAccessToken))

		resp, err = m.transport.RoundTrip(req)
		fmt.Println("RoundTrip second response code", resp.StatusCode)
	}

	return resp, err
}

type StravaApi struct {
	appClientId      string
	appClientSecret  string
	baseURl          string
	httpClient       *http.Client
	obalunenko       *strava.APIClient
	storeTokens      func(accessToken, refreshToken string) error
	transport        http.RoundTripper
	userAccessToken  string
	userRefreshToken string
}

type CreateStravaApiParams struct {
	AppClientId      string
	AppClientSecret  string
	StoreTokens      func(accessToken, refreshToken string) error
	UserAccessToken  string
	UserRefreshToken string
}

func NewStravaApi(params CreateStravaApiParams) *StravaApi {
	client, err := strava.NewAPIClient(params.UserAccessToken)
	if err != nil {
		fmt.Println("Failed to start obalunenko strava:", err)
	}

	stravaApi := &StravaApi{
		appClientId:      params.AppClientId,
		appClientSecret:  params.AppClientSecret,
		baseURl:          "https://www.strava.com/api/v3/",
		obalunenko:       client,
		storeTokens:      params.StoreTokens,
		transport:        http.DefaultTransport,
		userAccessToken:  params.UserAccessToken,
		userRefreshToken: params.UserRefreshToken,
	}
	stravaApi.httpClient = &http.Client{
		Transport: stravaApi,
	}

	return stravaApi
}

func NewStravaApiFromRequest(r *http.Request, appClientID string, appClientSecret string, stravaStore store.StravaStore) *StravaApi {
	user := utils.GetUserFromContext(r)
	fmt.Printf("user: %v, %v\n", *user.Strava.AccessToken, *user.Strava.RefreshToken)
	stravaApi := NewStravaApi(CreateStravaApiParams{
		AppClientId:     appClientID,
		AppClientSecret: appClientSecret,
		StoreTokens: func(accessToken string, refreshToken string) error {
			return stravaStore.UpdateUserStravaInfo(r.Context(), &model.UpdateStravaTokens{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
				UserId:       user.Id,
			})
		},
		UserAccessToken:  *user.Strava.AccessToken,
		UserRefreshToken: *user.Strava.RefreshToken,
	})

	return stravaApi
}
