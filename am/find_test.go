package am

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWillFindInArray(t *testing.T) {

	result := &AmResult{
		Result: []Result{
			{
				ID: "123",
			},
		},
	}

	idPredicate := func(r *Result) string {
		return r.ID
	}

	usernamePredicate := func(r *Result) string {
		return r.Username
	}

	assert.True(t, Find("123", result, idPredicate))
	assert.False(t, Find("abc", result, idPredicate))

	assert.False(t, Find("123", result, usernamePredicate))
}

func TestWillReturnId(t *testing.T) {
	result := &AmResult{
		Result: []Result{
			{
				ID:   "123",
				Name: "xyz",
			},
		},
	}

	predicate := func(r *Result) string {
		return r.Name
	}

	assert.Equal(t, "123", FindIdByName("xyz", result, predicate))
	assert.Equal(t, "", FindIdByName("doesnt exist", result, predicate))
}
