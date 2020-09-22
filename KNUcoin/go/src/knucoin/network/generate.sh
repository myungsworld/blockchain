#!/bin/sh
export FABRIC_CFG_PATH=${PWD}
CHANNEL_NAME=channelsales1

# remove previous crypto material and config transactions
rm -fr config/*
rm -fr crypto-config/*

# generate crypto material
./bin/cryptogen generate --config=./crypto-config.yaml

# generate genesis block for orderer
#mkdir config
./bin/configtxgen -profile OrdererGenesis -outputBlock ./config/genesis.block

# generate channel 1 configuration transaction
./bin/configtxgen -profile Channel1 -outputCreateChannelTx ./config/channel1.tx -channelID $CHANNEL_NAME

# generate anchor peer transaction of SalesOrg
./bin/configtxgen -profile Channel1 -outputAnchorPeersUpdate ./config/SalesOrganchors.tx -channelID $CHANNEL_NAME -asOrg SalesOrg

# generate anchor peer transaction of CustomerOrg
./bin/configtxgen -profile Channel1 -outputAnchorPeersUpdate ./config/CustomerOrganchors.tx -channelID $CHANNEL_NAME -asOrg CustomerOrg