GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=snapcontrol-backend
BINARY_PATH=./bin

all: test build

build:
	$(GOBUILD) -ldflags '-w -s' -a -installsuffix cgo -o $(BINARY_PATH)/$(BINARY_NAME) -v

test: 
	$(GOTEST) -coverprofile=coverage/coverage.out -v ./...

clean: 
	$(GOCLEAN)
	rm -rf $(BINARY_PATH)

run:
	$(GOBUILD) -o $(BINARY_PATH)/$(BINARY_NAME) -v
	. $(BINARY_PATH)/$(BINARY_NAME)

install:
	$(GOGET) -u github.com/labstack/echo/...
	$(GOGET) -u github.com/stretchr/testify/assert
	$(GOGET) -u github.com/appleboy/gofight
