package errors

import "net/http"

// HandlerFunc is an error handler function, for instance http.Error
type HandlerFunc func(w http.ResponseWriter, message string, code int)
