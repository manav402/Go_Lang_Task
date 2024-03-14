package main

import (
	"fmt"
	"sync"
)

// Write a Go program that is concurrently executed using two goroutines, a() and b(). a() prints 0 when the integer pointer is even, and b() prints 1 when the integer pointer it receives is odd. Both functions sleep for 100 milliseconds between iterations. The main function launches these goroutines and waits for 1 second before exiting.

func zero(c chan bool,wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Println("0")
		c <- true
		<-c
	}
}

func one(c chan bool,wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		<-c
		fmt.Println("1")
		c <- true
	}
}

func qution1() {
	var c chan bool = make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go zero(c,&wg)
	go one(c,&wg)
	wg.Wait()
}

// Write a program that prints the numbers from 1 to 30. But for multiples of three, print "Fizz" instead of the number, and for the multiples of five, print "Buzz". For numbers that are multiples of both three and five, print "FizzBuzz".
// Write three go routines for fizz, buzz, and fizzBuzz, and use channels to communicate between go routines. The output should be synchronized

var fbtofizz chan int = make(chan int)
var fizztobuzz chan int = make(chan int)
var input chan int = make(chan int)
var output chan bool = make(chan bool)

func fizz() {
	for {
		select {
		case v := <-fbtofizz:
			if v%3 == 0 {
				fmt.Println("fizz")
				output <- true
			} else {
				fizztobuzz <- v
			}
		}
	}
}

func buzz() {
	for  {
		select {
		case v := <-fizztobuzz:
			if v%5 == 0 {
				fmt.Println("buzz")
				output <- true
			} else {
				output <- false
			}
		}
	}
}

func fizzbuzz() {
	for {
		select {
		case v := <-input:
			if v%15 == 0 {
				fmt.Println("fizzbuzz")
				output <- true
			} else {
				fbtofizz <- v
			}
		}
	}
}

func qution2() {
	go fizz()
	go buzz()
	go fizzbuzz()

	for i := 1; i <= 30; i++ {
		input <- i
		x := <-output
		if !x {
			fmt.Println(i)
		}
	}
}

func main(){
	qution1()
	qution2()
}
