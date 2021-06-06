# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

all: build

build:
	$(GOBUILD) -ldflags "-X main.version=${DRONE_TAG} -X main.build=${DRONE_BUILD_NUMBER}" -o gce -v .
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "-X main.version=${DRONE_TAG} -X main.build=${DRONE_BUILD_NUMBER}" -o gce -v .

clean:
	$(GOCLEAN)
	rm -f gce