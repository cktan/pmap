package main

import (
	"fmt"
	"math/rand"
	"os"
	"pmap"
	"time"
)

func main() {

	var status [200]int

	work := func(k int) {
		fmt.Println(k, "start")
		x := rand.Intn(10)
		for i := 0; i < x; i++ {
			fmt.Println(k, "sleep", x-i)
			time.Sleep(time.Second)
		}
		fmt.Println(k, "fin")
		status[k] = 999
	}
	pmap.Pmap(work, 200, 20)

	for i := 0; i < 200; i++ {
		if status[i] != 999 {
			fmt.Println("ERROR")
			os.Exit(1)
		}
	}
}
