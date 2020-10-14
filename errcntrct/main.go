package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


// ./errcntrct-gen -i "errorContract.json" -p "contract" -o "example-generated/contract.go"
func main (){
	var contractMap map[string]objContract
	var err error

	contractMap, err = loadFile("errorContract.json")
	fmt.Println("Hello World")


	f, err := os.Create("errcntrct/example-generated/const.go")
	if err != nil {

	}
	defer f.Close()

	f.WriteString("package contract\n\n\n")
	f.WriteString(`// this file is example-generated using .json file contract`+"\n")
	f.WriteString("const (\n")

	for code, obj := range contractMap {
		f.WriteString("\t"+obj.ConstVar+" = "+`"`+code+`"`+"\n")
	}
 	f.WriteString(")\n")
	f.Sync()
}


func loadFile(pathfilename string) (map[string]objContract,error){
	var file []byte
	var err error
	var eC map[string]objContract

	file, err = ioutil.ReadFile(pathfilename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &eC)
	if err != nil {
		return nil, err
	}

	return eC, nil
}

type objContract struct {
	ConstVar 	string `json:"var"`
	Msg 		string `json:"msg"`
}