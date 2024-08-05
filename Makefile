GO111MODULE = on

DEST_PATH ?= internal/msgs

.ONESHELL:
SHELL = bash

all: configure build #default make target

configure:
	echo "" # make configure unimplemented. Unified deployment API between all software components.

build:
	cd cmd/rclgo-gen && go build

install:
	cd cmd/rclgo-gen && go install

.PHONY: test
test:
	go test -count=1 ./...

.PHONY: test-verbose
test-verbose:
	go test -v -count=1 ./...

generate:
	rm -rf "$(DEST_PATH)/"*
	go run ./cmd/rclgo-gen generate \
	    --root-path /usr \
	    --root-path /opt/ros/${ROS_DISTRO} \
	    --dest-path "$(DEST_PATH)" \
		--message-module-prefix "github.com/tiiuae/rclgo/$(DEST_PATH)" \
		--license-header-path ./license-header.txt \
		--include-go-package-deps ./... \
		--cgo-flags-path "" \
		|| exit 1
	rm "$(DEST_PATH)/msgs.gen.go" || exit 1
	go run ./cmd/rclgo-gen generate-rclgo \
		--root-path /usr \
		--root-path /opt/ros/${ROS_DISTRO} \
		--license-header-path ./license-header.txt

lint:
	golangci-lint run ./...

container-build-image:
	podman build --tag ros-rclgo:jazzy .

container-generate:
	podman run  --rm \
				--interactive \
				--tty \
				--volume $(PWD):/opt/rclgo:ro \
				--volume $(PWD)/internal/msgs:/opt/rclgo-msgs:rw \
				sh -c 'cd /opt/rclgo && make generate -e DEST_PATH=/opt/rclgo-msgs'
