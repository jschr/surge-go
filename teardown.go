package surge

import (
	"encoding/json"
	"net/http"
)

// Teardown will delete the project for the domain
func (s *Client) Teardown(domain string) error {
	var teardownError APIError

	token := s.Token()
	req, err := http.NewRequest("DELETE", s.getURL(domain), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth("token", token)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode >= 400 {
		err = decoder.Decode(&teardownError)
		if err != nil {
			return err
		}
		return &teardownError
	}

	return nil
}
