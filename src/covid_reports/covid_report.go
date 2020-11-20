package covid_reports

import "time"

type CovidReport struct {
	ID       string `json:"id" sql:",pk" validate:"required"`
	State    string `json:"state" validate:"max=10"`
	UF       string `json:"uf" validate:"max=50"`
	Country  string `json:"country" validate:"max=100"`
	Cases    int    `json:"cases" validate:"-"`
	Deaths   int    `json:"deaths" validate:"-"`
	Suspects int    `json:"suspects" validate:"-"`
	Refuses  int    `json:"refuses" validate:"-"`

	ReportedAt time.Time `json:"reported_at" validate:"-"`
}
