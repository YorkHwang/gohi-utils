package slices

import (
	"fmt"
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

func TestRange(t *testing.T) {

	strlist := []string{"foo", "bar", "far"}
	for i, v := range strlist {
		if i == 0 || v == "far" {
			strlist[0] = "kk"
		}

	}

	fmt.Printf("strlist=%v\n", strlist)

	r := Associate([]string{"foo", "bar", "far"}, func(it string) (string, string) {
		return it + "_key", it
	})
	t.Log(r)
	assert.Equal(t, map[string]string{
		"foo_key": "foo",
		"bar_key": "bar",
		"far_key": "far",
	}, r)

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 0, len(a))
	for i, v := range a {
		if i == 0 {
			a[2] = 1
			a[3] = 6
		}

		b = append(b, v)
	}

	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)

}
