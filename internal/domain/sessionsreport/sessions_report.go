package sessionsreport

import (
	"reflect"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
)

type SessionsReport struct {
	Sessions []session.Session
}

func (s SessionsReport) Equals(report SessionsReport) bool {
	return reflect.DeepEqual(s, report)
}

func (s SessionsReport) TotalDuration() time.Duration {
	totalDuration := time.Second * 0
	for _, session := range s.Sessions {
		totalDuration += session.Duration()
	}

	return totalDuration
}

func (s SessionsReport) ProjectsReport() map[string]time.Duration {
	projectsReport := make(map[string]time.Duration)
	for _, session := range s.Sessions {
		_, ok := projectsReport[session.Project]

		if ok {
			projectsReport[session.Project] += session.Duration()
		} else {
			projectsReport[session.Project] = session.Duration()
		}
	}

	return projectsReport
}

func (s SessionsReport) SplitSessionsByDay() map[time.Time][]session.Session {
	sessionMap := make(map[time.Time][]session.Session)

	for _, session := range s.Sessions {
		day := session.StartTime.Truncate(24 * time.Hour)
		sessionMap[day] = append(sessionMap[day], session)
	}

	return sessionMap
}
