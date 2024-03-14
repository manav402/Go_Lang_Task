package questions

import (
	"fmt"
	"sync"
)

func zero(c chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Println("0")
		c <- true
		<-c
	}
}

func one(c chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		<-c
		fmt.Println("1")
		c <- true
	}
}

func Question1() {
	var c chan bool = make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go zero(c, &wg)
	go one(c, &wg)
	wg.Wait()
}
