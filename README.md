# Wormtongue
An evil Prysm node for fuzzzing the ethereum CL request/response message domain. Named after [Grima Wormtongue](https://en.wikipedia.org/wiki/Gr%C3%ADma_Wormtongue)

## Getting Started
### Testing code modifications
Wormtongue is designed to be tested in local testnets via [Ethereum Testnet Bootstrapper](https://github.com/0xTylerHolmes/ethereum-testnet-bootstrapper/) (ETB) which builds wormtongue in docker. If you want to test that modifications build correctly you can do so locally on linux eg.:

```
mkdir builds
go build -o builds ./...
```
### Setting up ETB
You will need two repos to run Wormtongue in ETB, the main ETB repo and the wormtongue config/docker repo. Here is an example workflow of getting a Wormtongue fuzzer going:

Download the required repos
```
# NOTE this should change back to tyler's repo once he accepts the PR here:
# https://github.com/0xTylerHolmes/ethereum-testnet-bootstrapper/pull/14
# git clone git@github.com:0xTylerHolmes/ethereum-testnet-bootstrapper.git
git clone git@github.com:infosecual/ethereum-testnet-bootstrapper.git
# git clone git@github.com:0xTylerHolmes/etb-fuzzer-images.git
git clone git@github.com:infosecual/etb-fuzzer-images.git
```

Build the Wormtongue image
```
docker build -t etb-all-clients:latest -f etb-fuzzer-images/wormtongue/deps/dockers/etb-all-clients_mainnet-wormtongue.Dockerfile .
```

Move the configs and the launchers into ethereum-testnet-bootstrapper
```
cp -r etb-fuzzer-images/wormtongue/deps/* ethereum-testnet-bootstrapper/deps
cp etb-fuzzer-images/wormtongue/configs/mainnet-current-wormtongue.yaml ethereum-testnet-bootstrapper/configs
cp etb-fuzzer-images/wormtongue/apps/* ethereum-testnet-bootstrapper/apps
```

Prepair ETB for Wormtongue
```
cd ethereum-testnet-bootstrapper
make build-bootstrapper
make clean config=configs/mainnet-current-wormtongue.yaml
make init-testnet config=configs/mainnet-current-wormtongue.yaml
```

A small script to do all of this can be found in this directory: `setup_wormtongue_etb.sh`

### Running ETB

To bring up the network and start fuzzing (you may want to do this in a screen session if you plan to leave it running on a remote server or anything):
```
docker compose up --force-recreate --remove-orphans
```

To attach to the health checker:
```
docker attach status-check-0
```
