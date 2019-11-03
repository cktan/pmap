package pmap

func pmap(processitem func(idx int), maxidx int, maxworker int) {
	// notified when a go routine is done
	// must have maxidx reserved to avoid potential race
	fin := make(chan int, maxidx) 

	// the gate - controls #concurrent go routines at any time
	gate := make(chan int, maxworker)
	defer close(fin)
	defer close(gate)

	// let maxworker run
	for i := 0; i < maxworker; i++ {
		gate <- 1
	}
	for i := 0; i < maxidx; i++ {
		<-gate		// wait to launch 
		go func(idx int) {
			processitem(idx)
			gate <- 1 // let next guy run
			fin <- idx // notify done
		}(i)
	}

	// wait for all jobs to finish
	for i := 0; i < maxidx; i++ {
		<-fin
	}
}
