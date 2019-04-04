.PHONY: build fmt godep clean

APPNAME := pgb-reload

export GOPATH = $(PWD)/vendor:$(PWD)
export GOBIN = $(PWD)/vendor/bin:$(PWD)/bin

default: build

build: fmt clean
	go build -v -o ./${APPNAME}

godep:
	go get

fmt:
	go fmt

clean:
	rm -rf `find ./vendor/src -type d -name .git` && \
	rm -rf `find ./vendor/src -type d -name .hg` && \
	rm -rf `find ./vendor/src -type d -name .bzr` && \
	rm -rf `find ./vendor/src -type d -name .svn`
	rm -rf ./vendor/bin/*
	rm -rf ./bin/*
	rm -rf ./vendor/pkg/*
	rm -rf ./pkg/*
	rm -rf ${APPNAME}
