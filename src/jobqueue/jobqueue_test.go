package jobqueue

import (
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func TestJobQueue(t *testing.T) {
	N := 20
	status := make([]int, N)
	var active int32

	work := func(k int) {
		x := rand.Intn(10)
		cur := atomic.AddInt32(&active, 1)
		t.Logf("%d start sleep %d [%d active]\n", k, x, cur)
		time.Sleep(time.Duration(x*100) * time.Millisecond)
		status[k] = 999
		cur = atomic.AddInt32(&active, -1)
		t.Logf("%d fin [%d active]\n", k, cur)
	}

	jq := New(5)
	for i := 0; i < N; i++ {
		jq.Add(work, i)
	}

	jq.Destroy()

	for i := 0; i < N; i++ {
		if status[i] != 999 {
			t.Errorf("Some job is not done yet")
		}
	}
}
