package session

import (
	"fmt"
	"strings"
	"time"
)

const (
	FlowingStatus = "FLOWING"
	EndedStatus   = "ENDED"
)

type Session struct {
	Id        string
	StartTime time.Time
	EndTime   time.Time
	Project   string
	Tags      []string
}

func (s Session) GetFormattedStartTime() string {
	return s.StartTime.Format(time.DateTime)
}

func (s Session) GetFormattedEndTime() string {
	if s.EndTime.IsZero() {
		return "/"
	}

	return s.EndTime.Format(time.DateTime)
}

func (s Session) Duration() time.Duration {
	if s.EndTime.IsZero() {
		return 0
	}
	return s.EndTime.Sub(s.StartTime).Round(time.Second)
}

func (s Session) PrettyString() string {
	return fmt.Sprintf("StartTime: %s\n EndTime: %s\n Project: %s\n Tags: %s\n",
		s.GetFormattedStartTime(), s.GetFormattedEndTime(), s.Project, strings.Join(s.Tags, ", "))
}

func (s Session) Status() string {
	if s.EndTime.IsZero() {
		return FlowingStatus
	}

	return EndedStatus
}

func (s Session) Equals(session Session) bool {
	return s.StartTime.Equal(session.StartTime)
}
