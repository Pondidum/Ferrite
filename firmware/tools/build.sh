#!/bin/sh

set -eu

cd /west

cp /config/* /west/config

west build -s zmk/app -b "nice_nano_v2" --build-dir build/zephyr/left -- -DZMK_CONFIG=/west/config -DSHIELD="cradio_left"
mv build/zephyr/zmk.uf2 bin/left.uf2
rm -rf build

west build -s zmk/app -b "nice_nano_v2" -- -DZMK_CONFIG=/west/config -DSHIELD="cradio_right"
mv build/zephyr/zmk.uf2 bin/right.uf2
rm -rf build
