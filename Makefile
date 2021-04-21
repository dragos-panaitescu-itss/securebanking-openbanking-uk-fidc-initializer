.PHONY: all
all: mod build

mod:
	go mod download

build: clean
	go build -o setup

clean:
	rm -f setup

docker: clean mod
	env GOOS=linux GOARCH=amd64 go build -o setup
	docker build -t eu.gcr.io/sbat-gcr-develop/securebanking/secureopenbanking-uk-fidc-initializer:latest .
