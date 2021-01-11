#!/bin/bash
set -e
root=$(cd "$(dirname $0)/../.." && pwd)
github_url="https://api.github.com/repos/outscale/osc-api/releases"

if [ -z "$GITHUB_TOKEN" ]; then
    echo "GITHUB_TOKEN is not set, abort."
    exit 1
fi

osc_api_last_release=$(curl -s -u LolUserName:$GITHUB_TOKEN $github_url | jq ".[] | select(.prerelease == false) | select(.draft == false) | .tag_name" | sort -r | head -n 1 | cut -f 2 -d '"')
local_api_version=$(cat $root/api_version)

echo "last API version used: $local_api_version"
echo "last available version: $osc_api_last_release"

if [[ "$local_api_version" = "$osc_api_last_release" ]]; then
    echo "no update needed, exiting"
    exit 1
fi

echo $osc_api_last_release > $root/osc_sdk_go
