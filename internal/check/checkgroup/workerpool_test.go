package checkgroup_test

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ory/keto/internal/check/checkgroup"
)

func TestPool(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numWorkers := 5
	p := checkgroup.NewPool(
		checkgroup.WithWorkers(numWorkers),
		checkgroup.WithContext(ctx),
	)

	var (
		jobsCount int32
		wg        sync.WaitGroup
	)

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		p.Add(func() {
			defer wg.Done()
			if jobs := atomic.AddInt32(&jobsCount, 1); jobs > int32(numWorkers) {
				t.Errorf("%d jobs in flight, more than %d", jobs, numWorkers)
			}
			time.Sleep(1 * time.Millisecond)
			atomic.AddInt32(&jobsCount, -1)
		})
	}
	wg.Wait()
}
