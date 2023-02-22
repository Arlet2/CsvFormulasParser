go build -o bin/csvreader cmd/main.go
GOARCH=amd64 GOOS=windows go build -o bin/csvreader.exe cmd/main.go