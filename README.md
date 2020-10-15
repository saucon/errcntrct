# errcntrct

errcntrct is a library for create error contract and return simple struct. This is compatible with c.JSON in gin framework. You will define error contract in JSON format, check `errorContract.json` file for example.

## How to Use

1. Initiate the lib

`InitContract(pathfilename string)`

Call this function in your main project. This function will load your errorContract.json file.

2. User ErrorMessage

`ErrorMessage(statuscode int, codefamily string, err interface{})(httpCode int, errorData ErrorData)`

statuscode = http status code

codefamily = Your error code family, this is for `[]error`. You can just ignore it if you use `error`.

err = You can use `error` or `[]error`

## CLI errcntrct

Install cli 

        go install github.com/Saucon/errcntrct/errcntrct




A tool to convert your JSON contract into golang. 
For example:

       errcntrct gen -i errorContract.json -o example/const.go

With this you dont need to declare const in go, just copy 
the generated code to your project

Usage:
  errcntrct gen [flags]

Flags:
  -h, --help             help for gen
  -o, --output string    Output const file 
  -p, --package string   Package const file 
  -s, --source string    Source .json file

Global Flags:
      --config string   config file (default is $HOME/.errcntrct.yaml)
