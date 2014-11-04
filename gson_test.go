package gson_test

import (
	"testing"
	"encoding/json"
	"github.com/Jackong/gson"
	"fmt"
)

func TestDecode(t *testing.T) {
    jsonStr := []byte(`{
    	"int": -11,
    	"uint": 1231,
    	"float": 2131.12312,
    	"bool": true,
    	"string": "hello gson",
    	"strArr": ["elem1", "elem2"],
    	"intArr": [9, 0, 3],
    	"object": {
    		"key1": "value1"
    	}
    }`)
	var obj gson.Gson
	json.Unmarshal(jsonStr, &obj)
	if obj.Get("int").Int() != -11 {
		t.Error("failed to decode int")
	}

	if  obj.Get("uint").Uint() != 1231 {
		t.Error("failed to decode uint")
	}

	if obj.Get("float").Float() != 2131.12312 {
		t.Error("failed to decode float")
	}

	if !obj["bool"].Bool() {
		t.Error("failed to decode bool")
	}

	if obj["string"].String() != "hello gson" {
		t.Error("failed to decode string")
	}

	if len(obj["strArr"].Array()) != 2 {
		t.Error("failed to decode string array")
	}

	for idx, e := range obj["strArr"].Array() {
		if e.String() != fmt.Sprint("elem", idx + 1) {
			t.Error("failed to expect element")
		}
	}

	if len(obj["intArr"].Array()) != 3 {
		t.Error("failed to decode int array")
	}

	if  len(obj.Get("object").Map()) != 1 {
		t.Error("failed to decode object")
	}

	if obj["notExist"].String() != "" {
		t.Error("failed to decode not exists value")
	}

	if obj["notExist"].String("default") != "default" {
		t.Error("failed to decode not default value", obj["notExist"].String("default"))
	}
}
