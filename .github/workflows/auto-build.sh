#!/bin/env bash
set -e

major_sdk_version=$1
if [ -z "$major_sdk_version"]; then
    echo "usage: $0 MAJOR_SDK_VERSION"
    echo "exemple: $0 v1"
    exit 1
fi

echo "done"
exit 0
