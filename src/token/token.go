package token

type Kind string

type Token struct {
	Kind    Kind
	Literal string
}

const (
	Illegal Kind = "ILLEGAL"
	Eof     Kind = "EOF"

	// Identifiers + Literals
	Identifier Kind = "IDENTIFIER"
	Integer    Kind = "INTEGER"
	Boolean    Kind = "BOOLEAN"

	// Operators
	Assign      Kind = "="
	Plus        Kind = "+"
	Minus       Kind = "-"
	Bang        Kind = "!"
	Asterisk    Kind = "*"
	Slash       Kind = "/"
	LessThan    Kind = "<"
	GreaterThan Kind = ">"
	Equals      Kind = "=="
	NotEquals   Kind = "!="

	// Delimiters
	Comma     Kind = ","
	Semicolon Kind = ";"

	LeftParen  Kind = "("
	RightParen Kind = ")"
	LeftBrace  Kind = "{"
	RightBrace Kind = "}"

	// Keywords
	Function Kind = "FUNCTION"
	Let      Kind = "LET"
	If       Kind = "IF"
	Else     Kind = "ELSE"
	Return   Kind = "RETURN"
)

func NewToken(kind Kind, literal string) Token {
	return Token{Kind: kind, Literal: literal}
}

func NewTokenForIdentifier(literal string) Token {
	switch literal {
	case "fn":
		return NewToken(Function, string(Function))
	case "let":
		return NewToken(Let, string(Let))
	case "true", "false":
		return NewToken(Boolean, literal)
	case "if":
		return NewToken(If, string(If))
	case "else":
		return NewToken(Else, string(Else))
	case "return":
		return NewToken(Return, string(Return))
	default:
		return NewToken(Identifier, literal)
	}
}
