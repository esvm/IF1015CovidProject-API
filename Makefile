BASE_IMAGE=api-starter-golang

pipeline/docker/base:
	docker build -t $(BASE_IMAGE):latest .
	docker save --output $(BASE_IMAGE).docker $(BASE_IMAGE):latest

pipeline/lint:
	docker load --input ./$(BASE_IMAGE).docker
	docker-compose run $(BASE_IMAGE) golangci-lint run -v

pipeline/test:
	make infrastructure/raise
	make db/bootstrap
	docker load --input ./$(BASE_IMAGE).docker
	docker-compose run $(BASE_IMAGE)

infrastructure/raise:
	docker-compose up -d db

db/bootstrap:
	sleep 10
	docker-compose exec -T -e PGPASSWORD=${COVID_REPORT_DATABASE_PASS} db psql -h localhost -U ${COVID_REPORT_DATABASE_USER} -p ${COVID_REPORT_DATABASE_PORT} -c "CREATE DATABASE ${COVID_REPORT_DATABASE_NAME}"
	docker-compose exec -T -e PGPASSWORD=${COVID_REPORT_DATABASE_PASS} db psql -h localhost -d ${COVID_REPORT_DATABASE_NAME} -U ${COVID_REPORT_DATABASE_USER} -p ${COVID_REPORT_DATABASE_PORT} \
		-f /src/covid_report_service/store/postgres/schema.sql

start:
	docker load --input ./$(BASE_IMAGE).docker
	docker-compose run --service-ports $(BASE_IMAGE) goreman start

build_and_push:
	env GOOS=linux GOARCH=386 go build -a --ldflags="-s" -o cmd/api_gateway/bin/api_gateway cmd/api_gateway/api_gateway.go
	docker build -t covid-reports cmd/api_gateway/
	rm -rf cmd/api_gateway/bin/
	docker tag covid-reports:latest 349254485044.dkr.ecr.sa-east-1.amazonaws.com/covid-reports:latest
	docker push 349254485044.dkr.ecr.sa-east-1.amazonaws.com/covid-reports:latest
