# errcntrct

errcntrct (Error Contract) is a library for create error contract and return simple struct. This is compatible with c.JSON in gin framework. 
You will define error contract in JSON format, check `errorContract.json` file for example.

contractError.json
```json
{
    "1000": {"var": "ErrInvalidRequestFamily", "msg": "Invalid Request"},
    "1002": {"var": "ErrInvalidDateFormat", "msg": "Invalid date format"},
    "1003": {"var": "ErrEmailNotFound", "msg": "Could not find email"},

    "9999": {"var": "ErrUnexpectedError", "msg": "Unexpected Error"}
}
```



## Installation

1. The first need [Go](https://golang.org/), then you can use the below Go command to install errcntrct.

```sh
$ go get -u github.com/Saucon/errcntrct
```

2. Import it in your code:

```go
import "github.com/Saucon/errcntrct"
```

## Quick start


```go
package main

import (
"errors"
"fmt"
"github.com/Saucon/errcntrct"
"net/http"
)

const PATH_TO_JSONFILE = "errorContract.json"

func main() {
    err := errcntrct.InitContract(PATH_TO_JSONFILE)
    if err != nil {
        // handle this error
    }
    
    // This is with single error 
    httpstatuscode, errData := errcntrct.ErrorMessage(http.StatusBadRequest,"", errors.New("1001"))
    fmt.Println("http status code", httpstatuscode)
    fmt.Println("errData struct" , errData)
    
    // This is with []error
    // "1000" in contractError.json is the family code
    httpstatuscode, errData = errcntrct.ErrorMessage(http.StatusBadRequest, "1000", []error{errors.New("1001"),errors.New("1001"),})
    fmt.Println("http status code", httpstatuscode)
    fmt.Println("errData struct" , errData)  

    // you can use in gin with c.JSON(errcntrct.ErrorMessage(http.StatusBadRequest, "1000", []error{errors.New("1001"),errors.New("1001"),}))
    
}
```

statuscode = http status code

codefamily = Your error code family, this is for `[]error`. You can just ignore it if you use `error`.

err = You can use `error` or `[]error`

## CLI errcntrct

Install cli

```sh
$ go install github.com/Saucon/errcntrct/errcntrct
```

A tool to convert your JSON contract into golang. 
For example:

```sh
$ errcntrct gen -s errorContract.json -o errcntrct/example/const.go -p contract
```

With this you dont need to declare const in go, just copy 
the generated code to your project

To more detail about this cli use this command.
```sh
$ errcntrct gen -h                                                             
A tool to convert your JSON contract into golang. 
For example:

        errcntrct gen -i errorContract.json -o example/const.go -p contract


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
```

With this you can generate const in go with json file that you define and your code will be like this.

```go
package main

import (
"errors"
"fmt"
"github.com/Saucon/errcntrct"
"net/http"
)

const PATH_TO_JSONFILE = "errorContract.json"

func main() {
    err := errcntrct.InitContract(PATH_TO_JSONFILE)
    if err != nil {
        // handle this error
    }
    
    // This is with single error 
    httpstatuscode, errData := errcntrct.ErrorMessage(http.StatusBadRequest,"", errors.New("1001"))
    fmt.Println("http status code", httpstatuscode)
    fmt.Println("errData struct" , errData)
    
    // This is with []error
    // "1000" in contractError.json is the family code
    errors := []error{contract.ErrEmailNotFound,contract.ErrEmailNotFound,}
    httpstatuscode, errData = errcntrct.ErrorMessage(http.StatusBadRequest, contract.ErrInvalidRequestFamily, errors)
    fmt.Println("http status code", httpstatuscode)
    fmt.Println("errData struct" , errData)  

    // you can use in gin with c.JSON(errcntrct.ErrorMessage(http.StatusBadRequest, "1000", []error{errors.New("1001"),errors.New("1001"),}))
    
}
```