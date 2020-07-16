package errors

import (
	"fmt"
	"todomvc/core/codes"
)

type RPCError struct {
	Code codes.Code
	Desc string
}

func (rpcError RPCError) Error() string {
	msg := fmt.Sprint("%s,%s", rpcError.Code, rpcError.Desc)
	return msg
}
