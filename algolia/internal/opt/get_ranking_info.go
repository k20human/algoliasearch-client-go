// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractGetRankingInfo(opts ...interface{}) *opt.GetRankingInfoOption {
	for _, o := range opts {
		if v, ok := o.(*opt.GetRankingInfoOption); ok {
			return v
		}
	}
	return nil
}