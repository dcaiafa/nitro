lexer grammar NitroLexer;
import Strings;

AND: 'and';
ELSE: 'else';
END: 'end';
FALSE: 'false';
FUNC: 'func';
FOR: 'for';
IF: 'if';
IN: 'in';
META: 'meta';
NOT: 'not';
OR: 'or';
RETURN: 'return';
TRUE: 'true';
VAR: 'var';
WHILE: 'while';
DEFER: 'defer';
TRY: 'try';
CATCH: 'catch';

ASSIGN: '=';
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

WS: [ \t]+ -> skip;
NEWLINE: '\r'? '\n';

