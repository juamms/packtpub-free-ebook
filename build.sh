rm -rf "./bin"
mkdir "./bin"

env GOOS=linux GOARCH=amd64 go build -o ./bin/packtpub-linux-arm64 main.go
env GOOS=darwin GOARCH=amd64 go build -o ./bin/packtpub-macos-arm64 main.go
env GOOS=windows GOARCH=amd64 go build -o ./bin/packtpub-windows-arm64.exe main.go