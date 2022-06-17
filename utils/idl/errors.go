package idl

import "fmt"

type DecodeError struct {
	Types       Tuple
	Description string
}

func (e DecodeError) Error() string {
	return fmt.Sprintf("%s %s", e.Types.String(), e.Description)
}

type FormatError struct {
	Description string
}

func (e FormatError) Error() string {
	return fmt.Sprintf("() %s", e.Description)
}

var NOMAGICError = &FormatError{
	Description: "no magic bytes",
}
var WRONGMAGICError = &FormatError{
	Description: "wrong magic bytes",
}

var EMPTYError = &FormatError{
	Description: "empty",
}
