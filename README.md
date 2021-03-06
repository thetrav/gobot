Build for windows
```
 docker run --rm -it -v "$PWD/src":/usr/src/myapp -v "$PWD/bin":/usr/src/bin -w /usr/src/myapp -e GOOS=windows -e GOARCH=amd64 x1unix/go-mingw:1.14 go build -v -o /usr/src/bin/gobot.exe
```

Shell
```
docker run --rm -it -v "$PWD/src":/usr/src/myapp -v "$PWD/bin":/usr/src/bin -w /usr/src/myapp -e GOOS=windows -e GOARCH=amd64 x1unix/go-mingw:1.14 /bin/sh
```

Build for OSX
Assuming you already have go installed:
```
go build -v -o ../bin
```



Dependencies

The go library is poorly released and you need to explicitly tag it with master (should be fine now that it's in go.mod)
```
 go get github.com/g3n/engine@master
```

For windows you will need:
```
OpanAL32.dll
libogg.dll
libvorbis.dll
libvorbisfile.dll
```
which are available from https://github.com/g3n/engine/tree/master/audio/windows/bin

They just need to be on the path, but I like to put them in the same folder as the executable

for OSX you will need:
```
brew install libvorbis openal-soft
```