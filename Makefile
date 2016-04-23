#
# Simple Makefile
#

build: range.go
	go build -o bin/range range.go 

clean: range
	rm bin/range

install: range.go
	go install range.go

