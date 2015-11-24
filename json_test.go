package apidemic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestParseJSONData(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/sample.json")
	if err != nil {
		t.Fatal(err)
	}
	ob, err := parseJSONData(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}

	_, err = json.MarshalIndent(ob, "", "\t")
	if err != nil {
		t.Error(err)
	}
}
