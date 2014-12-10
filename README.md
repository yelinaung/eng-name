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

import "github.com/yelinaung/eng-name"
import "fmt"

func main() {
    fmt.Println()
    // You have to seed yourself. For example, like this.
    RandName := New(time.Now().UTC().UnixNano())
    fmt.Printf("Man name : %v \n", RandName.GetMenName())
    fmt.Printf("Woman name : %v \n", RandName.GetWomenName())
    fmt.Println()
}
```

License
-------
MIT
