package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"k8s.io/utils/feature"
)

var (
	libGates = &feature.GateSet{}
	f1       = libGates.AddOrDie(&feature.Gate{
		Name:    "Feature1",
		Default: false,
		Release: feature.Alpha,
	})
	f2 = libGates.AddOrDie(&feature.Gate{
		Name:    "Feature2",
		Default: true,
		Release: feature.Beta,
	})
	f3 = libGates.AddOrDie(&feature.Gate{
		Name:          "Feature3",
		Default:       true,
		Release:       feature.GA,
		LockToDefault: true,
	})
	f4 = &feature.Gate{
		Name:    "Feature4",
		Default: true,
		Release: feature.Beta,
	}
)

func main() {
	libGates.EnablePFlagControl("gates", pflag.CommandLine)
	libGates.EnableEnvControl("GATE_", func(err error) { panic(err.Error()) })
	pflag.Parse()

	fmt.Println(feature.Enabled(f1))
	fmt.Println(feature.Enabled(f2))
	fmt.Println(feature.Enabled(f3))
	fmt.Println(feature.Enabled(f4))
}
