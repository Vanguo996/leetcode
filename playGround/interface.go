package main

import "fmt"

type err interface{
	Error() string
}

type RPCError struct {
	Code int64
	Message string
}

func (e *RPCError) Error() string{
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func NewRPCError(code int64, msg string) error {
	return &RPCError{ // typecheck3
		Code:    code,
		Message: msg,
	}
}

func main() {
	er := NewRPCError(400, "un")
	fmt.Print(er)
}