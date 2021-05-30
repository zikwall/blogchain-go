package exceptions

type ErrApplicationLogic struct {
	Err error
}

func NewErrApplicationLogic(err error) *ErrApplicationLogic {
	return &ErrApplicationLogic{
		Err: err,
	}
}

func (e *ErrApplicationLogic) Error() string {
	return e.Err.Error()
}

type ErrDatabaseAccess struct {
	Err error
}

func NewErrDatabaseAccess(err error) *ErrDatabaseAccess {
	return &ErrDatabaseAccess{
		Err: err,
	}
}

func (e *ErrDatabaseAccess) Error() string {
	return e.Err.Error()
}
