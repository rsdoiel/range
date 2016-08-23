#
# Simple Makefile
#

build:
	go build -o bin/range cmds/range/range.go 

clean:
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f range-binary-release.zip ]; then /bin/rm range-binary-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/range/range.go

release:
	./mk-release.bash

save:
	./mk-website.bash
	git commit -am "quick save"
	git push origin master

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash
