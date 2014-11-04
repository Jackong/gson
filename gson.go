package gson

import (
	"encoding/json"
)

type Raw struct {
	json.RawMessage
}

type Gson map[string]Raw

func (g Gson) Get(key string) Raw {
	return g[key]
}

func (r Raw) Map(d ...Gson) (m Gson) {
	if err := r.Decode(&m); err != nil && len(d) > 0 {
		m = d[0]
	}
	return
}

func (r Raw) Array(d ...[]Raw) (arr []Raw) {
	if err := r.Decode(&arr); err != nil && len(d) > 0 {
		arr = d[0]
	}
	return
}

func (r Raw) String(d ...string) (str string) {
	if err := r.Decode(&str); err != nil && len(d) > 0 {
		str = d[0]
	}
	return
}

func (r Raw) Float(d ...float64) (f float64) {
	if err := r.Decode(&f); err != nil && len(d) > 0 {
		f = d[0]
	}
	return
}

func (r Raw) Int(d ...int64) (i int64) {
	if err := r.Decode(&i); err != nil && len(d) > 0 {
		i = d[0]
	}
	return
}

func (r Raw) Uint(d ...uint64) (ui uint64) {
	if err := r.Decode(&ui); err != nil && len(d) > 0 {
		ui = d[0]
	}
	return
}

func (r Raw) Bool(d ...bool) (b bool) {
	if err := r.Decode(&b); err != nil && len(d) > 0 {
		b = d[0]
	}
	return
}

func (r Raw) Decode(res interface {}) error {
     return json.Unmarshal(r.RawMessage, res)
}
