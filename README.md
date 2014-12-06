eng-name
=========

Ramdom Men/Women name generator. Golang port of [this lib](https://github.com/Zwenexsys/eng_name).

```bash

Man name :  HenryeSingleton
Womn name :  MargerreScott

```

Example
-------

```golang
package main

import "github.com/yelinaung/eng-name"

func main() {
    println("Man name : ", engname.GetMenName())
    println("Womn name : ", engname.GetWomenName())
}
```

License
-------
MIT
