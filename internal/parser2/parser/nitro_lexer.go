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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 58, 368,
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
	59, 4, 60, 9, 60, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3,
	36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41,
	3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3,
	46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 6, 50,
	293, 10, 50, 13, 50, 14, 50, 294, 3, 50, 3, 50, 6, 50, 299, 10, 50, 13,
	50, 14, 50, 300, 5, 50, 303, 10, 50, 3, 51, 3, 51, 7, 51, 307, 10, 51,
	12, 51, 14, 51, 310, 11, 51, 3, 52, 3, 52, 3, 52, 3, 52, 7, 52, 316, 10,
	52, 12, 52, 14, 52, 319, 11, 52, 3, 52, 3, 52, 3, 53, 5, 53, 324, 10, 53,
	3, 53, 3, 53, 3, 54, 3, 54, 7, 54, 330, 10, 54, 12, 54, 14, 54, 333, 11,
	54, 3, 54, 3, 54, 3, 55, 6, 55, 338, 10, 55, 13, 55, 14, 55, 339, 3, 55,
	3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3,
	58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 5, 58, 359, 10, 58, 3, 58, 3, 58,
	3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 2, 2, 61, 4, 3, 6, 4, 8, 5, 10,
	6, 12, 7, 14, 8, 16, 9, 18, 10, 20, 11, 22, 12, 24, 13, 26, 14, 28, 15,
	30, 16, 32, 17, 34, 18, 36, 19, 38, 20, 40, 21, 42, 22, 44, 23, 46, 24,
	48, 25, 50, 26, 52, 27, 54, 28, 56, 29, 58, 30, 60, 31, 62, 32, 64, 33,
	66, 34, 68, 35, 70, 36, 72, 37, 74, 38, 76, 39, 78, 40, 80, 41, 82, 42,
	84, 43, 86, 44, 88, 45, 90, 46, 92, 47, 94, 48, 96, 49, 98, 50, 100, 51,
	102, 52, 104, 53, 106, 54, 108, 55, 110, 56, 112, 58, 114, 2, 116, 2, 118,
	2, 120, 57, 4, 2, 3, 11, 3, 2, 50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 6,
	2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 98, 98, 3, 2, 12, 12, 4, 2, 11,
	11, 34, 34, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 7, 2, 36, 36, 94, 94,
	112, 112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72, 99, 104, 2, 374, 2,
	4, 3, 2, 2, 2, 2, 6, 3, 2, 2, 2, 2, 8, 3, 2, 2, 2, 2, 10, 3, 2, 2, 2, 2,
	12, 3, 2, 2, 2, 2, 14, 3, 2, 2, 2, 2, 16, 3, 2, 2, 2, 2, 18, 3, 2, 2, 2,
	2, 20, 3, 2, 2, 2, 2, 22, 3, 2, 2, 2, 2, 24, 3, 2, 2, 2, 2, 26, 3, 2, 2,
	2, 2, 28, 3, 2, 2, 2, 2, 30, 3, 2, 2, 2, 2, 32, 3, 2, 2, 2, 2, 34, 3, 2,
	2, 2, 2, 36, 3, 2, 2, 2, 2, 38, 3, 2, 2, 2, 2, 40, 3, 2, 2, 2, 2, 42, 3,
	2, 2, 2, 2, 44, 3, 2, 2, 2, 2, 46, 3, 2, 2, 2, 2, 48, 3, 2, 2, 2, 2, 50,
	3, 2, 2, 2, 2, 52, 3, 2, 2, 2, 2, 54, 3, 2, 2, 2, 2, 56, 3, 2, 2, 2, 2,
	58, 3, 2, 2, 2, 2, 60, 3, 2, 2, 2, 2, 62, 3, 2, 2, 2, 2, 64, 3, 2, 2, 2,
	2, 66, 3, 2, 2, 2, 2, 68, 3, 2, 2, 2, 2, 70, 3, 2, 2, 2, 2, 72, 3, 2, 2,
	2, 2, 74, 3, 2, 2, 2, 2, 76, 3, 2, 2, 2, 2, 78, 3, 2, 2, 2, 2, 80, 3, 2,
	2, 2, 2, 82, 3, 2, 2, 2, 2, 84, 3, 2, 2, 2, 2, 86, 3, 2, 2, 2, 2, 88, 3,
	2, 2, 2, 2, 90, 3, 2, 2, 2, 2, 92, 3, 2, 2, 2, 2, 94, 3, 2, 2, 2, 2, 96,
	3, 2, 2, 2, 2, 98, 3, 2, 2, 2, 2, 100, 3, 2, 2, 2, 2, 102, 3, 2, 2, 2,
	2, 104, 3, 2, 2, 2, 2, 106, 3, 2, 2, 2, 2, 108, 3, 2, 2, 2, 2, 110, 3,
	2, 2, 2, 2, 112, 3, 2, 2, 2, 3, 114, 3, 2, 2, 2, 3, 116, 3, 2, 2, 2, 3,
	120, 3, 2, 2, 2, 4, 122, 3, 2, 2, 2, 6, 126, 3, 2, 2, 2, 8, 132, 3, 2,
	2, 2, 10, 138, 3, 2, 2, 2, 12, 147, 3, 2, 2, 2, 14, 153, 3, 2, 2, 2, 16,
	158, 3, 2, 2, 2, 18, 164, 3, 2, 2, 2, 20, 168, 3, 2, 2, 2, 22, 173, 3,
	2, 2, 2, 24, 176, 3, 2, 2, 2, 26, 181, 3, 2, 2, 2, 28, 185, 3, 2, 2, 2,
	30, 189, 3, 2, 2, 2, 32, 192, 3, 2, 2, 2, 34, 199, 3, 2, 2, 2, 36, 205,
	3, 2, 2, 2, 38, 210, 3, 2, 2, 2, 40, 214, 3, 2, 2, 2, 42, 218, 3, 2, 2,
	2, 44, 224, 3, 2, 2, 2, 46, 230, 3, 2, 2, 2, 48, 232, 3, 2, 2, 2, 50, 235,
	3, 2, 2, 2, 52, 238, 3, 2, 2, 2, 54, 240, 3, 2, 2, 2, 56, 243, 3, 2, 2,
	2, 58, 245, 3, 2, 2, 2, 60, 248, 3, 2, 2, 2, 62, 250, 3, 2, 2, 2, 64, 252,
	3, 2, 2, 2, 66, 254, 3, 2, 2, 2, 68, 256, 3, 2, 2, 2, 70, 258, 3, 2, 2,
	2, 72, 260, 3, 2, 2, 2, 74, 262, 3, 2, 2, 2, 76, 264, 3, 2, 2, 2, 78, 266,
	3, 2, 2, 2, 80, 268, 3, 2, 2, 2, 82, 270, 3, 2, 2, 2, 84, 272, 3, 2, 2,
	2, 86, 274, 3, 2, 2, 2, 88, 276, 3, 2, 2, 2, 90, 278, 3, 2, 2, 2, 92, 280,
	3, 2, 2, 2, 94, 283, 3, 2, 2, 2, 96, 285, 3, 2, 2, 2, 98, 287, 3, 2, 2,
	2, 100, 292, 3, 2, 2, 2, 102, 304, 3, 2, 2, 2, 104, 311, 3, 2, 2, 2, 106,
	323, 3, 2, 2, 2, 108, 327, 3, 2, 2, 2, 110, 337, 3, 2, 2, 2, 112, 343,
	3, 2, 2, 2, 114, 348, 3, 2, 2, 2, 116, 352, 3, 2, 2, 2, 118, 362, 3, 2,
	2, 2, 120, 364, 3, 2, 2, 2, 122, 123, 7, 99, 2, 2, 123, 124, 7, 112, 2,
	2, 124, 125, 7, 102, 2, 2, 125, 5, 3, 2, 2, 2, 126, 127, 7, 100, 2, 2,
	127, 128, 7, 116, 2, 2, 128, 129, 7, 103, 2, 2, 129, 130, 7, 99, 2, 2,
	130, 131, 7, 109, 2, 2, 131, 7, 3, 2, 2, 2, 132, 133, 7, 101, 2, 2, 133,
	134, 7, 99, 2, 2, 134, 135, 7, 118, 2, 2, 135, 136, 7, 101, 2, 2, 136,
	137, 7, 106, 2, 2, 137, 9, 3, 2, 2, 2, 138, 139, 7, 101, 2, 2, 139, 140,
	7, 113, 2, 2, 140, 141, 7, 112, 2, 2, 141, 142, 7, 118, 2, 2, 142, 143,
	7, 107, 2, 2, 143, 144, 7, 112, 2, 2, 144, 145, 7, 119, 2, 2, 145, 146,
	7, 103, 2, 2, 146, 11, 3, 2, 2, 2, 147, 148, 7, 102, 2, 2, 148, 149, 7,
	103, 2, 2, 149, 150, 7, 104, 2, 2, 150, 151, 7, 103, 2, 2, 151, 152, 7,
	116, 2, 2, 152, 13, 3, 2, 2, 2, 153, 154, 7, 103, 2, 2, 154, 155, 7, 110,
	2, 2, 155, 156, 7, 117, 2, 2, 156, 157, 7, 103, 2, 2, 157, 15, 3, 2, 2,
	2, 158, 159, 7, 104, 2, 2, 159, 160, 7, 99, 2, 2, 160, 161, 7, 110, 2,
	2, 161, 162, 7, 117, 2, 2, 162, 163, 7, 103, 2, 2, 163, 17, 3, 2, 2, 2,
	164, 165, 7, 104, 2, 2, 165, 166, 7, 113, 2, 2, 166, 167, 7, 116, 2, 2,
	167, 19, 3, 2, 2, 2, 168, 169, 7, 104, 2, 2, 169, 170, 7, 119, 2, 2, 170,
	171, 7, 112, 2, 2, 171, 172, 7, 101, 2, 2, 172, 21, 3, 2, 2, 2, 173, 174,
	7, 107, 2, 2, 174, 175, 7, 104, 2, 2, 175, 23, 3, 2, 2, 2, 176, 177, 7,
	111, 2, 2, 177, 178, 7, 103, 2, 2, 178, 179, 7, 118, 2, 2, 179, 180, 7,
	99, 2, 2, 180, 25, 3, 2, 2, 2, 181, 182, 7, 112, 2, 2, 182, 183, 7, 107,
	2, 2, 183, 184, 7, 110, 2, 2, 184, 27, 3, 2, 2, 2, 185, 186, 7, 112, 2,
	2, 186, 187, 7, 113, 2, 2, 187, 188, 7, 118, 2, 2, 188, 29, 3, 2, 2, 2,
	189, 190, 7, 113, 2, 2, 190, 191, 7, 116, 2, 2, 191, 31, 3, 2, 2, 2, 192,
	193, 7, 116, 2, 2, 193, 194, 7, 103, 2, 2, 194, 195, 7, 118, 2, 2, 195,
	196, 7, 119, 2, 2, 196, 197, 7, 116, 2, 2, 197, 198, 7, 112, 2, 2, 198,
	33, 3, 2, 2, 2, 199, 200, 7, 118, 2, 2, 200, 201, 7, 106, 2, 2, 201, 202,
	7, 116, 2, 2, 202, 203, 7, 113, 2, 2, 203, 204, 7, 121, 2, 2, 204, 35,
	3, 2, 2, 2, 205, 206, 7, 118, 2, 2, 206, 207, 7, 116, 2, 2, 207, 208, 7,
	119, 2, 2, 208, 209, 7, 103, 2, 2, 209, 37, 3, 2, 2, 2, 210, 211, 7, 118,
	2, 2, 211, 212, 7, 116, 2, 2, 212, 213, 7, 123, 2, 2, 213, 39, 3, 2, 2,
	2, 214, 215, 7, 120, 2, 2, 215, 216, 7, 99, 2, 2, 216, 217, 7, 116, 2,
	2, 217, 41, 3, 2, 2, 2, 218, 219, 7, 121, 2, 2, 219, 220, 7, 106, 2, 2,
	220, 221, 7, 107, 2, 2, 221, 222, 7, 110, 2, 2, 222, 223, 7, 103, 2, 2,
	223, 43, 3, 2, 2, 2, 224, 225, 7, 123, 2, 2, 225, 226, 7, 107, 2, 2, 226,
	227, 7, 103, 2, 2, 227, 228, 7, 110, 2, 2, 228, 229, 7, 102, 2, 2, 229,
	45, 3, 2, 2, 2, 230, 231, 7, 63, 2, 2, 231, 47, 3, 2, 2, 2, 232, 233, 7,
	63, 2, 2, 233, 234, 7, 63, 2, 2, 234, 49, 3, 2, 2, 2, 235, 236, 7, 35,
	2, 2, 236, 237, 7, 63, 2, 2, 237, 51, 3, 2, 2, 2, 238, 239, 7, 62, 2, 2,
	239, 53, 3, 2, 2, 2, 240, 241, 7, 62, 2, 2, 241, 242, 7, 63, 2, 2, 242,
	55, 3, 2, 2, 2, 243, 244, 7, 64, 2, 2, 244, 57, 3, 2, 2, 2, 245, 246, 7,
	64, 2, 2, 246, 247, 7, 63, 2, 2, 247, 59, 3, 2, 2, 2, 248, 249, 7, 45,
	2, 2, 249, 61, 3, 2, 2, 2, 250, 251, 7, 47, 2, 2, 251, 63, 3, 2, 2, 2,
	252, 253, 7, 44, 2, 2, 253, 65, 3, 2, 2, 2, 254, 255, 7, 49, 2, 2, 255,
	67, 3, 2, 2, 2, 256, 257, 7, 39, 2, 2, 257, 69, 3, 2, 2, 2, 258, 259, 7,
	65, 2, 2, 259, 71, 3, 2, 2, 2, 260, 261, 7, 61, 2, 2, 261, 73, 3, 2, 2,
	2, 262, 263, 7, 46, 2, 2, 263, 75, 3, 2, 2, 2, 264, 265, 7, 60, 2, 2, 265,
	77, 3, 2, 2, 2, 266, 267, 7, 48, 2, 2, 267, 79, 3, 2, 2, 2, 268, 269, 7,
	42, 2, 2, 269, 81, 3, 2, 2, 2, 270, 271, 7, 43, 2, 2, 271, 83, 3, 2, 2,
	2, 272, 273, 7, 93, 2, 2, 273, 85, 3, 2, 2, 2, 274, 275, 7, 95, 2, 2, 275,
	87, 3, 2, 2, 2, 276, 277, 7, 125, 2, 2, 277, 89, 3, 2, 2, 2, 278, 279,
	7, 127, 2, 2, 279, 91, 3, 2, 2, 2, 280, 281, 7, 47, 2, 2, 281, 282, 7,
	64, 2, 2, 282, 93, 3, 2, 2, 2, 283, 284, 7, 40, 2, 2, 284, 95, 3, 2, 2,
	2, 285, 286, 7, 126, 2, 2, 286, 97, 3, 2, 2, 2, 287, 288, 7, 48, 2, 2,
	288, 289, 7, 48, 2, 2, 289, 290, 7, 48, 2, 2, 290, 99, 3, 2, 2, 2, 291,
	293, 9, 2, 2, 2, 292, 291, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294, 292,
	3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 302, 3, 2, 2, 2, 296, 298, 7, 48,
	2, 2, 297, 299, 9, 2, 2, 2, 298, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2,
	300, 298, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 303, 3, 2, 2, 2, 302,
	296, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 101, 3, 2, 2, 2, 304, 308,
	9, 3, 2, 2, 305, 307, 9, 4, 2, 2, 306, 305, 3, 2, 2, 2, 307, 310, 3, 2,
	2, 2, 308, 306, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 103, 3, 2, 2, 2,
	310, 308, 3, 2, 2, 2, 311, 312, 7, 116, 2, 2, 312, 313, 7, 98, 2, 2, 313,
	317, 3, 2, 2, 2, 314, 316, 10, 5, 2, 2, 315, 314, 3, 2, 2, 2, 316, 319,
	3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 320, 3, 2,
	2, 2, 319, 317, 3, 2, 2, 2, 320, 321, 7, 98, 2, 2, 321, 105, 3, 2, 2, 2,
	322, 324, 7, 15, 2, 2, 323, 322, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324,
	325, 3, 2, 2, 2, 325, 326, 7, 12, 2, 2, 326, 107, 3, 2, 2, 2, 327, 331,
	7, 37, 2, 2, 328, 330, 10, 6, 2, 2, 329, 328, 3, 2, 2, 2, 330, 333, 3,
	2, 2, 2, 331, 329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 334, 3, 2, 2,
	2, 333, 331, 3, 2, 2, 2, 334, 335, 8, 54, 2, 2, 335, 109, 3, 2, 2, 2, 336,
	338, 9, 7, 2, 2, 337, 336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 337,
	3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 342, 8, 55,
	2, 2, 342, 111, 3, 2, 2, 2, 343, 344, 7, 36, 2, 2, 344, 345, 3, 2, 2, 2,
	345, 346, 8, 56, 3, 2, 346, 347, 8, 56, 4, 2, 347, 113, 3, 2, 2, 2, 348,
	349, 10, 8, 2, 2, 349, 350, 3, 2, 2, 2, 350, 351, 8, 57, 3, 2, 351, 115,
	3, 2, 2, 2, 352, 358, 7, 94, 2, 2, 353, 359, 9, 9, 2, 2, 354, 355, 7, 122,
	2, 2, 355, 356, 5, 118, 59, 2, 356, 357, 5, 118, 59, 2, 357, 359, 3, 2,
	2, 2, 358, 353, 3, 2, 2, 2, 358, 354, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2,
	360, 361, 8, 58, 3, 2, 361, 117, 3, 2, 2, 2, 362, 363, 9, 10, 2, 2, 363,
	119, 3, 2, 2, 2, 364, 365, 7, 36, 2, 2, 365, 366, 3, 2, 2, 2, 366, 367,
	8, 60, 5, 2, 367, 121, 3, 2, 2, 2, 13, 2, 3, 294, 300, 302, 308, 317, 323,
	331, 339, 358, 6, 8, 2, 2, 5, 2, 2, 4, 3, 2, 4, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "QSTR",
}

var lexerLiteralNames = []string{
	"", "'and'", "'break'", "'catch'", "'continue'", "'defer'", "'else'", "'false'",
	"'for'", "'func'", "'if'", "'meta'", "'nil'", "'not'", "'or'", "'return'",
	"'throw'", "'true'", "'try'", "'var'", "'while'", "'yield'", "'='", "'=='",
	"'!='", "'<'", "'<='", "'>'", "'>='", "'+'", "'-'", "'*'", "'/'", "'%'",
	"'?'", "';'", "','", "':'", "'.'", "'('", "')'", "'['", "']'", "'{'", "'}'",
	"'->'", "'&'", "'|'", "'...'",
}

var lexerSymbolicNames = []string{
	"", "AND", "BREAK", "CATCH", "CONTINUE", "DEFER", "ELSE", "FALSE", "FOR",
	"FUNC", "IF", "META", "NIL", "NOT", "OR", "RETURN", "THROW", "TRUE", "TRY",
	"VAR", "WHILE", "YIELD", "ASSIGN", "EQ", "NE", "LT", "LE", "GT", "GE",
	"ADD", "SUB", "MUL", "DIV", "MOD", "QUESTION_MARK", "SEMICOLON", "COMMA",
	"COLON", "PERIOD", "OPAREN", "CPAREN", "OBRACKET", "CBRACKET", "OCURLY",
	"CCURLY", "ARROW", "LAMBDA", "PIPE", "EXPAND", "NUMBER", "ID", "REGEX",
	"NEWLINE", "COMMENT", "WS", "STRING", "LQUOTE",
}

var lexerRuleNames = []string{
	"AND", "BREAK", "CATCH", "CONTINUE", "DEFER", "ELSE", "FALSE", "FOR", "FUNC",
	"IF", "META", "NIL", "NOT", "OR", "RETURN", "THROW", "TRUE", "TRY", "VAR",
	"WHILE", "YIELD", "ASSIGN", "EQ", "NE", "LT", "LE", "GT", "GE", "ADD",
	"SUB", "MUL", "DIV", "MOD", "QUESTION_MARK", "SEMICOLON", "COMMA", "COLON",
	"PERIOD", "OPAREN", "CPAREN", "OBRACKET", "CBRACKET", "OCURLY", "CCURLY",
	"ARROW", "LAMBDA", "PIPE", "EXPAND", "NUMBER", "ID", "REGEX", "NEWLINE",
	"COMMENT", "WS", "LQUOTE", "QUOTED_TEXT", "QUOTED_ESCAPE", "HEX_DIGIT",
	"STRING",
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
	NitroLexerMETA          = 11
	NitroLexerNIL           = 12
	NitroLexerNOT           = 13
	NitroLexerOR            = 14
	NitroLexerRETURN        = 15
	NitroLexerTHROW         = 16
	NitroLexerTRUE          = 17
	NitroLexerTRY           = 18
	NitroLexerVAR           = 19
	NitroLexerWHILE         = 20
	NitroLexerYIELD         = 21
	NitroLexerASSIGN        = 22
	NitroLexerEQ            = 23
	NitroLexerNE            = 24
	NitroLexerLT            = 25
	NitroLexerLE            = 26
	NitroLexerGT            = 27
	NitroLexerGE            = 28
	NitroLexerADD           = 29
	NitroLexerSUB           = 30
	NitroLexerMUL           = 31
	NitroLexerDIV           = 32
	NitroLexerMOD           = 33
	NitroLexerQUESTION_MARK = 34
	NitroLexerSEMICOLON     = 35
	NitroLexerCOMMA         = 36
	NitroLexerCOLON         = 37
	NitroLexerPERIOD        = 38
	NitroLexerOPAREN        = 39
	NitroLexerCPAREN        = 40
	NitroLexerOBRACKET      = 41
	NitroLexerCBRACKET      = 42
	NitroLexerOCURLY        = 43
	NitroLexerCCURLY        = 44
	NitroLexerARROW         = 45
	NitroLexerLAMBDA        = 46
	NitroLexerPIPE          = 47
	NitroLexerEXPAND        = 48
	NitroLexerNUMBER        = 49
	NitroLexerID            = 50
	NitroLexerREGEX         = 51
	NitroLexerNEWLINE       = 52
	NitroLexerCOMMENT       = 53
	NitroLexerWS            = 54
	NitroLexerSTRING        = 55
	NitroLexerLQUOTE        = 56
)

// NitroLexerQSTR is the NitroLexer mode.
const NitroLexerQSTR = 1
