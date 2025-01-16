package lib

import "fmt"

type T_Error struct {
	uid     string
	message string
	code    int
}

const (
	CODE_THREE_PART = -300
	CODE_DB_ERROR   = -400
)

func NewT_Error(uid, message string, code int) *T_Error {
	return &T_Error{
		uid:     uid,
		message: message,
		code:    code,
	}
}

func (e *T_Error) Error() string {
	return e.message
}

func (e *T_Error) GetCode() int {
	return e.code
}

func (e *T_Error) PrintStackError() {
	fmt.Printf("T_Error | uid %s | code %d \n", e.uid, e.code)
}

// 各种类型的TYPE
func ThrowERROR_CODE_THREE_PART(uid, msg string) *T_Error {
	return NewT_Error(uid, msg, CODE_THREE_PART)
}

func ThrowERROR_CODE_DB_ERROR(uid, msg string) *T_Error {
	return NewT_Error(uid, msg, CODE_DB_ERROR)
}
