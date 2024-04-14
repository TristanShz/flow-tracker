package application

import "github.com/TristanSch1/flow/internal/domain/session"

type SessionRepository interface {
	Save(session session.Session) error
	FindLastSession() (*session.Session, error)
	FindAllSessions() ([]session.Session, error)
	FindAllByProject(project string) ([]session.Session, error)
	FindAllProjects() ([]string, error)
	FindAllProjectTags(project string) ([]string, error)
}
