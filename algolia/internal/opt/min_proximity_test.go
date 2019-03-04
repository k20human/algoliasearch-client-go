// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestMinProximity(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.MinProximityOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.MinProximity(1),
		},
		{
			opts:     []interface{}{opt.MinProximity(0)},
			expected: opt.MinProximity(0),
		},
		{
			opts:     []interface{}{opt.MinProximity(1)},
			expected: opt.MinProximity(1),
		},
		{
			opts:     []interface{}{opt.MinProximity(-42)},
			expected: opt.MinProximity(-42),
		},
	} {
		var (
			in  = ExtractMinProximity(c.opts...)
			out opt.MinProximityOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
