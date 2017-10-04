# Make a release with
# make -j4 release

VERSION=$(shell git describe)
LDFLAGS=-ldflags "-s -w -X main.version=${VERSION}"

.PHONY: build
build:
	go build ${LDFLAGS}

.PHONY: linuxarm
linuxarm:
	env GOOS=linux GOARCH=arm go build ${LDFLAGS} -o dist/readable_linux_arm
	#cd dist && upx --brute readable_linux_arm

.PHONY: linux64
linux64:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o dist/readable_linux_amd64
	#cd dist && upx --brute readable_linux_amd64

.PHONY: windows
windows:
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o dist/readable_windows_amd64.exe
	#cd dist && upx --brute readable_windows_amd64.exe

.PHONY: osx
osx:
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o dist/readable_osx_amd64
	#cd dist && upx --brute readable_osx_amd64

.PHONY: release
release: osx windows linux64 linuxarm


