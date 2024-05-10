package sessionsreport

import (
	"sort"
	"time"

	"github.com/TristanShz/flow/internal/domain/session"
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
	DurationByTag      map[string]time.Duration
	Project            string
	TotalDuration      time.Duration
	LastSessionEndTime time.Time
}

type SessionsReport struct {
	Sessions []session.Session
}

func NewSessionsReport(sessions []session.Session) SessionsReport {
	sort.Slice(sessions, func(i, j int) bool {
		return sessions[i].StartTime.Before(sessions[j].StartTime)
	})
	return SessionsReport{Sessions: sessions}
}

func (s SessionsReport) GetByDayReport() []DayReport {
	dayReports := []DayReport{}
	sessionsByDay := s.splitSessionsByDay()
	for day, sessions := range sessionsByDay {
		dayReports = append(dayReports, DayReport{Day: day, Sessions: sessions})
	}
	sort.Slice(dayReports, func(i, j int) bool {
		return dayReports[i].Day.Before(dayReports[j].Day)
	})
	return dayReports
}

func (s SessionsReport) GetByProjectReport() []ProjectReport {
	projectReports := []ProjectReport{}

	sessionsByProject := s.splitSessionsByProject()
	for project, sessions := range sessionsByProject {
		lastSession := sessions[len(sessions)-1]
		projectReports = append(projectReports, ProjectReport{
			Project:            project,
			DurationByTag:      s.durationByTag(sessions),
			TotalDuration:      s.duration(sessions),
			LastSessionEndTime: lastSession.EndTime,
		})
	}

	sort.Slice(projectReports, func(i, j int) bool {
		return projectReports[i].LastSessionEndTime.Before(projectReports[j].LastSessionEndTime)
	})

	return projectReports
}

func (s SessionsReport) duration(sessions []session.Session) time.Duration {
	totalDuration := time.Second * 0
	for _, session := range sessions {
		totalDuration += session.Duration()
	}
	return totalDuration
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
