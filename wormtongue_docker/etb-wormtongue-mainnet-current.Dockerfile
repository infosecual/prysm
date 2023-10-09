###############################################################################
#           Dockerfile to build all clients minimal mainnet preset.           #
###############################################################################
# Execution Clients
ARG GETH_REPO="https://github.com/ethereum/go-ethereum.git"
ARG GETH_BRANCH="master"

# All of the fuzzers we will be using
ARG TX_FUZZ_REPO="https://github.com/MariusVanDerWijden/tx-fuzz.git"
ARG TX_FUZZ_BRANCH="master"

ARG WORMTONGUE_REPO="https://github.com/infosecual/wormtongue.git"
ARG WORMTONGUE_BRANCH="wormtongue"
###############################################################################
# Builder to build all of the clients.
FROM debian:bullseye-slim AS etb-wormtongue-builder

# build deps
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    curl \
    libpcre3-dev \
    lsb-release \
    software-properties-common \
    apt-transport-https \
    openjdk-17-jdk \
    ca-certificates \
    wget \
    tzdata \
    bash \
    python3-dev \
    make \
    g++ \
    gnupg \
    cmake \
    libc6 \
    libc6-dev \
    libsnappy-dev \
    gradle \
    pkg-config \
    libssl-dev \
    git

WORKDIR /git

# set up go (geth+prysm)
RUN wget https://go.dev/dl/go1.20.3.linux-amd64.tar.gz
RUN tar -zxvf go1.20.3.linux-amd64.tar.gz -C /usr/local/
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go
RUN ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt
ENV PATH="$PATH:/root/go/bin"

############################# Execution  Clients  #############################
# Geth
FROM etb-wormtongue-builder AS geth-builder
ARG GETH_BRANCH
ARG GETH_REPO
RUN git clone "${GETH_REPO}" && \
    cd go-ethereum && \
    git checkout "${GETH_BRANCH}" && \
    git log -n 1 --format=format:"%H" > /geth.version

RUN cd go-ethereum && \
    go install ./...

# Wormtongue
FROM etb-wormtongue-builder AS wormtongue-builder
ARG WORMTONGUE_BRANCH
ARG WORMTONGUE_REPO

# uncomment the following line to force no-cache for wormtongue
ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" skipcache

RUN git clone "${WORMTONGUE_REPO}" && \
    cd wormtongue && \
    git checkout "${WORMTONGUE_BRANCH}" && \
    git log -n 1 --format=format:"%H" > /wormtongue.version && \
    mkdir bins && \
    go build -o bins ./...

############################### Misc.  Modules  ###############################
FROM etb-wormtongue-builder AS misc-builder
ARG TX_FUZZ_BRANCH
ARG TX_FUZZ_REPO
ARG BEACON_METRICS_GAZER_REPO
ARG BEACON_METRICS_GAZER_BRANCH

RUN go install github.com/wealdtech/ethereal/v2@latest \
    && go install github.com/wealdtech/ethdo@latest \
    && go install github.com/protolambda/eth2-val-tools@latest

RUN git clone "${TX_FUZZ_REPO}" && \
    cd tx-fuzz && \
    git checkout "${TX_FUZZ_BRANCH}"

RUN cd tx-fuzz && \
    cd cmd/livefuzzer && go build

########################### etb-all-clients runner  ###########################
FROM debian:bullseye-slim

WORKDIR /git

RUN apt update && apt install curl ca-certificates -y --no-install-recommends \
    wget \
    lsb-release \
    software-properties-common && \
    curl -sL https://deb.nodesource.com/setup_18.x | bash -

RUN wget https://packages.microsoft.com/config/debian/11/packages-microsoft-prod.deb -O packages-microsoft-prod.deb && \
    dpkg -i packages-microsoft-prod.deb && \
    rm packages-microsoft-prod.deb

RUN apt-get update && apt-get install -y --no-install-recommends \
    nodejs \
    libgflags-dev \
    libsnappy-dev \
    zlib1g-dev \
    libbz2-dev \
    liblz4-dev \
    libzstd-dev \
    openjdk-17-jre \
    dotnet-runtime-7.0 \
    aspnetcore-runtime-7.0 \
    python3-dev \
    python3-pip

RUN pip3 install ruamel.yaml web3

# for coverage artifacts and runtime libraries.
RUN wget --no-check-certificate https://apt.llvm.org/llvm.sh && \
    chmod +x llvm.sh && \
    ./llvm.sh 15

ENV LLVM_CONFIG=llvm-config-15

# misc tools used in etb
COPY --from=misc-builder /root/go/bin/ethereal /usr/local/bin/ethereal
COPY --from=misc-builder /root/go/bin/ethdo /usr/local/bin/ethdo
COPY --from=misc-builder /root/go/bin/eth2-val-tools /usr/local/bin/eth2-val-tools

# tx-fuzz
COPY --from=misc-builder /git/tx-fuzz/cmd/livefuzzer/livefuzzer /usr/local/bin/livefuzzer

# execution clients
COPY --from=geth-builder /geth.version /geth.version
COPY --from=geth-builder /root/go/bin/geth /usr/local/bin/geth

# prysm wormtongue
COPY --from=wormtongue-builder /git/wormtongue/bins/beacon-chain /usr/local/bin/wormtongue-beacon-chain
COPY --from=wormtongue-builder /git/wormtongue/bins/validator /usr/local/bin/wormtongue-validator
COPY --from=wormtongue-builder /wormtongue.version /wormtongue.version
