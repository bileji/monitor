APP = monitor
build:
	go build -o ../bin/${APP} -ldflags '-s -w' ./cmd/monitord.go
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/${APP} -ldflags '-s -w' ./cmd/monitord.go
run:
	go run *.go
clean:
	@rm ../bin/{$APP}