// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type ReplicasOption struct {
	value []string
}

func Replicas(v ...string) *ReplicasOption {
	return &ReplicasOption{v}
}

func (o ReplicasOption) Get() []string {
	return o.value
}

func (o ReplicasOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ReplicasOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
