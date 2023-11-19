BINARY_NAME=tokenizer
MAIN_PATH=cmd/tokenizer/main.go

build:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux ${MAIN_PATH}
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows.exe ${MAIN_PATH}
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin ${MAIN_PATH}

clean:
	rm bin/${BINARY_NAME}-linux
	rm bin/${BINARY_NAME}-windows.exe
	rm bin/${BINARY_NAME}-darwin

dep:
	go mod download

lint:
	golangci-lint run
