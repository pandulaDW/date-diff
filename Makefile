BINARY_NAME=date-diff

build:
	 GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
	 GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux main.go
	 GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows.exe main.go

test:
	go test -v ./...

run:
	go run main.go

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows.exe