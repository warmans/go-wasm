package main

import "syscall/js"

func main() {
	el := js.Global().Get("document").Call("getElementById", "output")
	el.Set("innerHTML", "Hello World")
}
