#!/bin/bash

build () {
    local os=$1
    local arch=${2:-amd64}

    echo Build: $os/$arch

    GOOS=$os GOARCH=$arch go build -o build/$os/$arch/
}

pack () {
    local os=$1
    local arch=${2:-amd64}
    local ext=$3

    echo Pack: $os/$arch

    pushd build/$os/$arch > /dev/null
    rm -f build.zip
    zip build.zip beatriz$ext > /dev/null
    popd > /dev/null
}

copy () {
    local os=$1
    local arch=${2:-amd64}
    local dest=${3:-$os}

    echo Copy: $os/$arch

    cp build/$os/$arch/build.zip build/$dest.zip
}

build linux amd64
build windows amd64
build darwin amd64
build darwin arm64

pack linux amd64
pack windows amd64 .exe
pack darwin amd64
pack darwin arm64

copy linux amd64
copy windows amd64
copy darwin arm64 m1
