// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type UnretrievableAttributesOption struct {
	value []string
}

func UnretrievableAttributes(v ...string) *UnretrievableAttributesOption {
	return &UnretrievableAttributesOption{v}
}

func (o UnretrievableAttributesOption) Get() []string {
	return o.value
}

func (o UnretrievableAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *UnretrievableAttributesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
