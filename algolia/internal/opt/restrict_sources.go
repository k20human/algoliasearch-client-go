// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractRestrictSources(opts ...interface{}) *opt.RestrictSourcesOption {
	for _, o := range opts {
		if v, ok := o.(*opt.RestrictSourcesOption); ok {
			return v
		}
	}
	return nil
}
