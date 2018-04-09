package apidemic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicEndpointFailsWithoutRegistration(t *testing.T) {
	s := NewServer()
	payload := registerPayload(t, "fixtures/sample_request.json")

	w := httptest.NewRecorder()
	req := JsonRequest("POST", "/api/test", payload)
	s.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDynamicEndpointWithGetRequest(t *testing.T) {
	s := NewServer()
	payload := registerPayload(t, "fixtures/sample_request.json")

	w := httptest.NewRecorder()
	req := JsonRequest("POST", "/register", payload)
	s.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req = JsonRequest("GET", "/api/test", nil)
	s.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDynamicEndpointWithPostRequest(t *testing.T) {
	s := NewServer()
	payload := registerPayload(t, "fixtures/sample_request.json")
	payload["http_method"] = "POST"

	w := httptest.NewRecorder()
	req := JsonRequest("POST", "/register", payload)
	s.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req = JsonRequest("POST", "/api/test", nil)

	s.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
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
