#!/bin/bash
set -e

exit_code=0
if [ -z "$OSC_ACCESS_KEY" ]; then
    echo "OSC_ACCESS_KEY var is not set, please fix"
    exit_code=1
fi
if [ -z "$OSC_SECRET_KEY" ]; then
    echo "OSC_SECRET_KEY var is not set, please fix"
    exit_code=1
fi
if [ -z "$OSC_REGION" ]; then
    echo "OSC_REGION var is not set, please fix"
    exit_code=1
fi
if [ -z "$OMI_ID" ]; then
    echo "OMI_ID var is not set, please fix"
    exit_code=1
fi
exit $exit_code