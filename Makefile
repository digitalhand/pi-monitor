build: 
	go build -o build/pimonitor cmd/pimonitor/main.go

compile: 
	GOOS=linux GOARCH=amd64 go build -o build/pimonitor-linux-amd64 cmd/pimonitor/main.go