package main

import (
	"log"
	"time"
)

func main() {
	n := 10000000
	times := 500

	sTForLoop(n, times)
	chanForLoop(n, times)
}

func singleThreadedFor(n int) int {
	counter := 0
	for i := 0; i < n; i++ {
		counter = counter + i
	}

	return counter
}

func chanFor(c chan int, n int) {
	go func() {
		for {
			c <- singleThreadedFor(n)
		}
	}()
}

func sTForLoop(n, times int) int {
	defer timeTrack(time.Now(), "singleThreadedFor")
	res := 0
	for i := 0; i < times; i++ {
		res = res + singleThreadedFor(n)
	}
	log.Println("result is ", res)
	return res
}

func chanForLoop(n, times int) {
	defer timeTrack(time.Now(), "chanFor")

	counterChan := make(chan (int))
	for i := 0; i < times; i++ {
		chanFor(counterChan, n)
	}
	res := 0
	for i := 0; i < times; i++ {
		res = res + <-counterChan
	}

	log.Println("result is ", res)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
