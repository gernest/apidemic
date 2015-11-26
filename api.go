package apidemic

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pmylund/go-cache"
)

const Version = "0.1"

var maxItemTime = cache.DefaultExpiration

var store = func() *cache.Cache {
	c := cache.New(5*time.Minute, 30*time.Second)
	return c
}()

type API struct {
	Endpoint string                 `json:"endpoint"`
	Payload  map[string]interface{} `json:"payload"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	details := make(map[string]interface{})
	details["app_name"] = "ApiDemic"
	details["version"] = Version
	details["details"] = "Fake JSON API response"
	RenderJSON(w, http.StatusOK, details)
	return
}

func RenderJSON(w http.ResponseWriter, code int, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RegisterEndpoint(w http.ResponseWriter, r *http.Request) {
	a := API{}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		RenderJSON(w, http.StatusInternalServerError, NewResponse(err.Error()))
		return
	}
	if _, ok := store.Get(a.Endpoint); ok {
		RenderJSON(w, http.StatusOK, NewResponse("endpoint already taken"))
		return
	}
	obj := NewObject()
	err = obj.Load(a.Payload)
	if err != nil {
		RenderJSON(w, http.StatusInternalServerError, NewResponse(err.Error()))
		return
	}
	store.Set(a.Endpoint, obj, maxItemTime)
	RenderJSON(w, http.StatusOK, NewResponse("cool"))
}

func GetEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	endpoint := vars["endpoint"]
	if eVal, ok := store.Get(endpoint); ok {
		RenderJSON(w, http.StatusOK, eVal)
		return
	}
	RenderJSON(w, http.StatusNotFound, NewResponse("apidemic: "+endpoint+" is not found"))
}

func NewResponse(message string) interface{} {
	return struct {
		Text string `json:"text"`
	}{
		message,
	}
}

func NewServer() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/", Home)
	m.HandleFunc("/register", RegisterEndpoint).Methods("POST")
	m.HandleFunc("/api/{endpoint}", GetEndpoint).Methods("GET")
	return m
}
