// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAttributesForFaceting(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.AttributesForFacetingOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AttributesForFaceting(),
		},
		{
			opts:     []interface{}{opt.AttributesForFaceting("value1")},
			expected: opt.AttributesForFaceting("value1"),
		},
		{
			opts:     []interface{}{opt.AttributesForFaceting("value1", "value2", "value3")},
			expected: opt.AttributesForFaceting("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractAttributesForFaceting(c.opts...)
			out opt.AttributesForFacetingOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
