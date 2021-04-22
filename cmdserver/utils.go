package cmdserver

import "encoding/json"

/**
 * Parse API request body into struct
 */
func ParseApiBody(body string) (*ApiRequest, error) {

	// create new request object
	req := ApiRequest{}

	// decode json
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}

	// return object
	return &req, nil
}
