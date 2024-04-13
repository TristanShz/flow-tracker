package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stop"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra"
)

type SessionFixture struct {
	T                 *testing.T
	SessionRepository *infra.InMemorySessionRepository
	DateProvider      *infra.StubDateProvider
	StartFlowSession  start.UseCase
	StopFlowSession   stop.UseCase
	ThrownError       error
}

func (s *SessionFixture) GivenPredefinedStartTime(t time.Time) {
	s.DateProvider.Now = t
}

func (s *SessionFixture) GivenSomeSessions(sessions []session.Session) {
	s.SessionRepository.Sessions = sessions
}

func (s *SessionFixture) WhenStartingFlowSession(command start.Command) {
	err := s.StartFlowSession.Execute(command)
	if err != nil {
		s.ThrownError = err
	}
}

func (s *SessionFixture) WhenStoppingFlowSession() {
	err := s.StopFlowSession.Execute()
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

func (s *SessionFixture) ThenSessionShouldBeStopped() {
	got, _ := s.SessionRepository.FindLastSession()

	if got.EndTime.IsZero() {
		s.T.Errorf("Found not stopped session '%v'", got.PrettyString())
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

	startFlowSession := start.NewStartFlowSessionUseCase(sessionRepository, dateProvider)
	stopFlowSession := stop.NewStopSessionUseCase(sessionRepository)

	return SessionFixture{
		T:                 t,
		SessionRepository: sessionRepository,
		DateProvider:      dateProvider,
		StartFlowSession:  startFlowSession,
		StopFlowSession:   stopFlowSession,
	}
}
