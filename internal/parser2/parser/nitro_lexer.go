// Code generated from NitroLexer.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 60, 391,
	8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6,
	4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12,
	9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9,
	17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22,
	4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4,
	28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33,
	9, 33, 4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9,
	38, 4, 39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43,
	4, 44, 9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4,
	49, 9, 49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54,
	9, 54, 4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9,
	59, 4, 60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37,
	3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3,
	42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47,
	3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 6,
	51, 304, 10, 51, 13, 51, 14, 51, 305, 3, 51, 3, 51, 6, 51, 310, 10, 51,
	13, 51, 14, 51, 311, 5, 51, 314, 10, 51, 3, 52, 3, 52, 7, 52, 318, 10,
	52, 12, 52, 14, 52, 321, 11, 52, 3, 53, 3, 53, 3, 53, 3, 53, 7, 53, 327,
	10, 53, 12, 53, 14, 53, 330, 11, 53, 3, 53, 3, 53, 3, 54, 5, 54, 335, 10,
	54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 6, 55, 343, 10, 55, 13, 55,
	14, 55, 344, 5, 55, 347, 10, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3,
	57, 7, 57, 355, 10, 57, 12, 57, 14, 57, 358, 11, 57, 3, 57, 3, 57, 3, 58,
	6, 58, 363, 10, 58, 13, 58, 14, 58, 364, 3, 58, 3, 58, 3, 59, 3, 59, 3,
	59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 61,
	3, 61, 3, 61, 5, 61, 384, 10, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3,
	62, 2, 2, 63, 4, 3, 6, 4, 8, 5, 10, 6, 12, 7, 14, 8, 16, 9, 18, 10, 20,
	11, 22, 12, 24, 13, 26, 14, 28, 15, 30, 16, 32, 17, 34, 18, 36, 19, 38,
	20, 40, 21, 42, 22, 44, 23, 46, 24, 48, 25, 50, 26, 52, 27, 54, 28, 56,
	29, 58, 30, 60, 31, 62, 32, 64, 33, 66, 34, 68, 35, 70, 36, 72, 37, 74,
	38, 76, 39, 78, 40, 80, 41, 82, 42, 84, 43, 86, 44, 88, 45, 90, 46, 92,
	47, 94, 48, 96, 49, 98, 50, 100, 51, 102, 52, 104, 53, 106, 54, 108, 55,
	110, 56, 112, 2, 114, 57, 116, 58, 118, 60, 120, 2, 122, 2, 124, 59, 4,
	2, 3, 13, 3, 2, 50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67,
	92, 97, 97, 99, 124, 3, 2, 98, 98, 7, 2, 41, 41, 94, 94, 112, 112, 116,
	116, 118, 118, 6, 2, 12, 12, 15, 15, 41, 41, 94, 94, 5, 2, 50, 59, 67,
	72, 99, 104, 3, 2, 12, 12, 4, 2, 11, 11, 34, 34, 6, 2, 12, 12, 15, 15,
	36, 36, 94, 94, 7, 2, 36, 36, 94, 94, 112, 112, 116, 116, 118, 118, 2,
	399, 2, 4, 3, 2, 2, 2, 2, 6, 3, 2, 2, 2, 2, 8, 3, 2, 2, 2, 2, 10, 3, 2,
	2, 2, 2, 12, 3, 2, 2, 2, 2, 14, 3, 2, 2, 2, 2, 16, 3, 2, 2, 2, 2, 18, 3,
	2, 2, 2, 2, 20, 3, 2, 2, 2, 2, 22, 3, 2, 2, 2, 2, 24, 3, 2, 2, 2, 2, 26,
	3, 2, 2, 2, 2, 28, 3, 2, 2, 2, 2, 30, 3, 2, 2, 2, 2, 32, 3, 2, 2, 2, 2,
	34, 3, 2, 2, 2, 2, 36, 3, 2, 2, 2, 2, 38, 3, 2, 2, 2, 2, 40, 3, 2, 2, 2,
	2, 42, 3, 2, 2, 2, 2, 44, 3, 2, 2, 2, 2, 46, 3, 2, 2, 2, 2, 48, 3, 2, 2,
	2, 2, 50, 3, 2, 2, 2, 2, 52, 3, 2, 2, 2, 2, 54, 3, 2, 2, 2, 2, 56, 3, 2,
	2, 2, 2, 58, 3, 2, 2, 2, 2, 60, 3, 2, 2, 2, 2, 62, 3, 2, 2, 2, 2, 64, 3,
	2, 2, 2, 2, 66, 3, 2, 2, 2, 2, 68, 3, 2, 2, 2, 2, 70, 3, 2, 2, 2, 2, 72,
	3, 2, 2, 2, 2, 74, 3, 2, 2, 2, 2, 76, 3, 2, 2, 2, 2, 78, 3, 2, 2, 2, 2,
	80, 3, 2, 2, 2, 2, 82, 3, 2, 2, 2, 2, 84, 3, 2, 2, 2, 2, 86, 3, 2, 2, 2,
	2, 88, 3, 2, 2, 2, 2, 90, 3, 2, 2, 2, 2, 92, 3, 2, 2, 2, 2, 94, 3, 2, 2,
	2, 2, 96, 3, 2, 2, 2, 2, 98, 3, 2, 2, 2, 2, 100, 3, 2, 2, 2, 2, 102, 3,
	2, 2, 2, 2, 104, 3, 2, 2, 2, 2, 106, 3, 2, 2, 2, 2, 108, 3, 2, 2, 2, 2,
	110, 3, 2, 2, 2, 2, 114, 3, 2, 2, 2, 2, 116, 3, 2, 2, 2, 2, 118, 3, 2,
	2, 2, 3, 120, 3, 2, 2, 2, 3, 122, 3, 2, 2, 2, 3, 124, 3, 2, 2, 2, 4, 126,
	3, 2, 2, 2, 6, 130, 3, 2, 2, 2, 8, 136, 3, 2, 2, 2, 10, 142, 3, 2, 2, 2,
	12, 151, 3, 2, 2, 2, 14, 157, 3, 2, 2, 2, 16, 162, 3, 2, 2, 2, 18, 168,
	3, 2, 2, 2, 20, 172, 3, 2, 2, 2, 22, 177, 3, 2, 2, 2, 24, 180, 3, 2, 2,
	2, 26, 187, 3, 2, 2, 2, 28, 192, 3, 2, 2, 2, 30, 196, 3, 2, 2, 2, 32, 200,
	3, 2, 2, 2, 34, 203, 3, 2, 2, 2, 36, 210, 3, 2, 2, 2, 38, 216, 3, 2, 2,
	2, 40, 221, 3, 2, 2, 2, 42, 225, 3, 2, 2, 2, 44, 229, 3, 2, 2, 2, 46, 235,
	3, 2, 2, 2, 48, 241, 3, 2, 2, 2, 50, 243, 3, 2, 2, 2, 52, 246, 3, 2, 2,
	2, 54, 249, 3, 2, 2, 2, 56, 251, 3, 2, 2, 2, 58, 254, 3, 2, 2, 2, 60, 256,
	3, 2, 2, 2, 62, 259, 3, 2, 2, 2, 64, 261, 3, 2, 2, 2, 66, 263, 3, 2, 2,
	2, 68, 265, 3, 2, 2, 2, 70, 267, 3, 2, 2, 2, 72, 269, 3, 2, 2, 2, 74, 271,
	3, 2, 2, 2, 76, 273, 3, 2, 2, 2, 78, 275, 3, 2, 2, 2, 80, 277, 3, 2, 2,
	2, 82, 279, 3, 2, 2, 2, 84, 281, 3, 2, 2, 2, 86, 283, 3, 2, 2, 2, 88, 285,
	3, 2, 2, 2, 90, 287, 3, 2, 2, 2, 92, 289, 3, 2, 2, 2, 94, 291, 3, 2, 2,
	2, 96, 294, 3, 2, 2, 2, 98, 296, 3, 2, 2, 2, 100, 298, 3, 2, 2, 2, 102,
	303, 3, 2, 2, 2, 104, 315, 3, 2, 2, 2, 106, 322, 3, 2, 2, 2, 108, 334,
	3, 2, 2, 2, 110, 338, 3, 2, 2, 2, 112, 350, 3, 2, 2, 2, 114, 352, 3, 2,
	2, 2, 116, 362, 3, 2, 2, 2, 118, 368, 3, 2, 2, 2, 120, 373, 3, 2, 2, 2,
	122, 377, 3, 2, 2, 2, 124, 387, 3, 2, 2, 2, 126, 127, 7, 99, 2, 2, 127,
	128, 7, 112, 2, 2, 128, 129, 7, 102, 2, 2, 129, 5, 3, 2, 2, 2, 130, 131,
	7, 100, 2, 2, 131, 132, 7, 116, 2, 2, 132, 133, 7, 103, 2, 2, 133, 134,
	7, 99, 2, 2, 134, 135, 7, 109, 2, 2, 135, 7, 3, 2, 2, 2, 136, 137, 7, 101,
	2, 2, 137, 138, 7, 99, 2, 2, 138, 139, 7, 118, 2, 2, 139, 140, 7, 101,
	2, 2, 140, 141, 7, 106, 2, 2, 141, 9, 3, 2, 2, 2, 142, 143, 7, 101, 2,
	2, 143, 144, 7, 113, 2, 2, 144, 145, 7, 112, 2, 2, 145, 146, 7, 118, 2,
	2, 146, 147, 7, 107, 2, 2, 147, 148, 7, 112, 2, 2, 148, 149, 7, 119, 2,
	2, 149, 150, 7, 103, 2, 2, 150, 11, 3, 2, 2, 2, 151, 152, 7, 102, 2, 2,
	152, 153, 7, 103, 2, 2, 153, 154, 7, 104, 2, 2, 154, 155, 7, 103, 2, 2,
	155, 156, 7, 116, 2, 2, 156, 13, 3, 2, 2, 2, 157, 158, 7, 103, 2, 2, 158,
	159, 7, 110, 2, 2, 159, 160, 7, 117, 2, 2, 160, 161, 7, 103, 2, 2, 161,
	15, 3, 2, 2, 2, 162, 163, 7, 104, 2, 2, 163, 164, 7, 99, 2, 2, 164, 165,
	7, 110, 2, 2, 165, 166, 7, 117, 2, 2, 166, 167, 7, 103, 2, 2, 167, 17,
	3, 2, 2, 2, 168, 169, 7, 104, 2, 2, 169, 170, 7, 113, 2, 2, 170, 171, 7,
	116, 2, 2, 171, 19, 3, 2, 2, 2, 172, 173, 7, 104, 2, 2, 173, 174, 7, 119,
	2, 2, 174, 175, 7, 112, 2, 2, 175, 176, 7, 101, 2, 2, 176, 21, 3, 2, 2,
	2, 177, 178, 7, 107, 2, 2, 178, 179, 7, 104, 2, 2, 179, 23, 3, 2, 2, 2,
	180, 181, 7, 107, 2, 2, 181, 182, 7, 111, 2, 2, 182, 183, 7, 114, 2, 2,
	183, 184, 7, 113, 2, 2, 184, 185, 7, 116, 2, 2, 185, 186, 7, 118, 2, 2,
	186, 25, 3, 2, 2, 2, 187, 188, 7, 111, 2, 2, 188, 189, 7, 103, 2, 2, 189,
	190, 7, 118, 2, 2, 190, 191, 7, 99, 2, 2, 191, 27, 3, 2, 2, 2, 192, 193,
	7, 112, 2, 2, 193, 194, 7, 107, 2, 2, 194, 195, 7, 110, 2, 2, 195, 29,
	3, 2, 2, 2, 196, 197, 7, 112, 2, 2, 197, 198, 7, 113, 2, 2, 198, 199, 7,
	118, 2, 2, 199, 31, 3, 2, 2, 2, 200, 201, 7, 113, 2, 2, 201, 202, 7, 116,
	2, 2, 202, 33, 3, 2, 2, 2, 203, 204, 7, 116, 2, 2, 204, 205, 7, 103, 2,
	2, 205, 206, 7, 118, 2, 2, 206, 207, 7, 119, 2, 2, 207, 208, 7, 116, 2,
	2, 208, 209, 7, 112, 2, 2, 209, 35, 3, 2, 2, 2, 210, 211, 7, 118, 2, 2,
	211, 212, 7, 106, 2, 2, 212, 213, 7, 116, 2, 2, 213, 214, 7, 113, 2, 2,
	214, 215, 7, 121, 2, 2, 215, 37, 3, 2, 2, 2, 216, 217, 7, 118, 2, 2, 217,
	218, 7, 116, 2, 2, 218, 219, 7, 119, 2, 2, 219, 220, 7, 103, 2, 2, 220,
	39, 3, 2, 2, 2, 221, 222, 7, 118, 2, 2, 222, 223, 7, 116, 2, 2, 223, 224,
	7, 123, 2, 2, 224, 41, 3, 2, 2, 2, 225, 226, 7, 120, 2, 2, 226, 227, 7,
	99, 2, 2, 227, 228, 7, 116, 2, 2, 228, 43, 3, 2, 2, 2, 229, 230, 7, 121,
	2, 2, 230, 231, 7, 106, 2, 2, 231, 232, 7, 107, 2, 2, 232, 233, 7, 110,
	2, 2, 233, 234, 7, 103, 2, 2, 234, 45, 3, 2, 2, 2, 235, 236, 7, 123, 2,
	2, 236, 237, 7, 107, 2, 2, 237, 238, 7, 103, 2, 2, 238, 239, 7, 110, 2,
	2, 239, 240, 7, 102, 2, 2, 240, 47, 3, 2, 2, 2, 241, 242, 7, 63, 2, 2,
	242, 49, 3, 2, 2, 2, 243, 244, 7, 63, 2, 2, 244, 245, 7, 63, 2, 2, 245,
	51, 3, 2, 2, 2, 246, 247, 7, 35, 2, 2, 247, 248, 7, 63, 2, 2, 248, 53,
	3, 2, 2, 2, 249, 250, 7, 62, 2, 2, 250, 55, 3, 2, 2, 2, 251, 252, 7, 62,
	2, 2, 252, 253, 7, 63, 2, 2, 253, 57, 3, 2, 2, 2, 254, 255, 7, 64, 2, 2,
	255, 59, 3, 2, 2, 2, 256, 257, 7, 64, 2, 2, 257, 258, 7, 63, 2, 2, 258,
	61, 3, 2, 2, 2, 259, 260, 7, 45, 2, 2, 260, 63, 3, 2, 2, 2, 261, 262, 7,
	47, 2, 2, 262, 65, 3, 2, 2, 2, 263, 264, 7, 44, 2, 2, 264, 67, 3, 2, 2,
	2, 265, 266, 7, 49, 2, 2, 266, 69, 3, 2, 2, 2, 267, 268, 7, 39, 2, 2, 268,
	71, 3, 2, 2, 2, 269, 270, 7, 65, 2, 2, 270, 73, 3, 2, 2, 2, 271, 272, 7,
	61, 2, 2, 272, 75, 3, 2, 2, 2, 273, 274, 7, 46, 2, 2, 274, 77, 3, 2, 2,
	2, 275, 276, 7, 60, 2, 2, 276, 79, 3, 2, 2, 2, 277, 278, 7, 48, 2, 2, 278,
	81, 3, 2, 2, 2, 279, 280, 7, 42, 2, 2, 280, 83, 3, 2, 2, 2, 281, 282, 7,
	43, 2, 2, 282, 85, 3, 2, 2, 2, 283, 284, 7, 93, 2, 2, 284, 87, 3, 2, 2,
	2, 285, 286, 7, 95, 2, 2, 286, 89, 3, 2, 2, 2, 287, 288, 7, 125, 2, 2,
	288, 91, 3, 2, 2, 2, 289, 290, 7, 127, 2, 2, 290, 93, 3, 2, 2, 2, 291,
	292, 7, 47, 2, 2, 292, 293, 7, 64, 2, 2, 293, 95, 3, 2, 2, 2, 294, 295,
	7, 40, 2, 2, 295, 97, 3, 2, 2, 2, 296, 297, 7, 126, 2, 2, 297, 99, 3, 2,
	2, 2, 298, 299, 7, 48, 2, 2, 299, 300, 7, 48, 2, 2, 300, 301, 7, 48, 2,
	2, 301, 101, 3, 2, 2, 2, 302, 304, 9, 2, 2, 2, 303, 302, 3, 2, 2, 2, 304,
	305, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 313,
	3, 2, 2, 2, 307, 309, 7, 48, 2, 2, 308, 310, 9, 2, 2, 2, 309, 308, 3, 2,
	2, 2, 310, 311, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2,
	312, 314, 3, 2, 2, 2, 313, 307, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314,
	103, 3, 2, 2, 2, 315, 319, 9, 3, 2, 2, 316, 318, 9, 4, 2, 2, 317, 316,
	3, 2, 2, 2, 318, 321, 3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 319, 320, 3, 2,
	2, 2, 320, 105, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 322, 323, 7, 116, 2,
	2, 323, 324, 7, 98, 2, 2, 324, 328, 3, 2, 2, 2, 325, 327, 10, 5, 2, 2,
	326, 325, 3, 2, 2, 2, 327, 330, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 328,
	329, 3, 2, 2, 2, 329, 331, 3, 2, 2, 2, 330, 328, 3, 2, 2, 2, 331, 332,
	7, 98, 2, 2, 332, 107, 3, 2, 2, 2, 333, 335, 7, 15, 2, 2, 334, 333, 3,
	2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 336, 3, 2, 2, 2, 336, 337, 7, 12, 2,
	2, 337, 109, 3, 2, 2, 2, 338, 346, 7, 41, 2, 2, 339, 340, 7, 94, 2, 2,
	340, 347, 9, 6, 2, 2, 341, 343, 10, 7, 2, 2, 342, 341, 3, 2, 2, 2, 343,
	344, 3, 2, 2, 2, 344, 342, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 347,
	3, 2, 2, 2, 346, 339, 3, 2, 2, 2, 346, 342, 3, 2, 2, 2, 347, 348, 3, 2,
	2, 2, 348, 349, 7, 41, 2, 2, 349, 111, 3, 2, 2, 2, 350, 351, 9, 8, 2, 2,
	351, 113, 3, 2, 2, 2, 352, 356, 7, 37, 2, 2, 353, 355, 10, 9, 2, 2, 354,
	353, 3, 2, 2, 2, 355, 358, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 356, 357,
	3, 2, 2, 2, 357, 359, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 359, 360, 8, 57,
	2, 2, 360, 115, 3, 2, 2, 2, 361, 363, 9, 10, 2, 2, 362, 361, 3, 2, 2, 2,
	363, 364, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365,
	366, 3, 2, 2, 2, 366, 367, 8, 58, 2, 2, 367, 117, 3, 2, 2, 2, 368, 369,
	7, 36, 2, 2, 369, 370, 3, 2, 2, 2, 370, 371, 8, 59, 3, 2, 371, 372, 8,
	59, 4, 2, 372, 119, 3, 2, 2, 2, 373, 374, 10, 11, 2, 2, 374, 375, 3, 2,
	2, 2, 375, 376, 8, 60, 3, 2, 376, 121, 3, 2, 2, 2, 377, 383, 7, 94, 2,
	2, 378, 384, 9, 12, 2, 2, 379, 380, 7, 122, 2, 2, 380, 381, 5, 112, 56,
	2, 381, 382, 5, 112, 56, 2, 382, 384, 3, 2, 2, 2, 383, 378, 3, 2, 2, 2,
	383, 379, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 386, 8, 61, 3, 2, 386,
	123, 3, 2, 2, 2, 387, 388, 7, 36, 2, 2, 388, 389, 3, 2, 2, 2, 389, 390,
	8, 62, 5, 2, 390, 125, 3, 2, 2, 2, 15, 2, 3, 305, 311, 313, 319, 328, 334,
	344, 346, 356, 364, 383, 6, 8, 2, 2, 5, 2, 2, 4, 3, 2, 4, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "QSTR",
}

var lexerLiteralNames = []string{
	"", "'and'", "'break'", "'catch'", "'continue'", "'defer'", "'else'", "'false'",
	"'for'", "'func'", "'if'", "'import'", "'meta'", "'nil'", "'not'", "'or'",
	"'return'", "'throw'", "'true'", "'try'", "'var'", "'while'", "'yield'",
	"'='", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'+'", "'-'", "'*'",
	"'/'", "'%'", "'?'", "';'", "','", "':'", "'.'", "'('", "')'", "'['", "']'",
	"'{'", "'}'", "'->'", "'&'", "'|'", "'...'",
}

var lexerSymbolicNames = []string{
	"", "AND", "BREAK", "CATCH", "CONTINUE", "DEFER", "ELSE", "FALSE", "FOR",
	"FUNC", "IF", "IMPORT", "META", "NIL", "NOT", "OR", "RETURN", "THROW",
	"TRUE", "TRY", "VAR", "WHILE", "YIELD", "ASSIGN", "EQ", "NE", "LT", "LE",
	"GT", "GE", "ADD", "SUB", "MUL", "DIV", "MOD", "QUESTION_MARK", "SEMICOLON",
	"COMMA", "COLON", "PERIOD", "OPAREN", "CPAREN", "OBRACKET", "CBRACKET",
	"OCURLY", "CCURLY", "ARROW", "LAMBDA", "PIPE", "EXPAND", "NUMBER", "ID",
	"REGEX", "NEWLINE", "CHAR", "COMMENT", "WS", "STRING", "LDQUOTE",
}

var lexerRuleNames = []string{
	"AND", "BREAK", "CATCH", "CONTINUE", "DEFER", "ELSE", "FALSE", "FOR", "FUNC",
	"IF", "IMPORT", "META", "NIL", "NOT", "OR", "RETURN", "THROW", "TRUE",
	"TRY", "VAR", "WHILE", "YIELD", "ASSIGN", "EQ", "NE", "LT", "LE", "GT",
	"GE", "ADD", "SUB", "MUL", "DIV", "MOD", "QUESTION_MARK", "SEMICOLON",
	"COMMA", "COLON", "PERIOD", "OPAREN", "CPAREN", "OBRACKET", "CBRACKET",
	"OCURLY", "CCURLY", "ARROW", "LAMBDA", "PIPE", "EXPAND", "NUMBER", "ID",
	"REGEX", "NEWLINE", "CHAR", "HEX_DIGIT", "COMMENT", "WS", "LDQUOTE", "QUOTED_TEXT",
	"QUOTED_ESCAPE", "STRING",
}

type NitroLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewNitroLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *NitroLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNitroLexer(input antlr.CharStream) *NitroLexer {
	l := new(NitroLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "NitroLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NitroLexer tokens.
const (
	NitroLexerAND           = 1
	NitroLexerBREAK         = 2
	NitroLexerCATCH         = 3
	NitroLexerCONTINUE      = 4
	NitroLexerDEFER         = 5
	NitroLexerELSE          = 6
	NitroLexerFALSE         = 7
	NitroLexerFOR           = 8
	NitroLexerFUNC          = 9
	NitroLexerIF            = 10
	NitroLexerIMPORT        = 11
	NitroLexerMETA          = 12
	NitroLexerNIL           = 13
	NitroLexerNOT           = 14
	NitroLexerOR            = 15
	NitroLexerRETURN        = 16
	NitroLexerTHROW         = 17
	NitroLexerTRUE          = 18
	NitroLexerTRY           = 19
	NitroLexerVAR           = 20
	NitroLexerWHILE         = 21
	NitroLexerYIELD         = 22
	NitroLexerASSIGN        = 23
	NitroLexerEQ            = 24
	NitroLexerNE            = 25
	NitroLexerLT            = 26
	NitroLexerLE            = 27
	NitroLexerGT            = 28
	NitroLexerGE            = 29
	NitroLexerADD           = 30
	NitroLexerSUB           = 31
	NitroLexerMUL           = 32
	NitroLexerDIV           = 33
	NitroLexerMOD           = 34
	NitroLexerQUESTION_MARK = 35
	NitroLexerSEMICOLON     = 36
	NitroLexerCOMMA         = 37
	NitroLexerCOLON         = 38
	NitroLexerPERIOD        = 39
	NitroLexerOPAREN        = 40
	NitroLexerCPAREN        = 41
	NitroLexerOBRACKET      = 42
	NitroLexerCBRACKET      = 43
	NitroLexerOCURLY        = 44
	NitroLexerCCURLY        = 45
	NitroLexerARROW         = 46
	NitroLexerLAMBDA        = 47
	NitroLexerPIPE          = 48
	NitroLexerEXPAND        = 49
	NitroLexerNUMBER        = 50
	NitroLexerID            = 51
	NitroLexerREGEX         = 52
	NitroLexerNEWLINE       = 53
	NitroLexerCHAR          = 54
	NitroLexerCOMMENT       = 55
	NitroLexerWS            = 56
	NitroLexerSTRING        = 57
	NitroLexerLDQUOTE       = 58
)

// NitroLexerQSTR is the NitroLexer mode.
const NitroLexerQSTR = 1
