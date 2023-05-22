package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	incremneter := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			m.Lock()
			v := incremneter
			v++
			incremneter = v
			fmt.Println(incremneter)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incremneter)
}
