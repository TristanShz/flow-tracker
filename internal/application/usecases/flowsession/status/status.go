package status

import (
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

	duration := s.dateProvider.GetNow().Sub(lastSession.StartTime).Round(time.Second)

	return fmt.Sprintf("You're in the flow for %v", duration.String()), nil
}

func NewFlowSessionStatusUseCase(sessionRepository application.SessionRepository, dateProvider application.DateProvider) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
		dateProvider:      dateProvider,
	}
}
