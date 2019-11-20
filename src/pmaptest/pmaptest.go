package main

import (
	"pmap"
	"fmt"
	"math/rand"
	"time"
)


func work(k int) {
	x := rand.Intn(10)
	for i := 0; i < x; i++ {
		fmt.Println(k, "sleep", x - i)
		time.Sleep(time.Second)
	}
	fmt.Println(k, "wake")
}

func main() {
	pmap.Pmap(work, 20, 2)
}

