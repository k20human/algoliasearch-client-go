// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractIgnorePlurals(opts ...interface{}) *opt.IgnorePluralsOption {
	for _, o := range opts {
		if v, ok := o.(*opt.IgnorePluralsOption); ok {
			return v
		}
	}
	return nil
}
