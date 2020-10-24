package peloton

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetUserWorkouts returns workouts for a given user
func (c *Client) GetUserWorkouts(userID string, limit int) (Workouts, error) {
	var w Workouts
	data, err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/workouts?limit=%d", UserEndpoint, userID, limit), nil)
	if err != nil {
		return w, err
	}
	err = json.Unmarshal(data, &w)
	if err != nil {
		return w, fmt.Errorf("failed to unmarshal user workouts response, %w", err)
	}

	return w, nil
}

// GetUserWorkoutsCSV returns csv of all workouts for a given user
func (c *Client) GetUserWorkoutsCSV(userID string) ([]byte, error) {
	data, err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/workout_history_csv", UserEndpoint, userID), nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}
