SHELL := /bin/bash
BASEDIR = $(shell pwd)

export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOPRIVATE=*.gitlab.com
export GOSUMDB=off

# params pass from cmd
APP_BRANCH = "master"
APP_ENV = "dev"
APP_KUBE_CONFIG = "mb-sz-test"

# params stable
APP_NAME=`cat package.json | grep name | head -1 | awk -F: '{ print $$2 }' | sed 's/[\",]//g' | tr -d '[[:space:]]'`
APP_VERSION=`cat package.json | grep version | head -1 | awk -F: '{ print $$2 }' | sed 's/[\",]//g' | tr -d '[[:space:]]'`
COMMIT_ID=`git rev-parse HEAD`
IMAGE_PREFIX="registry.cn-shenzhen.aliyuncs.com/jamin-cui/${APP_NAME}:v${APP_VERSION}"

all: fmt imports mod lint test
fmt:
	gofmt -w .
imports:
	goimports -w .
mod:
	go mod tidy
lint:
	golangci-lint run -c .golangci.yml
.PHONY: test
test: mod
	go test -gcflags=-l -coverpkg=./... -coverprofile=coverage.data ./...
.PHONY: build
build:
	IMAGE_NAME="${IMAGE_PREFIX}-${APP_BRANCH}"; \
	sh build/package/build.sh ${COMMIT_ID} $$IMAGE_NAME
build-master:
	make build APP_BRANCH=master
build-release:
	make build APP_BRANCH=release
cleanup:
	sh scripts/cleanup.sh
deploy-preview:
	NEW_IMAGE="${APP_NAME}=${IMAGE_PREFIX}-${APP_BRANCH}"; \
	sh deploy/kubernetes/deploy-preview.sh $$NEW_IMAGE ${COMMIT_ID} ${APP_ENV}
.PHONY: deploy
deploy:
	sh deploy/kubernetes/deploy.sh ${APP_KUBE_CONFIG} ${APP_ENV}
deploy-direct:
	make deploy-preview
	make deploy
deploy-dev:
	make deploy-direct APP_BRANCH=master APP_ENV=dev APP_KUBE_CONFIG=mb-sz-test
deploy-test:
	make deploy-direct APP_BRANCH=release APP_ENV=test APP_KUBE_CONFIG=mb-sz-test
deploy-we-test:
	make deploy-direct APP_BRANCH=release APP_ENV=we-test APP_KUBE_CONFIG=mb-we-test
deploy-pre:
	make deploy-direct APP_BRANCH=release APP_ENV=pre APP_KUBE_CONFIG=mb-sz-prod
deploy-prod-preview:
	make deploy-preview APP_BRANCH=release APP_ENV=prod
deploy-prod:
	make deploy APP_ENV=prod APP_KUBE_CONFIG=mb-sz-prod
deploy-we-prod-preview:
	make deploy-preview APP_BRANCH=release APP_ENV=we-prod
deploy-we-prod:
	make deploy APP_ENV=we-prod APP_KUBE_CONFIG=mb-we-prod
help:
	@echo "fmt - format the source code"
	@echo "imports - goimports"
	@echo "mod - go mod tidy"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"
	@echo "build - build docker image"
	@echo "build-master - build docker image for master branch"
	@echo "build-release - build docker image for release branch"
	@echo "cleanup - clean up the build binary"
	@echo "deploy-preview - preview deployment yaml"
	@echo "deploy - deploy to kubernetes"
	@echo "deploy-direct - deploy-preview & deploy"
	@echo "deploy-dev - deploy to shenzhen develop environment"
	@echo "deploy-test - deploy to shenzhen test environment"
	@echo "deploy-we-test - deploy to west europe test environment"
	@echo "deploy-pre - deploy to shenzhen preview environment"
	@echo "deploy-prod - deploy to shenzhen production environment"
	@echo "deploy-prod-preview - preview the yaml of shenzhen production environment"
	@echo "deploy-we-prod - deploy to west europe production environment"
	@echo "deploy-we-prod-preview - preview the yaml of west europe production environment"