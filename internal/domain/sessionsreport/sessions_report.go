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

type DayReport struct {
	Day      time.Time
	Sessions []session.Session
}

type ProjectReport struct {
	Project       string
	DurationByTag map[string]time.Duration
	TotalDuration time.Duration
}

type SessionsReport struct {
	Sessions []session.Session
}

func NewSessionsReport(sessions []session.Session) SessionsReport {
	return SessionsReport{Sessions: sessions}
}

func (s SessionsReport) Equals(report SessionsReport) bool {
	return reflect.DeepEqual(s.Sessions, report.Sessions)
}

func (s SessionsReport) GetByDayReport() []DayReport {
	dayReports := []DayReport{}
	sessionsByDay := s.splitSessionsByDay()
	for day, sessions := range sessionsByDay {
		dayReports = append(dayReports, DayReport{Day: day, Sessions: sessions})
	}
	return dayReports
}

func (s SessionsReport) GetByProjectReport() []ProjectReport {
	projectReports := []ProjectReport{}

	sessionsByProject := s.splitSessionsByProject()
	for project, sessions := range sessionsByProject {
		projectReports = append(projectReports, ProjectReport{
			Project:       project,
			DurationByTag: s.durationByTag(sessions),
			TotalDuration: s.duration(sessions),
		})
	}

	return projectReports
}

func (s SessionsReport) duration(sessions []session.Session) time.Duration {
	totalDuration := time.Second * 0
	for _, session := range sessions {
		totalDuration += session.Duration()
	}
	return totalDuration
}

func (s SessionsReport) totalDuration() time.Duration {
	return s.duration(s.Sessions)
}

func (s SessionsReport) splitSessionsByProject() map[string][]session.Session {
	projectsReport := make(map[string][]session.Session)
	for _, session := range s.Sessions {
		projectsReport[session.Project] = append(projectsReport[session.Project], session)
	}

	return projectsReport
}

func (s SessionsReport) splitSessionsByDay() map[time.Time][]session.Session {
	sessionMap := make(map[time.Time][]session.Session)

	for _, session := range s.Sessions {
		day := session.StartTime.Truncate(24 * time.Hour)
		sessionMap[day] = append(sessionMap[day], session)
	}

	return sessionMap
}

func (s SessionsReport) findUniqueTags(sessions []session.Session) []string {
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

func (s SessionsReport) durationByTag(sessions []session.Session) map[string]time.Duration {
	tags := s.findUniqueTags(sessions)

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
