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
				idx, ok := <-ticket
				if !ok {
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
		close(ticket)
	}()

	// wait for all jobs to finish
	for i := 0; i < N; i++ {
		<-fin
	}

	// Go Channel Closing Principle:
	//  - Don't close a channel from the receiver side and
	//  - Don't close a channel if the channel has multiple concurrent senders.

	// At this point, we gave out N tickets, and we received N fins.
	// All workers must be blocked waiting for ticket at this time.
	// They will next get the term signal and exit.
	
	// Even though we are receiver on fin, we can close it making sure
	// sure that no one will ever send to it again.
	close(fin)
}
