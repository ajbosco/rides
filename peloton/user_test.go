package peloton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	c := NewClient(testUser, testPassword)

	c.baseURL = mockResponse("user.json").URL

	user, err := c.GetUser()
	assert.NoError(t, err)

	assert.Equal(t, "Test", user.FirstName)
}
