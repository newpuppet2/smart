package main

/**
 * Shows how to use the history
 **/

import (
	// For printing messages on console
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// // peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"

	// JSON Encoding
	"encoding/json"

	// KV Interface
	//"github.com/hyperledger/fabric/protos/ledger/queryresult"

	// "strconv"
)

type LC struct {
}



// LetterCredit Represents our car asset
type LetterCredit struct {
	DocType			                   string  `json:"docType"`
	Date			                   string  `json:"date"`
	ImporterName 		               string   `json:"importerName"`
	ExporterName		               string   `json:"exporterName"`
	ImporterBankName                   string  `json:"importerBankName"`
	ExporterBankName			       string  `json:"exporterBankName"`
	ProductOrderId			           string    `json:"productOrderId"`
	ProductOrderDetails		           string  `json:"productOrderDetails"`	
	PaymentAmount                     int    `json:"paymentAmount"`
	State                              string  `json:"state"`
	Pendingstate                       string  `json:"pendingstate"`
}

const	DocType	= "LC"

func (loc *LC) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed in loc")

	// Setup the sample data
	loc.SetupSampleData(stub)

	// Return success
	return shim.Success(nil)
}

func (loc *LC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the function name and parameters
	funcName, args := stub.GetFunctionAndParameters()

	if funcName == "ApproveLC" {
		// Returns the vehicle's current state
		return loc.createlc(stub, args)

	} else if funcName == "GetLC" {
		// Invoke this function to transfer ownership of vehicle
		return loc.getlc(stub, args)

	} 

	// This is not good
	return shim.Error(("Bad Function Name = !!!"))
}

func (tradefinance *LC) ApproveLC(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	firstarg := "qwert"
        bytes, _ := stub.GetState(firstarg)
	if bytes == nil {
		return shim.Error("Provided ID not found!!!")
	}

	var lc1  LetterCredit
	_ = json.Unmarshal(bytes, &lc1)

	app := string(args[0])
	
    if app == "reject" {
		lc1.State = "rejected"
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(firstarg, jsonletter)
		return shim.Success([]byte("Transaction rejected"))
	}
 
	if app == "accept" && lc1.Pendingstate == "importer" {
		lc1.State = "pending"
		lc1.Pendingstate = "importerbank"
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(firstarg, jsonletter)	
		return shim.Success([]byte("Transaction approved"))
	}

	if app == "accept" && lc1.Pendingstate == "importerbank" {
		lc1.State = "pending"
		lc1.Pendingstate = "exporterbank"
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(firstarg, jsonletter)
		return shim.Success([]byte("Transaction approved"))
	}

	if app == "accept" && lc1.Pendingstate == "exporterbank" {
		lc1.State = "pending"
		lc1.Pendingstate = "exporter"
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(firstarg, jsonletter)
		return shim.Success([]byte("Transaction approved"))
	}

	if app == "accept" && lc1.Pendingstate == "exporter" {
		lc1.State = "pending"
		lc1.Pendingstate = "exportcustoms"
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(args[0], jsonletter)
		return shim.Success([]byte("Transaction approved"))
	}

	if app == "accept" && lc1.Pendingstate == "exportcustoms" {
		lc1.State = "pending"
		lc1.Pendingstate = "importcustoms"
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(args[0], jsonletter)
		return shim.Success([]byte("Transaction approved"))
	}

	if app == "accept" && lc1.Pendingstate == "importcustoms" {
		lc1.State = "complete"
		lc1.Pendingstate = "importcustoms"
        jsonletter, _ := json.Marshal(lc1)
		stub.PutState(args[0], jsonletter)
		return shim.Success([]byte("Transaction approved"))
	}
return shim.Success([]byte("success"))

}

func (loc *LC) Getlc(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	firstarg := "qwert"
	bytes, _ := stub.GetState(firstarg)

	if bytes == nil {
		return shim.Error("Provided ID not found!!!")
	}

	var lc1  LetterCredit
	_ = json.Unmarshal(bytes, &lc1)

	app := string(args[0])
 

       if app == "importer" && lc1.Pendingstate == "importer" {
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(firstarg, jsonletter)
        return shim.Success(jsonletter)
	}

	if app == "importerbank" && lc1.Pendingstate == "importerbank" {
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(firstarg, jsonletter)
        return shim.Success(jsonletter)
	}

	if app == "exporterbank" && lc1.Pendingstate == "exporterbank" {
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(firstarg, jsonletter)
        return shim.Success(jsonletter)
	}

	if app == "exporter" && lc1.Pendingstate == "exporter" {
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(firstarg, jsonletter)
        return shim.Success(jsonletter)
	}

	if app == "exportcustoms" && lc1.Pendingstate == "exportcustoms" {
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(firstarg, jsonletter)
        return shim.Success(jsonletter)
	}

	if app == "importcustoms" && lc1.Pendingstate == "importcustoms" {
		jsonletter, _ := json.Marshal(lc1)
		stub.PutState(firstarg, jsonletter)
        return shim.Success(jsonletter)
	}
return shim.Success([]byte("success"))
	


}


func (loc *LC) SetupSampleData(stub shim.ChaincodeStubInterface) {
	
	// This the car data for testing
	AddData(stub, "10-11-2019","Ays","Tencent","HDFC","Kotak","qwert","Batteries",10000,"new","importer")
	//AddData(stub, "12-11-2019","Ays","Tencent","HDFC","Kotak","qwert1","Hardware",10000,"new",100)
	//AddData(stub, "14-11-2019","Ays","Tencent","HDFC","Kotak","qwert2","Shutters",10000,"new",100)
	//AddData(stub, "16-11-2019","Ays","Tencent","HDFC","Kotak","qwert4","Lens",10000,"new",100)
	
	fmt.Println("Initialized with the sample data!!")
}

//AddData adds a car asset to the chaincode asset database
//Structure is created and initialized then it is marshalled to JSON for storage using PutState
func AddData(stub shim.ChaincodeStubInterface,date string, importer string, exporter string, importerbank string, exporterbank string, key string, productdes string,payment int, status string, pendingstate string) {
	letter := LetterCredit{DocType: DocType, Date: date, ImporterName: importer, ExporterName: exporter, ImporterBankName: importerbank, ExporterBankName: exporterbank, ProductOrderId: key, ProductOrderDetails: productdes, PaymentAmount: payment, State: status,Pendingstate: pendingstate}
	jsonletter, _ = json.Marshal(letter)
	// Key = VIN#, Value = Car's JSON representation
	stub.PutState(key, jsonletter)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. Letter of credit\n")
	err := shim.Start(new(LC))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}




