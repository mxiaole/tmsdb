package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 关键字
	SELECT = "SELECT"
	AND    = "AND"
	FROM   = "FROM"
	WHERE  = "WHERE"

	// 运算符
	EQUAL  = "="
	GREATE = ">"
	LESS   = "<"
)
