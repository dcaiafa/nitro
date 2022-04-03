lexer grammar NitroLexer;
import Strings;

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
