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

	"github.com/stretchr/testify/assert"
)

func TestDynamicEndpointFailsWithoutRegistration(t *testing.T) {
	s := NewServer()
	sample, err := ioutil.ReadFile("fixtures/sample_post_request.json")
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]interface{}
	err = json.NewDecoder(bytes.NewReader(sample)).Decode(&out)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := jsonRequest("POST", "/api/test", out)
	s.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestDynamicEndpointWithGetRequest(t *testing.T) {
	s := NewServer()
	sample, err := ioutil.ReadFile("fixtures/sample_request.json")
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]interface{}
	err = json.NewDecoder(bytes.NewReader(sample)).Decode(&out)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := jsonRequest("POST", "/register", out)
	s.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req = jsonRequest("GET", "/api/test", nil)
	s.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestDynamicEndpointWithPostRequest(t *testing.T) {
	s := NewServer()
	sample, err := ioutil.ReadFile("fixtures/sample_post_request.json")
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]interface{}
	err = json.NewDecoder(bytes.NewReader(sample)).Decode(&out)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := jsonRequest("POST", "/register", out)

	s.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req = jsonRequest("POST", "/api/test", nil)

	s.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusCreated)
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
