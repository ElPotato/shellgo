package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	ready := Parse(output)
	fmt.Printf(ready)
}

func Parse(input []rune) string {
	var buffer bytes.Buffer
	r, _ := regexp.Compile("\t[0-9a-fA-F]+")
	matched := r.FindAllString(string(input), -1)

	for _, e := range matched {
		buffer.WriteString(strings.Trim(e, "\t"))
	}

	return buffer.String()
}