#!/bin/sh
rm -rf parser/
antlr -Dlanguage=Go -o parser NitroLexer.g4
antlr -Dlanguage=Go -o parser NitroParser.g4
