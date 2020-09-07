package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Wallet struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

//조직 2개 다 PDC 접근 허용
type PersonalInfo struct {
	ObjectType string `json:"docType"`
	Name       string `json:"name"`
	Age        string `json:"age"`
}

//sales 조직만 접근 허용
type Price struct {
	ObjectType string `json:"docType"`
	Name       string `json:"name"`
	Price      string `json:"price"`
}

// func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
// 	return shim.Success(nil)
// }

// func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
// 	function, args := APIstub.GetFunctionAndParameters()

// 	if function == "initWallet" {
// 		return s.initWallet(ctx, args)
// 	} else if function == "chargeMoney" {
// 		return s.chargeMoney(APIstub, args)
// 	} else if function == "getWallet" {
// 		return s.getWallet(APIstub, args)
// 	} else if function == "transferMoney" {
// 		return s.transferMoney(APIstub, args)
// 	}
// 	fmt.Println("Please check your function : " + function)
// 	return shim.Error("Unknown function")
// }

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting private chaincode: %s", err.Error())
	}

}

func (s *SmartContract) initWallet(ctx contractapi.TransactionContextInterface, args []string) error {

	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return fmt.Errorf("Error getting transient: " + err.Error())
	}

	transientPIJSON, ok := transMap["PersonalInfo"]
	if !ok {
		return fmt.Errorf("marble not found in the transient map")
	}

	type infoTransientInput struct {
		Name  string `json:"name"`
		Age   string `json:"age"`
		Price string `json:"price"`
	}

	infoInput := infoTransientInput{
		Name:  args[0],
		Age:   args[1],
		Price: args[2],
	}

	err = json.Unmarshal(transientPIJSON, &infoInput)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}

	if len(infoInput.Name) == 0 {
		return fmt.Errorf("name field must be a non-empty string")
	}
	if len(infoInput.Price) == 0 {
		return fmt.Errorf("Price field must be a non-empty string")
	}
	if len(infoInput.Age) == 0 {
		return fmt.Errorf("Age field must be a non-empty string")
	}

	PIAsJSONBytes, err := ctx.GetStub().GetPrivateData("collectionPersonalInfo", infoInput.Name)
	if err != nil {
		return fmt.Errorf("Failed to get Personal Info " + err.Error())
	} else if PIAsJSONBytes != nil {
		fmt.Println("This Info is already exists " + infoInput.Name)
		return fmt.Errorf("This Info is already exists " + infoInput.Name)
	}

	// 개인정보 객체 생성
	personalInfo := &PersonalInfo{
		ObjectType: "Info",
		Name:       infoInput.Name,
		Age:        infoInput.Age,
	}

	personalInfoAsJSONBytes, err := json.Marshal(personalInfo)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	//개인정보 state에 저장
	err = ctx.GetStub().PutPrivateData("collectionPersonalInfo", infoInput.Name, personalInfoAsJSONBytes)
	if err != nil {
		return fmt.Errorf("failed to put Personal Info: %s ", err.Error())
	}

	// 세일즈조직에만 보여지는 Price 객체 생성
	price := &Price{
		ObjectType: "Price",
		Name:       infoInput.Name,
		Price:      infoInput.Price,
	}

	priceAsJSONBytes, err := json.Marshal(price)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	err = ctx.GetStub().PutPrivateData("collectionPrice", infoInput.Name, priceAsJSONBytes)
	if err != nil {
		return fmt.Errorf("failed to put Price: %s", err.Error())
	}

	// indexName := "name"
	// nameIndexKey, err := ctx.GetStub().CreateCompositeKey(indexName,[args[0]])\
	// if err != nil {
	// 	return err
	// }

	// err = ctx.GetStub().DelPrivateData("collectionPersonalInfo", nameIndexKey)
	// if err != nil {
	// 	return fmt.Errorf("Failed to delete personal Info:" + err.Error())
	// }

	// // Finally, delete private details of marble
	// err = ctx.GetStub().DelPrivateData("collectionPrice", marbleDeleteInput.Name)
	// if err != nil {
	// 		return err
	// }

	return nil

}

// func (s *SmartContract) chargeMoney(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 2")
// 	}

// 	WalletAsBytes, err := APIstub.GetState(args[0])
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	wallet := Wallet{}
// 	json.Unmarshal(WalletAsBytes, &wallet)

// 	var walletToken, _ = strconv.Atoi(string(wallet.Token))
// 	var amount, _ = strconv.Atoi(string(args[1]))

// 	wallet.Token = strconv.Itoa(walletToken + amount)

// 	updatedWalletAsBytes, _ := json.Marshal(wallet)

// 	APIstub.PutState(wallet.ID, updatedWalletAsBytes)

// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")

// 	buffer.WriteString("{\"Charging Token\":")
// 	buffer.WriteString("\"")
// 	buffer.WriteString(wallet.Token)
// 	buffer.WriteString("\"")

// 	buffer.WriteString("}")
// 	buffer.WriteString("]\n")

// 	return shim.Success(buffer.Bytes())
// }

// func (s *SmartContract) getWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

// 	walletAsBytes, err := APIstub.GetState(args[0])
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	wallet := Wallet{}
// 	json.Unmarshal(walletAsBytes, &wallet)

// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")
// 	bArrayMemberAlreadyWritten := false

// 	if bArrayMemberAlreadyWritten == true {
// 		buffer.WriteString(",")
// 	}

// 	buffer.WriteString("{\"ID\":")
// 	buffer.WriteString("\"")
// 	buffer.WriteString(wallet.ID)
// 	buffer.WriteString("\"")

// 	buffer.WriteString(", \"Token\":")
// 	buffer.WriteString("\"")
// 	buffer.WriteString(wallet.Token)
// 	buffer.WriteString("\"")

// 	buffer.WriteString("}")
// 	bArrayMemberAlreadyWritten = true
// 	buffer.WriteString("]\n")

// 	return shim.Success(buffer.Bytes())

// }

// func (s *SmartContract) transferMoney(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	var err error
// 	if len(args) != 3 {
// 		return shim.Error("Incorrect number of arguments. Expecting 3")
// 	}

// 	fmt.Println("Sender : " + args[0] + ", Receiver : " + args[1] + "amount : " + args[2])

// 	SenderAsBytes, err := APIstub.GetState(args[0])
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	sender := Wallet{}
// 	json.Unmarshal(SenderAsBytes, &sender)

// 	ReceiverAsBytes, err := APIstub.GetState(args[1])
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	receiver := Wallet{}
// 	json.Unmarshal(ReceiverAsBytes, &receiver)

// 	var senderToken, _ = strconv.Atoi(string(sender.Token))
// 	var receiverToken, _ = strconv.Atoi(string(receiver.Token))
// 	var amount, _ = strconv.Atoi(string(args[2]))

// 	sender.Token = strconv.Itoa(senderToken - amount)
// 	receiver.Token = strconv.Itoa(receiverToken + amount)

// 	updatedSenderAsBytes, _ := json.Marshal(sender)
// 	updatedReceiverAsBytes, _ := json.Marshal(receiver)

// 	APIstub.PutState(args[0], updatedSenderAsBytes)
// 	APIstub.PutState(args[1], updatedReceiverAsBytes)

// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")

// 	buffer.WriteString("{\"Receiver Token\":")
// 	buffer.WriteString("\"")
// 	buffer.WriteString(receiver.Token)
// 	buffer.WriteString("\"")

// 	buffer.WriteString(", \"Sender Token\":")
// 	buffer.WriteString("\"")
// 	buffer.WriteString(sender.Token)
// 	buffer.WriteString("\"")

// 	buffer.WriteString("}")
// 	buffer.WriteString("]\n")

// 	return shim.Success(buffer.Bytes())
// }
