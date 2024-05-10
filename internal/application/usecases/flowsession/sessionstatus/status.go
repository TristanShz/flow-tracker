package sessionstatus

import (
	"errors"
	"time"

	"github.com/TristanShz/flow/internal/application"
	"github.com/TristanShz/flow/internal/domain/session"
)

type SessionStatus struct {
	Session  session.Session
	Duration time.Duration
}

type UseCase struct {
	sessionRepository application.SessionRepository
	dateProvider      application.DateProvider
}

func (s *UseCase) Execute() (SessionStatus, error) {
	lastSession := s.sessionRepository.FindLastSession()

	if lastSession == nil || lastSession.Status() != session.FlowingStatus {
		return SessionStatus{}, ErrNoCurrentSession
	}

	duration := s.dateProvider.GetNow().Sub(lastSession.StartTime).Round(time.Second)

	return SessionStatus{
		Session:  *lastSession,
		Duration: duration,
	}, nil
}

var ErrNoCurrentSession = errors.New("there is no flow session in progress")

func NewFlowSessionStatusUseCase(sessionRepository application.SessionRepository, dateProvider application.DateProvider) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
		dateProvider:      dateProvider,
	}
}
