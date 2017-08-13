package console

import (
	"fmt"
	"os"

	"github.com/r0nni3/awstools/config"
)

// Output Describes the content of an output
type Output struct {
	Status    string
	Code      int
	DebugMsgs []interface{}
}

// PrintOutput Prints program output and exits program with specified exit code
func (output *Output) PrintOutput() {
	fmt.Println(output.Status)
	params := config.GetState()
	if *params.Debug {
		for _, msg := range output.DebugMsgs {
			fmt.Println(msg)
		}
	}
	os.Exit(output.Code)
}

// AddDebugMessage Appends a message to the debug message slice
func (output *Output) AddDebugMessage(value interface{}) {
	output.DebugMsgs = append(output.DebugMsgs, value)
}
