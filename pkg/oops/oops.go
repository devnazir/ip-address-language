package oops

type OopsError struct {
	message string
}

func (e *OopsError) Error() string {
	return e.message
}

func Oops(e OopsError) error {
	return &OopsError{
		message: e.message,
	}
}

func New(message string) error {
	return Oops(OopsError{
		message: message,
	})
}
