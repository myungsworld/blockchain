package main

import (
	"fmt"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//stub = 속이 빈 함수(가짜함수)
// 1 구현되지 않은 함수거나 라이브러리에서 제공하는 함수
// 2 함수가 반환하는 값을 임의로 생성
// 3 복잡한 논리 흐름을 가지는 경우 테스트를 단순화 할 목적으로 사용

type SimpleChaincode struct{}

/*체인코드가 처음 실행될때 실행되는 함수 (체인코드 인스턴스화 시킴)
SimpleChaincode struct를 사용하는 Init 메서드*/
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("ex init")
	 , agrs := stub.GetFunctionAndParameters()
	var A, b string //keys
	var Aval, Bval int // key values
	var err error
	
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments, Expecting 4")
	}
	
	A = args[0]
	Aval , err = strconv.Atoi(args[1]) //문자열을 정수로 변경
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	
	B = args[2]
	Bval , err = strconv.Atoi(argsp[4]) 
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}

	fmt.Printf("Aval: = %d, Bval = %d", Aval , Bval )

	//putState를 쓰면 world state를 자동으로 변경해주고 아래 query에서 getState로 조회 가능
	//putstate에선 정수를 bytes로 변환해서 써줌
	err = stub.Putstate(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.Putstate(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//변경
func (t *SimpleChaincode) Invoke(stub ChaincodeStubInterface) pb.response {}

//조회
func (t *SimpleChaincode) query(stub ChaincodeStubInterface, args []string) pb.response {
	var A string //Entities
	var err error

	if len(args) != 1 {
		return shim.error("Incorrect")
	}
	//Getstate는 가장 나중,최근의 키값을 가져온다
	Avalbytes, err := stub.Getstate(A)

}

func main() {
	//Init과 Inoke method implement
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
