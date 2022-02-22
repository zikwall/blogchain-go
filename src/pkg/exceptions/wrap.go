package exceptions

import "fmt"

// WrapError structure is the main wrapper for all errors in the application.
// After analyzing this error, you can later get more detailed information, which can be:
// - public (i.e. you can return it to the user in "raw" form)
// - or private (database errors, file system errors, etc.)
//
// It can be used in the global error handler (middleware)
//
// Examples:
//
// ```go
// 	if err != nil {
//		return exceptions.Wrap("failed parse number", exceptions.ThrowPublicError(err))
// 	}
// ```
//
// this is private error, database access, details of the connection to the database,
// as well as other confidential information, may be disclosed here
// ```go
//  if err := query.ScanStructsContext(context, &content); err != nil {
//		return exceptions.ThrowPrivateError(err)
//	}
// ```
type WrapError struct {
	Context string
	Err     error
}

func (e *WrapError) Error() string {
	return e.Err.Error()
}

func (e *WrapError) Unwrap() error {
	return e.Err
}

func Wrap(context string, err error) error {
	return fmt.Errorf("%v: %w", context, &WrapError{
		Context: context,
		Err:     err,
	})
}
