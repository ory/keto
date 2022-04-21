package hlc

import "time"

// PT is physical time, it uses time.Now().UnixNano().
// it's the default PT source in HLC.
type PT struct {
	// if true it'll use time.Now().Unix()
	Seconds bool
}

// Now returns the value of
func (p PT) Now() int64 {
	if p.Seconds {
		return time.Now().Unix()
	}
	return time.Now().UnixNano()
}
