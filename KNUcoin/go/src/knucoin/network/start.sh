#!/bin/bash


set -ev

function replacePrivateKey() {
    echo "ca key file exchange"
    cp docker-compose-template.yaml docker-compose.yaml
    PRIV_KEY1=$(ls crypto-config/peerOrganizations/sales.knucoin.com/ca/ | grep _sk)
    PRIV_KEY2=$(ls crypto-config/peerOrganizations/customer.knucoin.com/ca/ | grep _sk)
    sed -i '' "s/CA_SALES_PRIVATE_KEY/${PRIV_KEY1}/g" docker-compose.yaml
    sed -i '' "s/CA_CUSTOMER_PRIVATE_KEY/${PRIV_KEY2}/g" docker-compose.yaml
}

function checkPrereqs() {
    # check config dir
    if [ ! -d "crypto-config" ]; then
        echo "crypto-config dir missing"
        exit 1
    fi
    # check crypto-config dir
     if [ ! -d "config" ]; then
        echo "config dir missing"
        exit 1
    fi
}

checkPrereqs
replacePrivateKey

docker-compose -f docker-compose.yaml down

replacePrivateKey

docker-compose -f docker-compose.yaml up -d

# Create the channel
docker exec cli peer channel create -o orderer.knucoin.com:7050 -c channelsales1 -f /etc/hyperledger/configtx/channel1.tx
sleep 20
# Join peer0.sales.knucoin.com to the channel and Update the Anchor Peers in Channel1
docker exec cli peer channel join -b channelsales1.block
docker exec cli peer channel update -o orderer.knucoin.com:7050 -c channelsales1 -f /etc/hyperledger/configtx/SalesOrganchors.tx

# Join peer1.sales.knucoin.com to the channel
docker exec -e "CORE_PEER_ADDRESS=peer1.sales.knucoin.com:7051" cli peer channel join -b channelsales1.block

# Join peer0.customer.knucoin.com to the channel and update the Anchor Peers in Channel1
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.knucoin.com/users/Admin@customer.knucoin.com/msp" -e "CORE_PEER_ADDRESS=peer0.customer.knucoin.com:7051" cli peer channel join -b channelsales1.block
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.knucoin.com/users/Admin@customer.knucoin.com/msp" -e "CORE_PEER_ADDRESS=peer0.customer.knucoin.com:7051" cli peer channel update -o orderer.knucoin.com:7050 -c channelsales1 -f /etc/hyperledger/configtx/CustomerOrganchors.tx

# Join peer1.customer.knucoin.com to the channel
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.knucoin.com/users/Admin@customer.knucoin.com/msp" -e "CORE_PEER_ADDRESS=peer1.customer.knucoin.com:7051" cli peer channel join -b channelsales1.block

sleep 1
# Install and instantiate chaincode 

# peer0.sales.com
docker exec cli peer chaincode install -n knucoin-cc3 -v 1.04 -p chaincode/go
sleep 1

# peer1.sales.com
docker exec -e "CORE_PEER_ADDRESS=peer1.sales.knucoin.com:7051" cli peer chaincode install -n knucoin-cc3 -v 1.04 -p chaincode/go
sleep 1

# peer0.customer.com
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.knucoin.com/users/Admin@customer.knucoin.com/msp" -e "CORE_PEER_ADDRESS=peer0.customer.knucoin.com:7051" cli peer chaincode install -n knucoin-cc3 -v 1.04 -p chaincode/go
sleep 1

# peer1.customer.com
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.knucoin.com/users/Admin@customer.knucoin.com/msp" -e "CORE_PEER_ADDRESS=peer1.customer.knucoin.com:7051" cli peer chaincode install -n knucoin-cc3 -v 1.04 -p chaincode/go
sleep 1

docker exec cli peer chaincode instantiate -o orderer.knucoin.com:7050 -C channelsales1 -n knucoin-cc3 -v 1.04 -c '{"Args":[""]}' -P "OR ('SalesOrg.member','CustomerOrg.member')"

cd ../application/sdk
node enrollAdmin.js