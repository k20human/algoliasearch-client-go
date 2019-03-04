// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestRanking(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.RankingOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Ranking([]string{"typo", "geo", "words", "filters", "proximity", "attribute", "exact", "custom"}...),
		},
		{
			opts:     []interface{}{opt.Ranking("value1")},
			expected: opt.Ranking("value1"),
		},
		{
			opts:     []interface{}{opt.Ranking("value1", "value2", "value3")},
			expected: opt.Ranking("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractRanking(c.opts...)
			out opt.RankingOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
