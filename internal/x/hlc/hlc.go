package hlc

import "fmt"

type Clock interface {
	Now() int64
}

type HLC struct {
	clock Clock
	t     Timestamp
}

type Timestamp struct {
	ts    int64
	count int64
}

// Less compares timestamps by checking the timestamps first
// if they're equal then it uses the counters to break the tie.
// it returns true if incoming is less than t
func (t Timestamp) Less(incoming Timestamp) bool {
	return t.ts < incoming.ts || (t.ts == incoming.ts && t.count < incoming.count)
}

func (t Timestamp) String() string {
	return fmt.Sprintf("TS=%d Count=%d", t.ts, t.count)

}

func New() *HLC {
	return &HLC{t: Timestamp{}, clock: PT{}}
}

func NewWithPT(pt Clock) *HLC {
	return &HLC{t: Timestamp{}, clock: pt}
}

// Now should be called when sending an event or when a local event happens
func (h *HLC) Now() Timestamp {
	t := h.t
	h.t.ts = h.max(t.ts, h.clock.Now())

	if h.t.ts != t.ts {
		h.t.count = 0
		return h.t
	}

	h.t.count++
	return h.t
}

// Update should be called when receive an event
func (h *HLC) Update(incoming Timestamp) {
	t := h.t
	h.t.ts = h.max(t.ts, incoming.ts, h.clock.Now())

	if h.t.ts == t.ts && incoming.ts == t.ts {
		h.t.count = h.max(t.count, incoming.count) + 1
	} else if t.ts == h.t.ts {
		h.t.count++
	} else if h.t.ts == incoming.ts {
		h.t.count = incoming.count + 1
	} else {
		h.t.count = 0
	}
}

func (h *HLC) max(vals ...int64) int64 {
	if len(vals) == 0 {
		return 0
	}

	m := vals[0]
	for i := range vals {
		if vals[i] > m {
			m = vals[i]
		}
	}
	return m
}
