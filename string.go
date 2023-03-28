package nullable

import (
	"database/sql"
	"database/sql/driver"
)

// String is a type alias against the standard sql.NullString type.
type String sql.NullString

// NewString returns a String populated with the given string.
func NewString(val string) String {
	return String{String: val, Valid: true}
}

// Scan wraps the standard Scan function, which implements the Scanner interface.
func (s *String) Scan(value any) error {
	return s.Scan(value)
}

// Value wraps the standard Value function, which implements the driver Valuer interface.
func (s String) Value() (driver.Value, error) {
	return s.Value()
}

// Present returns true if the value is a non-empty string.
func (s String) Present() bool {
	return s.Valid && len(s.String) > 0
}

// Null returns true if the underlying value is NULL.
func (s String) Null() bool {
	return s.Valid == false
}

// Empty returns true if the underlying value is either NULL or an empty string.
// Use the Null() function if you want to test specifically for NULL.
func (s String) Empty() bool {
	return !s.Valid || len(s.String) == 0
}
