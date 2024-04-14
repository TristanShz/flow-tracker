package allsessionsreport

import (
	"fmt"
	"time"

	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
	"github.com/TristanSch1/flow/utils"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

type AllSessionsReport struct {
	Projects         map[string]time.Duration
	Total            time.Duration
	NumberOfSessions int
}

func (r AllSessionsReport) PrettyPrint() string {
	result := utils.HeaderStyle.Render("All Flow Sessions Report:\n\n")
	result += "Projects:\n"
	for project, duration := range r.Projects {
		result += fmt.Sprintf("~ %s: %s\n", project, duration.String())
	}
	result += fmt.Sprintf("Total flow time: %s\n", r.Total.String())
	result += fmt.Sprintf("Number of sessions: %d\n", r.NumberOfSessions)
	return result
}

func (s UseCase) Execute() (AllSessionsReport, error) {
	sessions, err := s.sessionRepository.FindAllSessions()
	if err != nil {
		return AllSessionsReport{}, err
	}

	sessionsReport := sessionsreport.SessionsReport{
		Sessions: sessions,
	}

	return AllSessionsReport{
		Projects:         sessionsReport.ProjectsReport(),
		Total:            sessionsReport.TotalDuration(),
		NumberOfSessions: len(sessions),
	}, nil
}

func NewFlowSessionsReportUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
