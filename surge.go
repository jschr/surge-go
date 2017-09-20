package surge

import (
	"fmt"
	"net/http"
)

const (
	// BaseURL is the default surge api base url
	BaseURL       = "https://surge.surge.sh"
	tokenEndpoint = "token"
)

// Surge defines the surge client api methods
type Surge interface {
	Token() string
	SetToken(token string)
	Login(username, password string) (string, error)
	List() ([]string, error)
	// Publish(projectPath, domain string) error
	Teardown(domain string) error
	getURL(endpoint string) string
}

// Client interacts with the surge api
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// APIError is an error returned by the surge API
type APIError struct {
	Messages []string
}

func (e *APIError) Error() string {
	return e.Messages[0]
}

// NewSurge creates a new surge client
func NewSurge(token string) Surge {
	s := Client{
		token:      token,
		baseURL:    BaseURL,
		httpClient: http.DefaultClient,
	}

	return &s
}

// Token returns the token used to authenticate to the api
func (s *Client) Token() string {
	return s.token
}

// SetToken sets the token
func (s *Client) SetToken(token string) {
	s.token = token
}

// func (s *Client) makeRequest(method, endpoint, username, password string, body io.Reader, resData, resError interface{}) error {
// 	req, err := http.NewRequest(method, s.getURL(endpoint), body)
// 	if err != nil {
// 		return err
// 	}

// 	req.SetBasicAuth(username, password)

// 	res, err := s.httpClient.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	defer res.Body.Close()
// 	decoder := json.NewDecoder(res.Body)
// 	if res.StatusCode >= 400 {
// 		err = decoder.Decode(&loginError)
// 		if err != nil {
// 			return err
// 		}
// 		return &loginError
// 	}

// 	err = decoder.Decode(&loginSuccess)
// 	if err != nil {
// 		return err
// 	}
// }

func (s *Client) getURL(endpoint string) string {
	return fmt.Sprintf("%s/%s", s.baseURL, endpoint)
}
