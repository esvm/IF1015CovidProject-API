package covid_report

import (
	"encoding/json"
	"net/http"

	"github.com/esvm/if1015covidproject-api/src/api_gateway/context"
	"github.com/esvm/if1015covidproject-api/src/covid_report_service"
	"github.com/labstack/echo"
)

const (
	contentType = "application/json"

	EntryPoint = "/reports"

	GetCovidReportsRoute   = "/"
	InsertCovidReportRoute = "/"
)

type CovidReportAPI struct {
	covidService covid_report_service.CovidReportService
}

func MakeCovidReportRoutes(
	g *echo.Group,
	covidService covid_report_service.CovidReportService,
) {
	api := &CovidReportAPI{covidService}

	g.GET(
		GetCovidReportsRoute,
		api.GetCovidReportsHandler,
		instrumentingMiddleware("GetCovidReports"),
	)
	g.POST(
		InsertCovidReportRoute,
		api.InsertCovidReportHandler,
		instrumentingMiddleware("InsertCovidReport"),
	)
}

func (api *CovidReportAPI) GetCovidReportsHandler(ctx echo.Context) error {
	c := context.GetContext(ctx)
	reports, err := api.covidService.GetCovidReports(c)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Covid Report service failed"}
	}

	body, err := json.Marshal(reports)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to get covid reports"}
	}

	return ctx.Blob(http.StatusOK, contentType, body)
}

func (api *CovidReportAPI) InsertCovidReportHandler(ctx echo.Context) error {
	report, err := UnmarshalCovidReport(ctx)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Failed to parse body"}
	}

	c := context.GetContext(ctx)

	created, err := api.covidService.InsertCovidReport(c, report)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Covid Report service failed"}
	}

	body, _ := json.Marshal(created)
	return ctx.Blob(http.StatusCreated, contentType, body)
}
