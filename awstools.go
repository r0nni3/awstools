package awstools

import (
	"github.com/r0nni3/awstools/config"
	"github.com/r0nni3/awstools/service"
)

func init() {
	config.Setup()
}

// CheckEC2Status Get instance status.
func CheckEC2Status() {
	output := service.EC2QueryRun()
	output.PrintOutput()
}
