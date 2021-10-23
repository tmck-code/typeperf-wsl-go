build:
	go build -o typeperf-wsl-go main.go

poc: build
	./typeperf-wsl-go

.PHONY: build poc
