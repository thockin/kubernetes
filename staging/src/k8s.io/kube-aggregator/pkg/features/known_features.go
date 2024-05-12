package features

import (
	"k8s.io/utils/feature"
)

var (
	libGates = &feature.GateSet{}

	// Every feature gate should add method here following this template:
	//
	// // owner: @username
	// // alpha: v1.4
	// MyFeature featuregate.Feature = "MyFeature"
	//
	// Feature gates should be listed in alphabetical, case-sensitive
	// (upper before any lower case character) order. This reduces the risk
	// of code conflicts because changes are more likely to be scattered
	// across the file.
)

// FeatureGates returns the set of feature gates exposed by this library.
func FeatureGates() *feature.GateSet {
	return libGates
}
