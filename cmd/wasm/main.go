package main

import (
	"fmt"
	"github.com/carv-protocol/verifier/pkg/dcap"
	"syscall/js"
)

func execCallback(args []js.Value, value interface{}, e error) {
	var last = args[len(args)-1]
	if last.Type() == js.TypeFunction {
		last.Invoke(e, value)
	} else {
		fmt.Println("no callback")
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
		panic("length of args should be 5")
	}

	for index, arg := range args {
		if index <= 3 {
			if arg.Type() != js.TypeString {
				panic(fmt.Sprintf("argument %d should be a string", index))
			}
		} else {
			if arg.Type() != js.TypeFunction {
				panic("callback should be a function")
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
	if err != nil {
		panic(err)
	}
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
