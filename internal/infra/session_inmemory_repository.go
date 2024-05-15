package infra

import (
	"slices"

	"github.com/TristanShz/flow/internal/application"
	"github.com/TristanShz/flow/internal/domain/session"
	"github.com/TristanShz/flow/pkg/timerange"
)

type InMemorySessionRepository struct {
	Sessions []session.Session
}

func (r *InMemorySessionRepository) FindById(id string) *session.Session {
	for _, session := range r.Sessions {
		if session.Id == id {
			return &session
		}
	}
	return nil
}

func (r *InMemorySessionRepository) Save(s session.Session) error {
	startedSessionIndex := slices.IndexFunc(r.Sessions, func(s session.Session) bool {
		return s.StartTime.Equal(s.StartTime)
	})

	if startedSessionIndex == -1 {
		r.Sessions = append(r.Sessions, s)
	} else {
		r.Sessions[startedSessionIndex] = s
	}

	return nil
}

func (r *InMemorySessionRepository) Delete(id string) error {
	sessionIndex := slices.IndexFunc(r.Sessions, func(s session.Session) bool {
		return s.Id == id
	})
	if sessionIndex == -1 {
		return nil
	}
	r.Sessions = append(r.Sessions[:sessionIndex], r.Sessions[sessionIndex+1:]...)
	return nil
}

func (r *InMemorySessionRepository) FindLastSession() *session.Session {
	if len(r.Sessions) == 0 {
		return nil
	}

	return &r.Sessions[len(r.Sessions)-1]
}

func (r *InMemorySessionRepository) FindAllSessions(filters *application.SessionsFilters) []session.Session {
	filteredSessions := r.Sessions

	if filters != nil {
		if !filters.Timerange.IsZero() {
			filteredSessions = r.filterByTimeRange(filteredSessions, filters.Timerange)
		}

		if filters.Project != "" {
			filteredSessions = r.filterByProject(filteredSessions, filters.Project)
		}
	}

	return filteredSessions
}

func (r *InMemorySessionRepository) FindAllProjects() []string {
	sessions := r.FindAllSessions(nil)

	projects := []string{}

	for _, session := range sessions {
		if slices.Contains(projects, session.Project) {
			continue
		}

		projects = append(projects, session.Project)
	}

	return projects
}

func (r *InMemorySessionRepository) FindAllProjectTags(project string) []string {
	sessionsForProject := r.FindAllSessions(&application.SessionsFilters{Project: project})

	tags := []string{}

	for _, session := range sessionsForProject {
		for _, tag := range session.Tags {
			if slices.Contains(tags, tag) {
				continue
			}

			tags = append(tags, tag)
		}
	}

	return tags
}

func (r *InMemorySessionRepository) filterByTimeRange(sessions []session.Session, timeRange timerange.TimeRange) []session.Session {
	filteredSessions := []session.Session{}

	for _, session := range sessions {
		if timeRange.Since.IsZero() && !timeRange.Until.IsZero() {
			if session.StartTime.Before(timeRange.Until) {
				filteredSessions = append(filteredSessions, session)
			}
		} else if !timeRange.Since.IsZero() && timeRange.Until.IsZero() {
			if session.StartTime.After(timeRange.Since) {
				filteredSessions = append(filteredSessions, session)
			}
		} else if !timeRange.Since.IsZero() && !timeRange.Until.IsZero() {
			if session.StartTime.After(timeRange.Since) && session.StartTime.Before(timeRange.Until) {
				filteredSessions = append(filteredSessions, session)
			}
		} else {
			filteredSessions = append(filteredSessions, session)
		}
	}
	return filteredSessions
}

func (r *InMemorySessionRepository) filterByProject(sessions []session.Session, project string) []session.Session {
	filteredSessions := []session.Session{}

	for _, session := range sessions {
		if session.Project == project {
			filteredSessions = append(filteredSessions, session)
		}
	}

	return filteredSessions
}
