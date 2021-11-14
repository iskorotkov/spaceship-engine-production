FROM golang:1.17 AS build
WORKDIR /go/src

COPY ["go.mod", "go.sum", "./"]
RUN go get -d -v ./...

COPY . .
ENV CGO_ENABLED=0
RUN go build -o app ./cmd/data-loader-server/...

FROM alpine AS runtime
WORKDIR /

COPY --from=build /go/src/app ./
COPY ./config/data-loader-server/config.yml ./config.yaml
EXPOSE 8080/tcp
ENTRYPOINT ["./app"]
