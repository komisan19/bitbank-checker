.PHONY: build
build: 
	docker build -t bitbank-check .

.PHONY: start
start: 
	docker run --rm --env-file=env bitbank-check
