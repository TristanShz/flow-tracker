package infra

import (
	"slices"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
)

type InMemorySessionRepository struct {
	Sessions []session.Session
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

func (r *InMemorySessionRepository) FindByStartTime(startTime time.Time) *session.Session {
	index := slices.IndexFunc(r.Sessions, func(s session.Session) bool {
		return s.StartTime.Equal(startTime)
	})

	return &r.Sessions[index]
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
