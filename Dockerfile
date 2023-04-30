FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download
 
RUN	CGO_ENABLED=0 GOOS=linux go build -o ./bin/crnt-data-manager ./cmd/crnt-data-manager/main.go

CMD ["./bin/crnt-data-manager"]
