package gson_test

import (
	"testing"
	"encoding/json"
	"github.com/Jackong/gson"
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
	if err, _int := obj["int"].Int(); err != nil || _int != -11 {
		t.Error("failed to decode int")
	}

	if err, _uint := obj["uint"].Uint(); err != nil || _uint != 1231 {
		t.Error("failed to decode uint")
	}

	if err, _float := obj["float"].Float(); err != nil || _float != 2131.12312 {
		t.Error("failed to decode float")
	}

	if err, _bool := obj["bool"].Bool(); err != nil || !_bool {
		t.Error("failed to decode bool")
	}

	if err, _string := obj["string"].String(); err != nil || _string != "hello gson" {
		t.Error("failed to decode string")
	}


	if err, _strArr := obj["strArr"].Array(); err != nil || len(_strArr) != 2 {
		t.Error("failed to decode string array")
	}

	if err, _intArr := obj["intArr"].Array(); err != nil || len(_intArr) != 3 {
		t.Error("failed to decode int array")
	}

	if err, _obj := obj["object"].Map(); err != nil || len(_obj) != 1 {
		t.Error("failed to decode object")
	}
}
