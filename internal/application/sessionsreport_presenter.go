package application

import (
	"github.com/TristanShz/flow/internal/domain/sessionsreport"
)

type SessionsReportPresenter interface {
	ShowByProject(sessionsReport sessionsreport.SessionsReport)
	ShowByDay(sessionsReport sessionsreport.SessionsReport)
}
