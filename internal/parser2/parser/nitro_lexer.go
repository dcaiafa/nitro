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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 57, 357,
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
	59, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3,
	41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45,
	3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 6, 49, 282,
	10, 49, 13, 49, 14, 49, 283, 3, 49, 3, 49, 6, 49, 288, 10, 49, 13, 49,
	14, 49, 289, 5, 49, 292, 10, 49, 3, 50, 3, 50, 7, 50, 296, 10, 50, 12,
	50, 14, 50, 299, 11, 50, 3, 51, 3, 51, 3, 51, 3, 51, 7, 51, 305, 10, 51,
	12, 51, 14, 51, 308, 11, 51, 3, 51, 3, 51, 3, 52, 5, 52, 313, 10, 52, 3,
	52, 3, 52, 3, 53, 3, 53, 7, 53, 319, 10, 53, 12, 53, 14, 53, 322, 11, 53,
	3, 53, 3, 53, 3, 54, 6, 54, 327, 10, 54, 13, 54, 14, 54, 328, 3, 54, 3,
	54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 5, 57, 348, 10, 57, 3, 57, 3, 57, 3,
	58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 2, 2, 60, 4, 3, 6, 4, 8, 5, 10,
	6, 12, 7, 14, 8, 16, 9, 18, 10, 20, 11, 22, 12, 24, 13, 26, 14, 28, 15,
	30, 16, 32, 17, 34, 18, 36, 19, 38, 20, 40, 21, 42, 22, 44, 23, 46, 24,
	48, 25, 50, 26, 52, 27, 54, 28, 56, 29, 58, 30, 60, 31, 62, 32, 64, 33,
	66, 34, 68, 35, 70, 36, 72, 37, 74, 38, 76, 39, 78, 40, 80, 41, 82, 42,
	84, 43, 86, 44, 88, 45, 90, 46, 92, 47, 94, 48, 96, 49, 98, 50, 100, 51,
	102, 52, 104, 53, 106, 54, 108, 55, 110, 57, 112, 2, 114, 2, 116, 2, 118,
	56, 4, 2, 3, 11, 3, 2, 50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50,
	59, 67, 92, 97, 97, 99, 124, 3, 2, 98, 98, 3, 2, 12, 12, 4, 2, 11, 11,
	34, 34, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 7, 2, 36, 36, 94, 94, 112,
	112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72, 99, 104, 2, 363, 2, 4, 3,
	2, 2, 2, 2, 6, 3, 2, 2, 2, 2, 8, 3, 2, 2, 2, 2, 10, 3, 2, 2, 2, 2, 12,
	3, 2, 2, 2, 2, 14, 3, 2, 2, 2, 2, 16, 3, 2, 2, 2, 2, 18, 3, 2, 2, 2, 2,
	20, 3, 2, 2, 2, 2, 22, 3, 2, 2, 2, 2, 24, 3, 2, 2, 2, 2, 26, 3, 2, 2, 2,
	2, 28, 3, 2, 2, 2, 2, 30, 3, 2, 2, 2, 2, 32, 3, 2, 2, 2, 2, 34, 3, 2, 2,
	2, 2, 36, 3, 2, 2, 2, 2, 38, 3, 2, 2, 2, 2, 40, 3, 2, 2, 2, 2, 42, 3, 2,
	2, 2, 2, 44, 3, 2, 2, 2, 2, 46, 3, 2, 2, 2, 2, 48, 3, 2, 2, 2, 2, 50, 3,
	2, 2, 2, 2, 52, 3, 2, 2, 2, 2, 54, 3, 2, 2, 2, 2, 56, 3, 2, 2, 2, 2, 58,
	3, 2, 2, 2, 2, 60, 3, 2, 2, 2, 2, 62, 3, 2, 2, 2, 2, 64, 3, 2, 2, 2, 2,
	66, 3, 2, 2, 2, 2, 68, 3, 2, 2, 2, 2, 70, 3, 2, 2, 2, 2, 72, 3, 2, 2, 2,
	2, 74, 3, 2, 2, 2, 2, 76, 3, 2, 2, 2, 2, 78, 3, 2, 2, 2, 2, 80, 3, 2, 2,
	2, 2, 82, 3, 2, 2, 2, 2, 84, 3, 2, 2, 2, 2, 86, 3, 2, 2, 2, 2, 88, 3, 2,
	2, 2, 2, 90, 3, 2, 2, 2, 2, 92, 3, 2, 2, 2, 2, 94, 3, 2, 2, 2, 2, 96, 3,
	2, 2, 2, 2, 98, 3, 2, 2, 2, 2, 100, 3, 2, 2, 2, 2, 102, 3, 2, 2, 2, 2,
	104, 3, 2, 2, 2, 2, 106, 3, 2, 2, 2, 2, 108, 3, 2, 2, 2, 2, 110, 3, 2,
	2, 2, 3, 112, 3, 2, 2, 2, 3, 114, 3, 2, 2, 2, 3, 118, 3, 2, 2, 2, 4, 120,
	3, 2, 2, 2, 6, 124, 3, 2, 2, 2, 8, 130, 3, 2, 2, 2, 10, 136, 3, 2, 2, 2,
	12, 142, 3, 2, 2, 2, 14, 147, 3, 2, 2, 2, 16, 153, 3, 2, 2, 2, 18, 157,
	3, 2, 2, 2, 20, 162, 3, 2, 2, 2, 22, 165, 3, 2, 2, 2, 24, 170, 3, 2, 2,
	2, 26, 174, 3, 2, 2, 2, 28, 178, 3, 2, 2, 2, 30, 181, 3, 2, 2, 2, 32, 188,
	3, 2, 2, 2, 34, 194, 3, 2, 2, 2, 36, 199, 3, 2, 2, 2, 38, 203, 3, 2, 2,
	2, 40, 207, 3, 2, 2, 2, 42, 213, 3, 2, 2, 2, 44, 219, 3, 2, 2, 2, 46, 221,
	3, 2, 2, 2, 48, 224, 3, 2, 2, 2, 50, 227, 3, 2, 2, 2, 52, 229, 3, 2, 2,
	2, 54, 232, 3, 2, 2, 2, 56, 234, 3, 2, 2, 2, 58, 237, 3, 2, 2, 2, 60, 239,
	3, 2, 2, 2, 62, 241, 3, 2, 2, 2, 64, 243, 3, 2, 2, 2, 66, 245, 3, 2, 2,
	2, 68, 247, 3, 2, 2, 2, 70, 249, 3, 2, 2, 2, 72, 251, 3, 2, 2, 2, 74, 253,
	3, 2, 2, 2, 76, 255, 3, 2, 2, 2, 78, 257, 3, 2, 2, 2, 80, 259, 3, 2, 2,
	2, 82, 261, 3, 2, 2, 2, 84, 263, 3, 2, 2, 2, 86, 265, 3, 2, 2, 2, 88, 267,
	3, 2, 2, 2, 90, 269, 3, 2, 2, 2, 92, 272, 3, 2, 2, 2, 94, 274, 3, 2, 2,
	2, 96, 276, 3, 2, 2, 2, 98, 281, 3, 2, 2, 2, 100, 293, 3, 2, 2, 2, 102,
	300, 3, 2, 2, 2, 104, 312, 3, 2, 2, 2, 106, 316, 3, 2, 2, 2, 108, 326,
	3, 2, 2, 2, 110, 332, 3, 2, 2, 2, 112, 337, 3, 2, 2, 2, 114, 341, 3, 2,
	2, 2, 116, 351, 3, 2, 2, 2, 118, 353, 3, 2, 2, 2, 120, 121, 7, 99, 2, 2,
	121, 122, 7, 112, 2, 2, 122, 123, 7, 102, 2, 2, 123, 5, 3, 2, 2, 2, 124,
	125, 7, 100, 2, 2, 125, 126, 7, 116, 2, 2, 126, 127, 7, 103, 2, 2, 127,
	128, 7, 99, 2, 2, 128, 129, 7, 109, 2, 2, 129, 7, 3, 2, 2, 2, 130, 131,
	7, 101, 2, 2, 131, 132, 7, 99, 2, 2, 132, 133, 7, 118, 2, 2, 133, 134,
	7, 101, 2, 2, 134, 135, 7, 106, 2, 2, 135, 9, 3, 2, 2, 2, 136, 137, 7,
	102, 2, 2, 137, 138, 7, 103, 2, 2, 138, 139, 7, 104, 2, 2, 139, 140, 7,
	103, 2, 2, 140, 141, 7, 116, 2, 2, 141, 11, 3, 2, 2, 2, 142, 143, 7, 103,
	2, 2, 143, 144, 7, 110, 2, 2, 144, 145, 7, 117, 2, 2, 145, 146, 7, 103,
	2, 2, 146, 13, 3, 2, 2, 2, 147, 148, 7, 104, 2, 2, 148, 149, 7, 99, 2,
	2, 149, 150, 7, 110, 2, 2, 150, 151, 7, 117, 2, 2, 151, 152, 7, 103, 2,
	2, 152, 15, 3, 2, 2, 2, 153, 154, 7, 104, 2, 2, 154, 155, 7, 113, 2, 2,
	155, 156, 7, 116, 2, 2, 156, 17, 3, 2, 2, 2, 157, 158, 7, 104, 2, 2, 158,
	159, 7, 119, 2, 2, 159, 160, 7, 112, 2, 2, 160, 161, 7, 101, 2, 2, 161,
	19, 3, 2, 2, 2, 162, 163, 7, 107, 2, 2, 163, 164, 7, 104, 2, 2, 164, 21,
	3, 2, 2, 2, 165, 166, 7, 111, 2, 2, 166, 167, 7, 103, 2, 2, 167, 168, 7,
	118, 2, 2, 168, 169, 7, 99, 2, 2, 169, 23, 3, 2, 2, 2, 170, 171, 7, 112,
	2, 2, 171, 172, 7, 107, 2, 2, 172, 173, 7, 110, 2, 2, 173, 25, 3, 2, 2,
	2, 174, 175, 7, 112, 2, 2, 175, 176, 7, 113, 2, 2, 176, 177, 7, 118, 2,
	2, 177, 27, 3, 2, 2, 2, 178, 179, 7, 113, 2, 2, 179, 180, 7, 116, 2, 2,
	180, 29, 3, 2, 2, 2, 181, 182, 7, 116, 2, 2, 182, 183, 7, 103, 2, 2, 183,
	184, 7, 118, 2, 2, 184, 185, 7, 119, 2, 2, 185, 186, 7, 116, 2, 2, 186,
	187, 7, 112, 2, 2, 187, 31, 3, 2, 2, 2, 188, 189, 7, 118, 2, 2, 189, 190,
	7, 106, 2, 2, 190, 191, 7, 116, 2, 2, 191, 192, 7, 113, 2, 2, 192, 193,
	7, 121, 2, 2, 193, 33, 3, 2, 2, 2, 194, 195, 7, 118, 2, 2, 195, 196, 7,
	116, 2, 2, 196, 197, 7, 119, 2, 2, 197, 198, 7, 103, 2, 2, 198, 35, 3,
	2, 2, 2, 199, 200, 7, 118, 2, 2, 200, 201, 7, 116, 2, 2, 201, 202, 7, 123,
	2, 2, 202, 37, 3, 2, 2, 2, 203, 204, 7, 120, 2, 2, 204, 205, 7, 99, 2,
	2, 205, 206, 7, 116, 2, 2, 206, 39, 3, 2, 2, 2, 207, 208, 7, 121, 2, 2,
	208, 209, 7, 106, 2, 2, 209, 210, 7, 107, 2, 2, 210, 211, 7, 110, 2, 2,
	211, 212, 7, 103, 2, 2, 212, 41, 3, 2, 2, 2, 213, 214, 7, 123, 2, 2, 214,
	215, 7, 107, 2, 2, 215, 216, 7, 103, 2, 2, 216, 217, 7, 110, 2, 2, 217,
	218, 7, 102, 2, 2, 218, 43, 3, 2, 2, 2, 219, 220, 7, 63, 2, 2, 220, 45,
	3, 2, 2, 2, 221, 222, 7, 63, 2, 2, 222, 223, 7, 63, 2, 2, 223, 47, 3, 2,
	2, 2, 224, 225, 7, 35, 2, 2, 225, 226, 7, 63, 2, 2, 226, 49, 3, 2, 2, 2,
	227, 228, 7, 62, 2, 2, 228, 51, 3, 2, 2, 2, 229, 230, 7, 62, 2, 2, 230,
	231, 7, 63, 2, 2, 231, 53, 3, 2, 2, 2, 232, 233, 7, 64, 2, 2, 233, 55,
	3, 2, 2, 2, 234, 235, 7, 64, 2, 2, 235, 236, 7, 63, 2, 2, 236, 57, 3, 2,
	2, 2, 237, 238, 7, 45, 2, 2, 238, 59, 3, 2, 2, 2, 239, 240, 7, 47, 2, 2,
	240, 61, 3, 2, 2, 2, 241, 242, 7, 44, 2, 2, 242, 63, 3, 2, 2, 2, 243, 244,
	7, 49, 2, 2, 244, 65, 3, 2, 2, 2, 245, 246, 7, 39, 2, 2, 246, 67, 3, 2,
	2, 2, 247, 248, 7, 65, 2, 2, 248, 69, 3, 2, 2, 2, 249, 250, 7, 61, 2, 2,
	250, 71, 3, 2, 2, 2, 251, 252, 7, 46, 2, 2, 252, 73, 3, 2, 2, 2, 253, 254,
	7, 60, 2, 2, 254, 75, 3, 2, 2, 2, 255, 256, 7, 48, 2, 2, 256, 77, 3, 2,
	2, 2, 257, 258, 7, 42, 2, 2, 258, 79, 3, 2, 2, 2, 259, 260, 7, 43, 2, 2,
	260, 81, 3, 2, 2, 2, 261, 262, 7, 93, 2, 2, 262, 83, 3, 2, 2, 2, 263, 264,
	7, 95, 2, 2, 264, 85, 3, 2, 2, 2, 265, 266, 7, 125, 2, 2, 266, 87, 3, 2,
	2, 2, 267, 268, 7, 127, 2, 2, 268, 89, 3, 2, 2, 2, 269, 270, 7, 47, 2,
	2, 270, 271, 7, 64, 2, 2, 271, 91, 3, 2, 2, 2, 272, 273, 7, 40, 2, 2, 273,
	93, 3, 2, 2, 2, 274, 275, 7, 126, 2, 2, 275, 95, 3, 2, 2, 2, 276, 277,
	7, 48, 2, 2, 277, 278, 7, 48, 2, 2, 278, 279, 7, 48, 2, 2, 279, 97, 3,
	2, 2, 2, 280, 282, 9, 2, 2, 2, 281, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2,
	2, 283, 281, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 291, 3, 2, 2, 2, 285,
	287, 7, 48, 2, 2, 286, 288, 9, 2, 2, 2, 287, 286, 3, 2, 2, 2, 288, 289,
	3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 292, 3, 2,
	2, 2, 291, 285, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 99, 3, 2, 2, 2,
	293, 297, 9, 3, 2, 2, 294, 296, 9, 4, 2, 2, 295, 294, 3, 2, 2, 2, 296,
	299, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 101,
	3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 300, 301, 7, 116, 2, 2, 301, 302, 7,
	98, 2, 2, 302, 306, 3, 2, 2, 2, 303, 305, 10, 5, 2, 2, 304, 303, 3, 2,
	2, 2, 305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2,
	307, 309, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 310, 7, 98, 2, 2, 310,
	103, 3, 2, 2, 2, 311, 313, 7, 15, 2, 2, 312, 311, 3, 2, 2, 2, 312, 313,
	3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 315, 7, 12, 2, 2, 315, 105, 3, 2,
	2, 2, 316, 320, 7, 37, 2, 2, 317, 319, 10, 6, 2, 2, 318, 317, 3, 2, 2,
	2, 319, 322, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321,
	323, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 323, 324, 8, 53, 2, 2, 324, 107,
	3, 2, 2, 2, 325, 327, 9, 7, 2, 2, 326, 325, 3, 2, 2, 2, 327, 328, 3, 2,
	2, 2, 328, 326, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2,
	330, 331, 8, 54, 2, 2, 331, 109, 3, 2, 2, 2, 332, 333, 7, 36, 2, 2, 333,
	334, 3, 2, 2, 2, 334, 335, 8, 55, 3, 2, 335, 336, 8, 55, 4, 2, 336, 111,
	3, 2, 2, 2, 337, 338, 10, 8, 2, 2, 338, 339, 3, 2, 2, 2, 339, 340, 8, 56,
	3, 2, 340, 113, 3, 2, 2, 2, 341, 347, 7, 94, 2, 2, 342, 348, 9, 9, 2, 2,
	343, 344, 7, 122, 2, 2, 344, 345, 5, 116, 58, 2, 345, 346, 5, 116, 58,
	2, 346, 348, 3, 2, 2, 2, 347, 342, 3, 2, 2, 2, 347, 343, 3, 2, 2, 2, 348,
	349, 3, 2, 2, 2, 349, 350, 8, 57, 3, 2, 350, 115, 3, 2, 2, 2, 351, 352,
	9, 10, 2, 2, 352, 117, 3, 2, 2, 2, 353, 354, 7, 36, 2, 2, 354, 355, 3,
	2, 2, 2, 355, 356, 8, 59, 5, 2, 356, 119, 3, 2, 2, 2, 13, 2, 3, 283, 289,
	291, 297, 306, 312, 320, 328, 347, 6, 8, 2, 2, 5, 2, 2, 4, 3, 2, 4, 2,
	2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "QSTR",
}

var lexerLiteralNames = []string{
	"", "'and'", "'break'", "'catch'", "'defer'", "'else'", "'false'", "'for'",
	"'func'", "'if'", "'meta'", "'nil'", "'not'", "'or'", "'return'", "'throw'",
	"'true'", "'try'", "'var'", "'while'", "'yield'", "'='", "'=='", "'!='",
	"'<'", "'<='", "'>'", "'>='", "'+'", "'-'", "'*'", "'/'", "'%'", "'?'",
	"';'", "','", "':'", "'.'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'->'",
	"'&'", "'|'", "'...'",
}

var lexerSymbolicNames = []string{
	"", "AND", "BREAK", "CATCH", "DEFER", "ELSE", "FALSE", "FOR", "FUNC", "IF",
	"META", "NIL", "NOT", "OR", "RETURN", "THROW", "TRUE", "TRY", "VAR", "WHILE",
	"YIELD", "ASSIGN", "EQ", "NE", "LT", "LE", "GT", "GE", "ADD", "SUB", "MUL",
	"DIV", "MOD", "QUESTION_MARK", "SEMICOLON", "COMMA", "COLON", "PERIOD",
	"OPAREN", "CPAREN", "OBRACKET", "CBRACKET", "OCURLY", "CCURLY", "ARROW",
	"LAMBDA", "PIPE", "EXPAND", "NUMBER", "ID", "REGEX", "NEWLINE", "COMMENT",
	"WS", "STRING", "LQUOTE",
}

var lexerRuleNames = []string{
	"AND", "BREAK", "CATCH", "DEFER", "ELSE", "FALSE", "FOR", "FUNC", "IF",
	"META", "NIL", "NOT", "OR", "RETURN", "THROW", "TRUE", "TRY", "VAR", "WHILE",
	"YIELD", "ASSIGN", "EQ", "NE", "LT", "LE", "GT", "GE", "ADD", "SUB", "MUL",
	"DIV", "MOD", "QUESTION_MARK", "SEMICOLON", "COMMA", "COLON", "PERIOD",
	"OPAREN", "CPAREN", "OBRACKET", "CBRACKET", "OCURLY", "CCURLY", "ARROW",
	"LAMBDA", "PIPE", "EXPAND", "NUMBER", "ID", "REGEX", "NEWLINE", "COMMENT",
	"WS", "LQUOTE", "QUOTED_TEXT", "QUOTED_ESCAPE", "HEX_DIGIT", "STRING",
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
	NitroLexerDEFER         = 4
	NitroLexerELSE          = 5
	NitroLexerFALSE         = 6
	NitroLexerFOR           = 7
	NitroLexerFUNC          = 8
	NitroLexerIF            = 9
	NitroLexerMETA          = 10
	NitroLexerNIL           = 11
	NitroLexerNOT           = 12
	NitroLexerOR            = 13
	NitroLexerRETURN        = 14
	NitroLexerTHROW         = 15
	NitroLexerTRUE          = 16
	NitroLexerTRY           = 17
	NitroLexerVAR           = 18
	NitroLexerWHILE         = 19
	NitroLexerYIELD         = 20
	NitroLexerASSIGN        = 21
	NitroLexerEQ            = 22
	NitroLexerNE            = 23
	NitroLexerLT            = 24
	NitroLexerLE            = 25
	NitroLexerGT            = 26
	NitroLexerGE            = 27
	NitroLexerADD           = 28
	NitroLexerSUB           = 29
	NitroLexerMUL           = 30
	NitroLexerDIV           = 31
	NitroLexerMOD           = 32
	NitroLexerQUESTION_MARK = 33
	NitroLexerSEMICOLON     = 34
	NitroLexerCOMMA         = 35
	NitroLexerCOLON         = 36
	NitroLexerPERIOD        = 37
	NitroLexerOPAREN        = 38
	NitroLexerCPAREN        = 39
	NitroLexerOBRACKET      = 40
	NitroLexerCBRACKET      = 41
	NitroLexerOCURLY        = 42
	NitroLexerCCURLY        = 43
	NitroLexerARROW         = 44
	NitroLexerLAMBDA        = 45
	NitroLexerPIPE          = 46
	NitroLexerEXPAND        = 47
	NitroLexerNUMBER        = 48
	NitroLexerID            = 49
	NitroLexerREGEX         = 50
	NitroLexerNEWLINE       = 51
	NitroLexerCOMMENT       = 52
	NitroLexerWS            = 53
	NitroLexerSTRING        = 54
	NitroLexerLQUOTE        = 55
)

// NitroLexerQSTR is the NitroLexer mode.
const NitroLexerQSTR = 1
