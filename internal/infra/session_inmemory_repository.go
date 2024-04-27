package infra

import (
	"slices"

	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/pkg/timerange"
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

func (r *InMemorySessionRepository) FindAllByProject(project string) []session.Session {
	sessions := []session.Session{}

	for _, session := range r.Sessions {
		if session.Project == project {
			sessions = append(sessions, session)
		}
	}

	return sessions
}

func (r *InMemorySessionRepository) FindLastSession() *session.Session {
	if len(r.Sessions) == 0 {
		return nil
	}

	return &r.Sessions[len(r.Sessions)-1]
}

func (r *InMemorySessionRepository) FindAllSessions() []session.Session {
	return r.Sessions
}

func (r *InMemorySessionRepository) FindAllProjects() []string {
	sessions := r.FindAllSessions()

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
	sessionsForProject := r.FindAllByProject(project)

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

func (r *InMemorySessionRepository) FindInTimeRange(timeRange timerange.TimeRange) []session.Session {
	sessions := []session.Session{}

	for _, session := range r.Sessions {
		if timeRange.Since.IsZero() && !timeRange.Until.IsZero() {
			if session.StartTime.Before(timeRange.Until) {
				sessions = append(sessions, session)
			}
		} else if !timeRange.Since.IsZero() && timeRange.Until.IsZero() {
			if session.StartTime.After(timeRange.Since) {
				sessions = append(sessions, session)
			}
		} else if !timeRange.Since.IsZero() && !timeRange.Until.IsZero() {
			if session.StartTime.After(timeRange.Since) && session.StartTime.Before(timeRange.Until) {
				sessions = append(sessions, session)
			}
		} else {
			sessions = append(sessions, session)
		}
	}
	return sessions
}
