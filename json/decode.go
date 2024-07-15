package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
)

type Unmarshaler = json.Unmarshaler

// Unmarshal data to value with type T, this shortcut can effectively prevent declaring a dummy variable by var.
func Unmarshal[T any, S ~[]byte | ~string](data S) (v T, err error) {
	err = json.Unmarshal([]byte(data), &v)
	return
}

// UnmarshalString
//
// deprecated: use Unmarshal instead
func UnmarshalString[T any](data string) (T, error) {
	return Unmarshal[T]([]byte(data))
}

// UnmarshalInto ref, ref must be a pointer, like Unmarshal in stdlib.
//
//	json.UnmarshalInto(`{"json":"value}`, &s.f)
func UnmarshalInto[T any, S ~[]byte | ~string](data S, ref *T) error {
	return json.Unmarshal([]byte(data), ref)
}

type (
	UnmarshalTypeError    = json.UnmarshalTypeError
	UnmarshalFieldError   = json.UnmarshalFieldError // Deprecated: No longer used; kept for compatibility.
	InvalidUnmarshalError = json.InvalidUnmarshalError
)

type Number = json.Number

type StringLike string

func (this *StringLike) UnmarshalJSON(buf []byte) error {
	decoder := NewDecoder(bytes.NewReader(buf))
	decoder.UseNumber()
	if token, err := decoder.Token(); err != nil {
		return err
	} else {
		switch token := token.(type) {
		case nil:
			*this = ""
		case Delim: // object or array
			return errors.New("string like requires a json value")
		case bool:
			if token {
				*this = "true"
			} else {
				*this = "false"
			}
		case Number:
			*this = StringLike(token.String())
		case string:
			*this = StringLike(token)
		default:
			panic("unexpected token type: " + reflect.TypeOf(token).String())
		}
		return nil
	}
}

func (this StringLike) String() string {
	return string(this)
}
