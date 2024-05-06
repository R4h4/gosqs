pretest:
	@docker run --rm -p 4100:4100 admiralpiett/goaws

test:
	@go test ./...

