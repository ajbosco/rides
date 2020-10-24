package peloton

import (
	"fmt"
	"net/http"
)

// CreateReservation adds workout to user's reservations
func (c *Client) CreateReservation(userID, workoutID string) error {
	req := ScheduleRequest{
		UserID:    userID,
		WorkoutID: workoutID,
	}
	_, err := c.doRequest(http.MethodPost, ReservationEndpoint, req)
	if err != nil {
		return err
	}
	return nil
}

// RemoveReservation removes a workout from a user's reservations
func (c *Client) RemoveReservation(reservationID string) error {
	_, err := c.doRequest(http.MethodDelete, fmt.Sprintf("%s/%s", ReservationEndpoint, reservationID), nil)
	if err != nil {
		return err
	}
	return nil
}
