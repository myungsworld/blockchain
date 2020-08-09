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
  - __조직__ : 조직별 피어노드 관리  


