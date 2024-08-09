// main_wasm.go
package main

import (
	"syscall/js"
)

func main() {
	// Bind a Go function to JavaScript
	js.Global().Set("wasmFunction", js.FuncOf(wasmFunction))

	// Keep the program running
	select {}
}

func wasmFunction(this js.Value, args []js.Value) interface{} {
	name := args[0].String()
	email := args[1].String()

	if name == "" || email == "" {
		js.Global().Get("alert").Invoke("Name and email are required!")
		return nil
	}

	js.Global().Get("console").Call("log", "Name:", name, "Email:", email)
	return nil
}

