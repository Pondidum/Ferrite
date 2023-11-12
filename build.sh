#!/bin/sh

docker run --rm -it -v $PWD/config:/config -v $PWD/bin:/west/bin zmkbuilder:local
