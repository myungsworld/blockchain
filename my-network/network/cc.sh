docker exec cli peer chaincode install -n sacc -v 1.0 -p github.com

peer chaincode list --installe

docker exec cli peer chaincode instantiate -o orderer.example.com:7050 -n sacc -v 1.0 -C mychannel -c '{"Args":["a","100"]}' -P 'OR ("Org1MSP.member","Org2MSP.member","Org3MSP.member")'

sleep 10

docker exec cli peer chaincode list --instantiated -C mychannel

docker exec cli peer chaincode query -n sacc -C mychannel -c '{"Args":["get","a"]}'
docker exec cli peer chaincode invoke -n sacc -C mychannel -c '{"Args":["set","a","200"]}'
