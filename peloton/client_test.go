package peloton

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

const (
	testUser     = "test-user"
	testPassword = "test-password"
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient(testUser, testPassword)
	client.baseURL = server.URL
}

func teardown() {
	server.Close()
}

func testClientDefaultBaseURL(t *testing.T, c *Client) {
	if c.baseURL == "" || c.baseURL != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, expected %v", c.baseURL, defaultBaseURL)
	}
}

func Test_NewClient(t *testing.T) {
	c := NewClient(testUser, testPassword)
	testClientDefaultBaseURL(t, c)
}

func Test_doRequest(t *testing.T) {
	setup()
	defer teardown()

	testData := `{"testing":"things"}`

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, testData)
	})

	actual, err := client.doRequest(http.MethodGet, "/test", nil)
	assert.NoError(t, err)

	expected := []byte(testData)
	assert.Equal(t, expected, actual)
}

func Test_doRequest_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	_, err := client.doRequest(http.MethodGet, "/", nil)
	assert.Error(t, err)
}

func mockResponse(paths ...string) *httptest.Server {
	parts := []string{".", "testdata"}
	filename := filepath.Join(append(parts, paths...)...)

	mockData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write(mockData)
	}))
}
