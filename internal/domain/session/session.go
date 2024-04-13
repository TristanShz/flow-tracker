package session

import (
	"fmt"
	"strings"
	"time"
)

type Session struct {
	StartTime time.Time
	EndTime   time.Time
	Project   string
	Tags      []string
}

func (s Session) Duration() time.Duration {
	return s.EndTime.Sub(s.StartTime).Round(time.Second)
}

func (s Session) PrettyString() string {
	return fmt.Sprintf("Session:\n  StartTime: %s\n  EndTime: %s\n  Project: %s\n  Tags: %s\n",
		s.StartTime, s.EndTime, s.Project, strings.Join(s.Tags, ", "))
}
