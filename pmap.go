package pmap

/**
 *  Process N items using M go routines
 */
func pmap(processitem func(n int), N int, M int) {
	// notified when a go routine is done
	// must have N reserved to avoid potential race
	fin := make(chan int, N) 

	// the gate with M resources - controls #concurrent go routines at any time
	gate := make(chan int, M)
	defer close(fin)
	defer close(gate)

	// let maxworker run
	for i := 0; i < M; i++ {
		gate <- 1
	}

	// launch jobs for workers
	for i := 0; i < N; i++ {
		<-gate		// wait to launch 
		go func(idx int) {
			processitem(idx)
			gate <- 1 // let next guy run
			fin <- idx // notify done
		}(i)
	}

	// wait for all jobs to finish
	for i := 0; i < N; i++ {
		<-fin
	}
}
