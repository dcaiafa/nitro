package parser


type TokenType int

const (
	EOF TokenType = 0
	ERROR TokenType = 1
	INFO TokenType = 2
	PARAM TokenType = 3
	FLAG TokenType = 4
	AND TokenType = 5
	BREAK TokenType = 6
	CATCH TokenType = 7
	CONTINUE TokenType = 8
	DEFER TokenType = 9
	ELSE TokenType = 10
	FALSE TokenType = 11
	FOR TokenType = 12
	FUNC TokenType = 13
	IF TokenType = 14
	IMPORT TokenType = 15
	NIL TokenType = 16
	NOT TokenType = 17
	OR TokenType = 18
	RETURN TokenType = 19
	THROW TokenType = 20
	TRUE TokenType = 21
	TRY TokenType = 22
	VAR TokenType = 23
	WHILE TokenType = 24
	YIELD TokenType = 25
	ASSIGN TokenType = 26
	ASSIGN_ADD TokenType = 27
	ASSIGN_SUB TokenType = 28
	ASSIGN_MUL TokenType = 29
	ASSIGN_DIV TokenType = 30
	ASSIGN_MOD TokenType = 31
	EQ TokenType = 32
	NE TokenType = 33
	LT TokenType = 34
	LE TokenType = 35
	GT TokenType = 36
	GE TokenType = 37
	ADD TokenType = 38
	SUB TokenType = 39
	MUL TokenType = 40
	DIV TokenType = 41
	MOD TokenType = 42
	INC TokenType = 43
	DEC TokenType = 44
	QUESTION_MARK TokenType = 45
	SEMICOLON TokenType = 46
	COMMA TokenType = 47
	COLON TokenType = 48
	PERIOD TokenType = 49
	OPAREN TokenType = 50
	CPAREN TokenType = 51
	OBRACKET TokenType = 52
	CBRACKET TokenType = 53
	OCURLY TokenType = 54
	CCURLY TokenType = 55
	ARROW TokenType = 56
	LAMBDA TokenType = 57
	PIPE TokenType = 58
	EXPAND TokenType = 59
	BACKTICK TokenType = 60
	NUMBER TokenType = 61
	STRING TokenType = 62
	ID TokenType = 63
	REGEX TokenType = 64
	NEWLINE TokenType = 65
	CHAR TokenType = 66
	EXEC_PREFIX TokenType = 67
	EXEC_LITERAL TokenType = 68
	EXEC_DQUOTE_LITERAL TokenType = 69
	EXEC_SQUOTE_LITERAL TokenType = 70
	EXEC_WS TokenType = 71
	EXEC_HOME TokenType = 72
)

func (t TokenType) String() string {
	switch t {
	case EOF: 
		return "EOF"
	case ERROR: 
		return "ERROR"
	case INFO: 
		return "INFO"
	case PARAM: 
		return "PARAM"
	case FLAG: 
		return "FLAG"
	case AND: 
		return "AND"
	case BREAK: 
		return "BREAK"
	case CATCH: 
		return "CATCH"
	case CONTINUE: 
		return "CONTINUE"
	case DEFER: 
		return "DEFER"
	case ELSE: 
		return "ELSE"
	case FALSE: 
		return "FALSE"
	case FOR: 
		return "FOR"
	case FUNC: 
		return "FUNC"
	case IF: 
		return "IF"
	case IMPORT: 
		return "IMPORT"
	case NIL: 
		return "NIL"
	case NOT: 
		return "NOT"
	case OR: 
		return "OR"
	case RETURN: 
		return "RETURN"
	case THROW: 
		return "THROW"
	case TRUE: 
		return "TRUE"
	case TRY: 
		return "TRY"
	case VAR: 
		return "VAR"
	case WHILE: 
		return "WHILE"
	case YIELD: 
		return "YIELD"
	case ASSIGN: 
		return "ASSIGN"
	case ASSIGN_ADD: 
		return "ASSIGN_ADD"
	case ASSIGN_SUB: 
		return "ASSIGN_SUB"
	case ASSIGN_MUL: 
		return "ASSIGN_MUL"
	case ASSIGN_DIV: 
		return "ASSIGN_DIV"
	case ASSIGN_MOD: 
		return "ASSIGN_MOD"
	case EQ: 
		return "EQ"
	case NE: 
		return "NE"
	case LT: 
		return "LT"
	case LE: 
		return "LE"
	case GT: 
		return "GT"
	case GE: 
		return "GE"
	case ADD: 
		return "ADD"
	case SUB: 
		return "SUB"
	case MUL: 
		return "MUL"
	case DIV: 
		return "DIV"
	case MOD: 
		return "MOD"
	case INC: 
		return "INC"
	case DEC: 
		return "DEC"
	case QUESTION_MARK: 
		return "QUESTION_MARK"
	case SEMICOLON: 
		return "SEMICOLON"
	case COMMA: 
		return "COMMA"
	case COLON: 
		return "COLON"
	case PERIOD: 
		return "PERIOD"
	case OPAREN: 
		return "OPAREN"
	case CPAREN: 
		return "CPAREN"
	case OBRACKET: 
		return "OBRACKET"
	case CBRACKET: 
		return "CBRACKET"
	case OCURLY: 
		return "OCURLY"
	case CCURLY: 
		return "CCURLY"
	case ARROW: 
		return "ARROW"
	case LAMBDA: 
		return "LAMBDA"
	case PIPE: 
		return "PIPE"
	case EXPAND: 
		return "EXPAND"
	case BACKTICK: 
		return "BACKTICK"
	case NUMBER: 
		return "NUMBER"
	case STRING: 
		return "STRING"
	case ID: 
		return "ID"
	case REGEX: 
		return "REGEX"
	case NEWLINE: 
		return "NEWLINE"
	case CHAR: 
		return "CHAR"
	case EXEC_PREFIX: 
		return "EXEC_PREFIX"
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
