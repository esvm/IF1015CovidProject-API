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

	InsertCovidReport(*covid_reports.CovidReport) (*covid_reports.CovidReport, error)
	GetCovidReports() ([]*covid_reports.CovidReport, error)
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

func (d covidReportDatabase) InsertCovidReport(covidReport *covid_reports.CovidReport) (*covid_reports.CovidReport, error) {
	db := d.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "Begin transaction failed")
	}

	err = tx.Insert(covidReport)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "Insert Covid Report query failed")
	}

	tx.Commit()

	return covidReport, nil
}

func (d covidReportDatabase) GetCovidReports() ([]*covid_reports.CovidReport, error) {
	db := d.GetConnection()

	covidReports := []*covid_reports.CovidReport{}
	if err := db.Model(&covidReports).Select(); err != nil {
		return nil, errors.Wrap(err, "Failed to select Covid Reports")
	}

	return covidReports, nil
}