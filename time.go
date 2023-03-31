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
	"time"
)

// Time is a type alias against the standard sql.NullTime type.
type Time sql.NullTime

// NewTime returns a Time populated with the given time.Time.
func NewTime(value time.Time) Time {
	return Time{Time: value, Valid: true}
}

// Scan wraps the standard Scan function, which implements the Scanner interface.
func (t *Time) Scan(value any) error {
	nt := sql.NullTime(*t)

	if err := nt.Scan(value); err != nil {
		return err
	}
	*t = Time(nt)

	return nil
}

// Value wraps the standard Value function, which implements the driver Valuer interface.
func (t Time) Value() (driver.Value, error) {
	return sql.NullTime(t).Value()
}

// Null returns true if the underlying value is NULL.
func (t Time) Null() bool {
	return !t.Valid
}

// Nil is an alias for Null() for those that prefer a more Go-like syntax.
func (t Time) Nil() bool {
	return t.Null()
}
