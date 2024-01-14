.PHONY: run
run:
	go run ./cmd/xyzmas --renderer console --fps 10 --mode test

.PHONY: benchmark
benchmark:
	go run ./cmd/xyzmas --renderer null --fps 99999999999 --benchmark --verbose

# Where to deploy to, using SSH
PI_HOST=xyzmas@berry.local
# Target platform configuration
PI_ARCH=arm
PI_ARM=6
PI_LDSO=/lib/ld-linux-armhf.so.3 # /lib/ld-linux.so.3 when not overridden
PI_PLATFORM=linux/arm/v6
# Name of docker build image
PI_BUILDER=xyzmas-builder-$(PI_ARCH)-$(PI_ARM)

# Local Go path.  The build will use the same package cache
GOPATH=$(shell go env GOPATH)

# Build outputs and intermediate files
PI_BUILD_OUTPUT=build/$(PI_ARCH)/$(PI_ARM)
PI_BUILDER_TARGET=$(PI_BUILD_OUTPUT)/builder.target

$(PI_BUILDER_TARGET): Dockerfile
	docker buildx build . \
		--platform $(PI_PLATFORM) \
		--tag $(PI_BUILDER) \
		--file Dockerfile
	mkdir -p $(PI_BUILD_OUTPUT)
	touch $(PI_BUILDER_TARGET)

.PHONY: build
build: $(PI_BUILDER_TARGET)
	docker run --rm \
		--platform $(PI_PLATFORM) \
		-v "$$PWD:/usr/src/app" \
		-v "$(GOPATH)/pkg/mod:/go/pkg/mod" \
		--env GOCACHE=/usr/src/app/$(PI_BUILD_OUTPUT)/cache \
		--env GOOS=linux \
		--env GOARCH=$(PI_ARCH) \
		--env GOARM=$(PI_ARM) \
		-w /usr/src/app $(PI_BUILDER):latest \
		go build -ldflags='-I $(PI_LDSO)' -tags ws281x -o $(PI_BUILD_OUTPUT) -v ./cmd/xyzmas

.PHONY: deploy
deploy: build
	rsync -ar $(PI_BUILD_OUTPUT)/xyzmas config.json $(PI_HOST):server

.PHONY: start
start:
	ssh $(PI_HOST) 'sudo systemctl start xyzmas.service'

.PHONY: stop
stop:
	ssh $(PI_HOST) 'sudo systemctl stop xyzmas.service'

.PHONY: restart
restart:
	ssh $(PI_HOST) 'sudo systemctl restart xyzmas.service'

.PHONY: shutdown
shutdown:
	ssh $(PI_HOST) 'sudo shutdown now'
