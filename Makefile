VERSION = 1.0.0
APP= alien-invasion
PACKAGE_DIR = ./cmd/${APP}
TEST_DIR=./tests

GO = /usr/local/go/bin/go
NOW := $(shell date -u +'%Y-%m-%d_%T')
LDFLAGS = "-X main.VersionString=$(VERSION) -X main.BuildTime=$(NOW)"
STATICCHECK_VER = 2022.1

all: module-verify build test
build:
	echo ${GO} build -v -o ./bin/${APP} -ldflags $(LDFLAGS) $(PACKAGE_DIR)
	${GO} build -v -o ./bin/${APP} -ldflags $(LDFLAGS) $(PACKAGE_DIR)
	cp ${TEST_DIR}/* ./bin/

clean:
	echo ${GO} clean $(PACKAGE_DIR)
	${GO} clean $(PACKAGE_DIR)
	rm -rf bin/*

module-verify:
	${GO} mod tidy

test:	build
	${GO} test -v ./...
