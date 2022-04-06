package main

import (
	"fmt"

	"github.com/evanw/esbuild/internal/compat"
)

var featuresData = []string{
	"ArbitraryModuleNamespaceNames",
	"ArraySpread",
	"Arrow",
	"AsyncAwait",
	"AsyncGenerator",
	"BigInt",
	"Class",
	"ClassField",
	"ClassPrivateAccessor",
	"ClassPrivateBrandCheck",
	"ClassPrivateField",
	"ClassPrivateMethod",
	"ClassPrivateStaticAccessor",
	"ClassPrivateStaticField",
	"ClassPrivateStaticMethod",
	"ClassStaticBlocks",
	"ClassStaticField",
	"Const",
	"DefaultArgument",
	"Destructuring",
	"DynamicImport",
	"ExponentOperator",
	"ExportStarAs",
	"ForAwait",
	"ForOf",
	"Generator",
	"Hashbang",
	"ImportAssertions",
	"ImportMeta",
	"Let",
	"LogicalAssignment",
	"NestedRestBinding",
	"NewTarget",
	"NodeColonPrefixImport",
	"NodeColonPrefixRequire",
	"NullishCoalescing",
	"ObjectAccessors",
	"ObjectExtensions",
	"ObjectRestSpread",
	"OptionalCatchBinding",
	"OptionalChain",
	"RestArgument",
	"TemplateLiteral",
	"TopLevelAwait",
	"TypeofExoticObjectIsObject",
	"UnicodeEscapes",
}

func main() {
	before := compat.UnsupportedJSFeatures(map[compat.Engine][]int{
		compat.ES:      {2019},
		compat.Edge:    {88},
		compat.Firefox: {78},
		compat.Chrome:  {87},
		compat.Safari:  {13, 1},
	})
	after := compat.UnsupportedJSFeatures(map[compat.Engine][]int{
		compat.ES:      {2020},
		compat.Edge:    {88},
		compat.Firefox: {78},
		compat.Chrome:  {87},
		compat.Safari:  {13, 1},
	})

	fmt.Println("Unsupported features change")
	diff := before ^ after
	for i := 0; i < 64; i++ {
		if ((diff >> i) & 0x1) == 1 {
			fmt.Printf("%s: %t -> %t\n", featuresData[i], before>>i&0x1 == 1, after>>i&0x1 == 1)
		}
	}
}
