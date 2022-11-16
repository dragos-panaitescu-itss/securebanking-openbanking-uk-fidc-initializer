service := secureopenbanking-uk-iam-initializer
gcr-repo := sbat-gcr-develop
binary-name := initialize

.PHONY: all
all: mod build

mod:
	go mod download

build: clean mod
	go build -o ${binary-name}

test:
	go test ./...

test-ci: mod
	$(eval localPath=$(shell pwd))
	curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | bash
	PATH=$(PATH):${localPath}/pact/bin go test ./...

clean:
	go clean
	rm -f ${binary-name}

docker: clean mod
ifndef tag
	$(warning No tag supplied, latest assumed. supply tag with make docker tag=x.x.x service=...)
	$(eval tag=latest)
endif
	env GOOS=linux GOARCH=amd64 go build -o initialize
	docker buildx build --platform linux/amd64  -t eu.gcr.io/${gcr-repo}/securebanking/${service}:${tag} .
	docker push eu.gcr.io/${gcr-repo}/securebanking/${service}:${tag}
ifdef release-repo
	docker tag eu.gcr.io/${gcr-repo}/securebanking/${service}:${tag} eu.gcr.io/${release-repo}/securebanking/${service}:${tag}
	docker push eu.gcr.io/${release-repo}/securebanking/${service}:${tag}
endif
