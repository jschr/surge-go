package surge

import (
	"encoding/json"
	"net/http"
)

// LoginSuccess is the successful response of a login request
type LoginSuccess struct {
	Email string
	Token string
}

// Login returns an auth token for a valid email and password
func (s *Client) Login(username, password string) (string, error) {
	var loginSuccess LoginSuccess
	var loginError APIError

	req, err := http.NewRequest("POST", s.getURL("token"), nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(username, password)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if res.StatusCode >= 400 {
		err = decoder.Decode(&loginError)
		if err != nil {
			return "", err
		}
		return "", &loginError
	}

	err = decoder.Decode(&loginSuccess)
	if err != nil {
		return "", err
	}

	return loginSuccess.Token, nil
}
