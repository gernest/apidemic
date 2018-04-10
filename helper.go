package apidemic

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// JsonRequest creates an HTTP request
func JsonRequest(method string, path string, body interface{}) *http.Request {
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
