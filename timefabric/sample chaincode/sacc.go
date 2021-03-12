package main


import (
	//"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	//pt "github.com/golang/protobuf/ptypes"
)


type SmartContract struct {


}

type Item struct {
	ItemCode int
	AvQty int
	ExpiryTime time.Time

}



func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)

}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fn, args := stub.GetFunctionAndParameters()
	var res string
	var err error
	if (fn=="createItem"){
		res,err = createItem(stub, args)
	} else if (fn=="executePurchase"){
		res,err = executePurchase(stub,args)
	} else if (fn=="getItem") {
		res,err = getItem(stub, args)
	} else if (fn=="checkTime") {
		res, err = checkTime(stub, args)
	} else {
		res, err = createTime(stub, args)
	}

	if err!=nil{
		return (shim.Error(err.Error()))
	}

	return shim.Success([]byte(res))
}


func createItem (stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args)!=3{
		return "", fmt.Errorf("Incorect arguments passed for %s",args[0])
	}	

	const customform = "02-Jan-2006, 3:04:05pm MST"
	code,_:= strconv.Atoi(args[0])
	qty,_:= strconv.Atoi(args[1])
	expTime,_:= time.Parse(customform, args[2])
	newItem := Item {ItemCode: code,
	AvQty: qty,
	ExpiryTime: expTime}
	storeItem, err1:= json.Marshal(newItem)
	if err1!=nil{
		return "", fmt.Errorf("error in parsing the value for %s with error %s",args[0],err1)

	}
	
	valid, err := ValidateTxTime(stub)
	if !(valid){
		return "", fmt.Errorf("Failed to create item for %s due to time diff",err)
		
	}




	err5:= stub.PutState(args[0],storeItem)
	if err5!=nil{
		return "",fmt.Errorf("Failed to create item for %s",args[0])
	}
	return args[0], nil


}

func getItem (stub shim.ChaincodeStubInterface, args []string) (string, error) {
	value, err := stub.GetState(args[0])
	if err != nil {
		return "",fmt.Errorf("Invalid item code: %s", args[0])
	}

	return string(value), nil

}


func executePurchase (stub shim.ChaincodeStubInterface, args []string) (string, error) {

	//parse args to query asset
	value,err := stub.GetState(args[0])
	if err!=nil{
		return "", fmt.Errorf("Could not retrieve item with code %s",args[0])
	}
	item := Item{}
	err1:= json.Unmarshal(value,&item)
	if err1!=nil{
		return "",fmt.Errorf("Error in unmarshaling item %s",args[0])
	}
	t2:= item.ExpiryTime
	t1,err2 := stub.GetTimenow()
	if err2!=nil{
		return "", fmt.Errorf("Error in retrieving tx timestamp from header")
	}
	var x time.Time
	err3 := x.UnmarshalBinary(t1)
	if err3!=nil{
		return "", fmt.Errorf("unmarshalling error %s", args[0])
	}
	
	valid, err19 := ValidateTxTime(stub)
	if !(valid){
		return "", fmt.Errorf("Failed to create item for %s due to time diff",err19)
		
	}
	
	if (t2.Unix()>x.Unix()){
		item.AvQty = item.AvQty-1
		val,err2:=json.Marshal(item)
		if err2!=nil{
			return "",fmt.Errorf("Error in re-marshaling item %s",args[0])
		}
		err3:= stub.PutState(args[0],val)
		if err3!=nil{
			return "",fmt.Errorf("Error in update item %s quantity",args[0])
		}

	} else{
		return "",fmt.Errorf("Failed to execute transaction for %s",args[0])
	}

	return args[0], nil
}


func checkTime (stub shim.ChaincodeStubInterface, args []string) (string, error) {

	value, err := stub.GetTimenow()
	if err!=nil{
		return "", fmt.Errorf("api not working %s", args[0])
	}
	var x time.Time
	err1 := x.UnmarshalBinary(value)
	if err1!=nil{
		return "", fmt.Errorf("unmarshalling error %s", args[0])
	}
	return x.String(),nil
}
func createTime (stub shim.ChaincodeStubInterface, args []string) (string, error) {

	z := time.Now()	
	x, err := stub.GetTimenow()
	if err!=nil{
		return "",fmt.Errorf("Error in fetching time %s", args[0])
	}

	var y time.Time
	err1:= y.UnmarshalBinary(x)
	if err1 !=nil{
		return "",fmt.Errorf("Unmarshalling error %s", args[0])
	}

	diff := (z.UnixNano()/int64(time.Millisecond)) - (y.UnixNano()/int64(time.Millisecond))
	res:= strconv.FormatInt(diff,10)
	err3:= stub.PutState(args[0],[]byte(res))
	if err3 !=nil{
		return "",fmt.Errorf("Error in update item %s quantity",args[0])
	}

	return args[0], nil
}


func ValidateTxTime (stub shim.ChaincodeStubInterface) (bool, string){

	var diff int64	
	value, err2 := stub.GetTimenow()
	if err2!=nil{
		return false,"Error fetching time"
	}
	var x time.Time
	err3 := x.UnmarshalBinary(value)
	if err3!=nil{
		return false, "Unmarshall Error"
	}
	value1 := x.Unix()
	
	t1,err4 := stub.GetTxTimestamp()
	if err4!=nil{
		return false,"Error in getting tx time"
	}
	value2 := t1.Seconds
	if (value2>value1){
		diff = value2-value1
	} else {
		diff = value1-value2
	}
	returnValue := strconv.FormatInt(diff,10)

	if (diff>10) {
		return false,returnValue 
	}

	return true,""
	

}





	


func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
