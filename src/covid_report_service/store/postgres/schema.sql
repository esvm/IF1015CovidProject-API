CREATE TABLE covid_report_brazil_states(
    id int,
    state text,
    uf text,
    cases bigint,
    deaths bigint,
    suspects bigint,
    refuses bigint,
    updated_at timestamp with time zone
);

CREATE UNIQUE INDEX idx_unique_covid_report_brazil_states
ON covid_report_brazil_states(state, updated_at);

CREATE INDEX idx_uf_covid_report_brazil_states
ON covid_report_brazil_states(uf);

CREATE INDEX idx_id_covid_report_brazil_states
ON covid_report_brazil_states(id);

CREATE TABLE covid_report_countries(
    country text,
    cases bigint,
    deaths bigint,
    suspects bigint,
    refuses bigint,
    updated_at timestamp with time zone
);

CREATE UNIQUE INDEX idx_unique_covid_report_countries 
ON covid_report_countries(country, updated_at);