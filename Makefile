release:
	mkdir -p bin/release/
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/release/otp-cli-darwin-amd64/otp-cli
	cp README.md bin/release/otp-cli-darwin-amd64/
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/release/otp-cli-linux-amd64/otp-cli
	cp README.md bin/release/otp-cli-linux-amd64/

all:
	mkdir -p /bin/dev
	go build -o bin/dev
