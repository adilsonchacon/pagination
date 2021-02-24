# Pagination

## Instalation

        $ go get github.com/adilsonchacon/pagination

## Test

        $ cd /go/path/github.com/adilsonchacon/pagination
        $ go test -v
	
## Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/adilsonchacon/pagination"
)

func main() {
	pagination := &pagination.PageInfo{CurrentPage: 11, TotalPages: 20, Around: 3, Boundaries: 3}
	err := pagination.Generate()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(pagination.ToString())
}
```
