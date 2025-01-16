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

// test

func TEST_errorHandler(retryNum int) (bool, *T_Error) {
	if retryNum >= 2 {
		return true, nil
	}
	return false, ThrowERROR_CODE_DB_ERROR(fmt.Sprint(retryNum), "数据库写入异常")
}

func TEST_serviceFunc(retryNum int) {
	isNext, err := TEST_errorHandler(retryNum)
	if err != nil && err.GetCode() == CODE_DB_ERROR {
		fmt.Printf("执行失败，进入重试 %v \n", retryNum)
		err.PrintStackError()
		TEST_serviceFunc(retryNum + 1)
		return
	}
	fmt.Println("执行成功", isNext)
}
