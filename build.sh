#!/bin/bash

build () {
    echo Build: $1/$2

    GOOS=$1 GOARCH=$2 go build -o build/$1/$2/
}

pack () {
    echo Pack: $1/$2

    pushd build/$1/$2 > /dev/null
    rm build.zip
    zip build.zip $3 > /dev/null
    popd > /dev/null
}

copy () {
    echo Copy: $1/$2

    cp build/$1/$2/build.zip build/$1.zip
}

build linux amd64
build windows amd64
build darwin amd64
build darwin arm64

pack linux amd64 beatriz
pack windows amd64 beatriz.exe
pack darwin amd64 beatriz
pack darwin arm64 beatriz

copy linux amd64
copy windows amd64
copy darwin arm64