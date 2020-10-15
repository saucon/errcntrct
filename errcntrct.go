package errcntrct

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type ErrorData struct {
	Code string      `json:"code"`
	Msg string       `json:"msg"`
	Data []ErrorData `json:"errors,omitempty"`
}

var (
	instance *errCntrct
	once sync.Once
)

func InitContract(pathfilename string) error{
	var err error
	var eC map[string]objContract

	if instance != nil {
		if instance.ErrContract == nil {
			instance.ErrContract, err = loadFile(pathfilename)
			if err != nil {
				return err
			}
		}
	}else{
		eC, err = loadFile(pathfilename)
		if err != nil {
			return err
		}
	}

	once.Do(func() {
		instance = &errCntrct{
			ErrContract: eC,
		}
	})

	return nil
}


// if err is single object then you can ignore codefamily using empty string ""
// if err is array then you can fill codefamily with actual error code
func ErrorMessage(statuscode int, codefamily string, err interface{}) (httpCode int, errorData ErrorData) {
	if getContract() == nil {
		errorData.Code = "109999"
		errorData.Msg = "Contract Error, check InitContract(pathfilename string) or .json format"
		httpCode = 500
		return
	}

	switch err.(type) {
	case error:
		contract := getContract().ErrContract[err.(error).Error()]
		if contract.Msg == "" || contract.ConstVar == "" {
			errorData.Code = "9999"
			errorData.Msg = "Unexpected Error"
			httpCode = 500
			return
		}

		errorData.Code = err.(error).Error()
		errorData.Msg = contract.Msg

		return statuscode, errorData
	case []error:
		contract := getContract().ErrContract[codefamily]
		if contract.Msg == "" || contract.ConstVar == "" {
			errorData.Code = "9999"
			errorData.Msg = "Unexpected Error"
			httpCode = 500
			for _ ,e := range err.([]error) {
				obj := getContract().ErrContract[e.Error()]
				errorData.Data = append(errorData.Data, ErrorData{
					Code: e.Error(),
					Msg: obj.Msg,
				})
			}
			return
		}

		errorData.Code = codefamily
		errorData.Msg = contract.Msg

		for _ ,e := range err.([]error) {
			obj := getContract().ErrContract[e.Error()]
			errorData.Data = append(errorData.Data, ErrorData{
				Code: e.Error(),
				Msg: obj.Msg,
			})
		}


		return statuscode, errorData

	default:
		errorData.Code = "109998"
		errorData.Msg = "Unknown err Type"
		httpCode = 500
	}

	return httpCode, errorData
}


type errCntrct struct {
	ErrContract map[string]objContract
}

type objContract struct {
	ConstVar 	string `json:"var"`
	Msg 		string `json:"msg"`
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

func getContract() *errCntrct {
	return instance
}


// for testing only
func setContractInstanceNil() {
	instance = nil
}

//for testing only
func resetContractMapAtInstance() {
	if instance != nil {
		instance.ErrContract = nil
	}
}