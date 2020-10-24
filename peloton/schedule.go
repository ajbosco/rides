package peloton

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetSchedule returns scheduled workouts for an date range
func (c *Client) GetSchedule(category string, start, end int) (Schedule, error) {
	var s Schedule
	scheduleEndpoint := fmt.Sprintf("/api/v3/ride/live?content_provider=studio&browse_category=%s&exclude_complete=true&start=%v&end=%v", category, start, end)
	data, err := c.doRequest(http.MethodGet, scheduleEndpoint, nil)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return s, fmt.Errorf("failed to unmarshal schedule response, %w", err)
	}

	return s, nil
}

// GetUserSchedule returns authenticated users's schedule
func (c *Client) GetUserSchedule(category string, start, end int) ([]ScheduledWorkout, error) {
	var workouts []ScheduledWorkout
	s, err := c.GetSchedule(category, start, end)
	if err != nil {
		return workouts, nil
	}
	for _, d := range s.Data {
		if d.AuthedUserReservationID != nil {
			workouts = append(workouts, d)
		}
	}
	return workouts, nil
}
