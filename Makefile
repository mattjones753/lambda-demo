all: deps build

default: all

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v

clean:
	rm -rf ./bin

build: clean
	mkdir -p bin
	set -e && for pkg in $$(ls src/lambdas); do \
		echo "\nbuilding: $$pkg\n"; \
		GOOS=linux CGO_ENABLED=0 go build -o ./bin/$$pkg ./src/lambdas/$$pkg; \
        zip -qj ./bin/$$pkg.zip ./configs/rclone.conf ./bin/$$pkg; \
	done