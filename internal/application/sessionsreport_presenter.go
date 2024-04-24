package application

import (
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
)

type SessionsReportPresenter interface {
	ShowByProject(sessionsReport sessionsreport.SessionsReport)
	ShowByDay(sessionsReport sessionsreport.SessionsReport)
}
