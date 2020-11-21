package covid_report

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esvm/if1015covidproject-api/src/api_gateway/context"
	"github.com/esvm/if1015covidproject-api/src/covid_report_service"
	"github.com/labstack/echo"
)

const (
	contentType = "application/json"

	EntryPoint = "/reports"

	GetCovidReportsRoute    = "/"
	InsertCovidReportsRoute = "/"
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
		InsertCovidReportsRoute,
		api.InsertCovidReportsHandler,
		instrumentingMiddleware("InsertCovidReports"),
	)
}

func (api *CovidReportAPI) GetCovidReportsHandler(ctx echo.Context) error {
	c := context.GetContext(ctx)
	reports, err := api.covidService.GetCovidReports(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Covid Report service failed: %s", err.Error()),
		}
	}

	body, err := json.Marshal(reports)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to get covid reports"}
	}

	return ctx.Blob(http.StatusOK, contentType, body)
}

func (api *CovidReportAPI) InsertCovidReportsHandler(ctx echo.Context) error {
	reports, err := UnmarshalCovidReport(ctx)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Failed to parse body"}
	}

	c := context.GetContext(ctx)

	err = api.covidService.InsertCovidReports(c, reports)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Covid Report service failed"}
	}

	return ctx.NoContent(http.StatusCreated)
}
