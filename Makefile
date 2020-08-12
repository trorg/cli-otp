release:
	mkdir -p bin/release/
	go build -ldflags "-s -w" -o bin/release/

all:
	mkdir -p /bin/dev
	go build -o bin/dev
