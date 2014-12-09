/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// The controller manager is responsible for monitoring replication
// controllers, and creating corresponding pods to achieve the desired
// state.  It uses the API to listen for new controllers and to create/delete
// pods.
package main

import (
	"flag"
	"math"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/cloudprovider"
	minionControllerPkg "github.com/GoogleCloudPlatform/kubernetes/pkg/cloudprovider/controller"
	replicationControllerPkg "github.com/GoogleCloudPlatform/kubernetes/pkg/controller"
	_ "github.com/GoogleCloudPlatform/kubernetes/pkg/healthz"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/master/ports"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/resources"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/service"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/version/verflag"
	"github.com/golang/glog"
)

var (
	port            = flag.Int("port", ports.ControllerManagerPort, "The port that the controller-manager's http service runs on")
	address         = util.IP(net.ParseIP("127.0.0.1"))
	clientConfig    = &client.Config{}
	cloudProvider   = flag.String("cloud_provider", "", "The provider for cloud services.  Empty string for no provider.")
	cloudConfigFile = flag.String("cloud_config", "", "The path to the cloud provider configuration file.  Empty string for no configuration file.")
	minionRegexp    = flag.String("minion_regexp", "", "If non empty, and -cloud_provider is specified, a regular expression for matching minion VMs.")
	machineList     util.StringList
	// TODO: Discover these by pinging the host machines, and rip out these flags.
	nodeMilliCPU = flag.Int64("node_milli_cpu", 1000, "The amount of MilliCPU provisioned on each node")
	nodeMemory   = flag.Int64("node_memory", 3*1024*1024*1024, "The amount of memory (in bytes) provisioned on each node")
)

func init() {
	flag.Var(&address, "address", "The IP address to serve on (set to 0.0.0.0 for all interfaces)")
	flag.Var(&machineList, "machines", "List of machines to schedule onto, comma separated.")
	client.BindClientConfigFlags(flag.CommandLine, clientConfig)
}

func verifyMinionFlags() {
	if *cloudProvider == "" || *minionRegexp == "" {
		if len(machineList) == 0 {
			glog.Info("No machines specified!")
		}
		return
	}
	if len(machineList) != 0 {
		glog.Info("-machines is overwritten by -minion_regexp")
	}
}

func main() {
	flag.Parse()
	util.InitLogs()
	defer util.FlushLogs()

	verflag.PrintAndExitIfRequested()
	verifyMinionFlags()

	if len(clientConfig.Host) == 0 {
		glog.Fatal("usage: controller-manager -master <master>")
	}

	kubeClient, err := client.New(clientConfig)
	if err != nil {
		glog.Fatalf("Invalid API configuration: %v", err)
	}

	if int64(int(*nodeMilliCPU)) != *nodeMilliCPU {
		glog.Warningf("node_milli_cpu is too big for platform. Clamping: %d -> %d",
			*nodeMilliCPU, math.MaxInt32)
		*nodeMilliCPU = math.MaxInt32
	}

	if int64(int(*nodeMemory)) != *nodeMemory {
		glog.Warningf("node_memory is too big for platform. Clamping: %d -> %d",
			*nodeMemory, math.MaxInt32)
		*nodeMemory = math.MaxInt32
	}

	go http.ListenAndServe(net.JoinHostPort(address.String(), strconv.Itoa(*port)), nil)

	endpoints := service.NewEndpointController(kubeClient)
	go util.Forever(func() { endpoints.SyncServiceEndpoints() }, time.Second*10)

	controllerManager := replicationControllerPkg.NewReplicationManager(kubeClient)
	controllerManager.Run(10 * time.Second)

	cloud := cloudprovider.InitCloudProvider(*cloudProvider, *cloudConfigFile)
	nodeResources := &api.NodeResources{
		Capacity: api.ResourceList{
			resources.CPU:    util.NewIntOrStringFromInt(int(*nodeMilliCPU)),
			resources.Memory: util.NewIntOrStringFromInt(int(*nodeMemory)),
		},
	}
	minionController := minionControllerPkg.NewMinionController(cloud, *minionRegexp, machineList, nodeResources, kubeClient)
	minionController.Run(10 * time.Second)

	select {}
}
