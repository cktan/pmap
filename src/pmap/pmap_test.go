package pmap

import (
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func test1_run(numworkers int, t *testing.T) {
	var status [200]int
	var active int32

	work := func(k int) {

		cur := atomic.AddInt32(&active, 1)
		x := rand.Intn(10)
		t.Logf("%d start sleep %d [%d active]\n", k, x, cur)
		time.Sleep(time.Duration(x*100) * time.Millisecond)
		status[k] = 999
		cur = atomic.AddInt32(&active, -1)
		t.Logf("%d fin [%d active]\n", k, cur)
	}

	Pmap(work, 200, numworkers)

	for i := 0; i < 200; i++ {
		if status[i] != 999 {
			t.Errorf("Some job is not done yet")
		}
	}
}

func TestPmap(t *testing.T) {
	test1_run(100, t)
	test1_run(20, t)
}
