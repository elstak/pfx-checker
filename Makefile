all:
	GOOS=linux GOARCH=amd64 go build -a -o pfx-checker-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -a -o pfx-checker-darwin-amd64
	GOOS=windows GOARCH=amd64 go build -a -o pfx-checker-windows-amd64
