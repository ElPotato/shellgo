package shellgo_test

import (
	sgo "github.com/ElPotato/shellgo"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const inputDir = "./test_inputs/"

func Test_parseOutputValidity(t *testing.T) {
	inputs := map[string]string{
		inputDir + "0_test_input.txt": "4883ec2048896c2418488d6c241848c74424280000000048c704240500000048c74424080a00000048c744241003000000488b0c24488b542408eb0048b8abaaaaaaaaaaaaaa4889d348f7ea4801da48d1fa48c1fb3f4829da4801d148894c2428488b6c24184883c420c3",
		inputDir + "1_test_input.txt": "48c74424080000000048c74424082f010000c3",
		inputDir + "2_test_input.txt": "48c74424080000000048c744240800000000c3"}

	for path, value := range inputs {
		file, _ := ioutil.ReadFile(path)
		parsed := sgo.Parse(string(file))
		assert.Equal(t, value, parsed)
	}
}
