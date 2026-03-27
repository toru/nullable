# Nullable

[![go workflow](https://github.com/toru/nullable/actions/workflows/go.yml/badge.svg)](https://github.com/toru/nullable/actions/workflows/go.yml)
[![go reference](https://pkg.go.dev/badge/github.com/toru/nullable.svg)](https://pkg.go.dev/github.com/toru/nullable)
[![mit license](https://img.shields.io/badge/license-MIT-green)](/LICENSE)

The nullable package provides an unobtrusive, [Active Support](https://guides.rubyonrails.org/active_support_core_extensions.html)
inspired syntax on top of Go's standard [database/sql](https://github.com/golang/go/tree/master/src/database/sql) types.
Because nullable wraps the standard package, it is a simple drop-in replacement.

## Quick Examples

Usage within struct definition.

```go
type IceCream struct {
  // Using the nullable package
  Flavor nullable.String `db:"flavor"`

  // Using the database/sql package
  Topping sql.NullString `db:"topping"`
}
```

### String

#### Initialization

```go
// Using the nullable package
nullable.NewString("hello")

// Equivalent using the standard type
sql.NullString{String: "hello", Valid: true}
```

#### Mutation

```go
// Overwrite the existing value
value.Set("hello there")
```

#### Checking State

```go
// Null() is also available
if value.Nil() {
  log.Println("stored as NULL in the database")
}

if value.Present() {
  log.Println("non-empty string")
}

if value.Empty() {
  log.Println("either NULL or empty string")
}
```

#### Utilities

```go
// Hexadecimal string representation of the underlying value
value.HexString()
```

### Binary

#### Initialization

```go
nullable.NewBinary([]byte("hello"))
```

#### Checking State

```go
if value.Nil() {
  log.Println("stored as NULL in the database")
}

if value.Present() {
  log.Println("non-empty byte slice")
}
```

#### Utilities

```go
// Hexadecimal string representation of the underlying bytes
value.HexString()
```

For all available types, see the [package documentation](https://pkg.go.dev/github.com/toru/nullable).

## Motivation

Go's standard library provides precisely what it should: well-thought-out
primitive building blocks. This package scratches the itch of those who avoid
ORMs but still miss the ergonomics of Rails when handling nullable values.

## Contributing

Contributions of any kind are welcome! If you are going to submit a PR, please
follow [Go's commit message structure](https://go.dev/wiki/CommitMessage).
