# blockchain

[6 strategies to successfully implement blockchain](https://www.fm-magazine.com/news/2018/nov/how-to-implement-blockchain-201820113.html)

## 블록체인 할짝

금융기관이 가진 문제점을 해결할수 있는 거의 모든 기술을 가짐   
블록체인 개념을 증명하려고 노력하나 많은 이유에서 연구가 중단되어지고 있음. 
전세계적으로 84퍼의 기관들이 블록체인 기술을 포함하고 사용하지만 그중 15퍼센트만이 **live project** 를 가짐 (그냥 존나 어렵다는 뜻)  
가트너가 2030년까지 블록체인 기술의 값어치는 3조달러를 넘는다고 말함  

## 그래서 블록체인이 뭔가?
잘 짜여진 블록체인은 말그대로 블록에 기록되어 저장된다.  
그 블록들은 암호화 기술을 적용해 디지털 데이터와 나뉘어진 원장(가계부같은거) 연결되어있다.
그 ledger는 접근하는 누구한테나 공유,증명 되어진다. (Third party의 개입없이) 중간 역할을 하는 기관을 말하는듯 암호화"signature"는 블록을 생성하고 각각의 이전의 블록과 연결하는데 이 가상의 블록은 남이 함부로 변경하지 못하게된다.  
블록체인기술은 보안에 있어서 아주 강하고 유연하지만 보이지 않고 기술에 대한 지식 부족이 이 기술의 장애요소로 남아있다. 
  
[learn blockchains by building one](https://hackernoon.com/learn-blockchains-by-building-one-117428612f46). 블록체인 만들기 ㄱ 

## 블록체인 예

```p
block = {
    'index': 1,
    'timestamp': 1506057125.900785,
    'transactions': [
        {
            'sender': "8527147fe1f5426f9dd545de4b27ee00",
            'recipient': "a77f5cdfa2934df3954a5c7c7da5df1f",
            'amount': 5,
        }
    ],
    'proof': 324984774000,
    'previous_hash': "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
}
```
각각의 블록은 인덱스, 타임스탬프, 트랜잭션,프루프,이전의 해시값을 가진다.    
여기서 체인에 대한 아이디어는 명백하다 각각의 새로운 블록은 이전블록의 해시를 포함한다.  
이것이 블록체인이 **변경 불가능** 하게 만들어 준단다.  
만약 공격자가 체인의 이전 블록들을 붕괴시키려 한다면, 모든 후속 블록들이 부정확한 해시 값을 참조할것이다.  
응 뭔소린지 몰라~  

# 하이퍼레저 패브릭

* 핵심 용어  
  - __채널__ : 트랜잭션의 접근 권한을 그룹별로 설정  
  - __조직__ : 조직별 피어노드 관리 보증 정책 수행  
  - __노드__ : 그냥 서버라고 보면됨 블록체인은 수많은 서버들이 관리함  
  - __오더링 서비스 노드__(오더러) : 네트워크에 참여하는 모든 피어 노드의 분기 및 정렬 역할 
  - __피어 노드__  
    - 엔도싱 피어노드 : 보증 정책에 따라 트랜잭션 검토 후 보증 사인을 첨부  
    - 커미팅 피어노드 : 트랜잭션 결과에 문제가 없으면 트랜잭션 확정한 후 블록체인에 업데이트  
    - 앵커 피어노드 : 채널 대표 피어 노드
    - 리더 피어노드 : 조직 대표 피어 노드 
 
  
  
## 플로우 차트
  ![Fabric Flowchart](https://user-images.githubusercontent.com/56465854/89724240-6a536f80-da3b-11ea-8799-13de458e43f3.png)
  
~~예제가 별로 맘에 안드는 느낌이라 내가 만들어봤는데 확실히 예제는 예제다. 예제를 보고하자 ^^~~  

1. MSP로부터 접근을 승인 받은 애플리케이션은 하이퍼레저 패브릭 네트워크에 트랜잭션을 발생시킬수 있다.  
2. 트랜잭션 제안서를 __Endoring peer Node__ 에게 보낸다. 내용으로는 클라이언트의 ID,트랜잭션을 제출하는 서명 등이 있다.  
3. 제안을 받은 엔도싱 피어는 서명이 유효한지 체크한 후 __체인코드를 실행__ 한다. 이후 RW세트를 포함하는 트랙잭션 결과를 생성해 반환한다.  
4. 트랜잭션의 신뢰성을 확보했으면, 트랜잭션을 RW세트와 함께 노드들의 서명,채널ID를 오더링 서비스에 전송한다.  
5. OSN에서 생성된 트랜잭션 블록을 엔도싱과 커미팅 피어 노드들에게 전송하고 __각 노드들은 RW세트와 현재 World State가 일치하는지 유효성 검사 진행__   
6. __유효성이 검증된 트랜잭션은 모든 피어노드 데이터베이스에 기록__ 하고 성공 또는 실패를 애플리케이션에게 알린다.   

## 하이퍼레저 핵심 도구 
- __Cryptogen__ 은 인증서 발급 도구  
  * __cryptogen generate__ <<args>>  
  * --config : 인증서 발급을 위한 설정 파일을 입력  
  * --output : 인증서 파일의 저장 위치 지정 default = crypto-config  
  * __cryptogen extend__ <<ages>>  
  * --input : 기존에 발급된 인증서 파일의 저장 위치 불러옴  

- __Configtxgen__ 은 블록 및 트랜잭션 파일을 생성하는 도구  
  * -asOrg : 조직이름  
  * -channelID : 채널ID  
  * -inspectBlock : 지정된 경로에 있는 블록에 포함된 구성  
  * -inspectChannelCreateTx : 채널 정보에 대한 트랜잭션 구성  
  * -outputAnchorPeersUpdate : 앵커 피어 노드 업데이트 정보가 담긴 트랜잭션 파일이 저장될 경로  
  * -outputBlock : 제네시스 블록 파일이 저장될 경로  
  * -outputCreateChannlTx : 채널 생성 정보가 담긴 트랜잭션 파일이 저장될 경로  
  
- __Configtxlator__ 블록 및 트랙잭션 구문 분석 도구(REST 서버 구성)
  * start : REST server 시작  
  * Proto_encode : JSON 형식을 protobuf로 변환  
  * proto_decode : protobuf 를 JSON 형식으로 변환   
  
_ __Peer__ 블록체인 데이터를 저장하고 유지보수 체인코드를 실행하는 마스터 노드 도구  
  * chaincode : 체인코드 관련 작업  
    - 패키지 배포 명령어  
    - package, signpackage, install  
    - 생성 및 업데이트  
    - instantiate, update  
    - 실행 명령어  
    - invoke,query  
  * channel : 채널 관련 작업  
    - create,fetch,getinfo,join,list,signconfigtx,update  
  * logging : 로그 수준 작업  
  * node : 노드 서버 시작  
  * version : 피어 노드 버전 정보 표시  
- __Orderer__ 트랜잭션 패키징,정렬 도구  

## MSP
하이퍼레저 모든 기능은 MSP정보를 토대로 동작한다.  
MSP는 쉽게 생각하면 인증서를 만들어주는 거고 CA는 그 인증서를 관리하는 서버다.  
네트워크,채널,피어,오더러 모두 MSP를 가지고 
피어와 오더러 MSP는 로컬에서 파일 시스템에 탑재 되어 사용하고  
네트워크와 채널 MSP는 글로벌로 모든 구성원들을 식별할때 사용되고 이 정보는 제네시스블록과 채널 트랜잭션에 포함된다.  

## HLF Ledger
Blockchain + database
read/write set are written to the blockchain


## [RAFT 오더러](https://github.com/IBM/raft-fabric-sample) 

## Private Data
조직 데이터를 비공개로 유지해야 하는 경우의 별도의 채널을 만들지 않고 Private 데이터를 수집,커밋,쿼리할 수 있는 채널의 조직에 대해 정의된 하위 집합을 허용하는 데이터 컬렉션  
쉽게말해 조직이 두개 있을때 하나의 조직에 프라이빗 데이터를 추가하면 조회를 했을때 그 조직에만 추가되서 보여주는 데이터를 말함  
프라이빗 데이터는 Side database에 저장된다.  
[Private Data 적용](https://miiingo.tistory.com/193)  

PDC JSON 설정파일  
```json
[
 {
   "name": "collectionMarbles",
   "policy": "OR('Org1MSP.member', 'Org2MSP.member')",
   "requiredPeerCount": 0,
   "maxPeerCount": 3,
   "blockToLive":1000000,
   "memberOnlyRead": true
},
 {
   "name": "collectionMarblePrivateDetails",
   "policy": "OR('Org1MSP.member')",
   "requiredPeerCount": 0,
   "maxPeerCount": 3,
   "blockToLive":3,
   "memberOnlyRead": true
 }
]
```
requiredPeerCount : 적어도 이정도 피어는 프라이빗 데이터를 가지고 있어야 한다.  
maxPeerCount : 최대 이정도 피어가 프라이빗 데이터를 복제하고 있으면 된다. 보안에 중요     


## Private Data Collections(PDC)
![PDC FLOW CHART](https://user-images.githubusercontent.com/56465854/92067009-3bcc7880-edde-11ea-8cdf-e9700cac66eb.PNG)

개인정보나 지울수있는 데이터는 private state에 저장한다.   
이 정보는 ledger에 해시값으로 저장된다. 
block to live 로 수명을 결정한다.  
## Protocol

Gossip: 속도가 느리고 조용함  
Brodcast : 속도가 빠르고 노이지함

## fabric-ca-client
1.webserer에서 prviate ,public 키를 생성  
2.ca가 그걸 받아서 ca 인증서를 생성해서 서명
3. webserver로 보냄 (인증서를 받음) 


## 개발 일지

### [Chaincode](https://github.com/myungsworld/blockchain/blob/master/go/src/coin.go) 

world state를 저장하는 데이터베이스는 couchDB  
document 기반 데이터베이스이고 몽고DB랑 비슷하며 Json이랑 연동된다.  
Init 메서드에서 putState 를 쓰면 자동으로 world state key를 업데이트 시켜준다.  
query 메서드에서 getState 로 DB에서 업데이트된 키를 가져온다.  


체인코드에선 putstate를 하고 난후 getstate를 바로 쓸수 없다  
putstate를 했다고 world state에 바로 업데이트 되는게 아니기 때문이다.  
putstate를 하고 난후 피어가 블록에 있는 트랜잭션을 읽은다음 변경된다.  
만약 put을 하고 바로 뒤에 get을 한다면 그 state는 이전의 state를 가져오게 된다.  

peer chaincode list --installed  
어떤 체인코드가 설치 되었는지 확인  
peer chaincode list --instantiated -C [채널이름]  
배포된 체인코드 확인  

```
docker exec cli peer chaincode instantiate -o orderer.knucoin.com:7050 -C channelsales1 -n coin-cc -v 1.0 -c '{"Args":[""]}' -P "OR ('SalesOrg.member','CustomerOrg.member')"
```
인스턴스화 과정에서 나오는 에러  
```
Error: could not assemble transaction, err proposal response was not successful, error code 500, msg error starting container: error starting container: Failed to generate platform-specific docker build: Failed to pull hyperledger/fabric-ccenv:latest: API error (404): manifest for hyperledger/fabric-ccenv:latest not found: manifest unknown: manifest unknown
```
Failed to pull hyperledger/fabric-ccenv:latest: API error (404): manifest for hyperledger/fabric-ccenv:latest not found: manifest unknown: manifest unknown  
이 에러 고치다가 블록체인에 대한 벽을 느꼇다. 다시 벽을 오르자;;    
그냥 기존에 있는 hyperledger/fabric-ccenv 뒤에있는 태그를 latest로 바꿔주면 됐다  
```
docker tag 3d31661a812a hyperledger/fabric-ccenv:latest
```


FABRIC_CFG_PATH : configtx.yaml 파일이 존재하는 경로 - configtxgen 도구가 위치한 경로로 configtxgen 명령 실행 필요 $ export FABRIC_CFG_PATH=$PWD

블록체인 설치과정에서 계속 내가 만든 체인코드가 설치가 되지않고 예제가 설치가 됐다. 고생을 몇시간 하고 난후 -v 1.0 이 버전부분을 고쳐야 비로서 내가 만든 체인코드가 설치가 된다는걸 알았다.  
(Px. 체인코드 이름을 바꾸고 도커 컨테이너와 이미지를 모두 지우고 해봐도 이미 저 버전으로 하면 예제가 설치된다.)  

docker-compose.yaml 파일을 도커 컨테이너로 올릴때 설정된 네트워크 이름이랑 달라서 나오는 에러  
```
API404 Error network not found. 
CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=를 network이름에 맞추면 된다.  
```

체인코드 카우치디비 연동  
yaml 파일에 각 피어마다 couchdb 연동. 
environment - CORE_LEDGER_STATE_STATEDATABASE=CouchDB.  
depens on - coundb(index)  

카우치 디비 컨테이너 생성  
```
couchdb(index):
  container_name: couchdb(index)
  image: hyperledger/fabric-couchdb
  environment:
    - COUCHDB_USER=
    - COUCHDB_PASSWORD=
  ports:
    - 8984:5984
  networks:
    - network name
```

걍 체인코드 잘못적으면 나오는 에러임 err 를 nil 에 유효성검사하는걸 그냥 err == nil에 넣어서 생김;;
```
Error: endorsement failure during invoke. response: status:500 message:"failed to execute transaction 163d61107ba8285a970dd305718df61d47744b5ed95c058f5f4d88051dddff05: error sending: timeout expired while executing transaction" 
```
채널을 나눠서 운영하는건 옛날방식, 체인코드를 여러개로 나누는것보다 Private Data를 사용하는게 좋다  
Fabric v 1.4 부터는 Kafka 대신 Raft 알고리즘 사용  

3f+1 -> PBFT  
L(leader)+f(follow) -> RAFT  
RAFT는 여러개의 오더러가 모여서 consensus 함 (5개), kafka (4개)  
위 두개 모두 블록의 내용이 같다는걸 증명하기 위함  

contractapi 라이브러리와 shim 라이브러리는 뭐가 다른가
