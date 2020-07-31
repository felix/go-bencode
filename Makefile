
.PHONY: test
test: lint
	go test -short -coverprofile=coverage.txt -covermode=atomic ./... \
		&& go tool cover -func=coverage.txt

.PHONY: lint
lint: ; revive ./...

.PHONY: bench
bench: ; go test -bench=.

.PHONY: clean
clean: ; rm -rf coverage*
