question -1:How would you design a concurrent Go program that prints alternating 0s and 1s indefinitely? The program should consist of two goroutines, a and b, each responsible for printing either a 0 or a 1. The main function should start both goroutines and ensure that they alternate printing 0s and 1s continuously.

question -2:https://goplay.tools/snippet/zd4DwTngeJu - reduce this code with concurrency
    answer :https://goplay.tools/snippet/9jAsuBymMOL

question -3: Write a program that prints the numbers from 1 to 30. But for multiples of three, print "Fizz" instead of the number, and for the multiples of five, print "Buzz". For numbers that are multiples of both three and five, print "FizzBuzz".Write three go routines for fizz, buzz, and fizzBuzz, and use channels to communicate between go routines. The output should be synchronized
