#FROM golang:1.23 AS build
#WORKDIR /app
#COPY . .
#
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd
#
#FROM debian:bookworm-slim
#WORKDIR /app
#COPY --from=build /app/cloudrun .
#EXPOSE 8080
#ENTRYPOINT [ "./cloudrun" ]

FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd

FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080
ENTRYPOINT [ "./cloudrun" ]
