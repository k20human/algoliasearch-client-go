// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type ExactOnSingleWordQueryOption struct {
	value string
}

func ExactOnSingleWordQuery(v string) *ExactOnSingleWordQueryOption {
	return &ExactOnSingleWordQueryOption{v}
}

func (o ExactOnSingleWordQueryOption) Get() string {
	return o.value
}

func (o ExactOnSingleWordQueryOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ExactOnSingleWordQueryOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = "attribute"
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
