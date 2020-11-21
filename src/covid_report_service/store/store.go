package store

import (
	"github.com/esvm/if1015covidproject-api/src/covid_report_service/store/postgres"
	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/go-kit/kit/log"
)

type Store interface {
	InsertCovidReportsBrazil([]*covid_reports.CovidReportBrazilState) error
	InsertCovidReportsCountries([]*covid_reports.CovidReportCountry) error
	GetCovidReportsBrazil() ([]*covid_reports.CovidReportBrazilState, error)
}

type basicStore struct {
	logger log.Logger
}

func New(logger log.Logger) Store {
	return basicStore{logger}
}

func (s basicStore) InsertCovidReportsBrazil(covidReports []*covid_reports.CovidReportBrazilState) error {
	database := postgres.NewDatabase(s.logger)
	return database.InsertCovidReportsBrazil(covidReports)
}

func (s basicStore) InsertCovidReportsCountries(covidReports []*covid_reports.CovidReportCountry) error {
	database := postgres.NewDatabase(s.logger)
	return database.InsertCovidReportsCountries(covidReports)
}

func (s basicStore) GetCovidReportsBrazil() ([]*covid_reports.CovidReportBrazilState, error) {
	database := postgres.NewDatabase(s.logger)
	return database.GetCovidReportsBrazil()
}
