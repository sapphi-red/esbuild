package css_parser

import (
	"strings"

	"github.com/evanw/esbuild/internal/css_ast"
	"github.com/evanw/esbuild/internal/css_lexer"
)

// https://www.w3.org/TR/css3-values/#keywords
func isSameKeyword(token css_ast.Token, loweredKeyword string) bool {
	return token.Kind == css_lexer.TIdent && strings.ToLower(token.Text) == loweredKeyword
}

var loweredSystemColorNames = map[string]struct{}{
	// <system-color> https://drafts.csswg.org/css-color-4/#typedef-system-color
	"canvas":           {},
	"canvastext":       {},
	"linktext":         {},
	"visitedtext":      {},
	"activetext":       {},
	"buttonface":       {},
	"buttontext":       {},
	"buttonborder":     {},
	"field":            {},
	"fieldtext":        {},
	"highlight":        {},
	"highlighttext":    {},
	"selecteditem":     {},
	"selecteditemtext": {},
	"mark":             {},
	"marktext":         {},
	"graytext":         {},
	// <deprecated-color> https://drafts.csswg.org/css-color-4/#typedef-deprecated-color
	"activeborder":        {},
	"activecaption":       {},
	"appworkspace":        {},
	"background":          {},
	"buttonhighlight":     {},
	"buttonshadow":        {},
	"captiontext":         {},
	"inactiveborder":      {},
	"inactivecaption":     {},
	"inactivecaptiontext": {},
	"infobackground":      {},
	"infotext":            {},
	"menu":                {},
	"menutext":            {},
	"scrollbar":           {},
	"threeddarkshadow":    {},
	"threedface":          {},
	"threedhighlight":     {},
	"threedlightshadow":   {},
	"threedshadow":        {},
	"window":              {},
	"windowframe":         {},
	"windowtext":          {},
}

// Specification: https://drafts.csswg.org/css-color-4/#color-syntax
func isColorType(token css_ast.Token) bool {
	loweredText := strings.ToLower(token.Text)

	switch token.Kind {
	case css_lexer.TIdent:
		if _, ok := colorNameToHex[loweredText]; ok {
			return true
		}
		if _, ok := loweredSystemColorNames[loweredText]; ok {
			return true
		}
		if loweredText == "transparent" {
			return true
		}
		if loweredText == "currentcolor" {
			return true
		}
	case css_lexer.THash:
		return true
	case css_lexer.TFunction:
		switch loweredText {
		case "device-cmyk":
			return true
		case "rgb", "rgba":
			return true
		case "hsl", "hsla":
			return true
		case "hwb":
			return true
		case "lab", "oklab":
			return true
		case "lch", "oklch":
			return true
		case "color":
			return true
		}
	}

	return false
}
