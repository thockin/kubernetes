package validation_test

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/cel/openapi/resolver"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	svctst "k8s.io/kubernetes/pkg/api/service/testing"
	"k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	"k8s.io/kubernetes/pkg/generated/openapi"
	apivalidationtesting "k8s.io/kubernetes/test/utils/apivalidation"

	// Ensure everything installed in schema
	_ "k8s.io/kubernetes/pkg/apis/core/install"
)

// FIXME: Create vs Update
func TestValidateServiceWithFramework(t *testing.T) {
	type options struct {
		// Service has no validation options
	}

	// FIXME: large numbers get converted to float: "9.344444e+06"
	// FIXME: error message says "in body": "spec.ports[0].port in body"
	// FIXME: error message says "should": spec.ports[0].port in body should be less than or equal to 65535"
	// FIXME: can't test ports by name yet
	cases := []apivalidationtesting.TestCase[*core.Service, options]{
		//
		// Test ports
		//
		{
			Name: "ClusterIP, one TCP port, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(svctst.MakeServicePort("", 93, intstr.FromInt32(76), core.ProtocolTCP))),
		}, {
			Name: "ClusterIP, one UDP port, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(svctst.MakeServicePort("", 93, intstr.FromInt32(76), core.ProtocolUDP))),
		}, {
			Name: "ClusterIP, one SCTP port, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(svctst.MakeServicePort("", 93, intstr.FromInt32(76), core.ProtocolSCTP))),
		}, {
			Name: "ClusterIP, one empty-protocol port, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(svctst.MakeServicePort("", 93, intstr.FromInt32(76), ""))),
			ExpectedErrors: apivalidationtesting.ExpectedErrorList{{
				Field: "spec.ports[0].protocol",
				Type:  field.ErrorTypeRequired,
			}},
		}, {
			Name: "ClusterIP, one invalid-protocol port, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(svctst.MakeServicePort("", 93, intstr.FromInt32(76), "not-a-valid-protocol"))),
			ExpectedErrors: apivalidationtesting.ExpectedErrorList{{
				Field:    "spec.ports[0].protocol",
				Type:     field.ErrorTypeNotSupported,
				BadValue: "not-a-valid-protocol",
			}},
		}, {
			Name: "ClusterIP, one named TCP port, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(svctst.MakeServicePort("p", 93, intstr.FromInt32(76), core.ProtocolTCP))),
		}, {
			Name: "ClusterIP, two named TCP ports, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(
					svctst.MakeServicePort("p", 93, intstr.FromInt32(76), core.ProtocolTCP),
					svctst.MakeServicePort("q", 76, intstr.FromInt32(93), core.ProtocolTCP))),
		}, {
			Name: "ClusterIP, two unnamed TCP ports, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(
					svctst.MakeServicePort("", 93, intstr.FromInt32(76), core.ProtocolTCP),
					svctst.MakeServicePort("", 76, intstr.FromInt32(93), core.ProtocolTCP))),
			ExpectedErrors: apivalidationtesting.ExpectedErrorList{{
				Field: "spec.ports[0].name",
				Type:  field.ErrorTypeRequired,
			}, {
				Field: "spec.ports[1].name",
				Type:  field.ErrorTypeRequired,
			}},
		}, {
			Name: "ClusterIP, one named, one unnamed TCP ports, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(
					svctst.MakeServicePort("p", 93, intstr.FromInt32(76), core.ProtocolTCP),
					svctst.MakeServicePort("", 76, intstr.FromInt32(93), core.ProtocolTCP))),
			ExpectedErrors: apivalidationtesting.ExpectedErrorList{{
				Field: "spec.ports[1].name",
				Type:  field.ErrorTypeRequired,
			}},
		}, {
			Name: "ClusterIP, duplicate named TCP ports, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(
					svctst.MakeServicePort("p", 93, intstr.FromInt32(76), core.ProtocolTCP),
					svctst.MakeServicePort("p", 76, intstr.FromInt32(93), core.ProtocolTCP))),
			ExpectedErrors: apivalidationtesting.ExpectedErrorList{{
				Field: "spec.ports[1].name",
				Type:  field.ErrorTypeRequired,
			}},
		}, {
			Name: "ClusterIP, invalidly named TCP ports, by number",
			Object: svctst.MakeService("foo",
				svctst.SetTypeClusterIP,
				svctst.SetPorts(
					svctst.MakeServicePort("  ", 93, intstr.FromInt32(76), core.ProtocolTCP),
					svctst.MakeServicePort(",-%!", 76, intstr.FromInt32(93), core.ProtocolTCP))),
			ExpectedErrors: apivalidationtesting.ExpectedErrorList{{
				Field: "spec.ports[1].name",
				Type:  field.ErrorTypeInvalid,
			}},
			// too long port name
			// dup port numbers

			/*
				{
					Name: "good-volume",
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "good-volume-with-capacity-unit",
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10Gi"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "good-volume-without-capacity-unit",
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "good-volume-with-storage-class",
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
						StorageClassName: "valid",
					}),
				},
				{
					Name: "good-volume-with-retain-policy",
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
						PersistentVolumeReclaimPolicy: core.PersistentVolumeReclaimRetain,
					}),
				},
				{
					Name: "good-volume-with-volume-mode",
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
						VolumeMode: &validMode,
					}),
				},
				{
					Name: "invalid-accessmode",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.accessModes", Type: field.ErrorTypeNotSupported},
					},
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{"fakemode"},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "invalid-reclaimpolicy",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.persistentVolumeReclaimPolicy", Type: field.ErrorTypeNotSupported},
					},
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
						PersistentVolumeReclaimPolicy: "fakeReclaimPolicy",
					}),
				},
				{
					Name: "invalid-volume-mode",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.volumeMode", Detail: `supported values: "Block", "Filesystem"`, Type: field.ErrorTypeNotSupported},
					},
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
						VolumeMode: &invalidMode,
					}),
				},
				{
					Name: "with-read-write-once-pod-feature-gate-enabled",
					Options: options{
						enableReadWriteOncePod: true,
					},
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{"ReadWriteOncePod"},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "with-read-write-once-pod-and-others-feature-gate-enabled",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.accessModes", Detail: "may not use ReadWriteOncePod with other access modes", Type: field.ErrorTypeForbidden},
					},
					Options: options{
						enableReadWriteOncePod: true,
					},
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{"ReadWriteOncePod", "ReadWriteMany"},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "unexpected-namespace",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "metadata.namespace", BadValue: "unexpected-namespace", Detail: "not allowed on this type", Type: field.ErrorTypeForbidden},
					},
					Object: testVolume("foo", "unexpected-namespace", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "missing-volume-source",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec", Detail: "must specify a volume type", Type: field.ErrorTypeRequired},
					},
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
					}),
				},
				{
					Name: "bad-name",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "metadata.name", BadValue: "123*Bad(Name", Type: field.ErrorTypeInvalid},
						{Field: "metadata.namespace", BadValue: "unexpected-namespace", Detail: "not allowed on this type", Type: field.ErrorTypeForbidden},
					},
					Object: testVolume("123*Bad(Name", "unexpected-namespace", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "missing-name",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "metadata.name", Detail: "name or generateName is required", Type: field.ErrorTypeRequired},
					},
					Object: testVolume("", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "missing-capacity",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.capacity", Type: field.ErrorTypeRequired},
						{Field: "spec.capacity", Type: field.ErrorTypeNotSupported, SchemaType: field.ErrorTypeInvalid},
					},
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "bad-volume-zero-capacity",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.capacity[storage]", Detail: "must be greater than zero", Type: field.ErrorTypeInvalid},
					},
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("0"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "missing-accessmodes",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.accessModes", Type: field.ErrorTypeRequired},
						{Field: "metadata.namespace", Detail: "not allowed on this type", Type: field.ErrorTypeForbidden},
					},
					Object: testVolume("goodname", "missing-accessmodes", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
					}),
				},
				{
					Name: "too-many-sources",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.accessModes", Type: field.ErrorTypeRequired},
						{Field: "spec.gcePersistentDisk", Detail: "may not specify more than 1 volume type", Type: field.ErrorTypeForbidden},
					},
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("5G"),
						},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
							GCEPersistentDisk: &core.GCEPersistentDiskVolumeSource{PDName: "foo", FSType: "ext4"},
						},
					}),
				},
				{
					Name: "host mount of / with recycle reclaim policy",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.persistentVolumeReclaimPolicy", Detail: `may not be 'recycle' for a hostPath mount of '/'`, Type: field.ErrorTypeForbidden},
					},
					Object: testVolume("bad-recycle-do-not-want", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
						PersistentVolumeReclaimPolicy: core.PersistentVolumeReclaimRecycle,
					}),
				},
				{
					Name: "host mount of / with recycle reclaim policy 2",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.hostPath.path", Detail: `must not contain '..'`, Type: field.ErrorTypeInvalid},
						{Field: "spec.persistentVolumeReclaimPolicy", Detail: `may not be 'recycle' for a hostPath mount of '/'`, Type: field.ErrorTypeForbidden},
					},
					Object: testVolume("bad-recycle-do-not-want", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/a/..",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
						PersistentVolumeReclaimPolicy: core.PersistentVolumeReclaimRecycle,
					}),
				},
				{
					Name: "invalid-storage-class-name",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.storageClassName", BadValue: "-invalid-", Type: field.ErrorTypeInvalid},
					},
					Object: testVolume("invalid-storage-class-name", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
						StorageClassName: "-invalid-",
					}),
				},
				{
					Name: "bad-hostpath-volume-backsteps",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.hostPath.path", Type: field.ErrorTypeInvalid},
					},
					Object: testVolume("foo", "", core.PersistentVolumeSpec{
						Capacity: core.ResourceList{
							core.ResourceName(core.ResourceStorage): resource.MustParse("10G"),
						},
						AccessModes: []core.PersistentVolumeAccessMode{core.ReadWriteOnce},
						PersistentVolumeSource: core.PersistentVolumeSource{
							HostPath: &core.HostPathVolumeSource{
								Path: "/foo/..",
								Type: newHostPathType(string(core.HostPathDirectory)),
							},
						},
						StorageClassName: "backstep-hostpath",
					}),
				},
				{
					Name:   "volume-node-affinity",
					Object: testVolumeWithNodeAffinity(simpleVolumeNodeAffinity("foo", "bar")),
				},
				{
					Name: "volume-empty-node-affinity",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.nodeAffinity.required", Type: field.ErrorTypeRequired},
					},
					Object: testVolumeWithNodeAffinity(&core.VolumeNodeAffinity{}),
				},
				{
					Name: "volume-bad-node-affinity",
					ExpectedErrors: apivalidationtesting.ExpectedErrorList{
						{Field: "spec.nodeAffinity.required.nodeSelectorTerms[0].matchExpressions[0].key", Detail: "name part must be non-empty", Type: field.ErrorTypeInvalid},
						{Field: "spec.nodeAffinity.required.nodeSelectorTerms[0].matchExpressions[0].key", Detail: "name part must consist of alphanumeric characters,", Type: field.ErrorTypeInvalid},
					},
					Object: testVolumeWithNodeAffinity(
						&core.VolumeNodeAffinity{
							Required: &core.NodeSelector{
								NodeSelectorTerms: []core.NodeSelectorTerm{{
									MatchExpressions: []core.NodeSelectorRequirement{{
										Operator: core.NodeSelectorOpIn,
										Values:   []string{"test-label-value"},
									}},
								}},
							},
						}),
				},
			*/
		}}

	defs := resolver.NewDefinitionsSchemaResolver(openapi.GetOpenAPIDefinitions, legacyscheme.Scheme)
	apivalidationtesting.TestValidate(t, legacyscheme.Scheme, defs, func(obj *core.Service, _ options) field.ErrorList {
		return validation.ValidateService(obj)
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
