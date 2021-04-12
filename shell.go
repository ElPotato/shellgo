package shellgo

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"
)

// Output keep raw which is default, unformatted bytecode.
type Output struct {
	raw string
}

// Format0x format output adding 0x at the beginning at every code.
func (o Output) Format0x() string {
	var buffer bytes.Buffer

	bytesArray := []byte(o.raw)

	for c := 0; c <= len(bytesArray)-2; {
		buffer.WriteString("0x")
		buffer.WriteString(string(bytesArray[c]))
		buffer.WriteString(string(bytesArray[c+1]))
		buffer.WriteString(", ")

		c += 2
	}

	return buffer.String()
}

// Default return unformatted raw output.
func (o Output) Default() string {
	return o.raw
}

// STDINReader append standard input characters.
func STDINReader() []rune {
	var output []rune
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}

		output = append(output, input)
	}

	return output
}

// Parse extracts bytecode from go tool objdump output.
func Parse(input string) Output {
	var buffer bytes.Buffer
	r, _ := regexp.Compile("\t[0-9a-f]+")
	matched := r.FindAllString(input, -1)

	for _, e := range matched {
		buffer.WriteString(strings.Trim(e, "\t"))
	}

	return Output{raw: buffer.String()}
}
