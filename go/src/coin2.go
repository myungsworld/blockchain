package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct{}

type Wallet struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "initWallet" {
		return s.initWallet(stub)
	} else if function == "sendMoney" {
		return s.sendMoney(stub, args)
	}
	return shim.Error("unknown function")
}

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

func (s *SmartContract) initWallet(stub shim.ChaincodeStubInterface) pb.Response {

	//user := Wallet{}
	// userAsJSONBytes, _ := json.Marshal(user)
	// err := stub.PutState(user.ID, userAsJSONBytes)
	// if err == nil {
	// 	return shim.Error("Failed to initializing :" + user.ID)
	// }
	user1 := Wallet{ID: "myung", Token: "100"}
	user1AsJSONBytes, _ := json.Marshal(user1)
	err := stub.PutState(user1.ID, user1AsJSONBytes)
	if err != nil {
		return shim.Error("Failed to initializing :" + user1.ID)
	}

	user2 := Wallet{ID: "min", Token: "200"}
	user2AsJSONBytes, _ := json.Marshal(user2)
	err = stub.PutState(user1.ID, user2AsJSONBytes)
	if err != nil {
		return shim.Error("Failed to initializing :" + user2.ID)
	}

	user3 := Wallet{ID: "musk", Token: "20000000"}
	user3AsJSONBytes, _ := json.Marshal(user3)
	err = stub.PutState(user3.ID, user3AsJSONBytes)
	if err != nil {
		return shim.Error("Failed to initializing :" + user3.ID)
	}

	return shim.Success(nil)
}

func (s *SmartContract) sendMoney(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var tokenFromSeller int
	var tokenFromCustomer int
	var price int

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	//판매자 정보 couchDB에서 가져옴
	sellerAsJSONBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if sellerAsJSONBytes == nil {
		return shim.Error("Seller not found")
	}
	seller := Wallet{}
	json.Unmarshal(sellerAsJSONBytes, &seller)
	tokenFromSeller, _ = strconv.Atoi(seller.Token)

	//구매자 정보 couchDB에서 가져옴
	customerAsJSONBytes, err := stub.GetState(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
	if customerAsJSONBytes == nil {
		return shim.Error("Cumtomer not found")
	}
	customer := Wallet{}
	json.Unmarshal(customerAsJSONBytes, &customer)
	tokenFromCustomer, _ = strconv.Atoi(customer.Token)

	//가격
	price, _ = strconv.Atoi(args[2])

	customer.Token = strconv.Itoa(tokenFromCustomer - price)
	seller.Token = strconv.Itoa(tokenFromSeller + price)

	updatedSellerAsJSONBytes, _ := json.Marshal(seller)
	stub.PutState(args[0], updatedSellerAsJSONBytes)
	updatedCustomerAsJSONBytes, _ := json.Marshal(customer)
	stub.PutState(args[1], updatedCustomerAsJSONBytes)

	var buffer bytes.Buffer
	buffer.WriteString(customer.ID)
	buffer.WriteString(" transferred ")
	buffer.WriteString(args[2])
	buffer.WriteString(" to ")
	buffer.WriteString(seller.ID)

	return shim.Success(buffer.Bytes())
}
