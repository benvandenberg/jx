// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	config "gopkg.in/src-d/go-git.v4/config"
)

func AnyPtrToConfigConfig() *config.Config {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*config.Config))(nil)).Elem()))
	var nullValue *config.Config
	return nullValue
}

func EqPtrToConfigConfig(value *config.Config) *config.Config {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *config.Config
	return nullValue
}
