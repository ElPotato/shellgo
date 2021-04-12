package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	sgo "github.com/ElPotato/shellgo"
)

// BuildVersion keep information about version tag along with commit ID
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

func saveInFile(path, dataOut, dataIn string) error {
	err := ioutil.WriteFile(path+".in", []byte(dataIn), 0600)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path+".out", []byte(dataOut), 0600)
	if err != nil {
		return err
	}

	return nil
}

func printOutVersion() {
	fmt.Println(BuildVersion)
}
