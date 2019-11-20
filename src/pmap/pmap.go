package pmap


/**
 *  Process N items using M go routines
 */
func Pmap(processItem func(n int), N int, M int) {
	if (M > N) {
		M = N
	}

	var fin chan int
	var ticket chan int
	if false {
		// for debug
		fin = make(chan int)
		ticket = make(chan int)
	} else {
		fin = make(chan int, 10)
		ticket = make(chan int, 10)
	} 
	defer func() {
		close(fin)
		close(ticket)
	}()

	// let M go routines run concurrently
	for i := 0; i < M; i++ {
		go func() {
			for {
				idx := <- ticket
				if idx == -1 {
					return
				}
				processItem(idx)
				fin <- idx
			}
		}()
	}

	// send the jobs
	go func() {
		for i := 0; i < N; i++ {
			ticket <- i
		}
	}()
	
	// wait for all jobs to finish
	for i := 0; i < N; i++ {
		<-fin
	}

	// send the terminate signal
	for i := 0; i < M; i++ {
		ticket <- -1
	}
}

