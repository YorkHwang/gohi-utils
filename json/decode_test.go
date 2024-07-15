package json

import (
	"testing"

	"github.com/jasonming/go-fuckerr/try"
	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Name string
	Age  int
}

func TestUnmarshal_struct(t *testing.T) {
	src := `{"name":"foo","age":18}`
	t.Run("struct", func(t *testing.T) {
		v := try.Must(Unmarshal[Foo](src))
		t.Log(v)
		assert.Equal(t, Foo{"foo", 18}, v)
	})
	t.Run("*struct", func(t *testing.T) {
		v := try.Must(Unmarshal[*Foo](src))
		t.Log(v)
		assert.Equal(t, &Foo{"foo", 18}, v)
	})
	t.Run("**struct", func(t *testing.T) {
		v := try.Must(Unmarshal[**Foo](src))
		t.Log(v, *v)
		expected := &Foo{"foo", 18}
		assert.Equal(t, &expected, v)
	})
}

func TestUnmarshal_slice(t *testing.T) {
	t.Run("[]string", func(t *testing.T) {
		v := try.Must(Unmarshal[[]string](`["foo","bar"]`))
		t.Log(v)
		assert.Equal(t, []string{"foo", "bar"}, v)
	})
	t.Run("*[]string", func(t *testing.T) {
		v := try.Must(Unmarshal[*[]string](`["foo","bar"]`))
		t.Log(v)
		assert.Equal(t, &([]string{"foo", "bar"}), v)
	})
	t.Run("**[]string", func(t *testing.T) {
		v := try.Must(Unmarshal[**[]string](`["foo","bar"]`))
		t.Log(v, *v)
		expected := &([]string{"foo", "bar"})
		assert.Equal(t, &expected, v)
	})
	t.Run("[]*string", func(t *testing.T) {
		v := try.Must(Unmarshal[[]*string](`["foo","bar"]`))
		t.Log(v)
		e0, e1 := "foo", "bar"
		assert.Equal(t, []*string{&e0, &e1}, v)
	})
	src := `[{"name":"foo","age":18},{"name":"bar","age":28}]`
	t.Run("[]struct", func(t *testing.T) {
		v := try.Must(Unmarshal[[]Foo](src))
		t.Log(v)
		assert.Equal(t, []Foo{
			{"foo", 18},
			{"bar", 28},
		}, v)
	})
	t.Run("[]*struct", func(t *testing.T) {
		v := try.Must(Unmarshal[[]*Foo](src))
		t.Log(v)
		assert.Equal(t, []*Foo{
			{"foo", 18},
			{"bar", 28},
		}, v)
	})
	t.Run("[]**struct", func(t *testing.T) {
		v := try.Must(Unmarshal[[]**Foo](src))
		t.Log(v)
		e0, e1 := &Foo{"foo", 18}, &Foo{"bar", 28}
		assert.Equal(t, []**Foo{&e0, &e1}, v)
	})
}

func TestStringLike(t *testing.T) {
	t.Run("undefined", func(t *testing.T) {
		_, err := Unmarshal[StringLike]([]byte{})
		assert.NotNil(t, err)
	})
	t.Run("null", func(t *testing.T) {
		assert.EqualValues(t, "", try.Must(Unmarshal[StringLike](`null`)))
	})
	t.Run("bool", func(t *testing.T) {
		assert.EqualValues(t, "true", try.Must(Unmarshal[StringLike](`true`)))
		assert.EqualValues(t, "false", try.Must(Unmarshal[StringLike](`false`)))
	})
	t.Run("string", func(t *testing.T) {
		assert.EqualValues(t, "foo", try.Must(Unmarshal[StringLike](`"foo"`)))
		assert.EqualValues(t, "", try.Must(Unmarshal[StringLike](`""`)))
	})
	t.Run("number", func(t *testing.T) {
		test := func(number string) {
			assert.EqualValues(t, number, try.Must(Unmarshal[StringLike](number)))
		}
		test("0")
		test("0.1")
		test("-0")
		test("-0.1")
		test("1.0100")
		test("1E2")
		test("1e2")
		test("1e-2")
		_, err := Unmarshal[StringLike]("0xFF")
		assert.NotNil(t, err)
	})
	t.Run("non-values", func(t *testing.T) {
		var err error
		_, err = Unmarshal[StringLike]("{}")
		assert.NotNil(t, err)
		_, err = Unmarshal[StringLike]("[]")
		assert.NotNil(t, err)
	})
}
