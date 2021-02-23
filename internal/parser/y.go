// Code generated by goyacc parser.y. DO NOT EDIT.

//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2

import (
	"github.com/dcaiafa/nitro/internal/ast"
	"github.com/dcaiafa/nitro/internal/token"
)

//line parser.y:11
type yySymType struct {
	yys   int
	tok   token.Token
	ast   ast.AST
	asts  ast.ASTs
	expr  ast.Expr
	exprs ast.Exprs
	other interface{}
}

const LEXERR = 57346
const kAND = 57347
const kDO = 57348
const kELIF = 57349
const kELSE = 57350
const kEND = 57351
const kFALSE = 57352
const kFN = 57353
const kFOR = 57354
const kIF = 57355
const kIN = 57356
const kNOT = 57357
const kOR = 57358
const kRETURN = 57359
const kTHEN = 57360
const kTRUE = 57361
const kVAR = 57362
const kWHILE = 57363
const NUMBER = 57364
const STRING = 57365
const ID = 57366
const EQ = 57367
const NE = 57368
const LE = 57369
const GE = 57370

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEXERR",
	"kAND",
	"kDO",
	"kELIF",
	"kELSE",
	"kEND",
	"kFALSE",
	"kFN",
	"kFOR",
	"kIF",
	"kIN",
	"kNOT",
	"kOR",
	"kRETURN",
	"kTHEN",
	"kTRUE",
	"kVAR",
	"kWHILE",
	"NUMBER",
	"STRING",
	"ID",
	"'='",
	"EQ",
	"NE",
	"'<'",
	"LE",
	"'>'",
	"GE",
	"'\"'",
	"'|'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"';'",
	"'('",
	"','",
	"'['",
	"':'",
	"'.'",
	"'{'",
	"')'",
	"']'",
	"'}'",
	"'%'",
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
	-1, 30,
	25, 75,
	40, 75,
	-2, 67,
	-1, 31,
	25, 76,
	40, 76,
	-2, 68,
	-1, 32,
	25, 77,
	40, 77,
	-2, 69,
	-1, 148,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	30, 0,
	31, 0,
	-2, 49,
	-1, 149,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	30, 0,
	31, 0,
	-2, 50,
	-1, 150,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	30, 0,
	31, 0,
	-2, 51,
	-1, 151,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	30, 0,
	31, 0,
	-2, 52,
	-1, 152,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	30, 0,
	31, 0,
	-2, 53,
	-1, 153,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	30, 0,
	31, 0,
	-2, 54,
	-1, 160,
	39, 74,
	41, 74,
	43, 74,
	-2, 18,
}

const yyPrivate = 57344

const yyLast = 838

var yyAct = [...]int{

	70, 211, 75, 67, 184, 69, 217, 77, 114, 124,
	130, 111, 123, 49, 3, 187, 186, 160, 158, 48,
	95, 96, 64, 65, 197, 23, 56, 32, 104, 66,
	134, 32, 97, 55, 31, 170, 179, 171, 31, 46,
	162, 93, 94, 95, 96, 88, 133, 173, 132, 102,
	103, 98, 99, 100, 101, 97, 90, 93, 94, 95,
	96, 116, 91, 118, 106, 107, 108, 54, 30, 85,
	32, 97, 30, 91, 128, 159, 126, 31, 125, 138,
	135, 120, 91, 136, 53, 20, 131, 110, 59, 20,
	139, 140, 84, 44, 143, 144, 145, 146, 147, 148,
	149, 150, 151, 152, 153, 154, 155, 142, 43, 58,
	116, 30, 89, 109, 129, 62, 42, 61, 156, 32,
	165, 166, 157, 137, 59, 60, 31, 62, 86, 61,
	189, 168, 169, 27, 57, 176, 167, 113, 50, 180,
	174, 175, 26, 141, 47, 25, 24, 37, 32, 83,
	82, 117, 45, 245, 242, 31, 241, 51, 52, 238,
	30, 78, 28, 190, 38, 192, 224, 39, 221, 191,
	27, 57, 220, 188, 196, 50, 195, 20, 79, 26,
	199, 198, 25, 24, 37, 32, 205, 207, 204, 30,
	194, 233, 31, 227, 51, 52, 200, 203, 214, 28,
	218, 38, 206, 212, 39, 185, 20, 127, 32, 161,
	1, 5, 228, 229, 32, 31, 40, 230, 222, 235,
	29, 31, 236, 234, 12, 226, 30, 225, 210, 209,
	32, 239, 72, 71, 68, 35, 243, 31, 237, 232,
	231, 216, 215, 20, 246, 81, 80, 247, 76, 30,
	32, 36, 104, 115, 112, 30, 33, 31, 34, 8,
	87, 15, 7, 105, 9, 202, 20, 201, 183, 182,
	10, 30, 20, 102, 103, 98, 99, 100, 101, 11,
	63, 93, 94, 95, 96, 13, 14, 4, 20, 164,
	104, 30, 2, 163, 0, 97, 0, 0, 0, 0,
	0, 105, 0, 244, 0, 0, 0, 0, 20, 0,
	0, 102, 103, 98, 99, 100, 101, 104, 0, 93,
	94, 95, 96, 0, 0, 0, 0, 0, 105, 0,
	240, 0, 0, 97, 0, 0, 0, 0, 102, 103,
	98, 99, 100, 101, 104, 0, 93, 94, 95, 96,
	27, 57, 0, 0, 0, 105, 0, 223, 0, 26,
	97, 0, 25, 24, 37, 102, 103, 98, 99, 100,
	101, 104, 219, 93, 94, 95, 96, 0, 0, 28,
	0, 38, 105, 0, 39, 0, 0, 97, 0, 0,
	0, 0, 102, 103, 98, 99, 100, 101, 104, 213,
	93, 94, 95, 96, 0, 0, 0, 0, 0, 105,
	0, 0, 0, 0, 97, 0, 0, 0, 0, 102,
	103, 98, 99, 100, 101, 104, 0, 93, 94, 95,
	96, 0, 0, 0, 0, 0, 105, 0, 0, 0,
	0, 97, 0, 0, 0, 0, 102, 103, 98, 99,
	100, 101, 104, 0, 93, 94, 95, 96, 0, 0,
	0, 0, 0, 105, 0, 0, 208, 0, 97, 0,
	0, 0, 0, 102, 103, 98, 99, 100, 101, 104,
	181, 93, 94, 95, 96, 0, 0, 0, 0, 0,
	105, 0, 0, 193, 0, 97, 0, 0, 0, 0,
	102, 103, 98, 99, 100, 101, 104, 0, 93, 94,
	95, 96, 0, 0, 0, 0, 0, 105, 0, 178,
	0, 0, 97, 0, 0, 0, 0, 102, 103, 98,
	99, 100, 101, 104, 0, 93, 94, 95, 96, 0,
	0, 0, 0, 0, 105, 0, 0, 0, 0, 97,
	0, 0, 0, 0, 102, 103, 98, 99, 100, 101,
	104, 0, 93, 94, 95, 96, 0, 0, 0, 0,
	0, 105, 0, 172, 177, 0, 97, 0, 0, 0,
	0, 102, 103, 98, 99, 100, 101, 104, 0, 93,
	94, 95, 96, 0, 0, 0, 0, 0, 105, 0,
	0, 0, 0, 97, 0, 0, 0, 0, 102, 103,
	98, 99, 100, 101, 104, 121, 93, 94, 95, 96,
	0, 0, 0, 0, 0, 105, 0, 122, 0, 0,
	97, 0, 0, 0, 0, 102, 103, 98, 99, 100,
	101, 104, 0, 93, 94, 95, 96, 0, 0, 0,
	0, 0, 105, 0, 92, 0, 0, 97, 0, 0,
	0, 0, 102, 103, 98, 99, 100, 101, 104, 0,
	93, 94, 95, 96, 0, 0, 0, 0, 0, 105,
	0, 0, 0, 0, 97, 0, 0, 0, 0, 102,
	103, 98, 99, 100, 101, 0, 0, 93, 94, 95,
	96, 0, 0, 27, 19, 17, 18, 0, 0, 0,
	21, 97, 26, 16, 22, 25, 24, 37, 102, 103,
	98, 99, 100, 101, 0, 0, 93, 94, 95, 96,
	0, 6, 28, 0, 38, 0, 0, 39, 0, 0,
	97, 27, 19, 17, 18, 0, 0, 0, 21, 0,
	26, 16, 22, 25, 24, 37, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 41,
	28, 0, 38, 0, 0, 39, 27, 57, 74, 73,
	0, 50, 0, 0, 0, 26, 0, 0, 25, 24,
	37, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	51, 52, 0, 27, 57, 28, 0, 38, 50, 0,
	39, 0, 26, 0, 0, 25, 24, 37, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 51, 52, 0,
	0, 0, 28, 0, 38, 119, 0, 39,
}
var yyPact = [...]int{

	693, -1000, -1000, -1000, 731, 78, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 68, 128, 120, 160, 85,
	86, 160, 160, -1000, -1000, -1000, -1000, -1000, 160, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 766, 137,
	54, -1000, -1000, 340, 160, 87, 42, -1000, 636, -1000,
	160, 160, 160, 74, -1000, -1000, -1000, 49, 48, 113,
	160, 127, 793, 41, 663, 609, 582, -34, 38, -1000,
	663, -1000, -1000, 160, 120, -37, 8, -1000, -12, 160,
	-1000, -1000, 160, 120, -1000, -1000, 74, 39, 663, 160,
	160, 119, 693, 160, 160, 160, 160, 160, 160, 160,
	160, 160, 160, 160, 160, 160, -1000, -1000, -1000, 160,
	113, -27, 35, -1000, -28, 0, 663, -1000, 247, 160,
	160, 693, -1000, -1000, -1000, 766, 766, -3, 555, 33,
	-1000, -1000, 137, 137, 160, 528, 501, 22, 160, 663,
	474, -1000, 198, -16, -16, -1000, -1000, -1000, 7, 7,
	7, 7, 7, 7, 692, 23, -29, -30, 693, 106,
	-1000, -1000, 160, -1000, 123, 447, 663, 181, -1000, -1000,
	-1000, -1000, 766, 160, -1000, -1000, 663, -18, 137, 160,
	663, 693, 189, 198, -1000, 160, -1000, 693, 178, -1000,
	663, -1000, 420, -1000, -1000, 196, 393, 160, 193, 366,
	163, 159, -1000, 693, -1000, 339, 157, -1000, -1000, 185,
	196, -1000, 160, 766, 663, 183, 193, -1000, 160, 137,
	-1000, -1000, -1000, 693, -1000, 150, -1000, 766, -1000, 312,
	147, 145, -1000, 137, -1000, 285, 144, -1000, -1000, -1000,
	766, -1000, -1000, -1000, 137, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 292, 14, 287, 211, 286, 285, 280, 279, 270,
	269, 268, 4, 267, 265, 264, 39, 262, 261, 260,
	259, 0, 13, 84, 258, 67, 33, 26, 256, 11,
	254, 8, 253, 251, 2, 248, 7, 246, 245, 242,
	241, 6, 240, 239, 235, 3, 234, 5, 233, 232,
	229, 228, 1, 227, 225, 25, 224, 220, 210, 209,
	9, 207,
}
var yyR1 = [...]int{

	0, 58, 1, 2, 2, 2, 3, 3, 3, 3,
	4, 4, 4, 4, 4, 4, 4, 4, 56, 5,
	6, 7, 7, 8, 9, 10, 10, 11, 11, 12,
	13, 13, 14, 15, 16, 16, 18, 18, 19, 19,
	17, 20, 20, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 22, 22, 22,
	22, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 57, 55, 55, 55, 25, 26,
	27, 28, 28, 28, 24, 29, 29, 30, 30, 31,
	31, 32, 32, 33, 34, 34, 35, 35, 35, 36,
	36, 36, 36, 37, 38, 39, 39, 40, 40, 41,
	42, 42, 43, 44, 45, 45, 46, 46, 46, 47,
	47, 47, 48, 49, 50, 50, 51, 51, 52, 53,
	53, 54, 60, 60, 61, 61, 61, 61, 59, 59,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 0, 3, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 4, 5,
	2, 3, 1, 7, 7, 1, 0, 2, 1, 4,
	1, 0, 2, 7, 3, 1, 3, 1, 3, 1,
	3, 2, 4, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	1, 1, 1, 1, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 4, 1, 1, 1, 1, 3,
	4, 5, 5, 6, 6, 1, 0, 3, 1, 2,
	0, 3, 1, 3, 2, 0, 3, 3, 1, 3,
	5, 1, 1, 7, 7, 1, 0, 2, 1, 4,
	1, 0, 2, 3, 2, 0, 3, 3, 1, 1,
	1, 1, 7, 7, 1, 0, 2, 1, 4, 1,
	0, 2, 1, 0, 2, 2, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -58, -1, -2, -3, -4, 38, -17, -20, -15,
	-9, -8, -56, -6, -5, -18, 20, 12, 13, 11,
	-23, 17, 21, -55, 23, 22, 19, 10, 39, -57,
	-25, -26, -27, -28, -24, -44, -33, 24, 41, 44,
	-4, 38, 38, 40, 25, 24, -16, 24, -21, -22,
	15, 34, 35, -23, -25, -26, -27, 11, 24, 39,
	39, 43, 41, -7, -21, -21, -21, -45, -46, -47,
	-21, -48, -49, 13, 12, -34, -35, -36, 24, 41,
	-37, -38, 13, 12, 38, -55, -23, -19, -21, 25,
	14, 40, 18, 34, 35, 36, 37, 48, 28, 29,
	30, 31, 26, 27, 5, 16, -22, -22, -22, 39,
	39, -29, -30, 24, -31, -32, -21, 24, -21, 42,
	40, 6, 45, 46, -60, 40, 38, -61, -21, -16,
	47, -60, 40, 38, 42, -21, -21, -16, 40, -21,
	-21, 24, -2, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -31, -29, 45, 40,
	45, -59, 40, 46, 42, -21, -21, -2, -47, -47,
	38, 40, 18, 14, -36, -36, -21, 46, 18, 14,
	-21, 6, -10, -11, -12, 7, 45, 45, -2, 24,
	-21, 46, -21, 46, 9, -45, -21, 42, -34, -21,
	-2, -13, -14, 8, -12, -21, -2, 9, 46, -50,
	-51, -52, 7, 6, -21, -39, -40, -41, 7, 6,
	9, 9, -2, 18, 9, -53, -54, 8, -52, -21,
	-45, -42, -43, 8, -41, -21, -34, -2, 9, -45,
	18, 9, 9, -34, 18, 9, -45, -34,
}
var yyDef = [...]int{

	5, -2, 1, 2, 3, 4, 9, 10, 11, 12,
	13, 14, 15, 16, 17, 0, 0, 0, 0, 0,
	0, 0, 0, 37, 61, 62, 63, 64, 0, 66,
	-2, -2, -2, 70, 71, 72, 73, 78, 115, 95,
	0, 7, 8, 0, 0, 41, 0, 35, 0, 43,
	0, 0, 0, 60, 67, 68, 69, 0, 0, 86,
	90, 0, 0, 20, 22, 0, 0, 0, 133, 118,
	119, 120, 121, 0, 0, 0, 133, 98, 0, 0,
	101, 102, 0, 0, 6, 36, 0, 40, 39, 0,
	0, 0, 5, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 57, 58, 59, 90,
	86, 0, 85, 88, 0, 139, 92, 79, 0, 0,
	0, 5, 65, 113, 114, 137, 136, 132, 0, 0,
	93, 94, 137, 136, 0, 0, 0, 0, 0, 42,
	0, 34, 26, 44, 45, 46, 47, 48, -2, -2,
	-2, -2, -2, -2, 55, 56, 0, 0, 5, 0,
	-2, 89, 138, 80, 0, 0, 21, 0, 116, 117,
	134, 135, 115, 0, 96, 97, 99, 0, 95, 0,
	38, 5, 31, 25, 28, 0, 74, 5, 0, 87,
	91, 81, 0, 82, 19, 125, 0, 0, 106, 0,
	0, 0, 30, 5, 27, 0, 0, 84, 83, 130,
	124, 127, 0, 115, 100, 111, 105, 108, 0, 95,
	33, 24, 32, 5, 23, 0, 129, 115, 126, 0,
	0, 0, 110, 95, 107, 0, 0, 29, 122, 131,
	115, 123, 103, 112, 95, 104, 128, 109,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 32, 3, 3, 48, 3, 3,
	39, 45, 36, 34, 40, 35, 43, 37, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 42, 38,
	28, 25, 30, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 41, 3, 46, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 44, 33, 47,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 26, 27, 29, 31,
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
//line parser.y:115
		{
			yylex.(*lex).Module = yyDollar[1].ast.(*ast.Module)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:118
		{
			m := &ast.Module{}
			m.Stmts = yyDollar[1].asts
			m.SetPos(yyDollar[1].asts.Pos())
			yyVAL.ast = m
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:126
		{
			yyVAL.asts = ast.ASTs{yyDollar[1].ast}
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:127
		{
			yyVAL.asts = ast.ASTs{}
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:129
		{
			l := yyDollar[1].asts
			yyVAL.asts = append(l, yyDollar[2].ast)
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:131
		{
			yyVAL.asts = ast.ASTs{yyDollar[1].ast}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:132
		{
			yyVAL.asts = ast.ASTs{}
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:144
		{
			yyVAL.ast = &ast.FuncCallStmt{Target: yyDollar[1].expr, Args: yyDollar[3].exprs}
			yyVAL.ast.SetPos(yyDollar[1].expr.Pos())
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:152
		{
			yyVAL.ast = &ast.WhileStmt{Predicate: yyDollar[2].expr, Stmts: yyDollar[4].asts}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:158
		{
			yyVAL.ast = &ast.ReturnStmt{Values: yyDollar[2].ast.(ast.Exprs)}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:164
		{
			l := yyDollar[1].ast.(ast.Exprs)
			yyVAL.ast = append(l, yyDollar[3].expr)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:169
		{
			yyVAL.ast = ast.Exprs{yyDollar[1].expr}
			yyVAL.ast.SetPos(yyDollar[1].expr.Pos())
		}
	case 23:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:177
		{
			fn := &ast.FuncStmt{}
			fn.Name = yyDollar[2].tok.Str
			fn.Params = yyDollar[4].asts
			fn.Stmts = yyDollar[6].asts
			fn.SetPos(yyDollar[1].tok.Pos)
			yyVAL.ast = fn
		}
	case 24:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:191
		{
			ifBlock := &ast.IfBlock{Pred: yyDollar[2].expr, Stmts: yyDollar[4].asts}
			ifStmt := &ast.IfStmt{Blocks: ast.ASTs{ifBlock}}
			if yyDollar[5].asts != nil {
				ifStmt.Blocks = append(ifStmt.Blocks, yyDollar[5].asts...)
			}
			if yyDollar[6].ast != nil {
				ifStmt.Blocks = append(ifStmt.Blocks, yyDollar[6].ast)
			}
			yyVAL.ast = ifStmt
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:206
		{
			yyVAL.asts = ast.ASTs{}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:211
		{
			l := yyDollar[1].asts
			yyVAL.asts = append(l, yyDollar[2].ast)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:216
		{
			yyVAL.asts = ast.ASTs{yyDollar[1].ast}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:222
		{
			yyVAL.ast = &ast.IfBlock{Pred: yyDollar[2].expr, Stmts: yyDollar[4].asts}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:229
		{
			yyVAL.ast = nil
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:235
		{
			yyVAL.ast = &ast.IfBlock{Pred: nil, Stmts: yyDollar[2].asts}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 33:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:243
		{
			yyVAL.ast = &ast.ForStmt{ForVars: yyDollar[2].asts, IterExpr: yyDollar[4].expr, Stmts: yyDollar[6].asts}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:249
		{
			yyVAL.asts = append(yyDollar[1].asts, &ast.ForVar{VarName: yyDollar[3].tok})
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:253
		{
			yyVAL.asts = ast.ASTs{&ast.ForVar{VarName: yyDollar[1].tok}}
			yyVAL.asts.SetPos(yyDollar[1].tok.Pos)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:259
		{
			yyVAL.asts = append(yyDollar[1].asts, &ast.LValue{Expr: yyDollar[3].expr})
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:263
		{
			lvalue := &ast.LValue{Expr: yyDollar[1].expr}
			lvalue.SetPos(yyDollar[1].expr.Pos())
			yyVAL.asts = ast.ASTs{lvalue}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:270
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:274
		{
			yyVAL.exprs = ast.Exprs{yyDollar[1].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:279
		{
			yyVAL.ast = &ast.AssignStmt{Lvalues: yyDollar[1].asts, Rvalues: yyDollar[3].exprs}
			yyVAL.ast.SetPos(yyDollar[1].asts.Pos())
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:285
		{
			yyVAL.ast = &ast.VarDeclStmt{VarName: yyDollar[2].tok, InitValue: nil}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:290
		{
			yyVAL.ast = &ast.VarDeclStmt{VarName: yyDollar[2].tok, InitValue: yyDollar[4].expr}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:297
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpPlus}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:302
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpMinus}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:307
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpMult}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:312
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpDiv}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:317
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpMod}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:322
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpLT}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:327
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpLE}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:332
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpGT}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:337
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpGE}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:342
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpEq}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:347
		{
			yyVAL.expr = &ast.BinaryExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Op: ast.BinOpNE}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:352
		{
			yyVAL.expr = &ast.AndExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:357
		{
			yyVAL.expr = &ast.OrExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:363
		{
			yyVAL.expr = &ast.UnaryExpr{Term: yyDollar[2].expr, Op: ast.UnaryOpNot}
			yyVAL.expr.SetPos(yyDollar[1].tok.Pos)
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:368
		{
			yyVAL.expr = &ast.UnaryExpr{Term: yyDollar[2].expr, Op: ast.UnaryOpPlus}
			yyVAL.expr.SetPos(yyDollar[1].tok.Pos)
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:373
		{
			yyVAL.expr = &ast.UnaryExpr{Term: yyDollar[2].expr, Op: ast.UnaryOpMinus}
			yyVAL.expr.SetPos(yyDollar[1].tok.Pos)
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:380
		{
			yyVAL.expr = &ast.LiteralExpr{Val: yyDollar[1].tok}
			yyVAL.expr.SetPos(yyDollar[1].tok.Pos)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:385
		{
			yyVAL.expr = &ast.LiteralExpr{Val: yyDollar[1].tok}
			yyVAL.expr.SetPos(yyDollar[1].tok.Pos)
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:390
		{
			yyVAL.expr = &ast.LiteralExpr{Val: yyDollar[1].tok}
			yyVAL.expr.SetPos(yyDollar[1].tok.Pos)
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:395
		{
			yyVAL.expr = &ast.LiteralExpr{Val: yyDollar[1].tok}
			yyVAL.expr.SetPos(yyDollar[1].tok.Pos)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:400
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:413
		{
			yyVAL.expr = &ast.FuncCallExpr{Target: yyDollar[1].expr, Args: yyDollar[3].exprs, RetN: 1}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:423
		{
			yyVAL.expr = &ast.SimpleRef{ID: yyDollar[1].tok}
			yyVAL.expr.SetPos(yyDollar[1].tok.Pos)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:429
		{
			yyVAL.expr = &ast.MemberAccess{Target: yyDollar[1].expr, Member: yyDollar[3].tok}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:435
		{
			yyVAL.expr = &ast.IndexExpr{Target: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:441
		{
			yyVAL.expr = &ast.SliceExpr{Target: yyDollar[1].expr, Begin: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:446
		{
			yyVAL.expr = &ast.SliceExpr{Target: yyDollar[1].expr, End: yyDollar[4].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 83:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:451
		{
			yyVAL.expr = &ast.SliceExpr{Target: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 84:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:459
		{
			l := &ast.LambdaExpr{}
			l.Params = yyDollar[3].asts
			l.Stmts = yyDollar[5].asts
			l.SetPos(yyDollar[1].tok.Pos)
			yyVAL.expr = l
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:469
		{
			yyVAL.asts = ast.ASTs{}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:474
		{
			param := &ast.FuncParam{Name: yyDollar[3].tok.Str}
			param.SetPos(yyDollar[3].tok.Pos)
			yyVAL.asts = append(yyDollar[1].asts, param)
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:480
		{
			yyVAL.asts = ast.ASTs{&ast.FuncParam{Name: yyDollar[1].tok.Str}}
			yyVAL.asts.SetPos(yyDollar[1].tok.Pos)
		}
	case 90:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:487
		{
			yyVAL.exprs = ast.Exprs{}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:492
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:496
		{
			yyVAL.exprs = ast.Exprs{yyDollar[1].expr}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:501
		{
			yyVAL.expr = &ast.ObjectLiteral{Fields: yyDollar[2].asts}
			yyVAL.expr.SetPos(yyDollar[1].tok.Pos)
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:508
		{
			yyVAL.asts = ast.ASTs{}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:513
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:517
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:521
		{
			yyVAL.asts = ast.ASTs{yyDollar[1].ast}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:526
		{
			yyVAL.ast = &ast.ObjectField{NameID: yyDollar[1].tok.Str, Val: yyDollar[3].expr}
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:530
		{
			yyVAL.ast = &ast.ObjectField{NameExpr: yyDollar[2].expr, Val: yyDollar[5].expr}
		}
	case 103:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:541
		{
			ifBlock := &ast.IfBlock{Pred: yyDollar[2].expr, Stmts: yyDollar[4].asts}
			ifStmt := &ast.IfStmt{Blocks: ast.ASTs{ifBlock}}
			if yyDollar[5].asts != nil {
				ifStmt.Blocks = append(ifStmt.Blocks, yyDollar[5].asts...)
			}
			if yyDollar[6].ast != nil {
				ifStmt.Blocks = append(ifStmt.Blocks, yyDollar[6].ast)
			}
			yyVAL.ast = ifStmt
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 104:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:557
		{
			yyVAL.ast = &ast.ForStmt{ForVars: yyDollar[2].asts, IterExpr: yyDollar[4].expr, Stmts: yyDollar[6].asts}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:564
		{
			yyVAL.asts = ast.ASTs{}
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:569
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[2].ast)
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:573
		{
			yyVAL.asts = ast.ASTs{yyDollar[1].ast}
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:579
		{
			yyVAL.ast = &ast.IfBlock{Pred: yyDollar[2].expr, Stmts: yyDollar[4].asts}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:586
		{
			yyVAL.ast = nil
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:592
		{
			yyVAL.ast = &ast.IfBlock{Pred: nil, Stmts: yyDollar[2].asts}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:598
		{
			yyVAL.expr = &ast.ArrayLiteral{Elements: yyDollar[2].asts}
		}
	case 115:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:604
		{
			yyVAL.asts = ast.ASTs{}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:609
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:613
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[3].ast)
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:617
		{
			yyVAL.asts = ast.ASTs{yyDollar[1].ast}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:622
		{
			yyVAL.ast = &ast.ArrayElement{Val: yyDollar[1].expr}
			yyVAL.ast.SetPos(yyDollar[1].expr.Pos())
		}
	case 122:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:634
		{
			ifBlock := &ast.IfBlock{Pred: yyDollar[2].expr, Stmts: yyDollar[4].asts}
			ifStmt := &ast.IfStmt{Blocks: ast.ASTs{ifBlock}}
			if yyDollar[5].asts != nil {
				ifStmt.Blocks = append(ifStmt.Blocks, yyDollar[5].asts...)
			}
			if yyDollar[6].ast != nil {
				ifStmt.Blocks = append(ifStmt.Blocks, yyDollar[6].ast)
			}
			yyVAL.ast = ifStmt
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 123:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:650
		{
			yyVAL.ast = &ast.ForStmt{ForVars: yyDollar[2].asts, IterExpr: yyDollar[4].expr, Stmts: yyDollar[6].asts}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 125:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:657
		{
			yyVAL.asts = ast.ASTs{}
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:662
		{
			yyVAL.asts = append(yyDollar[1].asts, yyDollar[2].ast)
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:666
		{
			yyVAL.asts = ast.ASTs{yyDollar[1].ast}
		}
	case 128:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:672
		{
			yyVAL.ast = &ast.IfBlock{Pred: yyDollar[2].expr, Stmts: yyDollar[4].asts}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	case 130:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:679
		{
			yyVAL.ast = nil
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:685
		{
			yyVAL.ast = &ast.IfBlock{Pred: nil, Stmts: yyDollar[2].asts}
			yyVAL.ast.SetPos(yyDollar[1].tok.Pos)
		}
	}
	goto yystack /* stack new state and value */
}
