package errors

import "errors"

// New returns an error that formats as the given text.
// Wraps original errors.New function to avoid importing "errors".
func New(text string) error {
	return errors.New(text)
}

// Is reports whether any error in err's chain matches target.
// Wraps original errors.Is function to avoid importing "errors".
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in err's chain that matches target, and if one is found,
// sets target to that error value and returns true. Otherwise, it returns false.
// Wraps original errors.As function to avoid importing "errors".
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's type contains
// an Unwrap method returning error. Otherwise, Unwrap returns nil.
// Wraps original errors.Unwrap function to avoid importing "errors".
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
