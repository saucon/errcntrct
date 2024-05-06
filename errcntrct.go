package errcntrct

import (
	"encoding/json"
	"os"
	"sync"
)

const (
	CODE_109999    = "109999"
	MESSAGE_109999 = "Contract Error, check InitContract(pathfilename string) or .json format"

	CODE_109998    = "109998"
	MESSAGE_109998 = "Unknown err Type"

	CODE_9999    = "9999"
	MESSAGE_9999 = "Unexpected Error"
)

var (
	instance    *errContractInstance = nil
	once        sync.Once
	lock        sync.Mutex
	errContract map[string]objContract = nil
)

type errContractInstance struct {
}

func (ec *errContractInstance) getErrContract() map[string]objContract {
	if ec == nil {
		return nil
	}
	return errContract
}

func (ec *errContractInstance) setErrContract(eC map[string]objContract) {
	if ec == nil {
		return
	}
	errContract = eC
}

type objContract struct {
	ConstVar string `json:"var"`
	Msg      string `json:"msg"`
}

type ErrorData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data []ErrorData `json:"errors,omitempty"`
}

func InitContract(pathFilename string) error {
	var err error
	var eC map[string]objContract

	if instance != nil {
		if instance.getErrContract() == nil {
			eC, err = loadFile(pathFilename)
			if err != nil {
				return err
			}
			instance.setErrContract(eC)
		}
	} else {
		eC, err = loadFile(pathFilename)
		if err != nil {
			return err
		}
	}

	once.Do(func() {
		instance = &errContractInstance{}
		instance.setErrContract(eC)
	})

	return nil
}

// ErrorMessage is a function to get error message from error code
// if err is single object then you can ignore codeFamily using empty string ""
// if err is array then you can fill codeFamily with actual error code
func ErrorMessage(statusCode int, codeFamily string, err interface{}) (int, ErrorData) {
	var (
		errorData ErrorData
		httpCode  int
	)
	if getContractInstance() == nil {
		errorData.Code = CODE_109999
		errorData.Msg = MESSAGE_109999
		httpCode = 500
		return httpCode, errorData
	}

	switch err.(type) {
	case error:
		contract := getContractInstance().getErrContract()[err.(error).Error()]
		if contract.Msg == "" || contract.ConstVar == "" {
			errorData.Code = CODE_9999
			errorData.Msg = MESSAGE_9999
			httpCode = 500
			return httpCode, errorData
		}

		errorData.Code = err.(error).Error()
		errorData.Msg = contract.Msg

		return statusCode, errorData
	case []error:
		contract := getContractInstance().getErrContract()[codeFamily]
		if contract.Msg == "" || contract.ConstVar == "" {
			errorData.Code = CODE_9999
			errorData.Msg = MESSAGE_9999
			httpCode = 500
			for _, e := range err.([]error) {
				obj := getContractInstance().getErrContract()[e.Error()]
				errorData.Data = append(errorData.Data, ErrorData{
					Code: e.Error(),
					Msg:  obj.Msg,
				})
			}
			return httpCode, errorData
		}

		errorData.Code = codeFamily
		errorData.Msg = contract.Msg

		for _, e := range err.([]error) {
			obj := getContractInstance().getErrContract()[e.Error()]
			errorData.Data = append(errorData.Data, ErrorData{
				Code: e.Error(),
				Msg:  obj.Msg,
			})
		}

		return statusCode, errorData

	default:
		errorData.Code = CODE_109998
		errorData.Msg = MESSAGE_109998
		httpCode = 500
	}

	return httpCode, errorData
}

func loadFile(pathFilename string) (map[string]objContract, error) {
	var file []byte
	var err error
	var eC map[string]objContract

	file, err = os.ReadFile(pathFilename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &eC)
	if err != nil {
		return nil, err
	}

	return eC, nil
}

func getContractInstance() *errContractInstance {
	return instance
}

// for testing only
func resetInstance() {
	lock.Lock()
	defer lock.Unlock()
	instance = nil     // Clear the existing instance
	once = sync.Once{} //
}

// for testing only
func resetContractMapAtInstance() {
	if instance != nil {
		instance.setErrContract(nil)
	}
}
