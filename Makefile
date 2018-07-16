# Include only if exist
-include .env
export

.PHONY: vendor help


# Docker labels-schema. See http://label-schema.org/rc1/
IMAGE_NAME := alextanhongpin/go-gin-starter
# TAG := $(shell git log -1 --pretty=%h)
# BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
# NAME := $(shell basename `git rev-parse --show-toplevel`)
# VCS_URL := $(shell git config --get remote.origin.url)
# VCS_REF := $(shell git rev-parse HEAD)
# VENDOR := $(shell whoami)
# VERSION := $(shell git rev-parse --short HEAD)

define HELP
usage: make (sub-commands ...)

Gin Starter commands.

commands:
  install  Install vgo and dependencies
  env      Describe the environment variables
  deploy   Deploy to the designated environment (production, staging, .etc)
  start    Start the application
  stop     Stop the application
  version  Current application version
  build    Build the application

endef

help: ## Display help
	@echo "$$HELP"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

vendor:	## Vendor the application
	@vgo mod -init # Creates the go.mod file
	@vgo mod -vendor # Create the vendor folder
	@vgo mod -sync # Syncs the current version

docker: ## Build the docker image
	@docker build -t ${IMAGE_NAME} .
	# --build-arg BUILD_DATE="${BUILD_DATE}" \
	# --build-arg IMAGE_NAME="${IMAGE_NAME}" \
	# --build-arg NAME="${NAME}" \
	# --build-arg VCS_URL="${VCS_URL}" \
	# --build-arg VCS_REF="${VCS_REF}" \
	# --build-arg VENDOR="${VENDOR}" \
	# --build-arg VERSION="${VERSION}" \

start:
	vgo run main.go
