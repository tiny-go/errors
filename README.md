# errors

[![Build Status][circleci-badge]][circleci-link]
[![Report Card][report-badge]][report-link]
[![GoCover][cover-badge]][cover-link]

Collection of HTTP error implementations

### HTTP error interface
```go
type Error interface{
  error
  Code() int
}
```

### Implementations
Code | Corresponding error type
-----|--------------------------
 400 | `BadRequest`  
 401 | `Unauthorized`
 403 | `Forbidden`
 404 | `NotFound`
 405 | `MethodNotAllowed`
 500 | `InternalServer`
 501 | `NotImplemented`

[circleci-badge]: https://circleci.com/gh/tiny-go/errors.svg?style=shield
[circleci-link]: https://circleci.com/gh/tiny-go/errors
[report-badge]: https://goreportcard.com/badge/github.com/tiny-go/errors
[report-link]: https://goreportcard.com/report/github.com/tiny-go/errors
[cover-badge]: https://gocover.io/_badge/github.com/tiny-go/errors
[cover-link]: https://gocover.io/github.com/tiny-go/errors
