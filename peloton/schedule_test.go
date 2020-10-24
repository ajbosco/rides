package peloton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSchedule(t *testing.T) {
	c := NewClient(testUser, testPassword)

	c.baseURL = mockResponse("schedule.json").URL

	schedule, err := c.GetSchedule("cycling", 0, 1603578228)
	assert.NoError(t, err)

	assert.Equal(t, 1, schedule.Count)
}

func TestGetUserSchedule(t *testing.T) {
	c := NewClient(testUser, testPassword)

	c.baseURL = mockResponse("schedule.json").URL

	scheduledWorkouts, err := c.GetUserSchedule("cycling", 0, 1603578228)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(scheduledWorkouts))
}
