# all will get required go packages
all: dependencies

# dependencies installs all of the dependencies that are required for building
# Sunfish.
dependencies:
	go get	./... 

# fmt calls go fmt on all packages.
fmt:
	go fmt ./...

.PHONY: dependencies fmt
