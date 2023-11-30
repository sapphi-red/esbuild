// This file was automatically generated by "js_table.ts"

package compat

type Engine uint8

const (
	Chrome Engine = iota
	Deno
	Edge
	ES
	Firefox
	Hermes
	IE
	IOS
	Node
	Opera
	Rhino
	Safari
)

func (e Engine) String() string {
	switch e {
	case Chrome:
		return "chrome"
	case Deno:
		return "deno"
	case Edge:
		return "edge"
	case ES:
		return "es"
	case Firefox:
		return "firefox"
	case Hermes:
		return "hermes"
	case IE:
		return "ie"
	case IOS:
		return "ios"
	case Node:
		return "node"
	case Opera:
		return "opera"
	case Rhino:
		return "rhino"
	case Safari:
		return "safari"
	}
	return ""
}

func (e Engine) IsBrowser() bool {
	switch e {
	case Chrome, Edge, Firefox, IE, IOS, Opera, Safari:
		return true
	}
	return false
}

type JSFeature uint64

const (
	ArbitraryModuleNamespaceNames JSFeature = 1 << iota
	ArraySpread
	Arrow
	AsyncAwait
	AsyncGenerator
	Bigint
	Class
	ClassField
	ClassPrivateAccessor
	ClassPrivateBrandCheck
	ClassPrivateField
	ClassPrivateMethod
	ClassPrivateStaticAccessor
	ClassPrivateStaticField
	ClassPrivateStaticMethod
	ClassStaticBlocks
	ClassStaticField
	ConstAndLet
	Decorators
	DefaultArgument
	Destructuring
	DynamicImport
	ExponentOperator
	ExportStarAs
	ForAwait
	ForOf
	FunctionNameConfigurable
	FunctionOrClassPropertyAccess
	Generator
	Hashbang
	ImportAssertions
	ImportAttributes
	ImportMeta
	InlineScript
	LogicalAssignment
	NestedRestBinding
	NewTarget
	NodeColonPrefixImport
	NodeColonPrefixRequire
	NullishCoalescing
	ObjectAccessors
	ObjectExtensions
	ObjectRestSpread
	OptionalCatchBinding
	OptionalChain
	RegexpDotAllFlag
	RegexpLookbehindAssertions
	RegexpMatchIndices
	RegexpNamedCaptureGroups
	RegexpSetNotation
	RegexpStickyAndUnicodeFlags
	RegexpUnicodePropertyEscapes
	RestArgument
	TemplateLiteral
	TopLevelAwait
	TypeofExoticObjectIsObject
	UnicodeEscapes
	Using
)

var StringToJSFeature = map[string]JSFeature{
	"arbitrary-module-namespace-names":  ArbitraryModuleNamespaceNames,
	"array-spread":                      ArraySpread,
	"arrow":                             Arrow,
	"async-await":                       AsyncAwait,
	"async-generator":                   AsyncGenerator,
	"bigint":                            Bigint,
	"class":                             Class,
	"class-field":                       ClassField,
	"class-private-accessor":            ClassPrivateAccessor,
	"class-private-brand-check":         ClassPrivateBrandCheck,
	"class-private-field":               ClassPrivateField,
	"class-private-method":              ClassPrivateMethod,
	"class-private-static-accessor":     ClassPrivateStaticAccessor,
	"class-private-static-field":        ClassPrivateStaticField,
	"class-private-static-method":       ClassPrivateStaticMethod,
	"class-static-blocks":               ClassStaticBlocks,
	"class-static-field":                ClassStaticField,
	"const-and-let":                     ConstAndLet,
	"decorators":                        Decorators,
	"default-argument":                  DefaultArgument,
	"destructuring":                     Destructuring,
	"dynamic-import":                    DynamicImport,
	"exponent-operator":                 ExponentOperator,
	"export-star-as":                    ExportStarAs,
	"for-await":                         ForAwait,
	"for-of":                            ForOf,
	"function-name-configurable":        FunctionNameConfigurable,
	"function-or-class-property-access": FunctionOrClassPropertyAccess,
	"generator":                         Generator,
	"hashbang":                          Hashbang,
	"import-assertions":                 ImportAssertions,
	"import-attributes":                 ImportAttributes,
	"import-meta":                       ImportMeta,
	"inline-script":                     InlineScript,
	"logical-assignment":                LogicalAssignment,
	"nested-rest-binding":               NestedRestBinding,
	"new-target":                        NewTarget,
	"node-colon-prefix-import":          NodeColonPrefixImport,
	"node-colon-prefix-require":         NodeColonPrefixRequire,
	"nullish-coalescing":                NullishCoalescing,
	"object-accessors":                  ObjectAccessors,
	"object-extensions":                 ObjectExtensions,
	"object-rest-spread":                ObjectRestSpread,
	"optional-catch-binding":            OptionalCatchBinding,
	"optional-chain":                    OptionalChain,
	"regexp-dot-all-flag":               RegexpDotAllFlag,
	"regexp-lookbehind-assertions":      RegexpLookbehindAssertions,
	"regexp-match-indices":              RegexpMatchIndices,
	"regexp-named-capture-groups":       RegexpNamedCaptureGroups,
	"regexp-set-notation":               RegexpSetNotation,
	"regexp-sticky-and-unicode-flags":   RegexpStickyAndUnicodeFlags,
	"regexp-unicode-property-escapes":   RegexpUnicodePropertyEscapes,
	"rest-argument":                     RestArgument,
	"template-literal":                  TemplateLiteral,
	"top-level-await":                   TopLevelAwait,
	"typeof-exotic-object-is-object":    TypeofExoticObjectIsObject,
	"unicode-escapes":                   UnicodeEscapes,
	"using":                             Using,
}

func (features JSFeature) Has(feature JSFeature) bool {
	return (features & feature) != 0
}

func (features JSFeature) ApplyOverrides(overrides JSFeature, mask JSFeature) JSFeature {
	return (features & ^mask) | (overrides & mask)
}

var jsTable = map[JSFeature]map[Engine][]versionRange{
	ArbitraryModuleNamespaceNames: {
		Chrome:  {{start: v{90, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{87, 0, 0}}},
		IOS:     {{start: v{14, 5, 0}}},
		Node:    {{start: v{16, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	ArraySpread: {
		// Note: The latest version of "IE" failed 15 tests including: spread syntax for iterable objects: spreading non-iterables is a runtime error
		// Note: The latest version of "Rhino" failed 15 tests including: spread syntax for iterable objects: spreading non-iterables is a runtime error
		Chrome:  {{start: v{46, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{36, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{5, 0, 0}}},
		Opera:   {{start: v{33, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	Arrow: {
		// Note: The latest version of "Hermes" failed 3 tests including: arrow functions: lexical "super" binding in constructors
		// Note: The latest version of "IE" failed 13 tests including: arrow functions: "this" unchanged by call or apply
		// Note: The latest version of "Rhino" failed 3 tests including: arrow functions: lexical "new.target" binding
		Chrome:  {{start: v{49, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{45, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	AsyncAwait: {
		// Note: The latest version of "Hermes" failed 4 tests including: async functions: async arrow functions
		// Note: The latest version of "IE" failed 16 tests including: async functions: async arrow functions
		// Note: The latest version of "Rhino" failed 16 tests including: async functions: async arrow functions
		Chrome:  {{start: v{55, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{15, 0, 0}}},
		ES:      {{start: v{2017, 0, 0}}},
		Firefox: {{start: v{52, 0, 0}}},
		IOS:     {{start: v{11, 0, 0}}},
		Node:    {{start: v{7, 6, 0}}},
		Opera:   {{start: v{42, 0, 0}}},
		Safari:  {{start: v{11, 0, 0}}},
	},
	AsyncGenerator: {
		// Note: The latest version of "Hermes" failed this test: Asynchronous Iterators: async generators
		// Note: The latest version of "IE" failed this test: Asynchronous Iterators: async generators
		// Note: The latest version of "Rhino" failed this test: Asynchronous Iterators: async generators
		Chrome:  {{start: v{63, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{57, 0, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{50, 0, 0}}},
		Safari:  {{start: v{12, 0, 0}}},
	},
	Bigint: {
		// Note: The latest version of "IE" failed this test: BigInt: basic functionality
		Chrome:  {{start: v{67, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{68, 0, 0}}},
		Hermes:  {{start: v{0, 12, 0}}},
		IOS:     {{start: v{14, 0, 0}}},
		Node:    {{start: v{10, 4, 0}}},
		Opera:   {{start: v{54, 0, 0}}},
		Rhino:   {{start: v{1, 7, 14}}},
		Safari:  {{start: v{14, 0, 0}}},
	},
	Class: {
		// Note: The latest version of "Hermes" failed 24 tests including: class: accessor properties
		// Note: The latest version of "IE" failed 24 tests including: class: accessor properties
		// Note: The latest version of "Rhino" failed 24 tests including: class: accessor properties
		Chrome:  {{start: v{49, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{45, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	ClassField: {
		// Note: The latest version of "Hermes" failed 2 tests including: instance class fields: computed instance class fields
		// Note: The latest version of "IE" failed 2 tests including: instance class fields: computed instance class fields
		// Note: The latest version of "Rhino" failed 2 tests including: instance class fields: computed instance class fields
		Chrome:  {{start: v{73, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{69, 0, 0}}},
		IOS:     {{start: v{14, 0, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{60, 0, 0}}},
		Safari:  {{start: v{14, 0, 0}}},
	},
	ClassPrivateAccessor: {
		// Note: The latest version of "Hermes" failed this test: private class methods: private accessor properties
		// Note: The latest version of "IE" failed this test: private class methods: private accessor properties
		// Note: The latest version of "Rhino" failed this test: private class methods: private accessor properties
		Chrome:  {{start: v{84, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateBrandCheck: {
		// Note: The latest version of "Hermes" failed this test: Ergonomic brand checks for private fields
		// Note: The latest version of "IE" failed this test: Ergonomic brand checks for private fields
		// Note: The latest version of "Rhino" failed this test: Ergonomic brand checks for private fields
		Chrome:  {{start: v{91, 0, 0}}},
		Deno:    {{start: v{1, 9, 0}}},
		Edge:    {{start: v{91, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{16, 4, 0}}},
		Opera:   {{start: v{77, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateField: {
		// Note: The latest version of "Hermes" failed 4 tests including: instance class fields: optional deep private instance class fields access
		// Note: The latest version of "IE" failed 4 tests including: instance class fields: optional deep private instance class fields access
		// Note: The latest version of "Rhino" failed 4 tests including: instance class fields: optional deep private instance class fields access
		Chrome:  {{start: v{84, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{14, 5, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	ClassPrivateMethod: {
		// Note: The latest version of "Hermes" failed this test: private class methods: private instance methods
		// Note: The latest version of "IE" failed this test: private class methods: private instance methods
		// Note: The latest version of "Rhino" failed this test: private class methods: private instance methods
		Chrome:  {{start: v{84, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateStaticAccessor: {
		// Note: The latest version of "Hermes" failed this test: private class methods: private static accessor properties
		// Note: The latest version of "IE" failed this test: private class methods: private static accessor properties
		// Note: The latest version of "Rhino" failed this test: private class methods: private static accessor properties
		Chrome:  {{start: v{84, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateStaticField: {
		// Note: The latest version of "Hermes" failed this test: static class fields: private static class fields
		// Note: The latest version of "IE" failed this test: static class fields: private static class fields
		// Note: The latest version of "Rhino" failed this test: static class fields: private static class fields
		Chrome:  {{start: v{74, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{14, 5, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{62, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	ClassPrivateStaticMethod: {
		// Note: The latest version of "Hermes" failed this test: private class methods: private static methods
		// Note: The latest version of "IE" failed this test: private class methods: private static methods
		// Note: The latest version of "Rhino" failed this test: private class methods: private static methods
		Chrome:  {{start: v{84, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassStaticBlocks: {
		Chrome:  {{start: v{91, 0, 0}}},
		Deno:    {{start: v{1, 14, 0}}},
		Edge:    {{start: v{94, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{93, 0, 0}}},
		IOS:     {{start: v{16, 4, 0}}},
		Node:    {{start: v{16, 11, 0}}},
		Opera:   {{start: v{80, 0, 0}}},
		Safari:  {{start: v{16, 4, 0}}},
	},
	ClassStaticField: {
		// Note: The latest version of "Hermes" failed 2 tests including: static class fields: computed static class fields
		// Note: The latest version of "IE" failed 2 tests including: static class fields: computed static class fields
		// Note: The latest version of "Rhino" failed 2 tests including: static class fields: computed static class fields
		Chrome:  {{start: v{73, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{75, 0, 0}}},
		IOS:     {{start: v{14, 5, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{60, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	ConstAndLet: {
		// Note: The latest version of "Hermes" failed 20 tests including: const: for loop statement scope
		// Note: The latest version of "IE" failed 6 tests including: const: for-in loop iteration scope
		// Note: The latest version of "Rhino" failed 22 tests including: const: cannot be in statements
		Chrome:  {{start: v{49, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{51, 0, 0}}},
		IOS:     {{start: v{11, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{11, 0, 0}}},
	},
	Decorators: {},
	DefaultArgument: {
		// Note: The latest version of "Hermes" failed 2 tests including: default function parameters: separate scope
		// Note: The latest version of "IE" failed 7 tests including: default function parameters: arguments object interaction
		// Note: The latest version of "Rhino" failed 7 tests including: default function parameters: arguments object interaction
		Chrome:  {{start: v{49, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	Destructuring: {
		// Note: The latest version of "Hermes" failed 3 tests including: destructuring, declarations: defaults, let temporal dead zone
		// Note: The latest version of "IE" failed 71 tests including: destructuring, assignment: chained iterable destructuring
		// Note: The latest version of "Rhino" failed 33 tests including: destructuring, assignment: computed properties
		Chrome:  {{start: v{51, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{18, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 5, 0}}},
		Opera:   {{start: v{38, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	DynamicImport: {
		Chrome:  {{start: v{63, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{67, 0, 0}}},
		IOS:     {{start: v{11, 0, 0}}},
		Node:    {{start: v{12, 20, 0}, end: v{13, 0, 0}}, {start: v{13, 2, 0}}},
		Opera:   {{start: v{50, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	ExponentOperator: {
		// Note: The latest version of "IE" failed 3 tests including: exponentiation (**) operator: assignment
		Chrome:  {{start: v{52, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2016, 0, 0}}},
		Firefox: {{start: v{52, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 3, 0}}},
		Node:    {{start: v{7, 0, 0}}},
		Opera:   {{start: v{39, 0, 0}}},
		Rhino:   {{start: v{1, 7, 14}}},
		Safari:  {{start: v{10, 1, 0}}},
	},
	ExportStarAs: {
		Chrome:  {{start: v{72, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{80, 0, 0}}},
		IOS:     {{start: v{14, 5, 0}}},
		Node:    {{start: v{13, 2, 0}}},
		Opera:   {{start: v{60, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	ForAwait: {
		// Note: The latest version of "Hermes" failed this test: Asynchronous Iterators: for-await-of loops
		// Note: The latest version of "IE" failed this test: Asynchronous Iterators: for-await-of loops
		// Note: The latest version of "Rhino" failed this test: Asynchronous Iterators: for-await-of loops
		Chrome:  {{start: v{63, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{57, 0, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{50, 0, 0}}},
		Safari:  {{start: v{12, 0, 0}}},
	},
	ForOf: {
		// Note: The latest version of "IE" failed 9 tests including: for..of loops: iterator closing, break
		// Note: The latest version of "Rhino" failed 4 tests including: for..of loops: iterator closing, break
		Chrome:  {{start: v{51, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{15, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 5, 0}}},
		Opera:   {{start: v{38, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	FunctionNameConfigurable: {
		// Note: The latest version of "IE" failed this test: function "name" property: isn't writable, is configurable
		// Note: The latest version of "Rhino" failed this test: function "name" property: isn't writable, is configurable
		Chrome:  {{start: v{43, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{38, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{4, 0, 0}}},
		Opera:   {{start: v{30, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	FunctionOrClassPropertyAccess: {
		Chrome:  {{start: v{0, 0, 0}}},
		Deno:    {{start: v{0, 0, 0}}},
		Edge:    {{start: v{0, 0, 0}}},
		ES:      {{start: v{0, 0, 0}}},
		Firefox: {{start: v{0, 0, 0}}},
		Hermes:  {{start: v{0, 0, 0}}},
		IE:      {{start: v{0, 0, 0}}},
		IOS:     {{start: v{0, 0, 0}}},
		Node:    {{start: v{0, 0, 0}}},
		Opera:   {{start: v{0, 0, 0}}},
		Rhino:   {{start: v{0, 0, 0}}},
		Safari:  {{start: v{16, 3, 0}}},
	},
	Generator: {
		// Note: The latest version of "Hermes" failed 3 tests including: generators: computed shorthand generators, classes
		// Note: The latest version of "IE" failed 27 tests including: generators: %GeneratorPrototype%
		// Note: The latest version of "Rhino" failed 15 tests including: generators: %GeneratorPrototype%
		Chrome:  {{start: v{50, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{37, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	Hashbang: {
		// Note: The latest version of "IE" failed this test: Hashbang Grammar
		// Note: The latest version of "Rhino" failed this test: Hashbang Grammar
		Chrome:  {{start: v{74, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		Firefox: {{start: v{67, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{13, 4, 0}}},
		Node:    {{start: v{12, 5, 0}}},
		Opera:   {{start: v{62, 0, 0}}},
		Safari:  {{start: v{13, 1, 0}}},
	},
	ImportAssertions: {
		Chrome: {{start: v{91, 0, 0}}},
		Deno:   {{start: v{1, 17, 0}}},
		Edge:   {{start: v{91, 0, 0}}},
		Node:   {{start: v{16, 14, 0}}},
	},
	ImportAttributes: {
		Deno: {{start: v{1, 37, 0}}},
		Node: {{start: v{21, 0, 0}}},
	},
	ImportMeta: {
		Chrome:  {{start: v{64, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{62, 0, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{10, 4, 0}}},
		Opera:   {{start: v{51, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	InlineScript: {},
	LogicalAssignment: {
		// Note: The latest version of "IE" failed 9 tests including: Logical Assignment: &&= basic support
		// Note: The latest version of "Rhino" failed 9 tests including: Logical Assignment: &&= basic support
		Chrome:  {{start: v{85, 0, 0}}},
		Deno:    {{start: v{1, 2, 0}}},
		Edge:    {{start: v{85, 0, 0}}},
		ES:      {{start: v{2021, 0, 0}}},
		Firefox: {{start: v{79, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{14, 0, 0}}},
		Node:    {{start: v{15, 0, 0}}},
		Opera:   {{start: v{71, 0, 0}}},
		Safari:  {{start: v{14, 0, 0}}},
	},
	NestedRestBinding: {
		// Note: The latest version of "IE" failed 2 tests including: nested rest destructuring, declarations
		// Note: The latest version of "Rhino" failed 2 tests including: nested rest destructuring, declarations
		Chrome:  {{start: v{49, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2016, 0, 0}}},
		Firefox: {{start: v{47, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 3, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 1, 0}}},
	},
	NewTarget: {
		// Note: The latest version of "IE" failed 2 tests including: new.target: assignment is an early error
		// Note: The latest version of "Rhino" failed 2 tests including: new.target: assignment is an early error
		Chrome:  {{start: v{46, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{41, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{5, 0, 0}}},
		Opera:   {{start: v{33, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	NodeColonPrefixImport: {
		Node: {{start: v{12, 20, 0}, end: v{13, 0, 0}}, {start: v{14, 13, 1}}},
	},
	NodeColonPrefixRequire: {
		Node: {{start: v{14, 18, 0}, end: v{15, 0, 0}}, {start: v{16, 0, 0}}},
	},
	NullishCoalescing: {
		// Note: The latest version of "IE" failed this test: nullish coalescing operator (??)
		// Note: The latest version of "Rhino" failed this test: nullish coalescing operator (??)
		Chrome:  {{start: v{80, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{80, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{72, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{13, 4, 0}}},
		Node:    {{start: v{14, 0, 0}}},
		Opera:   {{start: v{67, 0, 0}}},
		Safari:  {{start: v{13, 1, 0}}},
	},
	ObjectAccessors: {
		Chrome:  {{start: v{5, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{5, 0, 0}}},
		Firefox: {{start: v{2, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IE:      {{start: v{9, 0, 0}}},
		IOS:     {{start: v{6, 0, 0}}},
		Node:    {{start: v{0, 4, 0}}},
		Opera:   {{start: v{10, 10, 0}}},
		Rhino:   {{start: v{1, 7, 13}}},
		Safari:  {{start: v{3, 1, 0}}},
	},
	ObjectExtensions: {
		// Note: The latest version of "IE" failed 6 tests including: object literal extensions: computed accessors
		// Note: The latest version of "Rhino" failed 3 tests including: object literal extensions: computed accessors
		Chrome:  {{start: v{44, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{34, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{4, 0, 0}}},
		Opera:   {{start: v{31, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	ObjectRestSpread: {
		// Note: The latest version of "IE" failed 2 tests including: object rest/spread properties: object rest properties
		// Note: The latest version of "Rhino" failed 2 tests including: object rest/spread properties: object rest properties
		Chrome:  {{start: v{60, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{55, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Node:    {{start: v{8, 3, 0}}},
		Opera:   {{start: v{47, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	OptionalCatchBinding: {
		// Note: The latest version of "IE" failed 3 tests including: optional catch binding: await
		// Note: The latest version of "Rhino" failed 3 tests including: optional catch binding: await
		Chrome:  {{start: v{66, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2019, 0, 0}}},
		Firefox: {{start: v{58, 0, 0}}},
		Hermes:  {{start: v{0, 12, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{53, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	OptionalChain: {
		// Note: The latest version of "IE" failed 5 tests including: optional chaining operator (?.): optional bracket access
		// Note: The latest version of "Rhino" failed 5 tests including: optional chaining operator (?.): optional bracket access
		Chrome:  {{start: v{91, 0, 0}}},
		Deno:    {{start: v{1, 9, 0}}},
		Edge:    {{start: v{91, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{74, 0, 0}}},
		Hermes:  {{start: v{0, 12, 0}}},
		IOS:     {{start: v{13, 4, 0}}},
		Node:    {{start: v{16, 1, 0}}},
		Opera:   {{start: v{77, 0, 0}}},
		Safari:  {{start: v{13, 1, 0}}},
	},
	RegexpDotAllFlag: {
		// Note: The latest version of "IE" failed this test: s (dotAll) flag for regular expressions
		// Note: The latest version of "Rhino" failed this test: s (dotAll) flag for regular expressions
		Chrome:  {{start: v{62, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{78, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Node:    {{start: v{8, 10, 0}}},
		Opera:   {{start: v{49, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	RegexpLookbehindAssertions: {
		// Note: The latest version of "IE" failed this test: RegExp Lookbehind Assertions
		// Note: The latest version of "Rhino" failed this test: RegExp Lookbehind Assertions
		Chrome:  {{start: v{62, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{78, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{16, 4, 0}}},
		Node:    {{start: v{8, 10, 0}}},
		Opera:   {{start: v{49, 0, 0}}},
		Safari:  {{start: v{16, 4, 0}}},
	},
	RegexpMatchIndices: {
		Chrome:  {{start: v{90, 0, 0}}},
		Deno:    {{start: v{1, 8, 0}}},
		Edge:    {{start: v{90, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{88, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{16, 0, 0}}},
		Opera:   {{start: v{76, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	RegexpNamedCaptureGroups: {
		// Note: The latest version of "Hermes" failed this test: RegExp named capture groups
		// Note: The latest version of "IE" failed this test: RegExp named capture groups
		// Note: The latest version of "Rhino" failed this test: RegExp named capture groups
		Chrome:  {{start: v{64, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{78, 0, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{51, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	RegexpSetNotation: {},
	RegexpStickyAndUnicodeFlags: {
		// Note: The latest version of "IE" failed 6 tests including: RegExp "y" and "u" flags: "u" flag
		// Note: The latest version of "Rhino" failed 6 tests including: RegExp "y" and "u" flags: "u" flag
		Chrome:  {{start: v{50, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{46, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{37, 0, 0}}},
		Safari:  {{start: v{12, 0, 0}}},
	},
	RegexpUnicodePropertyEscapes: {
		// Note: The latest version of "Chrome" failed this test: RegExp Unicode Property Escapes: Unicode 15.1
		// Note: The latest version of "Firefox" failed this test: RegExp Unicode Property Escapes: Unicode 15.1
		// Note: The latest version of "Hermes" failed 8 tests including: RegExp Unicode Property Escapes: Unicode 11
		// Note: The latest version of "IE" failed 8 tests including: RegExp Unicode Property Escapes: Unicode 11
		// Note: The latest version of "IOS" failed this test: RegExp Unicode Property Escapes: Unicode 15.1
		// Note: The latest version of "Node" failed this test: RegExp Unicode Property Escapes: Unicode 15.1
		// Note: The latest version of "Rhino" failed 8 tests including: RegExp Unicode Property Escapes: Unicode 11
		// Note: The latest version of "Safari" failed this test: RegExp Unicode Property Escapes: Unicode 15.1
		Deno:  {{start: v{1, 31, 0}}},
		Edge:  {{start: v{110, 0, 0}}},
		ES:    {{start: v{2018, 0, 0}}},
		Opera: {{start: v{96, 0, 0}}},
	},
	RestArgument: {
		// Note: The latest version of "Hermes" failed this test: rest parameters: function 'length' property
		// Note: The latest version of "IE" failed 5 tests including: rest parameters: arguments object interaction
		// Note: The latest version of "Rhino" failed 5 tests including: rest parameters: arguments object interaction
		Chrome:  {{start: v{47, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{43, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{34, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	TemplateLiteral: {
		// Note: The latest version of "Hermes" failed this test: template literals: TemplateStrings call site caching
		// Note: The latest version of "IE" failed 7 tests including: template literals: TemplateStrings call site caching
		// Note: The latest version of "Rhino" failed 2 tests including: template literals: basic functionality
		Chrome:  {{start: v{41, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{34, 0, 0}}},
		IOS:     {{start: v{13, 0, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{28, 0, 0}}},
		Safari:  {{start: v{13, 0, 0}}},
	},
	TopLevelAwait: {
		Chrome:  {{start: v{89, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{89, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{89, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 8, 0}}},
		Opera:   {{start: v{75, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	TypeofExoticObjectIsObject: {
		Chrome:  {{start: v{0, 0, 0}}},
		Deno:    {{start: v{0, 0, 0}}},
		Edge:    {{start: v{0, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{0, 0, 0}}},
		Hermes:  {{start: v{0, 0, 0}}},
		IOS:     {{start: v{0, 0, 0}}},
		Node:    {{start: v{0, 0, 0}}},
		Opera:   {{start: v{0, 0, 0}}},
		Rhino:   {{start: v{0, 0, 0}}},
		Safari:  {{start: v{0, 0, 0}}},
	},
	UnicodeEscapes: {
		// Note: The latest version of "IE" failed 2 tests including: Unicode code point escapes: in identifiers
		// Note: The latest version of "Rhino" failed 4 tests including: Unicode code point escapes: in identifiers
		Chrome:  {{start: v{44, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{9, 0, 0}}},
		Node:    {{start: v{4, 0, 0}}},
		Opera:   {{start: v{31, 0, 0}}},
		Safari:  {{start: v{9, 0, 0}}},
	},
	Using: {},
}

// Return all features that are not available in at least one environment
func UnsupportedJSFeatures(constraints map[Engine]Semver) (unsupported JSFeature) {
	for feature, engines := range jsTable {
		if feature == InlineScript {
			continue // This is purely user-specified
		}
		for engine, version := range constraints {
			if versionRanges, ok := engines[engine]; !ok || !isVersionSupported(versionRanges, version) {
				unsupported |= feature
			}
		}
	}
	return
}
