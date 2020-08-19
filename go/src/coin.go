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
	Name  string `json:"name"`
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
	} else if function == "getWallet" {
		return s.getWallet(stub)
	} else if function == "setFood" {
		return s.setFood(stub)
	}
	fmt.Println("Please check your function: " + function)
	return shim.Error("Unknown function")
}
func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

func (s *SmartContract) initWallet(stub shim.ChaincodeStubInterface) pb.Response {

	//이건 데이터베이스에서 가져올꺼임 이건 예제라 이렇게 적은거임
	seller := Wallet{Name: "Song", ID: "thdehdaud", Token: "100"}
	customer := Wallet{Name: "Youmin", ID: "qkrdbals", Toekn: "200"}
	customer := Wallet{Name: "Sanghyun", ID: "qustkdgus", Toekn: "200"}
	customer := Wallet{Name: "Gangju", ID: "gjrkdwn", Toekn: "200"}
	customer := Wallet{Name: "Sanghoon", ID: "dltkdgns", Toekn: "200"}

	//카우치디비에 저장하기 위해 마샬링을 하는거임
	SellerAsJSONBytes, _ := json.Marshal(seller)
	err := stub.Putstate(seller.ID, SellerAsJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset" + seller.Name)
	}

	CustomerAsJSONBytes, _ := json.Marshal(customer)
	err := stub.Putstate(customer.ID, CustomerAsJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset" + customer.Name)
	}

	return shim.Success(nil)

}
func (s *SmartContract) getWallet(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	walletAsBytes, err := stub.GetState(args[0])
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
	buffer.WriteString("{\"Name\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.Name)
	buffer.WriteString("\"")

	buffer.WriteString(", \"ID\":")
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

type Food struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	WalletID string `json:"walletid"`
	Count    string `json:"count"`
}

type FoodKey struct {
	Key string
	Idx int
}

func generateKey(stub shim.ChaincodeStubInterface, key string) []byte {

	var isFirst bool = false

	foodkeyAsJSONBytes, err := stub.GetState(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	foodkey := FoodKey{}
	json.Unmarshal(foodkeyAsJSONBytes, &foodKey)
	var tempIdx string
	tempIdx = strconv.Itoa(foodkey.Idx)
	fmt.Println(foodkey)
	fmt.Println("Key is " + strconv.Itoa(len(foodkey.Key)))
	if len(foodkey.Key) == 0 || foodkey.Key == "" {
		isFirst = true
		foodkey.Key = "FS"
	}
	if !isFirst {
		foodkey.Idx = foodkey.Idx + 1
	}

	fmt.Println("Last FoodKey is " + foodkey.Key + " : " + tempIdx)

	returnValueBytes, _ := json.Marshal(foodkey)

	return returnValueBytes
}

func (s *SmartContract) setWallet(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(agrs) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	var wallet = Wallet{Name: args[0], ID: args[1], Token: args[2]}

	WalletAsJSONBytes, _ := json.Marshal(wallet)
	err := stub.Putstate(wallet.ID, WalletAsJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + wallet.Name)
	}
	return shim.Success(nil)
}

func (s *SmartContract) setFood(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var foodkey = FoodKey{}
	json.Unmarshal(generateKey(stub, "latestKey"), &foodkey)
	keyidx := strconv.Itoa(foodkey.Idx)
	fmt.Println("Key : " + foodkey.Key + ", Idx : " + keyidx)

	var food = Food{Name: args[0], Price: args[1], WalletID: args[2], Count: "0"}
	foodAsJSONBytes, _ := json.Marshal(food)

	var keyString = foodkey.Key + keyidx
	fmt.Println("foodkey is " + keyString)

	err := stub.Putstate(keyString, foodAsJSONBytes)
	if err != nil {
		return shim.Error(fmt.Sprint)
	}

	foodkeyAsJSONBytes, _ := json.Marshal(musickey)
	stub.Putstate("latestKey", foodkeyAsJSONBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getAllfood(stub shim.ChaincodeStubInterface) pb.Response {

	foodkeyAsJSONBytes, _ := stub.GetState("latestkey")
	foodkey := FoodKey{}
	json.Unmarshal(foodkeyAsJSONBytes, &foodkey)
	idxStr := strconv.Itoa(foodkey.Idx + 1)

	var startKey = "MS0"
	var endKey = musickey.key + idxStr
	fmt.Println(startKey)
	fmt.Println(endKey)

	resultsIter, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIter.HasNext() {
		queryResponse, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")

		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]\n")
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) buyFood(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var tokenFromKey, tokenToKey int
	var foodprice int
	var foodcount int
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	foodAsJSONBytes, err := stub.GetState(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}

	food := Food{}
	json.Unmarshal(foodAsJSONBytes, &food)
	foodprice, _ = strconv.Atoi(food.Price)
	foodcount, _ = strconv.Atoi(food.Count)

	SellerAsJSONBytes, _ := stub.GetState(food.WalletID)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if SellerAsJSONBytes == nil {
		return shim.Error("Entity not found")
	}
	seller := Wallet{}
	json.Unmarshal(SellerAsJSONBytes, &seller)
	tokenToKey, _ := strconv.Atoi(seller.Token)

	CustomerAsJSONBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if CustomerAsJSONBytes == nil {
		return shim.Error("Entity is not found")
	}

	customer := Wallet{}
	json.Unmarshal(CustomerAsJSONBytes, &customer)
	tokenFromKey, _ = strconv.Atoi(string(customer.Token))

	customer.Token = strconv.Itoa(tokenFromKey - foodprice)
	seller.Token = strconv.Itoa(tokenToKey + foodprice)
	food.Count = strconv.Itoa(foodcount + 1)
	updatedCustomerAsJSONBytes, _ := json.Marshal(customer)
	updatedSellerAsJSONBytes, _ := json.Marshal(seller)
	updatedFoodAsJSONBytes, _ := json.Marshal(food)
	stub.Putstate(args[0], updatedCustomerAsJSONBytes)
	stub.Putstate(food.WalletID, updatedSellerAsJSONBytes)
	stub.Putstate(args[1], updatedFoodAsJSONBytes)

	var buffer bytes.Buffer
	buffer.WriteString("[")

	buffer.WriteString("{\"Customer Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(customer.Token)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Seller Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(seller.Token)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) getFood(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	foodAsJSONBytes, err := stub.GetState(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}

	food := Food{}
	json.Unmarshal(foodAsJSONBytes, &food)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}
	buffer.WriteString("{\"Name\":")
	buffer.WriteString("\"")
	buffer.WriteString(food.Name)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Price\":")
	buffer.WriteString("\"")
	buffer.WriteString(food.Price)
	buffer.WriteString("\"")

	buffer.WriteString(", \"WalletID\":")
	buffer.WriteString("\"")
	buffer.WriteString(food.WalletID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"count\":")
	buffer.WriteString("\"")
	buffer.WriteString(food.Count)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())

}

func (s *SmartContract) changeFoodPrice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	foodAsJSONBytes, _ := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Could not locate music")
	}
	food := Food{}
	json.Unmarshal(foodAsJSONBytes, &food)

	food.Price = args[1]

	foodAsJSONBytes, _ = json.Marshal(food)
	err2 := stub.Putstate(args[0], foodAsJSONBytes)
	if err2 != nil {
		return shim.Error(fmt.Sprintf("Failed to change food price: %s", args[0]))
	}
	return shim.Success(nil)
}

func (s *SmartContract) deleteFood(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	//Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}
