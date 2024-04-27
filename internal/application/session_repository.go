package application

import (
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/pkg/timerange"
)

type SessionRepository interface {
	Save(session session.Session) error
	FindById(id string) *session.Session
	FindLastSession() *session.Session
	FindAllSessions() []session.Session
	FindAllByProject(project string) []session.Session
	FindAllProjects() []string
	FindAllProjectTags(project string) []string
	FindInTimeRange(timeRange timerange.TimeRange) []session.Session
}
