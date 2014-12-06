eng-name
=========

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
    fmt.Printf("Man name : %v \n", engname.GetMenName())
    fmt.Printf("Woman name : %v \n", engname.GetWomenName())
    fmt.Println()
}
```

License
-------
MIT
