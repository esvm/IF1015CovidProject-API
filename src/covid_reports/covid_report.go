package covid_reports

import "time"

type CovidReportBrazilState struct {
	ID       string `json:"uid" pg:"id" validate:"-"`
	State    string `json:"state" validate:"max=10"`
	UF       string `json:"uf" validate:"max=50"`
	Cases    int    `json:"cases" validate:"-"`
	Deaths   int    `json:"deaths" validate:"-"`
	Suspects int    `json:"suspects" validate:"-"`
	Refuses  int    `json:"refuses" validate:"-"`

	UpdatedAt time.Time `json:"updated_at" validate:"-"`
}

type CovidReportCountry struct {
	Country  string `json:"country" validate:"max=100"`
	Cases    int    `json:"cases" validate:"-"`
	Deaths   int    `json:"deaths" validate:"-"`
	Suspects int    `json:"suspects" validate:"-"`
	Refuses  int    `json:"refuses" validate:"-"`

	UpdatedAt time.Time `json:"updated_at" validate:"-"`
}
