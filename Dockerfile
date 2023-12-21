FROM golang:1.20.11-alpine3.18 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:latest AS run
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/configs/  ./configs/.
COPY --from=build /app/migrate.linux-amd64 ./migrate
COPY start.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]