package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func goRoutinesExec(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	sleepTime := time.Duration(rand.Intn(5)) * time.Second
	fmt.Printf("\ngoroutine número %d "+"dormiu por "+sleepTime.String(), i+1)
	time.Sleep(sleepTime)
	fmt.Println("\n------------------------------------")
	fmt.Printf("goroutine número %d "+"acordou!", i+1)
}

func main() {
	var numGoroutines int
	fmt.Print("\nQuantas goroutines deverão ser criadas? ")
	fmt.Scanf("%d", &numGoroutines)

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(n int) {
			goRoutinesExec(n, &wg)
		}(i)
	}

	wg.Wait()
	fmt.Println("\n------------------------------------")
	fmt.Printf("\nn: %d\n", numGoroutines)
}
