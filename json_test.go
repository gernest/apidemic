package apidemic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseJSONDataOfNonArray(t *testing.T) {
	ob := loadObject(t, "fixtures/sample.json")

	res, err := json.MarshalIndent(ob, "", "\t")
	require.NoError(t, err)

	str := string(res)
	assert.Equal(t, "{", str[:1])
	assert.Equal(t, "}", str[len(str)-1:])
}

func TestParseJSONDataOfArray(t *testing.T) {
	ob := loadObject(t, "fixtures/sample.json")
	ob.IsArray = true
	ob.MaxCount = 10

	res, err := json.MarshalIndent(ob, "", "\t")
	require.NoError(t, err)

	str := string(res)
	assert.Equal(t, "[", str[:1])
	assert.Equal(t, "]", str[len(str)-1:])
}

func loadObject(t *testing.T, fixture string) (*Object) {
	data, err := ioutil.ReadFile(fixture)
	require.NoError(t, err)

	ob, err :=  parseJSONData(bytes.NewReader(data))
	require.NoError(t, err)

	return ob
}