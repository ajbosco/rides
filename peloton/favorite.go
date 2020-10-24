package peloton

import (
	"fmt"
	"net/http"
)

// CreateFavorite adds workout to authenticated user's favorites
func (c *Client) CreateFavorite(rideID string) error {
	req := FavoritesRequest{
		RideID: rideID,
	}
	_, err := c.doRequest(http.MethodPost, fmt.Sprintf("%s/create", FavoritesEndpoint), req)
	if err != nil {
		return err
	}
	return nil
}

// RemoveFavorite removes workout from authenticated user's favorites
func (c *Client) RemoveFavorite(rideID string) error {
	req := FavoritesRequest{
		RideID: rideID,
	}
	_, err := c.doRequest(http.MethodPost, fmt.Sprintf("%s/delete", FavoritesEndpoint), req)
	if err != nil {
		return err
	}
	return nil
}
