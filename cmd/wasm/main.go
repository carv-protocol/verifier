package main

import (
	"errors"
	"fmt"
	"syscall/js"

	"github.com/carv-protocol/verifier/pkg/dcap"
)

func execCallback(args []js.Value, value interface{}, err error) {
	var last = args[len(args)-1]
	if last.Type() == js.TypeFunction {
		if err != nil {
			last.Invoke(err.Error(), value)
		} else {
			last.Invoke(nil, value)
		}
	} else {
		fmt.Println("[Go:execCallback] no callback")
	}
}

func register(this js.Value, args []js.Value) interface{} {
	var scope = args[0]
	scope.Set("verifyAttestation", js.FuncOf(VerifyAttestation))
	execCallback(args, scope, nil)
	return nil
}

func VerifyAttestation(this js.Value, args []js.Value) interface{} {
	if len(args) != 5 {
		execCallback(args, nil, errors.New("[Go:VerifyAttestation] length of args should be 5"))
		return nil
	}

	for index, arg := range args {
		if index <= 3 {
			if arg.Type() != js.TypeString {
				execCallback(args, nil, errors.New(fmt.Sprintf("[Go:VerifyAttestation] argument %d should be a string", index)))
				return nil
			}
		} else {
			if arg.Type() != js.TypeFunction {
				execCallback(args, nil, errors.New("[Go:VerifyAttestation] callback should be a function"))
				return nil
			}
		}
	}

	var (
		message      = args[0].String()
		identityJson = args[1].String()
		tcbJson      = args[2].String()
		trustedJson  = args[3].String()
	)

	result, err := dcap.VerifyAttestation(message, identityJson, tcbJson, trustedJson)
	println(args, result)
	execCallback(args, result, err)
	return nil
}
func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("__wasm_main__", js.FuncOf(register))
	fmt.Println("WASM Go Initialized")
	<-done
}
