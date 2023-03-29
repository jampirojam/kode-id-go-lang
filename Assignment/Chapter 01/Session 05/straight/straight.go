package straight

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var rwm sync.RWMutex
var chFinal = make(chan string, 2)
var chCoba = make(chan string)
var chBisa = make(chan string)

func StraightThread(coba, bisa []string) {
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go randomCoba(&wg, i, coba, &chCoba)
		go randomBisa(&wg, i, bisa, &chBisa)
	}

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go straight(&wg, &rwm, &chCoba, &chBisa, &chFinal)
	}

	for i := 1; i <= 8; i++ {
		fmt.Println(<-chFinal)
	}

	go func() {
		wg.Wait()
		close(chFinal)
	}()
}

func randomCoba(wg *sync.WaitGroup, i int, coba []string, chCoba *chan string) {
	defer wg.Done()
	*chCoba <- fmt.Sprintf("%v %d", coba, i)
}

func randomBisa(wg *sync.WaitGroup, i int, bisa []string, chBisa *chan string) {
	defer wg.Done()
	*chBisa <- fmt.Sprintf("%v %d", bisa, i)
}

func straight(wg *sync.WaitGroup, rwm *sync.RWMutex, chCoba *chan string, chBisa *chan string, chFinal *chan string) {
	defer wg.Done()
	rwm.Lock()
	*chFinal<-<-*chCoba
	*chFinal<-<-*chBisa
	rwm.Unlock()
}
