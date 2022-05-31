package trip

type ErrorType int

const (
	InternalError ErrorType = iota
	ClientRelatedError
)

type TripError struct {
	Err    error
	Type   ErrorType
	Reason string
}

func NewInternalTripError(err error) *TripError {
	return &TripError{
		Err:    err,
		Type:   InternalError,
		Reason: "Internal error",
	}
}

func NewClientRelatedTripError(err error, reason string) *TripError {
	return &TripError{
		Err:    err,
		Type:   ClientRelatedError,
		Reason: reason,
	}
}
