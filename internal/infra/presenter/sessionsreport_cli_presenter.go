package presenter

import (
	"fmt"
	"log"
	"strings"

	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
	"github.com/TristanSch1/flow/utils"
)

type SessionsReportCLIPresenter struct {
	Logger *log.Logger
}

func (s SessionsReportCLIPresenter) ShowByDay(sessionsReport sessionsreport.SessionsReport) {
	byDayReport := sessionsReport.GetByDayReport()
	text := "Sessions Report\n\n"

	for _, dayReport := range byDayReport {
		text += fmt.Sprintf("%v :\n", utils.HeaderStyle.Render(dayReport.Day.Format("Mon, 02 Jan 2006")))
		for _, session := range dayReport.Sessions {
			if session.EndTime.IsZero() {
				text += fmt.Sprintf(
					"    %v %v %v [%v]\n",
					session.Id,
					utils.TimeColor(session.StartTime.Format("15:04:05")),
					utils.ProjectColor(session.Project),
					utils.TagColor(strings.Join(session.Tags, ", ")),
				)
			} else {
				text += fmt.Sprintf(
					"    %v %v to %v %v %v [%v]\n",
					utils.Faint(session.Id),
					utils.TimeColor(session.StartTime.Format("15:04:05")),
					utils.TimeColor(session.EndTime.Format("15:04:05")),
					session.Duration().String(),
					utils.ProjectColor(session.Project),
					utils.TagColor(strings.Join(session.Tags, ", ")),
				)
			}
		}

		text += "\n"
	}

	s.Logger.Println(text)
}

func (s SessionsReportCLIPresenter) ShowByProject(sessionsReport sessionsreport.SessionsReport) {
	byProjectReport := sessionsReport.GetByProjectReport()
	text := "Sessions Report\n\n"

	for _, report := range byProjectReport {
		text += fmt.Sprintf("%v - %v\n", utils.ProjectColor(report.Project), utils.TimeColor(report.TotalDuration.String()))
		for tag, duration := range report.DurationByTag {
			text += fmt.Sprintf("    [%v] -> %v\n", utils.TagColor(tag), utils.TimeColor(duration.String()))
		}

		text += "\n"
	}

	s.Logger.Println(text)
}
