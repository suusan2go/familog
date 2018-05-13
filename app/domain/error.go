package domain

type EntityNotFoundError struct {
	Cause error
}

func (e *EntityNotFoundError) Error() string {
	return e.Cause.Error()
}
