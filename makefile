debug:
	go run .

release:
	go build -ldflags "-s -w"
