#
# Simple Makefile
#

build:
	go build -o bin/range cmds/range/range.go 

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

install:
	env GOBIN=$HOME/bin go install cmds/range/range.go

release:
	./mk-release.sh
