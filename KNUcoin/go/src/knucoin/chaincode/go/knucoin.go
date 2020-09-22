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

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

type Wallet struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()

	if function == "initWallet" {
		return s.initWallet(APIstub, args)
	} else if function == "chargeMoney" {
		return s.chargeMoney(APIstub, args)
	} else if function == "getWallet" {
		return s.getWallet(APIstub, args)
	} else if function == "transferMoney" {
		return s.transferMoney(APIstub, args)
	}
	fmt.Println("Please check your function : " + function)
	return shim.Error("Unknown function")
}

func (s *SmartContract) initWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	//Declare wallet
	wallet := Wallet{ID: args[0], Token: "0"}
	// Convert wallet to []byte
	WalletasJSONBytes, _ := json.Marshal(wallet)

	err := APIstub.PutState(wallet.ID, WalletasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + wallet.ID)
	}

	return shim.Success(nil)
}

func (s *SmartContract) chargeMoney(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	WalletAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	wallet := Wallet{}
	json.Unmarshal(WalletAsBytes, &wallet)

	var walletToken, _ = strconv.Atoi(string(wallet.Token))
	var amount, _ = strconv.Atoi(string(args[1]))

	wallet.Token = strconv.Itoa(walletToken + amount)

	updatedWalletAsBytes, _ := json.Marshal(wallet)

	APIstub.PutState(wallet.ID, updatedWalletAsBytes)

	var buffer bytes.Buffer
	buffer.WriteString("[")

	buffer.WriteString("{\"Charging Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.Token)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) getWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	walletAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}

	wallet := Wallet{}
	json.Unmarshal(walletAsBytes, &wallet)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}

	buffer.WriteString("{\"ID\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.ID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.Token)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())

}

func (s *SmartContract) transferMoney(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	var senderToken, receiverToken int
	var amount int
	var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	fmt.Println("Sender : " + args[0] + ", Receiver : " + args[1] + "amount : " + args[2])

	receiverAsBytes, err := APIstub.GetState(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}

	if receiverAsBytes == nil {
		return shim.Error("Entity not found")
	}
	receiver := Wallet{}
	json.Unmarshal(receiverAsBytes, &receiver)
	receiverToken, _ = strconv.Atoi(string(receiver.Token))

	senderAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if senderAsBytes == nil {
		return shim.Error("Entity not found")
	}

	sender := Wallet{}
	json.Unmarshal(senderAsBytes, &sender)
	senderToken, _ = strconv.Atoi(sender.Token)

	amount, _ = strconv.Atoi(args[2])
	if amount < 0 {
		return shim.Error("you can't transfer less than 0")
	}
	sender.Token = strconv.Itoa(senderToken - amount)
	if amount > senderToken {
		return shim.Error(args[0] + "doesn't have enough money to send")
	}
	receiver.Token = strconv.Itoa(receiverToken + amount)

	updatedSenderAsBytes, _ := json.Marshal(sender)
	updatedReceiverAsBytes, _ := json.Marshal(receiver)

	APIstub.PutState(args[0], updatedSenderAsBytes)
	APIstub.PutState(args[1], updatedReceiverAsBytes)

	var buffer bytes.Buffer
	buffer.WriteString("[")

	buffer.WriteString("{\"Receiver Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(receiver.Token)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Sender Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(sender.Token)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}
