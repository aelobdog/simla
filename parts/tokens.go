package parts

// a custom type to store the tokens' type
type TokenType uint8 

// listing out the token types as a const block
const (
    Integer         TokenType = iota
    Real
    Character
    String
    Boolean
    Lparen
    Rparen
    Lsquare
    Rsquare
    Lcurl
    Rcurl
    Plus
    Minus
    Star
    Fwrdsl
    Percent
    Equal
    Nequal
    Grthan
    Grequal
    Lsthan
    Lsequal
    Assign
    Comma
    Colon
    Dot
    Logand
    Logor
    Lognot
    Array
    Function
    Struct
    If
    Then
    Else
    End
    Loop
    Print
    Return
    True
    False
    EOF
    EOL
    Illegal
    Identifier
)

func (t *TokenType) StringFor() string {
    switch(*t) {
        case Integer: return "Integer"
        case Real: return "Real"
        case Character: return "Character"
        case String: return "String"
        case Boolean: return "Boolean"
        case Lparen: return "Lparen"
        case Rparen: return "Rparen"
        case Lsquare: return "Lsquare"
        case Rsquare: return "Rsquare"
        case Lcurl: return "Lcurl"
        case Rcurl: return "Rcurl"
        case Plus: return "Plus"
        case Minus: return "Minus"
        case Star: return "Star"
        case Fwrdsl: return "Fwrdsl"
        case Percent: return "Percent"
        case Equal: return "Equal"
        case Nequal: return "Nequal"
        case Grthan: return "Grthan"
        case Grequal: return "Grequal"
        case Lsthan: return "Lsthan"
        case Lsequal: return "Lsequal"
        case Assign: return "Assign"
        case Comma: return "Comma"
        case Colon: return "Colon"
        case Dot: return "Dot"
        case Logand: return "Logand"
        case Logor: return "Logor"
        case Lognot: return "Lognot"
        case Array: return "Array"
        case Function: return "Function"
        case Struct: return "Struct"
        case If: return "If"
        case Then: return "Then"
        case Else: return "Else"
        case End: return "End"
        case Loop: return "Loop"
        case Print: return "Print"
        case Return: return "Return"
        case True: return "true"
        case False: return "false"
        case EOF: return "EOF"
        case EOL: return "EOL"
        case Identifier: return "Identifier"
    }
    return "Illegal"
}

var Keywords map[string]TokenType = map[string]TokenType {
    "if": If,
    "then": Then,
    "end": End,
    "loop": Loop,
    "array": Array,
    "func": Function,
    "struct": Struct,
    "string": String,
    "int": Integer,
    "real": Real,
    "char": Character,
    "bool": Boolean,
    "print": Print,
    "return": Return,
    "true": True,
    "false": False,
}

type Token struct {
    Type TokenType
    Lexeme string
    Line int
}

func NewToken(tokenType TokenType, lexeme string, line int) Token {
    return Token{Type: tokenType, Lexeme: lexeme, Line: line}
}

func IsDigit(char byte) bool {
    return char >= '0' && char <= '9'
}

func IsLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char == '_')
}

func IsKeyword(word string) bool {
	_, ok := Keywords[word];
    return ok
}
