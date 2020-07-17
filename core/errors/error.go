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
	return fmt.Sprintf(rpcError.Desc)
}
