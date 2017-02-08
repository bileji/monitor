MONITOR = monitor
build:
	go build -o ../bin/${MONITOR} -ldflags '-s -w' ./main.go
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/${MONITOR} -ldflags '-s -w' ./main.go
run:
	go run *.go
clean:
	@rm -rf ../bin/${MONITOR}
