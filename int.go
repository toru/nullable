package nullable

import (
	"database/sql"
	"database/sql/driver"
)

// Int64 is a type alias against the standard sql.NullInt64 type.
type Int64 sql.NullInt64

// NewInt64 returns an Int64 populated with the given int64.
func NewInt64(val int64) Int64 {
	return Int64{Int64: val, Valid: true}
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
