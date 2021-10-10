# Add the following 'help' target to your Makefile
# And add help text after each target name starting with '\#\#'
.PHONY: help

help:           	## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

vet:       		## Runs "go vet" against project
	@go vet ./...

test:      		## Runs unit tests
	@go test ./...

format:			## Runs go fmt against code
	@go fmt ./...

build:			## Builds the binary
	@go build -o bin/
