package covid_report

import (
	"encoding/json"

	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/labstack/echo"
)

func UnmarshalCovidReportBrazil(ctx echo.Context) ([]*covid_reports.CovidReportBrazilState, error) {
	req := ctx.Request()

	covidReports := []*covid_reports.CovidReportBrazilState{}
	err := json.NewDecoder(req.Body).Decode(&covidReports)

	return covidReports, err
}

func UnmarshalCovidReportCountries(ctx echo.Context) ([]*covid_reports.CovidReportCountry, error) {
	req := ctx.Request()

	covidReports := []*covid_reports.CovidReportCountry{}
	err := json.NewDecoder(req.Body).Decode(&covidReports)

	return covidReports, err
}
