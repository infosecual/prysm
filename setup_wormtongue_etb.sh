# Download the required repos
# NOTE this should change back to tyler's repo once he accepts the PR here:
# https://github.com/0xTylerHolmes/ethereum-testnet-bootstrapper/pull/14
git clone git@github.com:infosecual/ethereum-testnet-bootstrapper.git
git clone git@github.com:infosecual/etb-fuzzer-images.git

# Build the GEneric client and Wormtongue images
docker build -t etb-all-clients:latest -f etb-fuzzer-images/wormtongue/deps/dockers/etb-all-clients_mainnet-wormtongue.Dockerfile .
docker build -t etb-wormtongue:latest -f etb-fuzzer-images/wormtongue/deps/dockers/wormtongue_mainnet.Dockerfile .

# Move the configs and the launchers into ethereum-testnet-bootstrapper
cp -r etb-fuzzer-images/wormtongue/deps/* ethereum-testnet-bootstrapper/deps
cp etb-fuzzer-images/wormtongue/configs/mainnet-current-wormtongue.yaml ethereum-testnet-bootstrapper/configs
cp etb-fuzzer-images/wormtongue/apps/* ethereum-testnet-bootstrapper/apps

# Prepair ETB for Wormtongue
cd ethereum-testnet-bootstrapper
make build-bootstrapper
make clean config=configs/mainnet-current-wormtongue.yaml
make init-testnet config=configs/mainnet-current-wormtongue.yaml
