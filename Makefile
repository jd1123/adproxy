default:
	all

test:
	go test ./...

build:
	go build

all:
	test
	build
