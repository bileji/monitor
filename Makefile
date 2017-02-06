MONITOR = monitor
MONITOR_SERVER = monitord
build:
	go build -o ../bin/${MONITOR} -ldflags '-s -w' ./cmd/monitor.go
	go build -o ../bin/${MONITOR_SERVER} -ldflags '-s -w' ./cmd/monitord.go
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/${MONITOR} -ldflags '-s -w' ./cmd/monitor.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/${MONITOR_SERVER} -ldflags '-s -w' ./cmd/monitord.go
run:
	go run *.go
clean:
	@rm -rf ../bin/${MONITOR} ../bin/${MONITOR_SERVER}
