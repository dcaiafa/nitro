lexer grammar Strings;

LDQUOTE: '"' -> more, mode(QSTR);

mode QSTR;

QUOTED_TEXT:   ~["\r\n\\]  -> more;
QUOTED_ESCAPE: '\\' ([nrt"\\] | ('x' HEX_DIGIT HEX_DIGIT)) -> more;
fragment HEX_DIGIT: [a-fA-F0-9];

STRING: '"' -> mode(DEFAULT_MODE);

mode DEFAULT_MODE;
