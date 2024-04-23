package viewsessionsreport

import (
	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

func (s UseCase) Execute(
	command Command,
	presenter application.SessionsReportPresenter,
) error {
	var sessions []session.Session

	if !command.Since.IsZero() || !command.Until.IsZero() {
		sessionsInTimeRange := s.sessionRepository.FindInTimeRange(application.TimeRange{
			Since: command.Since,
			Until: command.Until,
		})
		sessions = sessionsInTimeRange
	} else if command.Project != "" {
		sessionsForProject := s.sessionRepository.FindAllByProject(command.Project)

		sessions = sessionsForProject
	} else {
		allSessions := s.sessionRepository.FindAllSessions()

		sessions = allSessions
	}

	sessionsReport := sessionsreport.SessionsReport{
		Sessions: sessions,
	}

	if command.Format == sessionsreport.FormatByDay {
		presenter.ShowByDay(sessionsReport)
	} else {
		presenter.ShowByProject(sessionsReport)
	}

	return nil
}

func NewViewSessionsReportUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
