package main

import (
	"flag"
	"fmt"
	sgo "github.com/ElPotato/shellgo"
	"io/ioutil"
)

var BuildVersion string

func main() {
	dumpPath := flag.String("dump", "", "path for input/output data dump")
	version := flag.Bool("version", false, "print out software version")
	format0x := flag.Bool("0x", false, "format output e.g. 0x01, 0x02...")
	flag.Parse()

	if *version {
		printOutVersion()
		return
	}

	data := sgo.STDINReader()
	var ready string
	if *format0x {
		ready = sgo.Parse(string(data)).Format0x()
	} else {
		ready = sgo.Parse(string(data)).Default()
	}

	fmt.Println(ready)

	if *dumpPath != "" {
		err := saveInFile(*dumpPath, ready, string(data))
		if err != nil {
			panic(err)
		}
	}
}

func saveInFile(path string, dataOut, dataIn string) error {
	errIn := ioutil.WriteFile(path+".in", []byte(dataIn), 0644)
	if errIn != nil {
		return errIn
	}

	errOut := ioutil.WriteFile(path+".out", []byte(dataOut), 0644)
	if errOut != nil {
		return errOut
	}
	return nil
}

func printOutVersion() {
	fmt.Println(BuildVersion)
}
