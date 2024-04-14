package status

import (
	"errors"
	"fmt"
	"time"

	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/session"
)

type SessionStatus struct {
	StatusText string
	Session    session.Session
}

type UseCase struct {
	sessionRepository application.SessionRepository
	dateProvider      application.DateProvider
}

func (s *UseCase) Execute() (SessionStatus, error) {
	lastSession, err := s.sessionRepository.FindLastSession()
	if err != nil {
		return SessionStatus{}, err
	}

	if lastSession == nil || lastSession.Status() != session.FlowingStatus {
		return SessionStatus{}, ErrNoCurrentSession
	}

	duration := s.dateProvider.GetNow().Sub(lastSession.StartTime).Round(time.Second)

	return SessionStatus{
		Session:    *lastSession,
		StatusText: fmt.Sprintf("You're in the flow for %v", duration.String()),
	}, nil
}

var ErrNoCurrentSession = errors.New("there is no flow session in progress")

func NewFlowSessionStatusUseCase(sessionRepository application.SessionRepository, dateProvider application.DateProvider) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
		dateProvider:      dateProvider,
	}
}
