CREATE TABLE covid_reports(
    id text PRIMARY KEY,
    state text,
    uf text,
    country text,
    cases bigint,
    deaths bigint,
    suspects bigint,
    refuses bigint,
    reported_at timestamp with time zone
);