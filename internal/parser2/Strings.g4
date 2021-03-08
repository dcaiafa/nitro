lexer grammar Strings;

LQUOTE: '"' -> more, mode(QSTR);

mode QSTR;

QUOTED_TEXT:   ~["\r\n\\]  -> more;
QUOTED_ESCAPE: '\\'[\\"rnt] -> more;

STRING: '"' -> mode(DEFAULT_MODE);
