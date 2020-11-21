package postgres

import (
	"os"

	"github.com/esvm/if1015covidproject-api/src/covid_reports"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
)

var connection *pg.DB

type BasicDatabase interface {
	GetConnection() *pg.DB
	CloseConnection()
}

type CovidReportDatabase interface {
	BasicDatabase

	InsertCovidReportsBrazil([]*covid_reports.CovidReportBrazilState) error
	InsertCovidReportsCountries([]*covid_reports.CovidReportCountry) error
	GetCovidReportsBrazil() ([]*covid_reports.CovidReportBrazilState, error)
}

type covidReportDatabase struct {
	BasicDatabase

	logger log.Logger
}

func NewDatabase(logger log.Logger) CovidReportDatabase {
	var database CovidReportDatabase = covidReportDatabase{logger: logger}

	return database
}

func (d covidReportDatabase) CloseConnection() {
	if connection != nil {
		level.Info(d.logger).Log("message", "Close Postgres session")
		connection.Close()
	}
}

func (d covidReportDatabase) GetConnection() *pg.DB {
	if connection == nil {
		addr := os.Getenv("COVID_REPORT_DATABASE_ADDRESS")
		port := os.Getenv("COVID_REPORT_DATABASE_PORT")
		user := os.Getenv("COVID_REPORT_DATABASE_USER")
		pass := os.Getenv("COVID_REPORT_DATABASE_PASS")
		name := os.Getenv("COVID_REPORT_DATABASE_NAME")

		connection = pg.Connect(&pg.Options{
			User:     user,
			Password: pass,
			Database: name,
			Addr:     addr + ":" + port,
			PoolSize: 30,
		})
	}

	return connection
}

func (d covidReportDatabase) InsertCovidReportsBrazil(covidReports []*covid_reports.CovidReportBrazilState) error {
	db := d.GetConnection()

	_, err := db.Model(&covidReports).Insert()
	if err != nil {
		pgErr, _ := err.(pg.Error)
		if pgErr.IntegrityViolation() {
			return nil
		}
	}

	return err
}

func (d covidReportDatabase) InsertCovidReportsCountries(covidReports []*covid_reports.CovidReportCountry) error {
	db := d.GetConnection()

	_, err := db.Model(&covidReports).Insert()
	if err != nil {
		pgErr, _ := err.(pg.Error)
		if pgErr.IntegrityViolation() {
			return nil
		}
	}

	return err
}

func (d covidReportDatabase) GetCovidReportsBrazil() ([]*covid_reports.CovidReportBrazilState, error) {
	db := d.GetConnection()

	covidReports := []*covid_reports.CovidReportBrazilState{}
	if err := db.Model(&covidReports).Select(); err != nil {
		return nil, errors.Wrap(err, "Failed to select Covid Reports")
	}

	return covidReports, nil
}
