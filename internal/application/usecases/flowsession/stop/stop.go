package stop

import (
	"errors"
	"time"

	"github.com/TristanSch1/flow/internal/application"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

func (s UseCase) Execute() error {
	lastSession, err := s.sessionRepository.FindLastSession()
	if err != nil {
		return err
	}

	if lastSession == nil {
		return ErrNoCurrentSession
	}

	lastSession.EndTime = time.Now()

	s.sessionRepository.Save(*lastSession)

	return nil
}

var ErrNoCurrentSession = errors.New("there are no flow sessions in progress")

func NewStopSessionUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
