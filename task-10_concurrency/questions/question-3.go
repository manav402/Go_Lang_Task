package questions

import "fmt"

func fizz(fbtofizz, fizztobuzz chan int, output chan bool) {
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

func buzz(fizztobuzz chan int, output chan bool) {
	for {
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

func fizzbuzz(input, fbtofizz chan int, output chan bool) {
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

func Question3() {
	var fbtofizz chan int = make(chan int)
	var fizztobuzz chan int = make(chan int)
	var input chan int = make(chan int)
	var output chan bool = make(chan bool)

	go fizz(fbtofizz, fizztobuzz, output)
	go buzz(fizztobuzz, output)
	go fizzbuzz(input, fbtofizz, output)

	for i := 1; i <= 30; i++ {
		input <- i
		x := <-output
		if !x {
			fmt.Println(i)
		}
	}
}
