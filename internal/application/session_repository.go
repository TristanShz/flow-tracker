package application

import (
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
)

type TimeRange struct {
	Since time.Time
	Until time.Time
}

type SessionRepository interface {
	Save(session session.Session) error
	FindLastSession() (*session.Session, error)
	FindAllSessions() []session.Session
	FindAllByProject(project string) ([]session.Session, error)
	FindAllProjects() ([]string, error)
	FindAllProjectTags(project string) ([]string, error)
	FindInTimeRange(timeRange TimeRange) ([]session.Session, error)
}
