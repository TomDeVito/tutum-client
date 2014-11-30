# tutum-client

A (work-in-progress) Go client for [tutum](https://www.tutum.co/). The structure is
similar to the
[python-tutum](https://github.com/tutumcloud/python-tutum/) but is
instantiable rather than a single package.

# Installation

```sh
$ go get github.com/garslo/tutum-client
```

# Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/garslo/tutum-client"
)

func main() {
	api := tutum.NewTutumAPI("username", "api_key", "https://dashboard.tutum.co/api/v1")
	response, err := api.Provider().Fetch("digitalocean")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(response)
}
```
