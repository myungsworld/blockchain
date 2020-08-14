## 마샬링
논리적 구조를 로우 바이트로 변경하는 것을 Marshaling or Encoding 이라고 표현한다.  
구조체 같은 go value를 바이트 슬라이스로 변경하는 것이다.  

__json.Marshal__ 함수가 이 역할을 한다.  
__json,MarshalIndent__ 가독성을 높이고 싶다면 사용.  

ex. json.MarshalIndent([struct],""," ")

## 언마샬링
반대로 바이트 슬라이스나 문자열을 논리적 자료 구조로 변경하는 것을 언마샬링이라고 한다.
__json.Unmarshal__ 
```
var s = `{"name":"gopher","age":7}`
var u User
json.Unmarshal([]byte(s), &u)
fmt.Printf("%+v\n", u) // {Name:gopher Age:7}
```
JSON 문자열을 바이트 슬라이스 형태로 넘겨주고 User 타입 변수 u의 포인터를 전달한다,. 함수가 실행되면 문자열이 파싱되어 User 값이 생성된다.

## 인코더
많은 데이터를 처리할 때 스트림을 사용한다. __json.Encoder__ 타입은 스트림 기반으로 JSON 문자열을 만든다.  

마샬링과 인코딩의 차이  
바이트 슬라이스나 문자열을 사용하려면 Marshal/Unmarshal 함수가 적합하다. 만약 표준 입출력이나 파일 같은 Reader/Writer 인터페이스를 사용하여 스트림 기반으로 동작하려면 Encoder/Decoder를 사용한다.  

처리 속도는 스트림 방식이 더 낫다. 데이터 크기가 작다면 성능차이를 체감할 수 없지만 비교적 큰 데이터를 다룬다면 스트림 기반의 Encoder/Decoder가 거의 50% 정도 더 빠른 성능을 낸다.  
