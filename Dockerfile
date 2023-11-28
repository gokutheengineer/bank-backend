FROM golang:1.20.11-alpine3.18 AS build
WORKDIR /app
COPY . .
#RUN go mod tidy
#RUN go mod download
RUN go build -o main main.go

FROM alpine:latest AS run
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/configs/  ./configs/.
#COPY db/migration ./db/migration
EXPOSE 8080
CMD ["/app/main"]
