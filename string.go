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

// Package nullable provides an Active Support like syntax on top of the
// nullable data types provided by the standard database/sql package. The
// interface should feel natural to those who are familiar with Rails.
package nullable

import (
	"database/sql"
	"database/sql/driver"
	"encoding/hex"
)

// String is a type alias against the standard sql.NullString type.
type String sql.NullString

// NewString returns a String populated with the given string.
func NewString(val string) String {
	return String{String: val, Valid: true}
}

// Scan wraps the standard Scan function, which implements the Scanner interface.
func (s *String) Scan(value any) error {
	ns := sql.NullString(*s)
	return ns.Scan(value)
}

// Value wraps the standard Value function, which implements the driver Valuer interface.
func (s String) Value() (driver.Value, error) {
	return sql.NullString(s).Value()
}

// Present returns true if the value is a non-empty string.
func (s String) Present() bool {
	return s.Valid && len(s.String) > 0
}

// Null returns true if the underlying value is NULL.
func (s String) Null() bool {
	return !s.Valid
}

// Nil is an alias for Null() for those that prefer a more Go-like syntax.
func (s String) Nil() bool {
	return s.Null()
}

// Empty returns true if the underlying value is either NULL or an empty string.
// Use the Null() function if you want to test specifically for NULL.
func (s String) Empty() bool {
	return !s.Valid || len(s.String) == 0
}

// HexString returns a hexadecimal string representation of the underlying value.
func (s String) HexString() string {
	var src string

	// Defensively handle the NULL case instead of computing the digest of the
	// underlying value. For example, it is possible for Valid to be false while
	// the value is a non-empty string.
	if s.Valid {
		src = s.String
	} else {
		src = ""
	}

	return hex.EncodeToString([]byte(src))
}
