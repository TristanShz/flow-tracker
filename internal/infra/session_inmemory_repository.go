package infra

import (
	"github.com/TristanSch1/flow/internal/domain/session"
)

type InMemorySessionRepository struct {
	Sessions []session.Session
}

func (r *InMemorySessionRepository) Save(session session.Session) error {
	r.Sessions = append(r.Sessions, session)
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

func (r *InMemorySessionRepository) FindLastProjectSession(project string) *session.Session {
	allProjectSessions := r.FindAllByProject(project)

	if len(allProjectSessions) == 0 {
		return nil
	}

	return &allProjectSessions[len(allProjectSessions)-1]
}
