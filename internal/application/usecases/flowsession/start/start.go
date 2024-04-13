package start

import (
	"errors"

	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/session"
)

type UseCase struct {
	sessionRepository application.SessionRepository
	dateProvider      application.DateProvider
}

func (s UseCase) Execute(command Command) error {
	lastSession, err := s.sessionRepository.FindLastSession()
	if err != nil {
		return err
	}

	if lastSession != nil && lastSession.EndTime.IsZero() {
		return ErrSessionAlreadyStarted
	}

	session := session.Session{
		StartTime: s.dateProvider.GetNow(),
		Project:   command.Project,
		Tags:      command.Tags,
	}

	s.sessionRepository.Save(session)

	return nil
}

var ErrSessionAlreadyStarted = errors.New("there is already a session in progress")

func NewStartFlowSessionUseCase(sessionRepository application.SessionRepository, dateProvider application.DateProvider) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
		dateProvider:      dateProvider,
	}
}
