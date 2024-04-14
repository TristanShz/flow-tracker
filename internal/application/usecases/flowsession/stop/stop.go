package stop

import (
	"errors"

	"github.com/TristanSch1/flow/internal/application"
)

type UseCase struct {
	sessionRepository application.SessionRepository
	dateProvider      application.DateProvider
}

func (s UseCase) Execute() error {
	lastSession, err := s.sessionRepository.FindLastSession()
	if err != nil {
		return err
	}

	if lastSession == nil {
		return ErrNoCurrentSession
	}

	lastSession.EndTime = s.dateProvider.GetNow()

	s.sessionRepository.Save(*lastSession)

	return nil
}

var ErrNoCurrentSession = errors.New("there are no flow sessions in progress")

func NewStopSessionUseCase(sessionRepository application.SessionRepository, dateProvider application.DateProvider) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
		dateProvider:      dateProvider,
	}
}
