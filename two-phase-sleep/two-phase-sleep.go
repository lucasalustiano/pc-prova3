package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func goRoutinesExec(i int, dormidasSegundaFase []int) {
	sleepTime := time.Duration(rand.Intn(5)) * time.Second
	fmt.Printf("goroutine número %d "+"dormiu por "+sleepTime.String(), i+1)
	fmt.Println("\n------------------------------------")
	time.Sleep(sleepTime)

	fmt.Printf("goroutine número %d "+"acordou!", i+1)
	fmt.Println("\n------------------------------------")
	goRoutinesCount = goRoutinesCount - 1
	sleepTimePhase2 := rand.Intn(10)
	dormidasSegundaFase[i] = sleepTimePhase2

	fmt.Printf("goroutine número %d "+"escolheu %ds para segunda fase!", i+1, sleepTimePhase2)
	fmt.Println("\n------------------------------------")

	if goRoutinesCount != -1 {
		locker1.Done() // decrementa, quando a última decrementar, o contador vira zero, ai todo mundo acorda.
		locker1.Wait() // espera.
	}

	fmt.Printf("goroutine número %d "+"começou a segunda fase!", i+1)

	if i == 0 { // primeira goroutine
		sleepTime = time.Duration(dormidasSegundaFase[len(dormidasSegundaFase)-1]) * time.Second
		fmt.Printf("\ngoroutine número %d "+"dormiu por "+sleepTime.String(), i+1)
		fmt.Println("\n------------------------------------")
		time.Sleep(sleepTime)
	} else { // todas as outras goroutines
		sleepTime = time.Duration(dormidasSegundaFase[i-1]) * time.Second
		fmt.Printf("\ngoroutine número %d "+"dormiu por "+sleepTime.String(), i+1)
		fmt.Println("\n------------------------------------")
		time.Sleep(sleepTime)
	}

	defer wg.Done() // libera mainha
}

var dormidasSegundaFase []int
var wg, locker1 sync.WaitGroup
var numGoroutines, goRoutinesCount int

func main() {

	fmt.Print("\nQuantas goroutines deverão ser criadas? ")
	fmt.Scanf("%d", &numGoroutines)
	fmt.Print("\n")

	goRoutinesCount = numGoroutines
	dormidasSegundaFase = make([]int, numGoroutines)

	locker1.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(n int) {
			goRoutinesExec(n, dormidasSegundaFase)
		}(i)
	}

	wg.Wait()
	fmt.Printf("\nn: %d\n", numGoroutines)
}
