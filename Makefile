SRC_YAML?="swagger.yml"
CGO?=cgo

CWD?=$(shell pwd)

SHELL:=/bin/bash
GO_PIPELINE_LABEL?=BUILD_ID
ENVIRONMENT?=DEVELOPMENT

BUILD_NUMBER?=$(GO_PIPELINE_LABEL)
BUILD_ID?=$(ENVIRONMENT)

GO_TARGETS= ./server ./backend
GOR_TARGETS= ./server/... ./backend/...

SHA?=$(shell git rev-parse HEAD)
BUILT_BY?=$(shell whoami)
HOSTNAME?=$(shell hostname)
NOW?=$(shell date +%s)

VERSION?=$(shell git rev-parse --short HEAD)

run-clean: clean build run

run: terraform-server
	$(CWD)/terraform-server --port=8080 --scheme=http

update-vendor:
	glide update

clean: clean-client
	rm -rf $(CWD)/server/cmd/ \
		$(CWD)/server/models/ \
		$(CWD)/server/restapi/operations \
		$(CWD)/server/restapi/doc.go \
		$(CWD)/server/restapi/embedded_spec.go \
		$(CWD)/server/restapi/server.go \
		$(CWD)/terraform-server \
		$(CWD)/profile.txt \


clean-client:
	rm -rf $(CWD)/tfs-client

validate-swagger:
	swagger validate $(SRC_YAML)

build: generate-swagger terraform-server

test:
	go tool vet $(GO_TARGETS)
	go test $(GOR_TARGETS)

test-coverage:
	@sh $(CWD)/scripts/test-coverage.sh $(CWD) "$(GO_TARGETS)"
	go tool cover -html=$(CWD)/profile.out -o $(CWD)/coverage.html

format:
	go fmt $(shell go list ./...)

lint:
	diff -u <(echo -n) <(gofmt -d -s $(shell find server -type d))
	diff -u <(echo -n) <(gofmt -d -s $(shell find backend -type d))
	golint -set_exit_status . $(GOR_TARGETS)

generate-swagger: validate-swagger
	swagger generate server \
		--target=server \
		--principal=models.Principal \
		--name=TerraformServer \
		--spec=$(SRC_YAML)

generate:generate-swagger

terraform-server:
	go build \
		-a -installsuffix $(CGO) \
		-o ./terraform-server \
		github.com/drewsonne/terraform-server/server/cmd/terraform-server-server

build-docker:
	GOOS=linux $(MAKE) terraform-server
	docker build -t terraform-server .

run-docker: build-docker
	docker run -p 8080:8080 terraform-server

build-client:
	swagger generate client \
		-f swagger.yml \
		-A terraform-server-client \
		--existing-models github.com/drewsonne/terraform-server/server/models \
		--skip-models \
		--target client && \
	go build -ldflags "-X github.com/drewsonne/terraform-server/cli/cmd.version=$(VERSION)" -o tfs-client github.com/drewsonne/terraform-server/cli

client-clean: clean-client build-client