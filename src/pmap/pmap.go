package pmap


/**
 *  Process N items using M go routines
 */
func Pmap(processItem func(n int), N int, M int) {
	if M > N {
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

	// let M workers run concurrently
	for i := 0; i < M; i++ {
		go func() {
			for {
				idx := <-ticket
				if idx == -1 {
					return
				}
				processItem(idx)
				fin <- idx
			}
		}()
	}

	// send the jobs. Do this in a go routine so we don't have a
	// race between ticket and fin
	go func() {
		for i := 0; i < N; i++ {
			ticket <- i
		}
	}()

	// wait for all jobs to finish
	for i := 0; i < N; i++ {
		<-fin
	}
	
	// safe to close fin here because all workers must be blocked
	// waiting on ticket at this time
	close(fin)
	
	// send the terminate signal. each worker will get the term
	// signal and MUST NOT send to fin (it was closed above).
	for i := 0; i < M; i++ {
		ticket <- -1
	}

	close(ticket)
}
