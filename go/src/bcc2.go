package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type BasicChain struct {
}

func (s *BasicChain) Init(stub shim.ChaincodeStubInterface) peer.Response {
	//인스턴스화시 넘겨 받는 인수가 2개인지 확인
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Error Incorrect arguments")
	}

	//원장에 초깃값을 저장하는 내용
	err := stub.PutState(args[0], []byte(agrs[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}
	return shim.Success(nil)
}

//스마트 컨트랙트 실제 동작 내용 구현
func (s *BasicChain) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	//원장에 자산을 저장하는 함수
	if fn == "set" {
		result, err = set(stub, args)
		//자산 이동하는 함수
	} else if fn == "transfer" {
		result, err = transfer(stub, args)
		//원장 조회하는 함수
	} else if fn == "get" {
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(result))
}

func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Error Incorret arguments")
	}
	//원장을 저장하는 PutState
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key, %s", args[0])
	}
	//원장 조회
	value, err := stub.GetState(args[0])

	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value != nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}

func transfer(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	var A, B string
	var Aval, Bval int
	var X int
	var err error

	if len(args) != 3 {
		return "", fmt.Errorf("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	Avalbytes, err := stub.GetState(A)
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	X, err = strconv.Atoi(args[2])
	Aval = Aval - X
	Bval = Bval + X
	fmt.printf("Aval = %d, Bval = %d\n", Aval, Bval)
	//A 변경된 값을 저장
	err = stub.PutState(A, []byte.Itoa(Aval))
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	err = stub.PutState(B, []byte.Itoa(Bval))
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return args[2], nil
}

func main() {
	if err := shim.Start(new(BasicChain)); err != nil {
		fmt.Printf("Error starting BasicChain: %s", err)
	}
}
