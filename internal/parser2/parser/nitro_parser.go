// Code generated from NitroParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NitroParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type NitroParser struct {
	*antlr.BaseParser
}

var nitroparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nitroparserParserInit() {
	staticData := &nitroparserParserStaticData
	staticData.literalNames = []string{
		"", "'!info'", "'!param'", "'!flag'", "", "'break'", "'catch'", "'continue'",
		"'defer'", "'else'", "", "'for'", "'func'", "'if'", "'import'", "",
		"", "", "'return'", "'throw'", "", "'try'", "'var'", "'while'", "'yield'",
		"'='", "'+='", "'-='", "'*='", "'/='", "'%='", "", "", "", "", "", "",
		"", "", "", "", "", "'++'", "'--'", "", "';'", "", "", "", "", "", "",
		"", "'{'", "", "", "", "", "", "", "", "", "", "", "", "", "", "'e`'",
		"", "'~'", "'`'",
	}
	staticData.symbolicNames = []string{
		"", "M_INFO", "M_PARAM", "M_FLAG", "AND", "BREAK", "CATCH", "CONTINUE",
		"DEFER", "ELSE", "FALSE", "FOR", "FUNC", "IF", "IMPORT", "NIL", "NOT",
		"OR", "RETURN", "THROW", "TRUE", "TRY", "VAR", "WHILE", "YIELD", "ASSIGN",
		"ASSIGN_ADD", "ASSIGN_SUB", "ASSIGN_MUL", "ASSIGN_DIV", "ASSIGN_MOD",
		"EQ", "NE", "LT", "LE", "GT", "GE", "ADD", "SUB", "MUL", "DIV", "MOD",
		"INC", "DEC", "QUESTION_MARK", "SEMICOLON", "COMMA", "COLON", "PERIOD",
		"OPAREN", "CPAREN", "OBRACKET", "CBRACKET", "OCURLY", "CCURLY", "ARROW",
		"LAMBDA", "PIPE", "EXPAND", "NUMBER", "ID", "REGEX", "NEWLINE", "CHAR",
		"COMMENT", "WS", "STRING", "EXEC_PREFIX", "EXEC_WS", "EXEC_HOME", "EXEC_SUFFIX",
		"EXEC_LITERAL", "EXEC_DQUOTE_LITERAL", "EXEC_SQUOTE_LITERAL", "EXEC_EXPR_WS_",
		"DQUOTE_", "EXEC_SQUOTE_",
	}
	staticData.ruleNames = []string{
		"start", "unit", "prologue", "meta_directive", "meta_info", "meta_param",
		"meta_flag", "meta_attribs", "meta_attrib", "meta_literal", "import_stmt",
		"stmts", "stmt_list", "stmt", "assignment_stmt", "assignment_lvalues",
		"rvalues", "assignment_op_stmt", "var_decl_stmt", "var_decl_vars", "for_stmt",
		"for_vars", "while_stmt", "if_stmt", "if_elif", "if_else", "func_stmt",
		"param_list", "return_stmt", "try_catch_stmt", "throw_stmt", "defer_stmt",
		"yield_stmt", "break_stmt", "continue_stmt", "inc_dec_stmt", "expr",
		"expr3", "binary_expr", "unary_expr", "primary_expr", "simple_literal",
		"arg_list", "lvalue_expr", "lambda_expr", "short_lambda_expr", "exec_expr",
		"exec_expr_arg", "object_literal", "object_fields", "object_field",
		"array_literal", "array_elems", "array_elem", "id_or_keyword",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 76, 593, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		2, 5, 2, 118, 8, 2, 10, 2, 12, 2, 121, 9, 2, 1, 2, 5, 2, 124, 8, 2, 10,
		2, 12, 2, 127, 9, 2, 1, 3, 1, 3, 1, 3, 3, 3, 132, 8, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 139, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 147, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 153, 8, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 161, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 167, 8,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 174, 8, 7, 10, 7, 12, 7, 177, 9,
		7, 1, 7, 3, 7, 180, 8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 185, 8, 8, 1, 9, 1, 9,
		1, 10, 1, 10, 3, 10, 191, 8, 10, 1, 10, 1, 10, 1, 10, 1, 11, 3, 11, 197,
		8, 11, 1, 11, 5, 11, 200, 8, 11, 10, 11, 12, 11, 203, 9, 11, 1, 12, 1,
		12, 4, 12, 207, 8, 12, 11, 12, 12, 12, 208, 1, 12, 5, 12, 212, 8, 12, 10,
		12, 12, 12, 215, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 233,
		8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 5, 15, 242, 8,
		15, 10, 15, 12, 15, 245, 9, 15, 1, 16, 1, 16, 1, 16, 5, 16, 250, 8, 16,
		10, 16, 12, 16, 253, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 3, 18, 263, 8, 18, 1, 19, 1, 19, 1, 19, 5, 19, 268, 8, 19, 10,
		19, 12, 19, 271, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 5, 21, 284, 8, 21, 10, 21, 12, 21, 287, 9,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 5, 23, 301, 8, 23, 10, 23, 12, 23, 304, 9, 23, 1, 23, 3,
		23, 307, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 325, 8,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 335,
		8, 27, 10, 27, 12, 27, 338, 9, 27, 1, 28, 1, 28, 3, 28, 342, 8, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 350, 8, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 3, 36,
		374, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 5, 37, 385, 8, 37, 10, 37, 12, 37, 388, 9, 37, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 411, 8, 38,
		10, 38, 12, 38, 414, 9, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 3, 39, 423, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 437, 8, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 450,
		8, 40, 1, 40, 1, 40, 3, 40, 454, 8, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 3, 40, 461, 8, 40, 3, 40, 463, 8, 40, 1, 40, 5, 40, 466, 8, 40, 10,
		40, 12, 40, 469, 9, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 5, 42, 476,
		8, 42, 10, 42, 12, 42, 479, 9, 42, 1, 42, 3, 42, 482, 8, 42, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43, 494,
		8, 43, 1, 44, 1, 44, 1, 44, 3, 44, 499, 8, 44, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 45, 1, 45, 3, 45, 508, 8, 45, 1, 45, 1, 45, 1, 45, 1, 46,
		1, 46, 4, 46, 515, 8, 46, 11, 46, 12, 46, 516, 1, 46, 1, 46, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 529, 8, 47, 1, 47,
		1, 47, 3, 47, 533, 8, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1,
		49, 5, 49, 542, 8, 49, 10, 49, 12, 49, 545, 9, 49, 1, 49, 3, 49, 548, 8,
		49, 3, 49, 550, 8, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 3, 50, 565, 8, 50, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 5, 52, 574, 8, 52, 10, 52, 12, 52,
		577, 9, 52, 1, 52, 5, 52, 580, 8, 52, 10, 52, 12, 52, 583, 9, 52, 3, 52,
		585, 8, 52, 1, 53, 1, 53, 3, 53, 589, 8, 53, 1, 54, 1, 54, 1, 54, 0, 3,
		74, 76, 80, 55, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100,
		102, 104, 106, 108, 0, 8, 1, 0, 45, 46, 6, 0, 10, 10, 15, 15, 20, 20, 59,
		59, 63, 63, 66, 66, 1, 0, 26, 30, 1, 0, 42, 43, 1, 0, 39, 41, 1, 0, 37,
		38, 1, 0, 31, 36, 2, 0, 4, 24, 60, 60, 628, 0, 110, 1, 0, 0, 0, 2, 113,
		1, 0, 0, 0, 4, 119, 1, 0, 0, 0, 6, 131, 1, 0, 0, 0, 8, 133, 1, 0, 0, 0,
		10, 142, 1, 0, 0, 0, 12, 156, 1, 0, 0, 0, 14, 170, 1, 0, 0, 0, 16, 181,
		1, 0, 0, 0, 18, 186, 1, 0, 0, 0, 20, 188, 1, 0, 0, 0, 22, 196, 1, 0, 0,
		0, 24, 204, 1, 0, 0, 0, 26, 232, 1, 0, 0, 0, 28, 234, 1, 0, 0, 0, 30, 238,
		1, 0, 0, 0, 32, 246, 1, 0, 0, 0, 34, 254, 1, 0, 0, 0, 36, 258, 1, 0, 0,
		0, 38, 264, 1, 0, 0, 0, 40, 272, 1, 0, 0, 0, 42, 280, 1, 0, 0, 0, 44, 288,
		1, 0, 0, 0, 46, 294, 1, 0, 0, 0, 48, 308, 1, 0, 0, 0, 50, 315, 1, 0, 0,
		0, 52, 320, 1, 0, 0, 0, 54, 331, 1, 0, 0, 0, 56, 339, 1, 0, 0, 0, 58, 343,
		1, 0, 0, 0, 60, 355, 1, 0, 0, 0, 62, 358, 1, 0, 0, 0, 64, 361, 1, 0, 0,
		0, 66, 364, 1, 0, 0, 0, 68, 366, 1, 0, 0, 0, 70, 368, 1, 0, 0, 0, 72, 373,
		1, 0, 0, 0, 74, 375, 1, 0, 0, 0, 76, 389, 1, 0, 0, 0, 78, 422, 1, 0, 0,
		0, 80, 436, 1, 0, 0, 0, 82, 470, 1, 0, 0, 0, 84, 472, 1, 0, 0, 0, 86, 493,
		1, 0, 0, 0, 88, 495, 1, 0, 0, 0, 90, 505, 1, 0, 0, 0, 92, 512, 1, 0, 0,
		0, 94, 532, 1, 0, 0, 0, 96, 534, 1, 0, 0, 0, 98, 549, 1, 0, 0, 0, 100,
		564, 1, 0, 0, 0, 102, 566, 1, 0, 0, 0, 104, 584, 1, 0, 0, 0, 106, 586,
		1, 0, 0, 0, 108, 590, 1, 0, 0, 0, 110, 111, 3, 2, 1, 0, 111, 112, 5, 0,
		0, 1, 112, 1, 1, 0, 0, 0, 113, 114, 3, 4, 2, 0, 114, 115, 3, 22, 11, 0,
		115, 3, 1, 0, 0, 0, 116, 118, 3, 6, 3, 0, 117, 116, 1, 0, 0, 0, 118, 121,
		1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 125, 1, 0,
		0, 0, 121, 119, 1, 0, 0, 0, 122, 124, 3, 20, 10, 0, 123, 122, 1, 0, 0,
		0, 124, 127, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126,
		5, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128, 132, 3, 8, 4, 0, 129, 132, 3,
		10, 5, 0, 130, 132, 3, 12, 6, 0, 131, 128, 1, 0, 0, 0, 131, 129, 1, 0,
		0, 0, 131, 130, 1, 0, 0, 0, 132, 7, 1, 0, 0, 0, 133, 138, 5, 1, 0, 0, 134,
		135, 5, 53, 0, 0, 135, 136, 3, 14, 7, 0, 136, 137, 5, 54, 0, 0, 137, 139,
		1, 0, 0, 0, 138, 134, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 141, 5, 45, 0, 0, 141, 9, 1, 0, 0, 0, 142, 143, 5, 2, 0, 0,
		143, 146, 5, 60, 0, 0, 144, 145, 5, 25, 0, 0, 145, 147, 3, 72, 36, 0, 146,
		144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 152, 1, 0, 0, 0, 148, 149,
		5, 53, 0, 0, 149, 150, 3, 14, 7, 0, 150, 151, 5, 54, 0, 0, 151, 153, 1,
		0, 0, 0, 152, 148, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 1, 0, 0,
		0, 154, 155, 5, 45, 0, 0, 155, 11, 1, 0, 0, 0, 156, 157, 5, 3, 0, 0, 157,
		160, 5, 60, 0, 0, 158, 159, 5, 25, 0, 0, 159, 161, 3, 72, 36, 0, 160, 158,
		1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 166, 1, 0, 0, 0, 162, 163, 5, 53,
		0, 0, 163, 164, 3, 14, 7, 0, 164, 165, 5, 54, 0, 0, 165, 167, 1, 0, 0,
		0, 166, 162, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168,
		169, 5, 45, 0, 0, 169, 13, 1, 0, 0, 0, 170, 175, 3, 16, 8, 0, 171, 172,
		7, 0, 0, 0, 172, 174, 3, 16, 8, 0, 173, 171, 1, 0, 0, 0, 174, 177, 1, 0,
		0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0,
		177, 175, 1, 0, 0, 0, 178, 180, 7, 0, 0, 0, 179, 178, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 15, 1, 0, 0, 0, 181, 184, 3, 108, 54, 0, 182, 183,
		5, 47, 0, 0, 183, 185, 3, 18, 9, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1,
		0, 0, 0, 185, 17, 1, 0, 0, 0, 186, 187, 7, 1, 0, 0, 187, 19, 1, 0, 0, 0,
		188, 190, 5, 14, 0, 0, 189, 191, 5, 60, 0, 0, 190, 189, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 5, 66, 0, 0, 193, 194,
		5, 45, 0, 0, 194, 21, 1, 0, 0, 0, 195, 197, 3, 24, 12, 0, 196, 195, 1,
		0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 201, 1, 0, 0, 0, 198, 200, 5, 45, 0,
		0, 199, 198, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201,
		202, 1, 0, 0, 0, 202, 23, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 204, 213, 3,
		26, 13, 0, 205, 207, 5, 45, 0, 0, 206, 205, 1, 0, 0, 0, 207, 208, 1, 0,
		0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0,
		210, 212, 3, 26, 13, 0, 211, 206, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213,
		211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 25, 1, 0, 0, 0, 215, 213, 1,
		0, 0, 0, 216, 233, 3, 28, 14, 0, 217, 233, 3, 34, 17, 0, 218, 233, 3, 36,
		18, 0, 219, 233, 3, 40, 20, 0, 220, 233, 3, 44, 22, 0, 221, 233, 3, 46,
		23, 0, 222, 233, 3, 52, 26, 0, 223, 233, 3, 56, 28, 0, 224, 233, 3, 72,
		36, 0, 225, 233, 3, 58, 29, 0, 226, 233, 3, 60, 30, 0, 227, 233, 3, 62,
		31, 0, 228, 233, 3, 64, 32, 0, 229, 233, 3, 66, 33, 0, 230, 233, 3, 68,
		34, 0, 231, 233, 3, 70, 35, 0, 232, 216, 1, 0, 0, 0, 232, 217, 1, 0, 0,
		0, 232, 218, 1, 0, 0, 0, 232, 219, 1, 0, 0, 0, 232, 220, 1, 0, 0, 0, 232,
		221, 1, 0, 0, 0, 232, 222, 1, 0, 0, 0, 232, 223, 1, 0, 0, 0, 232, 224,
		1, 0, 0, 0, 232, 225, 1, 0, 0, 0, 232, 226, 1, 0, 0, 0, 232, 227, 1, 0,
		0, 0, 232, 228, 1, 0, 0, 0, 232, 229, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0,
		232, 231, 1, 0, 0, 0, 233, 27, 1, 0, 0, 0, 234, 235, 3, 30, 15, 0, 235,
		236, 5, 25, 0, 0, 236, 237, 3, 32, 16, 0, 237, 29, 1, 0, 0, 0, 238, 243,
		3, 86, 43, 0, 239, 240, 5, 46, 0, 0, 240, 242, 3, 86, 43, 0, 241, 239,
		1, 0, 0, 0, 242, 245, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0,
		0, 0, 244, 31, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 251, 3, 72, 36, 0,
		247, 248, 5, 46, 0, 0, 248, 250, 3, 72, 36, 0, 249, 247, 1, 0, 0, 0, 250,
		253, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 33, 1,
		0, 0, 0, 253, 251, 1, 0, 0, 0, 254, 255, 3, 86, 43, 0, 255, 256, 7, 2,
		0, 0, 256, 257, 3, 72, 36, 0, 257, 35, 1, 0, 0, 0, 258, 259, 5, 22, 0,
		0, 259, 262, 3, 38, 19, 0, 260, 261, 5, 25, 0, 0, 261, 263, 3, 32, 16,
		0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 37, 1, 0, 0, 0, 264,
		269, 5, 60, 0, 0, 265, 266, 5, 46, 0, 0, 266, 268, 5, 60, 0, 0, 267, 265,
		1, 0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0,
		0, 0, 270, 39, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 273, 5, 11, 0, 0,
		273, 274, 3, 42, 21, 0, 274, 275, 5, 60, 0, 0, 275, 276, 3, 72, 36, 0,
		276, 277, 5, 53, 0, 0, 277, 278, 3, 22, 11, 0, 278, 279, 5, 54, 0, 0, 279,
		41, 1, 0, 0, 0, 280, 285, 5, 60, 0, 0, 281, 282, 5, 46, 0, 0, 282, 284,
		5, 60, 0, 0, 283, 281, 1, 0, 0, 0, 284, 287, 1, 0, 0, 0, 285, 283, 1, 0,
		0, 0, 285, 286, 1, 0, 0, 0, 286, 43, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0,
		288, 289, 5, 23, 0, 0, 289, 290, 3, 72, 36, 0, 290, 291, 5, 53, 0, 0, 291,
		292, 3, 22, 11, 0, 292, 293, 5, 54, 0, 0, 293, 45, 1, 0, 0, 0, 294, 295,
		5, 13, 0, 0, 295, 296, 3, 72, 36, 0, 296, 297, 5, 53, 0, 0, 297, 298, 3,
		22, 11, 0, 298, 302, 5, 54, 0, 0, 299, 301, 3, 48, 24, 0, 300, 299, 1,
		0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0, 0,
		0, 303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 307, 3, 50, 25, 0,
		306, 305, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 47, 1, 0, 0, 0, 308, 309,
		5, 9, 0, 0, 309, 310, 5, 13, 0, 0, 310, 311, 3, 72, 36, 0, 311, 312, 5,
		53, 0, 0, 312, 313, 3, 22, 11, 0, 313, 314, 5, 54, 0, 0, 314, 49, 1, 0,
		0, 0, 315, 316, 5, 9, 0, 0, 316, 317, 5, 53, 0, 0, 317, 318, 3, 22, 11,
		0, 318, 319, 5, 54, 0, 0, 319, 51, 1, 0, 0, 0, 320, 321, 5, 12, 0, 0, 321,
		322, 5, 60, 0, 0, 322, 324, 5, 49, 0, 0, 323, 325, 3, 54, 27, 0, 324, 323,
		1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327, 5, 50,
		0, 0, 327, 328, 5, 53, 0, 0, 328, 329, 3, 22, 11, 0, 329, 330, 5, 54, 0,
		0, 330, 53, 1, 0, 0, 0, 331, 336, 5, 60, 0, 0, 332, 333, 5, 46, 0, 0, 333,
		335, 5, 60, 0, 0, 334, 332, 1, 0, 0, 0, 335, 338, 1, 0, 0, 0, 336, 334,
		1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 55, 1, 0, 0, 0, 338, 336, 1, 0,
		0, 0, 339, 341, 5, 18, 0, 0, 340, 342, 3, 32, 16, 0, 341, 340, 1, 0, 0,
		0, 341, 342, 1, 0, 0, 0, 342, 57, 1, 0, 0, 0, 343, 344, 5, 21, 0, 0, 344,
		345, 5, 53, 0, 0, 345, 346, 3, 22, 11, 0, 346, 347, 5, 54, 0, 0, 347, 349,
		5, 6, 0, 0, 348, 350, 5, 60, 0, 0, 349, 348, 1, 0, 0, 0, 349, 350, 1, 0,
		0, 0, 350, 351, 1, 0, 0, 0, 351, 352, 5, 53, 0, 0, 352, 353, 3, 22, 11,
		0, 353, 354, 5, 54, 0, 0, 354, 59, 1, 0, 0, 0, 355, 356, 5, 19, 0, 0, 356,
		357, 3, 72, 36, 0, 357, 61, 1, 0, 0, 0, 358, 359, 5, 8, 0, 0, 359, 360,
		3, 80, 40, 0, 360, 63, 1, 0, 0, 0, 361, 362, 5, 24, 0, 0, 362, 363, 3,
		32, 16, 0, 363, 65, 1, 0, 0, 0, 364, 365, 5, 5, 0, 0, 365, 67, 1, 0, 0,
		0, 366, 367, 5, 7, 0, 0, 367, 69, 1, 0, 0, 0, 368, 369, 3, 86, 43, 0, 369,
		370, 7, 3, 0, 0, 370, 71, 1, 0, 0, 0, 371, 374, 3, 90, 45, 0, 372, 374,
		3, 74, 37, 0, 373, 371, 1, 0, 0, 0, 373, 372, 1, 0, 0, 0, 374, 73, 1, 0,
		0, 0, 375, 376, 6, 37, -1, 0, 376, 377, 3, 76, 38, 0, 377, 386, 1, 0, 0,
		0, 378, 379, 10, 2, 0, 0, 379, 380, 5, 44, 0, 0, 380, 381, 3, 74, 37, 0,
		381, 382, 5, 47, 0, 0, 382, 383, 3, 74, 37, 2, 383, 385, 1, 0, 0, 0, 384,
		378, 1, 0, 0, 0, 385, 388, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 386, 387,
		1, 0, 0, 0, 387, 75, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 389, 390, 6, 38,
		-1, 0, 390, 391, 3, 78, 39, 0, 391, 412, 1, 0, 0, 0, 392, 393, 10, 6, 0,
		0, 393, 394, 7, 4, 0, 0, 394, 411, 3, 76, 38, 7, 395, 396, 10, 5, 0, 0,
		396, 397, 7, 5, 0, 0, 397, 411, 3, 76, 38, 6, 398, 399, 10, 4, 0, 0, 399,
		400, 5, 57, 0, 0, 400, 411, 3, 76, 38, 5, 401, 402, 10, 3, 0, 0, 402, 403,
		7, 6, 0, 0, 403, 411, 3, 76, 38, 4, 404, 405, 10, 2, 0, 0, 405, 406, 5,
		4, 0, 0, 406, 411, 3, 76, 38, 3, 407, 408, 10, 1, 0, 0, 408, 409, 5, 17,
		0, 0, 409, 411, 3, 76, 38, 2, 410, 392, 1, 0, 0, 0, 410, 395, 1, 0, 0,
		0, 410, 398, 1, 0, 0, 0, 410, 401, 1, 0, 0, 0, 410, 404, 1, 0, 0, 0, 410,
		407, 1, 0, 0, 0, 411, 414, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412, 413,
		1, 0, 0, 0, 413, 77, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 415, 416, 5, 16,
		0, 0, 416, 423, 3, 80, 40, 0, 417, 418, 5, 37, 0, 0, 418, 423, 3, 80, 40,
		0, 419, 420, 5, 38, 0, 0, 420, 423, 3, 80, 40, 0, 421, 423, 3, 80, 40,
		0, 422, 415, 1, 0, 0, 0, 422, 417, 1, 0, 0, 0, 422, 419, 1, 0, 0, 0, 422,
		421, 1, 0, 0, 0, 423, 79, 1, 0, 0, 0, 424, 425, 6, 40, -1, 0, 425, 437,
		5, 60, 0, 0, 426, 437, 3, 88, 44, 0, 427, 437, 3, 92, 46, 0, 428, 437,
		3, 96, 48, 0, 429, 437, 3, 102, 51, 0, 430, 437, 3, 82, 41, 0, 431, 437,
		5, 61, 0, 0, 432, 433, 5, 49, 0, 0, 433, 434, 3, 72, 36, 0, 434, 435, 5,
		50, 0, 0, 435, 437, 1, 0, 0, 0, 436, 424, 1, 0, 0, 0, 436, 426, 1, 0, 0,
		0, 436, 427, 1, 0, 0, 0, 436, 428, 1, 0, 0, 0, 436, 429, 1, 0, 0, 0, 436,
		430, 1, 0, 0, 0, 436, 431, 1, 0, 0, 0, 436, 432, 1, 0, 0, 0, 437, 467,
		1, 0, 0, 0, 438, 439, 10, 11, 0, 0, 439, 440, 5, 48, 0, 0, 440, 466, 5,
		60, 0, 0, 441, 442, 10, 10, 0, 0, 442, 443, 5, 51, 0, 0, 443, 444, 3, 72,
		36, 0, 444, 445, 5, 52, 0, 0, 445, 466, 1, 0, 0, 0, 446, 447, 10, 9, 0,
		0, 447, 449, 5, 51, 0, 0, 448, 450, 3, 72, 36, 0, 449, 448, 1, 0, 0, 0,
		449, 450, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 453, 5, 47, 0, 0, 452,
		454, 3, 72, 36, 0, 453, 452, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454, 455,
		1, 0, 0, 0, 455, 466, 5, 52, 0, 0, 456, 457, 10, 8, 0, 0, 457, 462, 5,
		49, 0, 0, 458, 460, 3, 84, 42, 0, 459, 461, 5, 58, 0, 0, 460, 459, 1, 0,
		0, 0, 460, 461, 1, 0, 0, 0, 461, 463, 1, 0, 0, 0, 462, 458, 1, 0, 0, 0,
		462, 463, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 466, 5, 50, 0, 0, 465,
		438, 1, 0, 0, 0, 465, 441, 1, 0, 0, 0, 465, 446, 1, 0, 0, 0, 465, 456,
		1, 0, 0, 0, 466, 469, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 467, 468, 1, 0,
		0, 0, 468, 81, 1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 470, 471, 7, 1, 0, 0,
		471, 83, 1, 0, 0, 0, 472, 477, 3, 72, 36, 0, 473, 474, 5, 46, 0, 0, 474,
		476, 3, 72, 36, 0, 475, 473, 1, 0, 0, 0, 476, 479, 1, 0, 0, 0, 477, 475,
		1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 481, 1, 0, 0, 0, 479, 477, 1, 0,
		0, 0, 480, 482, 7, 0, 0, 0, 481, 480, 1, 0, 0, 0, 481, 482, 1, 0, 0, 0,
		482, 85, 1, 0, 0, 0, 483, 494, 5, 60, 0, 0, 484, 485, 3, 80, 40, 0, 485,
		486, 5, 48, 0, 0, 486, 487, 5, 60, 0, 0, 487, 494, 1, 0, 0, 0, 488, 489,
		3, 80, 40, 0, 489, 490, 5, 51, 0, 0, 490, 491, 3, 72, 36, 0, 491, 492,
		5, 52, 0, 0, 492, 494, 1, 0, 0, 0, 493, 483, 1, 0, 0, 0, 493, 484, 1, 0,
		0, 0, 493, 488, 1, 0, 0, 0, 494, 87, 1, 0, 0, 0, 495, 496, 5, 12, 0, 0,
		496, 498, 5, 49, 0, 0, 497, 499, 3, 54, 27, 0, 498, 497, 1, 0, 0, 0, 498,
		499, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 501, 5, 50, 0, 0, 501, 502,
		5, 53, 0, 0, 502, 503, 3, 22, 11, 0, 503, 504, 5, 54, 0, 0, 504, 89, 1,
		0, 0, 0, 505, 507, 5, 56, 0, 0, 506, 508, 3, 54, 27, 0, 507, 506, 1, 0,
		0, 0, 507, 508, 1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 510, 5, 55, 0, 0,
		510, 511, 3, 74, 37, 0, 511, 91, 1, 0, 0, 0, 512, 514, 5, 67, 0, 0, 513,
		515, 3, 94, 47, 0, 514, 513, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 514,
		1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 519, 5, 70,
		0, 0, 519, 93, 1, 0, 0, 0, 520, 533, 5, 71, 0, 0, 521, 533, 5, 72, 0, 0,
		522, 533, 5, 73, 0, 0, 523, 533, 5, 68, 0, 0, 524, 533, 5, 69, 0, 0, 525,
		526, 5, 53, 0, 0, 526, 528, 3, 72, 36, 0, 527, 529, 5, 58, 0, 0, 528, 527,
		1, 0, 0, 0, 528, 529, 1, 0, 0, 0, 529, 530, 1, 0, 0, 0, 530, 531, 5, 54,
		0, 0, 531, 533, 1, 0, 0, 0, 532, 520, 1, 0, 0, 0, 532, 521, 1, 0, 0, 0,
		532, 522, 1, 0, 0, 0, 532, 523, 1, 0, 0, 0, 532, 524, 1, 0, 0, 0, 532,
		525, 1, 0, 0, 0, 533, 95, 1, 0, 0, 0, 534, 535, 5, 53, 0, 0, 535, 536,
		3, 98, 49, 0, 536, 537, 5, 54, 0, 0, 537, 97, 1, 0, 0, 0, 538, 543, 3,
		100, 50, 0, 539, 540, 7, 0, 0, 0, 540, 542, 3, 100, 50, 0, 541, 539, 1,
		0, 0, 0, 542, 545, 1, 0, 0, 0, 543, 541, 1, 0, 0, 0, 543, 544, 1, 0, 0,
		0, 544, 547, 1, 0, 0, 0, 545, 543, 1, 0, 0, 0, 546, 548, 7, 0, 0, 0, 547,
		546, 1, 0, 0, 0, 547, 548, 1, 0, 0, 0, 548, 550, 1, 0, 0, 0, 549, 538,
		1, 0, 0, 0, 549, 550, 1, 0, 0, 0, 550, 99, 1, 0, 0, 0, 551, 552, 3, 108,
		54, 0, 552, 553, 5, 47, 0, 0, 553, 554, 3, 72, 36, 0, 554, 565, 1, 0, 0,
		0, 555, 556, 5, 51, 0, 0, 556, 557, 3, 72, 36, 0, 557, 558, 5, 52, 0, 0,
		558, 559, 5, 47, 0, 0, 559, 560, 3, 72, 36, 0, 560, 565, 1, 0, 0, 0, 561,
		562, 3, 80, 40, 0, 562, 563, 5, 58, 0, 0, 563, 565, 1, 0, 0, 0, 564, 551,
		1, 0, 0, 0, 564, 555, 1, 0, 0, 0, 564, 561, 1, 0, 0, 0, 565, 101, 1, 0,
		0, 0, 566, 567, 5, 51, 0, 0, 567, 568, 3, 104, 52, 0, 568, 569, 5, 52,
		0, 0, 569, 103, 1, 0, 0, 0, 570, 575, 3, 106, 53, 0, 571, 572, 7, 0, 0,
		0, 572, 574, 3, 106, 53, 0, 573, 571, 1, 0, 0, 0, 574, 577, 1, 0, 0, 0,
		575, 573, 1, 0, 0, 0, 575, 576, 1, 0, 0, 0, 576, 581, 1, 0, 0, 0, 577,
		575, 1, 0, 0, 0, 578, 580, 7, 0, 0, 0, 579, 578, 1, 0, 0, 0, 580, 583,
		1, 0, 0, 0, 581, 579, 1, 0, 0, 0, 581, 582, 1, 0, 0, 0, 582, 585, 1, 0,
		0, 0, 583, 581, 1, 0, 0, 0, 584, 570, 1, 0, 0, 0, 584, 585, 1, 0, 0, 0,
		585, 105, 1, 0, 0, 0, 586, 588, 3, 72, 36, 0, 587, 589, 5, 58, 0, 0, 588,
		587, 1, 0, 0, 0, 588, 589, 1, 0, 0, 0, 589, 107, 1, 0, 0, 0, 590, 591,
		7, 7, 0, 0, 591, 109, 1, 0, 0, 0, 56, 119, 125, 131, 138, 146, 152, 160,
		166, 175, 179, 184, 190, 196, 201, 208, 213, 232, 243, 251, 262, 269, 285,
		302, 306, 324, 336, 341, 349, 373, 386, 410, 412, 422, 436, 449, 453, 460,
		462, 465, 467, 477, 481, 493, 498, 507, 516, 528, 532, 543, 547, 549, 564,
		575, 581, 584, 588,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// NitroParserInit initializes any static state used to implement NitroParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNitroParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NitroParserInit() {
	staticData := &nitroparserParserStaticData
	staticData.once.Do(nitroparserParserInit)
}

// NewNitroParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNitroParser(input antlr.TokenStream) *NitroParser {
	NitroParserInit()
	this := new(NitroParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &nitroparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "NitroParser.g4"

	return this
}

// NitroParser tokens.
const (
	NitroParserEOF                 = antlr.TokenEOF
	NitroParserM_INFO              = 1
	NitroParserM_PARAM             = 2
	NitroParserM_FLAG              = 3
	NitroParserAND                 = 4
	NitroParserBREAK               = 5
	NitroParserCATCH               = 6
	NitroParserCONTINUE            = 7
	NitroParserDEFER               = 8
	NitroParserELSE                = 9
	NitroParserFALSE               = 10
	NitroParserFOR                 = 11
	NitroParserFUNC                = 12
	NitroParserIF                  = 13
	NitroParserIMPORT              = 14
	NitroParserNIL                 = 15
	NitroParserNOT                 = 16
	NitroParserOR                  = 17
	NitroParserRETURN              = 18
	NitroParserTHROW               = 19
	NitroParserTRUE                = 20
	NitroParserTRY                 = 21
	NitroParserVAR                 = 22
	NitroParserWHILE               = 23
	NitroParserYIELD               = 24
	NitroParserASSIGN              = 25
	NitroParserASSIGN_ADD          = 26
	NitroParserASSIGN_SUB          = 27
	NitroParserASSIGN_MUL          = 28
	NitroParserASSIGN_DIV          = 29
	NitroParserASSIGN_MOD          = 30
	NitroParserEQ                  = 31
	NitroParserNE                  = 32
	NitroParserLT                  = 33
	NitroParserLE                  = 34
	NitroParserGT                  = 35
	NitroParserGE                  = 36
	NitroParserADD                 = 37
	NitroParserSUB                 = 38
	NitroParserMUL                 = 39
	NitroParserDIV                 = 40
	NitroParserMOD                 = 41
	NitroParserINC                 = 42
	NitroParserDEC                 = 43
	NitroParserQUESTION_MARK       = 44
	NitroParserSEMICOLON           = 45
	NitroParserCOMMA               = 46
	NitroParserCOLON               = 47
	NitroParserPERIOD              = 48
	NitroParserOPAREN              = 49
	NitroParserCPAREN              = 50
	NitroParserOBRACKET            = 51
	NitroParserCBRACKET            = 52
	NitroParserOCURLY              = 53
	NitroParserCCURLY              = 54
	NitroParserARROW               = 55
	NitroParserLAMBDA              = 56
	NitroParserPIPE                = 57
	NitroParserEXPAND              = 58
	NitroParserNUMBER              = 59
	NitroParserID                  = 60
	NitroParserREGEX               = 61
	NitroParserNEWLINE             = 62
	NitroParserCHAR                = 63
	NitroParserCOMMENT             = 64
	NitroParserWS                  = 65
	NitroParserSTRING              = 66
	NitroParserEXEC_PREFIX         = 67
	NitroParserEXEC_WS             = 68
	NitroParserEXEC_HOME           = 69
	NitroParserEXEC_SUFFIX         = 70
	NitroParserEXEC_LITERAL        = 71
	NitroParserEXEC_DQUOTE_LITERAL = 72
	NitroParserEXEC_SQUOTE_LITERAL = 73
	NitroParserEXEC_EXPR_WS_       = 74
	NitroParserDQUOTE_             = 75
	NitroParserEXEC_SQUOTE_        = 76
)

// NitroParser rules.
const (
	NitroParserRULE_start              = 0
	NitroParserRULE_unit               = 1
	NitroParserRULE_prologue           = 2
	NitroParserRULE_meta_directive     = 3
	NitroParserRULE_meta_info          = 4
	NitroParserRULE_meta_param         = 5
	NitroParserRULE_meta_flag          = 6
	NitroParserRULE_meta_attribs       = 7
	NitroParserRULE_meta_attrib        = 8
	NitroParserRULE_meta_literal       = 9
	NitroParserRULE_import_stmt        = 10
	NitroParserRULE_stmts              = 11
	NitroParserRULE_stmt_list          = 12
	NitroParserRULE_stmt               = 13
	NitroParserRULE_assignment_stmt    = 14
	NitroParserRULE_assignment_lvalues = 15
	NitroParserRULE_rvalues            = 16
	NitroParserRULE_assignment_op_stmt = 17
	NitroParserRULE_var_decl_stmt      = 18
	NitroParserRULE_var_decl_vars      = 19
	NitroParserRULE_for_stmt           = 20
	NitroParserRULE_for_vars           = 21
	NitroParserRULE_while_stmt         = 22
	NitroParserRULE_if_stmt            = 23
	NitroParserRULE_if_elif            = 24
	NitroParserRULE_if_else            = 25
	NitroParserRULE_func_stmt          = 26
	NitroParserRULE_param_list         = 27
	NitroParserRULE_return_stmt        = 28
	NitroParserRULE_try_catch_stmt     = 29
	NitroParserRULE_throw_stmt         = 30
	NitroParserRULE_defer_stmt         = 31
	NitroParserRULE_yield_stmt         = 32
	NitroParserRULE_break_stmt         = 33
	NitroParserRULE_continue_stmt      = 34
	NitroParserRULE_inc_dec_stmt       = 35
	NitroParserRULE_expr               = 36
	NitroParserRULE_expr3              = 37
	NitroParserRULE_binary_expr        = 38
	NitroParserRULE_unary_expr         = 39
	NitroParserRULE_primary_expr       = 40
	NitroParserRULE_simple_literal     = 41
	NitroParserRULE_arg_list           = 42
	NitroParserRULE_lvalue_expr        = 43
	NitroParserRULE_lambda_expr        = 44
	NitroParserRULE_short_lambda_expr  = 45
	NitroParserRULE_exec_expr          = 46
	NitroParserRULE_exec_expr_arg      = 47
	NitroParserRULE_object_literal     = 48
	NitroParserRULE_object_fields      = 49
	NitroParserRULE_object_field       = 50
	NitroParserRULE_array_literal      = 51
	NitroParserRULE_array_elems        = 52
	NitroParserRULE_array_elem         = 53
	NitroParserRULE_id_or_keyword      = 54
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Unit() IUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(NitroParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *NitroParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NitroParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Unit()
	}
	{
		p.SetState(111)
		p.Match(NitroParserEOF)
	}

	return localctx
}

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_unit
	return p
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) Prologue() IPrologueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrologueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrologueContext)
}

func (s *UnitContext) Stmts() IStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (p *NitroParser) Unit() (localctx IUnitContext) {
	this := p
	_ = this

	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NitroParserRULE_unit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Prologue()
	}
	{
		p.SetState(114)
		p.Stmts()
	}

	return localctx
}

// IPrologueContext is an interface to support dynamic dispatch.
type IPrologueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrologueContext differentiates from other interfaces.
	IsPrologueContext()
}

type PrologueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrologueContext() *PrologueContext {
	var p = new(PrologueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_prologue
	return p
}

func (*PrologueContext) IsPrologueContext() {}

func NewPrologueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrologueContext {
	var p = new(PrologueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_prologue

	return p
}

func (s *PrologueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrologueContext) AllMeta_directive() []IMeta_directiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMeta_directiveContext); ok {
			len++
		}
	}

	tst := make([]IMeta_directiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMeta_directiveContext); ok {
			tst[i] = t.(IMeta_directiveContext)
			i++
		}
	}

	return tst
}

func (s *PrologueContext) Meta_directive(i int) IMeta_directiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_directiveContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_directiveContext)
}

func (s *PrologueContext) AllImport_stmt() []IImport_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImport_stmtContext); ok {
			len++
		}
	}

	tst := make([]IImport_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImport_stmtContext); ok {
			tst[i] = t.(IImport_stmtContext)
			i++
		}
	}

	return tst
}

func (s *PrologueContext) Import_stmt(i int) IImport_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImport_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImport_stmtContext)
}

func (s *PrologueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrologueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrologueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrologue(s)
	}
}

func (s *PrologueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrologue(s)
	}
}

func (p *NitroParser) Prologue() (localctx IPrologueContext) {
	this := p
	_ = this

	localctx = NewPrologueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NitroParserRULE_prologue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserM_INFO)|(1<<NitroParserM_PARAM)|(1<<NitroParserM_FLAG))) != 0 {
		{
			p.SetState(116)
			p.Meta_directive()
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserIMPORT {
		{
			p.SetState(122)
			p.Import_stmt()
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMeta_directiveContext is an interface to support dynamic dispatch.
type IMeta_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_directiveContext differentiates from other interfaces.
	IsMeta_directiveContext()
}

type Meta_directiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_directiveContext() *Meta_directiveContext {
	var p = new(Meta_directiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_directive
	return p
}

func (*Meta_directiveContext) IsMeta_directiveContext() {}

func NewMeta_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_directiveContext {
	var p = new(Meta_directiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_directive

	return p
}

func (s *Meta_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_directiveContext) Meta_info() IMeta_infoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_infoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_infoContext)
}

func (s *Meta_directiveContext) Meta_param() IMeta_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_paramContext)
}

func (s *Meta_directiveContext) Meta_flag() IMeta_flagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_flagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_flagContext)
}

func (s *Meta_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_directiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_directive(s)
	}
}

func (s *Meta_directiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_directive(s)
	}
}

func (p *NitroParser) Meta_directive() (localctx IMeta_directiveContext) {
	this := p
	_ = this

	localctx = NewMeta_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NitroParserRULE_meta_directive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserM_INFO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Meta_info()
		}

	case NitroParserM_PARAM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Meta_param()
		}

	case NitroParserM_FLAG:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Meta_flag()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMeta_infoContext is an interface to support dynamic dispatch.
type IMeta_infoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_infoContext differentiates from other interfaces.
	IsMeta_infoContext()
}

type Meta_infoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_infoContext() *Meta_infoContext {
	var p = new(Meta_infoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_info
	return p
}

func (*Meta_infoContext) IsMeta_infoContext() {}

func NewMeta_infoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_infoContext {
	var p = new(Meta_infoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_info

	return p
}

func (s *Meta_infoContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_infoContext) M_INFO() antlr.TerminalNode {
	return s.GetToken(NitroParserM_INFO, 0)
}

func (s *Meta_infoContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Meta_infoContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Meta_infoContext) Meta_attribs() IMeta_attribsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_attribsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_attribsContext)
}

func (s *Meta_infoContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *Meta_infoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_infoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_infoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_info(s)
	}
}

func (s *Meta_infoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_info(s)
	}
}

func (p *NitroParser) Meta_info() (localctx IMeta_infoContext) {
	this := p
	_ = this

	localctx = NewMeta_infoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NitroParserRULE_meta_info)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(NitroParserM_INFO)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserOCURLY {
		{
			p.SetState(134)
			p.Match(NitroParserOCURLY)
		}
		{
			p.SetState(135)
			p.Meta_attribs()
		}
		{
			p.SetState(136)
			p.Match(NitroParserCCURLY)
		}

	}
	{
		p.SetState(140)
		p.Match(NitroParserSEMICOLON)
	}

	return localctx
}

// IMeta_paramContext is an interface to support dynamic dispatch.
type IMeta_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_paramContext differentiates from other interfaces.
	IsMeta_paramContext()
}

type Meta_paramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_paramContext() *Meta_paramContext {
	var p = new(Meta_paramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_param
	return p
}

func (*Meta_paramContext) IsMeta_paramContext() {}

func NewMeta_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_paramContext {
	var p = new(Meta_paramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_param

	return p
}

func (s *Meta_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_paramContext) M_PARAM() antlr.TerminalNode {
	return s.GetToken(NitroParserM_PARAM, 0)
}

func (s *Meta_paramContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Meta_paramContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Meta_paramContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN, 0)
}

func (s *Meta_paramContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Meta_paramContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Meta_paramContext) Meta_attribs() IMeta_attribsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_attribsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_attribsContext)
}

func (s *Meta_paramContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *Meta_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_param(s)
	}
}

func (s *Meta_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_param(s)
	}
}

func (p *NitroParser) Meta_param() (localctx IMeta_paramContext) {
	this := p
	_ = this

	localctx = NewMeta_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NitroParserRULE_meta_param)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(NitroParserM_PARAM)
	}
	{
		p.SetState(143)
		p.Match(NitroParserID)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserASSIGN {
		{
			p.SetState(144)
			p.Match(NitroParserASSIGN)
		}
		{
			p.SetState(145)
			p.Expr()
		}

	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserOCURLY {
		{
			p.SetState(148)
			p.Match(NitroParserOCURLY)
		}
		{
			p.SetState(149)
			p.Meta_attribs()
		}
		{
			p.SetState(150)
			p.Match(NitroParserCCURLY)
		}

	}
	{
		p.SetState(154)
		p.Match(NitroParserSEMICOLON)
	}

	return localctx
}

// IMeta_flagContext is an interface to support dynamic dispatch.
type IMeta_flagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_flagContext differentiates from other interfaces.
	IsMeta_flagContext()
}

type Meta_flagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_flagContext() *Meta_flagContext {
	var p = new(Meta_flagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_flag
	return p
}

func (*Meta_flagContext) IsMeta_flagContext() {}

func NewMeta_flagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_flagContext {
	var p = new(Meta_flagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_flag

	return p
}

func (s *Meta_flagContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_flagContext) M_FLAG() antlr.TerminalNode {
	return s.GetToken(NitroParserM_FLAG, 0)
}

func (s *Meta_flagContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Meta_flagContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Meta_flagContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN, 0)
}

func (s *Meta_flagContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Meta_flagContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Meta_flagContext) Meta_attribs() IMeta_attribsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_attribsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_attribsContext)
}

func (s *Meta_flagContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *Meta_flagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_flagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_flagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_flag(s)
	}
}

func (s *Meta_flagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_flag(s)
	}
}

func (p *NitroParser) Meta_flag() (localctx IMeta_flagContext) {
	this := p
	_ = this

	localctx = NewMeta_flagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NitroParserRULE_meta_flag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(NitroParserM_FLAG)
	}
	{
		p.SetState(157)
		p.Match(NitroParserID)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserASSIGN {
		{
			p.SetState(158)
			p.Match(NitroParserASSIGN)
		}
		{
			p.SetState(159)
			p.Expr()
		}

	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserOCURLY {
		{
			p.SetState(162)
			p.Match(NitroParserOCURLY)
		}
		{
			p.SetState(163)
			p.Meta_attribs()
		}
		{
			p.SetState(164)
			p.Match(NitroParserCCURLY)
		}

	}
	{
		p.SetState(168)
		p.Match(NitroParserSEMICOLON)
	}

	return localctx
}

// IMeta_attribsContext is an interface to support dynamic dispatch.
type IMeta_attribsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_attribsContext differentiates from other interfaces.
	IsMeta_attribsContext()
}

type Meta_attribsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_attribsContext() *Meta_attribsContext {
	var p = new(Meta_attribsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_attribs
	return p
}

func (*Meta_attribsContext) IsMeta_attribsContext() {}

func NewMeta_attribsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_attribsContext {
	var p = new(Meta_attribsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_attribs

	return p
}

func (s *Meta_attribsContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_attribsContext) AllMeta_attrib() []IMeta_attribContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMeta_attribContext); ok {
			len++
		}
	}

	tst := make([]IMeta_attribContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMeta_attribContext); ok {
			tst[i] = t.(IMeta_attribContext)
			i++
		}
	}

	return tst
}

func (s *Meta_attribsContext) Meta_attrib(i int) IMeta_attribContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_attribContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_attribContext)
}

func (s *Meta_attribsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCOMMA)
}

func (s *Meta_attribsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCOMMA, i)
}

func (s *Meta_attribsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(NitroParserSEMICOLON)
}

func (s *Meta_attribsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, i)
}

func (s *Meta_attribsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_attribsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_attribsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_attribs(s)
	}
}

func (s *Meta_attribsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_attribs(s)
	}
}

func (p *NitroParser) Meta_attribs() (localctx IMeta_attribsContext) {
	this := p
	_ = this

	localctx = NewMeta_attribsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NitroParserRULE_meta_attribs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Meta_attrib()
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(171)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(172)
				p.Meta_attrib()
			}

		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
		{
			p.SetState(178)
			_la = p.GetTokenStream().LA(1)

			if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IMeta_attribContext is an interface to support dynamic dispatch.
type IMeta_attribContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_attribContext differentiates from other interfaces.
	IsMeta_attribContext()
}

type Meta_attribContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_attribContext() *Meta_attribContext {
	var p = new(Meta_attribContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_attrib
	return p
}

func (*Meta_attribContext) IsMeta_attribContext() {}

func NewMeta_attribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_attribContext {
	var p = new(Meta_attribContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_attrib

	return p
}

func (s *Meta_attribContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_attribContext) Id_or_keyword() IId_or_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_or_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_or_keywordContext)
}

func (s *Meta_attribContext) COLON() antlr.TerminalNode {
	return s.GetToken(NitroParserCOLON, 0)
}

func (s *Meta_attribContext) Meta_literal() IMeta_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeta_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeta_literalContext)
}

func (s *Meta_attribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_attribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_attribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_attrib(s)
	}
}

func (s *Meta_attribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_attrib(s)
	}
}

func (p *NitroParser) Meta_attrib() (localctx IMeta_attribContext) {
	this := p
	_ = this

	localctx = NewMeta_attribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NitroParserRULE_meta_attrib)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Id_or_keyword()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserCOLON {
		{
			p.SetState(182)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(183)
			p.Meta_literal()
		}

	}

	return localctx
}

// IMeta_literalContext is an interface to support dynamic dispatch.
type IMeta_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVal returns the val token.
	GetVal() antlr.Token

	// SetVal sets the val token.
	SetVal(antlr.Token)

	// IsMeta_literalContext differentiates from other interfaces.
	IsMeta_literalContext()
}

type Meta_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	val    antlr.Token
}

func NewEmptyMeta_literalContext() *Meta_literalContext {
	var p = new(Meta_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_literal
	return p
}

func (*Meta_literalContext) IsMeta_literalContext() {}

func NewMeta_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_literalContext {
	var p = new(Meta_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_literal

	return p
}

func (s *Meta_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_literalContext) GetVal() antlr.Token { return s.val }

func (s *Meta_literalContext) SetVal(v antlr.Token) { s.val = v }

func (s *Meta_literalContext) STRING() antlr.TerminalNode {
	return s.GetToken(NitroParserSTRING, 0)
}

func (s *Meta_literalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NitroParserNUMBER, 0)
}

func (s *Meta_literalContext) CHAR() antlr.TerminalNode {
	return s.GetToken(NitroParserCHAR, 0)
}

func (s *Meta_literalContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NitroParserTRUE, 0)
}

func (s *Meta_literalContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NitroParserFALSE, 0)
}

func (s *Meta_literalContext) NIL() antlr.TerminalNode {
	return s.GetToken(NitroParserNIL, 0)
}

func (s *Meta_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_literal(s)
	}
}

func (s *Meta_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_literal(s)
	}
}

func (p *NitroParser) Meta_literal() (localctx IMeta_literalContext) {
	this := p
	_ = this

	localctx = NewMeta_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NitroParserRULE_meta_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Meta_literalContext).val = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserNIL)|(1<<NitroParserTRUE))) != 0) || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(NitroParserNUMBER-59))|(1<<(NitroParserCHAR-59))|(1<<(NitroParserSTRING-59)))) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Meta_literalContext).val = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IImport_stmtContext is an interface to support dynamic dispatch.
type IImport_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImport_stmtContext differentiates from other interfaces.
	IsImport_stmtContext()
}

type Import_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_stmtContext() *Import_stmtContext {
	var p = new(Import_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_import_stmt
	return p
}

func (*Import_stmtContext) IsImport_stmtContext() {}

func NewImport_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_stmtContext {
	var p = new(Import_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_import_stmt

	return p
}

func (s *Import_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_stmtContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(NitroParserIMPORT, 0)
}

func (s *Import_stmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(NitroParserSTRING, 0)
}

func (s *Import_stmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Import_stmtContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Import_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterImport_stmt(s)
	}
}

func (s *Import_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitImport_stmt(s)
	}
}

func (p *NitroParser) Import_stmt() (localctx IImport_stmtContext) {
	this := p
	_ = this

	localctx = NewImport_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NitroParserRULE_import_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(NitroParserIMPORT)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(189)
			p.Match(NitroParserID)
		}

	}
	{
		p.SetState(192)
		p.Match(NitroParserSTRING)
	}
	{
		p.SetState(193)
		p.Match(NitroParserSEMICOLON)
	}

	return localctx
}

// IStmtsContext is an interface to support dynamic dispatch.
type IStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtsContext differentiates from other interfaces.
	IsStmtsContext()
}

type StmtsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtsContext() *StmtsContext {
	var p = new(StmtsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_stmts
	return p
}

func (*StmtsContext) IsStmtsContext() {}

func NewStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtsContext {
	var p = new(StmtsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_stmts

	return p
}

func (s *StmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtsContext) Stmt_list() IStmt_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmt_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmt_listContext)
}

func (s *StmtsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(NitroParserSEMICOLON)
}

func (s *StmtsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, i)
}

func (s *StmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmts(s)
	}
}

func (s *StmtsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmts(s)
	}
}

func (p *NitroParser) Stmts() (localctx IStmtsContext) {
	this := p
	_ = this

	localctx = NewStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NitroParserRULE_stmts)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserBREAK)|(1<<NitroParserCONTINUE)|(1<<NitroParserDEFER)|(1<<NitroParserFALSE)|(1<<NitroParserFOR)|(1<<NitroParserFUNC)|(1<<NitroParserIF)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserRETURN)|(1<<NitroParserTHROW)|(1<<NitroParserTRUE)|(1<<NitroParserTRY)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserYIELD))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37))|(1<<(NitroParserEXEC_PREFIX-37)))) != 0) {
		{
			p.SetState(195)
			p.Stmt_list()
		}

	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserSEMICOLON {
		{
			p.SetState(198)
			p.Match(NitroParserSEMICOLON)
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStmt_listContext is an interface to support dynamic dispatch.
type IStmt_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmt_listContext differentiates from other interfaces.
	IsStmt_listContext()
}

type Stmt_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmt_listContext() *Stmt_listContext {
	var p = new(Stmt_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_stmt_list
	return p
}

func (*Stmt_listContext) IsStmt_listContext() {}

func NewStmt_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stmt_listContext {
	var p = new(Stmt_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_stmt_list

	return p
}

func (s *Stmt_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Stmt_listContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *Stmt_listContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *Stmt_listContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(NitroParserSEMICOLON)
}

func (s *Stmt_listContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, i)
}

func (s *Stmt_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stmt_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_list(s)
	}
}

func (s *Stmt_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_list(s)
	}
}

func (p *NitroParser) Stmt_list() (localctx IStmt_listContext) {
	this := p
	_ = this

	localctx = NewStmt_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NitroParserRULE_stmt_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Stmt()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == NitroParserSEMICOLON {
				{
					p.SetState(205)
					p.Match(NitroParserSEMICOLON)
				}

				p.SetState(208)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(210)
				p.Stmt()
			}

		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyFrom(ctx *StmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Stmt_op_assignContext struct {
	*StmtContext
}

func NewStmt_op_assignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_op_assignContext {
	var p = new(Stmt_op_assignContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_op_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_op_assignContext) Assignment_op_stmt() IAssignment_op_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_op_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_op_stmtContext)
}

func (s *Stmt_op_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_op_assign(s)
	}
}

func (s *Stmt_op_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_op_assign(s)
	}
}

type Stmt_ifContext struct {
	*StmtContext
}

func NewStmt_ifContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_ifContext {
	var p = new(Stmt_ifContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_ifContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *Stmt_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_if(s)
	}
}

func (s *Stmt_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_if(s)
	}
}

type Stmt_breakContext struct {
	*StmtContext
}

func NewStmt_breakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_breakContext {
	var p = new(Stmt_breakContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_breakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_breakContext) Break_stmt() IBreak_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreak_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreak_stmtContext)
}

func (s *Stmt_breakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_break(s)
	}
}

func (s *Stmt_breakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_break(s)
	}
}

type Stmt_inc_decContext struct {
	*StmtContext
}

func NewStmt_inc_decContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_inc_decContext {
	var p = new(Stmt_inc_decContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_inc_decContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_inc_decContext) Inc_dec_stmt() IInc_dec_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInc_dec_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInc_dec_stmtContext)
}

func (s *Stmt_inc_decContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_inc_dec(s)
	}
}

func (s *Stmt_inc_decContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_inc_dec(s)
	}
}

type Stmt_assignmentContext struct {
	*StmtContext
}

func NewStmt_assignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_assignmentContext {
	var p = new(Stmt_assignmentContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_assignmentContext) Assignment_stmt() IAssignment_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_stmtContext)
}

func (s *Stmt_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_assignment(s)
	}
}

func (s *Stmt_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_assignment(s)
	}
}

type Stmt_exprContext struct {
	*StmtContext
}

func NewStmt_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_exprContext {
	var p = new(Stmt_exprContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_exprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Stmt_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_expr(s)
	}
}

func (s *Stmt_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_expr(s)
	}
}

type Stmt_continueContext struct {
	*StmtContext
}

func NewStmt_continueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_continueContext {
	var p = new(Stmt_continueContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_continueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_continueContext) Continue_stmt() IContinue_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinue_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinue_stmtContext)
}

func (s *Stmt_continueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_continue(s)
	}
}

func (s *Stmt_continueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_continue(s)
	}
}

type Stmt_returnContext struct {
	*StmtContext
}

func NewStmt_returnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_returnContext {
	var p = new(Stmt_returnContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_returnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_returnContext) Return_stmt() IReturn_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_stmtContext)
}

func (s *Stmt_returnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_return(s)
	}
}

func (s *Stmt_returnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_return(s)
	}
}

type Stmt_throwContext struct {
	*StmtContext
}

func NewStmt_throwContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_throwContext {
	var p = new(Stmt_throwContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_throwContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_throwContext) Throw_stmt() IThrow_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThrow_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThrow_stmtContext)
}

func (s *Stmt_throwContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_throw(s)
	}
}

func (s *Stmt_throwContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_throw(s)
	}
}

type Stmt_var_decContext struct {
	*StmtContext
}

func NewStmt_var_decContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_var_decContext {
	var p = new(Stmt_var_decContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_var_decContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_var_decContext) Var_decl_stmt() IVar_decl_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_decl_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_decl_stmtContext)
}

func (s *Stmt_var_decContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_var_dec(s)
	}
}

func (s *Stmt_var_decContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_var_dec(s)
	}
}

type Stmt_funcContext struct {
	*StmtContext
}

func NewStmt_funcContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_funcContext {
	var p = new(Stmt_funcContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_funcContext) Func_stmt() IFunc_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_stmtContext)
}

func (s *Stmt_funcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_func(s)
	}
}

func (s *Stmt_funcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_func(s)
	}
}

type Stmt_whileContext struct {
	*StmtContext
}

func NewStmt_whileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_whileContext {
	var p = new(Stmt_whileContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_whileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_whileContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *Stmt_whileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_while(s)
	}
}

func (s *Stmt_whileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_while(s)
	}
}

type Stmt_forContext struct {
	*StmtContext
}

func NewStmt_forContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_forContext {
	var p = new(Stmt_forContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_forContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_forContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *Stmt_forContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_for(s)
	}
}

func (s *Stmt_forContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_for(s)
	}
}

type Stmt_deferContext struct {
	*StmtContext
}

func NewStmt_deferContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_deferContext {
	var p = new(Stmt_deferContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_deferContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_deferContext) Defer_stmt() IDefer_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefer_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefer_stmtContext)
}

func (s *Stmt_deferContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_defer(s)
	}
}

func (s *Stmt_deferContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_defer(s)
	}
}

type Stmt_try_catchContext struct {
	*StmtContext
}

func NewStmt_try_catchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_try_catchContext {
	var p = new(Stmt_try_catchContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_try_catchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_try_catchContext) Try_catch_stmt() ITry_catch_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITry_catch_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITry_catch_stmtContext)
}

func (s *Stmt_try_catchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_try_catch(s)
	}
}

func (s *Stmt_try_catchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_try_catch(s)
	}
}

type Stmt_yieldContext struct {
	*StmtContext
}

func NewStmt_yieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Stmt_yieldContext {
	var p = new(Stmt_yieldContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *Stmt_yieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_yieldContext) Yield_stmt() IYield_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IYield_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IYield_stmtContext)
}

func (s *Stmt_yieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt_yield(s)
	}
}

func (s *Stmt_yieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt_yield(s)
	}
}

func (p *NitroParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NitroParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStmt_assignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.Assignment_stmt()
		}

	case 2:
		localctx = NewStmt_op_assignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Assignment_op_stmt()
		}

	case 3:
		localctx = NewStmt_var_decContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(218)
			p.Var_decl_stmt()
		}

	case 4:
		localctx = NewStmt_forContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(219)
			p.For_stmt()
		}

	case 5:
		localctx = NewStmt_whileContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(220)
			p.While_stmt()
		}

	case 6:
		localctx = NewStmt_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(221)
			p.If_stmt()
		}

	case 7:
		localctx = NewStmt_funcContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(222)
			p.Func_stmt()
		}

	case 8:
		localctx = NewStmt_returnContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(223)
			p.Return_stmt()
		}

	case 9:
		localctx = NewStmt_exprContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(224)
			p.Expr()
		}

	case 10:
		localctx = NewStmt_try_catchContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(225)
			p.Try_catch_stmt()
		}

	case 11:
		localctx = NewStmt_throwContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(226)
			p.Throw_stmt()
		}

	case 12:
		localctx = NewStmt_deferContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(227)
			p.Defer_stmt()
		}

	case 13:
		localctx = NewStmt_yieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(228)
			p.Yield_stmt()
		}

	case 14:
		localctx = NewStmt_breakContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(229)
			p.Break_stmt()
		}

	case 15:
		localctx = NewStmt_continueContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(230)
			p.Continue_stmt()
		}

	case 16:
		localctx = NewStmt_inc_decContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(231)
			p.Inc_dec_stmt()
		}

	}

	return localctx
}

// IAssignment_stmtContext is an interface to support dynamic dispatch.
type IAssignment_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignment_stmtContext differentiates from other interfaces.
	IsAssignment_stmtContext()
}

type Assignment_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_stmtContext() *Assignment_stmtContext {
	var p = new(Assignment_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_assignment_stmt
	return p
}

func (*Assignment_stmtContext) IsAssignment_stmtContext() {}

func NewAssignment_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_stmtContext {
	var p = new(Assignment_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_assignment_stmt

	return p
}

func (s *Assignment_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_stmtContext) Assignment_lvalues() IAssignment_lvaluesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_lvaluesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_lvaluesContext)
}

func (s *Assignment_stmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN, 0)
}

func (s *Assignment_stmtContext) Rvalues() IRvaluesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvaluesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvaluesContext)
}

func (s *Assignment_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterAssignment_stmt(s)
	}
}

func (s *Assignment_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitAssignment_stmt(s)
	}
}

func (p *NitroParser) Assignment_stmt() (localctx IAssignment_stmtContext) {
	this := p
	_ = this

	localctx = NewAssignment_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NitroParserRULE_assignment_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Assignment_lvalues()
	}
	{
		p.SetState(235)
		p.Match(NitroParserASSIGN)
	}
	{
		p.SetState(236)
		p.Rvalues()
	}

	return localctx
}

// IAssignment_lvaluesContext is an interface to support dynamic dispatch.
type IAssignment_lvaluesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignment_lvaluesContext differentiates from other interfaces.
	IsAssignment_lvaluesContext()
}

type Assignment_lvaluesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_lvaluesContext() *Assignment_lvaluesContext {
	var p = new(Assignment_lvaluesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_assignment_lvalues
	return p
}

func (*Assignment_lvaluesContext) IsAssignment_lvaluesContext() {}

func NewAssignment_lvaluesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_lvaluesContext {
	var p = new(Assignment_lvaluesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_assignment_lvalues

	return p
}

func (s *Assignment_lvaluesContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_lvaluesContext) AllLvalue_expr() []ILvalue_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILvalue_exprContext); ok {
			len++
		}
	}

	tst := make([]ILvalue_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILvalue_exprContext); ok {
			tst[i] = t.(ILvalue_exprContext)
			i++
		}
	}

	return tst
}

func (s *Assignment_lvaluesContext) Lvalue_expr(i int) ILvalue_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalue_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalue_exprContext)
}

func (s *Assignment_lvaluesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCOMMA)
}

func (s *Assignment_lvaluesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCOMMA, i)
}

func (s *Assignment_lvaluesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_lvaluesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_lvaluesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterAssignment_lvalues(s)
	}
}

func (s *Assignment_lvaluesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitAssignment_lvalues(s)
	}
}

func (p *NitroParser) Assignment_lvalues() (localctx IAssignment_lvaluesContext) {
	this := p
	_ = this

	localctx = NewAssignment_lvaluesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NitroParserRULE_assignment_lvalues)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Lvalue_expr()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(239)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(240)
			p.Lvalue_expr()
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRvaluesContext is an interface to support dynamic dispatch.
type IRvaluesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRvaluesContext differentiates from other interfaces.
	IsRvaluesContext()
}

type RvaluesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRvaluesContext() *RvaluesContext {
	var p = new(RvaluesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_rvalues
	return p
}

func (*RvaluesContext) IsRvaluesContext() {}

func NewRvaluesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RvaluesContext {
	var p = new(RvaluesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_rvalues

	return p
}

func (s *RvaluesContext) GetParser() antlr.Parser { return s.parser }

func (s *RvaluesContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RvaluesContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RvaluesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCOMMA)
}

func (s *RvaluesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCOMMA, i)
}

func (s *RvaluesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RvaluesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RvaluesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterRvalues(s)
	}
}

func (s *RvaluesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitRvalues(s)
	}
}

func (p *NitroParser) Rvalues() (localctx IRvaluesContext) {
	this := p
	_ = this

	localctx = NewRvaluesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NitroParserRULE_rvalues)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Expr()
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(247)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(248)
			p.Expr()
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAssignment_op_stmtContext is an interface to support dynamic dispatch.
type IAssignment_op_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsAssignment_op_stmtContext differentiates from other interfaces.
	IsAssignment_op_stmtContext()
}

type Assignment_op_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAssignment_op_stmtContext() *Assignment_op_stmtContext {
	var p = new(Assignment_op_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_assignment_op_stmt
	return p
}

func (*Assignment_op_stmtContext) IsAssignment_op_stmtContext() {}

func NewAssignment_op_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_op_stmtContext {
	var p = new(Assignment_op_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_assignment_op_stmt

	return p
}

func (s *Assignment_op_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_op_stmtContext) GetOp() antlr.Token { return s.op }

func (s *Assignment_op_stmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *Assignment_op_stmtContext) Lvalue_expr() ILvalue_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalue_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalue_exprContext)
}

func (s *Assignment_op_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Assignment_op_stmtContext) ASSIGN_ADD() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN_ADD, 0)
}

func (s *Assignment_op_stmtContext) ASSIGN_SUB() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN_SUB, 0)
}

func (s *Assignment_op_stmtContext) ASSIGN_MUL() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN_MUL, 0)
}

func (s *Assignment_op_stmtContext) ASSIGN_DIV() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN_DIV, 0)
}

func (s *Assignment_op_stmtContext) ASSIGN_MOD() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN_MOD, 0)
}

func (s *Assignment_op_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_op_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_op_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterAssignment_op_stmt(s)
	}
}

func (s *Assignment_op_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitAssignment_op_stmt(s)
	}
}

func (p *NitroParser) Assignment_op_stmt() (localctx IAssignment_op_stmtContext) {
	this := p
	_ = this

	localctx = NewAssignment_op_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NitroParserRULE_assignment_op_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Lvalue_expr()
	}
	{
		p.SetState(255)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Assignment_op_stmtContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserASSIGN_ADD)|(1<<NitroParserASSIGN_SUB)|(1<<NitroParserASSIGN_MUL)|(1<<NitroParserASSIGN_DIV)|(1<<NitroParserASSIGN_MOD))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Assignment_op_stmtContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(256)
		p.Expr()
	}

	return localctx
}

// IVar_decl_stmtContext is an interface to support dynamic dispatch.
type IVar_decl_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_decl_stmtContext differentiates from other interfaces.
	IsVar_decl_stmtContext()
}

type Var_decl_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_decl_stmtContext() *Var_decl_stmtContext {
	var p = new(Var_decl_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_var_decl_stmt
	return p
}

func (*Var_decl_stmtContext) IsVar_decl_stmtContext() {}

func NewVar_decl_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_decl_stmtContext {
	var p = new(Var_decl_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_var_decl_stmt

	return p
}

func (s *Var_decl_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_decl_stmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(NitroParserVAR, 0)
}

func (s *Var_decl_stmtContext) Var_decl_vars() IVar_decl_varsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_decl_varsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_decl_varsContext)
}

func (s *Var_decl_stmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN, 0)
}

func (s *Var_decl_stmtContext) Rvalues() IRvaluesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvaluesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvaluesContext)
}

func (s *Var_decl_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_decl_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_decl_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterVar_decl_stmt(s)
	}
}

func (s *Var_decl_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitVar_decl_stmt(s)
	}
}

func (p *NitroParser) Var_decl_stmt() (localctx IVar_decl_stmtContext) {
	this := p
	_ = this

	localctx = NewVar_decl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, NitroParserRULE_var_decl_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(NitroParserVAR)
	}
	{
		p.SetState(259)
		p.Var_decl_vars()
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserASSIGN {
		{
			p.SetState(260)
			p.Match(NitroParserASSIGN)
		}
		{
			p.SetState(261)
			p.Rvalues()
		}

	}

	return localctx
}

// IVar_decl_varsContext is an interface to support dynamic dispatch.
type IVar_decl_varsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_decl_varsContext differentiates from other interfaces.
	IsVar_decl_varsContext()
}

type Var_decl_varsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_decl_varsContext() *Var_decl_varsContext {
	var p = new(Var_decl_varsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_var_decl_vars
	return p
}

func (*Var_decl_varsContext) IsVar_decl_varsContext() {}

func NewVar_decl_varsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_decl_varsContext {
	var p = new(Var_decl_varsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_var_decl_vars

	return p
}

func (s *Var_decl_varsContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_decl_varsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(NitroParserID)
}

func (s *Var_decl_varsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserID, i)
}

func (s *Var_decl_varsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCOMMA)
}

func (s *Var_decl_varsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCOMMA, i)
}

func (s *Var_decl_varsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_decl_varsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_decl_varsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterVar_decl_vars(s)
	}
}

func (s *Var_decl_varsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitVar_decl_vars(s)
	}
}

func (p *NitroParser) Var_decl_vars() (localctx IVar_decl_varsContext) {
	this := p
	_ = this

	localctx = NewVar_decl_varsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, NitroParserRULE_var_decl_vars)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(NitroParserID)
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(265)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(266)
			p.Match(NitroParserID)
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_for_stmt
	return p
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(NitroParserFOR, 0)
}

func (s *For_stmtContext) For_vars() IFor_varsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_varsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_varsContext)
}

func (s *For_stmtContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *For_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *For_stmtContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *For_stmtContext) Stmts() IStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *For_stmtContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterFor_stmt(s)
	}
}

func (s *For_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitFor_stmt(s)
	}
}

func (p *NitroParser) For_stmt() (localctx IFor_stmtContext) {
	this := p
	_ = this

	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, NitroParserRULE_for_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(273)
		p.For_vars()
	}
	{
		p.SetState(274)
		p.Match(NitroParserID)
	}
	{
		p.SetState(275)
		p.Expr()
	}
	{
		p.SetState(276)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(277)
		p.Stmts()
	}
	{
		p.SetState(278)
		p.Match(NitroParserCCURLY)
	}

	return localctx
}

// IFor_varsContext is an interface to support dynamic dispatch.
type IFor_varsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_varsContext differentiates from other interfaces.
	IsFor_varsContext()
}

type For_varsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_varsContext() *For_varsContext {
	var p = new(For_varsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_for_vars
	return p
}

func (*For_varsContext) IsFor_varsContext() {}

func NewFor_varsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_varsContext {
	var p = new(For_varsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_for_vars

	return p
}

func (s *For_varsContext) GetParser() antlr.Parser { return s.parser }

func (s *For_varsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(NitroParserID)
}

func (s *For_varsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserID, i)
}

func (s *For_varsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCOMMA)
}

func (s *For_varsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCOMMA, i)
}

func (s *For_varsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_varsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_varsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterFor_vars(s)
	}
}

func (s *For_varsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitFor_vars(s)
	}
}

func (p *NitroParser) For_vars() (localctx IFor_varsContext) {
	this := p
	_ = this

	localctx = NewFor_varsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, NitroParserRULE_for_vars)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Match(NitroParserID)
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(281)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(282)
			p.Match(NitroParserID)
		}

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_while_stmt
	return p
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(NitroParserWHILE, 0)
}

func (s *While_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *While_stmtContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *While_stmtContext) Stmts() IStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *While_stmtContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterWhile_stmt(s)
	}
}

func (s *While_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitWhile_stmt(s)
	}
}

func (p *NitroParser) While_stmt() (localctx IWhile_stmtContext) {
	this := p
	_ = this

	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, NitroParserRULE_while_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(NitroParserWHILE)
	}
	{
		p.SetState(289)
		p.Expr()
	}
	{
		p.SetState(290)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(291)
		p.Stmts()
	}
	{
		p.SetState(292)
		p.Match(NitroParserCCURLY)
	}

	return localctx
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_if_stmt
	return p
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) IF() antlr.TerminalNode {
	return s.GetToken(NitroParserIF, 0)
}

func (s *If_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_stmtContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *If_stmtContext) Stmts() IStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *If_stmtContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *If_stmtContext) AllIf_elif() []IIf_elifContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_elifContext); ok {
			len++
		}
	}

	tst := make([]IIf_elifContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_elifContext); ok {
			tst[i] = t.(IIf_elifContext)
			i++
		}
	}

	return tst
}

func (s *If_stmtContext) If_elif(i int) IIf_elifContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_elifContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_elifContext)
}

func (s *If_stmtContext) If_else() IIf_elseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_elseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_elseContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterIf_stmt(s)
	}
}

func (s *If_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitIf_stmt(s)
	}
}

func (p *NitroParser) If_stmt() (localctx IIf_stmtContext) {
	this := p
	_ = this

	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, NitroParserRULE_if_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(295)
		p.Expr()
	}
	{
		p.SetState(296)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(297)
		p.Stmts()
	}
	{
		p.SetState(298)
		p.Match(NitroParserCCURLY)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(299)
				p.If_elif()
			}

		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(305)
			p.If_else()
		}

	}

	return localctx
}

// IIf_elifContext is an interface to support dynamic dispatch.
type IIf_elifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_elifContext differentiates from other interfaces.
	IsIf_elifContext()
}

type If_elifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_elifContext() *If_elifContext {
	var p = new(If_elifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_if_elif
	return p
}

func (*If_elifContext) IsIf_elifContext() {}

func NewIf_elifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_elifContext {
	var p = new(If_elifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_if_elif

	return p
}

func (s *If_elifContext) GetParser() antlr.Parser { return s.parser }

func (s *If_elifContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NitroParserELSE, 0)
}

func (s *If_elifContext) IF() antlr.TerminalNode {
	return s.GetToken(NitroParserIF, 0)
}

func (s *If_elifContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_elifContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *If_elifContext) Stmts() IStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *If_elifContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *If_elifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_elifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_elifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterIf_elif(s)
	}
}

func (s *If_elifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitIf_elif(s)
	}
}

func (p *NitroParser) If_elif() (localctx IIf_elifContext) {
	this := p
	_ = this

	localctx = NewIf_elifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, NitroParserRULE_if_elif)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(309)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(310)
		p.Expr()
	}
	{
		p.SetState(311)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(312)
		p.Stmts()
	}
	{
		p.SetState(313)
		p.Match(NitroParserCCURLY)
	}

	return localctx
}

// IIf_elseContext is an interface to support dynamic dispatch.
type IIf_elseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_elseContext differentiates from other interfaces.
	IsIf_elseContext()
}

type If_elseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_elseContext() *If_elseContext {
	var p = new(If_elseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_if_else
	return p
}

func (*If_elseContext) IsIf_elseContext() {}

func NewIf_elseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_elseContext {
	var p = new(If_elseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_if_else

	return p
}

func (s *If_elseContext) GetParser() antlr.Parser { return s.parser }

func (s *If_elseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NitroParserELSE, 0)
}

func (s *If_elseContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *If_elseContext) Stmts() IStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *If_elseContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *If_elseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_elseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_elseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterIf_else(s)
	}
}

func (s *If_elseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitIf_else(s)
	}
}

func (p *NitroParser) If_else() (localctx IIf_elseContext) {
	this := p
	_ = this

	localctx = NewIf_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, NitroParserRULE_if_else)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(316)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(317)
		p.Stmts()
	}
	{
		p.SetState(318)
		p.Match(NitroParserCCURLY)
	}

	return localctx
}

// IFunc_stmtContext is an interface to support dynamic dispatch.
type IFunc_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_stmtContext differentiates from other interfaces.
	IsFunc_stmtContext()
}

type Func_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_stmtContext() *Func_stmtContext {
	var p = new(Func_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_func_stmt
	return p
}

func (*Func_stmtContext) IsFunc_stmtContext() {}

func NewFunc_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_stmtContext {
	var p = new(Func_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_func_stmt

	return p
}

func (s *Func_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_stmtContext) FUNC() antlr.TerminalNode {
	return s.GetToken(NitroParserFUNC, 0)
}

func (s *Func_stmtContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Func_stmtContext) OPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserOPAREN, 0)
}

func (s *Func_stmtContext) CPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserCPAREN, 0)
}

func (s *Func_stmtContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Func_stmtContext) Stmts() IStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *Func_stmtContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *Func_stmtContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *Func_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterFunc_stmt(s)
	}
}

func (s *Func_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitFunc_stmt(s)
	}
}

func (p *NitroParser) Func_stmt() (localctx IFunc_stmtContext) {
	this := p
	_ = this

	localctx = NewFunc_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, NitroParserRULE_func_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(NitroParserFUNC)
	}
	{
		p.SetState(321)
		p.Match(NitroParserID)
	}
	{
		p.SetState(322)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(323)
			p.Param_list()
		}

	}
	{
		p.SetState(326)
		p.Match(NitroParserCPAREN)
	}
	{
		p.SetState(327)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(328)
		p.Stmts()
	}
	{
		p.SetState(329)
		p.Match(NitroParserCCURLY)
	}

	return localctx
}

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_listContext() *Param_listContext {
	var p = new(Param_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_param_list
	return p
}

func (*Param_listContext) IsParam_listContext() {}

func NewParam_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_listContext {
	var p = new(Param_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_param_list

	return p
}

func (s *Param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_listContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(NitroParserID)
}

func (s *Param_listContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserID, i)
}

func (s *Param_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCOMMA)
}

func (s *Param_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCOMMA, i)
}

func (s *Param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterParam_list(s)
	}
}

func (s *Param_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitParam_list(s)
	}
}

func (p *NitroParser) Param_list() (localctx IParam_listContext) {
	this := p
	_ = this

	localctx = NewParam_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, NitroParserRULE_param_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(331)
		p.Match(NitroParserID)
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(332)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(333)
			p.Match(NitroParserID)
		}

		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IReturn_stmtContext is an interface to support dynamic dispatch.
type IReturn_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturn_stmtContext differentiates from other interfaces.
	IsReturn_stmtContext()
}

type Return_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_stmtContext() *Return_stmtContext {
	var p = new(Return_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_return_stmt
	return p
}

func (*Return_stmtContext) IsReturn_stmtContext() {}

func NewReturn_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_stmtContext {
	var p = new(Return_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_return_stmt

	return p
}

func (s *Return_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_stmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(NitroParserRETURN, 0)
}

func (s *Return_stmtContext) Rvalues() IRvaluesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvaluesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvaluesContext)
}

func (s *Return_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterReturn_stmt(s)
	}
}

func (s *Return_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitReturn_stmt(s)
	}
}

func (p *NitroParser) Return_stmt() (localctx IReturn_stmtContext) {
	this := p
	_ = this

	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, NitroParserRULE_return_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(NitroParserRETURN)
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37))|(1<<(NitroParserEXEC_PREFIX-37)))) != 0) {
		{
			p.SetState(340)
			p.Rvalues()
		}

	}

	return localctx
}

// ITry_catch_stmtContext is an interface to support dynamic dispatch.
type ITry_catch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTry_catch_stmtContext differentiates from other interfaces.
	IsTry_catch_stmtContext()
}

type Try_catch_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTry_catch_stmtContext() *Try_catch_stmtContext {
	var p = new(Try_catch_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_try_catch_stmt
	return p
}

func (*Try_catch_stmtContext) IsTry_catch_stmtContext() {}

func NewTry_catch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Try_catch_stmtContext {
	var p = new(Try_catch_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_try_catch_stmt

	return p
}

func (s *Try_catch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Try_catch_stmtContext) TRY() antlr.TerminalNode {
	return s.GetToken(NitroParserTRY, 0)
}

func (s *Try_catch_stmtContext) AllOCURLY() []antlr.TerminalNode {
	return s.GetTokens(NitroParserOCURLY)
}

func (s *Try_catch_stmtContext) OCURLY(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, i)
}

func (s *Try_catch_stmtContext) AllStmts() []IStmtsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtsContext); ok {
			len++
		}
	}

	tst := make([]IStmtsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtsContext); ok {
			tst[i] = t.(IStmtsContext)
			i++
		}
	}

	return tst
}

func (s *Try_catch_stmtContext) Stmts(i int) IStmtsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *Try_catch_stmtContext) AllCCURLY() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCCURLY)
}

func (s *Try_catch_stmtContext) CCURLY(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, i)
}

func (s *Try_catch_stmtContext) CATCH() antlr.TerminalNode {
	return s.GetToken(NitroParserCATCH, 0)
}

func (s *Try_catch_stmtContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Try_catch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Try_catch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Try_catch_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterTry_catch_stmt(s)
	}
}

func (s *Try_catch_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitTry_catch_stmt(s)
	}
}

func (p *NitroParser) Try_catch_stmt() (localctx ITry_catch_stmtContext) {
	this := p
	_ = this

	localctx = NewTry_catch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, NitroParserRULE_try_catch_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.Match(NitroParserTRY)
	}
	{
		p.SetState(344)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(345)
		p.Stmts()
	}
	{
		p.SetState(346)
		p.Match(NitroParserCCURLY)
	}
	{
		p.SetState(347)
		p.Match(NitroParserCATCH)
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(348)
			p.Match(NitroParserID)
		}

	}
	{
		p.SetState(351)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(352)
		p.Stmts()
	}
	{
		p.SetState(353)
		p.Match(NitroParserCCURLY)
	}

	return localctx
}

// IThrow_stmtContext is an interface to support dynamic dispatch.
type IThrow_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThrow_stmtContext differentiates from other interfaces.
	IsThrow_stmtContext()
}

type Throw_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThrow_stmtContext() *Throw_stmtContext {
	var p = new(Throw_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_throw_stmt
	return p
}

func (*Throw_stmtContext) IsThrow_stmtContext() {}

func NewThrow_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Throw_stmtContext {
	var p = new(Throw_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_throw_stmt

	return p
}

func (s *Throw_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Throw_stmtContext) THROW() antlr.TerminalNode {
	return s.GetToken(NitroParserTHROW, 0)
}

func (s *Throw_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Throw_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Throw_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Throw_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterThrow_stmt(s)
	}
}

func (s *Throw_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitThrow_stmt(s)
	}
}

func (p *NitroParser) Throw_stmt() (localctx IThrow_stmtContext) {
	this := p
	_ = this

	localctx = NewThrow_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, NitroParserRULE_throw_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(NitroParserTHROW)
	}
	{
		p.SetState(356)
		p.Expr()
	}

	return localctx
}

// IDefer_stmtContext is an interface to support dynamic dispatch.
type IDefer_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefer_stmtContext differentiates from other interfaces.
	IsDefer_stmtContext()
}

type Defer_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefer_stmtContext() *Defer_stmtContext {
	var p = new(Defer_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_defer_stmt
	return p
}

func (*Defer_stmtContext) IsDefer_stmtContext() {}

func NewDefer_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Defer_stmtContext {
	var p = new(Defer_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_defer_stmt

	return p
}

func (s *Defer_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Defer_stmtContext) DEFER() antlr.TerminalNode {
	return s.GetToken(NitroParserDEFER, 0)
}

func (s *Defer_stmtContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Defer_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Defer_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Defer_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterDefer_stmt(s)
	}
}

func (s *Defer_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitDefer_stmt(s)
	}
}

func (p *NitroParser) Defer_stmt() (localctx IDefer_stmtContext) {
	this := p
	_ = this

	localctx = NewDefer_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, NitroParserRULE_defer_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(NitroParserDEFER)
	}
	{
		p.SetState(359)
		p.primary_expr(0)
	}

	return localctx
}

// IYield_stmtContext is an interface to support dynamic dispatch.
type IYield_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYield_stmtContext differentiates from other interfaces.
	IsYield_stmtContext()
}

type Yield_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYield_stmtContext() *Yield_stmtContext {
	var p = new(Yield_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_yield_stmt
	return p
}

func (*Yield_stmtContext) IsYield_stmtContext() {}

func NewYield_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Yield_stmtContext {
	var p = new(Yield_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_yield_stmt

	return p
}

func (s *Yield_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Yield_stmtContext) YIELD() antlr.TerminalNode {
	return s.GetToken(NitroParserYIELD, 0)
}

func (s *Yield_stmtContext) Rvalues() IRvaluesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvaluesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvaluesContext)
}

func (s *Yield_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Yield_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Yield_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterYield_stmt(s)
	}
}

func (s *Yield_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitYield_stmt(s)
	}
}

func (p *NitroParser) Yield_stmt() (localctx IYield_stmtContext) {
	this := p
	_ = this

	localctx = NewYield_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, NitroParserRULE_yield_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Match(NitroParserYIELD)
	}
	{
		p.SetState(362)
		p.Rvalues()
	}

	return localctx
}

// IBreak_stmtContext is an interface to support dynamic dispatch.
type IBreak_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreak_stmtContext differentiates from other interfaces.
	IsBreak_stmtContext()
}

type Break_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_stmtContext() *Break_stmtContext {
	var p = new(Break_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_break_stmt
	return p
}

func (*Break_stmtContext) IsBreak_stmtContext() {}

func NewBreak_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_stmtContext {
	var p = new(Break_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_break_stmt

	return p
}

func (s *Break_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Break_stmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(NitroParserBREAK, 0)
}

func (s *Break_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterBreak_stmt(s)
	}
}

func (s *Break_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitBreak_stmt(s)
	}
}

func (p *NitroParser) Break_stmt() (localctx IBreak_stmtContext) {
	this := p
	_ = this

	localctx = NewBreak_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, NitroParserRULE_break_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(NitroParserBREAK)
	}

	return localctx
}

// IContinue_stmtContext is an interface to support dynamic dispatch.
type IContinue_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinue_stmtContext differentiates from other interfaces.
	IsContinue_stmtContext()
}

type Continue_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_stmtContext() *Continue_stmtContext {
	var p = new(Continue_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_continue_stmt
	return p
}

func (*Continue_stmtContext) IsContinue_stmtContext() {}

func NewContinue_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_stmtContext {
	var p = new(Continue_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_continue_stmt

	return p
}

func (s *Continue_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Continue_stmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(NitroParserCONTINUE, 0)
}

func (s *Continue_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterContinue_stmt(s)
	}
}

func (s *Continue_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitContinue_stmt(s)
	}
}

func (p *NitroParser) Continue_stmt() (localctx IContinue_stmtContext) {
	this := p
	_ = this

	localctx = NewContinue_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, NitroParserRULE_continue_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(NitroParserCONTINUE)
	}

	return localctx
}

// IInc_dec_stmtContext is an interface to support dynamic dispatch.
type IInc_dec_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsInc_dec_stmtContext differentiates from other interfaces.
	IsInc_dec_stmtContext()
}

type Inc_dec_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyInc_dec_stmtContext() *Inc_dec_stmtContext {
	var p = new(Inc_dec_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_inc_dec_stmt
	return p
}

func (*Inc_dec_stmtContext) IsInc_dec_stmtContext() {}

func NewInc_dec_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inc_dec_stmtContext {
	var p = new(Inc_dec_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_inc_dec_stmt

	return p
}

func (s *Inc_dec_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Inc_dec_stmtContext) GetOp() antlr.Token { return s.op }

func (s *Inc_dec_stmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *Inc_dec_stmtContext) Lvalue_expr() ILvalue_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalue_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalue_exprContext)
}

func (s *Inc_dec_stmtContext) INC() antlr.TerminalNode {
	return s.GetToken(NitroParserINC, 0)
}

func (s *Inc_dec_stmtContext) DEC() antlr.TerminalNode {
	return s.GetToken(NitroParserDEC, 0)
}

func (s *Inc_dec_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inc_dec_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inc_dec_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterInc_dec_stmt(s)
	}
}

func (s *Inc_dec_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitInc_dec_stmt(s)
	}
}

func (p *NitroParser) Inc_dec_stmt() (localctx IInc_dec_stmtContext) {
	this := p
	_ = this

	localctx = NewInc_dec_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, NitroParserRULE_inc_dec_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Lvalue_expr()
	}
	{
		p.SetState(369)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Inc_dec_stmtContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == NitroParserINC || _la == NitroParserDEC) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Inc_dec_stmtContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Short_lambda_expr() IShort_lambda_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShort_lambda_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShort_lambda_exprContext)
}

func (s *ExprContext) Expr3() IExpr3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr3Context)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *NitroParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, NitroParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(373)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserLAMBDA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.Short_lambda_expr()
		}

	case NitroParserFALSE, NitroParserFUNC, NitroParserNIL, NitroParserNOT, NitroParserTRUE, NitroParserADD, NitroParserSUB, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserNUMBER, NitroParserID, NitroParserREGEX, NitroParserCHAR, NitroParserSTRING, NitroParserEXEC_PREFIX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.expr3(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpr3Context is an interface to support dynamic dispatch.
type IExpr3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr3Context differentiates from other interfaces.
	IsExpr3Context()
}

type Expr3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr3Context() *Expr3Context {
	var p = new(Expr3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_expr3
	return p
}

func (*Expr3Context) IsExpr3Context() {}

func NewExpr3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr3Context {
	var p = new(Expr3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_expr3

	return p
}

func (s *Expr3Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr3Context) Binary_expr() IBinary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinary_exprContext)
}

func (s *Expr3Context) AllExpr3() []IExpr3Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpr3Context); ok {
			len++
		}
	}

	tst := make([]IExpr3Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpr3Context); ok {
			tst[i] = t.(IExpr3Context)
			i++
		}
	}

	return tst
}

func (s *Expr3Context) Expr3(i int) IExpr3Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr3Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr3Context)
}

func (s *Expr3Context) QUESTION_MARK() antlr.TerminalNode {
	return s.GetToken(NitroParserQUESTION_MARK, 0)
}

func (s *Expr3Context) COLON() antlr.TerminalNode {
	return s.GetToken(NitroParserCOLON, 0)
}

func (s *Expr3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterExpr3(s)
	}
}

func (s *Expr3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitExpr3(s)
	}
}

func (p *NitroParser) Expr3() (localctx IExpr3Context) {
	return p.expr3(0)
}

func (p *NitroParser) expr3(_p int) (localctx IExpr3Context) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr3Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr3Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 74
	p.EnterRecursionRule(localctx, 74, NitroParserRULE_expr3, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.binary_expr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr3Context(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_expr3)
			p.SetState(378)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(379)
				p.Match(NitroParserQUESTION_MARK)
			}
			{
				p.SetState(380)
				p.expr3(0)
			}
			{
				p.SetState(381)
				p.Match(NitroParserCOLON)
			}
			{
				p.SetState(382)
				p.expr3(2)
			}

		}
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// IBinary_exprContext is an interface to support dynamic dispatch.
type IBinary_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsBinary_exprContext differentiates from other interfaces.
	IsBinary_exprContext()
}

type Binary_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyBinary_exprContext() *Binary_exprContext {
	var p = new(Binary_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_binary_expr
	return p
}

func (*Binary_exprContext) IsBinary_exprContext() {}

func NewBinary_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exprContext {
	var p = new(Binary_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_binary_expr

	return p
}

func (s *Binary_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exprContext) GetOp() antlr.Token { return s.op }

func (s *Binary_exprContext) SetOp(v antlr.Token) { s.op = v }

func (s *Binary_exprContext) Unary_expr() IUnary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_exprContext)
}

func (s *Binary_exprContext) AllBinary_expr() []IBinary_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBinary_exprContext); ok {
			len++
		}
	}

	tst := make([]IBinary_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBinary_exprContext); ok {
			tst[i] = t.(IBinary_exprContext)
			i++
		}
	}

	return tst
}

func (s *Binary_exprContext) Binary_expr(i int) IBinary_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinary_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinary_exprContext)
}

func (s *Binary_exprContext) MUL() antlr.TerminalNode {
	return s.GetToken(NitroParserMUL, 0)
}

func (s *Binary_exprContext) DIV() antlr.TerminalNode {
	return s.GetToken(NitroParserDIV, 0)
}

func (s *Binary_exprContext) MOD() antlr.TerminalNode {
	return s.GetToken(NitroParserMOD, 0)
}

func (s *Binary_exprContext) ADD() antlr.TerminalNode {
	return s.GetToken(NitroParserADD, 0)
}

func (s *Binary_exprContext) SUB() antlr.TerminalNode {
	return s.GetToken(NitroParserSUB, 0)
}

func (s *Binary_exprContext) PIPE() antlr.TerminalNode {
	return s.GetToken(NitroParserPIPE, 0)
}

func (s *Binary_exprContext) LT() antlr.TerminalNode {
	return s.GetToken(NitroParserLT, 0)
}

func (s *Binary_exprContext) LE() antlr.TerminalNode {
	return s.GetToken(NitroParserLE, 0)
}

func (s *Binary_exprContext) GT() antlr.TerminalNode {
	return s.GetToken(NitroParserGT, 0)
}

func (s *Binary_exprContext) GE() antlr.TerminalNode {
	return s.GetToken(NitroParserGE, 0)
}

func (s *Binary_exprContext) EQ() antlr.TerminalNode {
	return s.GetToken(NitroParserEQ, 0)
}

func (s *Binary_exprContext) NE() antlr.TerminalNode {
	return s.GetToken(NitroParserNE, 0)
}

func (s *Binary_exprContext) AND() antlr.TerminalNode {
	return s.GetToken(NitroParserAND, 0)
}

func (s *Binary_exprContext) OR() antlr.TerminalNode {
	return s.GetToken(NitroParserOR, 0)
}

func (s *Binary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Binary_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterBinary_expr(s)
	}
}

func (s *Binary_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitBinary_expr(s)
	}
}

func (p *NitroParser) Binary_expr() (localctx IBinary_exprContext) {
	return p.binary_expr(0)
}

func (p *NitroParser) binary_expr(_p int) (localctx IBinary_exprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBinary_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBinary_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 76
	p.EnterRecursionRule(localctx, 76, NitroParserRULE_binary_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Unary_expr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(410)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(392)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(393)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Binary_exprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(NitroParserMUL-39))|(1<<(NitroParserDIV-39))|(1<<(NitroParserMOD-39)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Binary_exprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(394)
					p.binary_expr(7)
				}

			case 2:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(395)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(396)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Binary_exprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == NitroParserADD || _la == NitroParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Binary_exprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(397)
					p.binary_expr(6)
				}

			case 3:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(398)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(399)

					var _m = p.Match(NitroParserPIPE)

					localctx.(*Binary_exprContext).op = _m
				}
				{
					p.SetState(400)
					p.binary_expr(5)
				}

			case 4:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(401)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(402)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Binary_exprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(NitroParserEQ-31))|(1<<(NitroParserNE-31))|(1<<(NitroParserLT-31))|(1<<(NitroParserLE-31))|(1<<(NitroParserGT-31))|(1<<(NitroParserGE-31)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Binary_exprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(403)
					p.binary_expr(4)
				}

			case 5:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(404)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(405)

					var _m = p.Match(NitroParserAND)

					localctx.(*Binary_exprContext).op = _m
				}
				{
					p.SetState(406)
					p.binary_expr(3)
				}

			case 6:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(407)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(408)

					var _m = p.Match(NitroParserOR)

					localctx.(*Binary_exprContext).op = _m
				}
				{
					p.SetState(409)
					p.binary_expr(2)
				}

			}

		}
		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}

	return localctx
}

// IUnary_exprContext is an interface to support dynamic dispatch.
type IUnary_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsUnary_exprContext differentiates from other interfaces.
	IsUnary_exprContext()
}

type Unary_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyUnary_exprContext() *Unary_exprContext {
	var p = new(Unary_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_unary_expr
	return p
}

func (*Unary_exprContext) IsUnary_exprContext() {}

func NewUnary_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_exprContext {
	var p = new(Unary_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_unary_expr

	return p
}

func (s *Unary_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_exprContext) GetOp() antlr.Token { return s.op }

func (s *Unary_exprContext) SetOp(v antlr.Token) { s.op = v }

func (s *Unary_exprContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Unary_exprContext) NOT() antlr.TerminalNode {
	return s.GetToken(NitroParserNOT, 0)
}

func (s *Unary_exprContext) ADD() antlr.TerminalNode {
	return s.GetToken(NitroParserADD, 0)
}

func (s *Unary_exprContext) SUB() antlr.TerminalNode {
	return s.GetToken(NitroParserSUB, 0)
}

func (s *Unary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterUnary_expr(s)
	}
}

func (s *Unary_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitUnary_expr(s)
	}
}

func (p *NitroParser) Unary_expr() (localctx IUnary_exprContext) {
	this := p
	_ = this

	localctx = NewUnary_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, NitroParserRULE_unary_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(422)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(415)

			var _m = p.Match(NitroParserNOT)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(416)
			p.primary_expr(0)
		}

	case NitroParserADD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(417)

			var _m = p.Match(NitroParserADD)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(418)
			p.primary_expr(0)
		}

	case NitroParserSUB:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(419)

			var _m = p.Match(NitroParserSUB)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(420)
			p.primary_expr(0)
		}

	case NitroParserFALSE, NitroParserFUNC, NitroParserNIL, NitroParserTRUE, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserNUMBER, NitroParserID, NitroParserREGEX, NitroParserCHAR, NitroParserSTRING, NitroParserEXEC_PREFIX:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(421)
			p.primary_expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimary_exprContext is an interface to support dynamic dispatch.
type IPrimary_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimary_exprContext differentiates from other interfaces.
	IsPrimary_exprContext()
}

type Primary_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_exprContext() *Primary_exprContext {
	var p = new(Primary_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_primary_expr
	return p
}

func (*Primary_exprContext) IsPrimary_exprContext() {}

func NewPrimary_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_exprContext {
	var p = new(Primary_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_primary_expr

	return p
}

func (s *Primary_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_exprContext) CopyFrom(ctx *Primary_exprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Primary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Primary_expr_regexContext struct {
	*Primary_exprContext
}

func NewPrimary_expr_regexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_regexContext {
	var p = new(Primary_expr_regexContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_regexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_regexContext) REGEX() antlr.TerminalNode {
	return s.GetToken(NitroParserREGEX, 0)
}

func (s *Primary_expr_regexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_regex(s)
	}
}

func (s *Primary_expr_regexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_regex(s)
	}
}

type Primary_expr_objectContext struct {
	*Primary_exprContext
}

func NewPrimary_expr_objectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_objectContext {
	var p = new(Primary_expr_objectContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_objectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_objectContext) Object_literal() IObject_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_literalContext)
}

func (s *Primary_expr_objectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_object(s)
	}
}

func (s *Primary_expr_objectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_object(s)
	}
}

type Primary_expr_parenthesisContext struct {
	*Primary_exprContext
}

func NewPrimary_expr_parenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_parenthesisContext {
	var p = new(Primary_expr_parenthesisContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_parenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_parenthesisContext) OPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserOPAREN, 0)
}

func (s *Primary_expr_parenthesisContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Primary_expr_parenthesisContext) CPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserCPAREN, 0)
}

func (s *Primary_expr_parenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_parenthesis(s)
	}
}

func (s *Primary_expr_parenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_parenthesis(s)
	}
}

type Primary_expr_simple_refContext struct {
	*Primary_exprContext
}

func NewPrimary_expr_simple_refContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_simple_refContext {
	var p = new(Primary_expr_simple_refContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_simple_refContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_simple_refContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Primary_expr_simple_refContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_simple_ref(s)
	}
}

func (s *Primary_expr_simple_refContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_simple_ref(s)
	}
}

type Primary_expr_lambdaContext struct {
	*Primary_exprContext
}

func NewPrimary_expr_lambdaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_lambdaContext {
	var p = new(Primary_expr_lambdaContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_lambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_lambdaContext) Lambda_expr() ILambda_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambda_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambda_exprContext)
}

func (s *Primary_expr_lambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_lambda(s)
	}
}

func (s *Primary_expr_lambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_lambda(s)
	}
}

type Primary_expr_indexContext struct {
	*Primary_exprContext
}

func NewPrimary_expr_indexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_indexContext {
	var p = new(Primary_expr_indexContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_indexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_indexContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Primary_expr_indexContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserOBRACKET, 0)
}

func (s *Primary_expr_indexContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Primary_expr_indexContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserCBRACKET, 0)
}

func (s *Primary_expr_indexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_index(s)
	}
}

func (s *Primary_expr_indexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_index(s)
	}
}

type Primary_expr_literalContext struct {
	*Primary_exprContext
}

func NewPrimary_expr_literalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_literalContext {
	var p = new(Primary_expr_literalContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_literalContext) Simple_literal() ISimple_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_literalContext)
}

func (s *Primary_expr_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_literal(s)
	}
}

func (s *Primary_expr_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_literal(s)
	}
}

type Primary_expr_member_accessContext struct {
	*Primary_exprContext
}

func NewPrimary_expr_member_accessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_member_accessContext {
	var p = new(Primary_expr_member_accessContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_member_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_member_accessContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Primary_expr_member_accessContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(NitroParserPERIOD, 0)
}

func (s *Primary_expr_member_accessContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Primary_expr_member_accessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_member_access(s)
	}
}

func (s *Primary_expr_member_accessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_member_access(s)
	}
}

type Primary_exec_exprContext struct {
	*Primary_exprContext
}

func NewPrimary_exec_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_exec_exprContext {
	var p = new(Primary_exec_exprContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_exec_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_exec_exprContext) Exec_expr() IExec_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExec_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExec_exprContext)
}

func (s *Primary_exec_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_exec_expr(s)
	}
}

func (s *Primary_exec_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_exec_expr(s)
	}
}

type Primary_expr_arrayContext struct {
	*Primary_exprContext
}

func NewPrimary_expr_arrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_arrayContext {
	var p = new(Primary_expr_arrayContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_arrayContext) Array_literal() IArray_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_literalContext)
}

func (s *Primary_expr_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_array(s)
	}
}

func (s *Primary_expr_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_array(s)
	}
}

type Primary_expr_callContext struct {
	*Primary_exprContext
}

func NewPrimary_expr_callContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_callContext {
	var p = new(Primary_expr_callContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_callContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Primary_expr_callContext) OPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserOPAREN, 0)
}

func (s *Primary_expr_callContext) CPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserCPAREN, 0)
}

func (s *Primary_expr_callContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *Primary_expr_callContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(NitroParserEXPAND, 0)
}

func (s *Primary_expr_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_call(s)
	}
}

func (s *Primary_expr_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_call(s)
	}
}

type Primary_expr_sliceContext struct {
	*Primary_exprContext
	b IExprContext
	e IExprContext
}

func NewPrimary_expr_sliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primary_expr_sliceContext {
	var p = new(Primary_expr_sliceContext)

	p.Primary_exprContext = NewEmptyPrimary_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Primary_exprContext))

	return p
}

func (s *Primary_expr_sliceContext) GetB() IExprContext { return s.b }

func (s *Primary_expr_sliceContext) GetE() IExprContext { return s.e }

func (s *Primary_expr_sliceContext) SetB(v IExprContext) { s.b = v }

func (s *Primary_expr_sliceContext) SetE(v IExprContext) { s.e = v }

func (s *Primary_expr_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expr_sliceContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Primary_expr_sliceContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserOBRACKET, 0)
}

func (s *Primary_expr_sliceContext) COLON() antlr.TerminalNode {
	return s.GetToken(NitroParserCOLON, 0)
}

func (s *Primary_expr_sliceContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserCBRACKET, 0)
}

func (s *Primary_expr_sliceContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Primary_expr_sliceContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Primary_expr_sliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPrimary_expr_slice(s)
	}
}

func (s *Primary_expr_sliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPrimary_expr_slice(s)
	}
}

func (p *NitroParser) Primary_expr() (localctx IPrimary_exprContext) {
	return p.primary_expr(0)
}

func (p *NitroParser) primary_expr(_p int) (localctx IPrimary_exprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimary_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimary_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 80
	p.EnterRecursionRule(localctx, 80, NitroParserRULE_primary_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(436)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserID:
		localctx = NewPrimary_expr_simple_refContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(425)
			p.Match(NitroParserID)
		}

	case NitroParserFUNC:
		localctx = NewPrimary_expr_lambdaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(426)
			p.Lambda_expr()
		}

	case NitroParserEXEC_PREFIX:
		localctx = NewPrimary_exec_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(427)
			p.Exec_expr()
		}

	case NitroParserOCURLY:
		localctx = NewPrimary_expr_objectContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(428)
			p.Object_literal()
		}

	case NitroParserOBRACKET:
		localctx = NewPrimary_expr_arrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(429)
			p.Array_literal()
		}

	case NitroParserFALSE, NitroParserNIL, NitroParserTRUE, NitroParserNUMBER, NitroParserCHAR, NitroParserSTRING:
		localctx = NewPrimary_expr_literalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(430)
			p.Simple_literal()
		}

	case NitroParserREGEX:
		localctx = NewPrimary_expr_regexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(431)
			p.Match(NitroParserREGEX)
		}

	case NitroParserOPAREN:
		localctx = NewPrimary_expr_parenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(432)
			p.Match(NitroParserOPAREN)
		}
		{
			p.SetState(433)
			p.Expr()
		}
		{
			p.SetState(434)
			p.Match(NitroParserCPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(467)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(465)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimary_expr_member_accessContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(438)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(439)
					p.Match(NitroParserPERIOD)
				}
				{
					p.SetState(440)
					p.Match(NitroParserID)
				}

			case 2:
				localctx = NewPrimary_expr_indexContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(441)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(442)
					p.Match(NitroParserOBRACKET)
				}
				{
					p.SetState(443)
					p.Expr()
				}
				{
					p.SetState(444)
					p.Match(NitroParserCBRACKET)
				}

			case 3:
				localctx = NewPrimary_expr_sliceContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(446)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(447)
					p.Match(NitroParserOBRACKET)
				}
				p.SetState(449)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37))|(1<<(NitroParserEXEC_PREFIX-37)))) != 0) {
					{
						p.SetState(448)

						var _x = p.Expr()

						localctx.(*Primary_expr_sliceContext).b = _x
					}

				}
				{
					p.SetState(451)
					p.Match(NitroParserCOLON)
				}
				p.SetState(453)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37))|(1<<(NitroParserEXEC_PREFIX-37)))) != 0) {
					{
						p.SetState(452)

						var _x = p.Expr()

						localctx.(*Primary_expr_sliceContext).e = _x
					}

				}
				{
					p.SetState(455)
					p.Match(NitroParserCBRACKET)
				}

			case 4:
				localctx = NewPrimary_expr_callContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(456)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(457)
					p.Match(NitroParserOPAREN)
				}
				p.SetState(462)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37))|(1<<(NitroParserEXEC_PREFIX-37)))) != 0) {
					{
						p.SetState(458)
						p.Arg_list()
					}
					p.SetState(460)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if _la == NitroParserEXPAND {
						{
							p.SetState(459)
							p.Match(NitroParserEXPAND)
						}

					}

				}
				{
					p.SetState(464)
					p.Match(NitroParserCPAREN)
				}

			}

		}
		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

// ISimple_literalContext is an interface to support dynamic dispatch.
type ISimple_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVal returns the val token.
	GetVal() antlr.Token

	// SetVal sets the val token.
	SetVal(antlr.Token)

	// IsSimple_literalContext differentiates from other interfaces.
	IsSimple_literalContext()
}

type Simple_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	val    antlr.Token
}

func NewEmptySimple_literalContext() *Simple_literalContext {
	var p = new(Simple_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_simple_literal
	return p
}

func (*Simple_literalContext) IsSimple_literalContext() {}

func NewSimple_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_literalContext {
	var p = new(Simple_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_simple_literal

	return p
}

func (s *Simple_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_literalContext) GetVal() antlr.Token { return s.val }

func (s *Simple_literalContext) SetVal(v antlr.Token) { s.val = v }

func (s *Simple_literalContext) STRING() antlr.TerminalNode {
	return s.GetToken(NitroParserSTRING, 0)
}

func (s *Simple_literalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NitroParserNUMBER, 0)
}

func (s *Simple_literalContext) CHAR() antlr.TerminalNode {
	return s.GetToken(NitroParserCHAR, 0)
}

func (s *Simple_literalContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NitroParserTRUE, 0)
}

func (s *Simple_literalContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NitroParserFALSE, 0)
}

func (s *Simple_literalContext) NIL() antlr.TerminalNode {
	return s.GetToken(NitroParserNIL, 0)
}

func (s *Simple_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterSimple_literal(s)
	}
}

func (s *Simple_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitSimple_literal(s)
	}
}

func (p *NitroParser) Simple_literal() (localctx ISimple_literalContext) {
	this := p
	_ = this

	localctx = NewSimple_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, NitroParserRULE_simple_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(470)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Simple_literalContext).val = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserNIL)|(1<<NitroParserTRUE))) != 0) || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(NitroParserNUMBER-59))|(1<<(NitroParserCHAR-59))|(1<<(NitroParserSTRING-59)))) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Simple_literalContext).val = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_arg_list
	return p
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Arg_listContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Arg_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCOMMA)
}

func (s *Arg_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCOMMA, i)
}

func (s *Arg_listContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arg_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterArg_list(s)
	}
}

func (s *Arg_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitArg_list(s)
	}
}

func (p *NitroParser) Arg_list() (localctx IArg_listContext) {
	this := p
	_ = this

	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, NitroParserRULE_arg_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		p.Expr()
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(473)
				p.Match(NitroParserCOMMA)
			}
			{
				p.SetState(474)
				p.Expr()
			}

		}
		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
		{
			p.SetState(480)
			_la = p.GetTokenStream().LA(1)

			if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// ILvalue_exprContext is an interface to support dynamic dispatch.
type ILvalue_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLvalue_exprContext differentiates from other interfaces.
	IsLvalue_exprContext()
}

type Lvalue_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalue_exprContext() *Lvalue_exprContext {
	var p = new(Lvalue_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_lvalue_expr
	return p
}

func (*Lvalue_exprContext) IsLvalue_exprContext() {}

func NewLvalue_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lvalue_exprContext {
	var p = new(Lvalue_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_lvalue_expr

	return p
}

func (s *Lvalue_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Lvalue_exprContext) CopyFrom(ctx *Lvalue_exprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Lvalue_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lvalue_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Lvalue_expr_simple_refContext struct {
	*Lvalue_exprContext
}

func NewLvalue_expr_simple_refContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Lvalue_expr_simple_refContext {
	var p = new(Lvalue_expr_simple_refContext)

	p.Lvalue_exprContext = NewEmptyLvalue_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Lvalue_exprContext))

	return p
}

func (s *Lvalue_expr_simple_refContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lvalue_expr_simple_refContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Lvalue_expr_simple_refContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterLvalue_expr_simple_ref(s)
	}
}

func (s *Lvalue_expr_simple_refContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitLvalue_expr_simple_ref(s)
	}
}

type Lvalue_expr_indexContext struct {
	*Lvalue_exprContext
}

func NewLvalue_expr_indexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Lvalue_expr_indexContext {
	var p = new(Lvalue_expr_indexContext)

	p.Lvalue_exprContext = NewEmptyLvalue_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Lvalue_exprContext))

	return p
}

func (s *Lvalue_expr_indexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lvalue_expr_indexContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Lvalue_expr_indexContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserOBRACKET, 0)
}

func (s *Lvalue_expr_indexContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Lvalue_expr_indexContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserCBRACKET, 0)
}

func (s *Lvalue_expr_indexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterLvalue_expr_index(s)
	}
}

func (s *Lvalue_expr_indexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitLvalue_expr_index(s)
	}
}

type Lvalue_expr_member_accessContext struct {
	*Lvalue_exprContext
}

func NewLvalue_expr_member_accessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Lvalue_expr_member_accessContext {
	var p = new(Lvalue_expr_member_accessContext)

	p.Lvalue_exprContext = NewEmptyLvalue_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Lvalue_exprContext))

	return p
}

func (s *Lvalue_expr_member_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lvalue_expr_member_accessContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Lvalue_expr_member_accessContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(NitroParserPERIOD, 0)
}

func (s *Lvalue_expr_member_accessContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Lvalue_expr_member_accessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterLvalue_expr_member_access(s)
	}
}

func (s *Lvalue_expr_member_accessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitLvalue_expr_member_access(s)
	}
}

func (p *NitroParser) Lvalue_expr() (localctx ILvalue_exprContext) {
	this := p
	_ = this

	localctx = NewLvalue_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, NitroParserRULE_lvalue_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLvalue_expr_simple_refContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(483)
			p.Match(NitroParserID)
		}

	case 2:
		localctx = NewLvalue_expr_member_accessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(484)
			p.primary_expr(0)
		}
		{
			p.SetState(485)
			p.Match(NitroParserPERIOD)
		}
		{
			p.SetState(486)
			p.Match(NitroParserID)
		}

	case 3:
		localctx = NewLvalue_expr_indexContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(488)
			p.primary_expr(0)
		}
		{
			p.SetState(489)
			p.Match(NitroParserOBRACKET)
		}
		{
			p.SetState(490)
			p.Expr()
		}
		{
			p.SetState(491)
			p.Match(NitroParserCBRACKET)
		}

	}

	return localctx
}

// ILambda_exprContext is an interface to support dynamic dispatch.
type ILambda_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambda_exprContext differentiates from other interfaces.
	IsLambda_exprContext()
}

type Lambda_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambda_exprContext() *Lambda_exprContext {
	var p = new(Lambda_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_lambda_expr
	return p
}

func (*Lambda_exprContext) IsLambda_exprContext() {}

func NewLambda_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lambda_exprContext {
	var p = new(Lambda_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_lambda_expr

	return p
}

func (s *Lambda_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Lambda_exprContext) FUNC() antlr.TerminalNode {
	return s.GetToken(NitroParserFUNC, 0)
}

func (s *Lambda_exprContext) OPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserOPAREN, 0)
}

func (s *Lambda_exprContext) CPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserCPAREN, 0)
}

func (s *Lambda_exprContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Lambda_exprContext) Stmts() IStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *Lambda_exprContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *Lambda_exprContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *Lambda_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lambda_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lambda_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterLambda_expr(s)
	}
}

func (s *Lambda_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitLambda_expr(s)
	}
}

func (p *NitroParser) Lambda_expr() (localctx ILambda_exprContext) {
	this := p
	_ = this

	localctx = NewLambda_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, NitroParserRULE_lambda_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)
		p.Match(NitroParserFUNC)
	}
	{
		p.SetState(496)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(497)
			p.Param_list()
		}

	}
	{
		p.SetState(500)
		p.Match(NitroParserCPAREN)
	}
	{
		p.SetState(501)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(502)
		p.Stmts()
	}
	{
		p.SetState(503)
		p.Match(NitroParserCCURLY)
	}

	return localctx
}

// IShort_lambda_exprContext is an interface to support dynamic dispatch.
type IShort_lambda_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShort_lambda_exprContext differentiates from other interfaces.
	IsShort_lambda_exprContext()
}

type Short_lambda_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShort_lambda_exprContext() *Short_lambda_exprContext {
	var p = new(Short_lambda_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_short_lambda_expr
	return p
}

func (*Short_lambda_exprContext) IsShort_lambda_exprContext() {}

func NewShort_lambda_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Short_lambda_exprContext {
	var p = new(Short_lambda_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_short_lambda_expr

	return p
}

func (s *Short_lambda_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Short_lambda_exprContext) LAMBDA() antlr.TerminalNode {
	return s.GetToken(NitroParserLAMBDA, 0)
}

func (s *Short_lambda_exprContext) ARROW() antlr.TerminalNode {
	return s.GetToken(NitroParserARROW, 0)
}

func (s *Short_lambda_exprContext) Expr3() IExpr3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr3Context)
}

func (s *Short_lambda_exprContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *Short_lambda_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Short_lambda_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Short_lambda_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterShort_lambda_expr(s)
	}
}

func (s *Short_lambda_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitShort_lambda_expr(s)
	}
}

func (p *NitroParser) Short_lambda_expr() (localctx IShort_lambda_exprContext) {
	this := p
	_ = this

	localctx = NewShort_lambda_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, NitroParserRULE_short_lambda_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)
		p.Match(NitroParserLAMBDA)
	}
	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(506)
			p.Param_list()
		}

	}
	{
		p.SetState(509)
		p.Match(NitroParserARROW)
	}
	{
		p.SetState(510)
		p.expr3(0)
	}

	return localctx
}

// IExec_exprContext is an interface to support dynamic dispatch.
type IExec_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExec_exprContext differentiates from other interfaces.
	IsExec_exprContext()
}

type Exec_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExec_exprContext() *Exec_exprContext {
	var p = new(Exec_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_exec_expr
	return p
}

func (*Exec_exprContext) IsExec_exprContext() {}

func NewExec_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exec_exprContext {
	var p = new(Exec_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_exec_expr

	return p
}

func (s *Exec_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Exec_exprContext) EXEC_PREFIX() antlr.TerminalNode {
	return s.GetToken(NitroParserEXEC_PREFIX, 0)
}

func (s *Exec_exprContext) EXEC_SUFFIX() antlr.TerminalNode {
	return s.GetToken(NitroParserEXEC_SUFFIX, 0)
}

func (s *Exec_exprContext) AllExec_expr_arg() []IExec_expr_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExec_expr_argContext); ok {
			len++
		}
	}

	tst := make([]IExec_expr_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExec_expr_argContext); ok {
			tst[i] = t.(IExec_expr_argContext)
			i++
		}
	}

	return tst
}

func (s *Exec_exprContext) Exec_expr_arg(i int) IExec_expr_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExec_expr_argContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExec_expr_argContext)
}

func (s *Exec_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exec_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exec_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterExec_expr(s)
	}
}

func (s *Exec_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitExec_expr(s)
	}
}

func (p *NitroParser) Exec_expr() (localctx IExec_exprContext) {
	this := p
	_ = this

	localctx = NewExec_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, NitroParserRULE_exec_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.Match(NitroParserEXEC_PREFIX)
	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(NitroParserOCURLY-53))|(1<<(NitroParserEXEC_WS-53))|(1<<(NitroParserEXEC_HOME-53))|(1<<(NitroParserEXEC_LITERAL-53))|(1<<(NitroParserEXEC_DQUOTE_LITERAL-53))|(1<<(NitroParserEXEC_SQUOTE_LITERAL-53)))) != 0) {
		{
			p.SetState(513)
			p.Exec_expr_arg()
		}

		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(518)
		p.Match(NitroParserEXEC_SUFFIX)
	}

	return localctx
}

// IExec_expr_argContext is an interface to support dynamic dispatch.
type IExec_expr_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExec_expr_argContext differentiates from other interfaces.
	IsExec_expr_argContext()
}

type Exec_expr_argContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExec_expr_argContext() *Exec_expr_argContext {
	var p = new(Exec_expr_argContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_exec_expr_arg
	return p
}

func (*Exec_expr_argContext) IsExec_expr_argContext() {}

func NewExec_expr_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exec_expr_argContext {
	var p = new(Exec_expr_argContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_exec_expr_arg

	return p
}

func (s *Exec_expr_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Exec_expr_argContext) EXEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(NitroParserEXEC_LITERAL, 0)
}

func (s *Exec_expr_argContext) EXEC_DQUOTE_LITERAL() antlr.TerminalNode {
	return s.GetToken(NitroParserEXEC_DQUOTE_LITERAL, 0)
}

func (s *Exec_expr_argContext) EXEC_SQUOTE_LITERAL() antlr.TerminalNode {
	return s.GetToken(NitroParserEXEC_SQUOTE_LITERAL, 0)
}

func (s *Exec_expr_argContext) EXEC_WS() antlr.TerminalNode {
	return s.GetToken(NitroParserEXEC_WS, 0)
}

func (s *Exec_expr_argContext) EXEC_HOME() antlr.TerminalNode {
	return s.GetToken(NitroParserEXEC_HOME, 0)
}

func (s *Exec_expr_argContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Exec_expr_argContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Exec_expr_argContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *Exec_expr_argContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(NitroParserEXPAND, 0)
}

func (s *Exec_expr_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exec_expr_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exec_expr_argContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterExec_expr_arg(s)
	}
}

func (s *Exec_expr_argContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitExec_expr_arg(s)
	}
}

func (p *NitroParser) Exec_expr_arg() (localctx IExec_expr_argContext) {
	this := p
	_ = this

	localctx = NewExec_expr_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, NitroParserRULE_exec_expr_arg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(532)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserEXEC_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(520)
			p.Match(NitroParserEXEC_LITERAL)
		}

	case NitroParserEXEC_DQUOTE_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(521)
			p.Match(NitroParserEXEC_DQUOTE_LITERAL)
		}

	case NitroParserEXEC_SQUOTE_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(522)
			p.Match(NitroParserEXEC_SQUOTE_LITERAL)
		}

	case NitroParserEXEC_WS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(523)
			p.Match(NitroParserEXEC_WS)
		}

	case NitroParserEXEC_HOME:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(524)
			p.Match(NitroParserEXEC_HOME)
		}

	case NitroParserOCURLY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(525)
			p.Match(NitroParserOCURLY)
		}
		{
			p.SetState(526)
			p.Expr()
		}
		p.SetState(528)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NitroParserEXPAND {
			{
				p.SetState(527)
				p.Match(NitroParserEXPAND)
			}

		}
		{
			p.SetState(530)
			p.Match(NitroParserCCURLY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObject_literalContext is an interface to support dynamic dispatch.
type IObject_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_literalContext differentiates from other interfaces.
	IsObject_literalContext()
}

type Object_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_literalContext() *Object_literalContext {
	var p = new(Object_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_object_literal
	return p
}

func (*Object_literalContext) IsObject_literalContext() {}

func NewObject_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_literalContext {
	var p = new(Object_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_object_literal

	return p
}

func (s *Object_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_literalContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Object_literalContext) Object_fields() IObject_fieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_fieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_literalContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *Object_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_literal(s)
	}
}

func (s *Object_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_literal(s)
	}
}

func (p *NitroParser) Object_literal() (localctx IObject_literalContext) {
	this := p
	_ = this

	localctx = NewObject_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, NitroParserRULE_object_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(535)
		p.Object_fields()
	}
	{
		p.SetState(536)
		p.Match(NitroParserCCURLY)
	}

	return localctx
}

// IObject_fieldsContext is an interface to support dynamic dispatch.
type IObject_fieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_fieldsContext differentiates from other interfaces.
	IsObject_fieldsContext()
}

type Object_fieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_fieldsContext() *Object_fieldsContext {
	var p = new(Object_fieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_object_fields
	return p
}

func (*Object_fieldsContext) IsObject_fieldsContext() {}

func NewObject_fieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_fieldsContext {
	var p = new(Object_fieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_object_fields

	return p
}

func (s *Object_fieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_fieldsContext) AllObject_field() []IObject_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObject_fieldContext); ok {
			len++
		}
	}

	tst := make([]IObject_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObject_fieldContext); ok {
			tst[i] = t.(IObject_fieldContext)
			i++
		}
	}

	return tst
}

func (s *Object_fieldsContext) Object_field(i int) IObject_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_fieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_fieldContext)
}

func (s *Object_fieldsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCOMMA)
}

func (s *Object_fieldsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCOMMA, i)
}

func (s *Object_fieldsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(NitroParserSEMICOLON)
}

func (s *Object_fieldsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, i)
}

func (s *Object_fieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_fieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_fieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_fields(s)
	}
}

func (s *Object_fieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_fields(s)
	}
}

func (p *NitroParser) Object_fields() (localctx IObject_fieldsContext) {
	this := p
	_ = this

	localctx = NewObject_fieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, NitroParserRULE_object_fields)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserAND)|(1<<NitroParserBREAK)|(1<<NitroParserCATCH)|(1<<NitroParserCONTINUE)|(1<<NitroParserDEFER)|(1<<NitroParserELSE)|(1<<NitroParserFALSE)|(1<<NitroParserFOR)|(1<<NitroParserFUNC)|(1<<NitroParserIF)|(1<<NitroParserIMPORT)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserOR)|(1<<NitroParserRETURN)|(1<<NitroParserTHROW)|(1<<NitroParserTRUE)|(1<<NitroParserTRY)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserYIELD))) != 0) || (((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(NitroParserOPAREN-49))|(1<<(NitroParserOBRACKET-49))|(1<<(NitroParserOCURLY-49))|(1<<(NitroParserNUMBER-49))|(1<<(NitroParserID-49))|(1<<(NitroParserREGEX-49))|(1<<(NitroParserCHAR-49))|(1<<(NitroParserSTRING-49))|(1<<(NitroParserEXEC_PREFIX-49)))) != 0) {
		{
			p.SetState(538)
			p.Object_field()
		}
		p.SetState(543)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(539)
					_la = p.GetTokenStream().LA(1)

					if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(540)
					p.Object_field()
				}

			}
			p.SetState(545)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
		}
		p.SetState(547)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
			{
				p.SetState(546)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	}

	return localctx
}

// IObject_fieldContext is an interface to support dynamic dispatch.
type IObject_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_fieldContext differentiates from other interfaces.
	IsObject_fieldContext()
}

type Object_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_fieldContext() *Object_fieldContext {
	var p = new(Object_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_object_field
	return p
}

func (*Object_fieldContext) IsObject_fieldContext() {}

func NewObject_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_fieldContext {
	var p = new(Object_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_object_field

	return p
}

func (s *Object_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_fieldContext) CopyFrom(ctx *Object_fieldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Object_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Object_field_id_keyContext struct {
	*Object_fieldContext
}

func NewObject_field_id_keyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Object_field_id_keyContext {
	var p = new(Object_field_id_keyContext)

	p.Object_fieldContext = NewEmptyObject_fieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Object_fieldContext))

	return p
}

func (s *Object_field_id_keyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_field_id_keyContext) Id_or_keyword() IId_or_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_or_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_or_keywordContext)
}

func (s *Object_field_id_keyContext) COLON() antlr.TerminalNode {
	return s.GetToken(NitroParserCOLON, 0)
}

func (s *Object_field_id_keyContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Object_field_id_keyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_field_id_key(s)
	}
}

func (s *Object_field_id_keyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_field_id_key(s)
	}
}

type Object_field_expansionContext struct {
	*Object_fieldContext
}

func NewObject_field_expansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Object_field_expansionContext {
	var p = new(Object_field_expansionContext)

	p.Object_fieldContext = NewEmptyObject_fieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Object_fieldContext))

	return p
}

func (s *Object_field_expansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_field_expansionContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Object_field_expansionContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(NitroParserEXPAND, 0)
}

func (s *Object_field_expansionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_field_expansion(s)
	}
}

func (s *Object_field_expansionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_field_expansion(s)
	}
}

type Object_field_expr_keyContext struct {
	*Object_fieldContext
}

func NewObject_field_expr_keyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Object_field_expr_keyContext {
	var p = new(Object_field_expr_keyContext)

	p.Object_fieldContext = NewEmptyObject_fieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Object_fieldContext))

	return p
}

func (s *Object_field_expr_keyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_field_expr_keyContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserOBRACKET, 0)
}

func (s *Object_field_expr_keyContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Object_field_expr_keyContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Object_field_expr_keyContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserCBRACKET, 0)
}

func (s *Object_field_expr_keyContext) COLON() antlr.TerminalNode {
	return s.GetToken(NitroParserCOLON, 0)
}

func (s *Object_field_expr_keyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_field_expr_key(s)
	}
}

func (s *Object_field_expr_keyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_field_expr_key(s)
	}
}

func (p *NitroParser) Object_field() (localctx IObject_fieldContext) {
	this := p
	_ = this

	localctx = NewObject_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, NitroParserRULE_object_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(564)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		localctx = NewObject_field_id_keyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(551)
			p.Id_or_keyword()
		}
		{
			p.SetState(552)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(553)
			p.Expr()
		}

	case 2:
		localctx = NewObject_field_expr_keyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(555)
			p.Match(NitroParserOBRACKET)
		}
		{
			p.SetState(556)
			p.Expr()
		}
		{
			p.SetState(557)
			p.Match(NitroParserCBRACKET)
		}
		{
			p.SetState(558)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(559)
			p.Expr()
		}

	case 3:
		localctx = NewObject_field_expansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(561)
			p.primary_expr(0)
		}
		{
			p.SetState(562)
			p.Match(NitroParserEXPAND)
		}

	}

	return localctx
}

// IArray_literalContext is an interface to support dynamic dispatch.
type IArray_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_literalContext differentiates from other interfaces.
	IsArray_literalContext()
}

type Array_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_literalContext() *Array_literalContext {
	var p = new(Array_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_array_literal
	return p
}

func (*Array_literalContext) IsArray_literalContext() {}

func NewArray_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_literalContext {
	var p = new(Array_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_array_literal

	return p
}

func (s *Array_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_literalContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserOBRACKET, 0)
}

func (s *Array_literalContext) Array_elems() IArray_elemsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_elemsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_literalContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserCBRACKET, 0)
}

func (s *Array_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterArray_literal(s)
	}
}

func (s *Array_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitArray_literal(s)
	}
}

func (p *NitroParser) Array_literal() (localctx IArray_literalContext) {
	this := p
	_ = this

	localctx = NewArray_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, NitroParserRULE_array_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(566)
		p.Match(NitroParserOBRACKET)
	}
	{
		p.SetState(567)
		p.Array_elems()
	}
	{
		p.SetState(568)
		p.Match(NitroParserCBRACKET)
	}

	return localctx
}

// IArray_elemsContext is an interface to support dynamic dispatch.
type IArray_elemsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_elemsContext differentiates from other interfaces.
	IsArray_elemsContext()
}

type Array_elemsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_elemsContext() *Array_elemsContext {
	var p = new(Array_elemsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_array_elems
	return p
}

func (*Array_elemsContext) IsArray_elemsContext() {}

func NewArray_elemsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_elemsContext {
	var p = new(Array_elemsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_array_elems

	return p
}

func (s *Array_elemsContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_elemsContext) AllArray_elem() []IArray_elemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArray_elemContext); ok {
			len++
		}
	}

	tst := make([]IArray_elemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArray_elemContext); ok {
			tst[i] = t.(IArray_elemContext)
			i++
		}
	}

	return tst
}

func (s *Array_elemsContext) Array_elem(i int) IArray_elemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_elemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_elemContext)
}

func (s *Array_elemsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCOMMA)
}

func (s *Array_elemsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCOMMA, i)
}

func (s *Array_elemsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(NitroParserSEMICOLON)
}

func (s *Array_elemsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, i)
}

func (s *Array_elemsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_elemsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_elemsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterArray_elems(s)
	}
}

func (s *Array_elemsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitArray_elems(s)
	}
}

func (p *NitroParser) Array_elems() (localctx IArray_elemsContext) {
	this := p
	_ = this

	localctx = NewArray_elemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, NitroParserRULE_array_elems)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(584)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37))|(1<<(NitroParserEXEC_PREFIX-37)))) != 0) {
		{
			p.SetState(570)
			p.Array_elem()
		}
		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(571)
					_la = p.GetTokenStream().LA(1)

					if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(572)
					p.Array_elem()
				}

			}
			p.SetState(577)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())
		}
		p.SetState(581)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
			{
				p.SetState(578)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(583)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IArray_elemContext is an interface to support dynamic dispatch.
type IArray_elemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_elemContext differentiates from other interfaces.
	IsArray_elemContext()
}

type Array_elemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_elemContext() *Array_elemContext {
	var p = new(Array_elemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_array_elem
	return p
}

func (*Array_elemContext) IsArray_elemContext() {}

func NewArray_elemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_elemContext {
	var p = new(Array_elemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_array_elem

	return p
}

func (s *Array_elemContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_elemContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Array_elemContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(NitroParserEXPAND, 0)
}

func (s *Array_elemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_elemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_elemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterArray_elem(s)
	}
}

func (s *Array_elemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitArray_elem(s)
	}
}

func (p *NitroParser) Array_elem() (localctx IArray_elemContext) {
	this := p
	_ = this

	localctx = NewArray_elemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, NitroParserRULE_array_elem)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(586)
		p.Expr()
	}
	p.SetState(588)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserEXPAND {
		{
			p.SetState(587)
			p.Match(NitroParserEXPAND)
		}

	}

	return localctx
}

// IId_or_keywordContext is an interface to support dynamic dispatch.
type IId_or_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// IsId_or_keywordContext differentiates from other interfaces.
	IsId_or_keywordContext()
}

type Id_or_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	t      antlr.Token
}

func NewEmptyId_or_keywordContext() *Id_or_keywordContext {
	var p = new(Id_or_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_id_or_keyword
	return p
}

func (*Id_or_keywordContext) IsId_or_keywordContext() {}

func NewId_or_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_or_keywordContext {
	var p = new(Id_or_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_id_or_keyword

	return p
}

func (s *Id_or_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_or_keywordContext) GetT() antlr.Token { return s.t }

func (s *Id_or_keywordContext) SetT(v antlr.Token) { s.t = v }

func (s *Id_or_keywordContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Id_or_keywordContext) AND() antlr.TerminalNode {
	return s.GetToken(NitroParserAND, 0)
}

func (s *Id_or_keywordContext) BREAK() antlr.TerminalNode {
	return s.GetToken(NitroParserBREAK, 0)
}

func (s *Id_or_keywordContext) CATCH() antlr.TerminalNode {
	return s.GetToken(NitroParserCATCH, 0)
}

func (s *Id_or_keywordContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(NitroParserCONTINUE, 0)
}

func (s *Id_or_keywordContext) DEFER() antlr.TerminalNode {
	return s.GetToken(NitroParserDEFER, 0)
}

func (s *Id_or_keywordContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NitroParserELSE, 0)
}

func (s *Id_or_keywordContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NitroParserFALSE, 0)
}

func (s *Id_or_keywordContext) FOR() antlr.TerminalNode {
	return s.GetToken(NitroParserFOR, 0)
}

func (s *Id_or_keywordContext) FUNC() antlr.TerminalNode {
	return s.GetToken(NitroParserFUNC, 0)
}

func (s *Id_or_keywordContext) IF() antlr.TerminalNode {
	return s.GetToken(NitroParserIF, 0)
}

func (s *Id_or_keywordContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(NitroParserIMPORT, 0)
}

func (s *Id_or_keywordContext) NIL() antlr.TerminalNode {
	return s.GetToken(NitroParserNIL, 0)
}

func (s *Id_or_keywordContext) NOT() antlr.TerminalNode {
	return s.GetToken(NitroParserNOT, 0)
}

func (s *Id_or_keywordContext) OR() antlr.TerminalNode {
	return s.GetToken(NitroParserOR, 0)
}

func (s *Id_or_keywordContext) RETURN() antlr.TerminalNode {
	return s.GetToken(NitroParserRETURN, 0)
}

func (s *Id_or_keywordContext) THROW() antlr.TerminalNode {
	return s.GetToken(NitroParserTHROW, 0)
}

func (s *Id_or_keywordContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NitroParserTRUE, 0)
}

func (s *Id_or_keywordContext) TRY() antlr.TerminalNode {
	return s.GetToken(NitroParserTRY, 0)
}

func (s *Id_or_keywordContext) VAR() antlr.TerminalNode {
	return s.GetToken(NitroParserVAR, 0)
}

func (s *Id_or_keywordContext) WHILE() antlr.TerminalNode {
	return s.GetToken(NitroParserWHILE, 0)
}

func (s *Id_or_keywordContext) YIELD() antlr.TerminalNode {
	return s.GetToken(NitroParserYIELD, 0)
}

func (s *Id_or_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_or_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_or_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterId_or_keyword(s)
	}
}

func (s *Id_or_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitId_or_keyword(s)
	}
}

func (p *NitroParser) Id_or_keyword() (localctx IId_or_keywordContext) {
	this := p
	_ = this

	localctx = NewId_or_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, NitroParserRULE_id_or_keyword)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(590)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Id_or_keywordContext).t = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserAND)|(1<<NitroParserBREAK)|(1<<NitroParserCATCH)|(1<<NitroParserCONTINUE)|(1<<NitroParserDEFER)|(1<<NitroParserELSE)|(1<<NitroParserFALSE)|(1<<NitroParserFOR)|(1<<NitroParserFUNC)|(1<<NitroParserIF)|(1<<NitroParserIMPORT)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserOR)|(1<<NitroParserRETURN)|(1<<NitroParserTHROW)|(1<<NitroParserTRUE)|(1<<NitroParserTRY)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserYIELD))) != 0) || _la == NitroParserID) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Id_or_keywordContext).t = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *NitroParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 37:
		var t *Expr3Context = nil
		if localctx != nil {
			t = localctx.(*Expr3Context)
		}
		return p.Expr3_Sempred(t, predIndex)

	case 38:
		var t *Binary_exprContext = nil
		if localctx != nil {
			t = localctx.(*Binary_exprContext)
		}
		return p.Binary_expr_Sempred(t, predIndex)

	case 40:
		var t *Primary_exprContext = nil
		if localctx != nil {
			t = localctx.(*Primary_exprContext)
		}
		return p.Primary_expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NitroParser) Expr3_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *NitroParser) Binary_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *NitroParser) Primary_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
