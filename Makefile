.PHONY: help build test test-integration start-dev-jenkins stop-dev-jenkins

help:
	@echo "make help               Show this help"
	@echo "make build              Build the binaries in build directory"
	@echo "make test               Run the unit tests"
	@echo "make test-integration   Run the all the tests, including integration tests"
	@echo "make start-dev-infra    Start the development infrastructure"
	@echo "make stop-dev-infra     Stop the development infrastructure"

build: test
	GOOS=linux GOARCG=amd64 go build -o ./build/linux_amd64/opype .
	GOOS=darwin GOARCG=amd64 go build -o ./build/darwin_amd64/opype . 
	GOOS=windows GOARCG=amd64 go build -o ./build/windows_amd64/opype.exe . 

test:
	go fmt ./... && go vet ./... && golint --set_exit_status ./... && go test ./...

test_integration: test
	go test -tags=integration

start-dev-infra:
	UID=${UID} GID=${GID} cd docker/dev-infra && docker-compose up &

stop-dev-infra:
	cd docker/dev-infra && docker-compose down