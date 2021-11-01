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
