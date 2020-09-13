#!/bin/bash
set -ev

docker-compose -f docker-compose-ca.yaml down
docker-compose -f docker-compose-ca.yaml up -d

sleep 1
cd $GOPATH/src/knucoin/application/sdk
node enrollAdmin.js
sleep 1
node registUsers.js