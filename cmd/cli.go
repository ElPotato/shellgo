package main

import sgo "github.com/ElPotato/shellgo"
import "fmt"

func main() {
	data := sgo.Reader()
	ready := sgo.Parse(data)

	fmt.Printf(ready)
}