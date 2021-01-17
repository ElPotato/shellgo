package main

import sgo "github.com/ElPotato/shellgo"
import "fmt"

func main() {
	data := sgo.STDINReader()
	ready := sgo.Parse(string(data))

	fmt.Printf(ready)
}