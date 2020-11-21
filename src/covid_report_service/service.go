package covid_report_service

import (
	"context"

	"github.com/esvm/if1015covidproject-api/src/covid_report_service/store"
	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/go-kit/kit/log"
)

type CovidReportService interface {
	InsertCovidReportsBrazil(context.Context, []*covid_reports.CovidReportBrazilState) error
	InsertCovidReportsCountries(context.Context, []*covid_reports.CovidReportCountry) error
	GetCovidReportsBrazil(context.Context) ([]*covid_reports.CovidReportBrazilState, error)
}

type basicService struct {
	store store.Store
}

func NewCovidReportService(logger log.Logger) CovidReportService {
	var service CovidReportService
	service = basicService{store.New(logger)}
	service = loggingMiddleware{logger, service}
	return service
}

func (s basicService) InsertCovidReportsBrazil(ctx context.Context, reports []*covid_reports.CovidReportBrazilState) error {
	if err := ValidateCovidReportsBrazil(reports); err != nil {
		return err
	}

	return s.store.InsertCovidReportsBrazil(reports)
}

func (s basicService) InsertCovidReportsCountries(ctx context.Context, reports []*covid_reports.CovidReportCountry) error {
	if err := ValidateCovidReportsCountries(reports); err != nil {
		return err
	}

	return s.store.InsertCovidReportsCountries(reports)
}

func (s basicService) GetCovidReportsBrazil(ctx context.Context) ([]*covid_reports.CovidReportBrazilState, error) {
	return s.store.GetCovidReportsBrazil()
}
