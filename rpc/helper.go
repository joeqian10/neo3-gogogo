package rpc

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
)

func PopInvokeStack(response InvokeResultResponse) (*models.InvokeStack, error) {
	if response.HasError() {
		return nil, fmt.Errorf(response.GetErrorInfo())
	}
	if response.Result.State == "FAULT" {
		msg := "engine faulted"
		if len(response.Result.Exception) != 0 {
			msg += ", exception: " + response.Result.Exception
		}
		return nil, fmt.Errorf(msg)
	}
	if len(response.Result.Stack) == 0 {
		return nil, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	stack.Convert()
	return &stack, nil
}
