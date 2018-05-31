// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Tuna structure, with 4 properties.  
Structure tags are used by encoding/json library
*/
// type Tuna struct {
// 	Vessel string `json:"vessel"`
// 	Timestamp string `json:"timestamp"`
// 	Location  string `json:"location"`
// 	Holder  string `json:"holder"`
//}

type Piece struct {
    RefPiece string `json:"refPiece"`
	NomPiece string `json:"nomPiece"`
	NomFabricant  string `json:"nomFabricant"`
    NumeroLot string `json:"numeroLot"`
	DateExp string `json:"dateExp"`
	LocalisationVille string `json:"localisationVille"`
	LocalisationEtablissement string `json:"localisationEtablissement"`
	DateCreation string `json:"dateCreation"`
	DateReception string `json:"dateReception"`
	VenteClient string `json:"venteClient"`
}


/*
 * The Init method *
 called when the Smart Contract "tuna-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "tuna-chaincode"
 The app also specifies the specific smart contract function to call with args
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	fmt.Println(function)

	// Route to the appropriate handler function to interact with the ledger
	if function == "queryPiece" {
		return s.queryPiece(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordPiece" {
		return s.recordPiece(APIstub, args)
	} else if function == "queryAllPiece" {
		return s.queryAllPiece(APIstub)
	}

	return shim.Error("COINCOUN.")
}

/*
 * The queryTuna method *
Used to view the records of one particular tuna
It takes one argument -- the key for the tuna in question
 */
func (s *SmartContract) queryPiece(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	pieceAsBytes, _ := APIstub.GetState(args[0])
	if pieceAsBytes == nil {
		return shim.Error("Could not locate piece")
	}
	return shim.Success(pieceAsBytes)
}

/*
 * 	Ici, on initialise notre JDD sur le network
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	piece := []Piece{
		Piece{RefPiece: "", NomPiece:"Aile", NomFabricant: "AirFrance", NumeroLot: "ABCDEF1234567", DateExp: "16/02/2020", 
		LocalisationVille: "Paris", LocalisationEtablissement: "Etablissement", DateCreation:"06/08/2017", DateReception: "", VenteClient: "true"},

		Piece{RefPiece: "", NomPiece:"Aile", NomFabricant: "AirFrance", NumeroLot: "ABCDEF1234567", DateExp: "16/05/2020", 
		LocalisationVille: "Talence", LocalisationEtablissement: "Etablissement1", DateCreation:"06/08/2017", DateReception: "", VenteClient: "true"},

		Piece{RefPiece: "", NomPiece:"Aile", NomFabricant: "AirFrance", NumeroLot: "ABCDEF1234567", DateExp: "16/06/2020", 
		LocalisationVille: "Rennes", LocalisationEtablissement: "Etablissement2", DateCreation:"06/08/2017", DateReception: "", VenteClient: "true"},

		Piece{RefPiece: "", NomPiece:"Helice", NomFabricant: "AirFrance", NumeroLot: "AA1235458AZEE", DateExp: "06/08/2022", 
		LocalisationVille: "Tonneins", LocalisationEtablissement: "Etablissement1", DateCreation:"", DateReception: "", VenteClient: "true"},
	}

	i := 0
	for i < len(piece) {
		fmt.Println("i is ", i)
		pieceAsBytes, _ := json.Marshal(piece[i])
		APIstub.PutState(strconv.Itoa(i+1), pieceAsBytes)
		fmt.Println("Added", piece[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordTuna method *
Fisherman like Sarah would use to record each of her tuna catches. 
This method takes in five arguments (attributes to be saved in the ledger). 
 */
func (s *SmartContract) recordPiece(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	var piece = Piece{ RefPiece: args[1], NomPiece: args[2], NomFabricant: args[3],
	 NumeroLot: args[4], DateExp: args[5], LocalisationVille: args[6], LocalisationEtablissement: args[7], DateCreation: args[8], DateReception: args[9], VenteClient: args[10]}

	pieceAsBytes, _ := json.Marshal(piece)
	err := APIstub.PutState(args[0], pieceAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record piece catch: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllTuna method *
allows for assessing all the records added to the ledger(all tuna catches)
This method does not take any arguments. Returns JSON string containing results. 
 */
func (s *SmartContract) queryAllPiece(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllPiece:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
