package main

import (
	"fmt"

	"chelizichen.com/win/lib"
)

func main() {
	fmt.Println("hello world")
	isNext, err := errorHandler()
	if err != nil && err.GetCode() == lib.CODE_DB_ERROR {
		fmt.Println("handle db error")
	}
	fmt.Println("isNext", isNext)
}

func errorHandler() (bool, *lib.T_Error) {
	return false, lib.ThrowERROR_CODE_DB_ERROR("1234", "asmndkoasd")
}
