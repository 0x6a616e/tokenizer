BINARY_NAME=tokenizer
MAIN_PATH=cmd/tokenizer/main.go

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ${MAIN_PATH}
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows ${MAIN_PATH}
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin ${MAIN_PATH}

clean:
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows
	rm ${BINARY_NAME}-darwin

dep:
	go mod download

lint:
	golangci-lint run
