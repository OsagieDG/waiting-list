// main_wasm.go
package main

import (
	"syscall/js"
)

func main() {
	// Binds a Go function to JavaScript
	js.Global().Set("inputValidation", js.FuncOf(inputValidation))

	// Keeps the program running
	select {}
}

func inputValidation(this js.Value, args []js.Value) interface{} {
	name := args[0].String()
	email := args[1].String()

	if name == "" || email == "" {
		js.Global().Get("alert").Invoke("Name and email are required!")
		return nil
	}

	js.Global().Get("console").Call("log", "Name:", name, "Email:", email)
	return nil
}
