#!/bin/env bash
set -e

root=$(cd "$(dirname $0)/../.." && pwd)
if [ -e "$root/.auto-release-abort" ]; then
    echo "previous step triggered stop, abort"
    exit 0
fi
github_url="https://api.github.com/repos/outscale/osc-api/releases"

if [ -z "$GH_TOKEN" ]; then
    echo "GH_TOKEN is missing, abort."
    exit 1
fi

osc_api_last_release=$(curl -s -H "Authorization: token $GH_TOKEN" $github_url | jq ".[] | select(.prerelease == false) | select(.draft == false) | .tag_name" | sort -r --version-sort | head -n 1 | cut -f 2 -d '"')
local_api_version=$(cat $root/api_version)

echo "last API version used: $local_api_version"
echo "last available version: $osc_api_last_release"

if [[ "$local_api_version" = "$osc_api_last_release" ]]; then
    echo "no update needed, exiting"
    touch "$root/.auto-release-abort"
    exit 0
fi

echo "$osc_api_last_release" > $root/api_version
