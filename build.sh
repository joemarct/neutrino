#!/bin/bash

if [ -e ./a_main-packr.go ]; then
    rm a_main-packr.go
fi
packr

name="neutrino"
if [[ "$OSTYPE" == "darwin"* ]]; then
    go build -o dist/darwin/$name.app/Contents/MacOS/$name
elif [[ "$OSTYPE" == "linux-gnu" ]]; then
    go build -o dist/linux/$name
elif [[ "$OSTYPE" == "msys" ]]; then
    go build -ldflags="-H windowsgui" -o dist/windows/$name.exe
fi

open dist/
