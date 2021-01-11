#!/bin/env bash
set -e

root=$(cd "$(dirname $0)/../.." && pwd)
new_sdk_version=$(cat $root/sdk_version)
branch_name="autobuild-$new_sdk_version"

if [ -z "$GITHUB_TOKEN" ]; then
    echo "GITHUB_TOKEN is missing, abort."
    exit 1
fi

# TODO push branch_name

