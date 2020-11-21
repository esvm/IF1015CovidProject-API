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

	GetCovidReportsBrazilRoute          = "/brazil"
	InsertCovidReportsBrazilStatesRoute = "/brazil"
	GetCovidReportsBrazilPerDayRoute    = "/brazil"

	GetCovidReportsCountriesRoute    = "/countries"
	InsertCovidReportsCountriesRoute = "/countries"
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
		GetCovidReportsBrazilRoute,
		api.GetCovidReportsBrazilHandler,
		instrumentingMiddleware("GetCovidReportsBrazil"),
	)
	g.POST(
		InsertCovidReportsBrazilStatesRoute,
		api.InsertCovidReportsBrazilHandler,
		instrumentingMiddleware("InsertCovidReportsBrazil"),
	)

	g.GET(
		GetCovidReportsCountriesRoute,
		api.GetCovidReportsCountriesHandler,
		instrumentingMiddleware("GetCovidReportsCountries"),
	)
	g.POST(
		InsertCovidReportsCountriesRoute,
		api.InsertCovidReportsCountriesHandler,
		instrumentingMiddleware("InsertCovidReportsCountries"),
	)
}

func (api *CovidReportAPI) GetCovidReportsBrazilHandler(ctx echo.Context) error {
	c := context.GetContext(ctx)
	reports, err := api.covidService.GetCovidReportsBrazil(c)
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

func (api *CovidReportAPI) GetCovidReportsCountriesHandler(ctx echo.Context) error {
	c := context.GetContext(ctx)
	reports, err := api.covidService.GetCovidReportsCountries(c)
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

func (api *CovidReportAPI) InsertCovidReportsBrazilHandler(ctx echo.Context) error {
	reports, err := UnmarshalCovidReportBrazil(ctx)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Failed to parse body"}
	}

	c := context.GetContext(ctx)

	err = api.covidService.InsertCovidReportsBrazil(c, reports)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Covid Report service failed %s", err.Error()),
		}
	}

	return ctx.NoContent(http.StatusCreated)
}

func (api *CovidReportAPI) InsertCovidReportsCountriesHandler(ctx echo.Context) error {
	reports, err := UnmarshalCovidReportCountries(ctx)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Failed to parse body"}
	}

	c := context.GetContext(ctx)

	err = api.covidService.InsertCovidReportsCountries(c, reports)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Covid Report service failed"}
	}

	return ctx.NoContent(http.StatusCreated)
}
