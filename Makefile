build:
	@go get .
	@go mod tidy
	@go build -ldflags="-X 'main.version=$$(git describe --abbrev=0 --tags)'" -o bin/straw
