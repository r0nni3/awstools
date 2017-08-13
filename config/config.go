package config

import (
	"flag"
	"fmt"
	"os"
)

// State variables
type State struct {
	IID     *string
	DryRun  *bool
	ShowAll *bool
	Debug   *bool
	Args    []string
}

var parameters *State

// Setup Initialize command
func Setup() {
	parameters = &State{new(string), new(bool), new(bool), new(bool), []string{}}
	parameters.setupFlags()
	parameters.verifyRequiredFlags()
	parameters.setupArguments()
}

func (parameters *State) setupFlags() {
	flag.StringVar(parameters.IID, "instance-id", "", "Required flag representing the id of an ec2 instance. Example:\n\t\t--instance-id=i-xxxxxxx")
	flag.BoolVar(parameters.DryRun, "dry-run", false, "Optional flag default value is false. Example:\n\t\t--dry-run=true")
	flag.BoolVar(parameters.Debug, "debug", false, "Optional flag default value is false. Example:\n\t\t--debug=true")
	flag.BoolVar(parameters.ShowAll, "show-all", false, "Optional flag default value is false. Example:\n\t\t--show-all=true")
	flag.Parse()
}

func (parameters *State) setupArguments() {
	parameters.Args = flag.Args()
}

func (parameters *State) verifyRequiredFlags() {
	if *parameters.IID == "" {
		fmt.Println("Required flag missing.\n\nUSAGE:")
		flag.PrintDefaults()
		os.Exit(10)
	}
}

// GetState returns parameters values struct
func GetState() *State {
	return parameters
}
