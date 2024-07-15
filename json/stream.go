// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"encoding/json"
	"io"
	_ "unsafe"
)

type Decoder = json.Decoder

//go:linkname NewDecoder encoding/json.NewDecoder
func NewDecoder(r io.Reader) *Decoder

type Encoder = json.Encoder

//go:linkname NewEncoder encoding/json.NewEncoder
func NewEncoder(w io.Writer) *Encoder

type RawMessage = json.RawMessage

type Token = json.Token

type Delim = json.Delim
