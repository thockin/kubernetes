package namesdb

// RecordList is the type represented in each names.yaml file.
type RecordList []Record

// Record describes a single reserved name.
type Record struct {
	// Key is the full name being claimed, including the prefix.
	// Required
	Key string

	// SIG specifies which SIG owns this name.  Multiple SIGs might use a name,
	// but only one can own it.
	// Required
	SIG SIG

	// Description provides human-friendy information about what this name
	// means and how/when to use it.
	// Required
	Description string

	// DocsLink provides an HTTP link to more information about this name.
	// Optional
	DocsLink string

	// Context specifies where this name might be used.  Names which can be
	// used in more than one place should define multiple name records.
	// Subsequent fields' requiredness depend on the context.
	Context Context

	// ValueType specifies what "inner" type(s) of value are expected for
	// contexts which require a value.  For example, annotation values are
	// always strings, but the content of those strings might be expected to be
	// integers, booleans, or even JSON-encoded structs.  Some names accept
	// multiple types.
	// Required if context demands
	ValueType []ValueType

	// Resources specifies which API resource(s) this name might be used with,
	// in "apigroup/resource" syntax.
	// Required if context demands
	Resources []string
}

// SIG is a Kubernetes special interest group name.
type SIG string

const (
	SIGNetwork      = SIG("Network")
	SIGNode         = SIG("Node")
	SIGApps         = SIG("Apps")
	SIGStorage      = SIG("Storage")
	SIGArchitecture = SIG("Architecture")
	// more...
)

// Context describes where a name might be used.
type Context string

const (
	ContextLabel       = Context("Label")       // key+value
	ContextAnnotation  = Context("Annotation")  // key+value
	ContextTaint       = Context("Taint")       // key
	ContextCondition   = Context("Condition")   // key
	ContextAppProtocol = Context("AppProtocol") // key
	ContextAPIGroup    = Context("APIGroup")    // key
	// more...
)

// ValueType describes the datatype of a value for a key name.
type ValueType string

const (
	ValueTypeInt        = ValueType("Int")
	ValueTypeString     = ValueType("String")
	ValueTypeBool       = ValueType("Bool")
	ValueTypeJSONObject = ValueType("JSONObject")
	ValueTypeJSONList   = ValueType("JSONList")
	// more...
)
