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
const kAND = 57347
const kELIF = 57348
const kELSE = 57349
const kEND = 57350
const kFALSE = 57351
const kIF = 57352
const kIN = 57353
const kNOT = 57354
const kOR = 57355
const kTHEN = 57356
const kTRUE = 57357
const NUMBER = 57358
const STRING = 57359
const ID = 57360
const OR = 57361
const AND = 57362
const LE = 57363
const GE = 57364
const EQ = 57365
const NE = 57366

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEXERR",
	"kAND",
	"kELIF",
	"kELSE",
	"kEND",
	"kFALSE",
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
	"','",
	"'='",
	"'.'",
	"'('",
	"')'",
	"'{'",
	"'}'",
	"':'",
	"'['",
	"']'",
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
	-1, 7,
	33, 31,
	-2, 10,
	-1, 76,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 17,
	-1, 77,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 18,
	-1, 78,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 19,
	-1, 79,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 20,
	-1, 80,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 21,
	-1, 81,
	21, 0,
	22, 0,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 22,
}

const yyPrivate = 57344

const yyLast = 403

var yyAct = [...]int{

	41, 29, 89, 97, 3, 38, 67, 17, 91, 62,
	19, 60, 35, 2, 24, 11, 37, 11, 36, 44,
	32, 40, 27, 7, 18, 7, 28, 8, 59, 8,
	46, 47, 48, 49, 63, 31, 64, 65, 9, 48,
	49, 33, 106, 22, 70, 14, 12, 72, 73, 74,
	75, 76, 77, 78, 79, 80, 81, 82, 83, 71,
	11, 68, 15, 84, 10, 58, 61, 16, 7, 99,
	103, 98, 8, 42, 39, 13, 17, 1, 20, 101,
	95, 4, 34, 30, 102, 96, 11, 23, 92, 6,
	5, 0, 94, 93, 7, 56, 0, 100, 8, 105,
	104, 0, 0, 57, 0, 107, 0, 0, 0, 0,
	109, 50, 51, 52, 53, 54, 55, 46, 47, 48,
	49, 43, 0, 21, 0, 0, 0, 26, 14, 12,
	85, 0, 0, 0, 0, 56, 0, 0, 0, 25,
	0, 0, 0, 57, 0, 15, 0, 10, 0, 0,
	16, 50, 51, 52, 53, 54, 55, 46, 47, 48,
	49, 0, 21, 0, 0, 66, 26, 14, 12, 0,
	88, 0, 9, 0, 0, 0, 0, 0, 25, 14,
	12, 0, 0, 0, 15, 0, 10, 0, 0, 16,
	26, 14, 12, 0, 0, 0, 15, 0, 10, 56,
	0, 16, 25, 0, 0, 0, 0, 57, 15, 0,
	10, 0, 0, 16, 0, 50, 51, 52, 53, 54,
	55, 46, 47, 48, 49, 90, 56, 0, 0, 0,
	0, 0, 0, 0, 57, 0, 0, 0, 0, 0,
	0, 0, 50, 51, 52, 53, 54, 55, 46, 47,
	48, 49, 69, 56, 0, 0, 0, 0, 0, 0,
	0, 57, 108, 0, 0, 0, 0, 0, 0, 50,
	51, 52, 53, 54, 55, 46, 47, 48, 49, 56,
	0, 0, 0, 0, 0, 0, 0, 57, 87, 0,
	0, 0, 0, 0, 0, 50, 51, 52, 53, 54,
	55, 46, 47, 48, 49, 56, 0, 0, 0, 0,
	0, 0, 0, 57, 86, 0, 0, 0, 0, 0,
	0, 50, 51, 52, 53, 54, 55, 46, 47, 48,
	49, 56, 0, 0, 0, 0, 0, 0, 0, 57,
	45, 0, 0, 0, 0, 0, 0, 50, 51, 52,
	53, 54, 55, 46, 47, 48, 49, 56, 0, 0,
	0, 0, 0, 0, 0, 57, 0, 56, 0, 0,
	0, 0, 0, 50, 51, 52, 53, 54, 55, 46,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 46,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 46,
	47, 48, 49,
}
var yyPact = [...]int{

	28, -1000, 28, -1000, -1000, -1000, -1000, -1000, -8, 150,
	2, -15, -1000, -1000, -1000, 150, 111, -1000, 150, 326,
	-1000, 174, -1000, -1000, -15, 12, -1000, -1000, -1000, -26,
	2, -1000, -29, 150, -1000, 150, 19, 130, -34, 111,
	-1000, 221, -1000, 150, 352, 28, 150, 150, 150, 150,
	150, 150, 150, 150, 150, 150, 150, 150, -1000, -1000,
	-1000, -1000, 150, 90, 300, -1000, -1000, -1000, -1000, -1000,
	274, 162, 10, 10, -1000, -1000, 3, 3, 3, 3,
	3, 3, 372, 362, 194, -30, 2, 111, -1000, -1000,
	-1000, 150, 65, 61, 194, 63, 65, -1000, 150, -1000,
	-1000, 34, -1000, 2, -1000, 248, -1000, -1000, 2, -1000,
}
var yyPgo = [...]int{

	0, 90, 0, 89, 26, 87, 3, 85, 84, 35,
	83, 82, 22, 81, 80, 79, 1, 4, 13, 43,
	14, 78, 77, 2, 75, 5, 74, 21, 73,
}
var yyR1 = [...]int{

	0, 22, 18, 18, 23, 23, 17, 17, 17, 3,
	13, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 4, 4, 21, 21, 19,
	19, 20, 20, 20, 20, 20, 5, 5, 12, 16,
	16, 10, 10, 9, 9, 9, 11, 14, 14, 7,
	7, 6, 15, 15, 8, 24, 25, 25, 26, 26,
	27, 27, 28,
}
var yyR2 = [...]int{

	0, 1, 2, 1, 1, 0, 1, 1, 1, 5,
	1, 3, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 3, 2, 1, 3, 1,
	0, 2, 1, 4, 6, 1, 7, 1, 0, 2,
	1, 4, 1, 0, 2, 3, 1, 0, 2, 1,
	2, 1, 5,
}
var yyChk = [...]int{

	-1000, -22, -18, -17, -13, -1, -3, -12, -4, 10,
	36, -20, 18, -24, 17, 34, 39, -17, 32, -2,
	-21, 12, -19, -5, -20, 28, 16, -12, -4, -16,
	-10, -9, 18, 39, -11, 10, 33, -2, -25, -26,
	-27, -2, -28, 10, -2, 14, 27, 28, 29, 30,
	21, 22, 23, 24, 25, 26, 5, 13, -19, 16,
	37, -9, 38, -2, -2, 18, 35, 40, -27, 31,
	-2, -18, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, -2, -2, -2, -2, 40, 14, 14, 8, -23,
	31, 38, -16, -25, -2, -14, -7, -6, 6, 8,
	-23, -15, -8, 7, -6, -2, 8, -16, 14, -16,
}
var yyDef = [...]int{

	0, -2, 1, 3, 6, 7, 8, -2, 34, 0,
	40, 0, 26, 32, 33, 0, 57, 2, 0, 0,
	12, 0, 28, 29, 30, 0, 37, 31, 34, 0,
	39, 42, 0, 0, 45, 0, 0, 0, 0, 56,
	59, 0, 61, 0, 11, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 27, 36,
	38, 41, 0, 0, 0, 25, 35, 55, 58, 60,
	0, 0, 13, 14, 15, 16, -2, -2, -2, -2,
	-2, -2, 23, 24, 5, 0, 40, 57, 9, 43,
	4, 0, 48, 0, 5, 53, 47, 50, 0, 62,
	44, 0, 52, 40, 49, 0, 46, 54, 40, 51,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	34, 35, 29, 27, 31, 28, 33, 30, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 38, 3,
	21, 32, 23, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 39, 3, 40, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 36, 3, 37,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 22,
	24, 25, 26,
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
//line parser.y:66
		{
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:68
		{
			yyVAL.ast = nil
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:69
		{
			yyVAL.ast = nil
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:74
		{
			yyVAL.ast = nil
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:75
		{
			yyVAL.ast = nil
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:76
		{
			yyVAL.ast = nil
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:81
		{
			yyVAL.ast = nil
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:83
		{
			yyVAL.ast = nil
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:85
		{
			yyVAL.ast = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:87
		{
			yyVAL.ast = nil
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:88
		{
			yyVAL.ast = nil
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:89
		{
			yyVAL.ast = nil
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:90
		{
			yyVAL.ast = nil
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:91
		{
			yyVAL.ast = nil
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:92
		{
			yyVAL.ast = nil
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:93
		{
			yyVAL.ast = nil
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:94
		{
			yyVAL.ast = nil
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:95
		{
			yyVAL.ast = nil
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:96
		{
			yyVAL.ast = nil
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:97
		{
			yyVAL.ast = nil
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:98
		{
			yyVAL.ast = nil
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:99
		{
			yyVAL.ast = nil
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:101
		{
			yyVAL.ast = nil
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:102
		{
			yyVAL.ast = nil
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:104
		{
			yyVAL.ast = nil
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:105
		{
			yyVAL.ast = nil
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:107
		{
			yyVAL.ast = nil
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:108
		{
			yyVAL.ast = nil
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:110
		{
			yyVAL.ast = nil
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:111
		{
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:112
		{
			yyVAL.ast = nil
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:113
		{
			yyVAL.ast = nil
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:114
		{
			yyVAL.ast = nil
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:116
		{
			yyVAL.ast = nil
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:117
		{
			yyVAL.ast = nil
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:119
		{
			yyVAL.ast = nil
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:121
		{
			yyVAL.ast = nil
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:122
		{
			yyVAL.ast = nil
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:124
		{
			yyVAL.ast = nil
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:125
		{
			yyVAL.ast = nil
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:127
		{
			yyVAL.ast = nil
		}
	case 44:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:128
		{
			yyVAL.ast = nil
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:129
		{
			yyVAL.ast = nil
		}
	case 46:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:136
		{
			yyVAL.ast = nil
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:138
		{
			yyVAL.ast = nil
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:139
		{
			yyVAL.ast = nil
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:141
		{
			yyVAL.ast = nil
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:142
		{
			yyVAL.ast = nil
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:146
		{
			yyVAL.ast = nil
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:148
		{
			yyVAL.ast = nil
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:149
		{
			yyVAL.ast = nil
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:153
		{
			yyVAL.ast = nil
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:155
		{
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:157
		{
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:158
		{
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:160
		{
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:161
		{
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:163
		{
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:164
		{
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:169
		{
		}
	}
	goto yystack /* stack new state and value */
}
