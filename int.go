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

// Int64 is a type alias against the standard sql.NullInt64 type.
type Int64 sql.NullInt64

// NewInt64 returns an Int64 populated with the given int64.
func NewInt64(value int64) Int64 {
	return Int64{Int64: value, Valid: true}
}

// Scan wraps the standard Scan function, which implements the Scanner interface.
func (i *Int64) Scan(value any) error {
	ni := sql.NullInt64(*i)
	return ni.Scan(value)
}

// Value wraps the standard Value function, which implements the driver Valuer interface.
func (i Int64) Value() (driver.Value, error) {
	return sql.NullInt64(i).Value()
}

// Null returns true if the underlying value is NULL.
func (i Int64) Null() bool {
	return !i.Valid
}

// Nil is an alias for Null() for those that prefer a more Go-like syntax.
func (i Int64) Nil() bool {
	return i.Null()
}

// Int32 is a type alias against the standard sql.NullInt32 type.
type Int32 sql.NullInt32

// NewInt32 returns an Int64 populated with the given int32.
func NewInt32(value int32) Int32 {
	return Int32{Int32: value, Valid: true}
}

// Scan wraps the standard Scan function, which implements the Scanner interface.
func (i *Int32) Scan(value any) error {
	ni := sql.NullInt32(*i)
	return ni.Scan(value)
}

// Value wraps the standard Value function, which implements the driver Valuer interface.
func (i Int32) Value() (driver.Value, error) {
	return sql.NullInt32(i).Value()
}

// Null returns true if the underlying value is NULL.
func (i Int32) Null() bool {
	return !i.Valid
}

// Nil is an alias for Null() for those that prefer a more Go-like syntax.
func (i Int32) Nil() bool {
	return i.Null()
}

// Int16 is a type alias against the standard sql.NullInt16 type.
type Int16 sql.NullInt16

// NewInt16 returns an Int16 populated with the given int16.
func NewInt16(value int16) Int16 {
	return Int16{Int16: value, Valid: true}
}

// Scan wraps the standard Scan function, which implements the Scanner interface.
func (i *Int16) Scan(value any) error {
	ni := sql.NullInt16(*i)
	return ni.Scan(value)
}

// Value wraps the standard Value function, which implements the driver Valuer interface.
func (i Int16) Value() (driver.Value, error) {
	return sql.NullInt16(i).Value()
}

// Null returns true if the underlying value is NULL.
func (i Int16) Null() bool {
	return !i.Valid
}

// Nil is an alias for Null() for those that prefer a more Go-like syntax.
func (i Int16) Nil() bool {
	return i.Null()
}
