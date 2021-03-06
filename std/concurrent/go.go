package concurrent

import (
	"log"
	"time"
)

func goRun(i int) {
	log.Println(i)
}
func go1() {
	for i := 0; i < 10; i++ {
		go goRun(i)
	}
	time.Sleep(time.Second)
}

func go2() {
	for i := 0; i < 10; i++ {
		go func() {
			log.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
