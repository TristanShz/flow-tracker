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
		text += fmt.Sprintf("%v :\n", utils.GreenText(dayReport.Day.Format("2006-01-02")))
		for _, session := range dayReport.Sessions {
			text += fmt.Sprintf(
				"    From %v to %v %v %v [%v]\n",
				utils.GreenText(session.StartTime.Format("15:04:05")),
				utils.GreenText(session.EndTime.Format("15:04:05")),
				session.Duration().String(),
				utils.PurpleText(session.Project),
				utils.YellowText(strings.Join(session.Tags, ", ")),
			)
		}

	}

	fmt.Println(text)
}

func (s SessionsReportCLIPresenter) ShowByProject(sessionsReport sessionsreport.SessionsReport) {
	byProjectReport := sessionsReport.GetByProjectReport()
	text := "Sessions Report\n\n"

	for _, report := range byProjectReport {
		text += fmt.Sprintf("%v - %v\n", utils.YellowText(report.Project), utils.GreenText(report.TotalDuration.String()))
		for tag, duration := range report.DurationByTag {
			text += fmt.Sprintf("    %v -> %v\n", utils.YellowText(tag), utils.GreenText(duration.String()))
		}

	}

	fmt.Println(text)
}
