package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getRating(wg *sync.WaitGroup, channel chan int) {

	max := 10
	min := 1

	responseTime := rand.Intn(max-min+1) + min
	time.Sleep((time.Duration(responseTime)) * time.Second)
	channel <- rand.Intn(max-min+1) + min
	wg.Done()
}

func main() {

	wg := &sync.WaitGroup{}
	channel := make(chan int)

	n := 200
	for i := 0; i < n; i++ {
		wg.Add(1)
		go getRating(wg, channel)
	}
	sum := 0

	go func(channel chan int, n int) {

		for i := 0; i < n; i++ {
			sum = <-channel + sum
		}
	}(channel, n)

	wg.Wait()
	// fmt.Println(sum)
	var averageRating float32
	averageRating = float32(sum) / float32(n)
	fmt.Println("averageRating", averageRating)
}
