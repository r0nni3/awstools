package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/r0nni3/awstools/config"
	"github.com/r0nni3/awstools/console"
)

var ec2Svc *ec2.EC2

func init() {
	ec2Svc = setupEC2Service()
}

func setupEC2Service() *ec2.EC2 {
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new EC2 client
	return ec2.New(sess)
}

func queryInstanceStatus(ec2Svc *ec2.EC2) (result *ec2.DescribeInstanceStatusOutput) {
	params := config.GetState()
	instanceStatusInput := &ec2.DescribeInstanceStatusInput{
		DryRun:              params.DryRun,
		InstanceIds:         []*string{params.IID},
		IncludeAllInstances: params.ShowAll,
	}

	result, err := ec2Svc.DescribeInstanceStatus(instanceStatusInput)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(9)
	}

	return result
}

func parseInstanceStatus(result *ec2.DescribeInstanceStatusOutput) (output *console.Output) {
	output = &console.Output{}
	output.Code = 3
	output.Status = "Status: UNKNOWN"
	for _, instance := range result.InstanceStatuses {
		state := *instance.InstanceState.Name
		switch state {
		case ec2.InstanceStateNameRunning:
			output.Code = 0
		case ec2.InstanceStateNamePending:
			fallthrough
		case ec2.InstanceStateNameShuttingDown:
			fallthrough
		case ec2.InstanceStateNameStopping:
			output.Code = 1
		case ec2.InstanceStateNameStopped:
			fallthrough
		case ec2.InstanceStateNameTerminated:
			output.Code = 2
		default:
			output.Code = 3
		}
		output.Status = "Status " + strings.ToUpper(state)
		break
	}
	output.AddDebugMessage(result)
	return output
}

// EC2QueryRun EC2 Query run
func EC2QueryRun() *console.Output {
	result := queryInstanceStatus(ec2Svc)
	return parseInstanceStatus(result)

}

// GetEC2 returns current ec2 service
func GetEC2() *ec2.EC2 {
	return ec2Svc
}
