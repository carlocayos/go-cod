#!/bin/sh
set -eu

swagger generate client -A go-cod \
    -f specs/v1.0.0/cod-openapi-v1.0.0.yaml \
    -C specs/v1.0.0/swagger-gen-v1.0.0.yaml