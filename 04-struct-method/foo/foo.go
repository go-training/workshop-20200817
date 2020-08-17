package foo

import "fmt"

type Error struct {
	Cateogry Cateogry
	Message  string
	cause    error
}

func (e *Error) Error() string {
	return fmt.Sprintf("category: %s, message: %s", e.Cateogry, e.Message)
}

func (e *Error) Cause() string {
	if e.cause != nil {
		return fmt.Sprintf("category: %s, message: %s, err: %s", e.Cateogry, e.Message, e.cause)
	}
}

type Cateogry string

func (c Cateogry) String() string {
	switch c {
	case AMD:
		return "AMD"
	case Intel:
		return "INTEL"
	}
	return ""
}

func (c Cateogry) Code() int {
	switch c {
	case AMD:
		return 100
	case Intel:
		return 200
	}
	return 404
}

var (
	AMD   Cateogry = "amd"
	Intel Cateogry = "intel"
)

func New(c Cateogry, msg string, err error) error {
	return &Error{
		Cateogry: c,
		Message:  msg,
		cause:    err,
	}
}

func EAmd(msg string, err error) error {
	return &Error{
		Cateogry: AMD,
		Message:  msg,
		cause:    err,
	}
}

func EIntel(msg string, err error) error {
	return &Error{
		Cateogry: Intel,
		Message:  msg,
		cause:    err,
	}
}
