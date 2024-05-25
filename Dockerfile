FROM golang:1.22.2-alpine3.19 as stage_build

WORKDIR /app

ADD . /app
RUN go mod download

RUN go build

FROM alpine:3.20.0

WORKDIR /app

COPY --from=stage_build /app/api-jwt-go /app/
COPY --from=stage_build /app/.env /app/

ENTRYPOINT [ "./api-jwt-go" ]

# [error] failed to initialize database, got error failed to connect to 
# `host=localhost user=postgres database=postgres`: dial error (dial tcp [::1]:5432: connect: connection refused)
