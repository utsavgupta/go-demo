FROM golang:1.21 as build
WORKDIR /src
COPY go.mod go.sum ./
COPY go.sum ./
COPY calc/calc.go ./calc/
COPY calc/cmd/main.go ./calc/cmd/
COPY agg/agg.go ./agg/
COPY agg/cmd/main.go ./agg/cmd/
RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -o calc_app ./calc/cmd
FROM scratch
WORKDIR /app
COPY --from=build /src/calc_app ./
ENV nr_license=eu01xxeefc5eae67612e3dea126d9c71FFFFINAL
ENTRYPOINT [ "/app/calc_app" ]