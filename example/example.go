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
