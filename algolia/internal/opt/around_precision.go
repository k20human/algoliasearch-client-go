// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractAroundPrecision(opts ...interface{}) *opt.AroundPrecisionOption {
	for _, o := range opts {
		if v, ok := o.(*opt.AroundPrecisionOption); ok {
			return v
		}
	}
	return nil
}
