package parser2

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dcaiafa/nitro/internal/parser2/parser"
)

type augmentedLexer struct {
	*parser.NitroLexer

	stashed  antlr.Token
	eolToken bool
}

func newAugmentedLexer(input antlr.CharStream) *augmentedLexer {
	return &augmentedLexer{
		NitroLexer: parser.NewNitroLexer(input),
	}
}

func (l *augmentedLexer) NextToken() antlr.Token {
	var t antlr.Token

	for {
		if l.stashed != nil {
			t = l.stashed
			l.stashed = nil
		} else {
			t = l.NitroLexer.NextToken()
		}

		switch t.GetTokenType() {
		case parser.NitroLexerNEWLINE, antlr.TokenEOF:
			if l.eolToken {
				l.eolToken = false
				l.stashed = t
				t = l.NitroLexer.GetTokenFactory().Create(
					t.GetSource(),
					parser.NitroLexerSEMICOLON,
					"",
					t.GetChannel(),
					t.GetStart(),
					t.GetStop(),
					t.GetLine(),
					t.GetColumn())
			}

		case parser.NitroLexerCPAREN,
			parser.NitroLexerCBRACKET,
			parser.NitroLexerCCURLY,
			parser.NitroLexerEXPAND,
			parser.NitroLexerNUMBER,
			parser.NitroLexerSTRING,
			parser.NitroLexerID,
			parser.NitroLexerTRUE,
			parser.NitroLexerFALSE,
			parser.NitroLexerCATCH,
			parser.NitroLexerREGEX:
			l.eolToken = true

		default:
			l.eolToken = false
		}

		if t.GetTokenType() != parser.NitroLexerNEWLINE {
			break
		}
	}

	return t
}
