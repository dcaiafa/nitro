lexer grammar NitroLexer;

M_INFO: '!info';
M_PARAM: '!param';
M_FLAG: '!flag';

AND: 'and';
BREAK: 'break';
CATCH: 'catch';
CONTINUE: 'continue';
DEFER: 'defer';
ELSE: 'else';
FALSE: 'false';
FOR: 'for';
FUNC: 'func';
IF: 'if';
IMPORT: 'import';
NIL: 'nil';
NOT: 'not';
OR: 'or';
RETURN: 'return';
THROW: 'throw';
TRUE: 'true';
TRY: 'try';
VAR: 'var';
WHILE: 'while';
YIELD: 'yield';

ASSIGN: '=';
ASSIGN_ADD: '+=';
ASSIGN_SUB: '-=';
ASSIGN_MUL: '*=';
ASSIGN_DIV: '/=';
ASSIGN_MOD: '%=';

EQ: '==';
NE: '!=';
LT: '<';
LE: '<=';
GT: '>';
GE: '>=';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
MOD: '%';

INC: '++';
DEC: '--';

QUESTION_MARK: '?';
SEMICOLON: ';';
COMMA: ',';
COLON: ':';
PERIOD: '.';
OPAREN: '(';
CPAREN: ')';
OBRACKET: '[';
CBRACKET: ']';
OCURLY: '{';
CCURLY: '}';
ARROW: '->';
LAMBDA: '&';
PIPE: '|';
EXPAND: '...';

NUMBER: [0-9]+ ('.' [0-9]+)?;
ID: [a-zA-Z_] [a-zA-Z0-9_]*;
REGEX: 'r`' ~[`]* '`';
NEWLINE: '\r'? '\n';
CHAR: '\'' (('\\' [nrt'\\]) | ~['\r\n\\]+) '\'';
fragment HEX_DIGIT: [a-fA-F0-9];

COMMENT: '#' ~[\n]* -> skip;
WS: [ \t]+ -> skip;

// Quoted string

DQUOTE_: '"' -> more, pushMode(QSTR);

mode QSTR;

QUOTED_TEXT: ~["\r\n\\]  -> more;
QUOTED_ESCAPE: '\\' ([nrt"\\] | ('x' HEX_DIGIT HEX_DIGIT)) -> more;

STRING: '"' -> popMode;

mode DEFAULT_MODE;

// Exec

EXEC_PREFIX: 'e`' -> pushMode(EXEC);

mode EXEC;

EXEC_WS: [ \t\r\n]+;
EXEC_SUFFIX: '`' -> popMode;
EXEC_LITERAL: ~[ \t\r\n{"'`]+;
EXEC_DQUOTE_: '"' -> more, pushMode(EXEC_DQUOTE);
EXEC_SQUOTE_: '\'' -> more, pushMode(EXEC_SQUOTE);
EXEC_EXPR_PREFIX_: '{' -> type(OCURLY), pushMode(EXEC_EXPR);

mode EXEC_DQUOTE;

EXEC_DQUOTE_TEXT_: ~["\r\n\\]  -> more;
EXEC_DQUOTE_ESCAPE_: '\\' ([nrt"\\] | ('x' HEX_DIGIT HEX_DIGIT)) -> more;
EXEC_DQUOTE_LITERAL: '"' -> popMode;

mode EXEC_SQUOTE;

EXEC_SQUOTE_TEXT_: ~['\r\n\\]  -> more;
EXEC_SQUOTE_ESCAPE_: '\\' ([nrt'\\] | ('x' HEX_DIGIT HEX_DIGIT)) -> more;
EXEC_SQUOTE_LITERAL: '\'' -> popMode;

mode EXEC_EXPR;

EXEC_EXPR_SUFFIX_: '}' -> type(CCURLY), popMode;
EXEC_EXPR_EXEC_SUFFIX_: 'e`' -> type(EXEC_PREFIX), pushMode(EXEC);

EXEC_EXPR_AND_: 'and' -> type(AND);
EXEC_EXPR_FALSE_: 'false' -> type(FALSE);
EXEC_EXPR_NIL_: 'nil' -> type(NIL);
EXEC_EXPR_NOT_: 'not' -> type(NOT);
EXEC_EXPR_OR_: 'or' -> type(OR);
EXEC_EXPR_TRUE_: 'true' -> type(TRUE);

EXEC_EXPR_EQ_: '==' -> type(EQ);
EXEC_EXPR_NE_: '!=' -> type(NE);
EXEC_EXPR_LT_: '<' -> type(LT);
EXEC_EXPR_LE_: '<=' -> type(LE);
EXEC_EXPR_GT_: '>' -> type(GT);
EXEC_EXPR_GE_: '>=' -> type(GE);
EXEC_EXPR_ADD_: '+' -> type(ADD);
EXEC_EXPR_SUB_: '-' -> type(SUB);
EXEC_EXPR_MUL_: '*' -> type(MUL);
EXEC_EXPR_DIV_: '/' -> type(DIV);
EXEC_EXPR_MOD_: '%' -> type(MOD);

EXEC_EXPR_QUESTION_MARK_: '?' -> type(QUESTION_MARK);
EXEC_EXPR_COMMA_: ',' -> type(COMMA);
EXEC_EXPR_COLON_: ':' -> type(COLON);
EXEC_EXPR_PERIOD_: '.' -> type(PERIOD);
EXEC_EXPR_OPAREN_: '(' -> type(OPAREN);
EXEC_EXPR_CPAREN_: ')' -> type(CPAREN);
EXEC_EXPR_OBRACKET_: '[' -> type(OBRACKET);
EXEC_EXPR_CBRACKET_: ']' -> type(CBRACKET);
EXEC_EXPR_ARROW_: '->' -> type(ARROW);
EXEC_EXPR_LAMBDA_: '&' -> type(LAMBDA);
EXEC_EXPR_PIPE_: '|' -> type(PIPE);
EXEC_EXPR_EXPAND_: '...' -> type(EXPAND);

EXEC_EXPR_CHAR_: '\'' (('\\' [nrt'\\]) | ~['\r\n\\]+) '\'' -> type(CHAR);
EXEC_EXPR_NUMBER_: [0-9]+ ('.' [0-9]+)? -> type(NUMBER);
EXEC_EXPR_DQUOTE_: '"' -> more, pushMode(QSTR);
EXEC_EXPR_ID_: [a-zA-Z_] [a-zA-Z0-9_]* -> type(ID);
EXEC_EXPR_WS_: [ \t]+ -> skip;
