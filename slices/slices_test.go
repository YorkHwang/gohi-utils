package slices

import (
	"context"
	"fmt"
	"testing"
	"time"
	"website-plugin-service/json"

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

func TestMap(t *testing.T) {

	var kids []*Kid

	var m = 5
	for i := 0; i < m; i++ {
		kids = append(kids, &Kid{
			Id:    int64(i),
			Name:  fmt.Sprintf("name_%v", i),
			Age:   int32(i + 1),
			Birth: time.Now(),
		})
	}

	kids = append(kids, &Kid{
		Id:    int64(0),
		Name:  fmt.Sprintf("name_%v", 0),
		Age:   int32(0 + 1),
		Birth: time.Now(),
	})

	var kidMap = make(map[int64][]*Kid)
	for i := range kids {
		k := kids[i]
		kidMap[k.Id] = append(kidMap[k.Id], k)
	}

	fmt.Printf("kidMap=%v\n", kidMap)

	kmap := Associate(kids, func(it *Kid) (string, *Kid) {
		return it.Name, it
	})

	t.Log(json.MarshalStrErrorWriteLog(context.Background(), kmap))

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

type Kid struct {
	Id    int64
	Name  string
	Age   int32
	Birth time.Time
}
