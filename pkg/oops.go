package pkg

import (
	"strconv"
)

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func Oops(message string) error {
	return &CustomError{message}
}

func IllegalTokenError(token Token) error {
	panic(Oops("Illegal token: " + token.Value + " at position " + strconv.Itoa(token.Start)))
}
