package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssociate(t *testing.T) {
	r := Associate([]string{"foo", "bar", "far"}, func(it string) (string, string) {
		return it + "_key", it
	})
	t.Log(r)
	assert.Equal(t, map[string]string{
		"foo_key": "foo",
		"bar_key": "bar",
		"far_key": "far",
	}, r)
}
