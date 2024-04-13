package entry

import (
	"time"
)

type Entry struct {
	StartTime time.Time
	EndTime   time.Time
	Project   string
	Tags      []string
}

func (e Entry) Duration() time.Duration {
	return e.EndTime.Sub(e.StartTime).Round(time.Second)
}
