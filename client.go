package binocs

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const BaseURL = "https://api.binocs.sh"
const UserAgent = "BinocsClientGo"
const Version = "v0.1.0"

type Client struct {
	Checks  CheckService
	Channel ChannelService
}

type ClientConfig struct {
	AccessKey string
	SecretKey string
}

func New(config ClientConfig) (*Client, error) {
	r := resty.New()
	c := &Client{
		Checks:  CheckService{resty: r},
		Channel: ChannelService{resty: r},
	}
	r.SetBaseURL(BaseURL)
	r.SetHeader("Content-Type", "application/json")
	r.SetHeader("User-Agent", fmt.Sprintf("%s %s", UserAgent, Version))
	token, err := getAccessToken(config.AccessKey, config.SecretKey, r)
	if err != nil {
		return c, err
	}
	r.SetAuthToken(token)
	return c, nil
}

type AuthenticationResponse struct {
	AccessToken string `json:"access_token"`
}

func getAccessToken(accessKey, secretKey string, r *resty.Client) (string, error) {
	var result AuthenticationResponse
	resp, err := r.R().
		SetBody(fmt.Sprintf(`{"access_key":"%s", "secret_key":"%s"}`, accessKey, secretKey)).
		ForceContentType("application/json").
		SetResult(&result).
		Post("/authenticate")
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return result.AccessToken, nil
}
