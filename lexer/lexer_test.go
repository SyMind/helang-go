package lexer

import "testing"

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	t.Helper()
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func expectNumber(t *testing.T, contents string, expected int) {
	t.Helper()
	l := NewLexer(contents)
	assertEqual(t, l.Token, TNumber)
	assertEqual(t, l.Number, expected)
}

func TestNumber(t *testing.T) {
	expectNumber(t, "0", 0)
	expectNumber(t, "123", 123)
}

func expectIdentifier(t *testing.T, contents string, expected string) {
	t.Helper()
	l := NewLexer(contents)
	assertEqual(t, l.Token, TIdentifier)
	assertEqual(t, l.Identifier, expected)
}

func TestIdentifier(t *testing.T) {
	expectIdentifier(t, "a", "a")
}

func expectKeyword(t *testing.T, contents string, token Token) {
	t.Helper()
	l := NewLexer(contents)
	assertEqual(t, l.Token, token)
}

func TestKeyworkd(t *testing.T) {
	expectKeyword(t, "print", TPrint)
	expectKeyword(t, "u8", TU8)
	expectKeyword(t, "test5g", TTest5G)
	expectKeyword(t, "cyberspaces", TCyberSpaces)
	expectKeyword(t, "sprint", TSprint)
}
