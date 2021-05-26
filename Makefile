.PHONY: all
all: mod build

mod:
	go mod download

build: clean
	go build -o setup

test:
	go test ./...

test-ci: mod
	$(eval localPath=$(shell pwd))
	curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | bash
	PATH=$(PATH):${localPath}/pact/bin go test ./...

clean:
	rm -f setup

docker: clean mod
	env GOOS=linux GOARCH=amd64 go build -o setup
	docker build -t eu.gcr.io/sbat-gcr-develop/securebanking/secureopenbanking-uk-fidc-initializer:latest .
