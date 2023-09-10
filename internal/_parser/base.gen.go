package parser

type TokenType int

const (
	EOF                 TokenType = 0
	ERROR               TokenType = 1
	INFO                TokenType = 2
	PARAM               TokenType = 3
	FLAG                TokenType = 4
	AND                 TokenType = 5
	BREAK               TokenType = 6
	CATCH               TokenType = 7
	CONTINUE            TokenType = 8
	DEFER               TokenType = 9
	ELSE                TokenType = 10
	FALSE               TokenType = 11
	FOR                 TokenType = 12
	FUNC                TokenType = 13
	IF                  TokenType = 14
	IMPORT              TokenType = 15
	NIL                 TokenType = 16
	NOT                 TokenType = 17
	OR                  TokenType = 18
	RETURN              TokenType = 19
	THROW               TokenType = 20
	TRUE                TokenType = 21
	TRY                 TokenType = 22
	VAR                 TokenType = 23
	WHILE               TokenType = 24
	YIELD               TokenType = 25
	ASSIGN              TokenType = 26
	ASSIGN_ADD          TokenType = 27
	ASSIGN_SUB          TokenType = 28
	ASSIGN_MUL          TokenType = 29
	ASSIGN_DIV          TokenType = 30
	ASSIGN_MOD          TokenType = 31
	EQ                  TokenType = 32
	NE                  TokenType = 33
	LT                  TokenType = 34
	LE                  TokenType = 35
	GT                  TokenType = 36
	GE                  TokenType = 37
	ADD                 TokenType = 38
	SUB                 TokenType = 39
	MUL                 TokenType = 40
	DIV                 TokenType = 41
	MOD                 TokenType = 42
	INC                 TokenType = 43
	DEC                 TokenType = 44
	QUESTION_MARK       TokenType = 45
	SEMICOLON           TokenType = 46
	COMMA               TokenType = 47
	COLON               TokenType = 48
	PERIOD              TokenType = 49
	OPAREN              TokenType = 50
	CPAREN              TokenType = 51
	OBRACKET            TokenType = 52
	CBRACKET            TokenType = 53
	OCURLY              TokenType = 54
	CCURLY              TokenType = 55
	ARROW               TokenType = 56
	LAMBDA              TokenType = 57
	PIPE                TokenType = 58
	EXPAND              TokenType = 59
	BACKTICK            TokenType = 60
	NUMBER              TokenType = 61
	STRING              TokenType = 62
	ID                  TokenType = 63
	REGEX               TokenType = 64
	NEWLINE             TokenType = 65
	CHAR                TokenType = 66
	EXEC_PREFIX         TokenType = 67
	EXEC_LITERAL        TokenType = 68
	EXEC_DQUOTE_LITERAL TokenType = 69
	EXEC_SQUOTE_LITERAL TokenType = 70
	EXEC_WS             TokenType = 71
	EXEC_HOME           TokenType = 72
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "end-of-file"
	case ERROR:
		return "error"
	case INFO:
		return "!info"
	case PARAM:
		return "!param"
	case FLAG:
		return "!flag"
	case AND:
		return "and"
	case BREAK:
		return "break"
	case CATCH:
		return "catch"
	case CONTINUE:
		return "continue"
	case DEFER:
		return "defer"
	case ELSE:
		return "else"
	case FALSE:
		return "false"
	case FOR:
		return "for"
	case FUNC:
		return "func"
	case IF:
		return "if"
	case IMPORT:
		return "import"
	case NIL:
		return "nil"
	case NOT:
		return "not"
	case OR:
		return "or"
	case RETURN:
		return "return"
	case THROW:
		return "throw"
	case TRUE:
		return "true"
	case TRY:
		return "try"
	case VAR:
		return "var"
	case WHILE:
		return "while"
	case YIELD:
		return "yield"
	case ASSIGN:
		return "="
	case ASSIGN_ADD:
		return "+="
	case ASSIGN_SUB:
		return "-="
	case ASSIGN_MUL:
		return "*="
	case ASSIGN_DIV:
		return "/="
	case ASSIGN_MOD:
		return "%="
	case EQ:
		return "=="
	case NE:
		return "!="
	case LT:
		return "<"
	case LE:
		return "<="
	case GT:
		return ">"
	case GE:
		return ">="
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case MOD:
		return "%"
	case INC:
		return "++"
	case DEC:
		return "--"
	case QUESTION_MARK:
		return "?"
	case SEMICOLON:
		return ";"
	case COMMA:
		return ","
	case COLON:
		return ":"
	case PERIOD:
		return "."
	case OPAREN:
		return "("
	case CPAREN:
		return ")"
	case OBRACKET:
		return "["
	case CBRACKET:
		return "]"
	case OCURLY:
		return "{"
	case CCURLY:
		return "}"
	case ARROW:
		return "->"
	case LAMBDA:
		return "&"
	case PIPE:
		return "|"
	case EXPAND:
		return "..."
	case BACKTICK:
		return "`"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case ID:
		return "identifier"
	case REGEX:
		return "regex-literal"
	case NEWLINE:
		return "new-line"
	case CHAR:
		return "character"
	case EXEC_PREFIX:
		return "e`"
	case EXEC_LITERAL:
		return "EXEC_LITERAL"
	case EXEC_DQUOTE_LITERAL:
		return "EXEC_DQUOTE_LITERAL"
	case EXEC_SQUOTE_LITERAL:
		return "EXEC_SQUOTE_LITERAL"
	case EXEC_WS:
		return "EXEC_WS"
	case EXEC_HOME:
		return "EXEC_HOME"
	default:
		return "???"
	}
}

type _Bounds struct {
	Begin Token
	End   Token
}

type lexer interface {
	ReadToken() (Token, TokenType)
}
