package tests

import (
	"errors"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/sessionstatus"
	startsession "github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stopsession"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/application/usecases/project/list"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
	"github.com/TristanSch1/flow/internal/infra"
)

type TestPresenter struct {
	SessionsReportByDay     sessionsreport.SessionsReport
	SessionsReportByProject sessionsreport.SessionsReport
}

func (tp *TestPresenter) ShowByDay(sessionReport sessionsreport.SessionsReport) {
	tp.SessionsReportByDay = sessionReport
}

func (tp *TestPresenter) ShowByProject(sessionReport sessionsreport.SessionsReport) {
	tp.SessionsReportByProject = sessionReport
}

type SessionFixture struct {
	T                         *testing.T
	SessionRepository         *infra.InMemorySessionRepository
	DateProvider              *infra.StubDateProvider
	IdProvider                *infra.StubIDProvider
	StartFlowSessionUseCase   startsession.UseCase
	StopFlowSessionUseCase    stopsession.UseCase
	FlowSessionStatusUseCase  sessionstatus.UseCase
	ListProjectsUseCase       list.UseCase
	ViewSessionsReportUseCase viewsessionsreport.UseCase
	ThrownError               error
	FlowSessionStatus         sessionstatus.SessionStatus
	Projects                  []string
	SessionsReport            sessionsreport.SessionsReport
	SessionsReportPresenter   TestPresenter
}

func (s *SessionFixture) GivenNowIs(t time.Time) {
	s.DateProvider.Now = t
}

func (s *SessionFixture) GivenPredefinedIdentifier(id string) {
	s.IdProvider.Id = id
}

func (s *SessionFixture) GivenSomeSessions(sessions []session.Session) {
	s.SessionRepository.Sessions = sessions
}

func (s *SessionFixture) WhenStartingFlowSession(command startsession.Command) {
	err := s.StartFlowSessionUseCase.Execute(command)
	if err != nil {
		s.ThrownError = err
	}
}

func (s *SessionFixture) WhenStoppingFlowSession() {
	_, err := s.StopFlowSessionUseCase.Execute()
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

func (s *SessionFixture) WhenUserSeesSessionsReport(
	command viewsessionsreport.Command,
) {
	err := s.ViewSessionsReportUseCase.Execute(command, &s.SessionsReportPresenter)
	if err != nil {
		s.ThrownError = err
	}
}

func (s SessionFixture) ThenUserShouldSeeSessionsReport(expectedReport sessionsreport.SessionsReport, expectedFormat string) {
	var got sessionsreport.SessionsReport
	if expectedFormat == sessionsreport.FormatByDay {
		got = s.SessionsReportPresenter.SessionsReportByDay
	}
	if expectedFormat == sessionsreport.FormatByProject {
		got = s.SessionsReportPresenter.SessionsReportByProject
	}

	if !reflect.DeepEqual(got, expectedReport) {
		s.T.Errorf("Expected report with session ids '%v', but got '%v'", s.formatReportForError(expectedReport), s.formatReportForError(got))
	}
}

func (s SessionFixture) formatReportForError(expectedReport sessionsreport.SessionsReport) string {
	ids := make([]string, len(expectedReport.Sessions))
	for i, session := range expectedReport.Sessions {
		ids[i] = session.Id
	}

	return strings.Join(ids, ", ")
}

func (s *SessionFixture) ThenProjectsShouldBe(projects []string) {
	got := s.Projects

	if !slices.Equal(got, projects) {
		s.T.Errorf("Expected projects '%v', but go '%v'", projects, got)
	}
}

func (s *SessionFixture) ThenUserShouldSee(session session.Session, duration time.Duration) {
	got := s.FlowSessionStatus

	if got.Duration != duration {
		s.T.Errorf("Expected SessionStatus.StatusText '%v', but got '%v'", duration, got.Duration)
	}

	if !got.Session.Equals(session) {
		s.T.Errorf("Expected SessionStatus.Session '%v', but got '%v'", session, got.Session)
	}
}

func (s *SessionFixture) ThenSessionShouldBeSaved(session session.Session) {
	got := s.SessionRepository.Sessions[0]

	if !got.Equals(session) {
		s.T.Errorf("Expected '%v', but got '%v'", session, got)
	}
}

func (s *SessionFixture) ThenSessionShouldBeStopped() {
	got := s.SessionRepository.FindLastSession()

	if got.EndTime.IsZero() {
		s.T.Errorf("Found not stopped session '%v'", got)
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
	idProvider := &infra.StubIDProvider{}

	startFlowSession := startsession.NewStartFlowSessionUseCase(sessionRepository, dateProvider, idProvider)
	stopFlowSession := stopsession.NewStopSessionUseCase(sessionRepository, dateProvider)
	flowSessionStatus := sessionstatus.NewFlowSessionStatusUseCase(sessionRepository, dateProvider)

	viewSessionsReport := viewsessionsreport.NewViewSessionsReportUseCase(sessionRepository)
	sessionsReportPresenter := TestPresenter{}

	listProjects := list.NewListProjectsUseCase(sessionRepository)

	return SessionFixture{
		T:                         t,
		SessionRepository:         sessionRepository,
		IdProvider:                idProvider,
		DateProvider:              dateProvider,
		StartFlowSessionUseCase:   startFlowSession,
		StopFlowSessionUseCase:    stopFlowSession,
		FlowSessionStatusUseCase:  flowSessionStatus,
		ListProjectsUseCase:       listProjects,
		ViewSessionsReportUseCase: viewSessionsReport,
		SessionsReportPresenter:   sessionsReportPresenter,
	}
}
