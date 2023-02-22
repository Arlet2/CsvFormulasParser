cd bin
go build -o csvreader ../cmd/main.go
GOARCH=amd64 GOOS=windows go build -o csvreader.exe ../cmd/main.go