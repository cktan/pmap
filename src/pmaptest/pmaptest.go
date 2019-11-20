package main

import (
	"pmap"
	"fmt"
	"math/rand"
	"time"
	"os"
)



func main() {
	
	var status [20]int
	
	work := func(k int) {
		fmt.Println(k, "start")
		x := rand.Intn(10)
		for i := 0; i < x; i++ {
			fmt.Println(k, "sleep", x - i)
			time.Sleep(time.Second)
		}
		fmt.Println(k, "fin")
		status[k] = 999
	}
	pmap.Pmap(work, 20, 2)

	for i := 0; i < 20; i++ {
		if status[i] != 999 {
			fmt.Println("ERROR");
			os.Exit(1)
		}
	}
}

