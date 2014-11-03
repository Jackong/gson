package gson

import "encoding/json"

type Raw struct {
	json.RawMessage
}

type Gson map[string]Raw

func (r Raw) Map() (err error, m Gson) {
	return r.Decode(&m), m
}

func (r Raw) Array() (err error, arr []Raw) {
	return r.Decode(&arr), arr
}

func (r Raw) String() (err error, str string) {
	return r.Decode(&str), str
}

func (r Raw) Float() (err error, f float64) {
	return r.Decode(&f), f
}

func (r Raw) Int() (err error, i int64) {
	return r.Decode(&i), i
}

func (r Raw) Uint() (err error, ui uint64) {
    return r.Decode(&ui), ui
}

func (r Raw) Bool() (err error, b bool) {
	return r.Decode(&b), b
}

func (r Raw) Decode(res interface {}) error {
    return json.Unmarshal(r.RawMessage, res)
}
