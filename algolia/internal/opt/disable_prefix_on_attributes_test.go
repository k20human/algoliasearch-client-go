// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestDisablePrefixOnAttributes(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.DisablePrefixOnAttributesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.DisablePrefixOnAttributes(),
		},
		{
			opts:     []interface{}{opt.DisablePrefixOnAttributes("value1")},
			expected: opt.DisablePrefixOnAttributes("value1"),
		},
		{
			opts:     []interface{}{opt.DisablePrefixOnAttributes("value1", "value2", "value3")},
			expected: opt.DisablePrefixOnAttributes("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractDisablePrefixOnAttributes(c.opts...)
			out opt.DisablePrefixOnAttributesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
