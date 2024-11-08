#!/bin/bash

set -euo pipefail

build(){
    echo "======================="
    echo "Info: build image..."
    echo "=======================" && docker build -t $IMAGE:$TAG .
}

push () {
    echo "======================="
    echo "Info: upload artifact registery..."
    echo "=======================" && docker push $IMAGE:$TAG
}

export IMAGE=gcr.io/media17-prod/17media/migration-base

export TAG="v$(date +"%y.%m.%d.$(date +%s)")"

build
push
