#!/bin/env bash
set -e

root=$(cd "$(dirname $0)/../.." && pwd)

# build new version number
local_sdk_version=$(cat $root/sdk_version)
local_sdk_version_major=$(echo $local_sdk_version | cut -d '.' -f 1)
local_sdk_version_minor=$(echo $local_sdk_version | cut -d '.' -f 2)
new_sdk_version_minor=$(( local_sdk_version_minor + 1 ))
new_sdk_version="$local_sdk_version_major.$new_sdk_version_minor.0"

# TODO: check if version is not already pushed
branch_name="autobuild-$new_sdk_version"


git branch -m $branch_name

# build release notes
new_api_version=$(cat $root/api_version)
release_notes="# $new_sdk_version

 - SDK update for Outscale API v$new_api_version

"
echo "$release_notes$(cat $root/changelog.md)" > $root/changelog.md

# generate SDK
cd "$root"
make gen

# setup git && commit
git config --global user.name="Outscale Bot"
git config --global user.email="opensource+bot@outscale.com"
git add osc
git commit -asm "osc-sdk-go $new_sdk_version"
