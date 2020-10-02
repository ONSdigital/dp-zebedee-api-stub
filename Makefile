PHONY: build
build:
	go build -o dp-zebedee-api-stub

PHONY: debug
debug: build
	./dp-zebedee-api-stub