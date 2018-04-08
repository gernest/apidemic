package apidemic

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pmylund/go-cache"
)

// Version is the version of apidemic. Apidemic uses semver.
const Version = "0.3"

var maxItemTime = cache.DefaultExpiration

var store = func() *cache.Cache {
	c := cache.New(5*time.Minute, 30*time.Second)
	return c
}()

var allowedHttpMethods = []string{"OPTIONS", "GET", "POST", "PUT", "DELETE", "HEAD"}

// API is the struct for the json object that is passed to apidemic for registration.
type API struct {
	Endpoint   string                 `json:"endpoint"`
	HTTPMethod string                 `json:"http_method"`
	Payload    map[string]interface{} `json:"payload"`
}

// Home renders hopme page. It renders a json response with information about the service.
func Home(w http.ResponseWriter, r *http.Request) {
	details := make(map[string]interface{})
	details["app_name"] = "ApiDemic"
	details["version"] = Version
	details["details"] = "Fake JSON API response"
	RenderJSON(w, http.StatusOK, details)
	return
}

// RenderJSON helper for rendering JSON response, it marshals value into json and writes
// it into w.
func RenderJSON(w http.ResponseWriter, code int, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// RegisterEndpoint receives API objects and registers them. The payload from the request is
// transformed into a self aware Value that is capable of faking its own attribute.
func RegisterEndpoint(w http.ResponseWriter, r *http.Request) {
	var httpMethod string
	a := API{}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		RenderJSON(w, http.StatusBadRequest, NewResponse(err.Error()))
		return
	}

	if httpMethod, err = getAllowedMethod(a.HTTPMethod); err != nil {
		RenderJSON(w, http.StatusBadRequest, NewResponse(err.Error()))
		return
	}

	eKey := getCacheKeys(a.Endpoint, httpMethod)
	if _, ok := store.Get(eKey); ok {
		RenderJSON(w, http.StatusOK, NewResponse("endpoint already taken"))
		return
	}
	obj := NewObject()
	err = obj.Load(a.Payload)
	if err != nil {
		RenderJSON(w, http.StatusInternalServerError, NewResponse(err.Error()))
		return
	}
	store.Set(eKey, obj, maxItemTime)
	RenderJSON(w, http.StatusOK, NewResponse("cool"))
}

func getCacheKeys(endpoint, httpMethod string) string {
	eKey := fmt.Sprintf("%s-%v-e", endpoint, httpMethod)

	return eKey
}

func getAllowedMethod(method string) (string, error) {
	if method == "" {
		return "GET", nil
	}

	for _, m := range allowedHttpMethods {
		if method == m {
			return m, nil
		}
	}

	return "", errors.New("HTTP method is not allowed")
}

// DynamicEndpoint renders registered endpoints.
func DynamicEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := http.StatusOK

	eKey := getCacheKeys(vars["endpoint"], r.Method)
	if eVal, ok := store.Get(eKey); ok {
		if r.Method == "POST" {
			code = http.StatusCreated
		}

		RenderJSON(w, code, eVal)
		return
	}

	responseText := fmt.Sprintf("apidemic: %s has no %s endpoint", vars["endpoint"], r.Method)
	RenderJSON(w, http.StatusNotFound, NewResponse(responseText))
}

// NewResponse helper for response JSON message
func NewResponse(message string) interface{} {
	return struct {
		Text string `json:"text"`
	}{
		message,
	}
}

// NewServer returns a new apidemic server
func NewServer() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/", Home)
	m.HandleFunc("/register", RegisterEndpoint).Methods("POST")
	m.HandleFunc("/api/{endpoint}", DynamicEndpoint).Methods("OPTIONS", "GET", "POST", "PUT", "DELETE", "HEAD")
	return m
}
