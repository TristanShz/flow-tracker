package presenter

import (
	"fmt"
	"strings"

	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
	"github.com/TristanSch1/flow/utils"
)

type SessionsReportCLIPresenter struct{}

func (s SessionsReportCLIPresenter) Show(sessionsReport sessionsreport.SessionsReport) {
	sessionsByDay := sessionsReport.SplitSessionsByDay()
	text := "Sessions Report\n\n"

	for day, sessions := range sessionsByDay {
		text += fmt.Sprintf("%v :\n", utils.GreenText(day.Format("2006-01-02")))
		for _, session := range sessions {
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
