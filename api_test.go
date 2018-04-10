package apidemic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/pmylund/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDynamicEndpointFailsWithoutRegistration(t *testing.T) {
	s := setUp()
	payload := registerPayload(t, "fixtures/sample_request.json")

	w := httptest.NewRecorder()
	req := jsonRequest("POST", "/api/test", payload)
	s.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDynamicEndpointWithGetRequest(t *testing.T) {
	s := setUp()
	payload := registerPayload(t, "fixtures/sample_request.json")

	w := httptest.NewRecorder()
	req := jsonRequest("POST", "/register", payload)
	s.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	req = jsonRequest("GET", "/api/test", nil)
	s.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDynamicEndpointWithPostRequest(t *testing.T) {
	s := setUp()
	payload := registerPayload(t, "fixtures/sample_request.json")
	payload["http_method"] = "POST"

	w := httptest.NewRecorder()
	req := jsonRequest("POST", "/register", payload)
	s.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	req = jsonRequest("POST", "/api/test", nil)

	s.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDynamicEndpointWithForbiddenResponse(t *testing.T) {
	s := setUp()
	registerPayload := registerPayload(t, "fixtures/sample_request.json")
	registerPayload["response_code_probabilities"] = map[string]int{"403": 100}

	w := httptest.NewRecorder()
	req := jsonRequest("POST", "/register", registerPayload)
	s.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	req = jsonRequest("GET", "/api/test", nil)

	s.ServeHTTP(w, req)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func setUp() *mux.Router {
	store = cache.New(5*time.Minute, 30*time.Second)

	return NewServer()
}

func registerPayload(t *testing.T, fixtureFile string) map[string]interface{} {
	content, err := ioutil.ReadFile(fixtureFile)
	if err != nil {
		t.Fatal(err)
	}

	var api map[string]interface{}
	err = json.NewDecoder(bytes.NewReader(content)).Decode(&api)
	if err != nil {
		t.Fatal(err)
	}

	return api
}

func jsonRequest(method string, path string, body interface{}) *http.Request {
	var bEnd io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil
		}
		bEnd = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, path, bEnd)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}

func TestGenBool(t *testing.T) {
	fixSeed(0)
	b := genBool()

	assert.True(t, b)
}

func TestGenDecimal(t *testing.T) {
	var fixtures = []struct {
		seed            int64
		min             int32
		max             int32
		precision       int
		expectedToMatch string
	}{
		{0, 1, 1, 2, "^1.\\d{2}$"},
		{0, 2, 4, 2, "^2.\\d{2}$"},
		{0, -2, 4, 2, "^\\-2.\\d{2}$"},
		{0, -20, -10, 2, "^\\-16.\\d{2}$"},
		{13, -20, -10, 5, "^\\-18.\\d{5}$"},
	}

	for _, tt := range fixtures {
		fixSeed(tt.seed)

		actual := genDecimal(tt.min, tt.max, tt.precision)

		assert.Regexp(
			t,
			regexp.MustCompile(tt.expectedToMatch),
			actual,
			fmt.Sprintf(
				"genDecimal(%d, %d, %d): expected to match %s, actual %s",
				tt.min,
				tt.max,
				tt.precision,
				tt.expectedToMatch,
				actual,
			),
		)
	}
}

func TestInt(t *testing.T) {
	var fixtures = []struct {
		seed     int64
		min      int32
		max      int32
		expected int32
	}{
		{0, 1, 1, 1},
		{0, 2, 4, 2},
		{0, -2, 4, -2},
		{0, -20, -10, -16},
		{13, -20, -10, -18},
	}

	for _, tt := range fixtures {
		fixSeed(tt.seed)

		actual := genInt(tt.min, tt.max)
		if actual != tt.expected {
			t.Errorf("genInt(%d, %d): expected %d, actual %d", tt.min, tt.max, tt.expected, actual)
		}
	}
}

func TestFloat(t *testing.T) {
	var fixtures = []struct {
		seed     int64
		min      int32
		max      int32
		expected float64
	}{
		{0, 1, 1, 1.0},
		{0, 2, 4, 3.890392},
		{0, -2, 4, 3.671177},
		{0, -20, -10, -10.548039},
		{13, -20, -10, -17.975146},
	}

	for _, tt := range fixtures {
		fixSeed(tt.seed)

		actual := genFloat(tt.min, tt.max)

		assert.InDelta(t, tt.expected, actual, 0.0001)
	}
}

func fixSeed(seed int64) {
	seededRand = rand.New(rand.NewSource(seed))
}
