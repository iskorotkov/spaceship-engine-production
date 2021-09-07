FROM golang:1.17 AS build
WORKDIR /go/src

COPY ["go.mod", "go.sum", "./"]
RUN go get -d -v ./...

COPY . .
ENV CGO_ENABLED=0
RUN go build -o app cmd/report-printer/main.go

FROM alpine AS runtime
WORKDIR /

COPY --from=build /go/src/app ./
EXPOSE 8080/tcp
ENTRYPOINT ["./app"]
