package projectsessionsreport

import (
	"fmt"
	"time"

	"github.com/TristanSch1/flow/internal/application"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
)

type UseCase struct {
	sessionRepository application.SessionRepository
}

type ProjectSessionReport struct {
	Project          string
	Total            time.Duration
	NumberOfSessions int
}

func (r ProjectSessionReport) PrettyPrint() string {
	result := fmt.Sprintf("%v project sessions report : \n\n", r.Project)
	result += fmt.Sprintf("Total flow time: %s\n", r.Total.String())
	result += fmt.Sprintf("Number of sessions: %d\n", r.NumberOfSessions)
	return result
}

func (s UseCase) Execute(command Command) (ProjectSessionReport, error) {
	sessions, err := s.sessionRepository.FindAllByProject(command.Project)
	if err != nil {
		return ProjectSessionReport{}, err
	}

	sessionsReport := sessionsreport.SessionsReport{
		Sessions: sessions,
	}

	return ProjectSessionReport{
		Project:          command.Project,
		Total:            sessionsReport.TotalDuration(),
		NumberOfSessions: len(sessions),
	}, nil
}

func NewProjectSessionsReportUseCase(sessionRepository application.SessionRepository) UseCase {
	return UseCase{
		sessionRepository: sessionRepository,
	}
}
