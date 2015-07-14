# all will get required go packages
all: dependencies

# dependencies installs all of the dependencies that are required for building
# Sunfish.
dependencies:
	go install -race std
	go get -u gopkg.in/mgo.v2
	go get github.com/gorilla/mux

# fmt calls go fmt on all packages.
fmt:
	go fmt ./...

.PHONY: dependencies fmt
