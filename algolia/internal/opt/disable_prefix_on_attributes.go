// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractDisablePrefixOnAttributes(opts ...interface{}) *opt.DisablePrefixOnAttributesOption {
	for _, o := range opts {
		if v, ok := o.(*opt.DisablePrefixOnAttributesOption); ok {
			return v
		}
	}
	return nil
}
