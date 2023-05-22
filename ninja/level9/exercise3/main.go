package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incremneter := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incremneter
			runtime.Gosched()
			v++
			incremneter = v
			fmt.Println(incremneter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incremneter)
}
