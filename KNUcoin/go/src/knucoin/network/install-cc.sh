#!/bin/bash
set -ev

# install chaincode for channelsales1
docker exec cli peer chaincode install -n knucoin-cc -v 1.02 -p chaincode/go
sleep 1
docker exec cli peer chaincode instantiate -o orderer.knucoin.com:7050 -C channelsales1 -n knucoin-cc -v 1.02 -c '{"Args":[""]}' -P "OR ('SalesOrg.member','CustomerOrg.member')"
# instantiate chaincode for channelsales1
docker exec cli peer chaincode instantiate -o orderer.knucoin.com:7050 -C channelsales1 -n knucoin-cc -v 1.02 -c '{"Args":[""]}' -P "OR ('SalesOrg.member','CustomerOrg.member')"
sleep 10
# invoke chaincode for channelsales1
docker exec cli peer chaincode invoke -o orderer.knucoin.com:7050 -C channelsales1 -n knucoin-cc3 -c '{"function":"initWallet","Args":["song"]}'
docker exec cli peer chaincode invoke -o orderer.knucoin.com:7050 -C channelsales1 -n knucoin-cc3 -c '{"function":"chargeMoney","Args":["song","1000"]}'
docker exec cli peer chaincode invoke -o orderer.knucoin.com:7050 -C channelsales1 -n knucoin-cc3 -c '{"function":"transferMoney","Args":["song","min","100"]}'
sleep 3
# query chaincode for channelsales1
docker exec cli peer chaincode query -o orderer.knucoin.com:7050 -C channelsales1 -n knucoin-cc3 -c '{"function":"getWallet","Args":["song"]}'
