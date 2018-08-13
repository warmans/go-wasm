package main

import (
	"syscall/js"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	el := js.Global().Get("document").Call("getElementById", "output")

	wg.Add(1)
	go func() {
		defer wg.Done()
		counter := 0
		for {
			counter++
			el.Set("innerHTML", fmt.Sprintf("Hello World %d", counter))
			if counter == 10 {
				return
			}
			time.Sleep(time.Second)
		}

	}()

	wg.Wait()
	el.Set("innerHTML", "Hello World complete!")
}
