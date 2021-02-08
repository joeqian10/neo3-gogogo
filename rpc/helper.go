package rpc

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
)

func PopInvokeStack(response InvokeResultResponse) (*models.InvokeStack, error) {
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return nil, fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return nil, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return nil, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	stack.Convert()
	return &stack, nil
}
