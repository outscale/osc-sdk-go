#!/bin/env bash
set -e

root=$(cd "$(dirname $0)/../.." && pwd)
if [ -e "$root/.auto-release-abort" ]; then
    echo "previous step triggered stop, abort"
    exit 0
fi
# build new version number
local_sdk_version=$(cat $root/sdk_version)
local_sdk_version_major=$(echo $local_sdk_version | cut -d '.' -f 1)
local_sdk_version_minor=$(echo $local_sdk_version | cut -d '.' -f 2)
new_sdk_version_minor=$(( local_sdk_version_minor + 1 ))
new_sdk_version="$local_sdk_version_major.$new_sdk_version_minor.0"

branch_name="autobuild-$new_sdk_version"
git branch -m $branch_name

echo "$new_sdk_version" > $root/sdk_version

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
git config user.name "Outscale Bot"
git config user.email "opensource+bot@outscale.com"
for f in osc v2; do
    git add $f || true
done
git commit -asm "osc-sdk-go v$new_sdk_version"
