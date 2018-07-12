include .env
export

# Docker labels-schema. See http://label-schema.org/rc1/
IMAGE_NAME := alextanhongpin/gin-starter
# TAG := $(shell git log -1 --pretty=%h)
# BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
# NAME := $(shell basename `git rev-parse --show-toplevel`)
# VCS_URL := $(shell git config --get remote.origin.url)
# VCS_REF := $(shell git rev-parse HEAD)
# VENDOR := $(shell whoami)
# VERSION := $(shell git rev-parse --short HEAD)

define HELP
usage: make (sub-commands ...)

Get started with gin.

commands:
  install Install vgo and dependencies

endef

default:
	echo "$$HELP"

vendor:
	@vgo mod -vendor
	@vgo mod -sync

docker:
	@docker build -t ${IMAGE_NAME} .
	# --build-arg BUILD_DATE="${BUILD_DATE}" \
	# --build-arg IMAGE_NAME="${IMAGE_NAME}" \
	# --build-arg NAME="${NAME}" \
	# --build-arg VCS_URL="${VCS_URL}" \
	# --build-arg VCS_REF="${VCS_REF}" \
	# --build-arg VENDOR="${VENDOR}" \
	# --build-arg VERSION="${VERSION}" \

build:
	@echo Build verified!