package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GenerateCodeFile(source string, output string, packagename string, f func(string, string, string) error) error {
	if err := f(source, output, packagename); err != nil {
		return err
	}

	return nil
}

func TemplateJSONtoGolangConst(source string, output string, packagename string) error {
	var contractMap map[string]objContract
	var err error

	contractMap, err = loadFile(source)

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(`// This file is generated using errcntrct tool.` + "\n" +
		`// Check out for more info "https://github.com/Saucon/errcntrct"` +
		"\n" +
		"package " + packagename + "\n\n" +
		`import "errors"` + "\n\n")
	if err != nil {
		return err
	}

	fmt.Print()

	// print const block
	_, err = f.WriteString("const (\n")
	if err != nil {
		return err
	}
	for code, obj := range contractMap {
		_, err = f.WriteString(
			"\t" + obj.ConstVar + "_const" + " = " + `"` + code + `"` + "\n")
		if err != nil {
			return err
		}
	}
	_, err = f.WriteString(")\n")
	if err != nil {
		return err
	}

	// print error block
	_, err = f.WriteString("var (\n")
	if err != nil {
		return err
	}
	for _, obj := range contractMap {
		_, err = f.WriteString(
			"\t" + obj.ConstVar + " = " + `errors.New(` + obj.ConstVar + "_const" + `)` + "\n")
		if err != nil {
			return err
		}
	}
	_, err = f.WriteString(")\n")
	if err != nil {
		return err
	}

	// finish with all
	err = f.Sync()
	if err != nil {
		return err
	}

	return nil
}

func loadFile(pathfilename string) (map[string]objContract, error) {
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
	ConstVar string `json:"var"`
	Msg      string `json:"msg"`
}
