package oops

import "fmt"

type OopsError struct {
	Message string
}

func (e OopsError) Error() string {
	return e.Message
}

func Oops(e OopsError) error {
	return OopsError{
		Message: e.Message,
	}
}

func New(Message string) error {
	return Oops(OopsError{
		Message: Message,
	})
}

type LineGetter interface {
	GetLine() int
}

type Node struct {
	Line int
}

func (t Node) GetLine() int {
	return t.Line
}

func CreateErrorMessage[T LineGetter](node T, msg string, args ...interface{}) string {
	return fmt.Sprintf(msg+" at line %d", append(args, node.GetLine())...)
}
