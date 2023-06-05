package rysrv

import (
	"errors"
)

var (
	renderedParseError         = errors.New("Parse error")
	renderedInvalidRequest     = errors.New("Invalid Request")
	errorMessageInvalidParams  = errors.New("Invalid params")
	errorMessageMethodNotFound = errors.New("Method not found")
	errorMessageInternalError  = errors.New("Internal error")
)
