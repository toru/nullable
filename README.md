# Nullable

This repository is under heavy construction.

The nullable package provides an [Active Support](https://guides.rubyonrails.org/active_support_core_extensions.html)
inspired syntax on top of the nullable data types provided by the standard
[database/sql](https://github.com/golang/go/tree/master/src/database/sql) package.
The interface should feel natural to those who are familiar with [Rails](https://github.com/rails/rails).

## Quick Examples

```go
// Constructing the value is identical to the standard type.
nullable.String{Value: "hello", Valid: true}

// An equivalent convenience function is also provided.
nullable.NewString("hello")

// Check if the underlying string is NULL. There is also a Nil() function
// alias for those that prefer a more Go-like syntax.
if value.Null() || value.Nil() {
    log.Println("i am stored as NULL in the database")
}

// Check if the underlying value is a non-empty string.
if value.Present() {
    log.Println("i am a non-empty value")
}

// Check if the underlying value is either NULL or an empty string.
if value.Empty() {
    log.Println("i am an empty string")
}

// Get the hexadecimal string representation of the underlying value.
value.HexString()
```

## Continuous Integration

![go workflow](https://github.com/toru/nullable/actions/workflows/go.yml/badge.svg)

## Motivation

Go's standard library provides precisely what it should: Well thought-out
primitive building blocks for specific purposes. Syntax sugar like the ones
provided by this package is subjective, hence the inception of this package.
