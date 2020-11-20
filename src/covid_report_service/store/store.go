package store

import (
	"github.com/esvm/if1015covidproject-api/src/covid_report_service/store/postgres"
	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/go-kit/kit/log"
)

type Store interface {
	InsertCovidReport(*covid_reports.CovidReport) (*covid_reports.CovidReport, error)
	GetCovidReports() ([]*covid_reports.CovidReport, error)
}

type basicStore struct {
	logger log.Logger
}

func New(logger log.Logger) Store {
	return basicStore{logger}
}

func (s basicStore) InsertCovidReport(covidReport *covid_reports.CovidReport) (*covid_reports.CovidReport, error) {
	database := postgres.NewDatabase(s.logger)
	return database.InsertCovidReport(covidReport)
}

func (s basicStore) GetCovidReports() ([]*covid_reports.CovidReport, error) {
	database := postgres.NewDatabase(s.logger)
	return database.GetCovidReports()
}
