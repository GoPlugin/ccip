package main

import (
	"fmt"
	"testing"

	"github.com/goplugin/pluginv3.0/core/scripts/ccip/revert-reason/handler"
)

func TestRevertReason(t *testing.T) {
	errorCodeString := "e1cd55090000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000008408c379a00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002645524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

	decodedError, err := handler.DecodeErrorStringFromABI(errorCodeString)
	if err != nil {
		fmt.Printf("Error decoding error string: %v\n", err)
		return
	}

	fmt.Println(decodedError)
}