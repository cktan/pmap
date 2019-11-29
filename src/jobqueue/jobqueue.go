package jobqueue

import (
	"sync"
	"sync/atomic"
)


type Item struct {
	idx int
	processItem  func(idx int)
}

type JobQueue struct {
	nworker int32
	backlog chan *Item
	waitGroup sync.WaitGroup
}

func New(nworker int) *JobQueue {
	jq := &JobQueue{ backlog: make(chan *Item, 100) }
	jq.SetNWorker(nworker)
	return jq
}

func (jq *JobQueue) Destroy() {
	close(jq.backlog)
	jq.waitGroup.Wait()
}

func (jq *JobQueue) NWorker() int {
	return int(jq.nworker)
}

func (jq *JobQueue) SetNWorker(n int) {
	if n < 0 {
		return
	}
	N := int32(n)
	for {
		k := atomic.LoadInt32(&jq.nworker)
		if k == N {
			break
		}
		if k < N {
			jq.addWorker()
		} else {
			jq.dropWorker()
		}
	}
}

func (jq *JobQueue) addWorker() {
	jq.waitGroup.Add(1)
	go jq.run()
}

func (jq *JobQueue) dropWorker() {
	n := atomic.LoadInt32(&jq.nworker) - 1
	if n > 0 {
		atomic.StoreInt32(&jq.nworker, n)
	}
}


func (jq *JobQueue) run() {
	id := atomic.AddInt32(&jq.nworker, 1) 
	for item := range jq.backlog {
		item.processItem(item.idx)
		n := atomic.LoadInt32(&jq.nworker)
		if (id > n) {
			// i am no longer needed
			break
		}
	}
	jq.waitGroup.Done()
}


func (jq *JobQueue) Add(processItem func(idx int), idx int) {
	item := &Item{idx, processItem}
	jq.backlog <- item
}


