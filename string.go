// Package nullable provides an Active Support like syntax on top of the
// nullable data types provided by the standard database/sql package. The
// interface should feel natural to those who are familiar with Rails.
package nullable

import (
	"crypto/sha256"
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
	ns := sql.NullString(s)
	return ns.Value()
}

// Present returns true if the value is a non-empty string.
func (s String) Present() bool {
	return s.Valid && len(s.String) > 0
}

// Null returns true if the underlying value is NULL.
func (s String) Null() bool {
	return !s.Valid
}

// Empty returns true if the underlying value is either NULL or an empty string.
// Use the Null() function if you want to test specifically for NULL.
func (s String) Empty() bool {
	return !s.Valid || len(s.String) == 0
}

// HexString returns a SHA256 hexadecimal string representation of the
// underlying value. Other algorithms may be supported in the future.
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
	digest := sha256.Sum256([]byte(src))

	return hex.EncodeToString(digest[:])
}
