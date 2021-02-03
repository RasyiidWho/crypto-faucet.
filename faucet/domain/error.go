package domain
import (
	"fmt"
)
type Error struct {
	Code    int
	Err     error
	Message interface{}
}
func (e *Error) Error() string {
	return fmt.Sprintf("Code %+v, Error: %s", e.Code, e.Err)
}
func (e *Error) ErrorWithMessage() string {
	return fmt.Sprintf("%s", e.Message)
}
func (e *Error) NewError(err error) *Error {
	newE := &Error{}
	newE.Code = e.Code
	newE.Message = e.Message
	newE.Err = err
	return newE
}
var (
	ErrIncognitoService   		= 	&Error{Code: -1001, Message: "incognito service error"}
	ErrTransactionLog           = 	&Error{Code: -4000, Message: "transaction log error"}
	ErrTransactionLogRepository = 	&Error{Code: -4001, Message: "transaction log repository error"}
	ErrPaymentAddress           = 	&Error{Code: -4002, Message: "address is invalid"}
	ErrInvalidTransaction       = 	&Error{Code: -4003, Message: "transaction is invalid"}
)
