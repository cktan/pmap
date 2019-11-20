package main

import (
	"fmt"
	"math/rand"
	"os"
	"pmap"
	"time"
	"sync/atomic"
)

func run(numworkers int) {
	var status [200]int
	var active int32

	work := func(k int) {

		cur := atomic.AddInt32(&active, 1)
		fmt.Printf("%d start [%d active]\n", k, cur)
		x := rand.Intn(10)
		for i := 0; i < x; i++ {
			fmt.Println(k, "sleep", x-i)
			time.Sleep(time.Second)
		}
		status[k] = 999
		cur = atomic.AddInt32(&active, -1)
		fmt.Printf("%d fin [%d active]\n", k, cur)
	}

	pmap.Pmap(work, 200, numworkers)

	for i := 0; i < 200; i++ {
		if status[i] != 999 {
			fmt.Println("ERROR")
			os.Exit(1)
		}
	}
}

func main() {

	fmt.Println("numworkers 100")
	run(100)
	fmt.Println("-------------------------")
	fmt.Println()
	
	fmt.Println("numworkers 50")
	run(50)
	fmt.Println("-------------------------")
	fmt.Println()
	
	fmt.Println("numworkers 20")
	run(20)
	fmt.Println("-------------------------")
	fmt.Println()
	
}
