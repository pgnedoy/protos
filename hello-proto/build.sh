#!/bin/sh
set -e

VERSION=1.9.0

docker pull uber/prototool:$VERSION
docker run --rm -it -v "$(pwd)":/defs --workdir /defs uber/prototool:$VERSION prototool generate .
