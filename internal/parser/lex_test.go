package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLex(t *testing.T) {
	run := func(name string, input string, res ...interface{}) {
		t.Run(name, func(t *testing.T) {
			l := newLex(input)

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
				case NUMBER:
					require.Equal(t, res[1].(float64), v.num)
				case STRING, ID:
					require.Equal(t, res[1].(string), v.str)
				default:
				}

				res = res[2:]
			}
		})
	}

	run("logic", `> >= < <= == !=`,
		int('>'), 0, GE, 0, int('<'), 0, LE, 0, EQ, 0, NE, 0)
	run("logic_keywords", `and or not`,
		kAND, "and", kOR, "or", kNOT, "not")
	run("number0", "1234567890", NUMBER, float64(1234567890))
	run("number1", "1", NUMBER, float64(1))
	run("number2", "3.14", NUMBER, float64(3.14))
	run("numberErr1", ".14", LEXERR)
	run("numberErr2", "3.14.15", NUMBER, float64(3.14), LEXERR)
	run("stringEmpty", `""`, STRING, "")
	run("string0", `"abcd*&fooo"`, STRING, "abcd*&fooo")
	run("stringEscape", `"\"\\\""`, STRING, `"\"`)
	run("true_false", "true false", kTRUE, "true", kFALSE, "false")
	run("id", `foobar1+_barFoo`, ID, "foobar1", int('+'), 0, ID, "_barFoo")
	run("mix", `123+foobar`, NUMBER, float64(123), int('+'), 0, ID, "foobar")
	run("in", "seg in [ONE, TWO]", ID, "seg", kIN, "in", int('['), 0, ID, "ONE", int(','), 0, ID, "TWO", int(']'), 0)
}
