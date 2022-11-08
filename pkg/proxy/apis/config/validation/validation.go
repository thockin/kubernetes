/*
Copyright 2017 The Kubernetes Authors.

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

package validation

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/api/validate"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	componentbaseconfig "k8s.io/component-base/config"
	"k8s.io/component-base/metrics"
	kubeproxyconfig "k8s.io/kubernetes/pkg/proxy/apis/config"
	netutils "k8s.io/utils/net"
)

// Validate validates the configuration of kube-proxy
func Validate(config *kubeproxyconfig.KubeProxyConfiguration) field.ErrorList {
	allErrs := field.ErrorList{}

	newPath := field.NewPath("KubeProxyConfiguration")

	effectiveFeatures := utilfeature.DefaultFeatureGate.DeepCopy()
	if err := effectiveFeatures.SetFromMap(config.FeatureGates); err != nil {
		allErrs = append(allErrs, field.Invalid(newPath.Child("featureGates"), config.FeatureGates, err.Error()))
	}

	allErrs = append(allErrs, validateKubeProxyIPTablesConfiguration(config.IPTables, newPath.Child("KubeProxyIPTablesConfiguration"))...)
	if config.Mode == kubeproxyconfig.ProxyModeIPVS {
		allErrs = append(allErrs, validateKubeProxyIPVSConfiguration(config.IPVS, newPath.Child("KubeProxyIPVSConfiguration"))...)
	}
	allErrs = append(allErrs, validateKubeProxyConntrackConfiguration(config.Conntrack, newPath.Child("KubeProxyConntrackConfiguration"))...)
	allErrs = append(allErrs, validateProxyMode(config.Mode, newPath.Child("Mode"))...)
	allErrs = append(allErrs, validateClientConnectionConfiguration(config.ClientConnection, newPath.Child("ClientConnection"))...)

	if config.OOMScoreAdj != nil && (*config.OOMScoreAdj < -1000 || *config.OOMScoreAdj > 1000) {
		allErrs = append(allErrs, field.Invalid(newPath.Child("OOMScoreAdj"), *config.OOMScoreAdj, "must be within the range [-1000, 1000]"))
	}

	allErrs = append(allErrs, validate.GTZ(config.ConfigSyncPeriod.Duration, newPath.Child("ConfigSyncPeriod"))...)

	if netutils.ParseIPSloppy(config.BindAddress) == nil {
		allErrs = append(allErrs, field.Invalid(newPath.Child("BindAddress"), config.BindAddress, "not a valid textual representation of an IP address"))
	}

	if config.HealthzBindAddress != "" {
		allErrs = append(allErrs, validateHostPort(config.HealthzBindAddress, newPath.Child("HealthzBindAddress"))...)
	}
	allErrs = append(allErrs, validateHostPort(config.MetricsBindAddress, newPath.Child("MetricsBindAddress"))...)

	if config.ClusterCIDR != "" {
		cidrs := strings.Split(config.ClusterCIDR, ",")
		switch {
		case len(cidrs) > 2:
			allErrs = append(allErrs, field.Invalid(newPath.Child("ClusterCIDR"), config.ClusterCIDR, "only one CIDR allowed or a valid DualStack CIDR (e.g. 10.100.0.0/16,fde4:8dba:82e1::/48)"))
		// if DualStack and two cidrs validate if there is at least one of each IP family
		case len(cidrs) == 2:
			isDual, err := netutils.IsDualStackCIDRStrings(cidrs)
			if err != nil || !isDual {
				allErrs = append(allErrs, field.Invalid(newPath.Child("ClusterCIDR"), config.ClusterCIDR, "must be a valid DualStack CIDR (e.g. 10.100.0.0/16,fde4:8dba:82e1::/48)"))
			}
		// if we are here means that len(cidrs) == 1, we need to validate it
		default:
			if _, _, err := netutils.ParseCIDRSloppy(config.ClusterCIDR); err != nil {
				allErrs = append(allErrs, field.Invalid(newPath.Child("ClusterCIDR"), config.ClusterCIDR, "must be a valid CIDR block (e.g. 10.100.0.0/16 or fde4:8dba:82e1::/48)"))
			}
		}
	}

	if _, err := utilnet.ParsePortRange(config.PortRange); err != nil {
		allErrs = append(allErrs, field.Invalid(newPath.Child("PortRange"), config.PortRange, "must be a valid port range (e.g. 300-2000)"))
	}

	allErrs = append(allErrs, validateKubeProxyNodePortAddress(config.NodePortAddresses, newPath.Child("NodePortAddresses"))...)
	allErrs = append(allErrs, validateShowHiddenMetricsVersion(config.ShowHiddenMetricsForVersion, newPath.Child("ShowHiddenMetricsForVersion"))...)
	if config.DetectLocalMode == kubeproxyconfig.LocalModeBridgeInterface {
		allErrs = append(allErrs, validateInterface(config.DetectLocal.BridgeInterface, newPath.Child("InterfaceName"))...)
	}
	if config.DetectLocalMode == kubeproxyconfig.LocalModeInterfaceNamePrefix {
		allErrs = append(allErrs, validateInterface(config.DetectLocal.InterfaceNamePrefix, newPath.Child("InterfacePrefix"))...)
	}

	return allErrs
}

func validateKubeProxyIPTablesConfiguration(config kubeproxyconfig.KubeProxyIPTablesConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if config.MasqueradeBit != nil && (*config.MasqueradeBit < 0 || *config.MasqueradeBit > 31) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("MasqueradeBit"), config.MasqueradeBit, "must be within the range [0, 31]"))
	}

	allErrs = append(allErrs, validate.GTZ(config.SyncPeriod.Duration, fldPath.Child("SyncPeriod"))...)
	allErrs = append(allErrs, validate.GEZ(config.MinSyncPeriod.Duration, fldPath.Child("MinSyncPeriod"))...)

	if config.MinSyncPeriod.Duration > config.SyncPeriod.Duration {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("SyncPeriod"), config.MinSyncPeriod, fmt.Sprintf("must be greater than or equal to %s", fldPath.Child("MinSyncPeriod").String())))
	}

	return allErrs
}

func validateKubeProxyIPVSConfiguration(config kubeproxyconfig.KubeProxyIPVSConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, validate.GTZ(config.SyncPeriod.Duration, fldPath.Child("SyncPeriod"))...)
	allErrs = append(allErrs, validate.GEZ(config.MinSyncPeriod.Duration, fldPath.Child("MinSyncPeriod"))...)

	if config.MinSyncPeriod.Duration > config.SyncPeriod.Duration {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("SyncPeriod"), config.MinSyncPeriod, fmt.Sprintf("must be greater than or equal to %s", fldPath.Child("MinSyncPeriod").String())))
	}

	allErrs = append(allErrs, validateIPVSTimeout(config, fldPath)...)
	allErrs = append(allErrs, validateIPVSSchedulerMethod(kubeproxyconfig.IPVSSchedulerMethod(config.Scheduler), fldPath.Child("Scheduler"))...)
	allErrs = append(allErrs, validateIPVSExcludeCIDRs(config.ExcludeCIDRs, fldPath.Child("ExcludeCidrs"))...)

	return allErrs
}

func validateKubeProxyConntrackConfiguration(config kubeproxyconfig.KubeProxyConntrackConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if config.MaxPerCore != nil {
		allErrs = append(allErrs, validate.GEZ(*config.MaxPerCore, fldPath.Child("MaxPerCore"))...)
	}

	if config.Min != nil {
		allErrs = append(allErrs, validate.GEZ(*config.Min, fldPath.Child("Min"))...)
	}

	allErrs = append(allErrs, validate.GEZ(config.TCPEstablishedTimeout.Duration, fldPath.Child("TCPEstablishedTimeout"))...)
	allErrs = append(allErrs, validate.GEZ(config.TCPCloseWaitTimeout.Duration, fldPath.Child("TCPCloseWaitTimeout"))...)

	return allErrs
}

func validateProxyMode(mode kubeproxyconfig.ProxyMode, fldPath *field.Path) field.ErrorList {
	if runtime.GOOS == "windows" {
		return validateProxyModeWindows(mode, fldPath)
	}

	return validateProxyModeLinux(mode, fldPath)
}

func validateProxyModeLinux(mode kubeproxyconfig.ProxyMode, fldPath *field.Path) field.ErrorList {
	validModes := sets.NewString(
		string(kubeproxyconfig.ProxyModeIPTables),
		string(kubeproxyconfig.ProxyModeIPVS),
	)

	if mode == "" || validModes.Has(string(mode)) {
		return nil
	}

	errMsg := fmt.Sprintf("must be %s or blank (blank means the best-available proxy [currently iptables])", strings.Join(validModes.List(), ","))
	return field.ErrorList{field.Invalid(fldPath.Child("ProxyMode"), string(mode), errMsg)}
}

func validateProxyModeWindows(mode kubeproxyconfig.ProxyMode, fldPath *field.Path) field.ErrorList {
	validModes := sets.NewString(
		string(kubeproxyconfig.ProxyModeKernelspace),
	)

	if mode == "" || validModes.Has(string(mode)) {
		return nil
	}

	errMsg := fmt.Sprintf("must be %s or blank (blank means the most-available proxy [currently 'kernelspace'])", strings.Join(validModes.List(), ","))
	return field.ErrorList{field.Invalid(fldPath.Child("ProxyMode"), string(mode), errMsg)}
}

func validateClientConnectionConfiguration(config componentbaseconfig.ClientConnectionConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, validate.GEZ(config.Burst, fldPath.Child("Burst"))...)
	return allErrs
}

func validateHostPort(input string, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	hostIP, port, err := net.SplitHostPort(input)
	if err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, input, "must be IP:port"))
		return allErrs
	}

	if ip := netutils.ParseIPSloppy(hostIP); ip == nil {
		allErrs = append(allErrs, field.Invalid(fldPath, hostIP, "must be a valid IP"))
	}

	if p, err := strconv.Atoi(port); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, port, "must be a valid port"))
	} else if p < 1 || p > 65535 {
		allErrs = append(allErrs, field.Invalid(fldPath, port, "must be a valid port"))
	}

	return allErrs
}

func validateIPVSSchedulerMethod(scheduler kubeproxyconfig.IPVSSchedulerMethod, fldPath *field.Path) field.ErrorList {
	supportedMethod := []kubeproxyconfig.IPVSSchedulerMethod{
		kubeproxyconfig.RoundRobin,
		kubeproxyconfig.WeightedRoundRobin,
		kubeproxyconfig.LeastConnection,
		kubeproxyconfig.WeightedLeastConnection,
		kubeproxyconfig.LocalityBasedLeastConnection,
		kubeproxyconfig.LocalityBasedLeastConnectionWithReplication,
		kubeproxyconfig.SourceHashing,
		kubeproxyconfig.DestinationHashing,
		kubeproxyconfig.ShortestExpectedDelay,
		kubeproxyconfig.NeverQueue,
		"",
	}
	allErrs := field.ErrorList{}
	var found bool
	for i := range supportedMethod {
		if scheduler == supportedMethod[i] {
			found = true
			break
		}
	}
	// Not found
	if !found {
		errMsg := fmt.Sprintf("must be in %v, blank means the default algorithm method (currently rr)", supportedMethod)
		allErrs = append(allErrs, field.Invalid(fldPath.Child("Scheduler"), string(scheduler), errMsg))
	}
	return allErrs
}

func validateKubeProxyNodePortAddress(nodePortAddresses []string, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	for i := range nodePortAddresses {
		if _, _, err := netutils.ParseCIDRSloppy(nodePortAddresses[i]); err != nil {
			allErrs = append(allErrs, field.Invalid(fldPath.Index(i), nodePortAddresses[i], "must be a valid CIDR"))
		}
	}

	return allErrs
}

func validateIPVSTimeout(config kubeproxyconfig.KubeProxyIPVSConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, validate.GEZ(config.TCPTimeout.Duration, fldPath.Child("TCPTimeout"))...)
	allErrs = append(allErrs, validate.GEZ(config.TCPFinTimeout.Duration, fldPath.Child("TCPFinTimeout"))...)
	allErrs = append(allErrs, validate.GEZ(config.UDPTimeout.Duration, fldPath.Child("UDPTimeout"))...)
	return allErrs
}

func validateIPVSExcludeCIDRs(excludeCIDRs []string, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	for i := range excludeCIDRs {
		if _, _, err := netutils.ParseCIDRSloppy(excludeCIDRs[i]); err != nil {
			allErrs = append(allErrs, field.Invalid(fldPath.Index(i), excludeCIDRs[i], "must be a valid CIDR"))
		}
	}
	return allErrs
}

func validateShowHiddenMetricsVersion(version string, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	errs := metrics.ValidateShowHiddenMetricsVersion(version)
	for _, e := range errs {
		allErrs = append(allErrs, field.Invalid(fldPath, version, e.Error()))
	}

	return allErrs
}

func validateInterface(iface string, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if len(iface) == 0 {
		allErrs = append(allErrs, field.Invalid(fldPath, iface, "must not be empty"))
	}
	return allErrs
}
