// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestRestrictHighlightAndSnippetArrays(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.RestrictHighlightAndSnippetArraysOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.RestrictHighlightAndSnippetArrays(false),
		},
		{
			opts:     []interface{}{opt.RestrictHighlightAndSnippetArrays(true)},
			expected: opt.RestrictHighlightAndSnippetArrays(true),
		},
		{
			opts:     []interface{}{opt.RestrictHighlightAndSnippetArrays(false)},
			expected: opt.RestrictHighlightAndSnippetArrays(false),
		},
	} {
		var (
			in  = ExtractRestrictHighlightAndSnippetArrays(c.opts...)
			out opt.RestrictHighlightAndSnippetArraysOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
