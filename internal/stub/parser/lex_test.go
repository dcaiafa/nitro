package parser

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/token"
	"github.com/stretchr/testify/require"
)

func TestLex(t *testing.T) {
	run := func(name string, input string, res ...interface{}) {
		t.Run(name, func(t *testing.T) {
			l := newLex([]byte(input))

			for {
				var v yySymType
				tk := l.Lex(&v)

				if len(res) == 0 {
					require.Equal(t, tk, 0)
					break
				}

				require.Equal(t, res[0].(int), tk)

				if tk == LEXERR {
					break
				}

				switch tk {
				case INT:
					require.Equal(t, token.Int, v.tok.Type)
					require.Equal(t, res[1].(int64), v.tok.Int)
				case FLOAT:
					require.Equal(t, token.Float, v.tok.Type)
					require.Equal(t, res[1].(float64), v.tok.Float)
				case STRING, ID:
					require.Equal(t, res[1].(string), v.tok.Str)
				default:
				}

				res = res[2:]
			}
		})
	}

	run("number0", "1234567890", INT, int64(1234567890))
	run("number1", "1", INT, int64(1))
	run("number2", "3.14", FLOAT, float64(3.14))
	run("stringEmpty", `""`, STRING, "")
	run("string0", `"abcd*&fooo"`, STRING, "abcd*&fooo")
	run("stringEscape", `"\"\\\""`, STRING, `"\"`)
	run("id", `foobar1._barFoo`, ID, "foobar1", int('.'), 0, ID, "_barFoo")
	run("mix", `123[foobar`, INT, int64(123), int('['), 0, ID, "foobar")
	run("keywords", "const func type", kCONST, 0, kFUNC, 0, kTYPE, 0)
}
