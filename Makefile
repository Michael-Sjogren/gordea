APP_NAME=gordea
build:
	CGO_ENABLED=0 go build -o bin/$(APP_NAME)

run: build
	./bin/$(APP_NAME)

test:
	go test -v ./...