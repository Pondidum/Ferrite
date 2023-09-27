#!/bin/sh

LEFT="1949D9C15328C68B"
RIGHT="49E171EC1481C453"

indent='s/^/    /g'
dev_path=/dev/disk/by-label/NICENANO

while true; do

  echo "==> Waiting for a NiceNano to be plugged in"

  while ! [ -h "${dev_path}" ]; do
    sleep 1s
  done

  echo "--> Plugged in"

  serial=$(/bin/udevadm info --name="${dev_path}" | sed -n 's/.*SERIAL_SHORT=\(.*\)/\1/p')
  firmware=""

  echo "    Serial: ${serial}"

  if [ "${serial}" = "${LEFT}" ]; then
    echo "    Left Half"
    firmware="bin/left.uf2"
  elif [ "${serial}" = "${RIGHT}" ]; then
    echo "    Right Half"
    firmware="bin/right.uf2"
  else
    echo "==> unknown keyboard half detected"
    exit 1
  fi

  if ! [ -d /media/andy/NICENANO ]; then
    udisksctl mount --block-device "${dev_path}" | sed "${indent}"
  fi

  echo "--> Copying ${firmware} to device"
  cp "${firmware}" "/media/andy/NICENANO/"
  echo "    Done"

  echo "--> Waiting for disconnect"
  while [ -h "${dev_path}" ]; do
    sleep 1s
  done
  echo "    Done"

done