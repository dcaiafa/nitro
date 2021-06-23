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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 60, 606,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 117, 10, 3, 12, 3, 14,
	3, 120, 11, 3, 3, 3, 7, 3, 123, 10, 3, 12, 3, 14, 3, 126, 11, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 135, 10, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 141, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 7, 5, 148, 10, 5,
	12, 5, 14, 5, 151, 11, 5, 3, 5, 5, 5, 154, 10, 5, 3, 6, 3, 6, 3, 6, 5,
	6, 159, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 165, 10, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 5, 9, 171, 10, 9, 3, 9, 7, 9, 174, 10, 9, 12, 9, 14, 9, 177, 11,
	9, 3, 10, 3, 10, 6, 10, 181, 10, 10, 13, 10, 14, 10, 182, 3, 10, 7, 10,
	186, 10, 10, 12, 10, 14, 10, 189, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11,
	205, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 7, 13, 214,
	10, 13, 12, 13, 14, 13, 217, 11, 13, 3, 14, 3, 14, 3, 14, 7, 14, 222, 10,
	14, 12, 14, 14, 14, 225, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 231,
	10, 15, 3, 16, 3, 16, 3, 16, 7, 16, 236, 10, 16, 12, 16, 14, 16, 239, 11,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 18, 7, 18, 252, 10, 18, 12, 18, 14, 18, 255, 11, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20,
	269, 10, 20, 12, 20, 14, 20, 272, 11, 20, 3, 20, 5, 20, 275, 10, 20, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 293, 10, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 7, 24, 303, 10, 24, 12, 24, 14,
	24, 306, 11, 24, 3, 25, 3, 25, 5, 25, 310, 10, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 5, 26, 318, 10, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 7, 32, 349, 10, 32, 12, 32, 14, 32, 352, 11, 32,
	3, 33, 3, 33, 5, 33, 356, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 7, 34, 376, 10, 34, 12, 34, 14, 34, 379, 11, 34, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 388, 10, 35, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 401,
	10, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 5, 36, 414, 10, 36, 3, 36, 3, 36, 5, 36, 418, 10, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 425, 10, 36, 5, 36, 427, 10, 36,
	3, 36, 7, 36, 430, 10, 36, 12, 36, 14, 36, 433, 11, 36, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 7, 38, 440, 10, 38, 12, 38, 14, 38, 443, 11, 38, 3, 38,
	5, 38, 446, 10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 5, 39, 458, 10, 39, 3, 40, 3, 40, 3, 40, 5, 40, 463,
	10, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 5, 41, 472, 10,
	41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43,
	7, 43, 484, 10, 43, 12, 43, 14, 43, 487, 11, 43, 3, 43, 5, 43, 490, 10,
	43, 5, 43, 492, 10, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 509, 10,
	44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 7, 45, 517, 10, 45, 12, 45,
	14, 45, 520, 11, 45, 3, 45, 5, 45, 523, 10, 45, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3,
	50, 3, 50, 3, 50, 7, 50, 552, 10, 50, 12, 50, 14, 50, 555, 11, 50, 3, 50,
	7, 50, 558, 10, 50, 12, 50, 14, 50, 561, 11, 50, 5, 50, 563, 10, 50, 3,
	51, 3, 51, 3, 51, 5, 51, 568, 10, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52,
	3, 52, 7, 52, 576, 10, 52, 12, 52, 14, 52, 579, 11, 52, 3, 52, 5, 52, 582,
	10, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54,
	3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3,
	55, 3, 56, 3, 56, 3, 56, 2, 5, 62, 66, 70, 57, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
	88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 2, 8, 3, 2, 38, 39,
	8, 2, 9, 9, 15, 15, 20, 20, 52, 52, 56, 56, 59, 59, 3, 2, 34, 36, 3, 2,
	32, 33, 3, 2, 26, 31, 4, 2, 3, 24, 53, 53, 2, 633, 2, 112, 3, 2, 2, 2,
	4, 118, 3, 2, 2, 2, 6, 129, 3, 2, 2, 2, 8, 144, 3, 2, 2, 2, 10, 155, 3,
	2, 2, 2, 12, 160, 3, 2, 2, 2, 14, 162, 3, 2, 2, 2, 16, 170, 3, 2, 2, 2,
	18, 178, 3, 2, 2, 2, 20, 204, 3, 2, 2, 2, 22, 206, 3, 2, 2, 2, 24, 210,
	3, 2, 2, 2, 26, 218, 3, 2, 2, 2, 28, 226, 3, 2, 2, 2, 30, 232, 3, 2, 2,
	2, 32, 240, 3, 2, 2, 2, 34, 248, 3, 2, 2, 2, 36, 256, 3, 2, 2, 2, 38, 262,
	3, 2, 2, 2, 40, 276, 3, 2, 2, 2, 42, 283, 3, 2, 2, 2, 44, 288, 3, 2, 2,
	2, 46, 299, 3, 2, 2, 2, 48, 307, 3, 2, 2, 2, 50, 311, 3, 2, 2, 2, 52, 323,
	3, 2, 2, 2, 54, 326, 3, 2, 2, 2, 56, 329, 3, 2, 2, 2, 58, 332, 3, 2, 2,
	2, 60, 334, 3, 2, 2, 2, 62, 336, 3, 2, 2, 2, 64, 355, 3, 2, 2, 2, 66, 357,
	3, 2, 2, 2, 68, 387, 3, 2, 2, 2, 70, 400, 3, 2, 2, 2, 72, 434, 3, 2, 2,
	2, 74, 436, 3, 2, 2, 2, 76, 457, 3, 2, 2, 2, 78, 459, 3, 2, 2, 2, 80, 469,
	3, 2, 2, 2, 82, 476, 3, 2, 2, 2, 84, 491, 3, 2, 2, 2, 86, 508, 3, 2, 2,
	2, 88, 510, 3, 2, 2, 2, 90, 524, 3, 2, 2, 2, 92, 531, 3, 2, 2, 2, 94, 536,
	3, 2, 2, 2, 96, 544, 3, 2, 2, 2, 98, 562, 3, 2, 2, 2, 100, 567, 3, 2, 2,
	2, 102, 569, 3, 2, 2, 2, 104, 583, 3, 2, 2, 2, 106, 590, 3, 2, 2, 2, 108,
	595, 3, 2, 2, 2, 110, 603, 3, 2, 2, 2, 112, 113, 5, 4, 3, 2, 113, 114,
	7, 2, 2, 3, 114, 3, 3, 2, 2, 2, 115, 117, 5, 6, 4, 2, 116, 115, 3, 2, 2,
	2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119,
	124, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 123, 5, 14, 8, 2, 122, 121,
	3, 2, 2, 2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2,
	2, 2, 125, 127, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 128, 5, 16, 9, 2,
	128, 5, 3, 2, 2, 2, 129, 130, 7, 14, 2, 2, 130, 131, 7, 53, 2, 2, 131,
	134, 7, 53, 2, 2, 132, 133, 7, 25, 2, 2, 133, 135, 5, 62, 32, 2, 134, 132,
	3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 140, 3, 2, 2, 2, 136, 137, 7, 46,
	2, 2, 137, 138, 5, 8, 5, 2, 138, 139, 7, 47, 2, 2, 139, 141, 3, 2, 2, 2,
	140, 136, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142,
	143, 7, 38, 2, 2, 143, 7, 3, 2, 2, 2, 144, 149, 5, 10, 6, 2, 145, 146,
	9, 2, 2, 2, 146, 148, 5, 10, 6, 2, 147, 145, 3, 2, 2, 2, 148, 151, 3, 2,
	2, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 153, 3, 2, 2, 2,
	151, 149, 3, 2, 2, 2, 152, 154, 9, 2, 2, 2, 153, 152, 3, 2, 2, 2, 153,
	154, 3, 2, 2, 2, 154, 9, 3, 2, 2, 2, 155, 158, 5, 110, 56, 2, 156, 157,
	7, 40, 2, 2, 157, 159, 5, 12, 7, 2, 158, 156, 3, 2, 2, 2, 158, 159, 3,
	2, 2, 2, 159, 11, 3, 2, 2, 2, 160, 161, 9, 3, 2, 2, 161, 13, 3, 2, 2, 2,
	162, 164, 7, 13, 2, 2, 163, 165, 7, 53, 2, 2, 164, 163, 3, 2, 2, 2, 164,
	165, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 167, 7, 59, 2, 2, 167, 168,
	7, 38, 2, 2, 168, 15, 3, 2, 2, 2, 169, 171, 5, 18, 10, 2, 170, 169, 3,
	2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 175, 3, 2, 2, 2, 172, 174, 7, 38, 2,
	2, 173, 172, 3, 2, 2, 2, 174, 177, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175,
	176, 3, 2, 2, 2, 176, 17, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 178, 187, 5,
	20, 11, 2, 179, 181, 7, 38, 2, 2, 180, 179, 3, 2, 2, 2, 181, 182, 3, 2,
	2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2,
	184, 186, 5, 20, 11, 2, 185, 180, 3, 2, 2, 2, 186, 189, 3, 2, 2, 2, 187,
	185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 19, 3, 2, 2, 2, 189, 187, 3,
	2, 2, 2, 190, 205, 5, 22, 12, 2, 191, 205, 5, 28, 15, 2, 192, 205, 5, 32,
	17, 2, 193, 205, 5, 36, 19, 2, 194, 205, 5, 38, 20, 2, 195, 205, 5, 44,
	23, 2, 196, 205, 5, 48, 25, 2, 197, 205, 5, 62, 32, 2, 198, 205, 5, 50,
	26, 2, 199, 205, 5, 52, 27, 2, 200, 205, 5, 54, 28, 2, 201, 205, 5, 56,
	29, 2, 202, 205, 5, 58, 30, 2, 203, 205, 5, 60, 31, 2, 204, 190, 3, 2,
	2, 2, 204, 191, 3, 2, 2, 2, 204, 192, 3, 2, 2, 2, 204, 193, 3, 2, 2, 2,
	204, 194, 3, 2, 2, 2, 204, 195, 3, 2, 2, 2, 204, 196, 3, 2, 2, 2, 204,
	197, 3, 2, 2, 2, 204, 198, 3, 2, 2, 2, 204, 199, 3, 2, 2, 2, 204, 200,
	3, 2, 2, 2, 204, 201, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204, 203, 3, 2,
	2, 2, 205, 21, 3, 2, 2, 2, 206, 207, 5, 24, 13, 2, 207, 208, 7, 25, 2,
	2, 208, 209, 5, 26, 14, 2, 209, 23, 3, 2, 2, 2, 210, 215, 5, 76, 39, 2,
	211, 212, 7, 39, 2, 2, 212, 214, 5, 76, 39, 2, 213, 211, 3, 2, 2, 2, 214,
	217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 25, 3,
	2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 223, 5, 62, 32, 2, 219, 220, 7, 39,
	2, 2, 220, 222, 5, 62, 32, 2, 221, 219, 3, 2, 2, 2, 222, 225, 3, 2, 2,
	2, 223, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 27, 3, 2, 2, 2, 225,
	223, 3, 2, 2, 2, 226, 227, 7, 22, 2, 2, 227, 230, 5, 30, 16, 2, 228, 229,
	7, 25, 2, 2, 229, 231, 5, 26, 14, 2, 230, 228, 3, 2, 2, 2, 230, 231, 3,
	2, 2, 2, 231, 29, 3, 2, 2, 2, 232, 237, 7, 53, 2, 2, 233, 234, 7, 39, 2,
	2, 234, 236, 7, 53, 2, 2, 235, 233, 3, 2, 2, 2, 236, 239, 3, 2, 2, 2, 237,
	235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 31, 3, 2, 2, 2, 239, 237, 3,
	2, 2, 2, 240, 241, 7, 10, 2, 2, 241, 242, 5, 34, 18, 2, 242, 243, 7, 53,
	2, 2, 243, 244, 5, 62, 32, 2, 244, 245, 7, 46, 2, 2, 245, 246, 5, 16, 9,
	2, 246, 247, 7, 47, 2, 2, 247, 33, 3, 2, 2, 2, 248, 253, 7, 53, 2, 2, 249,
	250, 7, 39, 2, 2, 250, 252, 7, 53, 2, 2, 251, 249, 3, 2, 2, 2, 252, 255,
	3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 35, 3, 2,
	2, 2, 255, 253, 3, 2, 2, 2, 256, 257, 7, 23, 2, 2, 257, 258, 5, 62, 32,
	2, 258, 259, 7, 46, 2, 2, 259, 260, 5, 16, 9, 2, 260, 261, 7, 47, 2, 2,
	261, 37, 3, 2, 2, 2, 262, 263, 7, 12, 2, 2, 263, 264, 5, 62, 32, 2, 264,
	265, 7, 46, 2, 2, 265, 266, 5, 16, 9, 2, 266, 270, 7, 47, 2, 2, 267, 269,
	5, 40, 21, 2, 268, 267, 3, 2, 2, 2, 269, 272, 3, 2, 2, 2, 270, 268, 3,
	2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2,
	2, 273, 275, 5, 42, 22, 2, 274, 273, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2,
	275, 39, 3, 2, 2, 2, 276, 277, 7, 8, 2, 2, 277, 278, 7, 12, 2, 2, 278,
	279, 5, 62, 32, 2, 279, 280, 7, 46, 2, 2, 280, 281, 5, 16, 9, 2, 281, 282,
	7, 47, 2, 2, 282, 41, 3, 2, 2, 2, 283, 284, 7, 8, 2, 2, 284, 285, 7, 46,
	2, 2, 285, 286, 5, 16, 9, 2, 286, 287, 7, 47, 2, 2, 287, 43, 3, 2, 2, 2,
	288, 289, 7, 11, 2, 2, 289, 290, 7, 53, 2, 2, 290, 292, 7, 42, 2, 2, 291,
	293, 5, 46, 24, 2, 292, 291, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 294,
	3, 2, 2, 2, 294, 295, 7, 43, 2, 2, 295, 296, 7, 46, 2, 2, 296, 297, 5,
	16, 9, 2, 297, 298, 7, 47, 2, 2, 298, 45, 3, 2, 2, 2, 299, 304, 7, 53,
	2, 2, 300, 301, 7, 39, 2, 2, 301, 303, 7, 53, 2, 2, 302, 300, 3, 2, 2,
	2, 303, 306, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305,
	47, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 307, 309, 7, 18, 2, 2, 308, 310,
	5, 26, 14, 2, 309, 308, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 49, 3, 2,
	2, 2, 311, 312, 7, 21, 2, 2, 312, 313, 7, 46, 2, 2, 313, 314, 5, 16, 9,
	2, 314, 315, 7, 47, 2, 2, 315, 317, 7, 5, 2, 2, 316, 318, 7, 53, 2, 2,
	317, 316, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319,
	320, 7, 46, 2, 2, 320, 321, 5, 16, 9, 2, 321, 322, 7, 47, 2, 2, 322, 51,
	3, 2, 2, 2, 323, 324, 7, 19, 2, 2, 324, 325, 5, 62, 32, 2, 325, 53, 3,
	2, 2, 2, 326, 327, 7, 7, 2, 2, 327, 328, 5, 70, 36, 2, 328, 55, 3, 2, 2,
	2, 329, 330, 7, 24, 2, 2, 330, 331, 5, 26, 14, 2, 331, 57, 3, 2, 2, 2,
	332, 333, 7, 4, 2, 2, 333, 59, 3, 2, 2, 2, 334, 335, 7, 6, 2, 2, 335, 61,
	3, 2, 2, 2, 336, 337, 8, 32, 1, 2, 337, 338, 5, 64, 33, 2, 338, 350, 3,
	2, 2, 2, 339, 340, 12, 5, 2, 2, 340, 341, 7, 37, 2, 2, 341, 342, 5, 62,
	32, 2, 342, 343, 7, 40, 2, 2, 343, 344, 5, 62, 32, 5, 344, 349, 3, 2, 2,
	2, 345, 346, 12, 4, 2, 2, 346, 347, 7, 50, 2, 2, 347, 349, 5, 62, 32, 5,
	348, 339, 3, 2, 2, 2, 348, 345, 3, 2, 2, 2, 349, 352, 3, 2, 2, 2, 350,
	348, 3, 2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 63, 3, 2, 2, 2, 352, 350, 3,
	2, 2, 2, 353, 356, 5, 66, 34, 2, 354, 356, 5, 80, 41, 2, 355, 353, 3, 2,
	2, 2, 355, 354, 3, 2, 2, 2, 356, 65, 3, 2, 2, 2, 357, 358, 8, 34, 1, 2,
	358, 359, 5, 68, 35, 2, 359, 377, 3, 2, 2, 2, 360, 361, 12, 7, 2, 2, 361,
	362, 9, 4, 2, 2, 362, 376, 5, 66, 34, 8, 363, 364, 12, 6, 2, 2, 364, 365,
	9, 5, 2, 2, 365, 376, 5, 66, 34, 7, 366, 367, 12, 5, 2, 2, 367, 368, 9,
	6, 2, 2, 368, 376, 5, 66, 34, 6, 369, 370, 12, 4, 2, 2, 370, 371, 7, 3,
	2, 2, 371, 376, 5, 66, 34, 5, 372, 373, 12, 3, 2, 2, 373, 374, 7, 17, 2,
	2, 374, 376, 5, 66, 34, 4, 375, 360, 3, 2, 2, 2, 375, 363, 3, 2, 2, 2,
	375, 366, 3, 2, 2, 2, 375, 369, 3, 2, 2, 2, 375, 372, 3, 2, 2, 2, 376,
	379, 3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 67, 3,
	2, 2, 2, 379, 377, 3, 2, 2, 2, 380, 381, 7, 16, 2, 2, 381, 388, 5, 68,
	35, 2, 382, 383, 7, 32, 2, 2, 383, 388, 5, 68, 35, 2, 384, 385, 7, 33,
	2, 2, 385, 388, 5, 68, 35, 2, 386, 388, 5, 70, 36, 2, 387, 380, 3, 2, 2,
	2, 387, 382, 3, 2, 2, 2, 387, 384, 3, 2, 2, 2, 387, 386, 3, 2, 2, 2, 388,
	69, 3, 2, 2, 2, 389, 390, 8, 36, 1, 2, 390, 401, 7, 53, 2, 2, 391, 401,
	5, 78, 40, 2, 392, 401, 5, 82, 42, 2, 393, 401, 5, 96, 49, 2, 394, 401,
	5, 72, 37, 2, 395, 401, 7, 54, 2, 2, 396, 397, 7, 42, 2, 2, 397, 398, 5,
	62, 32, 2, 398, 399, 7, 43, 2, 2, 399, 401, 3, 2, 2, 2, 400, 389, 3, 2,
	2, 2, 400, 391, 3, 2, 2, 2, 400, 392, 3, 2, 2, 2, 400, 393, 3, 2, 2, 2,
	400, 394, 3, 2, 2, 2, 400, 395, 3, 2, 2, 2, 400, 396, 3, 2, 2, 2, 401,
	431, 3, 2, 2, 2, 402, 403, 12, 12, 2, 2, 403, 404, 7, 41, 2, 2, 404, 430,
	7, 53, 2, 2, 405, 406, 12, 11, 2, 2, 406, 407, 7, 44, 2, 2, 407, 408, 5,
	62, 32, 2, 408, 409, 7, 45, 2, 2, 409, 430, 3, 2, 2, 2, 410, 411, 12, 10,
	2, 2, 411, 413, 7, 44, 2, 2, 412, 414, 5, 62, 32, 2, 413, 412, 3, 2, 2,
	2, 413, 414, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 417, 7, 40, 2, 2, 416,
	418, 5, 62, 32, 2, 417, 416, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 419,
	3, 2, 2, 2, 419, 430, 7, 45, 2, 2, 420, 421, 12, 9, 2, 2, 421, 426, 7,
	42, 2, 2, 422, 424, 5, 74, 38, 2, 423, 425, 7, 51, 2, 2, 424, 423, 3, 2,
	2, 2, 424, 425, 3, 2, 2, 2, 425, 427, 3, 2, 2, 2, 426, 422, 3, 2, 2, 2,
	426, 427, 3, 2, 2, 2, 427, 428, 3, 2, 2, 2, 428, 430, 7, 43, 2, 2, 429,
	402, 3, 2, 2, 2, 429, 405, 3, 2, 2, 2, 429, 410, 3, 2, 2, 2, 429, 420,
	3, 2, 2, 2, 430, 433, 3, 2, 2, 2, 431, 429, 3, 2, 2, 2, 431, 432, 3, 2,
	2, 2, 432, 71, 3, 2, 2, 2, 433, 431, 3, 2, 2, 2, 434, 435, 9, 3, 2, 2,
	435, 73, 3, 2, 2, 2, 436, 441, 5, 62, 32, 2, 437, 438, 7, 39, 2, 2, 438,
	440, 5, 62, 32, 2, 439, 437, 3, 2, 2, 2, 440, 443, 3, 2, 2, 2, 441, 439,
	3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 445, 3, 2, 2, 2, 443, 441, 3, 2,
	2, 2, 444, 446, 9, 2, 2, 2, 445, 444, 3, 2, 2, 2, 445, 446, 3, 2, 2, 2,
	446, 75, 3, 2, 2, 2, 447, 458, 7, 53, 2, 2, 448, 449, 5, 70, 36, 2, 449,
	450, 7, 41, 2, 2, 450, 451, 7, 53, 2, 2, 451, 458, 3, 2, 2, 2, 452, 453,
	5, 70, 36, 2, 453, 454, 7, 44, 2, 2, 454, 455, 5, 62, 32, 2, 455, 456,
	7, 45, 2, 2, 456, 458, 3, 2, 2, 2, 457, 447, 3, 2, 2, 2, 457, 448, 3, 2,
	2, 2, 457, 452, 3, 2, 2, 2, 458, 77, 3, 2, 2, 2, 459, 460, 7, 11, 2, 2,
	460, 462, 7, 42, 2, 2, 461, 463, 5, 46, 24, 2, 462, 461, 3, 2, 2, 2, 462,
	463, 3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 464, 465, 7, 43, 2, 2, 465, 466,
	7, 46, 2, 2, 466, 467, 5, 16, 9, 2, 467, 468, 7, 47, 2, 2, 468, 79, 3,
	2, 2, 2, 469, 471, 7, 49, 2, 2, 470, 472, 5, 46, 24, 2, 471, 470, 3, 2,
	2, 2, 471, 472, 3, 2, 2, 2, 472, 473, 3, 2, 2, 2, 473, 474, 7, 48, 2, 2,
	474, 475, 5, 66, 34, 2, 475, 81, 3, 2, 2, 2, 476, 477, 7, 46, 2, 2, 477,
	478, 5, 84, 43, 2, 478, 479, 7, 47, 2, 2, 479, 83, 3, 2, 2, 2, 480, 485,
	5, 86, 44, 2, 481, 482, 9, 2, 2, 2, 482, 484, 5, 86, 44, 2, 483, 481, 3,
	2, 2, 2, 484, 487, 3, 2, 2, 2, 485, 483, 3, 2, 2, 2, 485, 486, 3, 2, 2,
	2, 486, 489, 3, 2, 2, 2, 487, 485, 3, 2, 2, 2, 488, 490, 9, 2, 2, 2, 489,
	488, 3, 2, 2, 2, 489, 490, 3, 2, 2, 2, 490, 492, 3, 2, 2, 2, 491, 480,
	3, 2, 2, 2, 491, 492, 3, 2, 2, 2, 492, 85, 3, 2, 2, 2, 493, 494, 5, 110,
	56, 2, 494, 495, 7, 40, 2, 2, 495, 496, 5, 62, 32, 2, 496, 509, 3, 2, 2,
	2, 497, 498, 7, 44, 2, 2, 498, 499, 5, 62, 32, 2, 499, 500, 7, 45, 2, 2,
	500, 501, 7, 40, 2, 2, 501, 502, 5, 62, 32, 2, 502, 509, 3, 2, 2, 2, 503,
	504, 5, 70, 36, 2, 504, 505, 7, 51, 2, 2, 505, 509, 3, 2, 2, 2, 506, 509,
	5, 88, 45, 2, 507, 509, 5, 94, 48, 2, 508, 493, 3, 2, 2, 2, 508, 497, 3,
	2, 2, 2, 508, 503, 3, 2, 2, 2, 508, 506, 3, 2, 2, 2, 508, 507, 3, 2, 2,
	2, 509, 87, 3, 2, 2, 2, 510, 511, 7, 12, 2, 2, 511, 512, 5, 62, 32, 2,
	512, 513, 7, 46, 2, 2, 513, 514, 5, 84, 43, 2, 514, 518, 7, 47, 2, 2, 515,
	517, 5, 90, 46, 2, 516, 515, 3, 2, 2, 2, 517, 520, 3, 2, 2, 2, 518, 516,
	3, 2, 2, 2, 518, 519, 3, 2, 2, 2, 519, 522, 3, 2, 2, 2, 520, 518, 3, 2,
	2, 2, 521, 523, 5, 92, 47, 2, 522, 521, 3, 2, 2, 2, 522, 523, 3, 2, 2,
	2, 523, 89, 3, 2, 2, 2, 524, 525, 7, 8, 2, 2, 525, 526, 7, 12, 2, 2, 526,
	527, 5, 62, 32, 2, 527, 528, 7, 46, 2, 2, 528, 529, 5, 84, 43, 2, 529,
	530, 7, 47, 2, 2, 530, 91, 3, 2, 2, 2, 531, 532, 7, 8, 2, 2, 532, 533,
	7, 46, 2, 2, 533, 534, 5, 84, 43, 2, 534, 535, 7, 47, 2, 2, 535, 93, 3,
	2, 2, 2, 536, 537, 7, 10, 2, 2, 537, 538, 5, 34, 18, 2, 538, 539, 7, 53,
	2, 2, 539, 540, 5, 62, 32, 2, 540, 541, 7, 46, 2, 2, 541, 542, 5, 84, 43,
	2, 542, 543, 7, 47, 2, 2, 543, 95, 3, 2, 2, 2, 544, 545, 7, 44, 2, 2, 545,
	546, 5, 98, 50, 2, 546, 547, 7, 45, 2, 2, 547, 97, 3, 2, 2, 2, 548, 553,
	5, 100, 51, 2, 549, 550, 9, 2, 2, 2, 550, 552, 5, 100, 51, 2, 551, 549,
	3, 2, 2, 2, 552, 555, 3, 2, 2, 2, 553, 551, 3, 2, 2, 2, 553, 554, 3, 2,
	2, 2, 554, 559, 3, 2, 2, 2, 555, 553, 3, 2, 2, 2, 556, 558, 9, 2, 2, 2,
	557, 556, 3, 2, 2, 2, 558, 561, 3, 2, 2, 2, 559, 557, 3, 2, 2, 2, 559,
	560, 3, 2, 2, 2, 560, 563, 3, 2, 2, 2, 561, 559, 3, 2, 2, 2, 562, 548,
	3, 2, 2, 2, 562, 563, 3, 2, 2, 2, 563, 99, 3, 2, 2, 2, 564, 568, 5, 62,
	32, 2, 565, 568, 5, 102, 52, 2, 566, 568, 5, 108, 55, 2, 567, 564, 3, 2,
	2, 2, 567, 565, 3, 2, 2, 2, 567, 566, 3, 2, 2, 2, 568, 101, 3, 2, 2, 2,
	569, 570, 7, 12, 2, 2, 570, 571, 5, 62, 32, 2, 571, 572, 7, 46, 2, 2, 572,
	573, 5, 98, 50, 2, 573, 577, 7, 47, 2, 2, 574, 576, 5, 104, 53, 2, 575,
	574, 3, 2, 2, 2, 576, 579, 3, 2, 2, 2, 577, 575, 3, 2, 2, 2, 577, 578,
	3, 2, 2, 2, 578, 581, 3, 2, 2, 2, 579, 577, 3, 2, 2, 2, 580, 582, 5, 106,
	54, 2, 581, 580, 3, 2, 2, 2, 581, 582, 3, 2, 2, 2, 582, 103, 3, 2, 2, 2,
	583, 584, 7, 8, 2, 2, 584, 585, 7, 12, 2, 2, 585, 586, 5, 62, 32, 2, 586,
	587, 7, 46, 2, 2, 587, 588, 5, 98, 50, 2, 588, 589, 7, 47, 2, 2, 589, 105,
	3, 2, 2, 2, 590, 591, 7, 8, 2, 2, 591, 592, 7, 46, 2, 2, 592, 593, 5, 98,
	50, 2, 593, 594, 7, 47, 2, 2, 594, 107, 3, 2, 2, 2, 595, 596, 7, 10, 2,
	2, 596, 597, 5, 34, 18, 2, 597, 598, 7, 53, 2, 2, 598, 599, 5, 62, 32,
	2, 599, 600, 7, 46, 2, 2, 600, 601, 5, 98, 50, 2, 601, 602, 7, 47, 2, 2,
	602, 109, 3, 2, 2, 2, 603, 604, 9, 7, 2, 2, 604, 111, 3, 2, 2, 2, 56, 118,
	124, 134, 140, 149, 153, 158, 164, 170, 175, 182, 187, 204, 215, 223, 230,
	237, 253, 270, 274, 292, 304, 309, 317, 348, 350, 355, 375, 377, 387, 400,
	413, 417, 424, 426, 429, 431, 441, 445, 457, 462, 471, 485, 489, 491, 508,
	518, 522, 553, 559, 562, 567, 577, 581,
}
var literalNames = []string{
	"", "'and'", "'break'", "'catch'", "'continue'", "'defer'", "'else'", "'false'",
	"'for'", "'func'", "'if'", "'import'", "'meta'", "'nil'", "'not'", "'or'",
	"'return'", "'throw'", "'true'", "'try'", "'var'", "'while'", "'yield'",
	"'='", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'+'", "'-'", "'*'",
	"'/'", "'%'", "'?'", "';'", "','", "':'", "'.'", "'('", "')'", "'['", "']'",
	"'{'", "'}'", "'->'", "'&'", "'|'", "'...'",
}
var symbolicNames = []string{
	"", "AND", "BREAK", "CATCH", "CONTINUE", "DEFER", "ELSE", "FALSE", "FOR",
	"FUNC", "IF", "IMPORT", "META", "NIL", "NOT", "OR", "RETURN", "THROW",
	"TRUE", "TRY", "VAR", "WHILE", "YIELD", "ASSIGN", "EQ", "NE", "LT", "LE",
	"GT", "GE", "ADD", "SUB", "MUL", "DIV", "MOD", "QUESTION_MARK", "SEMICOLON",
	"COMMA", "COLON", "PERIOD", "OPAREN", "CPAREN", "OBRACKET", "CBRACKET",
	"OCURLY", "CCURLY", "ARROW", "LAMBDA", "PIPE", "EXPAND", "NUMBER", "ID",
	"REGEX", "NEWLINE", "CHAR", "COMMENT", "WS", "STRING", "LDQUOTE",
}

var ruleNames = []string{
	"start", "module", "meta_directive", "meta_attribs", "meta_attrib", "meta_literal",
	"import_stmt", "stmts", "stmt_list", "stmt", "assignment_stmt", "assignment_lvalues",
	"rvalues", "var_decl_stmt", "var_decl_vars", "for_stmt", "for_vars", "while_stmt",
	"if_stmt", "if_elif", "if_else", "func_stmt", "param_list", "return_stmt",
	"try_catch_stmt", "throw_stmt", "defer_stmt", "yield_stmt", "break_stmt",
	"continue_stmt", "expr", "pipeline_term_expr", "binary_expr", "unary_expr",
	"primary_expr", "simple_literal", "arg_list", "lvalue_expr", "lambda_expr",
	"short_lambda_expr", "object_literal", "object_fields", "object_field",
	"object_if", "object_elif", "object_else", "object_for", "array_literal",
	"array_elems", "array_elem", "array_if", "array_elif", "array_else", "array_for",
	"id_or_keyword",
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
	NitroParserEOF           = antlr.TokenEOF
	NitroParserAND           = 1
	NitroParserBREAK         = 2
	NitroParserCATCH         = 3
	NitroParserCONTINUE      = 4
	NitroParserDEFER         = 5
	NitroParserELSE          = 6
	NitroParserFALSE         = 7
	NitroParserFOR           = 8
	NitroParserFUNC          = 9
	NitroParserIF            = 10
	NitroParserIMPORT        = 11
	NitroParserMETA          = 12
	NitroParserNIL           = 13
	NitroParserNOT           = 14
	NitroParserOR            = 15
	NitroParserRETURN        = 16
	NitroParserTHROW         = 17
	NitroParserTRUE          = 18
	NitroParserTRY           = 19
	NitroParserVAR           = 20
	NitroParserWHILE         = 21
	NitroParserYIELD         = 22
	NitroParserASSIGN        = 23
	NitroParserEQ            = 24
	NitroParserNE            = 25
	NitroParserLT            = 26
	NitroParserLE            = 27
	NitroParserGT            = 28
	NitroParserGE            = 29
	NitroParserADD           = 30
	NitroParserSUB           = 31
	NitroParserMUL           = 32
	NitroParserDIV           = 33
	NitroParserMOD           = 34
	NitroParserQUESTION_MARK = 35
	NitroParserSEMICOLON     = 36
	NitroParserCOMMA         = 37
	NitroParserCOLON         = 38
	NitroParserPERIOD        = 39
	NitroParserOPAREN        = 40
	NitroParserCPAREN        = 41
	NitroParserOBRACKET      = 42
	NitroParserCBRACKET      = 43
	NitroParserOCURLY        = 44
	NitroParserCCURLY        = 45
	NitroParserARROW         = 46
	NitroParserLAMBDA        = 47
	NitroParserPIPE          = 48
	NitroParserEXPAND        = 49
	NitroParserNUMBER        = 50
	NitroParserID            = 51
	NitroParserREGEX         = 52
	NitroParserNEWLINE       = 53
	NitroParserCHAR          = 54
	NitroParserCOMMENT       = 55
	NitroParserWS            = 56
	NitroParserSTRING        = 57
	NitroParserLDQUOTE       = 58
)

// NitroParser rules.
const (
	NitroParserRULE_start              = 0
	NitroParserRULE_module             = 1
	NitroParserRULE_meta_directive     = 2
	NitroParserRULE_meta_attribs       = 3
	NitroParserRULE_meta_attrib        = 4
	NitroParserRULE_meta_literal       = 5
	NitroParserRULE_import_stmt        = 6
	NitroParserRULE_stmts              = 7
	NitroParserRULE_stmt_list          = 8
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
	NitroParserRULE_return_stmt        = 23
	NitroParserRULE_try_catch_stmt     = 24
	NitroParserRULE_throw_stmt         = 25
	NitroParserRULE_defer_stmt         = 26
	NitroParserRULE_yield_stmt         = 27
	NitroParserRULE_break_stmt         = 28
	NitroParserRULE_continue_stmt      = 29
	NitroParserRULE_expr               = 30
	NitroParserRULE_pipeline_term_expr = 31
	NitroParserRULE_binary_expr        = 32
	NitroParserRULE_unary_expr         = 33
	NitroParserRULE_primary_expr       = 34
	NitroParserRULE_simple_literal     = 35
	NitroParserRULE_arg_list           = 36
	NitroParserRULE_lvalue_expr        = 37
	NitroParserRULE_lambda_expr        = 38
	NitroParserRULE_short_lambda_expr  = 39
	NitroParserRULE_object_literal     = 40
	NitroParserRULE_object_fields      = 41
	NitroParserRULE_object_field       = 42
	NitroParserRULE_object_if          = 43
	NitroParserRULE_object_elif        = 44
	NitroParserRULE_object_else        = 45
	NitroParserRULE_object_for         = 46
	NitroParserRULE_array_literal      = 47
	NitroParserRULE_array_elems        = 48
	NitroParserRULE_array_elem         = 49
	NitroParserRULE_array_if           = 50
	NitroParserRULE_array_elif         = 51
	NitroParserRULE_array_else         = 52
	NitroParserRULE_array_for          = 53
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
		p.SetState(110)
		p.Module()
	}
	{
		p.SetState(111)
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

func (s *ModuleContext) AllMeta_directive() []IMeta_directiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMeta_directiveContext)(nil)).Elem())
	var tst = make([]IMeta_directiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMeta_directiveContext)
		}
	}

	return tst
}

func (s *ModuleContext) Meta_directive(i int) IMeta_directiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_directiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMeta_directiveContext)
}

func (s *ModuleContext) AllImport_stmt() []IImport_stmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImport_stmtContext)(nil)).Elem())
	var tst = make([]IImport_stmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImport_stmtContext)
		}
	}

	return tst
}

func (s *ModuleContext) Import_stmt(i int) IImport_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImport_stmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImport_stmtContext)
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
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserMETA {
		{
			p.SetState(113)
			p.Meta_directive()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserIMPORT {
		{
			p.SetState(119)
			p.Import_stmt()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.Stmts()
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

func (s *Meta_directiveContext) META() antlr.TerminalNode {
	return s.GetToken(NitroParserMETA, 0)
}

func (s *Meta_directiveContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(NitroParserID)
}

func (s *Meta_directiveContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserID, i)
}

func (s *Meta_directiveContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Meta_directiveContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(NitroParserASSIGN, 0)
}

func (s *Meta_directiveContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Meta_directiveContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Meta_directiveContext) Meta_attribs() IMeta_attribsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_attribsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeta_attribsContext)
}

func (s *Meta_directiveContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
	localctx = NewMeta_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NitroParserRULE_meta_directive)
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
		p.SetState(127)
		p.Match(NitroParserMETA)
	}
	{
		p.SetState(128)
		p.Match(NitroParserID)
	}
	{
		p.SetState(129)
		p.Match(NitroParserID)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserASSIGN {
		{
			p.SetState(130)
			p.Match(NitroParserASSIGN)
		}
		{
			p.SetState(131)
			p.expr(0)
		}

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMeta_attribContext)(nil)).Elem())
	var tst = make([]IMeta_attribContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMeta_attribContext)
		}
	}

	return tst
}

func (s *Meta_attribsContext) Meta_attrib(i int) IMeta_attribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_attribContext)(nil)).Elem(), i)

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
	localctx = NewMeta_attribsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NitroParserRULE_meta_attribs)
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
		p.SetState(142)
		p.Meta_attrib()
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(143)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(144)
				p.Meta_attrib()
			}

		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
		{
			p.SetState(150)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_or_keywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_or_keywordContext)
}

func (s *Meta_attribContext) COLON() antlr.TerminalNode {
	return s.GetToken(NitroParserCOLON, 0)
}

func (s *Meta_attribContext) Meta_literal() IMeta_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_literalContext)(nil)).Elem(), 0)

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
	localctx = NewMeta_attribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NitroParserRULE_meta_attrib)
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
		p.SetState(153)
		p.Id_or_keyword()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserCOLON {
		{
			p.SetState(154)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(155)
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
	localctx = NewMeta_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NitroParserRULE_meta_literal)
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
		p.SetState(158)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Meta_literalContext).val = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserNIL)|(1<<NitroParserTRUE))) != 0) || (((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(NitroParserNUMBER-50))|(1<<(NitroParserCHAR-50))|(1<<(NitroParserSTRING-50)))) != 0)) {
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
	localctx = NewImport_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NitroParserRULE_import_stmt)
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
		p.SetState(160)
		p.Match(NitroParserIMPORT)
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(161)
			p.Match(NitroParserID)
		}

	}
	{
		p.SetState(164)
		p.Match(NitroParserSTRING)
	}
	{
		p.SetState(165)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmt_listContext)(nil)).Elem(), 0)

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
	localctx = NewStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NitroParserRULE_stmts)
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
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserBREAK)|(1<<NitroParserCONTINUE)|(1<<NitroParserDEFER)|(1<<NitroParserFALSE)|(1<<NitroParserFOR)|(1<<NitroParserFUNC)|(1<<NitroParserIF)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserRETURN)|(1<<NitroParserTHROW)|(1<<NitroParserTRUE)|(1<<NitroParserTRY)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserYIELD)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(NitroParserOPAREN-40))|(1<<(NitroParserOBRACKET-40))|(1<<(NitroParserOCURLY-40))|(1<<(NitroParserLAMBDA-40))|(1<<(NitroParserNUMBER-40))|(1<<(NitroParserID-40))|(1<<(NitroParserREGEX-40))|(1<<(NitroParserCHAR-40))|(1<<(NitroParserSTRING-40)))) != 0) {
		{
			p.SetState(167)
			p.Stmt_list()
		}

	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserSEMICOLON {
		{
			p.SetState(170)
			p.Match(NitroParserSEMICOLON)
		}

		p.SetState(175)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *Stmt_listContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

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
	localctx = NewStmt_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NitroParserRULE_stmt_list)
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
		p.SetState(176)
		p.Stmt()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == NitroParserSEMICOLON {
				{
					p.SetState(177)
					p.Match(NitroParserSEMICOLON)
				}

				p.SetState(180)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(182)
				p.Stmt()
			}

		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreak_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignment_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinue_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThrow_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_decl_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefer_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITry_catch_stmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYield_stmtContext)(nil)).Elem(), 0)

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

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStmt_assignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Assignment_stmt()
		}

	case 2:
		localctx = NewStmt_var_decContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Var_decl_stmt()
		}

	case 3:
		localctx = NewStmt_forContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(190)
			p.For_stmt()
		}

	case 4:
		localctx = NewStmt_whileContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(191)
			p.While_stmt()
		}

	case 5:
		localctx = NewStmt_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(192)
			p.If_stmt()
		}

	case 6:
		localctx = NewStmt_funcContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(193)
			p.Func_stmt()
		}

	case 7:
		localctx = NewStmt_returnContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(194)
			p.Return_stmt()
		}

	case 8:
		localctx = NewStmt_exprContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(195)
			p.expr(0)
		}

	case 9:
		localctx = NewStmt_try_catchContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(196)
			p.Try_catch_stmt()
		}

	case 10:
		localctx = NewStmt_throwContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(197)
			p.Throw_stmt()
		}

	case 11:
		localctx = NewStmt_deferContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(198)
			p.Defer_stmt()
		}

	case 12:
		localctx = NewStmt_yieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(199)
			p.Yield_stmt()
		}

	case 13:
		localctx = NewStmt_breakContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(200)
			p.Break_stmt()
		}

	case 14:
		localctx = NewStmt_continueContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(201)
			p.Continue_stmt()
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
		p.SetState(204)
		p.Assignment_lvalues()
	}
	{
		p.SetState(205)
		p.Match(NitroParserASSIGN)
	}
	{
		p.SetState(206)
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
		p.SetState(208)
		p.Lvalue_expr()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(209)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(210)
			p.Lvalue_expr()
		}

		p.SetState(215)
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
		p.SetState(216)
		p.expr(0)
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(217)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(218)
			p.expr(0)
		}

		p.SetState(223)
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
		p.SetState(224)
		p.Match(NitroParserVAR)
	}
	{
		p.SetState(225)
		p.Var_decl_vars()
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserASSIGN {
		{
			p.SetState(226)
			p.Match(NitroParserASSIGN)
		}
		{
			p.SetState(227)
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
		p.SetState(230)
		p.Match(NitroParserID)
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(231)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(232)
			p.Match(NitroParserID)
		}

		p.SetState(237)
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

func (s *For_stmtContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *For_stmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *For_stmtContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *For_stmtContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

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
		p.SetState(238)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(239)
		p.For_vars()
	}
	{
		p.SetState(240)
		p.Match(NitroParserID)
	}
	{
		p.SetState(241)
		p.expr(0)
	}
	{
		p.SetState(242)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(243)
		p.Stmts()
	}
	{
		p.SetState(244)
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
		p.SetState(246)
		p.Match(NitroParserID)
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
			p.Match(NitroParserID)
		}

		p.SetState(253)
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

func (s *While_stmtContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *While_stmtContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

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
		p.SetState(254)
		p.Match(NitroParserWHILE)
	}
	{
		p.SetState(255)
		p.expr(0)
	}
	{
		p.SetState(256)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(257)
		p.Stmts()
	}
	{
		p.SetState(258)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_stmtContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *If_stmtContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *If_stmtContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(261)
		p.expr(0)
	}
	{
		p.SetState(262)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(263)
		p.Stmts()
	}
	{
		p.SetState(264)
		p.Match(NitroParserCCURLY)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(265)
				p.If_elif()
			}

		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(271)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_elifContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *If_elifContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

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
		p.SetState(274)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(275)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(276)
		p.expr(0)
	}
	{
		p.SetState(277)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(278)
		p.Stmts()
	}
	{
		p.SetState(279)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

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
		p.SetState(281)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(282)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(283)
		p.Stmts()
	}
	{
		p.SetState(284)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *Func_stmtContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
		p.SetState(286)
		p.Match(NitroParserFUNC)
	}
	{
		p.SetState(287)
		p.Match(NitroParserID)
	}
	{
		p.SetState(288)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(289)
			p.Param_list()
		}

	}
	{
		p.SetState(292)
		p.Match(NitroParserCPAREN)
	}
	{
		p.SetState(293)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(294)
		p.Stmts()
	}
	{
		p.SetState(295)
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
		p.SetState(297)
		p.Match(NitroParserID)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(298)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(299)
			p.Match(NitroParserID)
		}

		p.SetState(304)
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
	p.EnterRule(localctx, 46, NitroParserRULE_return_stmt)
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
		p.SetState(305)
		p.Match(NitroParserRETURN)
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(NitroParserOPAREN-40))|(1<<(NitroParserOBRACKET-40))|(1<<(NitroParserOCURLY-40))|(1<<(NitroParserLAMBDA-40))|(1<<(NitroParserNUMBER-40))|(1<<(NitroParserID-40))|(1<<(NitroParserREGEX-40))|(1<<(NitroParserCHAR-40))|(1<<(NitroParserSTRING-40)))) != 0) {
		{
			p.SetState(306)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtsContext)(nil)).Elem())
	var tst = make([]IStmtsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtsContext)
		}
	}

	return tst
}

func (s *Try_catch_stmtContext) Stmts(i int) IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), i)

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
	localctx = NewTry_catch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, NitroParserRULE_try_catch_stmt)
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
		p.SetState(309)
		p.Match(NitroParserTRY)
	}
	{
		p.SetState(310)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(311)
		p.Stmts()
	}
	{
		p.SetState(312)
		p.Match(NitroParserCCURLY)
	}
	{
		p.SetState(313)
		p.Match(NitroParserCATCH)
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(314)
			p.Match(NitroParserID)
		}

	}
	{
		p.SetState(317)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(318)
		p.Stmts()
	}
	{
		p.SetState(319)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

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
	localctx = NewThrow_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, NitroParserRULE_throw_stmt)

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
		p.SetState(321)
		p.Match(NitroParserTHROW)
	}
	{
		p.SetState(322)
		p.expr(0)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

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
	localctx = NewDefer_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, NitroParserRULE_defer_stmt)

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
		p.SetState(324)
		p.Match(NitroParserDEFER)
	}
	{
		p.SetState(325)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRvaluesContext)(nil)).Elem(), 0)

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
	localctx = NewYield_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, NitroParserRULE_yield_stmt)

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
		p.SetState(327)
		p.Match(NitroParserYIELD)
	}
	{
		p.SetState(328)
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
	localctx = NewBreak_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, NitroParserRULE_break_stmt)

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
		p.SetState(330)
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
	localctx = NewContinue_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, NitroParserRULE_continue_stmt)

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
		p.SetState(332)
		p.Match(NitroParserCONTINUE)
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

func (s *ExprContext) Pipeline_term_expr() IPipeline_term_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeline_term_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeline_term_exprContext)
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

func (s *ExprContext) QUESTION_MARK() antlr.TerminalNode {
	return s.GetToken(NitroParserQUESTION_MARK, 0)
}

func (s *ExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(NitroParserCOLON, 0)
}

func (s *ExprContext) PIPE() antlr.TerminalNode {
	return s.GetToken(NitroParserPIPE, 0)
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
	_startState := 60
	p.EnterRecursionRule(localctx, 60, NitroParserRULE_expr, _p)

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
		p.SetState(335)
		p.Pipeline_term_expr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(346)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_expr)
				p.SetState(337)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(338)
					p.Match(NitroParserQUESTION_MARK)
				}
				{
					p.SetState(339)
					p.expr(0)
				}
				{
					p.SetState(340)
					p.Match(NitroParserCOLON)
				}
				{
					p.SetState(341)
					p.expr(3)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_expr)
				p.SetState(343)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(344)
					p.Match(NitroParserPIPE)
				}
				{
					p.SetState(345)
					p.expr(3)
				}

			}

		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IPipeline_term_exprContext is an interface to support dynamic dispatch.
type IPipeline_term_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeline_term_exprContext differentiates from other interfaces.
	IsPipeline_term_exprContext()
}

type Pipeline_term_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeline_term_exprContext() *Pipeline_term_exprContext {
	var p = new(Pipeline_term_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_pipeline_term_expr
	return p
}

func (*Pipeline_term_exprContext) IsPipeline_term_exprContext() {}

func NewPipeline_term_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pipeline_term_exprContext {
	var p = new(Pipeline_term_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_pipeline_term_expr

	return p
}

func (s *Pipeline_term_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Pipeline_term_exprContext) CopyFrom(ctx *Pipeline_term_exprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Pipeline_term_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pipeline_term_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Pipeline_term_expr_short_lambdaContext struct {
	*Pipeline_term_exprContext
}

func NewPipeline_term_expr_short_lambdaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Pipeline_term_expr_short_lambdaContext {
	var p = new(Pipeline_term_expr_short_lambdaContext)

	p.Pipeline_term_exprContext = NewEmptyPipeline_term_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Pipeline_term_exprContext))

	return p
}

func (s *Pipeline_term_expr_short_lambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pipeline_term_expr_short_lambdaContext) Short_lambda_expr() IShort_lambda_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShort_lambda_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShort_lambda_exprContext)
}

func (s *Pipeline_term_expr_short_lambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPipeline_term_expr_short_lambda(s)
	}
}

func (s *Pipeline_term_expr_short_lambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPipeline_term_expr_short_lambda(s)
	}
}

type Pipeline_term_expr_binaryContext struct {
	*Pipeline_term_exprContext
}

func NewPipeline_term_expr_binaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Pipeline_term_expr_binaryContext {
	var p = new(Pipeline_term_expr_binaryContext)

	p.Pipeline_term_exprContext = NewEmptyPipeline_term_exprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Pipeline_term_exprContext))

	return p
}

func (s *Pipeline_term_expr_binaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pipeline_term_expr_binaryContext) Binary_expr() IBinary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exprContext)
}

func (s *Pipeline_term_expr_binaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPipeline_term_expr_binary(s)
	}
}

func (s *Pipeline_term_expr_binaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPipeline_term_expr_binary(s)
	}
}

func (p *NitroParser) Pipeline_term_expr() (localctx IPipeline_term_exprContext) {
	localctx = NewPipeline_term_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, NitroParserRULE_pipeline_term_expr)

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

	p.SetState(353)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserFALSE, NitroParserFUNC, NitroParserNIL, NitroParserNOT, NitroParserTRUE, NitroParserADD, NitroParserSUB, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserNUMBER, NitroParserID, NitroParserREGEX, NitroParserCHAR, NitroParserSTRING:
		localctx = NewPipeline_term_expr_binaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.binary_expr(0)
		}

	case NitroParserLAMBDA:
		localctx = NewPipeline_term_expr_short_lambdaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(352)
			p.Short_lambda_expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnary_exprContext)
}

func (s *Binary_exprContext) AllBinary_expr() []IBinary_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBinary_exprContext)(nil)).Elem())
	var tst = make([]IBinary_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBinary_exprContext)
		}
	}

	return tst
}

func (s *Binary_exprContext) Binary_expr(i int) IBinary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exprContext)(nil)).Elem(), i)

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
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBinary_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBinary_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 64
	p.EnterRecursionRule(localctx, 64, NitroParserRULE_binary_expr, _p)
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
		p.SetState(356)
		p.Unary_expr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(373)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(358)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(359)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Binary_exprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(NitroParserMUL-32))|(1<<(NitroParserDIV-32))|(1<<(NitroParserMOD-32)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Binary_exprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(360)
					p.binary_expr(6)
				}

			case 2:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(361)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(362)

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
					p.SetState(363)
					p.binary_expr(5)
				}

			case 3:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(364)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(365)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Binary_exprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserEQ)|(1<<NitroParserNE)|(1<<NitroParserLT)|(1<<NitroParserLE)|(1<<NitroParserGT)|(1<<NitroParserGE))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Binary_exprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(366)
					p.binary_expr(4)
				}

			case 4:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(367)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(368)

					var _m = p.Match(NitroParserAND)

					localctx.(*Binary_exprContext).op = _m
				}
				{
					p.SetState(369)
					p.binary_expr(3)
				}

			case 5:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(370)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(371)

					var _m = p.Match(NitroParserOR)

					localctx.(*Binary_exprContext).op = _m
				}
				{
					p.SetState(372)
					p.binary_expr(2)
				}

			}

		}
		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 66, NitroParserRULE_unary_expr)

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

	p.SetState(385)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)

			var _m = p.Match(NitroParserNOT)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(379)
			p.Unary_expr()
		}

	case NitroParserADD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)

			var _m = p.Match(NitroParserADD)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(381)
			p.Unary_expr()
		}

	case NitroParserSUB:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(382)

			var _m = p.Match(NitroParserSUB)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(383)
			p.Unary_expr()
		}

	case NitroParserFALSE, NitroParserFUNC, NitroParserNIL, NitroParserTRUE, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserNUMBER, NitroParserID, NitroParserREGEX, NitroParserCHAR, NitroParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(384)
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
	_startState := 68
	p.EnterRecursionRule(localctx, 68, NitroParserRULE_primary_expr, _p)
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
	p.SetState(398)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserID:
		localctx = NewPrimary_expr_simple_refContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(388)
			p.Match(NitroParserID)
		}

	case NitroParserFUNC:
		localctx = NewPrimary_expr_lambdaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(389)
			p.Lambda_expr()
		}

	case NitroParserOCURLY:
		localctx = NewPrimary_expr_objectContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(390)
			p.Object_literal()
		}

	case NitroParserOBRACKET:
		localctx = NewPrimary_expr_arrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(391)
			p.Array_literal()
		}

	case NitroParserFALSE, NitroParserNIL, NitroParserTRUE, NitroParserNUMBER, NitroParserCHAR, NitroParserSTRING:
		localctx = NewPrimary_expr_literalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(392)
			p.Simple_literal()
		}

	case NitroParserREGEX:
		localctx = NewPrimary_expr_regexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(393)
			p.Match(NitroParserREGEX)
		}

	case NitroParserOPAREN:
		localctx = NewPrimary_expr_parenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(394)
			p.Match(NitroParserOPAREN)
		}
		{
			p.SetState(395)
			p.expr(0)
		}
		{
			p.SetState(396)
			p.Match(NitroParserCPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(427)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimary_expr_member_accessContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(400)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(401)
					p.Match(NitroParserPERIOD)
				}
				{
					p.SetState(402)
					p.Match(NitroParserID)
				}

			case 2:
				localctx = NewPrimary_expr_indexContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(403)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(404)
					p.Match(NitroParserOBRACKET)
				}
				{
					p.SetState(405)
					p.expr(0)
				}
				{
					p.SetState(406)
					p.Match(NitroParserCBRACKET)
				}

			case 3:
				localctx = NewPrimary_expr_sliceContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(408)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(409)
					p.Match(NitroParserOBRACKET)
				}
				p.SetState(411)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(NitroParserOPAREN-40))|(1<<(NitroParserOBRACKET-40))|(1<<(NitroParserOCURLY-40))|(1<<(NitroParserLAMBDA-40))|(1<<(NitroParserNUMBER-40))|(1<<(NitroParserID-40))|(1<<(NitroParserREGEX-40))|(1<<(NitroParserCHAR-40))|(1<<(NitroParserSTRING-40)))) != 0) {
					{
						p.SetState(410)

						var _x = p.expr(0)

						localctx.(*Primary_expr_sliceContext).b = _x
					}

				}
				{
					p.SetState(413)
					p.Match(NitroParserCOLON)
				}
				p.SetState(415)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(NitroParserOPAREN-40))|(1<<(NitroParserOBRACKET-40))|(1<<(NitroParserOCURLY-40))|(1<<(NitroParserLAMBDA-40))|(1<<(NitroParserNUMBER-40))|(1<<(NitroParserID-40))|(1<<(NitroParserREGEX-40))|(1<<(NitroParserCHAR-40))|(1<<(NitroParserSTRING-40)))) != 0) {
					{
						p.SetState(414)

						var _x = p.expr(0)

						localctx.(*Primary_expr_sliceContext).e = _x
					}

				}
				{
					p.SetState(417)
					p.Match(NitroParserCBRACKET)
				}

			case 4:
				localctx = NewPrimary_expr_callContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(418)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(419)
					p.Match(NitroParserOPAREN)
				}
				p.SetState(424)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(NitroParserOPAREN-40))|(1<<(NitroParserOBRACKET-40))|(1<<(NitroParserOCURLY-40))|(1<<(NitroParserLAMBDA-40))|(1<<(NitroParserNUMBER-40))|(1<<(NitroParserID-40))|(1<<(NitroParserREGEX-40))|(1<<(NitroParserCHAR-40))|(1<<(NitroParserSTRING-40)))) != 0) {
					{
						p.SetState(420)
						p.Arg_list()
					}
					p.SetState(422)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if _la == NitroParserEXPAND {
						{
							p.SetState(421)
							p.Match(NitroParserEXPAND)
						}

					}

				}
				{
					p.SetState(426)
					p.Match(NitroParserCPAREN)
				}

			}

		}
		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
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
	localctx = NewSimple_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, NitroParserRULE_simple_literal)
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
		p.SetState(432)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Simple_literalContext).val = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserNIL)|(1<<NitroParserTRUE))) != 0) || (((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(NitroParserNUMBER-50))|(1<<(NitroParserCHAR-50))|(1<<(NitroParserSTRING-50)))) != 0)) {
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
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, NitroParserRULE_arg_list)
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
		p.SetState(434)
		p.expr(0)
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(435)
				p.Match(NitroParserCOMMA)
			}
			{
				p.SetState(436)
				p.expr(0)
			}

		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
		{
			p.SetState(442)
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
	p.EnterRule(localctx, 74, NitroParserRULE_lvalue_expr)

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

	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLvalue_expr_simple_refContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(445)
			p.Match(NitroParserID)
		}

	case 2:
		localctx = NewLvalue_expr_member_accessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(446)
			p.primary_expr(0)
		}
		{
			p.SetState(447)
			p.Match(NitroParserPERIOD)
		}
		{
			p.SetState(448)
			p.Match(NitroParserID)
		}

	case 3:
		localctx = NewLvalue_expr_indexContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(450)
			p.primary_expr(0)
		}
		{
			p.SetState(451)
			p.Match(NitroParserOBRACKET)
		}
		{
			p.SetState(452)
			p.expr(0)
		}
		{
			p.SetState(453)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *Lambda_exprContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
	p.EnterRule(localctx, 76, NitroParserRULE_lambda_expr)
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
		p.SetState(457)
		p.Match(NitroParserFUNC)
	}
	{
		p.SetState(458)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(459)
			p.Param_list()
		}

	}
	{
		p.SetState(462)
		p.Match(NitroParserCPAREN)
	}
	{
		p.SetState(463)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(464)
		p.Stmts()
	}
	{
		p.SetState(465)
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

func (s *Short_lambda_exprContext) Binary_expr() IBinary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exprContext)
}

func (s *Short_lambda_exprContext) Param_list() IParam_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParam_listContext)(nil)).Elem(), 0)

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
	localctx = NewShort_lambda_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, NitroParserRULE_short_lambda_expr)
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
		p.SetState(467)
		p.Match(NitroParserLAMBDA)
	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(468)
			p.Param_list()
		}

	}
	{
		p.SetState(471)
		p.Match(NitroParserARROW)
	}
	{
		p.SetState(472)
		p.binary_expr(0)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

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
	localctx = NewObject_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, NitroParserRULE_object_literal)

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
		p.SetState(474)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(475)
		p.Object_fields()
	}
	{
		p.SetState(476)
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
	p.EnterRule(localctx, 82, NitroParserRULE_object_fields)
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
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserAND)|(1<<NitroParserBREAK)|(1<<NitroParserCATCH)|(1<<NitroParserCONTINUE)|(1<<NitroParserDEFER)|(1<<NitroParserELSE)|(1<<NitroParserFALSE)|(1<<NitroParserFOR)|(1<<NitroParserFUNC)|(1<<NitroParserIF)|(1<<NitroParserIMPORT)|(1<<NitroParserMETA)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserOR)|(1<<NitroParserRETURN)|(1<<NitroParserTHROW)|(1<<NitroParserTRUE)|(1<<NitroParserTRY)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserYIELD))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(NitroParserOPAREN-40))|(1<<(NitroParserOBRACKET-40))|(1<<(NitroParserOCURLY-40))|(1<<(NitroParserNUMBER-40))|(1<<(NitroParserID-40))|(1<<(NitroParserREGEX-40))|(1<<(NitroParserCHAR-40))|(1<<(NitroParserSTRING-40)))) != 0) {
		{
			p.SetState(478)
			p.Object_field()
		}
		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(479)
					_la = p.GetTokenStream().LA(1)

					if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(480)
					p.Object_field()
				}

			}
			p.SetState(485)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
		}
		p.SetState(487)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
			{
				p.SetState(486)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 84, NitroParserRULE_object_field)

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

	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		localctx = NewObject_field_id_keyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(491)
			p.Id_or_keyword()
		}
		{
			p.SetState(492)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(493)
			p.expr(0)
		}

	case 2:
		localctx = NewObject_field_expr_keyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(495)
			p.Match(NitroParserOBRACKET)
		}
		{
			p.SetState(496)
			p.expr(0)
		}
		{
			p.SetState(497)
			p.Match(NitroParserCBRACKET)
		}
		{
			p.SetState(498)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(499)
			p.expr(0)
		}

	case 3:
		localctx = NewObject_field_expansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(501)
			p.primary_expr(0)
		}
		{
			p.SetState(502)
			p.Match(NitroParserEXPAND)
		}

	case 4:
		localctx = NewObject_field_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(504)
			p.Object_if()
		}

	case 5:
		localctx = NewObject_field_forContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(505)
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

func (s *Object_ifContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Object_ifContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_ifContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
	p.EnterRule(localctx, 86, NitroParserRULE_object_if)
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
		p.SetState(508)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(509)
		p.expr(0)
	}
	{
		p.SetState(510)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(511)
		p.Object_fields()
	}
	{
		p.SetState(512)
		p.Match(NitroParserCCURLY)
	}
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(513)
				p.Object_elif()
			}

		}
		p.SetState(518)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}
	p.SetState(520)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(519)
			p.Object_else()
		}

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

func (s *Object_elifContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NitroParserELSE, 0)
}

func (s *Object_elifContext) IF() antlr.TerminalNode {
	return s.GetToken(NitroParserIF, 0)
}

func (s *Object_elifContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Object_elifContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Object_elifContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_elifContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
	p.EnterRule(localctx, 88, NitroParserRULE_object_elif)

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
		p.SetState(522)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(523)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(524)
		p.expr(0)
	}
	{
		p.SetState(525)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(526)
		p.Object_fields()
	}
	{
		p.SetState(527)
		p.Match(NitroParserCCURLY)
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

func (s *Object_elseContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Object_elseContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_elseContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
	p.EnterRule(localctx, 90, NitroParserRULE_object_else)

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
		p.SetState(529)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(530)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(531)
		p.Object_fields()
	}
	{
		p.SetState(532)
		p.Match(NitroParserCCURLY)
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

func (s *Object_forContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Object_forContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Object_forContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Object_forContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_forContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
	p.EnterRule(localctx, 92, NitroParserRULE_object_for)

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
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(535)
		p.For_vars()
	}
	{
		p.SetState(536)
		p.Match(NitroParserID)
	}
	{
		p.SetState(537)
		p.expr(0)
	}
	{
		p.SetState(538)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(539)
		p.Object_fields()
	}
	{
		p.SetState(540)
		p.Match(NitroParserCCURLY)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

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
	localctx = NewArray_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, NitroParserRULE_array_literal)

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
		p.SetState(542)
		p.Match(NitroParserOBRACKET)
	}
	{
		p.SetState(543)
		p.Array_elems()
	}
	{
		p.SetState(544)
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
	p.EnterRule(localctx, 96, NitroParserRULE_array_elems)
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
	p.SetState(560)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFOR)|(1<<NitroParserFUNC)|(1<<NitroParserIF)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(NitroParserOPAREN-40))|(1<<(NitroParserOBRACKET-40))|(1<<(NitroParserOCURLY-40))|(1<<(NitroParserLAMBDA-40))|(1<<(NitroParserNUMBER-40))|(1<<(NitroParserID-40))|(1<<(NitroParserREGEX-40))|(1<<(NitroParserCHAR-40))|(1<<(NitroParserSTRING-40)))) != 0) {
		{
			p.SetState(546)
			p.Array_elem()
		}
		p.SetState(551)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(547)
					_la = p.GetTokenStream().LA(1)

					if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(548)
					p.Array_elem()
				}

			}
			p.SetState(553)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
		}
		p.SetState(557)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
			{
				p.SetState(554)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(559)
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
	p.EnterRule(localctx, 98, NitroParserRULE_array_elem)

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

	p.SetState(565)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserFALSE, NitroParserFUNC, NitroParserNIL, NitroParserNOT, NitroParserTRUE, NitroParserADD, NitroParserSUB, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserLAMBDA, NitroParserNUMBER, NitroParserID, NitroParserREGEX, NitroParserCHAR, NitroParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(562)
			p.expr(0)
		}

	case NitroParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(563)
			p.Array_if()
		}

	case NitroParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(564)
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

func (s *Array_ifContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Array_ifContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_ifContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
	p.EnterRule(localctx, 100, NitroParserRULE_array_if)
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
		p.SetState(567)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(568)
		p.expr(0)
	}
	{
		p.SetState(569)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(570)
		p.Array_elems()
	}
	{
		p.SetState(571)
		p.Match(NitroParserCCURLY)
	}
	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(572)
				p.Array_elif()
			}

		}
		p.SetState(577)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())
	}
	p.SetState(579)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(578)
			p.Array_else()
		}

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

func (s *Array_elifContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NitroParserELSE, 0)
}

func (s *Array_elifContext) IF() antlr.TerminalNode {
	return s.GetToken(NitroParserIF, 0)
}

func (s *Array_elifContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Array_elifContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Array_elifContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_elifContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
	p.EnterRule(localctx, 102, NitroParserRULE_array_elif)

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
		p.SetState(581)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(582)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(583)
		p.expr(0)
	}
	{
		p.SetState(584)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(585)
		p.Array_elems()
	}
	{
		p.SetState(586)
		p.Match(NitroParserCCURLY)
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

func (s *Array_elseContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Array_elseContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_elseContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
	p.EnterRule(localctx, 104, NitroParserRULE_array_else)

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
		p.SetState(588)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(589)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(590)
		p.Array_elems()
	}
	{
		p.SetState(591)
		p.Match(NitroParserCCURLY)
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

func (s *Array_forContext) ID() antlr.TerminalNode {
	return s.GetToken(NitroParserID, 0)
}

func (s *Array_forContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Array_forContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Array_forContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_forContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserCCURLY, 0)
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
	p.EnterRule(localctx, 106, NitroParserRULE_array_for)

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
		p.SetState(593)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(594)
		p.For_vars()
	}
	{
		p.SetState(595)
		p.Match(NitroParserID)
	}
	{
		p.SetState(596)
		p.expr(0)
	}
	{
		p.SetState(597)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(598)
		p.Array_elems()
	}
	{
		p.SetState(599)
		p.Match(NitroParserCCURLY)
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

func (s *Id_or_keywordContext) META() antlr.TerminalNode {
	return s.GetToken(NitroParserMETA, 0)
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
		p.SetState(601)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Id_or_keywordContext).t = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserAND)|(1<<NitroParserBREAK)|(1<<NitroParserCATCH)|(1<<NitroParserCONTINUE)|(1<<NitroParserDEFER)|(1<<NitroParserELSE)|(1<<NitroParserFALSE)|(1<<NitroParserFOR)|(1<<NitroParserFUNC)|(1<<NitroParserIF)|(1<<NitroParserIMPORT)|(1<<NitroParserMETA)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserOR)|(1<<NitroParserRETURN)|(1<<NitroParserTHROW)|(1<<NitroParserTRUE)|(1<<NitroParserTRY)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserYIELD))) != 0) || _la == NitroParserID) {
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
	case 30:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 32:
		var t *Binary_exprContext = nil
		if localctx != nil {
			t = localctx.(*Binary_exprContext)
		}
		return p.Binary_expr_Sempred(t, predIndex)

	case 34:
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
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *NitroParser) Binary_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
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
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
