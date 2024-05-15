package viewsessionsreport

import (
	"github.com/TristanShz/flow/internal/application"
	"github.com/TristanShz/flow/internal/domain/sessionsreport"
	"github.com/TristanShz/flow/pkg/timerange"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

func (s UseCase) Execute(
	command Command,
	presenter application.SessionsReportPresenter,
) error {
	filters := &application.SessionsFilters{}

	if command.Project != "" {
		filters.Project = command.Project
	}

	if !command.Since.IsZero() || !command.Until.IsZero() {
		filters.Timerange = timerange.TimeRange{
			Since: command.Since,
			Until: command.Until,
		}
	}

	sessions := s.sessionRepository.FindAllSessions(filters)

	sessionsReport := sessionsreport.SessionsReport{
		Sessions: sessions,
	}

	if command.Format == sessionsreport.FormatByProject {
		presenter.ShowByProject(sessionsReport)
	} else {
		presenter.ShowByDay(sessionsReport)
	}

	return nil
}

func NewViewSessionsReportUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
