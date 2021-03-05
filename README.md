
```
docker run --rm -v "$PWD/src":/usr/src/myapp -v "$PWD/bin":/usr/src/bin -w /usr/src/myapp -e GOOS=windows -e GOARCH=amd64 golang:1.14 go build -v -o /usr/src/bin/gobot.exe
```