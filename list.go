package surge

import (
	"encoding/json"
	"net/http"
)

// List returns all the domains for the authenticated user
func (s *Client) List() ([]string, error) {
	var listSuccess []string
	var listError APIError

	token := s.Token()
	req, err := http.NewRequest("GET", s.getURL("list"), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("token", token)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode >= 400 {
		err = decoder.Decode(&listError)
		if err != nil {
			return nil, err
		}
		return nil, &listError
	}

	err = decoder.Decode(&listSuccess)
	if err != nil {
		return nil, err
	}

	return listSuccess, nil
}
