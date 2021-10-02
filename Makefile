all:
	gofmt -s -w .
	go build -o codegen
	
build:
	go build -o codegen
	
run:
	go run main.go
