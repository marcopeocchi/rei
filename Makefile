multiarch:
	GOOS=linux GOARCH=arm go build -o valeera_linux-arm *.go
	GOOS=linux GOARCH=arm64 go build -o valeera_linux-arm64 *.go
	GOOS=linux GOARCH=amd64 go build -o valeera_linux-amd64 *.go
	mkdir -p build
	mv valeera_* build