#!/bin/bash

# Check that an argument was provided
if [ $# -eq 0 ]
  then
    echo "Please provide a vistim client to test."
    echo "If 'all' is provided then all CL clients will be tested with various ELs."
    echo "If a specific CL is provided then Geth will be the EL client."
    echo ""
    echo "Usage:"
    echo '    ./setup_wormtongue_etb.sh <prysm|lighthouse|teku|nimbus|lodestar|all>'
    exit 1
fi

# Check that the argument is a supported CL client or "all"
supported=false
if [ $1 != "prysm" -a $1 != 'dir' -a $1 != 'lighthouse' -a $1 != 'teku' -a $1 != 'nimbus' -a $1 != 'lodestar' -a $1 != 'all' ]
 then
    echo "Err: '$1' is not a supported option."
    echo ""
    echo "Please provide a vistim client to test."
    echo "If 'all' is provided then all CL clients will be tested with various ELs."
    echo "If a specific CL is provided then Geth will be the EL client."
    echo ""
    echo "Usage:"
    echo '    ./setup_wormtongue_etb.sh <prysm|lighthouse|teku|nimbus|lodestar|all>'
    exit 1
fi

# Download the required repos
# NOTE this should change back to tyler's repo once he accepts the PR here:
# https://github.com/0xTylerHolmes/ethereum-testnet-bootstrapper/pull/14
git clone git@github.com:infosecual/ethereum-testnet-bootstrapper.git
git clone git@github.com:infosecual/etb-fuzzer-images.git

# Build the Generic client and Wormtongue images
docker build -t etb-all-clients:latest -f etb-fuzzer-images/wormtongue/deps/dockers/etb-all-clients_generic_no-peer-scoring.Dockerfile .
docker build -t etb-wormtongue:latest -f etb-fuzzer-images/wormtongue/deps/dockers/wormtongue.Dockerfile .

# Move the configs and the launchers into ethereum-testnet-bootstrapper
cp -r etb-fuzzer-images/wormtongue/deps/* ethereum-testnet-bootstrapper/deps
cp etb-fuzzer-images/wormtongue/configs/mainnet-current-wormtongue-$1.yaml ethereum-testnet-bootstrapper/configs
cp etb-fuzzer-images/wormtongue/apps/* ethereum-testnet-bootstrapper/apps

# Prepair ETB for Wormtongue
cd ethereum-testnet-bootstrapper
make build-bootstrapper
make clean config=configs/mainnet-current-wormtongue-$1.yaml
make init-testnet config=configs/mainnet-current-wormtongue-$1.yaml

# print instuctions to run
echo "#########################################################################"
echo "# ETB is now set up to run. Run the following to start the network:"
echo "#########################################################################"
echo ""
echo "cd ethereum-testnet-bootstrapper/ && docker compose up --force-recreate --remove-orphans"
echo ""
echo "#########################################################################"
echo "# To attach to the health checker:"
echo "#########################################################################"
echo ""
echo "docker logs status-check-0 -f"
echo ""
