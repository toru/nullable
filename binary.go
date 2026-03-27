// Copyright 2026 Toru Maesaka
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package nullable

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
)

// Binary holds a nullable byte slice value.
type Binary struct {
	Bytes []byte
	Valid bool
}

// NewBinary returns a Binary populated with the given byte slice.
func NewBinary(value []byte) Binary {
	return Binary{Bytes: value, Valid: true}
}

// Scan implements the sql.Scanner interface.
func (b *Binary) Scan(value any) error {
	if value == nil {
		b.Bytes, b.Valid = nil, false
		return nil
	}

	switch v := value.(type) {
	case []byte:
		b.Bytes = make([]byte, len(v))
		copy(b.Bytes, v)
		b.Valid = true

		return nil
	}

	return fmt.Errorf("cannot scan type %T into binary", value)
}

// Value implements the driver.Valuer interface.
func (b Binary) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil
	}

	return b.Bytes, nil
}

// Present returns true if the value is a non-empty byte slice.
func (b Binary) Present() bool {
	return b.Valid && len(b.Bytes) > 0
}

// Null returns true if the underlying value is NULL.
func (b Binary) Null() bool {
	return !b.Valid
}

// Nil is an alias for Null() for those who prefer a more Go-like syntax.
func (b Binary) Nil() bool {
	return b.Null()
}

// HexString returns a hexadecimal string representation of the underlying value.
func (b Binary) HexString() string {
	if !b.Valid {
		return hex.EncodeToString([]byte{})
	}

	return hex.EncodeToString(b.Bytes)
}
