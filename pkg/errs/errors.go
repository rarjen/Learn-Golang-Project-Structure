package errs

import "fmt"

var (
	ERR_PING = newError("err_ping")

	ERR_INIT_NOT_FOUND = newError("err_init_not_found")

	//USER
	ERR_CREATE_USER    = newError("err_create_user")
	ERR_UPDATE_USER    = newError("err_update_user")
	ERR_USER_NOT_FOUND = newError("err_user_not_found")
	ERR_DELETE_USER    = newError("err_delete_user")

	//PROGRAM
	ERR_CREATE_PROGRAM   = newError("err_create_program")
	ERR_GET_ALL_PROGRAMS = newError("err_get_all_programs")
	ERR_GET_ONE_PROGRAM  = newError("err_get_one_program")
	ERR_UPDATE_PROGRAM   = newError("err_update_program")
)

type Error struct {
	errCode string
	title   string
	message string
}

func newError(errCode string) *Error {
	return &Error{
		errCode: errCode,
	}
}

func (err *Error) Error() string {
	return fmt.Sprintf("Error Code: %s", err.errCode)
}

func (err *Error) Title() string {
	return err.title
}

func (err *Error) ErrorCode() string {
	return err.errCode
}

func (err *Error) Message() string {
	return err.message
}
