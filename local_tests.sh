#!/bin/sh

# Copyright (c) Outscale SAS
#
# SPDX-License-Identifier: BSD-3-Clause

export OSC_PASSWORD='ashita wa dochida'
export OSC_LOGIN=joe
export OSC_TEST_SECRET_KEY=0000001111112222223333334444445555555666
export OSC_TEST_ACCESS_KEY=11112211111110000000

export OMI_ID="ami-90067666"

export OSC_SECRET_KEY=0000001111112222223333334444445555555666
export OSC_ACCESS_KEY=11112211111110000000

export OSC_ENDPOINT_ICU="http://127.0.0.1:3000/icu/"
export OSC_ENDPOINT_FCU="http://127.0.0.1:3000"
export OSC_ENDPOINT_API="http://127.0.0.1:3000/api/v1"
export OSC_ENDPOINT_API_NOPROTO="127.0.0.1:3000/api/v1"
export OSC_REGION="vp-ware-3"
export OSC_PROTOCOL="http"

export OSC_USING_RICOCHET="oui"

if [ "$#" -eq 0 ]; then

    if [ ! -d "osc-ricochet-2" ]; then
	git clone https://github.com/outscale-mgo/osc-ricochet-2
    fi

    cd osc-ricochet-2
    pkill ricochet

    cargo build --profile 'sdks'
    cargo run --profile 'sdks' -- ./ricochet.json &> /dev/null  &
    cd ..

    sleep 5
fi

set -e

make go-test
