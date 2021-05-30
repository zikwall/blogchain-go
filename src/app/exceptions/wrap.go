package exceptions

import "fmt"

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
