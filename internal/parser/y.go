// Code generated by goyacc parser.y. DO NOT EDIT.

//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2

import ()

//line parser.y:9
type yySymType struct {
	yys  int
	num  float64
	str  string
	ast  interface{}
	expr interface{}
}

const LEXERR = 57346
const kVAR = 57347
const kAND = 57348
const kDO = 57349
const kELIF = 57350
const kELSE = 57351
const kEND = 57352
const kFALSE = 57353
const kFOR = 57354
const kIF = 57355
const kIN = 57356
const kNOT = 57357
const kOR = 57358
const kTHEN = 57359
const kTRUE = 57360
const NUMBER = 57361
const STRING = 57362
const ID = 57363
const OR = 57364
const AND = 57365
const LE = 57366
const GE = 57367
const EQ = 57368
const NE = 57369

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEXERR",
	"kVAR",
	"kAND",
	"kDO",
	"kELIF",
	"kELSE",
	"kEND",
	"kFALSE",
	"kFOR",
	"kIF",
	"kIN",
	"kNOT",
	"kOR",
	"kTHEN",
	"kTRUE",
	"NUMBER",
	"STRING",
	"ID",
	"OR",
	"AND",
	"'<'",
	"LE",
	"'>'",
	"GE",
	"EQ",
	"NE",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"';'",
	"','",
	"'='",
	"'['",
	"']'",
	"'.'",
	"'('",
	"')'",
	"'{'",
	"'}'",
	"':'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 73,
	24, 0,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	-2, 31,
	-1, 74,
	24, 0,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	-2, 32,
	-1, 75,
	24, 0,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	-2, 33,
	-1, 76,
	24, 0,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	-2, 34,
	-1, 77,
	24, 0,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	-2, 35,
	-1, 78,
	24, 0,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	-2, 36,
}

const yyPrivate = 57344

const yyLast = 507

var yyAct = [...]int{

	7, 60, 54, 135, 130, 117, 62, 90, 56, 120,
	2, 98, 94, 87, 46, 35, 36, 37, 38, 39,
	40, 31, 32, 33, 34, 89, 53, 57, 97, 96,
	67, 68, 69, 70, 71, 72, 73, 74, 75, 76,
	77, 78, 79, 80, 50, 82, 51, 52, 92, 91,
	29, 85, 41, 81, 33, 34, 66, 103, 86, 45,
	93, 43, 42, 155, 63, 99, 83, 100, 148, 95,
	35, 36, 37, 38, 39, 40, 31, 32, 33, 34,
	64, 152, 101, 102, 138, 137, 143, 88, 31, 32,
	33, 34, 57, 57, 125, 104, 136, 11, 22, 111,
	106, 107, 14, 109, 110, 21, 19, 18, 20, 57,
	131, 119, 47, 48, 49, 121, 118, 15, 16, 127,
	142, 132, 126, 129, 26, 122, 4, 25, 141, 27,
	28, 128, 145, 58, 144, 55, 139, 150, 149, 147,
	134, 146, 133, 65, 57, 61, 153, 24, 23, 17,
	156, 151, 44, 124, 10, 57, 116, 158, 123, 159,
	22, 12, 13, 115, 14, 9, 8, 21, 19, 18,
	20, 6, 5, 3, 22, 1, 59, 0, 14, 15,
	16, 21, 19, 18, 20, 0, 26, 0, 41, 25,
	0, 27, 0, 15, 16, 0, 0, 0, 42, 0,
	26, 0, 0, 25, 0, 27, 35, 36, 37, 38,
	39, 40, 31, 32, 33, 34, 41, 0, 0, 0,
	112, 0, 0, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 0, 35, 36, 37, 38, 39, 40,
	31, 32, 33, 34, 41, 0, 0, 0, 105, 0,
	0, 0, 0, 0, 42, 0, 0, 0, 0, 0,
	0, 0, 35, 36, 37, 38, 39, 40, 31, 32,
	33, 34, 41, 0, 30, 0, 0, 0, 0, 0,
	0, 0, 42, 157, 0, 0, 0, 0, 0, 0,
	35, 36, 37, 38, 39, 40, 31, 32, 33, 34,
	41, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	42, 154, 0, 0, 0, 0, 0, 0, 35, 36,
	37, 38, 39, 40, 31, 32, 33, 34, 41, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 42, 140,
	0, 0, 0, 0, 0, 0, 35, 36, 37, 38,
	39, 40, 31, 32, 33, 34, 41, 114, 0, 0,
	0, 0, 0, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 0, 35, 36, 37, 38, 39, 40,
	31, 32, 33, 34, 41, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 42, 113, 0, 0, 0, 0,
	0, 0, 35, 36, 37, 38, 39, 40, 31, 32,
	33, 34, 41, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 42, 108, 0, 0, 0, 0, 0, 0,
	35, 36, 37, 38, 39, 40, 31, 32, 33, 34,
	41, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	42, 84, 0, 0, 0, 0, 0, 0, 35, 36,
	37, 38, 39, 40, 31, 32, 33, 34, 41, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 42, 41,
	0, 0, 0, 0, 0, 0, 35, 36, 37, 38,
	39, 40, 31, 32, 33, 34, 0, 35, 36, 37,
	38, 39, 40, 31, 32, 33, 34,
}
var yyPact = [...]int{

	149, -1000, -1000, 149, 16, -1000, -1000, 238, -1000, -1000,
	40, -1000, 38, 87, 87, 87, 87, 7, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 87, 163, 43, -4, -1000,
	87, 87, 87, 87, 87, 87, 87, 87, 87, 87,
	87, 87, 87, 17, 31, -1000, 434, -1000, -1000, -1000,
	87, 37, -28, 46, -13, 14, -1000, 462, -1000, 87,
	-31, -6, -1000, -33, 87, -1000, 87, -1000, 462, 22,
	22, -1000, -1000, 58, 58, 58, 58, 58, 58, -9,
	473, 87, 87, 36, 149, 210, -1000, -1000, -1000, -1000,
	-1000, 163, 163, 406, -1000, -1000, 43, 43, 87, 182,
	378, 462, 350, -1000, 108, -1000, -1000, -1000, 163, -1000,
	-1000, 462, -35, 43, 149, 85, 108, -1000, 87, 102,
	87, 88, 75, 74, -1000, 149, -1000, 322, 77, 102,
	-1000, 87, 462, 59, 88, -1000, 87, -1000, -1000, -1000,
	149, 71, -1000, 163, -1000, 294, 53, -1000, 43, -1000,
	266, -1000, -1000, -1000, 163, -1000, -1000, 43, -1000, -1000,
}
var yyPgo = [...]int{

	0, 175, 10, 173, 126, 172, 171, 0, 166, 165,
	163, 158, 156, 5, 153, 152, 97, 149, 148, 147,
	1, 145, 7, 6, 143, 142, 141, 140, 3, 139,
	2, 135, 8, 133, 131, 128, 123, 4, 120,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 3, 3, 4, 4, 4, 4,
	4, 9, 10, 10, 12, 12, 13, 11, 11, 14,
	8, 15, 15, 5, 6, 6, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 16,
	16, 16, 16, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 19, 20, 20, 21, 21, 21,
	23, 23, 23, 24, 25, 25, 27, 27, 28, 26,
	26, 29, 18, 30, 30, 31, 31, 31, 32, 32,
	33, 34, 34, 36, 36, 37, 35, 35, 38, 22,
	22,
}
var yyR2 = [...]int{

	0, 1, 1, 0, 3, 2, 1, 1, 1, 1,
	1, 7, 1, 0, 2, 1, 4, 1, 0, 2,
	7, 3, 1, 3, 2, 4, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 3, 3, 3, 3, 2, 0, 3, 3, 1,
	3, 5, 1, 7, 1, 0, 2, 1, 4, 1,
	0, 2, 3, 2, 0, 3, 3, 1, 1, 1,
	7, 1, 0, 2, 1, 4, 1, 0, 2, 1,
	1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	5, -16, 12, 13, 15, 30, 31, -17, 20, 19,
	21, 18, 11, -18, -19, 40, 37, 42, -4, 34,
	36, 30, 31, 32, 33, 24, 25, 26, 27, 28,
	29, 6, 16, 21, -15, 21, -7, -16, -16, -16,
	37, 39, 40, -7, -30, -31, -32, -7, -33, 13,
	-20, -21, -23, 21, 37, -24, 13, 34, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, 36, 14, 35, 17, -7, 21, 41, 41, 38,
	-22, 35, 34, -7, 43, -22, 35, 34, 44, -7,
	-7, -7, -7, 21, -2, 38, -32, -32, 17, -23,
	-23, -7, 38, 17, 7, -10, -12, -13, 8, -30,
	44, -20, -2, -11, -14, 9, -13, -7, -34, -36,
	-37, 8, -7, -25, -27, -28, 8, 10, 10, -2,
	17, -35, -38, 9, -37, -7, -26, -29, 9, -28,
	-7, -2, 10, -30, 17, 10, -20, 17, -30, -20,
}
var yyDef = [...]int{

	3, -2, 1, 2, 0, 6, 7, 8, 9, 10,
	0, 26, 0, 0, 0, 0, 0, 42, 43, 44,
	45, 46, 47, 48, 49, 0, 74, 56, 0, 5,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 24, 0, 22, 0, 39, 40, 41,
	0, 0, 0, 0, 0, 0, 77, 78, 79, 0,
	0, 0, 59, 0, 0, 62, 0, 4, 23, 27,
	28, 29, 30, -2, -2, -2, -2, -2, -2, 37,
	38, 0, 0, 0, 3, 0, 51, 52, 53, 72,
	73, 90, 89, 0, 54, 55, 90, 89, 0, 0,
	0, 25, 0, 21, 13, 50, 75, 76, 74, 57,
	58, 60, 0, 56, 3, 18, 12, 15, 0, 82,
	0, 65, 0, 0, 17, 3, 14, 0, 87, 81,
	84, 0, 61, 70, 64, 67, 0, 20, 11, 19,
	3, 0, 86, 74, 83, 0, 0, 69, 56, 66,
	0, 16, 80, 88, 74, 63, 71, 56, 85, 68,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	40, 41, 32, 30, 35, 31, 39, 33, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 44, 34,
	24, 36, 26, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 37, 3, 38, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 42, 3, 43,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 25, 27, 28, 29,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:47
		{
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:49
		{
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:50
		{
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:52
		{
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:53
		{
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:55
		{
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:56
		{
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:57
		{
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:58
		{
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:59
		{
		}
	case 11:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:66
		{
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:68
		{
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:69
		{
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:71
		{
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:72
		{
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:76
		{
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:78
		{
		}
	case 18:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:79
		{
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:83
		{
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:88
		{
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:90
		{
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:91
		{
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:93
		{
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:95
		{
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:96
		{
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:98
		{
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:99
		{
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:100
		{
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:101
		{
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:102
		{
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:103
		{
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:104
		{
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:105
		{
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:106
		{
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:107
		{
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:108
		{
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:109
		{
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:110
		{
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:112
		{
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:113
		{
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:114
		{
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:115
		{
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:117
		{
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:118
		{
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:119
		{
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:120
		{
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:121
		{
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:122
		{
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:123
		{
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:124
		{
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:125
		{
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:126
		{
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:127
		{
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:129
		{
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:131
		{
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:132
		{
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:134
		{
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:135
		{
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:136
		{
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:138
		{
		}
	case 61:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:139
		{
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:140
		{
		}
	case 63:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:147
		{
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:149
		{
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:150
		{
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:152
		{
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:153
		{
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:157
		{
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:159
		{
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:160
		{
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:164
		{
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:166
		{
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:168
		{
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:169
		{
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:171
		{
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:172
		{
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:173
		{
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:175
		{
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:176
		{
		}
	case 80:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:183
		{
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:185
		{
		}
	case 82:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:186
		{
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:193
		{
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:195
		{
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:196
		{
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:200
		{
		}
	}
	goto yystack /* stack new state and value */
}
