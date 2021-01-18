package main

import (
	"flag"
	sgo "github.com/ElPotato/shellgo"
	"io/ioutil"
)

var BuildVersion string

func main() {
	dumpPath := flag.String("dump", "", "path for input/output data dump")
	version := flag.Bool("version", false, "print out software version")
	flag.Parse()

	if *version {
		printOutVersion()
		return
	}

	data := sgo.STDINReader()
	ready := sgo.Parse(string(data))

	println(ready)

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
	println(BuildVersion)
}