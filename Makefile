BINARY_NAME=minder

build:
	go build -o ${BINARY_NAME} main.go

run: build 
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test ${BINARY_NAME}/src/test -v

dep:
	go mod download