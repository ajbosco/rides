package peloton

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"golang.org/x/net/publicsuffix"
)

const (
	defaultBaseURL = "https://api.onepeloton.com"
	mediaType      = "application/json"
)

// Client manages communication with Peloton API.
type Client struct {
	baseURL  string
	user     string
	password string
	client   *http.Client
}

// NewClient creates a new Peloton API client.
func NewClient(user string, password string) *Client {
	return &Client{
		baseURL:  defaultBaseURL,
		user:     user,
		password: password,
		client:   http.DefaultClient,
	}
}

// Authenticate authenticates the client using username and password
func (c *Client) Authenticate() error {
	req := auth{
		User:     c.user,
		Password: c.password,
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}
	c.client.Jar = jar

	_, err = c.doRequest(http.MethodPost, AuthEndpoint, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) doRequest(method, endpoint string, data interface{}) ([]byte, error) {
	// Encode data if we are passed an object.
	b := bytes.NewBuffer(nil)
	if data != nil {
		// Create the encoder.
		enc := json.NewEncoder(b)
		if err := enc.Encode(data); err != nil {
			return nil, fmt.Errorf("failed to encode json, %w", err)
		}
	}

	// Create the request.
	uri := fmt.Sprintf("%s/%s", c.baseURL, strings.Trim(endpoint, "/"))
	fmt.Println(uri)
	req, err := http.NewRequest(method, uri, b)
	if err != nil {
		return nil, fmt.Errorf("creating %s request to %s failed, %w", method, uri, err)
	}

	// Set the proper headers.
	req.Header.Set("Content-Type", mediaType)

	// Do the request.
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing %s request to %s failed, %w", method, uri, err)
	}
	defer resp.Body.Close()

	// Check that the response status code was OK.
	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("unauthorized access to endpoint")
	case http.StatusForbidden:
		return nil, fmt.Errorf("unauthorized access to endpoint")
	case http.StatusNotFound:
		return nil, fmt.Errorf("the requested uri does not exist")
	case http.StatusBadRequest:
		return nil, fmt.Errorf("the request is invalid")
	default:
		return nil, fmt.Errorf("bad response code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("decoding response from %s request to %s failed: body -> %s\n, %w", method, uri, string(body), err)
	}

	return body, nil
}
