package rpc

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
)

func PopInvokeStacks(response InvokeResultResponse) ([]models.InvokeStack, error) {
	if response.HasError() {
		return nil, fmt.Errorf(response.GetErrorInfo())
	}
	result := response.Result
	if result.State == "FAULT" {
		msg := "engine faulted"
		if len(result.Exception) != 0 {
			msg += ", exception: " + response.Result.Exception
		}
		return nil, fmt.Errorf(msg)
	}
	// json["stack"] = "error: invalid operation"
	if result.Stack == nil {
		return []models.InvokeStack{}, fmt.Errorf("error: invalid operation")
	}

	return result.Stack, nil
}
