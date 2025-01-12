build:
	go build -o bin/inkube cmd/inkube/main.go

build-arm:
	GOOS=linux GOARCH=arm go build -o bin/inkube cmd/inkube/main.go

run:
	go run cmd/inkube/main.go
