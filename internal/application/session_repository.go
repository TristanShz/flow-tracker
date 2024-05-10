package application

import (
	"github.com/TristanShz/flow/internal/domain/session"
	"github.com/TristanShz/flow/pkg/timerange"
)

type SessionRepository interface {
	Save(session session.Session) error
	Delete(id string) error
	FindById(id string) *session.Session
	FindLastSession() *session.Session
	FindAllSessions() []session.Session
	FindAllByProject(project string) []session.Session
	FindAllProjects() []string
	FindAllProjectTags(project string) []string
	FindInTimeRange(timeRange timerange.TimeRange) []session.Session
}
