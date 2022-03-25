package yapi

import (
	"encoding/json"
	"testing"
)

func TestWrite(t *testing.T) {
	obj := new(yApiObject)
	err := json.Unmarshal([]byte(demo), obj)

	t.Log(obj.Title, err)
	t.Log(obj.YType)
	t.Log(obj.Description)

	for k, v := range obj.Properties {
		t.Log(k) // field Name
		t.Log(v.YType)
	}
}
