package application

import "github.com/TristanSch1/flow/internal/domain/session"

type SessionRepository interface {
	Save(session session.Session) error
	FindAllByProject(project string) []session.Session
	FindLastSession() *session.Session
}
