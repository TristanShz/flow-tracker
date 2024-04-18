package application

import "github.com/TristanSch1/flow/internal/domain/sessionsreport"

type SessionsReportPresenter interface {
	Show(sessionsReport sessionsreport.SessionsReport)
}
