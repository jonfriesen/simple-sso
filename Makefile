


# Tidy dependencies are all up to date
tidy:
	@go mod tidy
.PHONY: ensure

setup: generate_certs
	@go get -u github.com/jteeuwen/go-bindata/...
.PHONY: setup

test:
	@go test -failfast -race -timeout 30s -covermode=atomic -coverprofile=coverage.txt -cover -json ./... | tparse
.PHONY: test

fmt:
	@find . -name '*.go' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
.PHONY: fmt

cover: test
	@go tool cover -html=coverage.txt
.PHONY: cover

run:
	@go run cmd/simple-sso-oldap/*

run-example:
	@go run example_app/main.go

generate_certs:
	@mkdir -p ssl_certs
	@openssl req \
       -newkey rsa:2048 -nodes -keyout ssl_certs/ssl_private.key \
       -x509 -days 365 -out ssl_certs/ssl_public.crt \
       -subj "/C=CA/ST=British Columbia/L=Vancouver/O=Sample SSL Certificate/CN=localhost"
	@openssl req \
       -newkey rsa:2048 -nodes -keyout key_pair/sso_private.key \
       -x509 -days 365 -out key_pair/sso_public.crt \
       -subj "/C=CA/ST=British Columbia/L=Vancouver/O=Sample SSL Certificate/CN=localhost"
