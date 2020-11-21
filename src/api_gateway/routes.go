package api_gateway

import (
	"github.com/esvm/if1015covidproject-api/src/api_gateway/covid_report"
	"github.com/esvm/if1015covidproject-api/src/covid_report_service"
	"github.com/labstack/echo"
)

type Clients struct {
	CovidReportService covid_report_service.CovidReportService
}

func MakeRoutes(app *echo.Echo, clients Clients) {
	covidReportRoutes := app.Group(covid_report.EntryPoint)
	covid_report.MakeCovidReportRoutes(covidReportRoutes, clients.CovidReportService)
}
