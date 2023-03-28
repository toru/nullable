# Nullable

This repository is under heavy construction.

The nullable package provides an [Active Support](https://guides.rubyonrails.org/active_support_core_extensions.html)
inspired syntax sugar on top of the nullable data types provided by the standard
[database/sql](https://github.com/golang/go/tree/master/src/database/sql) package.
The interface should feel natural to those who are familiar with [Rails](https://github.com/rails/rails).

## Quick Examples

```go
// Check if the underlying string is NULL.
if value.Null() {
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
```

## Motivation

Go's standard library provides precisely what it should: Well thought-out
primitive building blocks for specific purposes. Syntax sugar like the ones
provided by this package is subjective, hence the inception of this package.