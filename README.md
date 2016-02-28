# Simple Go struct validation library

This is a just for fun library to try out Go reflect library and tags.

## How to use

Simply add a `validate` tag as shown below

```go
type MyStruct struct {
  RequiredInt int `validate:"required"`
  RequiredString  string `validate:"required"`
  OptionalVariable int
}
```

## Future work?

- Nested validation currently fails
- Email validation
