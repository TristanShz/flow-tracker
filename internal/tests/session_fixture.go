package tests

import (
	"errors"
	"testing"
	"time"

	startflowsession "github.com/TristanSch1/flow/internal/application/usecases/start_flow_session"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra"
)

type SessionFixture struct {
	T                 *testing.T
	SessionRepository *infra.InMemorySessionRepository
	DateProvider      *infra.StubDateProvider
	UseCase           startflowsession.UseCase
	ThrownError       error
}

func (s *SessionFixture) GivenPredefinedStartTime(t time.Time) {
	s.DateProvider.Now = t
}

func (s *SessionFixture) GivenSomeSessions(sessions []session.Session) {
	s.SessionRepository.Sessions = sessions
}

func (s *SessionFixture) WhenStartingFlowSession(command startflowsession.Command) {
	err := s.UseCase.Execute(command)
	if err != nil {
		s.ThrownError = err
	}
}

func (s *SessionFixture) ThenSessionWithGivenStartTimeShouldBeSaved(expectedStartTime time.Time) {
	got := s.SessionRepository.Sessions[0].StartTime

	if got != expectedStartTime {
		s.T.Errorf("Expected '%v', but got '%v'", expectedStartTime, got)
	}
}

func (s *SessionFixture) ThenErrorShouldBe(e error) {
	if !errors.Is(s.ThrownError, e) {
		s.T.Errorf("Expected error '%v', but got '%v'", e, s.ThrownError)
	}
}

func GetSessionFixture(t *testing.T) SessionFixture {
	sessionRepository := &infra.InMemorySessionRepository{}
	dateProvider := &infra.StubDateProvider{}

	useCase := startflowsession.NewStartFlowSessionUseCase(sessionRepository, dateProvider)

	return SessionFixture{
		T:                 t,
		SessionRepository: sessionRepository,
		DateProvider:      dateProvider,
		UseCase:           useCase,
	}
}
