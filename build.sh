#!/bin/bash

build () {
    echo Build: $1/$2

    GOOS=$1 GOARCH=$2 go build -o build/$1/$2/
}

pack () {
    echo Pack: $1/$2

    pushd build/$1/$2 > /dev/null
    rm -f build.zip
    zip build.zip beatriz$3 > /dev/null
    popd > /dev/null
}

copy () {
    echo Copy: $1/$2

    cp build/$1/$2/build.zip build/$3.zip
}

build linux amd64
build windows amd64
build darwin amd64
build darwin arm64

pack linux amd64
pack windows amd64 .exe
pack darwin amd64
pack darwin arm64

copy linux amd64 linux
copy windows amd64 windows
copy darwin arm64 m1
