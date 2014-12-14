eng-name  [![Build Status](https://travis-ci.org/yelinaung/eng-name.svg?branch=master)](https://travis-ci.org/yelinaung/eng-name)
=============

Ramdom Men/Women name generator. Golang port of [this lib](https://github.com/Zwenexsys/eng_name).

```bash

Man name : Hughe Craster
Woman name : Mary Lighten

```

View the [docs](https://godoc.org/github.com/yelinaung/eng-name). 

Example
-------

```go
package main

import (
	"fmt"
	"github.com/yelinaung/eng-name"
	"time"
)

func main() {
	fmt.Println()
	// You have to seed yourself
	RandName := engname.New(time.Now().UTC().UnixNano())
	// Generating random men names
	fmt.Printf("Man name : %v \n", RandName.GetMenName())
	// Generating random women names
	fmt.Printf("Woman name : %v \n", RandName.GetWomenName())
	fmt.Println()
}
```

License
-------
MIT
