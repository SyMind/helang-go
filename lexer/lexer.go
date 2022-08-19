package lexer

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Token uint

const (
	TEndOfFile Token = iota

	// Literal tokens
	TNumber

	// Punctuation
	TBar          // |
	TOpenParen    // (
	TCloseParen   // )
	TOpenBrace    // {
	TCloseBrace   // }
	TOpenBracket  // [
	TCloseBracket // ]
	TComma        // ,
	TColon        // ;
	TMinus        // -
	TPlus         // +
	TPlusPlus     // ++
	TLessThan     // <

	// Assignments
	TEquals

	// Identifiers
	TIdentifier

	// Reserved words
	TPrint
	TU8
	TTest5G
	TCyberSpaces
	TSprint
)

var Keywords = map[string]Token{
	"print":       TPrint,
	"u8":          TU8,
	"test5g":      TTest5G,
	"cyberspaces": TCyberSpaces,
	"sprint":      TSprint,
}

type Loc struct {
	// This is the 0-based index of this location from the start of the file, in bytes
	Start int32
}

type LexerPanic struct {
	Msg string
	Loc Loc
}

func IsWhitespace(codePoint rune) bool {
	switch codePoint {
	case
		'\u0009', // character tabulation
		'\u000B', // line tabulation
		'\u000C', // form feed
		'\u0020', // space
		'\u00A0': // no-break space

		return true

	default:
		return false
	}
}

type Lexer struct {
	source     string
	current    int
	start      int
	end        int
	codePoint  rune
	Token      Token
	Number     int
	Identifier string
}

func NewLexer(source string) Lexer {
	lexer := Lexer{
		source: source,
	}
	lexer.step()
	lexer.Next()
	return lexer
}

func (l *Lexer) Next() {
	for {
		l.start = l.end
		l.Token = 0

		switch l.codePoint {
		case -1:
			l.Token = TEndOfFile

		case '\r', '\n', '\u2028', '\u2029':
			l.step()

		case '[':
			l.step()
			l.Token = TOpenBracket

		case ']':
			l.step()
			l.Token = TCloseBracket

		case '{':
			l.step()
			l.Token = TOpenBrace

		case '}':
			l.step()
			l.Token = TCloseBrace

		case ',':
			l.step()
			l.Token = TComma

		case ':':
			l.step()
			l.Token = TColon

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.parseNumber()

		case '_', '$',
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
			'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
			'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			l.step()
			l.parseIdentifier()

		default:
			// Check for insignificant whitespace characters
			if IsWhitespace(l.codePoint) {
				l.step()
				continue
			}

			l.SyntaxError(fmt.Sprintf("Unexpected token %c", l.source[l.end]))
		}

		return
	}
}

func (l *Lexer) parseNumber() {
	first := l.codePoint
	l.step()
	if first == '0' && l.codePoint == '0' {
		l.SyntaxError("Unexpected number")
	}

	for {
		if l.codePoint < '0' || l.codePoint > '9' {
			break
		}
		l.step()
	}

	l.Token = TNumber
	number := l.source[l.start:l.end]
	if value, err := strconv.ParseInt(number, 10, 32); err == nil {
		l.Number = int(value)
	}
}

func (l *Lexer) parseIdentifier() {
	for {
		l.step()
		if IsWhitespace(l.codePoint) || l.codePoint == -1 {
			l.step()
			break
		}
	}

	text := l.source[l.start:l.end]
	if Keywords[text] != 0 {
		l.Token = Keywords[text]
	} else {
		l.Token = TIdentifier
	}
	l.Identifier = text

}

func (l *Lexer) SyntaxError(msg string) {
	loc := Loc{Start: int32(l.end)}
	panic(LexerPanic{
		Msg: msg,
		Loc: loc,
	})
}

func (l *Lexer) step() {
	codePoint, width := utf8.DecodeRuneInString(l.source[l.current:])

	// 使用 -1 表示文件的结尾
	if width == 0 {
		codePoint = -1
	}

	l.codePoint = codePoint
	l.end = l.current
	l.current += width
}
