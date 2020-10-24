package peloton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserWorkouts(t *testing.T) {
	c := NewClient(testUser, testPassword)

	c.baseURL = mockResponse("workouts.json").URL

	workouts, err := c.GetUserWorkouts("1234567890", 10)
	assert.NoError(t, err)

	assert.Equal(t, 10, workouts.Count)
}
