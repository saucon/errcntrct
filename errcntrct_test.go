package errcntrct

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestLoadContract(t *testing.T) {
	for i := 0; i < 2; i++ {
		var err error
		resetInstance()
		resetContractMapAtInstance()
		err = InitContract("errorContract.json")
		if err != nil {
			fmt.Println(err)
		}

		assert.NoError(t, err)

		statusCode, errorData := ErrorMessage(500, "", errors.New("1002"))
		assert.Equal(t, 500, statusCode)
		assert.Equal(t, "1002", errorData.Code)
	}
}

func BenchmarkLoadContract(b *testing.B) {
	var err error
	resetInstance()
	resetContractMapAtInstance()
	err = InitContract("errorContract.json")
	if err != nil {
		fmt.Println(err)
	}

	assert.NoError(b, err)

	statusCode, errorData := ErrorMessage(500, "", errors.New("1002"))
	assert.Equal(b, 500, statusCode)
	assert.Equal(b, "1002", errorData.Code)
}

func TestErrorMessageWithGoRoutine(t *testing.T) {
	var err error
	resetContractMapAtInstance()
	err = InitContract("errorContract.json")
	if err != nil {
		fmt.Println(err)
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		fmt.Println(ErrorMessage(500, "", errors.New("1")))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println(ErrorMessage(500, "", errors.New("1001")))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println(ErrorMessage(500, "", errors.New("1002")))
		wg.Done()
	}()

	wg.Wait()

	assert.NoError(t, err)
}

func TestWrongPathFileName(t *testing.T) {
	var err error
	resetInstance()
	resetContractMapAtInstance()
	err = InitContract(".json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ErrorMessage(500, "", errors.New("1002")))

	assert.Error(t, err)
	assert.Nil(t, getContractInstance())
	assert.Nil(t, getContractInstance().getErrContract())
}

func TestWrongJsonFormat(t *testing.T) {
	var err error
	resetContractMapAtInstance()
	err = InitContract("test/errorContract_test.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ErrorMessage(500, "", errors.New("1002")))
	fmt.Println(getContractInstance())

	assert.NoError(t, err)
}

func TestWrongJsonFormat2(t *testing.T) {
	var err error
	resetContractMapAtInstance()
	err = InitContract("test/errorContractWrongFormat_test.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ErrorMessage(500, "", errors.New("1002")))
	fmt.Println(getContractInstance())

	assert.Error(t, err)
}

func TestErrorArray(t *testing.T) {
	var err error
	resetContractMapAtInstance()
	err = InitContract("errorContract.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ErrorMessage(500, "1000", []error{
		errors.New("1002"), errors.New("1002"), errors.New("1002"),
	}))
	fmt.Println(getContractInstance())

	assert.NoError(t, err)
}

func TestErrorArrayCodeFamilyNotFound(t *testing.T) {
	var err error
	resetContractMapAtInstance()
	err = InitContract("errorContract.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ErrorMessage(500, "6969", []error{
		errors.New("1002"), errors.New("6969"), errors.New("1002"),
	}))
	fmt.Println(getContractInstance())

	assert.NoError(t, err)
}

func TestUnknownErrorType(t *testing.T) {
	var err error
	resetContractMapAtInstance()
	err = InitContract("errorContract.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ErrorMessage(500, "", 10))
	fmt.Println(getContractInstance())

	assert.NoError(t, err)
}

func TestNilInstance(t *testing.T) {
	var err error
	resetContractMapAtInstance()
	resetInstance()
	err = InitContract("test/errorContractWrongFormat_test.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ErrorMessage(500, "", 10))
	fmt.Println(getContractInstance())

	assert.Error(t, err)
	assert.Nil(t, getContractInstance())
}
