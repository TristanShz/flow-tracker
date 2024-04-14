package status

import (
	"errors"
	"fmt"
	"time"

	"github.com/TristanSch1/flow/internal/application"
)

type UseCase struct {
	sessionRepository application.SessionRepository
	dateProvider      application.DateProvider
}

func (s *UseCase) Execute() (string, error) {
	lastSession, err := s.sessionRepository.FindLastSession()
	if err != nil {
		return "", err
	}

	if lastSession == nil {
		return "", ErrNoCurrentSession
	}

	duration := s.dateProvider.GetNow().Sub(lastSession.StartTime).Round(time.Second)

	return fmt.Sprintf("You're in the flow for %v", duration.String()), nil
}

var ErrNoCurrentSession = errors.New("there is no flow session in progress")

func NewFlowSessionStatusUseCase(sessionRepository application.SessionRepository, dateProvider application.DateProvider) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
		dateProvider:      dateProvider,
	}
}
