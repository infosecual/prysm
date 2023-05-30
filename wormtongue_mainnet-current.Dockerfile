###############################################################################
#           Dockerfile to build wormtongue minimal                            #
###############################################################################
# Consensus Clients

ARG PRYSM_REPO="https://github.com/infosecual/prysm.git"
ARG PRYSM_BRANCH="wormtongue"

# Execution Clients
ARG GETH_REPO="https://github.com/ethereum/go-ethereum.git"
ARG GETH_BRANCH="master"

###############################################################################
# Builder to the wormtongue client
FROM debian:bullseye-slim AS wormtongue-builder

# build deps
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    curl \
    apt-transport-https \
    ca-certificates \
    wget \
    bash \
    git

#WORKDIR /git

# set up go (geth+prysm)
RUN wget https://go.dev/dl/go1.20.4.linux-amd64.tar.gz
RUN tar -zxvf go1.20.4.linux-amd64.tar.gz -C /usr/local/
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go
RUN ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt
ENV PATH="$PATH:/root/go/bin"

############################# Execution  Clients  #############################
# Geth
ARG GETH_BRANCH
ARG GETH_REPO
RUN git clone "${GETH_REPO}" && \
    cd go-ethereum && \
    git checkout "${GETH_BRANCH}" && \
    git log -n 1 --format=format:"%H" > /geth.version

RUN cd go-ethereum && \
    go install ./...

############################# Consensus  Clients  #############################
# PRYSM

ARG PRYSM_BRANCH
ARG PRYSM_REPO
RUN git clone "${PRYSM_REPO}" && \
    cd prysm && \
    git checkout "${PRYSM_BRANCH}" && \
    git log -n 1 --format=format:"%H" > /prysm.version

RUN cd prysm && mkdir bins && go build -o bins ./...

########################### etb-all-clients runner  ###########################
FROM debian:bullseye-slim

WORKDIR /git

# prysm wormtongue
COPY --from=wormtongue-builder /prysm/bins/beacon-chain /usr/local/bin/beacon-chain
COPY --from=wormtongue-builder /prysm/bins/validator /usr/local/bin/validator
COPY --from=wormtongue-builder /prysm.version /prysm.version

# execution clients
COPY --from=wormtongue-builder /geth.version /geth.version
COPY --from=wormtongue-builder /root/go/bin/geth /usr/local/bin/geth