package app

import (
	"fmt"
	"ojam-test/c3/s1/app/service"
	"os"
	"time"
)

func Boot() {
	i := 0
	lastPostId := 101 // Actual PoisId have maximum 100, so we set 101 for the first time
	for {
		i+=1
		fmt.Printf("\nTask Weather Check %d\n", i)
		service.WeatherCheck()
		fmt.Printf("\nTask Post %d\n", i)
		lastPostId = service.PostStatus(lastPostId)
		fmt.Printf("\n========================================")
		time.Sleep(2 * time.Second)
		lastPostId+=1

		if i == 100 {
			os.Exit(1)
		}
	}
}