package validation_test

import (
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/cel/openapi/resolver"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	"k8s.io/kubernetes/pkg/generated/openapi"
	apivalidationtesting "k8s.io/kubernetes/test/utils/apivalidation"
	"k8s.io/utils/ptr"

	// Ensure everything installed in schema
	_ "k8s.io/kubernetes/pkg/apis/core/install"
)

// FIXME: Create vs Update
func TestValidateConfigMapWithFramework(t *testing.T) {
	newConfigMap := func(name string, tweaks ...func(cm *core.ConfigMap)) *core.ConfigMap {
		cm := &core.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: "default",
			},
			Data:       map[string]string{},
			BinaryData: map[string][]byte{},
			Immutable:  ptr.To(false),
		}
		for _, t := range tweaks {
			t(cm)
		}
		return cm
	}

	addData := func(data map[string]string) func(cm *core.ConfigMap) {
		return func(cm *core.ConfigMap) {
			for k, v := range data {
				cm.Data[k] = v
			}
		}
	}

	type options struct {
		// ConfigMap has no validation options
	}

	//FIXME: test name validation
	cases := []apivalidationtesting.TestCase[*core.ConfigMap, options]{
		//
		// Test data
		//
		{
			Name:   "No data",
			Object: newConfigMap("foo"),
		}, {
			Name:   "One valid data item",
			Object: newConfigMap("foo", addData(map[string]string{"key": "val"})),
		}, {
			Name: "Many valid data items",
			Object: newConfigMap("foo", addData(map[string]string{
				"key":                    "val", // alphabetic
				"123":                    "val", // numeric
				"abc123":                 "val", // alphanum
				"123abc":                 "val", // numalpha
				"-":                      "val", // dash
				"---":                    "val", // dashes
				"-a1":                    "val", // dash-alpha-num
				"---a1":                  "val", // dashes-alpha-num
				"a-1":                    "val", // alpha-dash-num
				"a---1":                  "val", // alpha-dashes-num
				"a1-":                    "val", // alpha-num-dash
				"a1---":                  "val", // alpha-num-dashes
				".a1":                    "val", // dot-alpha-num
				"a.1":                    "val", // alpha-dot-num
				"a...1":                  "val", // alpha-dots-num
				"a1.":                    "val", // alpha-num-dot
				"a1...":                  "val", // alpha-num-dots
				"_":                      "val", // under
				"___":                    "val", // unders
				"_a1":                    "val", // under-alpha-num
				"___a1":                  "val", // under-alpha-num
				"a_1":                    "val", // alpha-under-num
				"a___1":                  "val", // alpha-unders-num
				"a1_":                    "val", // alpha-num-under
				"a1___":                  "val", // alpha-num-unders
				strings.Repeat("x", 253): "val", // long
			})),
		}, {
			Name:   "Empty key",
			Object: newConfigMap("foo", addData(map[string]string{"": "val"})),
			ExpectedErrors: apivalidationtesting.ExpectedErrorList{{
				Field:    "data", //FIXME: should be `data[""]`
				Type:     field.ErrorTypeInvalid,
				BadValue: "",
				Detail:   `key must not be empty`,
			}},
		}, {
			Name:   "Too-long key",
			Object: newConfigMap("foo", addData(map[string]string{strings.Repeat("x", 254): "val"})),
			ExpectedErrors: apivalidationtesting.ExpectedErrorList{{
				Field:    "data", //FIXME: should be `data[""]`
				Type:     field.ErrorTypeInvalid,
				BadValue: "",
				Detail:   `key must not be more than 253 characters`,
			}},
		}, {
			Name: "Invalid keys",
			Object: newConfigMap("foo", addData(map[string]string{
				"!@#$":  "val",
				"%^&*":  "val",
				"()[]":  "val",
				"{}<>":  "val",
				" \t\n": "val",
			})),
			ExpectedErrors: apivalidationtesting.ExpectedErrorList{{
				Field:    "data", //FIXME: should be `data[""]`
				Type:     field.ErrorTypeInvalid,
				BadValue: "",
				Detail:   `must consist of alphanumeric characters, '-', '_', or '.'`,
			}},
			// TODO: .
			// TODO: ..
			// TODO: ..*
		}}

	defs := resolver.NewDefinitionsSchemaResolver(openapi.GetOpenAPIDefinitions, legacyscheme.Scheme)
	apivalidationtesting.TestValidate(t, legacyscheme.Scheme, defs, func(obj *core.ConfigMap, _ options) field.ErrorList {
		return validation.ValidateConfigMap(obj)
	}, cases...)
}

/*
type topologyPair struct {
	key   string
	value string
}

func newHostPathType(pathType string) *core.HostPathType {
	hostPathType := new(core.HostPathType)
	*hostPathType = core.HostPathType(pathType)
	return hostPathType
}

func testVolume(name string, namespace string, spec core.PersistentVolumeSpec) *core.PersistentVolume {
	objMeta := metav1.ObjectMeta{Name: name}
	if namespace != "" {
		objMeta.Namespace = namespace
	}

	return &core.PersistentVolume{
		ObjectMeta: objMeta,
		Spec:       spec,
	}
}

func testVolumeWithNodeAffinity(affinity *core.VolumeNodeAffinity) *core.PersistentVolume {
	return testVolume("test-affinity-volume", "",
		core.PersistentVolumeSpec{
			Capacity: core.ResourceList{
				core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
			},
			AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
			PersistentVolumeSource: core.PersistentVolumeSource{
				GCEPersistentDisk: &core.GCEPersistentDiskVolumeSource{
					PDName: "foo",
				},
			},
			StorageClassName: "test-storage-class",
			NodeAffinity:     affinity,
		})
}

func simpleVolumeNodeAffinity(key, value string) *core.VolumeNodeAffinity {
	return &core.VolumeNodeAffinity{
		Required: &core.NodeSelector{
			NodeSelectorTerms: []core.NodeSelectorTerm{{
				MatchExpressions: []core.NodeSelectorRequirement{{
					Key:      key,
					Operator: core.NodeSelectorOpIn,
					Values:   []string{value},
				}},
			}},
		},
	}
}

func multipleVolumeNodeAffinity(terms [][]topologyPair) *core.VolumeNodeAffinity {
	nodeSelectorTerms := []core.NodeSelectorTerm{}
	for _, term := range terms {
		matchExpressions := []core.NodeSelectorRequirement{}
		for _, topology := range term {
			matchExpressions = append(matchExpressions, core.NodeSelectorRequirement{
				Key:      topology.key,
				Operator: core.NodeSelectorOpIn,
				Values:   []string{topology.value},
			})
		}
		nodeSelectorTerms = append(nodeSelectorTerms, core.NodeSelectorTerm{
			MatchExpressions: matchExpressions,
		})
	}

	return &core.VolumeNodeAffinity{
		Required: &core.NodeSelector{
			NodeSelectorTerms: nodeSelectorTerms,
		},
	}
}
*/
