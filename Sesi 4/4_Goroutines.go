package main

import (
	"fmt"
)

func goroutine() {
	fmt.Println("Hello")
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i = ", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j = ", j)
	}
	fmt.Println("Second process func ended")
}

func main() {
	// go goroutine()

	// fmt.Println("Main execution started")
	// go firstProcess(8)
	// secondProcess(8)
	// fmt.Println("No. Of Goroutines : ", runtime.NumGoroutine())
	// fmt.Println("Main execution ended")

	// fmt.Println("Main execution started")
	// go firstProcess(8)
	// secondProcess(8)
	// fmt.Println("No. Of Goroutines : ", runtime.NumGoroutine())
	// time.Sleep(time.Second * 2)
	// fmt.Println("Main execution ended")
}