package application

import (
	"github.com/TristanShz/flow/internal/domain/session"
	"github.com/TristanShz/flow/pkg/timerange"
)

type SessionsFilters struct {
	Timerange timerange.TimeRange
	Project   string
}

type SessionRepository interface {
	Save(session session.Session) error
	Delete(id string) error
	FindById(id string) *session.Session
	FindLastSession() *session.Session
	FindAllSessions(filters *SessionsFilters) []session.Session
	FindAllProjects() []string
	FindAllProjectTags(project string) []string
}
