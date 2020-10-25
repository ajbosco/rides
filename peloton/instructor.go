package peloton

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetInstructorByID returns instrutor by ID
func (c *Client) GetInstructorByID(instructorID string) (Instructor, error) {
	var i Instructor
	data, err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s", InstructorEndpoint, instructorID), nil)
	if err != nil {
		return i, err
	}
	err = json.Unmarshal(data, &i)
	if err != nil {
		return i, fmt.Errorf("failed to unmarshal instructor response, %w", err)
	}

	return i, nil
}
