// Code generated from NitroParser.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // NitroParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 522,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3, 2, 3,
	2, 3, 3, 5, 3, 101, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 107, 10, 4, 12,
	4, 14, 4, 110, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 119,
	10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 126, 10, 7, 12, 7, 14, 7, 129,
	11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 140,
	10, 9, 3, 10, 7, 10, 143, 10, 10, 12, 10, 14, 10, 146, 11, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 5, 11, 173, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 7, 13, 182, 10, 13, 12, 13, 14, 13, 185, 11, 13, 3, 14, 3,
	14, 3, 14, 7, 14, 190, 10, 14, 12, 14, 14, 14, 193, 11, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 5, 15, 199, 10, 15, 3, 16, 3, 16, 3, 16, 7, 16, 204, 10,
	16, 12, 16, 14, 16, 207, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 7, 18, 220, 10, 18, 12, 18, 14,
	18, 223, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 7, 20, 236, 10, 20, 12, 20, 14, 20, 239, 11, 20, 3,
	20, 5, 20, 242, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 258, 10, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 7, 24, 267, 10, 24, 12, 24,
	14, 24, 270, 11, 24, 3, 25, 3, 25, 3, 25, 5, 25, 275, 10, 25, 3, 25, 3,
	25, 3, 26, 3, 26, 5, 26, 281, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 7, 27, 301, 10, 27, 12, 27, 14, 27, 304, 11, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 313, 10, 28, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 325,
	10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 5, 29, 338, 10, 29, 3, 29, 3, 29, 5, 29, 342, 10, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 5, 29, 348, 10, 29, 3, 29, 7, 29, 351, 10, 29,
	12, 29, 14, 29, 354, 11, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 7, 31,
	361, 10, 31, 12, 31, 14, 31, 364, 11, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 376, 10, 32, 3, 33, 3, 33,
	3, 33, 5, 33, 381, 10, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 5,
	34, 389, 10, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 7, 35, 396, 10, 35,
	12, 35, 14, 35, 399, 11, 35, 3, 35, 7, 35, 402, 10, 35, 12, 35, 14, 35,
	405, 11, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 5, 36, 419, 10, 36, 3, 37, 3, 37, 3, 37, 3, 37,
	5, 37, 425, 10, 37, 3, 37, 7, 37, 428, 10, 37, 12, 37, 14, 37, 431, 11,
	37, 3, 37, 5, 37, 434, 10, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38,
	5, 38, 442, 10, 38, 3, 39, 3, 39, 5, 39, 446, 10, 39, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 40, 5, 40, 454, 10, 40, 3, 40, 3, 40, 3, 41, 3, 41,
	5, 41, 460, 10, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 7, 42, 467, 10,
	42, 12, 42, 14, 42, 470, 11, 42, 3, 42, 7, 42, 473, 10, 42, 12, 42, 14,
	42, 476, 11, 42, 3, 43, 3, 43, 3, 43, 5, 43, 481, 10, 43, 3, 44, 3, 44,
	3, 44, 3, 44, 5, 44, 487, 10, 44, 3, 44, 7, 44, 490, 10, 44, 12, 44, 14,
	44, 493, 11, 44, 3, 44, 5, 44, 496, 10, 44, 3, 44, 3, 44, 3, 45, 3, 45,
	3, 45, 3, 45, 5, 45, 504, 10, 45, 3, 46, 3, 46, 5, 46, 508, 10, 46, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47, 516, 10, 47, 3, 47, 3, 47,
	3, 48, 3, 48, 3, 48, 2, 4, 52, 56, 49, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90,
	92, 94, 2, 8, 3, 2, 33, 34, 3, 2, 30, 32, 3, 2, 28, 29, 3, 2, 22, 27, 6,
	2, 8, 8, 18, 18, 43, 43, 47, 47, 4, 2, 3, 20, 44, 44, 2, 549, 2, 96, 3,
	2, 2, 2, 4, 100, 3, 2, 2, 2, 6, 104, 3, 2, 2, 2, 8, 113, 3, 2, 2, 2, 10,
	116, 3, 2, 2, 2, 12, 122, 3, 2, 2, 2, 14, 130, 3, 2, 2, 2, 16, 139, 3,
	2, 2, 2, 18, 144, 3, 2, 2, 2, 20, 172, 3, 2, 2, 2, 22, 174, 3, 2, 2, 2,
	24, 178, 3, 2, 2, 2, 26, 186, 3, 2, 2, 2, 28, 194, 3, 2, 2, 2, 30, 200,
	3, 2, 2, 2, 32, 208, 3, 2, 2, 2, 34, 216, 3, 2, 2, 2, 36, 224, 3, 2, 2,
	2, 38, 230, 3, 2, 2, 2, 40, 245, 3, 2, 2, 2, 42, 250, 3, 2, 2, 2, 44, 253,
	3, 2, 2, 2, 46, 263, 3, 2, 2, 2, 48, 271, 3, 2, 2, 2, 50, 278, 3, 2, 2,
	2, 52, 282, 3, 2, 2, 2, 54, 312, 3, 2, 2, 2, 56, 324, 3, 2, 2, 2, 58, 355,
	3, 2, 2, 2, 60, 357, 3, 2, 2, 2, 62, 375, 3, 2, 2, 2, 64, 377, 3, 2, 2,
	2, 66, 386, 3, 2, 2, 2, 68, 392, 3, 2, 2, 2, 70, 418, 3, 2, 2, 2, 72, 420,
	3, 2, 2, 2, 74, 437, 3, 2, 2, 2, 76, 443, 3, 2, 2, 2, 78, 447, 3, 2, 2,
	2, 80, 457, 3, 2, 2, 2, 82, 463, 3, 2, 2, 2, 84, 480, 3, 2, 2, 2, 86, 482,
	3, 2, 2, 2, 88, 499, 3, 2, 2, 2, 90, 505, 3, 2, 2, 2, 92, 509, 3, 2, 2,
	2, 94, 519, 3, 2, 2, 2, 96, 97, 5, 4, 3, 2, 97, 98, 7, 2, 2, 3, 98, 3,
	3, 2, 2, 2, 99, 101, 5, 6, 4, 2, 100, 99, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 102, 3, 2, 2, 2, 102, 103, 5, 18, 10, 2, 103, 5, 3, 2, 2, 2, 104,
	108, 7, 13, 2, 2, 105, 107, 5, 8, 5, 2, 106, 105, 3, 2, 2, 2, 107, 110,
	3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 111, 3, 2,
	2, 2, 110, 108, 3, 2, 2, 2, 111, 112, 7, 7, 2, 2, 112, 7, 3, 2, 2, 2, 113,
	114, 7, 44, 2, 2, 114, 115, 5, 10, 6, 2, 115, 9, 3, 2, 2, 2, 116, 118,
	7, 41, 2, 2, 117, 119, 5, 12, 7, 2, 118, 117, 3, 2, 2, 2, 118, 119, 3,
	2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 121, 7, 42, 2, 2, 121, 11, 3, 2, 2,
	2, 122, 127, 5, 14, 8, 2, 123, 124, 9, 2, 2, 2, 124, 126, 5, 14, 8, 2,
	125, 123, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127,
	128, 3, 2, 2, 2, 128, 13, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 131, 7,
	44, 2, 2, 131, 132, 7, 35, 2, 2, 132, 133, 5, 16, 9, 2, 133, 15, 3, 2,
	2, 2, 134, 140, 7, 47, 2, 2, 135, 140, 7, 43, 2, 2, 136, 140, 7, 18, 2,
	2, 137, 140, 7, 8, 2, 2, 138, 140, 5, 10, 6, 2, 139, 134, 3, 2, 2, 2, 139,
	135, 3, 2, 2, 2, 139, 136, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 138,
	3, 2, 2, 2, 140, 17, 3, 2, 2, 2, 141, 143, 5, 20, 11, 2, 142, 141, 3, 2,
	2, 2, 143, 146, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2,
	145, 19, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 147, 148, 5, 22, 12, 2, 148,
	149, 7, 33, 2, 2, 149, 173, 3, 2, 2, 2, 150, 151, 5, 28, 15, 2, 151, 152,
	7, 33, 2, 2, 152, 173, 3, 2, 2, 2, 153, 154, 5, 32, 17, 2, 154, 155, 7,
	33, 2, 2, 155, 173, 3, 2, 2, 2, 156, 157, 5, 36, 19, 2, 157, 158, 7, 33,
	2, 2, 158, 173, 3, 2, 2, 2, 159, 160, 5, 38, 20, 2, 160, 161, 7, 33, 2,
	2, 161, 173, 3, 2, 2, 2, 162, 163, 5, 44, 23, 2, 163, 164, 7, 33, 2, 2,
	164, 173, 3, 2, 2, 2, 165, 166, 5, 50, 26, 2, 166, 167, 7, 33, 2, 2, 167,
	173, 3, 2, 2, 2, 168, 169, 5, 56, 29, 2, 169, 170, 7, 33, 2, 2, 170, 173,
	3, 2, 2, 2, 171, 173, 7, 33, 2, 2, 172, 147, 3, 2, 2, 2, 172, 150, 3, 2,
	2, 2, 172, 153, 3, 2, 2, 2, 172, 156, 3, 2, 2, 2, 172, 159, 3, 2, 2, 2,
	172, 162, 3, 2, 2, 2, 172, 165, 3, 2, 2, 2, 172, 168, 3, 2, 2, 2, 172,
	171, 3, 2, 2, 2, 173, 21, 3, 2, 2, 2, 174, 175, 5, 24, 13, 2, 175, 176,
	7, 21, 2, 2, 176, 177, 5, 26, 14, 2, 177, 23, 3, 2, 2, 2, 178, 183, 5,
	62, 32, 2, 179, 180, 7, 34, 2, 2, 180, 182, 5, 62, 32, 2, 181, 179, 3,
	2, 2, 2, 182, 185, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2,
	2, 184, 25, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 186, 191, 5, 52, 27, 2, 187,
	188, 7, 34, 2, 2, 188, 190, 5, 52, 27, 2, 189, 187, 3, 2, 2, 2, 190, 193,
	3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 27, 3, 2,
	2, 2, 193, 191, 3, 2, 2, 2, 194, 195, 7, 19, 2, 2, 195, 198, 5, 30, 16,
	2, 196, 197, 7, 21, 2, 2, 197, 199, 5, 26, 14, 2, 198, 196, 3, 2, 2, 2,
	198, 199, 3, 2, 2, 2, 199, 29, 3, 2, 2, 2, 200, 205, 7, 44, 2, 2, 201,
	202, 7, 34, 2, 2, 202, 204, 7, 44, 2, 2, 203, 201, 3, 2, 2, 2, 204, 207,
	3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 31, 3, 2,
	2, 2, 207, 205, 3, 2, 2, 2, 208, 209, 7, 10, 2, 2, 209, 210, 5, 34, 18,
	2, 210, 211, 7, 12, 2, 2, 211, 212, 5, 52, 27, 2, 212, 213, 7, 4, 2, 2,
	213, 214, 5, 18, 10, 2, 214, 215, 7, 7, 2, 2, 215, 33, 3, 2, 2, 2, 216,
	221, 7, 44, 2, 2, 217, 218, 7, 34, 2, 2, 218, 220, 7, 44, 2, 2, 219, 217,
	3, 2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221, 222, 3, 2,
	2, 2, 222, 35, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 224, 225, 7, 20, 2, 2,
	225, 226, 5, 52, 27, 2, 226, 227, 7, 4, 2, 2, 227, 228, 5, 18, 10, 2, 228,
	229, 7, 7, 2, 2, 229, 37, 3, 2, 2, 2, 230, 231, 7, 11, 2, 2, 231, 232,
	5, 52, 27, 2, 232, 233, 7, 17, 2, 2, 233, 237, 5, 18, 10, 2, 234, 236,
	5, 40, 21, 2, 235, 234, 3, 2, 2, 2, 236, 239, 3, 2, 2, 2, 237, 235, 3,
	2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 241, 3, 2, 2, 2, 239, 237, 3, 2, 2,
	2, 240, 242, 5, 42, 22, 2, 241, 240, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2,
	242, 243, 3, 2, 2, 2, 243, 244, 7, 7, 2, 2, 244, 39, 3, 2, 2, 2, 245, 246,
	7, 5, 2, 2, 246, 247, 5, 52, 27, 2, 247, 248, 7, 17, 2, 2, 248, 249, 5,
	18, 10, 2, 249, 41, 3, 2, 2, 2, 250, 251, 7, 6, 2, 2, 251, 252, 5, 18,
	10, 2, 252, 43, 3, 2, 2, 2, 253, 254, 7, 9, 2, 2, 254, 255, 7, 44, 2, 2,
	255, 257, 7, 37, 2, 2, 256, 258, 5, 46, 24, 2, 257, 256, 3, 2, 2, 2, 257,
	258, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 260, 7, 38, 2, 2, 260, 261,
	5, 18, 10, 2, 261, 262, 7, 7, 2, 2, 262, 45, 3, 2, 2, 2, 263, 268, 7, 44,
	2, 2, 264, 265, 7, 34, 2, 2, 265, 267, 7, 44, 2, 2, 266, 264, 3, 2, 2,
	2, 267, 270, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269,
	47, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 271, 272, 5, 56, 29, 2, 272, 274,
	7, 37, 2, 2, 273, 275, 5, 60, 31, 2, 274, 273, 3, 2, 2, 2, 274, 275, 3,
	2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 277, 7, 38, 2, 2, 277, 49, 3, 2, 2,
	2, 278, 280, 7, 16, 2, 2, 279, 281, 5, 26, 14, 2, 280, 279, 3, 2, 2, 2,
	280, 281, 3, 2, 2, 2, 281, 51, 3, 2, 2, 2, 282, 283, 8, 27, 1, 2, 283,
	284, 5, 54, 28, 2, 284, 302, 3, 2, 2, 2, 285, 286, 12, 7, 2, 2, 286, 287,
	9, 3, 2, 2, 287, 301, 5, 52, 27, 8, 288, 289, 12, 6, 2, 2, 289, 290, 9,
	4, 2, 2, 290, 301, 5, 52, 27, 7, 291, 292, 12, 5, 2, 2, 292, 293, 9, 5,
	2, 2, 293, 301, 5, 52, 27, 6, 294, 295, 12, 4, 2, 2, 295, 296, 7, 3, 2,
	2, 296, 301, 5, 52, 27, 5, 297, 298, 12, 3, 2, 2, 298, 299, 7, 15, 2, 2,
	299, 301, 5, 52, 27, 4, 300, 285, 3, 2, 2, 2, 300, 288, 3, 2, 2, 2, 300,
	291, 3, 2, 2, 2, 300, 294, 3, 2, 2, 2, 300, 297, 3, 2, 2, 2, 301, 304,
	3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 53, 3, 2,
	2, 2, 304, 302, 3, 2, 2, 2, 305, 306, 7, 14, 2, 2, 306, 313, 5, 54, 28,
	2, 307, 308, 7, 28, 2, 2, 308, 313, 5, 54, 28, 2, 309, 310, 7, 29, 2, 2,
	310, 313, 5, 54, 28, 2, 311, 313, 5, 56, 29, 2, 312, 305, 3, 2, 2, 2, 312,
	307, 3, 2, 2, 2, 312, 309, 3, 2, 2, 2, 312, 311, 3, 2, 2, 2, 313, 55, 3,
	2, 2, 2, 314, 315, 8, 29, 1, 2, 315, 325, 7, 44, 2, 2, 316, 325, 5, 64,
	33, 2, 317, 325, 5, 66, 34, 2, 318, 325, 5, 80, 41, 2, 319, 325, 5, 58,
	30, 2, 320, 321, 7, 37, 2, 2, 321, 322, 5, 52, 27, 2, 322, 323, 7, 38,
	2, 2, 323, 325, 3, 2, 2, 2, 324, 314, 3, 2, 2, 2, 324, 316, 3, 2, 2, 2,
	324, 317, 3, 2, 2, 2, 324, 318, 3, 2, 2, 2, 324, 319, 3, 2, 2, 2, 324,
	320, 3, 2, 2, 2, 325, 352, 3, 2, 2, 2, 326, 327, 12, 11, 2, 2, 327, 328,
	7, 36, 2, 2, 328, 351, 7, 44, 2, 2, 329, 330, 12, 10, 2, 2, 330, 331, 7,
	39, 2, 2, 331, 332, 5, 52, 27, 2, 332, 333, 7, 40, 2, 2, 333, 351, 3, 2,
	2, 2, 334, 335, 12, 9, 2, 2, 335, 337, 7, 39, 2, 2, 336, 338, 5, 52, 27,
	2, 337, 336, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339,
	341, 7, 35, 2, 2, 340, 342, 5, 52, 27, 2, 341, 340, 3, 2, 2, 2, 341, 342,
	3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 351, 7, 40, 2, 2, 344, 345, 12,
	8, 2, 2, 345, 347, 7, 37, 2, 2, 346, 348, 5, 60, 31, 2, 347, 346, 3, 2,
	2, 2, 347, 348, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 351, 7, 38, 2, 2,
	350, 326, 3, 2, 2, 2, 350, 329, 3, 2, 2, 2, 350, 334, 3, 2, 2, 2, 350,
	344, 3, 2, 2, 2, 351, 354, 3, 2, 2, 2, 352, 350, 3, 2, 2, 2, 352, 353,
	3, 2, 2, 2, 353, 57, 3, 2, 2, 2, 354, 352, 3, 2, 2, 2, 355, 356, 9, 6,
	2, 2, 356, 59, 3, 2, 2, 2, 357, 362, 5, 52, 27, 2, 358, 359, 7, 34, 2,
	2, 359, 361, 5, 52, 27, 2, 360, 358, 3, 2, 2, 2, 361, 364, 3, 2, 2, 2,
	362, 360, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 61, 3, 2, 2, 2, 364, 362,
	3, 2, 2, 2, 365, 376, 7, 44, 2, 2, 366, 367, 5, 56, 29, 2, 367, 368, 7,
	36, 2, 2, 368, 369, 7, 44, 2, 2, 369, 376, 3, 2, 2, 2, 370, 371, 5, 56,
	29, 2, 371, 372, 7, 39, 2, 2, 372, 373, 5, 52, 27, 2, 373, 374, 7, 40,
	2, 2, 374, 376, 3, 2, 2, 2, 375, 365, 3, 2, 2, 2, 375, 366, 3, 2, 2, 2,
	375, 370, 3, 2, 2, 2, 376, 63, 3, 2, 2, 2, 377, 378, 7, 9, 2, 2, 378, 380,
	7, 37, 2, 2, 379, 381, 5, 46, 24, 2, 380, 379, 3, 2, 2, 2, 380, 381, 3,
	2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 383, 7, 38, 2, 2, 383, 384, 5, 18,
	10, 2, 384, 385, 7, 7, 2, 2, 385, 65, 3, 2, 2, 2, 386, 388, 7, 41, 2, 2,
	387, 389, 5, 68, 35, 2, 388, 387, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 389,
	390, 3, 2, 2, 2, 390, 391, 7, 42, 2, 2, 391, 67, 3, 2, 2, 2, 392, 397,
	5, 70, 36, 2, 393, 394, 9, 2, 2, 2, 394, 396, 5, 70, 36, 2, 395, 393, 3,
	2, 2, 2, 396, 399, 3, 2, 2, 2, 397, 395, 3, 2, 2, 2, 397, 398, 3, 2, 2,
	2, 398, 403, 3, 2, 2, 2, 399, 397, 3, 2, 2, 2, 400, 402, 9, 2, 2, 2, 401,
	400, 3, 2, 2, 2, 402, 405, 3, 2, 2, 2, 403, 401, 3, 2, 2, 2, 403, 404,
	3, 2, 2, 2, 404, 69, 3, 2, 2, 2, 405, 403, 3, 2, 2, 2, 406, 407, 5, 94,
	48, 2, 407, 408, 7, 35, 2, 2, 408, 409, 5, 52, 27, 2, 409, 419, 3, 2, 2,
	2, 410, 411, 7, 39, 2, 2, 411, 412, 5, 52, 27, 2, 412, 413, 7, 40, 2, 2,
	413, 414, 7, 35, 2, 2, 414, 415, 5, 52, 27, 2, 415, 419, 3, 2, 2, 2, 416,
	419, 5, 72, 37, 2, 417, 419, 5, 78, 40, 2, 418, 406, 3, 2, 2, 2, 418, 410,
	3, 2, 2, 2, 418, 416, 3, 2, 2, 2, 418, 417, 3, 2, 2, 2, 419, 71, 3, 2,
	2, 2, 420, 421, 7, 11, 2, 2, 421, 422, 5, 52, 27, 2, 422, 424, 7, 17, 2,
	2, 423, 425, 5, 68, 35, 2, 424, 423, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2,
	425, 429, 3, 2, 2, 2, 426, 428, 5, 74, 38, 2, 427, 426, 3, 2, 2, 2, 428,
	431, 3, 2, 2, 2, 429, 427, 3, 2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 433,
	3, 2, 2, 2, 431, 429, 3, 2, 2, 2, 432, 434, 5, 76, 39, 2, 433, 432, 3,
	2, 2, 2, 433, 434, 3, 2, 2, 2, 434, 435, 3, 2, 2, 2, 435, 436, 7, 7, 2,
	2, 436, 73, 3, 2, 2, 2, 437, 438, 7, 5, 2, 2, 438, 439, 5, 52, 27, 2, 439,
	441, 7, 17, 2, 2, 440, 442, 5, 68, 35, 2, 441, 440, 3, 2, 2, 2, 441, 442,
	3, 2, 2, 2, 442, 75, 3, 2, 2, 2, 443, 445, 7, 6, 2, 2, 444, 446, 5, 68,
	35, 2, 445, 444, 3, 2, 2, 2, 445, 446, 3, 2, 2, 2, 446, 77, 3, 2, 2, 2,
	447, 448, 7, 10, 2, 2, 448, 449, 5, 34, 18, 2, 449, 450, 7, 12, 2, 2, 450,
	451, 5, 52, 27, 2, 451, 453, 7, 4, 2, 2, 452, 454, 5, 68, 35, 2, 453, 452,
	3, 2, 2, 2, 453, 454, 3, 2, 2, 2, 454, 455, 3, 2, 2, 2, 455, 456, 7, 7,
	2, 2, 456, 79, 3, 2, 2, 2, 457, 459, 7, 39, 2, 2, 458, 460, 5, 82, 42,
	2, 459, 458, 3, 2, 2, 2, 459, 460, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461,
	462, 7, 40, 2, 2, 462, 81, 3, 2, 2, 2, 463, 468, 5, 84, 43, 2, 464, 465,
	9, 2, 2, 2, 465, 467, 5, 84, 43, 2, 466, 464, 3, 2, 2, 2, 467, 470, 3,
	2, 2, 2, 468, 466, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 474, 3, 2, 2,
	2, 470, 468, 3, 2, 2, 2, 471, 473, 9, 2, 2, 2, 472, 471, 3, 2, 2, 2, 473,
	476, 3, 2, 2, 2, 474, 472, 3, 2, 2, 2, 474, 475, 3, 2, 2, 2, 475, 83, 3,
	2, 2, 2, 476, 474, 3, 2, 2, 2, 477, 481, 5, 52, 27, 2, 478, 481, 5, 86,
	44, 2, 479, 481, 5, 92, 47, 2, 480, 477, 3, 2, 2, 2, 480, 478, 3, 2, 2,
	2, 480, 479, 3, 2, 2, 2, 481, 85, 3, 2, 2, 2, 482, 483, 7, 11, 2, 2, 483,
	484, 5, 52, 27, 2, 484, 486, 7, 17, 2, 2, 485, 487, 5, 82, 42, 2, 486,
	485, 3, 2, 2, 2, 486, 487, 3, 2, 2, 2, 487, 491, 3, 2, 2, 2, 488, 490,
	5, 88, 45, 2, 489, 488, 3, 2, 2, 2, 490, 493, 3, 2, 2, 2, 491, 489, 3,
	2, 2, 2, 491, 492, 3, 2, 2, 2, 492, 495, 3, 2, 2, 2, 493, 491, 3, 2, 2,
	2, 494, 496, 5, 90, 46, 2, 495, 494, 3, 2, 2, 2, 495, 496, 3, 2, 2, 2,
	496, 497, 3, 2, 2, 2, 497, 498, 7, 7, 2, 2, 498, 87, 3, 2, 2, 2, 499, 500,
	7, 5, 2, 2, 500, 501, 5, 52, 27, 2, 501, 503, 7, 17, 2, 2, 502, 504, 5,
	82, 42, 2, 503, 502, 3, 2, 2, 2, 503, 504, 3, 2, 2, 2, 504, 89, 3, 2, 2,
	2, 505, 507, 7, 6, 2, 2, 506, 508, 5, 82, 42, 2, 507, 506, 3, 2, 2, 2,
	507, 508, 3, 2, 2, 2, 508, 91, 3, 2, 2, 2, 509, 510, 7, 10, 2, 2, 510,
	511, 5, 34, 18, 2, 511, 512, 7, 12, 2, 2, 512, 513, 5, 52, 27, 2, 513,
	515, 7, 4, 2, 2, 514, 516, 5, 82, 42, 2, 515, 514, 3, 2, 2, 2, 515, 516,
	3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517, 518, 7, 7, 2, 2, 518, 93, 3, 2,
	2, 2, 519, 520, 9, 7, 2, 2, 520, 95, 3, 2, 2, 2, 52, 100, 108, 118, 127,
	139, 144, 172, 183, 191, 198, 205, 221, 237, 241, 257, 268, 274, 280, 300,
	302, 312, 324, 337, 341, 347, 350, 352, 362, 375, 380, 388, 397, 403, 418,
	424, 429, 433, 441, 445, 453, 459, 468, 474, 480, 486, 491, 495, 503, 507,
	515,
}
var literalNames = []string{
	"", "'and'", "'do'", "'elif'", "'else'", "'end'", "'false'", "'fn'", "'for'",
	"'if'", "'in'", "'meta'", "'not'", "'or'", "'return'", "'then'", "'true'",
	"'var'", "'while'", "'='", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "';'", "','", "':'", "'.'", "'('", "')'",
	"'['", "']'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "AND", "DO", "ELIF", "ELSE", "END", "FALSE", "FN", "FOR", "IF", "IN",
	"META", "NOT", "OR", "RETURN", "THEN", "TRUE", "VAR", "WHILE", "ASSIGN",
	"EQ", "NE", "LT", "LE", "GT", "GE", "ADD", "SUB", "MUL", "DIV", "MOD",
	"SEMICOLON", "COMMA", "COLON", "PERIOD", "OPAREN", "CPAREN", "OBRACKET",
	"CBRACKET", "OCURLY", "CCURLY", "NUMBER", "ID", "WS", "NEWLINE", "STRING",
	"LQUOTE",
}

var ruleNames = []string{
	"start", "module", "meta_section", "meta_entry", "meta_object", "meta_fields",
	"meta_field", "meta_field_value", "stmts", "stmt", "assignment_stmt", "assignment_lvalues",
	"rvalues", "var_decl_stmt", "var_decl_vars", "for_stmt", "for_vars", "while_stmt",
	"if_stmt", "if_elif", "if_else", "func_stmt", "param_list", "func_call_stmt",
	"return_stmt", "expr", "unary_expr", "primary_expr", "simple_literal",
	"arg_list", "lvalue_expr", "lambda_expr", "object_literal", "object_fields",
	"object_field", "object_if", "object_elif", "object_else", "object_for",
	"array_literal", "array_elems", "array_elem", "array_if", "array_elif",
	"array_else", "array_for", "id_or_keyword",
}

type NitroParser struct {
	*antlr.BaseParser
}

// NewNitroParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *NitroParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNitroParser(input antlr.TokenStream) *NitroParser {
	this := new(NitroParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "NitroParser.g4"

	return this
}

// NitroParser tokens.
const (
	NitroParserEOF       = antlr.TokenEOF
	NitroParserAND       = 1
	NitroParserDO        = 2
	NitroParserELIF      = 3
	NitroParserELSE      = 4
	NitroParserEND       = 5
	NitroParserFALSE     = 6
	NitroParserFN        = 7
	NitroParserFOR       = 8
	NitroParserIF        = 9
	NitroParserIN        = 10
	NitroParserMETA      = 11
	NitroParserNOT       = 12
	NitroParserOR        = 13
	NitroParserRETURN    = 14
	NitroParserTHEN      = 15
	NitroParserTRUE      = 16
	NitroParserVAR       = 17
	NitroParserWHILE     = 18
	NitroParserASSIGN    = 19
	NitroParserEQ        = 20
	NitroParserNE        = 21
	NitroParserLT        = 22
	NitroParserLE        = 23
	NitroParserGT        = 24
	NitroParserGE        = 25
	NitroParserADD       = 26
	NitroParserSUB       = 27
	NitroParserMUL       = 28
	NitroParserDIV       = 29
	NitroParserMOD       = 30
	NitroParserSEMICOLON = 31
	NitroParserCOMMA     = 32
	NitroParserCOLON     = 33
	NitroParserPERIOD    = 34
	NitroParserOPAREN    = 35
	NitroParserCPAREN    = 36
	NitroParserOBRACKET  = 37
	NitroParserCBRACKET  = 38
	NitroParserOCURLY    = 39
	NitroParserCCURLY    = 40
	NitroParserNUMBER    = 41
	NitroParserID        = 42
	NitroParserWS        = 43
	NitroParserNEWLINE   = 44
	NitroParserSTRING    = 45
	NitroParserLQUOTE    = 46
)

// NitroParser rules.
const (
	NitroParserRULE_start              = 0
	NitroParserRULE_module             = 1
	NitroParserRULE_meta_section       = 2
	NitroParserRULE_meta_entry         = 3
	NitroParserRULE_meta_object        = 4
	NitroParserRULE_meta_fields        = 5
	NitroParserRULE_meta_field         = 6
	NitroParserRULE_meta_field_value   = 7
	NitroParserRULE_stmts              = 8
	NitroParserRULE_stmt               = 9
	NitroParserRULE_assignment_stmt    = 10
	NitroParserRULE_assignment_lvalues = 11
	NitroParserRULE_rvalues            = 12
	NitroParserRULE_var_decl_stmt      = 13
	NitroParserRULE_var_decl_vars      = 14
	NitroParserRULE_for_stmt           = 15
	NitroParserRULE_for_vars           = 16
	NitroParserRULE_while_stmt         = 17
	NitroParserRULE_if_stmt            = 18
	NitroParserRULE_if_elif            = 19
	NitroParserRULE_if_else            = 20
	NitroParserRULE_func_stmt          = 21
	NitroParserRULE_param_list         = 22
	NitroParserRULE_func_call_stmt     = 23
	NitroParserRULE_return_stmt        = 24
	NitroParserRULE_expr               = 25
	NitroParserRULE_unary_expr         = 26
	NitroParserRULE_primary_expr       = 27
	NitroParserRULE_simple_literal     = 28
	NitroParserRULE_arg_list           = 29
	NitroParserRULE_lvalue_expr        = 30
	NitroParserRULE_lambda_expr        = 31
	NitroParserRULE_object_literal     = 32
	NitroParserRULE_object_fields      = 33
	NitroParserRULE_object_field       = 34
	NitroParserRULE_object_if          = 35
	NitroParserRULE_object_elif        = 36
	NitroParserRULE_object_else        = 37
	NitroParserRULE_object_for         = 38
	NitroParserRULE_array_literal      = 39
	NitroParserRULE_array_elems        = 40
	NitroParserRULE_array_elem         = 41
	NitroParserRULE_array_if           = 42
	NitroParserRULE_array_elif         = 43
	NitroParserRULE_array_else         = 44
	NitroParserRULE_array_for          = 45
	NitroParserRULE_id_or_keyword      = 46
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

func (s *StartContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
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
		p.SetState(94)
		p.Module()
	}
	{
		p.SetState(95)
		p.Match(NitroParserEOF)
	}

	return localctx
}

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *ModuleContext) Meta_section() IMeta_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeta_sectionContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitModule(s)
	}
}

func (p *NitroParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NitroParserRULE_module)
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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserMETA {
		{
			p.SetState(97)
			p.Meta_section()
		}

	}
	{
		p.SetState(100)
		p.Stmts()
	}

	return localctx
}

// IMeta_sectionContext is an interface to support dynamic dispatch.
type IMeta_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_sectionContext differentiates from other interfaces.
	IsMeta_sectionContext()
}

type Meta_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_sectionContext() *Meta_sectionContext {
	var p = new(Meta_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_section
	return p
}

func (*Meta_sectionContext) IsMeta_sectionContext() {}

func NewMeta_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_sectionContext {
	var p = new(Meta_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_section

	return p
}

func (s *Meta_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_sectionContext) META() antlr.TerminalNode {
	return s.GetToken(NitroParserMETA, 0)
}

func (s *Meta_sectionContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
}

func (s *Meta_sectionContext) AllMeta_entry() []IMeta_entryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMeta_entryContext)(nil)).Elem())
	var tst = make([]IMeta_entryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMeta_entryContext)
		}
	}

	return tst
}

func (s *Meta_sectionContext) Meta_entry(i int) IMeta_entryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_entryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMeta_entryContext)
}

func (s *Meta_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_section(s)
	}
}

func (s *Meta_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_section(s)
	}
}

func (p *NitroParser) Meta_section() (localctx IMeta_sectionContext) {
	localctx = NewMeta_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NitroParserRULE_meta_section)
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
		p.SetState(102)
		p.Match(NitroParserMETA)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserID {
		{
			p.SetState(103)
			p.Meta_entry()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
		p.Match(NitroParserEND)
	}

	return localctx
}

// IMeta_entryContext is an interface to support dynamic dispatch.
type IMeta_entryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_entryContext differentiates from other interfaces.
	IsMeta_entryContext()
}

type Meta_entryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_entryContext() *Meta_entryContext {
	var p = new(Meta_entryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_entry
	return p
}

func (*Meta_entryContext) IsMeta_entryContext() {}

func NewMeta_entryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_entryContext {
	var p = new(Meta_entryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_entry

	return p
}

func (s *Meta_entryContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_entryContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Meta_entryContext) Meta_object() IMeta_objectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_objectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeta_objectContext)
}

func (s *Meta_entryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_entryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_entryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_entry(s)
	}
}

func (s *Meta_entryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_entry(s)
	}
}

func (p *NitroParser) Meta_entry() (localctx IMeta_entryContext) {
	localctx = NewMeta_entryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NitroParserRULE_meta_entry)

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
		p.SetState(111)
		p.Match(NitroParserID)
	}
	{
		p.SetState(112)
		p.Meta_object()
	}

	return localctx
}

// IMeta_objectContext is an interface to support dynamic dispatch.
type IMeta_objectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_objectContext differentiates from other interfaces.
	IsMeta_objectContext()
}

type Meta_objectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_objectContext() *Meta_objectContext {
	var p = new(Meta_objectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_object
	return p
}

func (*Meta_objectContext) IsMeta_objectContext() {}

func NewMeta_objectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_objectContext {
	var p = new(Meta_objectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_object

	return p
}

func (s *Meta_objectContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_objectContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Meta_objectContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *Meta_objectContext) Meta_fields() IMeta_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeta_fieldsContext)
}

func (s *Meta_objectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_objectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_objectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_object(s)
	}
}

func (s *Meta_objectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_object(s)
	}
}

func (p *NitroParser) Meta_object() (localctx IMeta_objectContext) {
	localctx = NewMeta_objectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NitroParserRULE_meta_object)
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
		p.SetState(114)
		p.Match(NitroParserOCURLY)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(115)
			p.Meta_fields()
		}

	}
	{
		p.SetState(118)
		p.Match(NitroParserCCURLY)
	}

	return localctx
}

// IMeta_fieldsContext is an interface to support dynamic dispatch.
type IMeta_fieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_fieldsContext differentiates from other interfaces.
	IsMeta_fieldsContext()
}

type Meta_fieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_fieldsContext() *Meta_fieldsContext {
	var p = new(Meta_fieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_fields
	return p
}

func (*Meta_fieldsContext) IsMeta_fieldsContext() {}

func NewMeta_fieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_fieldsContext {
	var p = new(Meta_fieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_fields

	return p
}

func (s *Meta_fieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_fieldsContext) AllMeta_field() []IMeta_fieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMeta_fieldContext)(nil)).Elem())
	var tst = make([]IMeta_fieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMeta_fieldContext)
		}
	}

	return tst
}

func (s *Meta_fieldsContext) Meta_field(i int) IMeta_fieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_fieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMeta_fieldContext)
}

func (s *Meta_fieldsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NitroParserCOMMA)
}

func (s *Meta_fieldsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserCOMMA, i)
}

func (s *Meta_fieldsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(NitroParserSEMICOLON)
}

func (s *Meta_fieldsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, i)
}

func (s *Meta_fieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_fieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_fieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_fields(s)
	}
}

func (s *Meta_fieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_fields(s)
	}
}

func (p *NitroParser) Meta_fields() (localctx IMeta_fieldsContext) {
	localctx = NewMeta_fieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NitroParserRULE_meta_fields)
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
		p.SetState(120)
		p.Meta_field()
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
		{
			p.SetState(121)
			_la = p.GetTokenStream().LA(1)

			if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(122)
			p.Meta_field()
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMeta_fieldContext is an interface to support dynamic dispatch.
type IMeta_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_fieldContext differentiates from other interfaces.
	IsMeta_fieldContext()
}

type Meta_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_fieldContext() *Meta_fieldContext {
	var p = new(Meta_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_field
	return p
}

func (*Meta_fieldContext) IsMeta_fieldContext() {}

func NewMeta_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_fieldContext {
	var p = new(Meta_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_field

	return p
}

func (s *Meta_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_fieldContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Meta_fieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(NitroParserCOLON, 0)
}

func (s *Meta_fieldContext) Meta_field_value() IMeta_field_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_field_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeta_field_valueContext)
}

func (s *Meta_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_field(s)
	}
}

func (s *Meta_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_field(s)
	}
}

func (p *NitroParser) Meta_field() (localctx IMeta_fieldContext) {
	localctx = NewMeta_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NitroParserRULE_meta_field)

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
		p.SetState(128)
		p.Match(NitroParserID)
	}
	{
		p.SetState(129)
		p.Match(NitroParserCOLON)
	}
	{
		p.SetState(130)
		p.Meta_field_value()
	}

	return localctx
}

// IMeta_field_valueContext is an interface to support dynamic dispatch.
type IMeta_field_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_field_valueContext differentiates from other interfaces.
	IsMeta_field_valueContext()
}

type Meta_field_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_field_valueContext() *Meta_field_valueContext {
	var p = new(Meta_field_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_meta_field_value
	return p
}

func (*Meta_field_valueContext) IsMeta_field_valueContext() {}

func NewMeta_field_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_field_valueContext {
	var p = new(Meta_field_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_meta_field_value

	return p
}

func (s *Meta_field_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_field_valueContext) STRING() antlr.TerminalNode {
	return s.GetToken(NitroParserSTRING, 0)
}

func (s *Meta_field_valueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NitroParserNUMBER, 0)
}

func (s *Meta_field_valueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NitroParserTRUE, 0)
}

func (s *Meta_field_valueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NitroParserFALSE, 0)
}

func (s *Meta_field_valueContext) Meta_object() IMeta_objectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_objectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeta_objectContext)
}

func (s *Meta_field_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_field_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_field_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterMeta_field_value(s)
	}
}

func (s *Meta_field_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitMeta_field_value(s)
	}
}

func (p *NitroParser) Meta_field_value() (localctx IMeta_field_valueContext) {
	localctx = NewMeta_field_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NitroParserRULE_meta_field_value)

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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(NitroParserSTRING)
		}

	case NitroParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(NitroParserNUMBER)
		}

	case NitroParserTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Match(NitroParserTRUE)
		}

	case NitroParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(135)
			p.Match(NitroParserFALSE)
		}

	case NitroParserOCURLY:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(136)
			p.Meta_object()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *StmtsContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *StmtsContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
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
	localctx = NewStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NitroParserRULE_stmts)
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
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserRETURN)|(1<<NitroParserTRUE)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserSEMICOLON))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(139)
			p.Stmt()
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *StmtContext) Assignment_stmt() IAssignment_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignment_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignment_stmtContext)
}

func (s *StmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *StmtContext) Var_decl_stmt() IVar_decl_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_decl_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_decl_stmtContext)
}

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) While_stmt() IWhile_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) Func_stmt() IFunc_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_stmtContext)
}

func (s *StmtContext) Return_stmt() IReturn_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_stmtContext)
}

func (s *StmtContext) Primary_expr() IPrimary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *NitroParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NitroParserRULE_stmt)

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

	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Assignment_stmt()
		}
		{
			p.SetState(146)
			p.Match(NitroParserSEMICOLON)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Var_decl_stmt()
		}
		{
			p.SetState(149)
			p.Match(NitroParserSEMICOLON)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.For_stmt()
		}
		{
			p.SetState(152)
			p.Match(NitroParserSEMICOLON)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(154)
			p.While_stmt()
		}
		{
			p.SetState(155)
			p.Match(NitroParserSEMICOLON)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(157)
			p.If_stmt()
		}
		{
			p.SetState(158)
			p.Match(NitroParserSEMICOLON)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(160)
			p.Func_stmt()
		}
		{
			p.SetState(161)
			p.Match(NitroParserSEMICOLON)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(163)
			p.Return_stmt()
		}
		{
			p.SetState(164)
			p.Match(NitroParserSEMICOLON)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(166)
			p.primary_expr(0)
		}
		{
			p.SetState(167)
			p.Match(NitroParserSEMICOLON)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(169)
			p.Match(NitroParserSEMICOLON)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignment_lvaluesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignment_lvaluesContext)
}

func (s *Assignment_stmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN, 0)
}

func (s *Assignment_stmtContext) Rvalues() IRvaluesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvaluesContext)(nil)).Elem(), 0)

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
	localctx = NewAssignment_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NitroParserRULE_assignment_stmt)

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
		p.SetState(172)
		p.Assignment_lvalues()
	}
	{
		p.SetState(173)
		p.Match(NitroParserASSIGN)
	}
	{
		p.SetState(174)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILvalue_exprContext)(nil)).Elem())
	var tst = make([]ILvalue_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILvalue_exprContext)
		}
	}

	return tst
}

func (s *Assignment_lvaluesContext) Lvalue_expr(i int) ILvalue_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILvalue_exprContext)(nil)).Elem(), i)

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
	localctx = NewAssignment_lvaluesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NitroParserRULE_assignment_lvalues)
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
		p.SetState(176)
		p.Lvalue_expr()
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(177)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(178)
			p.Lvalue_expr()
		}

		p.SetState(183)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RvaluesContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

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
	localctx = NewRvaluesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NitroParserRULE_rvalues)
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
		p.SetState(184)
		p.expr(0)
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(185)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(186)
			p.expr(0)
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_decl_varsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_decl_varsContext)
}

func (s *Var_decl_stmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN, 0)
}

func (s *Var_decl_stmtContext) Rvalues() IRvaluesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvaluesContext)(nil)).Elem(), 0)

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
	localctx = NewVar_decl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NitroParserRULE_var_decl_stmt)
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
		p.SetState(192)
		p.Match(NitroParserVAR)
	}
	{
		p.SetState(193)
		p.Var_decl_vars()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserASSIGN {
		{
			p.SetState(194)
			p.Match(NitroParserASSIGN)
		}
		{
			p.SetState(195)
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
	localctx = NewVar_decl_varsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NitroParserRULE_var_decl_vars)
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
		p.SetState(198)
		p.Match(NitroParserID)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(199)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(200)
			p.Match(NitroParserID)
		}

		p.SetState(205)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_varsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_varsContext)
}

func (s *For_stmtContext) IN() antlr.TerminalNode {
	return s.GetToken(NitroParserIN, 0)
}

func (s *For_stmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *For_stmtContext) DO() antlr.TerminalNode {
	return s.GetToken(NitroParserDO, 0)
}

func (s *For_stmtContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *For_stmtContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
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
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NitroParserRULE_for_stmt)

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
		p.SetState(206)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(207)
		p.For_vars()
	}
	{
		p.SetState(208)
		p.Match(NitroParserIN)
	}
	{
		p.SetState(209)
		p.expr(0)
	}
	{
		p.SetState(210)
		p.Match(NitroParserDO)
	}
	{
		p.SetState(211)
		p.Stmts()
	}
	{
		p.SetState(212)
		p.Match(NitroParserEND)
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
	localctx = NewFor_varsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NitroParserRULE_for_vars)
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
		p.SetState(214)
		p.Match(NitroParserID)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(215)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(216)
			p.Match(NitroParserID)
		}

		p.SetState(221)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *While_stmtContext) DO() antlr.TerminalNode {
	return s.GetToken(NitroParserDO, 0)
}

func (s *While_stmtContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *While_stmtContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
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
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NitroParserRULE_while_stmt)

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
		p.SetState(222)
		p.Match(NitroParserWHILE)
	}
	{
		p.SetState(223)
		p.expr(0)
	}
	{
		p.SetState(224)
		p.Match(NitroParserDO)
	}
	{
		p.SetState(225)
		p.Stmts()
	}
	{
		p.SetState(226)
		p.Match(NitroParserEND)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_stmtContext) THEN() antlr.TerminalNode {
	return s.GetToken(NitroParserTHEN, 0)
}

func (s *If_stmtContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *If_stmtContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
}

func (s *If_stmtContext) AllIf_elif() []IIf_elifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIf_elifContext)(nil)).Elem())
	var tst = make([]IIf_elifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIf_elifContext)
		}
	}

	return tst
}

func (s *If_stmtContext) If_elif(i int) IIf_elifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_elifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIf_elifContext)
}

func (s *If_stmtContext) If_else() IIf_elseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_elseContext)(nil)).Elem(), 0)

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
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, NitroParserRULE_if_stmt)
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
		p.SetState(228)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(229)
		p.expr(0)
	}
	{
		p.SetState(230)
		p.Match(NitroParserTHEN)
	}
	{
		p.SetState(231)
		p.Stmts()
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserELIF {
		{
			p.SetState(232)
			p.If_elif()
		}

		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(238)
			p.If_else()
		}

	}
	{
		p.SetState(241)
		p.Match(NitroParserEND)
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

func (s *If_elifContext) ELIF() antlr.TerminalNode {
	return s.GetToken(NitroParserELIF, 0)
}

func (s *If_elifContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_elifContext) THEN() antlr.TerminalNode {
	return s.GetToken(NitroParserTHEN, 0)
}

func (s *If_elifContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
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
	localctx = NewIf_elifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, NitroParserRULE_if_elif)

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
		p.SetState(243)
		p.Match(NitroParserELIF)
	}
	{
		p.SetState(244)
		p.expr(0)
	}
	{
		p.SetState(245)
		p.Match(NitroParserTHEN)
	}
	{
		p.SetState(246)
		p.Stmts()
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

func (s *If_elseContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
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
	localctx = NewIf_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, NitroParserRULE_if_else)

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
		p.SetState(248)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(249)
		p.Stmts()
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

func (s *Func_stmtContext) FN() antlr.TerminalNode {
	return s.GetToken(NitroParserFN, 0)
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

func (s *Func_stmtContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *Func_stmtContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
}

func (s *Func_stmtContext) Param_list() IParam_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParam_listContext)(nil)).Elem(), 0)

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
	localctx = NewFunc_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, NitroParserRULE_func_stmt)
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
		p.SetState(251)
		p.Match(NitroParserFN)
	}
	{
		p.SetState(252)
		p.Match(NitroParserID)
	}
	{
		p.SetState(253)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(254)
			p.Param_list()
		}

	}
	{
		p.SetState(257)
		p.Match(NitroParserCPAREN)
	}
	{
		p.SetState(258)
		p.Stmts()
	}
	{
		p.SetState(259)
		p.Match(NitroParserEND)
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
	localctx = NewParam_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, NitroParserRULE_param_list)
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
		p.SetState(261)
		p.Match(NitroParserID)
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(262)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(263)
			p.Match(NitroParserID)
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunc_call_stmtContext is an interface to support dynamic dispatch.
type IFunc_call_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_call_stmtContext differentiates from other interfaces.
	IsFunc_call_stmtContext()
}

type Func_call_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_call_stmtContext() *Func_call_stmtContext {
	var p = new(Func_call_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_func_call_stmt
	return p
}

func (*Func_call_stmtContext) IsFunc_call_stmtContext() {}

func NewFunc_call_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_call_stmtContext {
	var p = new(Func_call_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_func_call_stmt

	return p
}

func (s *Func_call_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_call_stmtContext) Primary_expr() IPrimary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Func_call_stmtContext) OPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserOPAREN, 0)
}

func (s *Func_call_stmtContext) CPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserCPAREN, 0)
}

func (s *Func_call_stmtContext) Arg_list() IArg_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArg_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *Func_call_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_call_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_call_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterFunc_call_stmt(s)
	}
}

func (s *Func_call_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitFunc_call_stmt(s)
	}
}

func (p *NitroParser) Func_call_stmt() (localctx IFunc_call_stmtContext) {
	localctx = NewFunc_call_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, NitroParserRULE_func_call_stmt)
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
		p.SetState(269)
		p.primary_expr(0)
	}
	{
		p.SetState(270)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(271)
			p.Arg_list()
		}

	}
	{
		p.SetState(274)
		p.Match(NitroParserCPAREN)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvaluesContext)(nil)).Elem(), 0)

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
	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, NitroParserRULE_return_stmt)
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
		p.SetState(276)
		p.Match(NitroParserRETURN)
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(277)
			p.Rvalues()
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
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

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Unary_expr() IUnary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnary_exprContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(NitroParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(NitroParserDIV, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(NitroParserMOD, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(NitroParserADD, 0)
}

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(NitroParserSUB, 0)
}

func (s *ExprContext) LT() antlr.TerminalNode {
	return s.GetToken(NitroParserLT, 0)
}

func (s *ExprContext) LE() antlr.TerminalNode {
	return s.GetToken(NitroParserLE, 0)
}

func (s *ExprContext) GT() antlr.TerminalNode {
	return s.GetToken(NitroParserGT, 0)
}

func (s *ExprContext) GE() antlr.TerminalNode {
	return s.GetToken(NitroParserGE, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(NitroParserEQ, 0)
}

func (s *ExprContext) NE() antlr.TerminalNode {
	return s.GetToken(NitroParserNE, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(NitroParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(NitroParserOR, 0)
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
	return p.expr(0)
}

func (p *NitroParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, NitroParserRULE_expr, _p)
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
		p.SetState(281)
		p.Unary_expr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(298)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_expr)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(284)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserMUL)|(1<<NitroParserDIV)|(1<<NitroParserMOD))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(285)
					p.expr(6)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_expr)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(287)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == NitroParserADD || _la == NitroParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(288)
					p.expr(5)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_expr)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(290)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserEQ)|(1<<NitroParserNE)|(1<<NitroParserLT)|(1<<NitroParserLE)|(1<<NitroParserGT)|(1<<NitroParserGE))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(291)
					p.expr(4)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_expr)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(293)

					var _m = p.Match(NitroParserAND)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(294)
					p.expr(3)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_expr)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(296)

					var _m = p.Match(NitroParserOR)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(297)
					p.expr(2)
				}

			}

		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
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

func (s *Unary_exprContext) Unary_expr() IUnary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnary_exprContext)
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

func (s *Unary_exprContext) Primary_expr() IPrimary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
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
	localctx = NewUnary_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, NitroParserRULE_unary_expr)

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

	p.SetState(310)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(303)

			var _m = p.Match(NitroParserNOT)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(304)
			p.Unary_expr()
		}

	case NitroParserADD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)

			var _m = p.Match(NitroParserADD)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(306)
			p.Unary_expr()
		}

	case NitroParserSUB:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(307)

			var _m = p.Match(NitroParserSUB)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(308)
			p.Unary_expr()
		}

	case NitroParserFALSE, NitroParserFN, NitroParserTRUE, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserNUMBER, NitroParserID, NitroParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(309)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_literalContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambda_exprContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Primary_expr_indexContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserOBRACKET, 0)
}

func (s *Primary_expr_indexContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_literalContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_literalContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArg_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Primary_expr_sliceContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

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
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimary_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimary_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, NitroParserRULE_primary_expr, _p)
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
	p.SetState(322)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserID:
		localctx = NewPrimary_expr_simple_refContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(313)
			p.Match(NitroParserID)
		}

	case NitroParserFN:
		localctx = NewPrimary_expr_lambdaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(314)
			p.Lambda_expr()
		}

	case NitroParserOCURLY:
		localctx = NewPrimary_expr_objectContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(315)
			p.Object_literal()
		}

	case NitroParserOBRACKET:
		localctx = NewPrimary_expr_arrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(316)
			p.Array_literal()
		}

	case NitroParserFALSE, NitroParserTRUE, NitroParserNUMBER, NitroParserSTRING:
		localctx = NewPrimary_expr_literalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(317)
			p.Simple_literal()
		}

	case NitroParserOPAREN:
		localctx = NewPrimary_expr_parenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(318)
			p.Match(NitroParserOPAREN)
		}
		{
			p.SetState(319)
			p.expr(0)
		}
		{
			p.SetState(320)
			p.Match(NitroParserCPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(348)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimary_expr_member_accessContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(324)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(325)
					p.Match(NitroParserPERIOD)
				}
				{
					p.SetState(326)
					p.Match(NitroParserID)
				}

			case 2:
				localctx = NewPrimary_expr_indexContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(327)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(328)
					p.Match(NitroParserOBRACKET)
				}
				{
					p.SetState(329)
					p.expr(0)
				}
				{
					p.SetState(330)
					p.Match(NitroParserCBRACKET)
				}

			case 3:
				localctx = NewPrimary_expr_sliceContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(332)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(333)
					p.Match(NitroParserOBRACKET)
				}
				p.SetState(335)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
					{
						p.SetState(334)

						var _x = p.expr(0)

						localctx.(*Primary_expr_sliceContext).b = _x
					}

				}
				{
					p.SetState(337)
					p.Match(NitroParserCOLON)
				}
				p.SetState(339)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
					{
						p.SetState(338)

						var _x = p.expr(0)

						localctx.(*Primary_expr_sliceContext).e = _x
					}

				}
				{
					p.SetState(341)
					p.Match(NitroParserCBRACKET)
				}

			case 4:
				localctx = NewPrimary_expr_callContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(342)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(343)
					p.Match(NitroParserOPAREN)
				}
				p.SetState(345)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
					{
						p.SetState(344)
						p.Arg_list()
					}

				}
				{
					p.SetState(347)
					p.Match(NitroParserCPAREN)
				}

			}

		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
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

func (s *Simple_literalContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NitroParserTRUE, 0)
}

func (s *Simple_literalContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NitroParserFALSE, 0)
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
	localctx = NewSimple_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, NitroParserRULE_simple_literal)
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
		p.SetState(353)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Simple_literalContext).val = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == NitroParserFALSE || _la == NitroParserTRUE || _la == NitroParserNUMBER || _la == NitroParserSTRING) {
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Arg_listContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

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
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, NitroParserRULE_arg_list)
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
		p.SetState(355)
		p.expr(0)
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(356)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(357)
			p.expr(0)
		}

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Lvalue_expr_indexContext) OBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserOBRACKET, 0)
}

func (s *Lvalue_expr_indexContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

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
	localctx = NewLvalue_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, NitroParserRULE_lvalue_expr)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLvalue_expr_simple_refContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(363)
			p.Match(NitroParserID)
		}

	case 2:
		localctx = NewLvalue_expr_member_accessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(364)
			p.primary_expr(0)
		}
		{
			p.SetState(365)
			p.Match(NitroParserPERIOD)
		}
		{
			p.SetState(366)
			p.Match(NitroParserID)
		}

	case 3:
		localctx = NewLvalue_expr_indexContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(368)
			p.primary_expr(0)
		}
		{
			p.SetState(369)
			p.Match(NitroParserOBRACKET)
		}
		{
			p.SetState(370)
			p.expr(0)
		}
		{
			p.SetState(371)
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

func (s *Lambda_exprContext) FN() antlr.TerminalNode {
	return s.GetToken(NitroParserFN, 0)
}

func (s *Lambda_exprContext) OPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserOPAREN, 0)
}

func (s *Lambda_exprContext) CPAREN() antlr.TerminalNode {
	return s.GetToken(NitroParserCPAREN, 0)
}

func (s *Lambda_exprContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *Lambda_exprContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
}

func (s *Lambda_exprContext) Param_list() IParam_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParam_listContext)(nil)).Elem(), 0)

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
	localctx = NewLambda_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, NitroParserRULE_lambda_expr)
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
		p.SetState(375)
		p.Match(NitroParserFN)
	}
	{
		p.SetState(376)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(377)
			p.Param_list()
		}

	}
	{
		p.SetState(380)
		p.Match(NitroParserCPAREN)
	}
	{
		p.SetState(381)
		p.Stmts()
	}
	{
		p.SetState(382)
		p.Match(NitroParserEND)
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

func (s *Object_literalContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
}

func (s *Object_literalContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
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
	localctx = NewObject_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, NitroParserRULE_object_literal)
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
		p.SetState(384)
		p.Match(NitroParserOCURLY)
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserAND)|(1<<NitroParserDO)|(1<<NitroParserELIF)|(1<<NitroParserELSE)|(1<<NitroParserEND)|(1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserIN)|(1<<NitroParserMETA)|(1<<NitroParserNOT)|(1<<NitroParserOR)|(1<<NitroParserRETURN)|(1<<NitroParserTHEN)|(1<<NitroParserTRUE)|(1<<NitroParserVAR)|(1<<NitroParserWHILE))) != 0) || _la == NitroParserOBRACKET || _la == NitroParserID {
		{
			p.SetState(385)
			p.Object_fields()
		}

	}
	{
		p.SetState(388)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObject_fieldContext)(nil)).Elem())
	var tst = make([]IObject_fieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObject_fieldContext)
		}
	}

	return tst
}

func (s *Object_fieldsContext) Object_field(i int) IObject_fieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldContext)(nil)).Elem(), i)

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
	localctx = NewObject_fieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, NitroParserRULE_object_fields)
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
		p.SetState(390)
		p.Object_field()
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(391)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(392)
				p.Object_field()
			}

		}
		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
		{
			p.SetState(398)
			_la = p.GetTokenStream().LA(1)

			if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_or_keywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_or_keywordContext)
}

func (s *Object_field_id_keyContext) COLON() antlr.TerminalNode {
	return s.GetToken(NitroParserCOLON, 0)
}

func (s *Object_field_id_keyContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Object_field_expr_keyContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

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

type Object_field_forContext struct {
	*Object_fieldContext
}

func NewObject_field_forContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Object_field_forContext {
	var p = new(Object_field_forContext)

	p.Object_fieldContext = NewEmptyObject_fieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Object_fieldContext))

	return p
}

func (s *Object_field_forContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_field_forContext) Object_for() IObject_forContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_forContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_forContext)
}

func (s *Object_field_forContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_field_for(s)
	}
}

func (s *Object_field_forContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_field_for(s)
	}
}

type Object_field_ifContext struct {
	*Object_fieldContext
}

func NewObject_field_ifContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Object_field_ifContext {
	var p = new(Object_field_ifContext)

	p.Object_fieldContext = NewEmptyObject_fieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Object_fieldContext))

	return p
}

func (s *Object_field_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_field_ifContext) Object_if() IObject_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_ifContext)
}

func (s *Object_field_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_field_if(s)
	}
}

func (s *Object_field_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_field_if(s)
	}
}

func (p *NitroParser) Object_field() (localctx IObject_fieldContext) {
	localctx = NewObject_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, NitroParserRULE_object_field)

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

	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewObject_field_id_keyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)
			p.Id_or_keyword()
		}
		{
			p.SetState(405)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(406)
			p.expr(0)
		}

	case 2:
		localctx = NewObject_field_expr_keyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(408)
			p.Match(NitroParserOBRACKET)
		}
		{
			p.SetState(409)
			p.expr(0)
		}
		{
			p.SetState(410)
			p.Match(NitroParserCBRACKET)
		}
		{
			p.SetState(411)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(412)
			p.expr(0)
		}

	case 3:
		localctx = NewObject_field_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(414)
			p.Object_if()
		}

	case 4:
		localctx = NewObject_field_forContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(415)
			p.Object_for()
		}

	}

	return localctx
}

// IObject_ifContext is an interface to support dynamic dispatch.
type IObject_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_ifContext differentiates from other interfaces.
	IsObject_ifContext()
}

type Object_ifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_ifContext() *Object_ifContext {
	var p = new(Object_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_object_if
	return p
}

func (*Object_ifContext) IsObject_ifContext() {}

func NewObject_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_ifContext {
	var p = new(Object_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_object_if

	return p
}

func (s *Object_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_ifContext) IF() antlr.TerminalNode {
	return s.GetToken(NitroParserIF, 0)
}

func (s *Object_ifContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Object_ifContext) THEN() antlr.TerminalNode {
	return s.GetToken(NitroParserTHEN, 0)
}

func (s *Object_ifContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
}

func (s *Object_ifContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_ifContext) AllObject_elif() []IObject_elifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObject_elifContext)(nil)).Elem())
	var tst = make([]IObject_elifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObject_elifContext)
		}
	}

	return tst
}

func (s *Object_ifContext) Object_elif(i int) IObject_elifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_elifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObject_elifContext)
}

func (s *Object_ifContext) Object_else() IObject_elseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_elseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_elseContext)
}

func (s *Object_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_if(s)
	}
}

func (s *Object_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_if(s)
	}
}

func (p *NitroParser) Object_if() (localctx IObject_ifContext) {
	localctx = NewObject_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, NitroParserRULE_object_if)
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
		p.SetState(418)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(419)
		p.expr(0)
	}
	{
		p.SetState(420)
		p.Match(NitroParserTHEN)
	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(421)
			p.Object_fields()
		}

	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserELIF {
		{
			p.SetState(424)
			p.Object_elif()
		}

		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(430)
			p.Object_else()
		}

	}
	{
		p.SetState(433)
		p.Match(NitroParserEND)
	}

	return localctx
}

// IObject_elifContext is an interface to support dynamic dispatch.
type IObject_elifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_elifContext differentiates from other interfaces.
	IsObject_elifContext()
}

type Object_elifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_elifContext() *Object_elifContext {
	var p = new(Object_elifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_object_elif
	return p
}

func (*Object_elifContext) IsObject_elifContext() {}

func NewObject_elifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_elifContext {
	var p = new(Object_elifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_object_elif

	return p
}

func (s *Object_elifContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_elifContext) ELIF() antlr.TerminalNode {
	return s.GetToken(NitroParserELIF, 0)
}

func (s *Object_elifContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Object_elifContext) THEN() antlr.TerminalNode {
	return s.GetToken(NitroParserTHEN, 0)
}

func (s *Object_elifContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_elifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_elifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_elifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_elif(s)
	}
}

func (s *Object_elifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_elif(s)
	}
}

func (p *NitroParser) Object_elif() (localctx IObject_elifContext) {
	localctx = NewObject_elifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, NitroParserRULE_object_elif)

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
		p.SetState(435)
		p.Match(NitroParserELIF)
	}
	{
		p.SetState(436)
		p.expr(0)
	}
	{
		p.SetState(437)
		p.Match(NitroParserTHEN)
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(438)
			p.Object_fields()
		}

	}

	return localctx
}

// IObject_elseContext is an interface to support dynamic dispatch.
type IObject_elseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_elseContext differentiates from other interfaces.
	IsObject_elseContext()
}

type Object_elseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_elseContext() *Object_elseContext {
	var p = new(Object_elseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_object_else
	return p
}

func (*Object_elseContext) IsObject_elseContext() {}

func NewObject_elseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_elseContext {
	var p = new(Object_elseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_object_else

	return p
}

func (s *Object_elseContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_elseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NitroParserELSE, 0)
}

func (s *Object_elseContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_elseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_elseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_elseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_else(s)
	}
}

func (s *Object_elseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_else(s)
	}
}

func (p *NitroParser) Object_else() (localctx IObject_elseContext) {
	localctx = NewObject_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, NitroParserRULE_object_else)

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
		p.SetState(441)
		p.Match(NitroParserELSE)
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(442)
			p.Object_fields()
		}

	}

	return localctx
}

// IObject_forContext is an interface to support dynamic dispatch.
type IObject_forContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_forContext differentiates from other interfaces.
	IsObject_forContext()
}

type Object_forContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_forContext() *Object_forContext {
	var p = new(Object_forContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_object_for
	return p
}

func (*Object_forContext) IsObject_forContext() {}

func NewObject_forContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_forContext {
	var p = new(Object_forContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_object_for

	return p
}

func (s *Object_forContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_forContext) FOR() antlr.TerminalNode {
	return s.GetToken(NitroParserFOR, 0)
}

func (s *Object_forContext) For_vars() IFor_varsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_varsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_varsContext)
}

func (s *Object_forContext) IN() antlr.TerminalNode {
	return s.GetToken(NitroParserIN, 0)
}

func (s *Object_forContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Object_forContext) DO() antlr.TerminalNode {
	return s.GetToken(NitroParserDO, 0)
}

func (s *Object_forContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
}

func (s *Object_forContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_forContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_forContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_forContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterObject_for(s)
	}
}

func (s *Object_forContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitObject_for(s)
	}
}

func (p *NitroParser) Object_for() (localctx IObject_forContext) {
	localctx = NewObject_forContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, NitroParserRULE_object_for)

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
		p.SetState(445)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(446)
		p.For_vars()
	}
	{
		p.SetState(447)
		p.Match(NitroParserIN)
	}
	{
		p.SetState(448)
		p.expr(0)
	}
	{
		p.SetState(449)
		p.Match(NitroParserDO)
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(450)
			p.Object_fields()
		}

	}
	{
		p.SetState(453)
		p.Match(NitroParserEND)
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

func (s *Array_literalContext) CBRACKET() antlr.TerminalNode {
	return s.GetToken(NitroParserCBRACKET, 0)
}

func (s *Array_literalContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
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
	localctx = NewArray_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, NitroParserRULE_array_literal)
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
		p.SetState(455)
		p.Match(NitroParserOBRACKET)
	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(456)
			p.Array_elems()
		}

	}
	{
		p.SetState(459)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArray_elemContext)(nil)).Elem())
	var tst = make([]IArray_elemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArray_elemContext)
		}
	}

	return tst
}

func (s *Array_elemsContext) Array_elem(i int) IArray_elemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemContext)(nil)).Elem(), i)

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
	localctx = NewArray_elemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, NitroParserRULE_array_elems)
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
		p.SetState(461)
		p.Array_elem()
	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(462)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(463)
				p.Array_elem()
			}

		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
		{
			p.SetState(469)
			_la = p.GetTokenStream().LA(1)

			if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Array_elemContext) Array_if() IArray_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_ifContext)
}

func (s *Array_elemContext) Array_for() IArray_forContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_forContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_forContext)
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
	localctx = NewArray_elemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, NitroParserRULE_array_elem)

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

	p.SetState(478)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserFALSE, NitroParserFN, NitroParserNOT, NitroParserTRUE, NitroParserADD, NitroParserSUB, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserNUMBER, NitroParserID, NitroParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(475)
			p.expr(0)
		}

	case NitroParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(476)
			p.Array_if()
		}

	case NitroParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(477)
			p.Array_for()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArray_ifContext is an interface to support dynamic dispatch.
type IArray_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_ifContext differentiates from other interfaces.
	IsArray_ifContext()
}

type Array_ifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_ifContext() *Array_ifContext {
	var p = new(Array_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_array_if
	return p
}

func (*Array_ifContext) IsArray_ifContext() {}

func NewArray_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_ifContext {
	var p = new(Array_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_array_if

	return p
}

func (s *Array_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_ifContext) IF() antlr.TerminalNode {
	return s.GetToken(NitroParserIF, 0)
}

func (s *Array_ifContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Array_ifContext) THEN() antlr.TerminalNode {
	return s.GetToken(NitroParserTHEN, 0)
}

func (s *Array_ifContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
}

func (s *Array_ifContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_ifContext) AllArray_elif() []IArray_elifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArray_elifContext)(nil)).Elem())
	var tst = make([]IArray_elifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArray_elifContext)
		}
	}

	return tst
}

func (s *Array_ifContext) Array_elif(i int) IArray_elifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArray_elifContext)
}

func (s *Array_ifContext) Array_else() IArray_elseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elseContext)
}

func (s *Array_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterArray_if(s)
	}
}

func (s *Array_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitArray_if(s)
	}
}

func (p *NitroParser) Array_if() (localctx IArray_ifContext) {
	localctx = NewArray_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, NitroParserRULE_array_if)
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
		p.SetState(480)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(481)
		p.expr(0)
	}
	{
		p.SetState(482)
		p.Match(NitroParserTHEN)
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(483)
			p.Array_elems()
		}

	}
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserELIF {
		{
			p.SetState(486)
			p.Array_elif()
		}

		p.SetState(491)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(492)
			p.Array_else()
		}

	}
	{
		p.SetState(495)
		p.Match(NitroParserEND)
	}

	return localctx
}

// IArray_elifContext is an interface to support dynamic dispatch.
type IArray_elifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_elifContext differentiates from other interfaces.
	IsArray_elifContext()
}

type Array_elifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_elifContext() *Array_elifContext {
	var p = new(Array_elifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_array_elif
	return p
}

func (*Array_elifContext) IsArray_elifContext() {}

func NewArray_elifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_elifContext {
	var p = new(Array_elifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_array_elif

	return p
}

func (s *Array_elifContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_elifContext) ELIF() antlr.TerminalNode {
	return s.GetToken(NitroParserELIF, 0)
}

func (s *Array_elifContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Array_elifContext) THEN() antlr.TerminalNode {
	return s.GetToken(NitroParserTHEN, 0)
}

func (s *Array_elifContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_elifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_elifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_elifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterArray_elif(s)
	}
}

func (s *Array_elifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitArray_elif(s)
	}
}

func (p *NitroParser) Array_elif() (localctx IArray_elifContext) {
	localctx = NewArray_elifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, NitroParserRULE_array_elif)
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
		p.SetState(497)
		p.Match(NitroParserELIF)
	}
	{
		p.SetState(498)
		p.expr(0)
	}
	{
		p.SetState(499)
		p.Match(NitroParserTHEN)
	}
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(500)
			p.Array_elems()
		}

	}

	return localctx
}

// IArray_elseContext is an interface to support dynamic dispatch.
type IArray_elseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_elseContext differentiates from other interfaces.
	IsArray_elseContext()
}

type Array_elseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_elseContext() *Array_elseContext {
	var p = new(Array_elseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_array_else
	return p
}

func (*Array_elseContext) IsArray_elseContext() {}

func NewArray_elseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_elseContext {
	var p = new(Array_elseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_array_else

	return p
}

func (s *Array_elseContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_elseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NitroParserELSE, 0)
}

func (s *Array_elseContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_elseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_elseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_elseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterArray_else(s)
	}
}

func (s *Array_elseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitArray_else(s)
	}
}

func (p *NitroParser) Array_else() (localctx IArray_elseContext) {
	localctx = NewArray_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, NitroParserRULE_array_else)
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
		p.SetState(503)
		p.Match(NitroParserELSE)
	}
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(504)
			p.Array_elems()
		}

	}

	return localctx
}

// IArray_forContext is an interface to support dynamic dispatch.
type IArray_forContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_forContext differentiates from other interfaces.
	IsArray_forContext()
}

type Array_forContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_forContext() *Array_forContext {
	var p = new(Array_forContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_array_for
	return p
}

func (*Array_forContext) IsArray_forContext() {}

func NewArray_forContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_forContext {
	var p = new(Array_forContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_array_for

	return p
}

func (s *Array_forContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_forContext) FOR() antlr.TerminalNode {
	return s.GetToken(NitroParserFOR, 0)
}

func (s *Array_forContext) For_vars() IFor_varsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_varsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_varsContext)
}

func (s *Array_forContext) IN() antlr.TerminalNode {
	return s.GetToken(NitroParserIN, 0)
}

func (s *Array_forContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Array_forContext) DO() antlr.TerminalNode {
	return s.GetToken(NitroParserDO, 0)
}

func (s *Array_forContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
}

func (s *Array_forContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_forContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_forContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_forContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterArray_for(s)
	}
}

func (s *Array_forContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitArray_for(s)
	}
}

func (p *NitroParser) Array_for() (localctx IArray_forContext) {
	localctx = NewArray_forContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, NitroParserRULE_array_for)
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
		p.SetState(507)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(508)
		p.For_vars()
	}
	{
		p.SetState(509)
		p.Match(NitroParserIN)
	}
	{
		p.SetState(510)
		p.expr(0)
	}
	{
		p.SetState(511)
		p.Match(NitroParserDO)
	}
	p.SetState(513)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(512)
			p.Array_elems()
		}

	}
	{
		p.SetState(515)
		p.Match(NitroParserEND)
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

func (s *Id_or_keywordContext) DO() antlr.TerminalNode {
	return s.GetToken(NitroParserDO, 0)
}

func (s *Id_or_keywordContext) ELIF() antlr.TerminalNode {
	return s.GetToken(NitroParserELIF, 0)
}

func (s *Id_or_keywordContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NitroParserELSE, 0)
}

func (s *Id_or_keywordContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
}

func (s *Id_or_keywordContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NitroParserFALSE, 0)
}

func (s *Id_or_keywordContext) FN() antlr.TerminalNode {
	return s.GetToken(NitroParserFN, 0)
}

func (s *Id_or_keywordContext) FOR() antlr.TerminalNode {
	return s.GetToken(NitroParserFOR, 0)
}

func (s *Id_or_keywordContext) IF() antlr.TerminalNode {
	return s.GetToken(NitroParserIF, 0)
}

func (s *Id_or_keywordContext) IN() antlr.TerminalNode {
	return s.GetToken(NitroParserIN, 0)
}

func (s *Id_or_keywordContext) META() antlr.TerminalNode {
	return s.GetToken(NitroParserMETA, 0)
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

func (s *Id_or_keywordContext) THEN() antlr.TerminalNode {
	return s.GetToken(NitroParserTHEN, 0)
}

func (s *Id_or_keywordContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NitroParserTRUE, 0)
}

func (s *Id_or_keywordContext) VAR() antlr.TerminalNode {
	return s.GetToken(NitroParserVAR, 0)
}

func (s *Id_or_keywordContext) WHILE() antlr.TerminalNode {
	return s.GetToken(NitroParserWHILE, 0)
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
	localctx = NewId_or_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, NitroParserRULE_id_or_keyword)
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
		p.SetState(517)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Id_or_keywordContext).t = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserAND)|(1<<NitroParserDO)|(1<<NitroParserELIF)|(1<<NitroParserELSE)|(1<<NitroParserEND)|(1<<NitroParserFALSE)|(1<<NitroParserFN)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserIN)|(1<<NitroParserMETA)|(1<<NitroParserNOT)|(1<<NitroParserOR)|(1<<NitroParserRETURN)|(1<<NitroParserTHEN)|(1<<NitroParserTRUE)|(1<<NitroParserVAR)|(1<<NitroParserWHILE))) != 0) || _la == NitroParserID) {
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
	case 25:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 27:
		var t *Primary_exprContext = nil
		if localctx != nil {
			t = localctx.(*Primary_exprContext)
		}
		return p.Primary_expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NitroParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *NitroParser) Primary_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
