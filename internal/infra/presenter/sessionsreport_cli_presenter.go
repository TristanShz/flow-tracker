package presenter

import (
	"fmt"
	"strings"

	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
	"github.com/TristanSch1/flow/utils"
)

type SessionsReportCLIPresenter struct{}

func (s SessionsReportCLIPresenter) ShowByDay(sessionsReport sessionsreport.SessionsReport) {
	byDayReport := sessionsReport.GetByDayReport()
	text := "Sessions Report\n\n"

	for _, dayReport := range byDayReport {
		text += fmt.Sprintf("%v :\n", utils.TimeColor(dayReport.Day.Format("2006-01-02")))
		for _, session := range dayReport.Sessions {
			text += fmt.Sprintf(
				"    From %v to %v %v %v [%v]\n",
				utils.TimeColor(session.StartTime.Format("15:04:05")),
				utils.TimeColor(session.EndTime.Format("15:04:05")),
				session.Duration().String(),
				utils.ProjectColor(session.Project),
				utils.TagColor(strings.Join(session.Tags, ", ")),
			)
		}

	}

	fmt.Println(text)
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

	fmt.Println(text)
}
