// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractMaxFacetHits(opts ...interface{}) *opt.MaxFacetHitsOption {
	for _, o := range opts {
		if v, ok := o.(*opt.MaxFacetHitsOption); ok {
			return v
		}
	}
	return nil
}
