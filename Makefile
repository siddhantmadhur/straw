build:
	@go get .
	@go mod tidy
	@go build -o bin/straw
