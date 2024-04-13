package session

import (
	"time"
)

type Session struct {
	StartTime time.Time
	EndTime   time.Time
	Project   string
	Tags      []string
}

func (e Session) Duration() time.Duration {
	return e.EndTime.Sub(e.StartTime).Round(time.Second)
}
