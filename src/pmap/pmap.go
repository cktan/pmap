package pmap

import "sync"

/**
 *  Process N items using M go routines
 */
func Pmap(processItem func(n int), N int, M int) {
	if M > N {
		M = N
	}

	wg := &sync.WaitGroup{}
	ticket := make(chan int, 10)

	// let M workers run concurrently
	wg.Add(M)
	for i := 0; i < M; i++ {
		go func() {
			for idx := range ticket {
				processItem(idx)
			}
			wg.Done()
		}()
	}

	// send the jobs
	for i := 0; i < N; i++ {
		ticket <- i
	}
	close(ticket)

	// wait for all jobs to finish
	wg.Wait()
}
