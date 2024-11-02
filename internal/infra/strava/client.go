package strava

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	obalunenko "github.com/obalunenko/strava-api/client"
	"github.com/pttrulez/activitypeople/internal/domain"
)

func PrintJSON(v any) {
	indent, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("PrintJSON err:\n", err)
	}
	fmt.Println(string(indent))
}

func (api *StravaClient) ObaGetAthleteActivities(ctx context.Context, userAccessToken string) ([]domain.ActivityInfo, error) {
	client, err := obalunenko.NewAPIClient(userAccessToken)
	if err != nil {
		return nil, fmt.Errorf("obalunenko.NewAPIClient() err: %w", err)
	}

	activities, err := client.Activities.GetLoggedInAthleteActivities(ctx)
	if err != nil {
		fmt.Println("ObaGetAthleteActivities err:\n", err)
	}
	fmt.Println("ObaGetAthleteActivities printJSON:")

	PrintJSON(activities)
	return nil, nil
}

func (api *StravaClient) GetAthleteActivities(ctx context.Context, userAccessToken string,
) ([]domain.ActivityInfo, error) {
	req, _ := http.NewRequest(http.MethodGet, api.baseURl+"athlete/activities", nil)
	req = req.WithContext(ctx)
	req.Header.Add("Authorization", "Bearer "+userAccessToken)
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetAthleteActivities httpClient.Do(req): \n%s", err)
	}
	var activityData []domain.ActivityInfo
	err = json.NewDecoder(resp.Body).Decode(&activityData)
	if err != nil {
		return nil, fmt.Errorf("GetAthleteActivities json.NewDecoder: \n%s", err)
	}
	return activityData, nil
}

func (api *StravaClient) OAuth(userCode string) (*OAuthResponse, error) {
	// This funcion is used after user was redirected to Strava and authorized our app, we got userCode from callback
	url := fmt.Sprintf("https://www.strava.com/oauth/token?client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code",
		api.appClientId, api.appClientSecret, userCode,
	)

	b := bytes.NewBuffer([]byte(""))
	r, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		return nil, fmt.Errorf("stravaApi OAuth() http.Get(url): \n%s", err)
	}
	defer r.Body.Close()

	// Почему-то не декодится в StravaOAuthResponse
	responseData := &OAuthResponse{}
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

type StravaClient struct {
	appClientId     string
	appClientSecret string
	baseURl         string
	httpClient      *http.Client
	// transport       http.RoundTripper
}

func NewStravaClient(appClientId string, appClientSecret string) *StravaClient {
	return &StravaClient{
		appClientId:     appClientId,
		appClientSecret: appClientSecret,
		baseURl:         "https://www.strava.com/api/v3/",
		httpClient:      http.DefaultClient,
		// transport:       http.DefaultTransport,
	}
}

// func (m *StravaApi) RoundTrip(req *http.Request) (*http.Response, error) {
// 	resp, err := m.transport.RoundTrip(req)
// 	fmt.Println("RoundTrip first req URL:", req.URL.String())
// 	fmt.Println("RoundTrip first response code", resp.StatusCode)

// 	if resp.StatusCode == http.StatusUnauthorized {
// 		fmt.Println("Startting Refresh:")
// 		refreshTokenRequest, _ := http.NewRequest(http.MethodPost, "https://www.strava.com/oauth/token", nil)
// 		params := url.Values{}
// 		params.Add("client_id", m.appClientId)
// 		params.Add("client_secret", m.appClientSecret)
// 		params.Add("refresh_token", m.userRefreshToken)
// 		params.Add("grant_type", "refresh_token")
// 		refreshTokenRequest.URL.RawQuery = params.Encode()

// 		resp, err = m.transport.RoundTrip(refreshTokenRequest)
// 		if err != nil {
// 			return nil, fmt.Errorf("StravaApi RoundTrip refreshTokenRequest err: \n%s", err)
// 		}
// 		fmt.Println("RoundTrip refreshTokenRequest response code", resp.StatusCode)
// 		refreshData := &OAuthRefreshTokenResponse{}
// 		err = json.NewDecoder(resp.Body).Decode(refreshData)
// 		if err != nil {
// 			return nil, fmt.Errorf("StravaApi RoundTrip refreshTokenRequest decode err: \n%s", err)
// 		}
// 		fmt.Printf("RoundTrip refreshData %+v\n", refreshData)

// 		if resp.StatusCode == http.StatusOK {

// 			// Store tokens
// 			m.storeTokens(refreshData.AccessToken, refreshData.RefreshToken)
// 			m.userAccessToken = refreshData.AccessToken
// 			m.userRefreshToken = refreshData.RefreshToken

// 			//Repeat initial Request
// 			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.userAccessToken))

// 			resp, err = m.transport.RoundTrip(req)
// 			fmt.Println("RoundTrip second response code", resp.StatusCode)
// 		} else {
// 			return nil, fmt.Errorf("StravaApi RoundTrip refreshTokenRequest status code: %d", resp.StatusCode)
// 		}
// 	}

// 	return resp, err
// }
