package exceptions

type ErrPublic struct {
	Err error
}

func ThrowPublicError(err error) *ErrPublic {
	return &ErrPublic{
		Err: err,
	}
}

func (e *ErrPublic) Error() string {
	return e.Err.Error()
}

type ErrPrivate struct {
	Err error
}

func ThrowPrivateError(err error) *ErrPrivate {
	return &ErrPrivate{
		Err: err,
	}
}

func (e *ErrPrivate) Error() string {
	return e.Err.Error()
}
