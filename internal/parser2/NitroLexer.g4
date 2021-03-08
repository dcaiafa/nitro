lexer grammar NitroLexer;
import Strings;

AND: 'and';
DO: 'do';
ELIF: 'elif';
ELSE: 'else';
END: 'end';
FALSE: 'false';
FN: 'fn';
FOR: 'for';
IF: 'if';
IN: 'in';
META: 'meta';
NOT: 'not';
OR: 'or';
RETURN: 'return';
THEN: 'then';
TRUE: 'true';
VAR: 'var';
WHILE: 'while';

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

NUMBER: [0-9]+ ('.' [0-9]+)?;
ID: [a-zA-Z_] [a-zA-Z0-9_]*;

WS: [ \t]+ -> skip;
NEWLINE: '\r'? '\n';

