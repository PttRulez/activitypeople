package strava

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func (api *Client) GetActivity(ctx context.Context, id int) (
	ActivityResponse, error) {
	req, _ := http.NewRequest(http.MethodGet, api.baseURl+fmt.Sprintf(
		"activities/%d", id), nil)
	req = req.WithContext(ctx)
	req.Header.Add("Authorization", "Bearer "+api.userAccessToken)

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return ActivityResponse{}, fmt.Errorf(
			"GetAthleteActivities httpClient.Do(req): \n%s", err)
	}

	var activityData ActivityResponse
	err = json.NewDecoder(resp.Body).Decode(&activityData)
	if err != nil {
		return ActivityResponse{}, fmt.Errorf("GetAthleteActivities json.NewDecoder: \n%s", err)
	}

	return activityData, nil
}

func (api *Client) GetAthleteActivities(ctx context.Context, after *int64) (
	[]AthleteActivityResponse, error) {
	req, _ := http.NewRequest(http.MethodGet, api.baseURl+"athlete/activities", nil)
	req = req.WithContext(ctx)
	req.Header.Add("Authorization", "Bearer "+api.userAccessToken)

	if after != nil {
		q := req.URL.Query()
		q.Add("after", strconv.FormatInt(*after, 10))
		req.URL.RawQuery = q.Encode()
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetAthleteActivities httpClient.Do(req): \n%s", err)
	}

	var activityData []AthleteActivityResponse
	err = json.NewDecoder(resp.Body).Decode(&activityData)
	if err != nil {
		return nil, fmt.Errorf("GetAthleteActivities json.NewDecoder: \n%s", err)
	}

	return activityData, nil
}

func (api *Client) OAuth(userCode string) (*OAuthResponse, error) {
	// This function is used after user was redirected to Strava and authorized our app, we got userCode from callback
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

type Base struct {
	appClientId     string
	appClientSecret string
	baseURl         string
}

func NewStrava(appClientId string, appClientSecret string) *Base {
	return &Base{
		appClientId:     appClientId,
		appClientSecret: appClientSecret,
		baseURl:         "https://www.strava.com/api/v3/",
	}
}

func (b *Base) NewClient(
	accessToken, refreshToken string,
	storeTokens func(accessToken string, refreshToken string) error,
) *Client {
	client := &Client{
		Base:             *b,
		storeTokens:      storeTokens,
		transport:        http.DefaultTransport,
		userAccessToken:  accessToken,
		userRefreshToken: refreshToken,
	}

	client.httpClient = &http.Client{
		Transport: client,
	}

	return client
}

type Client struct {
	Base
	httpClient       *http.Client
	storeTokens      func(accessToken, refreshToken string) error
	transport        http.RoundTripper
	userAccessToken  string
	userRefreshToken string
}

func (m *Client) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := m.transport.RoundTrip(req)

	if resp.StatusCode == http.StatusUnauthorized {
		refreshTokenRequest, _ := http.NewRequest(http.MethodPost,
			"https://www.strava.com/oauth/token", nil)

		params := url.Values{}
		params.Add("client_id", m.appClientId)
		params.Add("client_secret", m.appClientSecret)
		params.Add("refresh_token", m.userRefreshToken)
		params.Add("grant_type", "refresh_token")
		refreshTokenRequest.URL.RawQuery = params.Encode()

		resp, err = m.transport.RoundTrip(refreshTokenRequest)
		if err != nil {
			return nil, fmt.Errorf("StravaApi RoundTrip refreshTokenRequest err: \n%s",
				err)
		}

		refreshData := &OAuthRefreshTokenResponse{}
		err = json.NewDecoder(resp.Body).Decode(refreshData)
		if err != nil {
			return nil, fmt.Errorf("StravaApi RoundTrip refreshTokenRequest decode err: \n%s",
				err)
		}

		if resp.StatusCode == http.StatusOK {
			// Store tokens
			err := m.storeTokens(refreshData.AccessToken, refreshData.RefreshToken)
			if err != nil {
				return nil, fmt.Errorf(
					"StravaApi RoundTrip refreshTokenRequest storeTokens err: \n%s",
					err)
			}
			m.userAccessToken = refreshData.AccessToken
			m.userRefreshToken = refreshData.RefreshToken

			//Repeat initial Request
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.userAccessToken))

			resp, err = m.transport.RoundTrip(req)
			if err != nil {
				return nil, fmt.Errorf("StravaApi RoundTrip second request err: \n%s",
					err)
			}
		} else {
			return nil, fmt.Errorf("StravaApi RoundTrip refreshTokenRequest status code: %d",
				resp.StatusCode)
		}
	}

	return resp, err
}
