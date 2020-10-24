package peloton

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetUser returns data about current user
func (c *Client) GetUser() (User, error) {
	var u User
	data, err := c.doRequest(http.MethodGet, MeEndpoint, nil)
	if err != nil {
		return u, err
	}
	err = json.Unmarshal(data, &u)
	if err != nil {
		return u, fmt.Errorf("failed to unmarshal user response, %w", err)
	}

	return u, nil
}
