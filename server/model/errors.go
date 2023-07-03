package model

import "errors"

var ErrDeckNotFound = errors.New("deck not found")

type Err struct {
	Error string `json:"error"`
}

type MyCustomErrors struct {
	ErrorText      string
	ErrorCode      int
	AdditionalInfo string
}

func (err MyCustomErrors) Error() string {
	return err.ErrorText
}
