package shellgo

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"
)

type Output struct {
	RAW string
}

func (o Output) Format0x() string {
	var buffer bytes.Buffer
	bytesArray := []byte(o.RAW)

	for c := 0; c <= len(bytesArray)-2; {
		buffer.WriteString("0x")
		buffer.WriteString(string(bytesArray[c]))
		buffer.WriteString(string(bytesArray[c+1]))
		buffer.WriteString(", ")
		c = c + 2
	}

	return buffer.String()
}

func (o Output) Default() string {
	return o.RAW
}

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

func Parse(input string) Output {
	var buffer bytes.Buffer
	r, _ := regexp.Compile("\t[0-9a-f]+")
	matched := r.FindAllString(input, -1)

	for _, e := range matched {
		buffer.WriteString(strings.Trim(e, "\t"))
	}

	return Output{RAW: buffer.String()}
}
