build: 
	go build -o build/pimonitor cmd/pimonitor/main.go

compile: 
	GOOS=linux GOARCH=amd64 go build -o build/pimonitor-linux-amd64 cmd/pimonitor/main.go

compile-darwin: 
	GOOS=darwin GOARCH=arm64 go build -o build/pimonitor-darwin-arm64 cmd/pimonitor/main.go

compile-raspberry-pi:
	GOOS=linux GOARCH=arm GOARM=7 go build
