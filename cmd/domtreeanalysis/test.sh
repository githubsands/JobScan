#!/usr/bin/env bash

# run script as ./domtreeanalysis <test #>
TEST=${1}

FILE="domtreeanalysis"
if [[ ! -f "$FILE" ]]; then
    go build
fi

if [ $TEST -eq 1 ]; then
    # Test 1: url
    PARAMS=$(echo \
        "-url=http://yahoo.com" \
    )
    RES=$(./domtreeanalysis $PARAMS)
    echo $RES
fi


if [ $TEST -eq 2 ]; then
    URL="https://www.predictit.org/"
    PARAMS=$(echo \
        "-stdout=on" \
        "-url=${URL}" \
    )
    RES=$(./domtreeanalysis $PARAMS)
    echo $RES
fi


if [ $TEST -eq 3 ]; then
    TAG="pageKey"
    URL="https://www.linkedin.com/jobs/search/?keywords=denver%20software%20engineer"

    PARAMS=$(echo \
        "-tag=${TAG}" \
        "-url=${URL}" \
    )
    RES=$(./domtreeanalysis $PARAMS)
    echo $RES
fi

if [ $TEST -eq 4 ]; then
    STRING="test"

    URL="https://www.linkedin.com/jobs/search/?keywords=denver%20software%20engineer"
    PARAMS=$(echo \
        "-stdout=on" \
        "-url=${URL}" \
        "-parse=${STRING}" \
    )
    RES=$(./domtreeanalysis $PARAMS)
    echo $RES
fi
