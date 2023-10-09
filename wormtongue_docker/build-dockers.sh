#!/usr/bin/env bash
# Usage:
# $1: Image tag
# $2: Dockerfile
function build_image() {
    if [ "${REBUILD_IMAGES}" == 1 ]; then
        container_builder build --no-cache -t $1 -f $2 || echo "failed to rebuild $1"
    else
        container_builder build -t $1 -f $2 || echo "failed to build $1"
    fi
}
# Usage:
# $1: logging string.
function log_step() {
    echo "[ethereum-testnet-bootstrapper] — $1"
}

function container_builder() {
    if hash docker 2>/dev/null; then
        BUILDKIT=1 docker "$@" .
    else
        podman "$@"
    fi
}

log_step "building etb-all-clients:wormtongue-mainnet-current"
build_image "etb-all-clients:wormtongue-mainnet-current" "./wormtongue_docker/etb-wormtongue-mainnet-current.Dockerfile"
