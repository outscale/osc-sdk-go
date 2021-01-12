#!/bin/env bash
set -e

root=$(cd "$(dirname $0)/../.." && pwd)
# build new version number
local_sdk_version=$(cat $root/sdk_version)
local_sdk_version_major=$(echo $local_sdk_version | cut -d '.' -f 1)
local_sdk_version_minor=$(echo $local_sdk_version | cut -d '.' -f 2)
new_sdk_version_minor=$(( local_sdk_version_minor + 1 ))
new_sdk_version="$local_sdk_version_major.$new_sdk_version_minor.0"

branch_name="autobuild-$new_sdk_version"

if [ -z "$GH_TOKEN" ]; then
    echo "GH_TOKEN is missing, abort."
    exit 1
fi

result=$(curl -s -u LolUserName:$GH_TOKEN "https://api.github.com/repos/outscale/osc-sdk-go/pulls" | jq ".[] | select(.title == \"$new_sdk_version version\") | .title")

if [ ! -z "$result" ]; then
    echo "Pull request seems to alread exist, abort."
    exit 1
fi    
