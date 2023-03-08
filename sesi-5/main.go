package main

import (
	"fmt"
	"sync"
	"time"
)

// var heroes = []string{}

func main() {
	// fmt.Println("func main started")
	//
	// go firstProcess()
	//
	// go secondProcess()
	//
	// fmt.Println("Num of gorountines:", runtime.NumGoroutine())
	//
	// time.Sleep(1 * time.Second)
	//
	// fmt.Println("func main ended")

	// var wg sync.WaitGroup
	// var mutex sync.Mutex
	//
	// heroes1 := []string{"Batman", "Superman", "Wonder Woman", "Flash", "Green Arrow"}
	//
	// for _, hero := range heroes1 {
	// 	wg.Add(1)
	//
	// 	go func(value string) {
	// 		mutex.Lock()
	// 		heroes = append(heroes, value)
	// 		mutex.Unlock()
	//
	// 		wg.Done()
	// 	}(hero)
	// }
	//
	// wg.Wait()
	//
	// fmt.Printf("Heroes %#v\n", heroes)
	const loops int = 5

	ch := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= loops; i++ {
		wg.Add(1)

		go func(num int) {
			fmt.Println("Start sending data:", num)
			ch <- num
			fmt.Println("After sending data:", num)
			wg.Done()
		}(i)
	}

	go func() {
		fmt.Println("Wait")
		wg.Wait()
		close(ch)
		fmt.Println("Close channel")
	}()

	for c := range ch {
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", c)
	}
}
