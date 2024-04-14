package tests

import (
	"errors"
	"slices"
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stop"
	"github.com/TristanSch1/flow/internal/application/usecases/project/list"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra"
)

type SessionFixture struct {
	T                        *testing.T
	SessionRepository        *infra.InMemorySessionRepository
	DateProvider             *infra.StubDateProvider
	StartFlowSessionUseCase  start.UseCase
	StopFlowSessionUseCase   stop.UseCase
	FlowSessionStatusUseCase status.UseCase
	ListProjectsUseCase      list.UseCase
	ThrownError              error
	FlowSessionStatus        status.SessionStatus
	Projects                 []string
}

func (s *SessionFixture) GivenNowIs(t time.Time) {
	s.DateProvider.Now = t
}

func (s *SessionFixture) GivenSomeSessions(sessions []session.Session) {
	s.SessionRepository.Sessions = sessions
}

func (s *SessionFixture) WhenStartingFlowSession(command start.Command) {
	err := s.StartFlowSessionUseCase.Execute(command)
	if err != nil {
		s.ThrownError = err
	}
}

func (s *SessionFixture) WhenStoppingFlowSession() {
	err := s.StopFlowSessionUseCase.Execute()
	if err != nil {
		s.ThrownError = err
	}
}

func (s *SessionFixture) WhenUserSeesTheCurrentSessionStatus() {
	status, err := s.FlowSessionStatusUseCase.Execute()
	if err != nil {
		s.ThrownError = err
	}

	s.FlowSessionStatus = status
}

func (s *SessionFixture) WhenGettingListOfProjects() {
	projects, err := s.ListProjectsUseCase.Execute()
	if err != nil {
		s.ThrownError = err
	}

	s.Projects = projects
}

func (s *SessionFixture) ThenProjectsShouldBe(projects []string) {
	got := s.Projects

	if !slices.Equal(got, projects) {
		s.T.Errorf("Expected projects '%v', but go '%v'", projects, got)
	}
}

func (s *SessionFixture) ThenUserShouldSee(session session.Session, statusText string) {
	got := s.FlowSessionStatus

	if got.StatusText != statusText {
		s.T.Errorf("Expected SessionStatus.StatusText '%v', but got '%v'", statusText, got.StatusText)
	}

	if !got.Session.Equals(session) {
		s.T.Errorf("Expected SessionStatus.Session '%v', but got '%v'", session.PrettyString(), got.Session.PrettyString())
	}
}

func (s *SessionFixture) ThenSessionShouldBeSaved(session session.Session) {
	got := s.SessionRepository.Sessions[0]

	if !got.Equals(session) {
		s.T.Errorf("Expected '%v', but got '%v'", session, got)
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
	dateProvider := infra.NewStubDateProvider()

	startFlowSession := start.NewStartFlowSessionUseCase(sessionRepository, dateProvider)
	stopFlowSession := stop.NewStopSessionUseCase(sessionRepository, dateProvider)
	flowSessionStatus := status.NewFlowSessionStatusUseCase(sessionRepository, dateProvider)

	listProjects := list.NewListProjectsUseCase(sessionRepository)

	return SessionFixture{
		T:                        t,
		SessionRepository:        sessionRepository,
		DateProvider:             dateProvider,
		StartFlowSessionUseCase:  startFlowSession,
		StopFlowSessionUseCase:   stopFlowSession,
		FlowSessionStatusUseCase: flowSessionStatus,
		ListProjectsUseCase:      listProjects,
	}
}
