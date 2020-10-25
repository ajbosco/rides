package peloton

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocarina/gocsv"
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
func (c *Client) GetUserWorkoutsCSV(userID string) ([]WorkoutCSV, error) {
	w := []WorkoutCSV{}
	data, err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/workout_history_csv", UserEndpoint, userID), nil)
	if err != nil {
		return w, err
	}
	reader := csv.NewReader(bytes.NewBuffer(data))
	err = gocsv.UnmarshalCSV(reader, &w)
	if err != nil {
		return w, err
	}
	return w, nil
}
