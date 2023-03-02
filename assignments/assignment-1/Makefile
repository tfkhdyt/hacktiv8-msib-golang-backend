BINARY_NAME=biodata
BINARY_DIR=bin

build:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_DIR}/${BINARY_NAME}-linux-amd64 ${BINARY_NAME}.go
	GOOS=darwin GOARCH=amd64 go build -o ${BINARY_DIR}/${BINARY_NAME}-darwin-amd64 ${BINARY_NAME}.go
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_DIR}/${BINARY_NAME}-windows-amd64.exe ${BINARY_NAME}.go

clean:
	go clean
	rm -f ${BINARY_DIR}/${BINARY_NAME}*

.PHONY: build clean
