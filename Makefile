SRC_YAML?="swagger.yml"
COMPILER?="cgo"

SHELL:=/bin/bash
GO_PIPELINE_LABEL?=BUILD_ID
ENVIRONMENT?=DEVELOPMENT

BUILD_NUMBER?=$(GO_PIPELINE_LABEL)
BUILD_ID?=$(ENVIRONMENT)

GO_TARGETS= ./server ./backend

doc:
	@sh scripts/generate-doc.sh

update-vendor:
	glide update

run: terraform-server
	./terraform-server --scheme=http

validate-swagger:
	swagger validate $(SRC_YAML)

build: generate-swagger terraform-server

test:
	go tool vet $(GO_TARGETS)
	go test $(GO_TARGETS)

test-coverage:
	goverage -v -race -coverprofile=profile.txt -covermode=atomic $(GO_TARGETS)

format:
	go fmt $(shell go list ./...)

lint:
	diff -u <(echo -n) <(gofmt -d -s main.go $(GO_TARGETS))
	golint -set_exit_status . $(GO_TARGETS)

generate-swagger: validate-swagger
	swagger generate server \
		--target=server \
		--principal=models.Principal \
		--name=TerraformServer \
		--spec=$(SRC_YAML)

terraform-server:
	go build \
		-a -installsuffix $(COMPILER) \
		-ldflags "-X github.com/zeebox/terraform-server/server/restapi.builtWhen=$(shell date +%s) \
				-X github.com/zeebox/terraform-server/server/restapi.buildMachine=$(shell hostname) \
				-X github.com/zeebox/terraform-server/server/restapi.buildNumber=$(BUILD_NUMBER) \
				-X github.com/zeebox/terraform-server/server/restapi.builtBy=$(shell whoami) \
				-X github.com/zeebox/terraform-server/server/restapi.buildId=$(BUILD_ID)\
				-X github.com/zeebox/terraform-server/server/restapi.compiler=$(COMPILER) \
				-X github.com/zeebox/terraform-server/server/restapi.sha=$(shell git rev-parse HEAD)" \
		-o ./terraform-server \
		github.com/zeebox/terraform-server/server/cmd/terraform-server-server

build-docker:
	GOOS=linux $(MAKE) terraform-server
	docker build -t terraform-server .

run-docker: build-docker
	docker run -p 8080:8080 terraform-server