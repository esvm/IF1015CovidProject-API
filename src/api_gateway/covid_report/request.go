package covid_report

import (
	"encoding/json"

	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/labstack/echo"
)

func UnmarshalCovidReport(ctx echo.Context) (*covid_reports.CovidReport, error) {
	req := ctx.Request()

	covidReport := covid_reports.CovidReport{}
	err := json.NewDecoder(req.Body).Decode(&covidReport)

	return &covidReport, err
}