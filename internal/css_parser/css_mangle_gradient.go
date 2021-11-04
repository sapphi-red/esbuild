package css_parser

import (
	"strings"

	"github.com/evanw/esbuild/internal/css_ast"
	"github.com/evanw/esbuild/internal/css_lexer"
)

func (p *parser) mangleGradientFunctions(token css_ast.Token) css_ast.Token {
	switch strings.ToLower(token.Text) {
	case "linear-gradient", "repeating-linear-gradient":
		children := p.mangleLinearGradient(*token.Children)
		token.Children = &children
	case "radial-gradient", "repeating-radial-gradient":
		children := p.mangleRadialGradient(*token.Children)
		token.Children = &children
	case "conic-gradient", "repeating-conic-gradient":
		children := p.mangleConicGradient(*token.Children)
		token.Children = &children
	}

	return token
}

func (p *parser) mangleLinearGradient(tokens []css_ast.Token) []css_ast.Token {
	// Specification: https://drafts.csswg.org/css-images-3/#linear-gradients
	// https://www.w3.org/TR/css-images-4/#linear-gradients
	res := make([]css_ast.Token, 0, len(tokens))
	splittedTokens, ok := getTokensSplittedByComma(tokens)
	if !ok {
		return tokens
	}

	isFirst := true
	for i, t := range splittedTokens {
		isLast := i == len(splittedTokens)-1

		if i == 0 && getColorIndex(t) == -1 {
			// it is not <linear-color-stop>
			t = p.mangleLinearGradientDirection(t)
		} else {
			t = p.mangleLinearColorStopOrHint(t, isFirst, isLast)
			isFirst = false
		}

		if len(t) > 0 {
			if len(res) > 0 {
				res = append(res, p.commaToken())
			}
			t[0].Whitespace &= ^css_ast.WhitespaceBefore
			res = append(res, t...)
		}
	}

	return res
}

func (p *parser) mangleRadialGradient(tokens []css_ast.Token) []css_ast.Token {
	// Specification: https://drafts.csswg.org/css-images-3/#radial-gradients
	// https://www.w3.org/TR/css-images-4/#radial-gradients
	res := make([]css_ast.Token, 0, len(tokens))
	splittedTokens, ok := getTokensSplittedByComma(tokens)
	if !ok {
		return tokens
	}

	isFirst := true
	for i, t := range splittedTokens {
		isLast := i == len(splittedTokens)-1

		if i == 0 && getColorIndex(t) == -1 {
			// it is not <linear-color-stop>
			t = p.mangleRadialGradientDirection(t)
		} else {
			t = p.mangleLinearColorStopOrHint(t, isFirst, isLast)
			isFirst = false
		}

		if len(t) > 0 {
			if len(res) > 0 {
				res = append(res, p.commaToken())
			}
			t[0].Whitespace &= ^css_ast.WhitespaceBefore
			res = append(res, t...)
		}
	}

	return res
}

func (p *parser) mangleConicGradient(tokens []css_ast.Token) []css_ast.Token {
	// Specification: https://drafts.csswg.org/css-images-3/#conic-gradients
	// https://www.w3.org/TR/css-images-4/#conic-gradients
	res := make([]css_ast.Token, 0, len(tokens))
	splittedTokens, ok := getTokensSplittedByComma(tokens)
	if !ok {
		return tokens
	}

	isFirst := true
	for i, t := range splittedTokens {
		isLast := i == len(splittedTokens)-1

		if i == 0 && getColorIndex(t) == -1 {
			// it is not <angular-color-stop>
			t = p.mangleConicGradientDirection(t)
		} else {
			t = p.mangleAngularColorStopOrHint(t, isFirst, isLast)
			isFirst = false
		}

		if len(t) > 0 {
			if len(res) > 0 {
				res = append(res, p.commaToken())
			}
			t[0].Whitespace &= ^css_ast.WhitespaceBefore
			res = append(res, t...)
		}
	}

	return res
}

func getTokensSplittedByComma(tokens []css_ast.Token) ([][]css_ast.Token, bool) {
	result := make([][]css_ast.Token, 0)

	start := 0
	for i := range tokens {
		if tokens[i].Kind == css_lexer.TComma {
			result = append(result, tokens[start:i])
			start = i + 1
			continue
		}

		// var() and env() may include comma
		if tokens[i].Kind == css_lexer.TFunction {
			switch strings.ToLower(tokens[i].Text) {
			case "var", "env":
				return [][]css_ast.Token{}, false
			}
		}
	}
	result = append(result, tokens[start:])
	return result, true
}

// [ <angle> | to <side-or-corner> ]?
func (p *parser) mangleLinearGradientDirection(tokens []css_ast.Token) []css_ast.Token {
	if len(tokens) >= 2 {
		// to <side-or-corner>
		if isSameKeyword(tokens[0], "to") && tokens[1].Kind == css_lexer.TIdent {
			// `to top left` should not be mangled
			if len(tokens) == 2 {
				switch strings.ToLower(tokens[1].Text) {
				case "top": // "to top" to "0deg"
					return []css_ast.Token{
						{Text: "0deg", Kind: css_lexer.TDimension, UnitOffset: 1},
					}

				case "right": // "to right" to "90deg"
					return []css_ast.Token{
						{Text: "90deg", Kind: css_lexer.TDimension, UnitOffset: 2},
					}

				case "bottom": // "to bottom" to ""
					return []css_ast.Token{}

				case "left": // "to left" to "270deg"
					return []css_ast.Token{
						{Text: "270deg", Kind: css_lexer.TDimension, UnitOffset: 3},
					}
				}
			}
		}
	}

	return tokens
}

// [ <ending-shape> || <size> ]? [ at <position> ]?
func (p *parser) mangleRadialGradientDirection(tokens []css_ast.Token) []css_ast.Token {
	pos := 0
	newTokens := make([]css_ast.Token, 0, len(tokens))

	endingShape := ""
	size := make([]css_ast.Token, 0, 2)
	extentKeyword := ""

	// [ <ending-shape> || <size> ]?
	for ; pos < len(tokens); pos++ {
		token := tokens[pos]
		loweredText := strings.ToLower(token.Text)
		if isSameKeyword(token, "at") {
			break
		}

		// <ending-shape>
		if endingShape == "" && token.Kind == css_lexer.TIdent {
			switch loweredText {
			case "circle", "ellipse":
				endingShape = loweredText
				continue
			}
		}

		// <size>
		if extentKeyword == "" && len(size) == 0 {
			if token.Kind == css_lexer.TIdent {
				switch loweredText {
				case "closest-side", "farthest-side", "closest-corner", "farthest-corner":
					extentKeyword = loweredText
					continue
				}
			} else if isLengthTypeOrPercentage(token) {
				size = append(size, token)
				if pos+1 < len(tokens) && isLengthTypeOrPercentage(tokens[pos+1]) {
					size = append(size, tokens[pos+1])
					pos++
				}
				continue
			}
		}

		// unknown token found
		return tokens
	}

	// omit ending-shape
	if len(size) == 1 {
		if endingShape == "circle" {
			endingShape = ""
		}
	} else {
		if endingShape == "ellipse" {
			endingShape = ""
		}
	}

	// omit size
	if extentKeyword == "farthest-corner" {
		extentKeyword = ""
	}

	if len(size) > 0 {
		size[len(size)-1].Whitespace &= ^css_ast.WhitespaceAfter
	}

	newTokens = append(newTokens, size...)
	if extentKeyword != "" {
		var whitespace css_ast.WhitespaceFlags
		if !p.options.RemoveWhitespace && len(newTokens) > 0 {
			whitespace = css_ast.WhitespaceBefore
		}

		newTokens = append(newTokens, css_ast.Token{
			Kind:       css_lexer.TIdent,
			Text:       extentKeyword,
			Whitespace: whitespace,
		})
	}
	if endingShape != "" {
		var whitespace css_ast.WhitespaceFlags
		if !p.options.RemoveWhitespace && len(newTokens) > 0 {
			whitespace = css_ast.WhitespaceBefore
		}

		newTokens = append(newTokens, css_ast.Token{
			Kind:       css_lexer.TIdent,
			Text:       endingShape,
			Whitespace: whitespace,
		})
	}

	if !(pos < len(tokens)) {
		return newTokens
	}

	// at <position>
	atPos := pos
	token := tokens[atPos]
	if atPos+1 < len(tokens) && isSameKeyword(token, "at") {
		if !canOmitPosition(tokens[atPos+1:]) {
			// cf. at left top
			newTokens = append(newTokens, tokens[atPos:]...)
		}
		return newTokens
	}

	// unknown token found
	return tokens
}

// [ from <angle> ]? [ at <position> ]?
func (p *parser) mangleConicGradientDirection(tokens []css_ast.Token) []css_ast.Token {
	newTokens := make([]css_ast.Token, 0, len(tokens))

	pos := 0
	token := tokens[pos]

	if isSameKeyword(token, "from") {
		pos++
		angleToken := tokens[pos]
		pos++
		if IsAngleType(angleToken) {
			// omit if `0deg` or other
			if angleToken.DimensionValue() != "0" {
				newTokens = append(newTokens, token)
				newTokens = append(newTokens, angleToken)
			}
		} else {
			// unknown token found
			return tokens
		}
	}

	if !(pos < len(tokens)) {
		return newTokens
	}

	token = tokens[pos]
	if isSameKeyword(token, "at") {
		// omit if `at center`
		if canOmitPosition(tokens[pos+1:]) {
			if !p.options.RemoveWhitespace && len(newTokens) > 0 {
				newTokens[len(newTokens)-1].Whitespace &= ^css_ast.WhitespaceAfter
			}
		} else {
			newTokens = append(newTokens, tokens[pos:]...)
		}
		return newTokens
	}

	// unknown token found
	return tokens
}

func canOmitPosition(tokens []css_ast.Token) bool {
	firstToken := tokens[0]
	if isCenterOr50Percent(firstToken) {
		// `at center` (not `at 50%`)
		if len(tokens) == 1 && firstToken.Kind == css_lexer.TIdent {
			return true
		}
		// at center center
		if len(tokens) == 2 && isCenterOr50Percent(tokens[1]) {
			return true
		}
	}
	return false
}

func isCenterOr50Percent(token css_ast.Token) bool {
	return isSameKeyword(token, "center") || isEqualPercentage(token, "50")
}

// <linear-color-stop> or <linear-color-hint>
func (p *parser) mangleLinearColorStopOrHint(tokens []css_ast.Token, isFirst bool, isLast bool) []css_ast.Token {
	// Specification: https://drafts.csswg.org/css-images-3/#color-stop-syntax
	// https://www.w3.org/TR/css-images-4/#color-stop-syntax
	if len(tokens) == 1 {
		// <color> inside <linear-color-stop> or <length-percentage> inside <linear-color-hint>
		if hex, ok := parseColor(tokens[0]); ok {
			tokens[0] = p.mangleColor(tokens[0], hex)
		}
		return tokens
	}
	if len(tokens) > 3 {
		return tokens
	}

	colorI := getColorIndex(tokens)
	if colorI == -1 {
		// color should have been found
		return tokens
	}
	if hex, ok := parseColor(tokens[colorI]); ok {
		tokens[colorI] = p.mangleColor(tokens[colorI], hex)
	}

	if len(tokens) == 2 {
		positionToken := tokens[0]
		if colorI == 0 {
			positionToken = tokens[1]
		}

		if isFirst && isEqualPercentage(positionToken, "0") {
			tokens[colorI].Whitespace = 0
			return []css_ast.Token{tokens[colorI]}
		}
		if isLast && isEqualPercentage(positionToken, "100") {
			tokens[colorI].Whitespace = 0
			return []css_ast.Token{tokens[colorI]}
		}
	}

	return tokens
}

// <angular-color-stop> or <angular-color-hint>
func (p *parser) mangleAngularColorStopOrHint(tokens []css_ast.Token, isFirst bool, isLast bool) []css_ast.Token {
	// Specification: https://drafts.csswg.org/css-images-3/#color-stop-syntax
	// https://www.w3.org/TR/css-images-4/#color-stop-syntax
	if len(tokens) == 1 {
		// <color> inside <angular-color-stop> or <angle-percentage> inside <angluar-color-hint>
		if hex, ok := parseColor(tokens[0]); ok {
			tokens[0] = p.mangleColor(tokens[0], hex)
		}
		return tokens
	}
	if len(tokens) > 3 {
		return tokens
	}

	colorI := getColorIndex(tokens)
	if colorI == -1 {
		// color should have been found
		return tokens
	}
	if hex, ok := parseColor(tokens[colorI]); ok {
		tokens[colorI] = p.mangleColor(tokens[colorI], hex)
	}

	if len(tokens) == 2 {
		positionToken := tokens[0]
		if colorI == 0 {
			positionToken = tokens[1]
		}

		if isFirst && (isEqualPercentage(positionToken, "0") || isEqualDegreesAngle(positionToken, 0)) {
			tokens[colorI].Whitespace = 0
			return []css_ast.Token{tokens[colorI]}
		}
		if isLast && (isEqualPercentage(positionToken, "100") || isEqualDegreesAngle(positionToken, 360)) {
			tokens[colorI].Whitespace = 0
			return []css_ast.Token{tokens[colorI]}
		}
	}

	return tokens
}

func getColorIndex(tokens []css_ast.Token) int {
	for i := range tokens {
		if isColorType(tokens[i]) {
			return i
		}
	}
	return -1
}

func isEqualPercentage(token css_ast.Token, val string) bool {
	if token.Kind != css_lexer.TPercentage {
		return false
	}
	return token.PercentageValue() == val
}

func isEqualDegreesAngle(token css_ast.Token, deg float64) bool {
	if token.Kind != css_lexer.TDimension {
		return false
	}

	val, ok := degreesForAngle(token)
	if !ok {
		return false
	}
	return val == deg
}
