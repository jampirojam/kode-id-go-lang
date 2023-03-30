package random

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var ch = make(chan string)

func RandomThread(coba, bisa []string) {
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go random(&wg, i, coba, ch)
		go random(&wg, i, bisa, ch)
	}

	for i := 1; i <= 8; i++ {
		fmt.Println(<-ch)
	}

	wg.Wait()
}

func random(wg *sync.WaitGroup, i int, data []string, ch chan string) {
	defer wg.Done()
	ch <- fmt.Sprintf("%v %d", data, i)
}
