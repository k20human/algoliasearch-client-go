// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractMinWordSizeFor1Typo(opts ...interface{}) *opt.MinWordSizeFor1TypoOption {
	for _, o := range opts {
		if v, ok := o.(*opt.MinWordSizeFor1TypoOption); ok {
			return v
		}
	}
	return nil
}
