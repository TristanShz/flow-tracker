package stop

import (
	"errors"
	"time"

	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/session"
)

type UseCase struct {
	sessionRepository application.SessionRepository
	dateProvider      application.DateProvider
}

func (s UseCase) Execute() (time.Duration, error) {
	lastSession, err := s.sessionRepository.FindLastSession()
	if err != nil {
		return 0, err
	}

	if lastSession == nil || lastSession.Status() != session.FlowingStatus {
		return 0, ErrNoCurrentSession
	}

	lastSession.EndTime = s.dateProvider.GetNow()

	s.sessionRepository.Save(*lastSession)

	return lastSession.Duration(), nil
}

var ErrNoCurrentSession = errors.New("there is no flow session in progress")

func NewStopSessionUseCase(sessionRepository application.SessionRepository, dateProvider application.DateProvider) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
		dateProvider:      dateProvider,
	}
}
