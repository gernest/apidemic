package apidemic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fatih/structs"
)

func TestAPI(t *testing.T) {
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
	req.Header.Set("Contet-Type", "application/json")
	return req
}

func TestCrap(t *testing.T) {
	s := structs.New(fieldTags)
	for _, v := range s.Values() {
		fmt.Printf(" %s | %s \n", v, strings.Replace(fmt.Sprint(v), "_", " ", -1))
	}
}
