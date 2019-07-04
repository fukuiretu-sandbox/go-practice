clean:
	rm -rf dest/main

build: clean
	env GOARCH="amd64" GOOS="linux" go build -o dist/main
