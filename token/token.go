package token

// トークンの種類
type TokenType string

type Token struct {
	Type TokenType // トークン名
	Literal string // 具体的な値
}

const (
	// 未知なトークン、文字
	ILLEGAL = "ILLEGAL"

	// ファイルの終端
	EOF = "EOF"

	// 識別子
	IDENT = "IDENT"

	// リテラル
	INT = "INT"

	// 演算子
	ASSIGN = "="
	PLUS = "+"

	// デリミタ
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET = "LET"
)