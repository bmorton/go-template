GO_VERSION = 1.6.0
WORKING_DIR = `pwd`


.PHONY: all clean build

all: clean build run

build:
	construct --template=`pwd` --source-path="tmp" new testapp
	go get github.com/tools/godep
	cd tmp/src/github.com/bmorton/testapp && GOPATH=`grealpath ./../../../..` godep restore
	cd tmp/src/github.com/bmorton/testapp && GOPATH=`grealpath ./../../../..` godep save
	cd tmp/src/github.com/bmorton/testapp && GOPATH=`grealpath ./../../../..` go build

clean:
	rm -rf tmp/src

run: 
	./tmp/src/github.com/bmorton/testapp/testapp --development server
