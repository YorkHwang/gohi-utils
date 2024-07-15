package json

import (
	"bytes"
	"encoding/json"
	_ "unsafe"
)

type Marshaler = json.Marshaler

//go:linkname Marshal encoding/json.Marshal
func Marshal(v any) ([]byte, error)

func MarshalString(v any) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

//go:linkname MarshalIndent encoding/json.MarshalIndent
func MarshalIndent(v any, prefix string, indent string) ([]byte, error)

func MarshalIndentString(v any, prefix string, indent string) (string, error) {
	b, err := json.MarshalIndent(v, prefix, indent)
	return string(b), err
}

//go:linkname HTMLEscape encoding/json.HTMLEscape
func HTMLEscape(dst *bytes.Buffer, src []byte)

type (
	UnsupportedTypeError  = json.UnsupportedTypeError
	UnsupportedValueError = json.UnsupportedValueError
	InvalidUTF8Error      = json.InvalidUTF8Error
	MarshalerError        = json.MarshalerError
)
