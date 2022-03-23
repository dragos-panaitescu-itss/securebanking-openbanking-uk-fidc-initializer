package common

import (
	"go.uber.org/zap"
	"secure-banking-uk-initializer/pkg/types"
)

// Find look at the Result of an AmResult object, commonly returned by the OpenAm API.
//  If result exists then return true. An AmResult contains an array of Result. which contains the fields
//  ID, Name and Username.
//  Eg. If an AmResult has the result with Username = abc
//  then calling Find("abc", theAmResultObject, func(r *Result) string {
//		return r.Username
//	}) will return true
func Find(arg string, ob *types.AmResult, fn func(*types.Result) string) bool {
	for _, e := range ob.Result {
		if fn(&e) == arg {
			zap.S().Infow("Argument found", "arg", arg)
			return true
		}
	}
	return false
}

func FindIdByName(name string, ob *types.AmResult, fn func(*types.Result) string) string {
	for _, e := range ob.Result {
		if fn(&e) == name {
			zap.S().Infow("ID found", "id", e.ID)
			return e.ID
		}
	}
	return ""
}
