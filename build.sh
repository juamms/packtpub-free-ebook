rm -rf "./bin"
mkdir "./bin"

env GOOS=linux GOARCH=amd64 go build -o ./bin/packtpub-linux-amd64 main.go
env GOOS=darwin GOARCH=amd64 go build -o ./bin/packtpub-macos-amd64 main.go
env GOOS=windows GOARCH=amd64 go build -o ./bin/packtpub-windows-amd64.exe main.go