/*
Copyright 2024 The Kubernetes Authors.

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

// Package feature provides functions and types for managing feature gates,
// which are used to control whether specific functionality is enabled.
package feature

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

//FIXME: unclear if we want to make the API in terms of "Gates" or "Features".
//"Features" is so abstract...

// GateSet is a group of feature gates, which can be controlled through
// various mechanisms.  All Gates in a GateSet must have unique
// names.
type GateSet struct {
	gateMap   map[string]*Gate
	providers []provider
}

// Add adds the specified Gate to the GateSet.
func (gs *GateSet) Add(gate *Gate) (*Gate, error) {
	if gate == nil {
		return nil, fmt.Errorf("Gate must be specified")
	}
	if gate.Name == "" {
		return nil, fmt.Errorf("Gate name must be specified")
	}
	if !gate.Release.valid() {
		return nil, fmt.Errorf("Gate release %q is unsupported", gate.Release)
	}

	if gs.gateMap == nil {
		gs.gateMap = map[string]*Gate{}
	}

	if _, found := gs.gateMap[gate.Name]; found {
		return nil, fmt.Errorf("Gate %q was redefined", gate.Name)
	}
	gs.gateMap[gate.Name] = gate
	gate.set = gs
	return gate, nil
}

// AddOrDie adds the specified Gate to the GateSet, or panics on error.
func (gs *GateSet) AddOrDie(gate *Gate) *Gate {
	g, err := gs.Add(gate)
	if err != nil {
		panic(fmt.Sprintf("GateSet.AddOrDie: %s", err.Error()))
	}
	return g
}

// Merge moves all of the Gates in other to this GateSet.
func (gs *GateSet) Merge(other *GateSet) error {
	for _, g := range other.gateMap {
		if _, err := gs.Add(g); err != nil {
			return err
		}
		delete(other.gateMap, g.Name)
	}
	return nil
}

// MergeOrDie moves all of the Gates in other to this GateSet, or panics on error.
func (gs *GateSet) MergeOrDie(other *GateSet) {
	err := gs.Merge(other)
	if err != nil {
		panic(fmt.Sprintf("GateSet.MergeOrDie: %s", err.Error()))
	}
}

// provider allows gates to be enabled or disabled by some mechanism.  A
// GateSet holds a list of providers and will consult them when code needs to
// know if a specific gate is enabled.
type provider interface {
	// Enabled tests whether a named gate is configured by this provider and if
	// so whether this provider thinks the gate is enabled or disabled.  If a
	// provider does not have a configured state for the specified gate, it
	// should indicate "not found".
	Enabled(name string) (enabled, found bool)
}

const (
	// allAlphaGates is a global toggle for alpha features. Per-feature key
	// values override the default set by allAlphaGates. Examples:
	//   AllAlpha=false,NewFeature=true  will result in newFeature=true
	//   AllAlpha=true,NewFeature=false  will result in newFeature=false
	allAlphaGates = "AllAlpha"

	// allBetaGates is a global toggle for beta features. Per-feature key
	// values override the default set by allBetaGates. Examples:
	//   AllBeta=false,NewFeature=true  will result in NewFeature=true
	//   AllBeta=true,NewFeature=false  will result in NewFeature=false
	allBetaGates = "AllBeta"
)

// EnablePFlagControl exposes this GateSet through flags (using
// github.com/spf13/pflag).  Users can enable or disable gates from the
// commandline.  If an invalid gate name or value is specified, flag parsing
// will fail.
// FIXME: merge with existing gates flag and config
func (gs *GateSet) EnablePFlagControl(flagName string, fs *pflag.FlagSet) {
	helpStrings := []string{
		fmt.Sprintf("%s=true|false (%s - default=%t)", allAlphaGates, Alpha, false),
		fmt.Sprintf("%s=true|false (%s - default=%t)", allBetaGates, Beta, false),
	}
	for _, gate := range gs.gateMap {
		helpStrings = append(helpStrings, fmt.Sprintf("%s=true|false (%s - default=%t)", gate.Name, gate.Release, gate.Default))
	}
	sort.Strings(helpStrings)

	//FIXME: nicer to have a different flag (--list-feature-gates) to enumerate?
	fp := &flagProvider{gates: gs}
	fs.Var(fp, flagName, ""+
		"A set of key=value pairs that describe feature gates for alpha/experimental features. "+
		"Options are:\n"+strings.Join(helpStrings, "\n")+"\n")

	gs.providers = append(gs.providers, fp)
}

// flagProvider implements pflag.Value and provider.
type flagProvider struct {
	gates  *GateSet
	values map[string]bool
}

var _ pflag.Value = (*flagProvider)(nil)
var _ provider = (*flagProvider)(nil)

// Enabled tests whether the named gate is enabled by flags.
// This method is part of the provider interface.
func (fp *flagProvider) Enabled(name string) (bool, bool) {
	val, found := fp.values[name]
	return val, found
}

func (fp *flagProvider) setOne(k string, v bool) {
	if fp.values == nil {
		fp.values = map[string]bool{}
	}
	fp.values[k] = v
}

// Set parses a string of the form "key1=value1,key2=value2,..." into a
// map[string]bool of known keys or returns an error.
// This method is part of the pflag.Value interface.
func (fp *flagProvider) Set(value string) error {
	m := make(map[string]bool)
	for _, s := range strings.Split(value, ",") {
		if len(s) == 0 {
			continue
		}
		arr := strings.SplitN(s, "=", 2)
		k := strings.TrimSpace(arr[0])
		if len(arr) != 2 {
			return fmt.Errorf("missing bool value for %s", k)
		}
		v := strings.TrimSpace(arr[1])
		boolValue, err := strconv.ParseBool(v)
		if err != nil {
			return fmt.Errorf("invalid value of %s=%s, err: %v", k, v, err)
		}
		m[k] = boolValue
	}

	for name, enabled := range m {
		switch {
		case name == allAlphaGates:
			for _, gate := range fp.gates.gateMap {
				if gate.Release == Alpha {
					fp.setOne(gate.Name, enabled)
				}
			}
			continue
		case name == allBetaGates:
			for _, gate := range fp.gates.gateMap {
				if gate.Release == Beta {
					fp.setOne(gate.Name, enabled)
				}
			}
			continue
		}

		// Ensure that specified gates are known, so admins get an error if
		// they use an obsolete or misspelled gate name.
		gate, exists := fp.gates.gateMap[name]
		if !exists {
			return fmt.Errorf("unknown feature gate %q", name)
		}
		// Ensure that specified gates are allowed to be changed.
		if gate.LockToDefault && gate.Default != enabled {
			return fmt.Errorf("feature gate %q (=%t) is locked", name, gate.Default)
		}
		fp.setOne(name, enabled)
	}

	return nil
}

// String returns a string containing all enabled feature gates, formatted as
// "key1=value1,key2=value2,...".
// This method is part of the pflag.Value interface.
func (fp *flagProvider) String() string {
	special := []string{
		fmt.Sprintf("%s=false", allAlphaGates),
		fmt.Sprintf("%s=false", allBetaGates),
	}
	pairs := []string{}
	for _, gate := range fp.gates.gateMap {
		pairs = append(pairs, fmt.Sprintf("%s=%t", gate.Name, Enabled(gate)))
	}
	sort.Strings(pairs)

	return strings.Join(append(special, pairs...), ",")
}

// This method is part of the pflag.Value interface.
func (fp *flagProvider) Type() string {
	return "mapStringBool"
}

// EnableEnvControl exposes this GateSet through environment variables.
// If an invalid gate name or value is specified, errfn will be called.
func (gs *GateSet) EnableEnvControl(prefix string, errfn func(error)) {
	ep := &envProvider{prefix: prefix}

	// Do validation here, since we know env vars can't be changed from
	// outside.
	env := os.Environ()
	for _, e := range env {
		if !strings.HasPrefix(e, ep.prefix) {
			continue
		}
		parts := strings.SplitN(e, "=", 2)
		k := strings.TrimPrefix(parts[0], ep.prefix)
		v := parts[1]

		// Ensure that specified gates are known, so admins get an error if
		// they use an obsolete or misspelled gate name.
		gate, exists := gs.gateMap[k]
		if !exists {
			errfn(fmt.Errorf("unknown feature gate %q", k))
			return
		}

		// Ensure that we can parse the value.
		enabled, err := strconv.ParseBool(v)
		if err != nil {
			errfn(fmt.Errorf("invalid boolean value %q for feature gate %q: %v", v, k, err))
		}

		// Ensure that specified gates are allowed to be changed.
		if gate.LockToDefault && gate.Default != enabled {
			errfn(fmt.Errorf("feature gate %q (=%t) is locked", k, gate.Default))
			return
		}
	}
	gs.providers = append(gs.providers, ep)
}

// envProvider implements provider.
type envProvider struct {
	prefix string
	values map[string]bool
}

var _ provider = (*envProvider)(nil)

// Enabled tests whether the named gate is enabled by environment variables.
// This method is part of the provider interface.
func (ep *envProvider) Enabled(name string) (bool, bool) {
	if ep.values == nil {
		ep.values = map[string]bool{}
	}
	key := fmt.Sprintf("%s%s", ep.prefix, name)
	strVal, found := os.LookupEnv(key)
	if !found {
		return false, false
	}

	// Already validated
	boolVal, _ := strconv.ParseBool(strVal)
	return boolVal, true
}

// Enabled tests whether the named gate is enabled by this provider.
func (gs *GateSet) Enabled(name string) (enabled bool, found bool) {
	foundAny := false
	for _, provider := range gs.providers {
		en, found := provider.Enabled(name)
		if !found {
			continue
		}
		foundAny = true
		if en {
			return true, true
		}
	}
	if foundAny {
		// We didn't early-return so we must have found a `false`
		return false, true
	}
	return false, false
}

// Gate is a single feature gate definition.
type Gate struct {
	Name          string
	Default       bool
	Release       StabilityLevel
	LockToDefault bool
	forcedValue   *bool
	set           *GateSet
}

// StabilityLevel indicates which phase of the feature lifecycle a given gate is
// currently on.
type StabilityLevel string

const (
	Alpha      = StabilityLevel("ALPHA")
	Beta       = StabilityLevel("BETA")
	GA         = StabilityLevel("GA")
	Deprecated = StabilityLevel("DEPRECATED")
)

func (s StabilityLevel) valid() bool {
	switch s {
	case Alpha, Beta, GA, Deprecated:
		return true
	}
	return false
}

// Set sets the value of this gate.  If this method is called, the provided
// value will be used in subsequent calls to Enabled, regardless of what
// providers might be registered to this gate's GateSet.
func (g *Gate) Set(val bool) {
	g.forcedValue = &val
}

// MergeGateSets combines a list of GateSets into a new GateSet.
func MergeGateSets(gatesets ...*GateSet) *GateSet {
	result := &GateSet{}
	for _, gs := range gatesets {
		result.Merge(gs)
	}
	return result
}

// FIXME: This could be a function here or a method on Gate, but let's not do both?
// Enabled returns whether the specified Gate is enabled.
func Enabled(g *Gate) bool {
	// If some code called Set(), we always respect that.
	if g.forcedValue != nil {
		return *(g.forcedValue)
	}
	// If the value is locked or this gate is not in a set, use the default.
	if g.LockToDefault || g.set == nil {
		return g.Default
	}
	// Otherwise let the set's configured providers have an opinion.
	en, found := g.set.Enabled(g.Name)
	if found {
		return en
	}
	// Otherwise, use the default.
	return g.Default
}

// testHelper is the locally-used subset of testing.T needed in SetForTesting.
type testHelper interface {
	Helper()
}

// FIXME: This could be a function here or a method on Gate, but let's not do both?
// SetForTesting sets the value of the specified Gate and returns a function
// which resets it to the previosu value.
// FIXME: The old mechanism for this has changed, and this needs to be
// re-worked.  See https://github.com/kubernetes/kubernetes/pull/123732
func SetForTesting(g *Gate, t testHelper, enabled bool) func() {
	t.Helper()

	prev := g.forcedValue
	g.forcedValue = &enabled

	return func() {
		g.forcedValue = prev
	}
}
