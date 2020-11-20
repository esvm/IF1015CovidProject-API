FROM alpine:latest
RUN apk add --update ca-certificates
COPY bin/covid_report /
RUN chmod +x covid_report
ENTRYPOINT ["/covid_report"]