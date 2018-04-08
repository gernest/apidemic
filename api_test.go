package apidemic

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
