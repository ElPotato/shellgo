package main

import (
	"bytes"
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	_, err := os.Stdin.Stat()
    if err != nil {
        panic(err)
    }

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	str := string(output)
	r, _ := regexp.Compile("\t[0-9a-fA-F]+")
	matched := r.FindAllString(str, -1)
	var buffer bytes.Buffer

	for _, e := range matched {
		buffer.WriteString(strings.Trim(e, "\t"))
	}

	fmt.Printf(buffer.String())
}