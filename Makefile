.PHONY: build clean

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/dcsv ./
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/dcsv-mac ./
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/dcsv-win ./
	@echo Done!
clean:
	rm -rf ./build/dcsv
