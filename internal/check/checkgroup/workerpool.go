package checkgroup

import "context"

type (
	PoolOption func(*workerPool)
	ctxKey     string
)

const poolCtxKey ctxKey = "pool"

// WithPool returns a new context that contains the pool. The pool will be used by the checkgroup and the binary operators (or, and) when spawning subchecks.
func WithPool(ctx context.Context, pool Pool) context.Context {
	return context.WithValue(ctx, poolCtxKey, pool)
}

// PoolFromContext returns the pool from the context, or a pool that does not
// limit the number of parallel jobs if none found.
func PoolFromContext(ctx context.Context) Pool {
	if p, ok := ctx.Value(poolCtxKey).(*workerPool); !ok {
		return new(limitlessPool)
	} else {
		return p
	}
}

// NewPool creates a new worker pool. With no options, this yields a pool with
// exactly one worker, meaning that all tasks that are added will run
// sequentially.
func NewPool(opts ...PoolOption) Pool {
	pool := &workerPool{
		numWorkers: 1,
	}
	for _, opt := range opts {
		opt(pool)
	}

	pool.jobs = make(chan func(), pool.numWorkers)
	for i := 0; i < pool.numWorkers; i++ {
		go worker(pool.jobs)
	}

	if pool.ctx != nil {
		go func() {
			<-pool.ctx.Done()
			close(pool.jobs)
		}()
	}

	return pool
}

func worker(jobs <-chan func()) {
	for job := range jobs {
		job()
	}
}

func WithWorkers(count int) PoolOption {
	return func(p *workerPool) { p.numWorkers = count }
}
func WithContext(ctx context.Context) PoolOption {
	return func(p *workerPool) { p.ctx = ctx }
}

// Add adds the function to the pool and schedules it. The function will only be
// run if there is a free worker available in the pool, thus limiting the
// concurrent workloads in flight.
func (p *workerPool) Add(check func()) {
	p.jobs <- check
}

// Add on a limitless pool just runs the function in a go routine.
func (p *limitlessPool) Add(check func()) {
	go check()
}
