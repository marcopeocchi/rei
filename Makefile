multiarch:
	GOOS=linux GOARCH=arm go build -o rei_linux-arm *.go
	GOOS=linux GOARCH=arm64 go build -o rei_linux-arm64 *.go
	GOOS=linux GOARCH=amd64 go build -o rei_linux-amd64 *.go
	mkdir -p build
	mv rei_* build