package sessionsreport

import (
	"reflect"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
)

const (
	FormatByDay     = "by-day"
	FormatByProject = "by-project"
)

type SessionsReport struct {
	Sessions []session.Session
}

func NewSessionsReport(sessions []session.Session) SessionsReport {
	return SessionsReport{Sessions: sessions}
}

func (s SessionsReport) Equals(report SessionsReport) bool {
	return reflect.DeepEqual(s.Sessions, report.Sessions)
}

func (s SessionsReport) Duration(sessions []session.Session) time.Duration {
	totalDuration := time.Second * 0
	for _, session := range sessions {
		totalDuration += session.Duration()
	}
	return totalDuration
}

func (s SessionsReport) TotalDuration() time.Duration {
	return s.Duration(s.Sessions)
}

func (s SessionsReport) SplitSessionsByProject() map[string][]session.Session {
	projectsReport := make(map[string][]session.Session)
	for _, session := range s.Sessions {
		projectsReport[session.Project] = append(projectsReport[session.Project], session)
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

func (s SessionsReport) FindUniqueTags(sessions []session.Session) []string {
	tags := make(map[string]bool)
	for _, session := range sessions {
		for _, tag := range session.Tags {
			tags[tag] = true
		}
	}
	uniqueTags := make([]string, 0, len(tags))
	for tag := range tags {
		uniqueTags = append(uniqueTags, tag)
	}
	return uniqueTags
}

func (s SessionsReport) DurationByTag(sessions []session.Session) map[string]time.Duration {
	tags := s.FindUniqueTags(sessions)

	tagsDuration := make(map[string]time.Duration)
	for _, tag := range tags {
		tagDuration := time.Second * 0
		for _, session := range sessions {
			if session.HasTag(tag) {
				tagDuration += session.Duration()
			}
		}
		tagsDuration[tag] = tagDuration
	}

	return tagsDuration
}
