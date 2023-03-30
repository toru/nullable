// Copyright 2023 Toru Maesaka
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
	"database/sql"
	"database/sql/driver"
)

// Byte is a type alias against the standard sql.NullByte type.
type Byte sql.NullByte

// NewByte returns a Byte populated with the given byte.
func NewByte(value byte) Byte {
	return Byte{Byte: value, Valid: true}
}

// Scan wraps the standard Scan function, which implements the Scanner interface.
func (b *Byte) Scan(value any) error {
	nb := sql.NullByte(*b)
	return nb.Scan(value)
}

// Value wraps the standard Value function, which implements the driver Valuer interface.
func (b Byte) Value() (driver.Value, error) {
	return sql.NullByte(b).Value()
}

// Null returns true if the underlying value is NULL.
func (b Byte) Null() bool {
	return !b.Valid
}

// Nil is an alias for Null() for those that prefer a more Go-like syntax.
func (b Byte) Nil() bool {
	return b.Null()
}
