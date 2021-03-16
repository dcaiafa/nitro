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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 557,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 3, 2, 3, 2, 3, 2, 3, 3, 5, 3, 109,
	10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 115, 10, 4, 12, 4, 14, 4, 118, 11,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 127, 10, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 7, 7, 134, 10, 7, 12, 7, 14, 7, 137, 11, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 148, 10, 9, 3, 10,
	7, 10, 151, 10, 10, 12, 10, 14, 10, 154, 11, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 187, 10, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 7, 13, 196, 10, 13, 12, 13, 14,
	13, 199, 11, 13, 3, 14, 3, 14, 3, 14, 7, 14, 204, 10, 14, 12, 14, 14, 14,
	207, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 213, 10, 15, 3, 16, 3,
	16, 3, 16, 7, 16, 218, 10, 16, 12, 16, 14, 16, 221, 11, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 7, 18, 234,
	10, 18, 12, 18, 14, 18, 237, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 250, 10, 20, 12, 20, 14,
	20, 253, 11, 20, 3, 20, 5, 20, 256, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 5, 23, 273, 10, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24,
	7, 24, 282, 10, 24, 12, 24, 14, 24, 285, 11, 24, 3, 25, 3, 25, 5, 25, 289,
	10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 295, 10, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 5, 28, 307, 10, 28,
	3, 29, 5, 29, 310, 10, 29, 3, 29, 3, 29, 6, 29, 314, 10, 29, 13, 29, 14,
	29, 315, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 336,
	10, 30, 12, 30, 14, 30, 339, 11, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 5, 31, 348, 10, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 360, 10, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 373, 10,
	32, 3, 32, 3, 32, 5, 32, 377, 10, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32,
	383, 10, 32, 3, 32, 7, 32, 386, 10, 32, 12, 32, 14, 32, 389, 11, 32, 3,
	33, 3, 33, 3, 34, 3, 34, 3, 34, 7, 34, 396, 10, 34, 12, 34, 14, 34, 399,
	11, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 5, 35, 411, 10, 35, 3, 36, 3, 36, 3, 36, 5, 36, 416, 10, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 5, 37, 424, 10, 37, 3, 37, 3, 37,
	3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 7, 39, 436, 10,
	39, 12, 39, 14, 39, 439, 11, 39, 3, 39, 7, 39, 442, 10, 39, 12, 39, 14,
	39, 445, 11, 39, 5, 39, 447, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5,
	40, 464, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 7, 41, 471, 10, 41,
	12, 41, 14, 41, 474, 11, 41, 3, 41, 5, 41, 477, 10, 41, 3, 41, 3, 41, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	46, 3, 46, 3, 46, 7, 46, 505, 10, 46, 12, 46, 14, 46, 508, 11, 46, 3, 46,
	7, 46, 511, 10, 46, 12, 46, 14, 46, 514, 11, 46, 5, 46, 516, 10, 46, 3,
	47, 3, 47, 3, 47, 5, 47, 521, 10, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48,
	7, 48, 528, 10, 48, 12, 48, 14, 48, 531, 11, 48, 3, 48, 5, 48, 534, 10,
	48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50,
	3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3,
	52, 3, 52, 2, 4, 58, 62, 53, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96,
	98, 100, 102, 2, 8, 3, 2, 33, 34, 3, 2, 30, 32, 3, 2, 28, 29, 3, 2, 22,
	27, 6, 2, 6, 6, 15, 15, 47, 47, 51, 51, 4, 2, 3, 20, 48, 48, 2, 580, 2,
	104, 3, 2, 2, 2, 4, 108, 3, 2, 2, 2, 6, 112, 3, 2, 2, 2, 8, 121, 3, 2,
	2, 2, 10, 124, 3, 2, 2, 2, 12, 130, 3, 2, 2, 2, 14, 138, 3, 2, 2, 2, 16,
	147, 3, 2, 2, 2, 18, 152, 3, 2, 2, 2, 20, 186, 3, 2, 2, 2, 22, 188, 3,
	2, 2, 2, 24, 192, 3, 2, 2, 2, 26, 200, 3, 2, 2, 2, 28, 208, 3, 2, 2, 2,
	30, 214, 3, 2, 2, 2, 32, 222, 3, 2, 2, 2, 34, 230, 3, 2, 2, 2, 36, 238,
	3, 2, 2, 2, 38, 244, 3, 2, 2, 2, 40, 259, 3, 2, 2, 2, 42, 265, 3, 2, 2,
	2, 44, 268, 3, 2, 2, 2, 46, 278, 3, 2, 2, 2, 48, 286, 3, 2, 2, 2, 50, 290,
	3, 2, 2, 2, 52, 300, 3, 2, 2, 2, 54, 306, 3, 2, 2, 2, 56, 309, 3, 2, 2,
	2, 58, 317, 3, 2, 2, 2, 60, 347, 3, 2, 2, 2, 62, 359, 3, 2, 2, 2, 64, 390,
	3, 2, 2, 2, 66, 392, 3, 2, 2, 2, 68, 410, 3, 2, 2, 2, 70, 412, 3, 2, 2,
	2, 72, 421, 3, 2, 2, 2, 74, 428, 3, 2, 2, 2, 76, 446, 3, 2, 2, 2, 78, 463,
	3, 2, 2, 2, 80, 465, 3, 2, 2, 2, 82, 480, 3, 2, 2, 2, 84, 486, 3, 2, 2,
	2, 86, 489, 3, 2, 2, 2, 88, 497, 3, 2, 2, 2, 90, 515, 3, 2, 2, 2, 92, 520,
	3, 2, 2, 2, 94, 522, 3, 2, 2, 2, 96, 537, 3, 2, 2, 2, 98, 543, 3, 2, 2,
	2, 100, 546, 3, 2, 2, 2, 102, 554, 3, 2, 2, 2, 104, 105, 5, 4, 3, 2, 105,
	106, 7, 2, 2, 3, 106, 3, 3, 2, 2, 2, 107, 109, 5, 6, 4, 2, 108, 107, 3,
	2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111, 5, 18, 10,
	2, 111, 5, 3, 2, 2, 2, 112, 116, 7, 11, 2, 2, 113, 115, 5, 8, 5, 2, 114,
	113, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117,
	3, 2, 2, 2, 117, 119, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 120, 7, 5,
	2, 2, 120, 7, 3, 2, 2, 2, 121, 122, 7, 48, 2, 2, 122, 123, 5, 10, 6, 2,
	123, 9, 3, 2, 2, 2, 124, 126, 7, 41, 2, 2, 125, 127, 5, 12, 7, 2, 126,
	125, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129,
	7, 42, 2, 2, 129, 11, 3, 2, 2, 2, 130, 135, 5, 14, 8, 2, 131, 132, 9, 2,
	2, 2, 132, 134, 5, 14, 8, 2, 133, 131, 3, 2, 2, 2, 134, 137, 3, 2, 2, 2,
	135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 13, 3, 2, 2, 2, 137, 135,
	3, 2, 2, 2, 138, 139, 7, 48, 2, 2, 139, 140, 7, 35, 2, 2, 140, 141, 5,
	16, 9, 2, 141, 15, 3, 2, 2, 2, 142, 148, 7, 51, 2, 2, 143, 148, 7, 47,
	2, 2, 144, 148, 7, 15, 2, 2, 145, 148, 7, 6, 2, 2, 146, 148, 5, 10, 6,
	2, 147, 142, 3, 2, 2, 2, 147, 143, 3, 2, 2, 2, 147, 144, 3, 2, 2, 2, 147,
	145, 3, 2, 2, 2, 147, 146, 3, 2, 2, 2, 148, 17, 3, 2, 2, 2, 149, 151, 5,
	20, 11, 2, 150, 149, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2, 152, 150, 3, 2,
	2, 2, 152, 153, 3, 2, 2, 2, 153, 19, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2,
	155, 156, 5, 22, 12, 2, 156, 157, 7, 33, 2, 2, 157, 187, 3, 2, 2, 2, 158,
	159, 5, 28, 15, 2, 159, 160, 7, 33, 2, 2, 160, 187, 3, 2, 2, 2, 161, 162,
	5, 32, 17, 2, 162, 163, 7, 33, 2, 2, 163, 187, 3, 2, 2, 2, 164, 165, 5,
	36, 19, 2, 165, 166, 7, 33, 2, 2, 166, 187, 3, 2, 2, 2, 167, 168, 5, 38,
	20, 2, 168, 169, 7, 33, 2, 2, 169, 187, 3, 2, 2, 2, 170, 171, 5, 44, 23,
	2, 171, 172, 7, 33, 2, 2, 172, 187, 3, 2, 2, 2, 173, 174, 5, 48, 25, 2,
	174, 175, 7, 33, 2, 2, 175, 187, 3, 2, 2, 2, 176, 177, 5, 54, 28, 2, 177,
	178, 7, 33, 2, 2, 178, 187, 3, 2, 2, 2, 179, 180, 5, 50, 26, 2, 180, 181,
	7, 33, 2, 2, 181, 187, 3, 2, 2, 2, 182, 183, 5, 52, 27, 2, 183, 184, 7,
	33, 2, 2, 184, 187, 3, 2, 2, 2, 185, 187, 7, 33, 2, 2, 186, 155, 3, 2,
	2, 2, 186, 158, 3, 2, 2, 2, 186, 161, 3, 2, 2, 2, 186, 164, 3, 2, 2, 2,
	186, 167, 3, 2, 2, 2, 186, 170, 3, 2, 2, 2, 186, 173, 3, 2, 2, 2, 186,
	176, 3, 2, 2, 2, 186, 179, 3, 2, 2, 2, 186, 182, 3, 2, 2, 2, 186, 185,
	3, 2, 2, 2, 187, 21, 3, 2, 2, 2, 188, 189, 5, 24, 13, 2, 189, 190, 7, 21,
	2, 2, 190, 191, 5, 26, 14, 2, 191, 23, 3, 2, 2, 2, 192, 197, 5, 68, 35,
	2, 193, 194, 7, 34, 2, 2, 194, 196, 5, 68, 35, 2, 195, 193, 3, 2, 2, 2,
	196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198,
	25, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 205, 5, 54, 28, 2, 201, 202,
	7, 34, 2, 2, 202, 204, 5, 54, 28, 2, 203, 201, 3, 2, 2, 2, 204, 207, 3,
	2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 27, 3, 2, 2,
	2, 207, 205, 3, 2, 2, 2, 208, 209, 7, 16, 2, 2, 209, 212, 5, 30, 16, 2,
	210, 211, 7, 21, 2, 2, 211, 213, 5, 26, 14, 2, 212, 210, 3, 2, 2, 2, 212,
	213, 3, 2, 2, 2, 213, 29, 3, 2, 2, 2, 214, 219, 7, 48, 2, 2, 215, 216,
	7, 34, 2, 2, 216, 218, 7, 48, 2, 2, 217, 215, 3, 2, 2, 2, 218, 221, 3,
	2, 2, 2, 219, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 31, 3, 2, 2,
	2, 221, 219, 3, 2, 2, 2, 222, 223, 7, 8, 2, 2, 223, 224, 5, 34, 18, 2,
	224, 225, 7, 10, 2, 2, 225, 226, 5, 54, 28, 2, 226, 227, 7, 33, 2, 2, 227,
	228, 5, 18, 10, 2, 228, 229, 7, 5, 2, 2, 229, 33, 3, 2, 2, 2, 230, 235,
	7, 48, 2, 2, 231, 232, 7, 34, 2, 2, 232, 234, 7, 48, 2, 2, 233, 231, 3,
	2, 2, 2, 234, 237, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 235, 236, 3, 2, 2,
	2, 236, 35, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 238, 239, 7, 17, 2, 2, 239,
	240, 5, 54, 28, 2, 240, 241, 7, 33, 2, 2, 241, 242, 5, 18, 10, 2, 242,
	243, 7, 5, 2, 2, 243, 37, 3, 2, 2, 2, 244, 245, 7, 9, 2, 2, 245, 246, 5,
	54, 28, 2, 246, 247, 7, 33, 2, 2, 247, 251, 5, 18, 10, 2, 248, 250, 5,
	40, 21, 2, 249, 248, 3, 2, 2, 2, 250, 253, 3, 2, 2, 2, 251, 249, 3, 2,
	2, 2, 251, 252, 3, 2, 2, 2, 252, 255, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2,
	254, 256, 5, 42, 22, 2, 255, 254, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256,
	257, 3, 2, 2, 2, 257, 258, 7, 5, 2, 2, 258, 39, 3, 2, 2, 2, 259, 260, 7,
	4, 2, 2, 260, 261, 7, 9, 2, 2, 261, 262, 5, 54, 28, 2, 262, 263, 7, 33,
	2, 2, 263, 264, 5, 18, 10, 2, 264, 41, 3, 2, 2, 2, 265, 266, 7, 4, 2, 2,
	266, 267, 5, 18, 10, 2, 267, 43, 3, 2, 2, 2, 268, 269, 7, 7, 2, 2, 269,
	270, 7, 48, 2, 2, 270, 272, 7, 37, 2, 2, 271, 273, 5, 46, 24, 2, 272, 271,
	3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 275, 7, 38,
	2, 2, 275, 276, 5, 18, 10, 2, 276, 277, 7, 5, 2, 2, 277, 45, 3, 2, 2, 2,
	278, 283, 7, 48, 2, 2, 279, 280, 7, 34, 2, 2, 280, 282, 7, 48, 2, 2, 281,
	279, 3, 2, 2, 2, 282, 285, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 284,
	3, 2, 2, 2, 284, 47, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 286, 288, 7, 14,
	2, 2, 287, 289, 5, 26, 14, 2, 288, 287, 3, 2, 2, 2, 288, 289, 3, 2, 2,
	2, 289, 49, 3, 2, 2, 2, 290, 291, 7, 19, 2, 2, 291, 292, 5, 18, 10, 2,
	292, 294, 7, 20, 2, 2, 293, 295, 7, 48, 2, 2, 294, 293, 3, 2, 2, 2, 294,
	295, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 297, 7, 33, 2, 2, 297, 298,
	5, 18, 10, 2, 298, 299, 7, 5, 2, 2, 299, 51, 3, 2, 2, 2, 300, 301, 7, 18,
	2, 2, 301, 302, 5, 62, 32, 2, 302, 53, 3, 2, 2, 2, 303, 307, 5, 58, 30,
	2, 304, 307, 5, 72, 37, 2, 305, 307, 5, 56, 29, 2, 306, 303, 3, 2, 2, 2,
	306, 304, 3, 2, 2, 2, 306, 305, 3, 2, 2, 2, 307, 55, 3, 2, 2, 2, 308, 310,
	5, 58, 30, 2, 309, 308, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 313, 3,
	2, 2, 2, 311, 312, 7, 45, 2, 2, 312, 314, 5, 58, 30, 2, 313, 311, 3, 2,
	2, 2, 314, 315, 3, 2, 2, 2, 315, 313, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2,
	316, 57, 3, 2, 2, 2, 317, 318, 8, 30, 1, 2, 318, 319, 5, 60, 31, 2, 319,
	337, 3, 2, 2, 2, 320, 321, 12, 7, 2, 2, 321, 322, 9, 3, 2, 2, 322, 336,
	5, 58, 30, 8, 323, 324, 12, 6, 2, 2, 324, 325, 9, 4, 2, 2, 325, 336, 5,
	58, 30, 7, 326, 327, 12, 5, 2, 2, 327, 328, 9, 5, 2, 2, 328, 336, 5, 58,
	30, 6, 329, 330, 12, 4, 2, 2, 330, 331, 7, 3, 2, 2, 331, 336, 5, 58, 30,
	5, 332, 333, 12, 3, 2, 2, 333, 334, 7, 13, 2, 2, 334, 336, 5, 58, 30, 4,
	335, 320, 3, 2, 2, 2, 335, 323, 3, 2, 2, 2, 335, 326, 3, 2, 2, 2, 335,
	329, 3, 2, 2, 2, 335, 332, 3, 2, 2, 2, 336, 339, 3, 2, 2, 2, 337, 335,
	3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 59, 3, 2, 2, 2, 339, 337, 3, 2,
	2, 2, 340, 341, 7, 12, 2, 2, 341, 348, 5, 60, 31, 2, 342, 343, 7, 28, 2,
	2, 343, 348, 5, 60, 31, 2, 344, 345, 7, 29, 2, 2, 345, 348, 5, 60, 31,
	2, 346, 348, 5, 62, 32, 2, 347, 340, 3, 2, 2, 2, 347, 342, 3, 2, 2, 2,
	347, 344, 3, 2, 2, 2, 347, 346, 3, 2, 2, 2, 348, 61, 3, 2, 2, 2, 349, 350,
	8, 32, 1, 2, 350, 360, 7, 48, 2, 2, 351, 360, 5, 70, 36, 2, 352, 360, 5,
	74, 38, 2, 353, 360, 5, 88, 45, 2, 354, 360, 5, 64, 33, 2, 355, 356, 7,
	37, 2, 2, 356, 357, 5, 54, 28, 2, 357, 358, 7, 38, 2, 2, 358, 360, 3, 2,
	2, 2, 359, 349, 3, 2, 2, 2, 359, 351, 3, 2, 2, 2, 359, 352, 3, 2, 2, 2,
	359, 353, 3, 2, 2, 2, 359, 354, 3, 2, 2, 2, 359, 355, 3, 2, 2, 2, 360,
	387, 3, 2, 2, 2, 361, 362, 12, 11, 2, 2, 362, 363, 7, 36, 2, 2, 363, 386,
	7, 48, 2, 2, 364, 365, 12, 10, 2, 2, 365, 366, 7, 39, 2, 2, 366, 367, 5,
	54, 28, 2, 367, 368, 7, 40, 2, 2, 368, 386, 3, 2, 2, 2, 369, 370, 12, 9,
	2, 2, 370, 372, 7, 39, 2, 2, 371, 373, 5, 54, 28, 2, 372, 371, 3, 2, 2,
	2, 372, 373, 3, 2, 2, 2, 373, 374, 3, 2, 2, 2, 374, 376, 7, 35, 2, 2, 375,
	377, 5, 54, 28, 2, 376, 375, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 378,
	3, 2, 2, 2, 378, 386, 7, 40, 2, 2, 379, 380, 12, 8, 2, 2, 380, 382, 7,
	37, 2, 2, 381, 383, 5, 66, 34, 2, 382, 381, 3, 2, 2, 2, 382, 383, 3, 2,
	2, 2, 383, 384, 3, 2, 2, 2, 384, 386, 7, 38, 2, 2, 385, 361, 3, 2, 2, 2,
	385, 364, 3, 2, 2, 2, 385, 369, 3, 2, 2, 2, 385, 379, 3, 2, 2, 2, 386,
	389, 3, 2, 2, 2, 387, 385, 3, 2, 2, 2, 387, 388, 3, 2, 2, 2, 388, 63, 3,
	2, 2, 2, 389, 387, 3, 2, 2, 2, 390, 391, 9, 6, 2, 2, 391, 65, 3, 2, 2,
	2, 392, 397, 5, 54, 28, 2, 393, 394, 7, 34, 2, 2, 394, 396, 5, 54, 28,
	2, 395, 393, 3, 2, 2, 2, 396, 399, 3, 2, 2, 2, 397, 395, 3, 2, 2, 2, 397,
	398, 3, 2, 2, 2, 398, 67, 3, 2, 2, 2, 399, 397, 3, 2, 2, 2, 400, 411, 7,
	48, 2, 2, 401, 402, 5, 62, 32, 2, 402, 403, 7, 36, 2, 2, 403, 404, 7, 48,
	2, 2, 404, 411, 3, 2, 2, 2, 405, 406, 5, 62, 32, 2, 406, 407, 7, 39, 2,
	2, 407, 408, 5, 54, 28, 2, 408, 409, 7, 40, 2, 2, 409, 411, 3, 2, 2, 2,
	410, 400, 3, 2, 2, 2, 410, 401, 3, 2, 2, 2, 410, 405, 3, 2, 2, 2, 411,
	69, 3, 2, 2, 2, 412, 413, 7, 7, 2, 2, 413, 415, 7, 37, 2, 2, 414, 416,
	5, 46, 24, 2, 415, 414, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 417, 3,
	2, 2, 2, 417, 418, 7, 38, 2, 2, 418, 419, 5, 18, 10, 2, 419, 420, 7, 5,
	2, 2, 420, 71, 3, 2, 2, 2, 421, 423, 7, 44, 2, 2, 422, 424, 5, 46, 24,
	2, 423, 422, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425,
	426, 7, 43, 2, 2, 426, 427, 5, 58, 30, 2, 427, 73, 3, 2, 2, 2, 428, 429,
	7, 41, 2, 2, 429, 430, 5, 76, 39, 2, 430, 431, 7, 42, 2, 2, 431, 75, 3,
	2, 2, 2, 432, 437, 5, 78, 40, 2, 433, 434, 9, 2, 2, 2, 434, 436, 5, 78,
	40, 2, 435, 433, 3, 2, 2, 2, 436, 439, 3, 2, 2, 2, 437, 435, 3, 2, 2, 2,
	437, 438, 3, 2, 2, 2, 438, 443, 3, 2, 2, 2, 439, 437, 3, 2, 2, 2, 440,
	442, 9, 2, 2, 2, 441, 440, 3, 2, 2, 2, 442, 445, 3, 2, 2, 2, 443, 441,
	3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 447, 3, 2, 2, 2, 445, 443, 3, 2,
	2, 2, 446, 432, 3, 2, 2, 2, 446, 447, 3, 2, 2, 2, 447, 77, 3, 2, 2, 2,
	448, 449, 5, 102, 52, 2, 449, 450, 7, 35, 2, 2, 450, 451, 5, 54, 28, 2,
	451, 464, 3, 2, 2, 2, 452, 453, 7, 39, 2, 2, 453, 454, 5, 54, 28, 2, 454,
	455, 7, 40, 2, 2, 455, 456, 7, 35, 2, 2, 456, 457, 5, 54, 28, 2, 457, 464,
	3, 2, 2, 2, 458, 459, 5, 62, 32, 2, 459, 460, 7, 46, 2, 2, 460, 464, 3,
	2, 2, 2, 461, 464, 5, 80, 41, 2, 462, 464, 5, 86, 44, 2, 463, 448, 3, 2,
	2, 2, 463, 452, 3, 2, 2, 2, 463, 458, 3, 2, 2, 2, 463, 461, 3, 2, 2, 2,
	463, 462, 3, 2, 2, 2, 464, 79, 3, 2, 2, 2, 465, 466, 7, 9, 2, 2, 466, 467,
	5, 54, 28, 2, 467, 468, 7, 33, 2, 2, 468, 472, 5, 76, 39, 2, 469, 471,
	5, 82, 42, 2, 470, 469, 3, 2, 2, 2, 471, 474, 3, 2, 2, 2, 472, 470, 3,
	2, 2, 2, 472, 473, 3, 2, 2, 2, 473, 476, 3, 2, 2, 2, 474, 472, 3, 2, 2,
	2, 475, 477, 5, 84, 43, 2, 476, 475, 3, 2, 2, 2, 476, 477, 3, 2, 2, 2,
	477, 478, 3, 2, 2, 2, 478, 479, 7, 5, 2, 2, 479, 81, 3, 2, 2, 2, 480, 481,
	7, 4, 2, 2, 481, 482, 7, 9, 2, 2, 482, 483, 5, 54, 28, 2, 483, 484, 7,
	33, 2, 2, 484, 485, 5, 76, 39, 2, 485, 83, 3, 2, 2, 2, 486, 487, 7, 4,
	2, 2, 487, 488, 5, 76, 39, 2, 488, 85, 3, 2, 2, 2, 489, 490, 7, 8, 2, 2,
	490, 491, 5, 34, 18, 2, 491, 492, 7, 10, 2, 2, 492, 493, 5, 54, 28, 2,
	493, 494, 7, 33, 2, 2, 494, 495, 5, 76, 39, 2, 495, 496, 7, 5, 2, 2, 496,
	87, 3, 2, 2, 2, 497, 498, 7, 39, 2, 2, 498, 499, 5, 90, 46, 2, 499, 500,
	7, 40, 2, 2, 500, 89, 3, 2, 2, 2, 501, 506, 5, 92, 47, 2, 502, 503, 9,
	2, 2, 2, 503, 505, 5, 92, 47, 2, 504, 502, 3, 2, 2, 2, 505, 508, 3, 2,
	2, 2, 506, 504, 3, 2, 2, 2, 506, 507, 3, 2, 2, 2, 507, 512, 3, 2, 2, 2,
	508, 506, 3, 2, 2, 2, 509, 511, 9, 2, 2, 2, 510, 509, 3, 2, 2, 2, 511,
	514, 3, 2, 2, 2, 512, 510, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2, 513, 516,
	3, 2, 2, 2, 514, 512, 3, 2, 2, 2, 515, 501, 3, 2, 2, 2, 515, 516, 3, 2,
	2, 2, 516, 91, 3, 2, 2, 2, 517, 521, 5, 54, 28, 2, 518, 521, 5, 94, 48,
	2, 519, 521, 5, 100, 51, 2, 520, 517, 3, 2, 2, 2, 520, 518, 3, 2, 2, 2,
	520, 519, 3, 2, 2, 2, 521, 93, 3, 2, 2, 2, 522, 523, 7, 9, 2, 2, 523, 524,
	5, 54, 28, 2, 524, 525, 7, 33, 2, 2, 525, 529, 5, 90, 46, 2, 526, 528,
	5, 96, 49, 2, 527, 526, 3, 2, 2, 2, 528, 531, 3, 2, 2, 2, 529, 527, 3,
	2, 2, 2, 529, 530, 3, 2, 2, 2, 530, 533, 3, 2, 2, 2, 531, 529, 3, 2, 2,
	2, 532, 534, 5, 98, 50, 2, 533, 532, 3, 2, 2, 2, 533, 534, 3, 2, 2, 2,
	534, 535, 3, 2, 2, 2, 535, 536, 7, 5, 2, 2, 536, 95, 3, 2, 2, 2, 537, 538,
	7, 4, 2, 2, 538, 539, 7, 9, 2, 2, 539, 540, 5, 54, 28, 2, 540, 541, 7,
	33, 2, 2, 541, 542, 5, 90, 46, 2, 542, 97, 3, 2, 2, 2, 543, 544, 7, 4,
	2, 2, 544, 545, 5, 90, 46, 2, 545, 99, 3, 2, 2, 2, 546, 547, 7, 8, 2, 2,
	547, 548, 5, 34, 18, 2, 548, 549, 7, 10, 2, 2, 549, 550, 5, 54, 28, 2,
	550, 551, 7, 33, 2, 2, 551, 552, 5, 90, 46, 2, 552, 553, 7, 5, 2, 2, 553,
	101, 3, 2, 2, 2, 554, 555, 9, 7, 2, 2, 555, 103, 3, 2, 2, 2, 48, 108, 116,
	126, 135, 147, 152, 186, 197, 205, 212, 219, 235, 251, 255, 272, 283, 288,
	294, 306, 309, 315, 335, 337, 347, 359, 372, 376, 382, 385, 387, 397, 410,
	415, 423, 437, 443, 446, 463, 472, 476, 506, 512, 515, 520, 529, 533,
}
var literalNames = []string{
	"", "'and'", "'else'", "'end'", "'false'", "'func'", "'for'", "'if'", "'in'",
	"'meta'", "'not'", "'or'", "'return'", "'true'", "'var'", "'while'", "'defer'",
	"'try'", "'catch'", "'='", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "';'", "','", "':'", "'.'", "'('", "')'",
	"'['", "']'", "'{'", "'}'", "'->'", "'&'", "'|'", "'...'",
}
var symbolicNames = []string{
	"", "AND", "ELSE", "END", "FALSE", "FUNC", "FOR", "IF", "IN", "META", "NOT",
	"OR", "RETURN", "TRUE", "VAR", "WHILE", "DEFER", "TRY", "CATCH", "ASSIGN",
	"EQ", "NE", "LT", "LE", "GT", "GE", "ADD", "SUB", "MUL", "DIV", "MOD",
	"SEMICOLON", "COMMA", "COLON", "PERIOD", "OPAREN", "CPAREN", "OBRACKET",
	"CBRACKET", "OCURLY", "CCURLY", "ARROW", "LAMBDA", "PIPE", "EXPAND", "NUMBER",
	"ID", "WS", "NEWLINE", "STRING", "LQUOTE",
}

var ruleNames = []string{
	"start", "module", "meta_section", "meta_entry", "meta_object", "meta_fields",
	"meta_field", "meta_field_value", "stmts", "stmt", "assignment_stmt", "assignment_lvalues",
	"rvalues", "var_decl_stmt", "var_decl_vars", "for_stmt", "for_vars", "while_stmt",
	"if_stmt", "if_elif", "if_else", "func_stmt", "param_list", "return_stmt",
	"try_catch_stmt", "defer_stmt", "expr", "pipeline_expr", "binary_expr",
	"unary_expr", "primary_expr", "simple_literal", "arg_list", "lvalue_expr",
	"lambda_expr", "short_lambda_expr", "object_literal", "object_fields",
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
	NitroParserELSE      = 2
	NitroParserEND       = 3
	NitroParserFALSE     = 4
	NitroParserFUNC      = 5
	NitroParserFOR       = 6
	NitroParserIF        = 7
	NitroParserIN        = 8
	NitroParserMETA      = 9
	NitroParserNOT       = 10
	NitroParserOR        = 11
	NitroParserRETURN    = 12
	NitroParserTRUE      = 13
	NitroParserVAR       = 14
	NitroParserWHILE     = 15
	NitroParserDEFER     = 16
	NitroParserTRY       = 17
	NitroParserCATCH     = 18
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
	NitroParserARROW     = 41
	NitroParserLAMBDA    = 42
	NitroParserPIPE      = 43
	NitroParserEXPAND    = 44
	NitroParserNUMBER    = 45
	NitroParserID        = 46
	NitroParserWS        = 47
	NitroParserNEWLINE   = 48
	NitroParserSTRING    = 49
	NitroParserLQUOTE    = 50
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
	NitroParserRULE_return_stmt        = 23
	NitroParserRULE_try_catch_stmt     = 24
	NitroParserRULE_defer_stmt         = 25
	NitroParserRULE_expr               = 26
	NitroParserRULE_pipeline_expr      = 27
	NitroParserRULE_binary_expr        = 28
	NitroParserRULE_unary_expr         = 29
	NitroParserRULE_primary_expr       = 30
	NitroParserRULE_simple_literal     = 31
	NitroParserRULE_arg_list           = 32
	NitroParserRULE_lvalue_expr        = 33
	NitroParserRULE_lambda_expr        = 34
	NitroParserRULE_short_lambda_expr  = 35
	NitroParserRULE_object_literal     = 36
	NitroParserRULE_object_fields      = 37
	NitroParserRULE_object_field       = 38
	NitroParserRULE_object_if          = 39
	NitroParserRULE_object_elif        = 40
	NitroParserRULE_object_else        = 41
	NitroParserRULE_object_for         = 42
	NitroParserRULE_array_literal      = 43
	NitroParserRULE_array_elems        = 44
	NitroParserRULE_array_elem         = 45
	NitroParserRULE_array_if           = 46
	NitroParserRULE_array_elif         = 47
	NitroParserRULE_array_else         = 48
	NitroParserRULE_array_for          = 49
	NitroParserRULE_id_or_keyword      = 50
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
		p.SetState(102)
		p.Module()
	}
	{
		p.SetState(103)
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
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserMETA {
		{
			p.SetState(105)
			p.Meta_section()
		}

	}
	{
		p.SetState(108)
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
		p.SetState(110)
		p.Match(NitroParserMETA)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserID {
		{
			p.SetState(111)
			p.Meta_entry()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
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
		p.SetState(119)
		p.Match(NitroParserID)
	}
	{
		p.SetState(120)
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
		p.SetState(122)
		p.Match(NitroParserOCURLY)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(123)
			p.Meta_fields()
		}

	}
	{
		p.SetState(126)
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
		p.SetState(128)
		p.Meta_field()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
		{
			p.SetState(129)
			_la = p.GetTokenStream().LA(1)

			if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(130)
			p.Meta_field()
		}

		p.SetState(135)
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
		p.SetState(136)
		p.Match(NitroParserID)
	}
	{
		p.SetState(137)
		p.Match(NitroParserCOLON)
	}
	{
		p.SetState(138)
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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Match(NitroParserSTRING)
		}

	case NitroParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(NitroParserNUMBER)
		}

	case NitroParserTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(142)
			p.Match(NitroParserTRUE)
		}

	case NitroParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(143)
			p.Match(NitroParserFALSE)
		}

	case NitroParserOCURLY:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(144)
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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserNOT)|(1<<NitroParserRETURN)|(1<<NitroParserTRUE)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserDEFER)|(1<<NitroParserTRY)|(1<<NitroParserADD)|(1<<NitroParserSUB)|(1<<NitroParserSEMICOLON))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserLAMBDA-35))|(1<<(NitroParserPIPE-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(147)
			p.Stmt()
		}

		p.SetState(152)
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

func (s *StmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtContext) Try_catch_stmt() ITry_catch_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITry_catch_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITry_catch_stmtContext)
}

func (s *StmtContext) Defer_stmt() IDefer_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefer_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefer_stmtContext)
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

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Assignment_stmt()
		}
		{
			p.SetState(154)
			p.Match(NitroParserSEMICOLON)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Var_decl_stmt()
		}
		{
			p.SetState(157)
			p.Match(NitroParserSEMICOLON)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.For_stmt()
		}
		{
			p.SetState(160)
			p.Match(NitroParserSEMICOLON)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(162)
			p.While_stmt()
		}
		{
			p.SetState(163)
			p.Match(NitroParserSEMICOLON)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(165)
			p.If_stmt()
		}
		{
			p.SetState(166)
			p.Match(NitroParserSEMICOLON)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(168)
			p.Func_stmt()
		}
		{
			p.SetState(169)
			p.Match(NitroParserSEMICOLON)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(171)
			p.Return_stmt()
		}
		{
			p.SetState(172)
			p.Match(NitroParserSEMICOLON)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(174)
			p.Expr()
		}
		{
			p.SetState(175)
			p.Match(NitroParserSEMICOLON)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(177)
			p.Try_catch_stmt()
		}
		{
			p.SetState(178)
			p.Match(NitroParserSEMICOLON)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(180)
			p.Defer_stmt()
		}
		{
			p.SetState(181)
			p.Match(NitroParserSEMICOLON)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(183)
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
		p.SetState(186)
		p.Assignment_lvalues()
	}
	{
		p.SetState(187)
		p.Match(NitroParserASSIGN)
	}
	{
		p.SetState(188)
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
		p.SetState(190)
		p.Lvalue_expr()
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(191)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(192)
			p.Lvalue_expr()
		}

		p.SetState(197)
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
		p.SetState(198)
		p.Expr()
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
			p.Expr()
		}

		p.SetState(205)
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
		p.SetState(206)
		p.Match(NitroParserVAR)
	}
	{
		p.SetState(207)
		p.Var_decl_vars()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserASSIGN {
		{
			p.SetState(208)
			p.Match(NitroParserASSIGN)
		}
		{
			p.SetState(209)
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
		p.SetState(212)
		p.Match(NitroParserID)
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(213)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(214)
			p.Match(NitroParserID)
		}

		p.SetState(219)
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

func (s *For_stmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
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
		p.SetState(220)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(221)
		p.For_vars()
	}
	{
		p.SetState(222)
		p.Match(NitroParserIN)
	}
	{
		p.SetState(223)
		p.Expr()
	}
	{
		p.SetState(224)
		p.Match(NitroParserSEMICOLON)
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
		p.SetState(228)
		p.Match(NitroParserID)
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(229)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(230)
			p.Match(NitroParserID)
		}

		p.SetState(235)
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

func (s *While_stmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
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
		p.SetState(236)
		p.Match(NitroParserWHILE)
	}
	{
		p.SetState(237)
		p.Expr()
	}
	{
		p.SetState(238)
		p.Match(NitroParserSEMICOLON)
	}
	{
		p.SetState(239)
		p.Stmts()
	}
	{
		p.SetState(240)
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

func (s *If_stmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(243)
		p.Expr()
	}
	{
		p.SetState(244)
		p.Match(NitroParserSEMICOLON)
	}
	{
		p.SetState(245)
		p.Stmts()
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(246)
				p.If_elif()
			}

		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(252)
			p.If_else()
		}

	}
	{
		p.SetState(255)
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

func (s *If_elifContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
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
		p.SetState(257)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(258)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(259)
		p.Expr()
	}
	{
		p.SetState(260)
		p.Match(NitroParserSEMICOLON)
	}
	{
		p.SetState(261)
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
		p.SetState(263)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(264)
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
		p.SetState(266)
		p.Match(NitroParserFUNC)
	}
	{
		p.SetState(267)
		p.Match(NitroParserID)
	}
	{
		p.SetState(268)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(269)
			p.Param_list()
		}

	}
	{
		p.SetState(272)
		p.Match(NitroParserCPAREN)
	}
	{
		p.SetState(273)
		p.Stmts()
	}
	{
		p.SetState(274)
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
		p.SetState(276)
		p.Match(NitroParserID)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(277)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(278)
			p.Match(NitroParserID)
		}

		p.SetState(283)
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
		p.SetState(284)
		p.Match(NitroParserRETURN)
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserLAMBDA-35))|(1<<(NitroParserPIPE-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(285)
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

func (s *Try_catch_stmtContext) CATCH() antlr.TerminalNode {
	return s.GetToken(NitroParserCATCH, 0)
}

func (s *Try_catch_stmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Try_catch_stmtContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
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
		p.SetState(288)
		p.Match(NitroParserTRY)
	}
	{
		p.SetState(289)
		p.Stmts()
	}
	{
		p.SetState(290)
		p.Match(NitroParserCATCH)
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(291)
			p.Match(NitroParserID)
		}

	}
	{
		p.SetState(294)
		p.Match(NitroParserSEMICOLON)
	}
	{
		p.SetState(295)
		p.Stmts()
	}
	{
		p.SetState(296)
		p.Match(NitroParserEND)
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
	p.EnterRule(localctx, 50, NitroParserRULE_defer_stmt)

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
		p.SetState(298)
		p.Match(NitroParserDEFER)
	}
	{
		p.SetState(299)
		p.primary_expr(0)
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

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Expr_pipelineContext struct {
	*ExprContext
}

func NewExpr_pipelineContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_pipelineContext {
	var p = new(Expr_pipelineContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Expr_pipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_pipelineContext) Pipeline_expr() IPipeline_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeline_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeline_exprContext)
}

func (s *Expr_pipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterExpr_pipeline(s)
	}
}

func (s *Expr_pipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitExpr_pipeline(s)
	}
}

type Expr_binaryContext struct {
	*ExprContext
}

func NewExpr_binaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_binaryContext {
	var p = new(Expr_binaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Expr_binaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_binaryContext) Binary_expr() IBinary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exprContext)
}

func (s *Expr_binaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterExpr_binary(s)
	}
}

func (s *Expr_binaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitExpr_binary(s)
	}
}

type Expr_short_lambdaContext struct {
	*ExprContext
}

func NewExpr_short_lambdaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_short_lambdaContext {
	var p = new(Expr_short_lambdaContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Expr_short_lambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_short_lambdaContext) Short_lambda_expr() IShort_lambda_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShort_lambda_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShort_lambda_exprContext)
}

func (s *Expr_short_lambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterExpr_short_lambda(s)
	}
}

func (s *Expr_short_lambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitExpr_short_lambda(s)
	}
}

func (p *NitroParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, NitroParserRULE_expr)

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

	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpr_binaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)
			p.binary_expr(0)
		}

	case 2:
		localctx = NewExpr_short_lambdaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)
			p.Short_lambda_expr()
		}

	case 3:
		localctx = NewExpr_pipelineContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(303)
			p.Pipeline_expr()
		}

	}

	return localctx
}

// IPipeline_exprContext is an interface to support dynamic dispatch.
type IPipeline_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeline_exprContext differentiates from other interfaces.
	IsPipeline_exprContext()
}

type Pipeline_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeline_exprContext() *Pipeline_exprContext {
	var p = new(Pipeline_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_pipeline_expr
	return p
}

func (*Pipeline_exprContext) IsPipeline_exprContext() {}

func NewPipeline_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pipeline_exprContext {
	var p = new(Pipeline_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_pipeline_expr

	return p
}

func (s *Pipeline_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Pipeline_exprContext) AllBinary_expr() []IBinary_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBinary_exprContext)(nil)).Elem())
	var tst = make([]IBinary_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBinary_exprContext)
		}
	}

	return tst
}

func (s *Pipeline_exprContext) Binary_expr(i int) IBinary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBinary_exprContext)
}

func (s *Pipeline_exprContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(NitroParserPIPE)
}

func (s *Pipeline_exprContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(NitroParserPIPE, i)
}

func (s *Pipeline_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pipeline_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pipeline_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterPipeline_expr(s)
	}
}

func (s *Pipeline_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitPipeline_expr(s)
	}
}

func (p *NitroParser) Pipeline_expr() (localctx IPipeline_exprContext) {
	localctx = NewPipeline_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, NitroParserRULE_pipeline_expr)
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
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(306)
			p.binary_expr(0)
		}

	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NitroParserPIPE {
		{
			p.SetState(309)
			p.Match(NitroParserPIPE)
		}
		{
			p.SetState(310)
			p.binary_expr(0)
		}

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	_startState := 56
	p.EnterRecursionRule(localctx, 56, NitroParserRULE_binary_expr, _p)
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
		p.SetState(316)
		p.Unary_expr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(333)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(318)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(319)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Binary_exprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserMUL)|(1<<NitroParserDIV)|(1<<NitroParserMOD))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Binary_exprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(320)
					p.binary_expr(6)
				}

			case 2:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(321)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(322)

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
					p.SetState(323)
					p.binary_expr(5)
				}

			case 3:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(324)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(325)

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
					p.SetState(326)
					p.binary_expr(4)
				}

			case 4:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(327)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(328)

					var _m = p.Match(NitroParserAND)

					localctx.(*Binary_exprContext).op = _m
				}
				{
					p.SetState(329)
					p.binary_expr(3)
				}

			case 5:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(330)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(331)

					var _m = p.Match(NitroParserOR)

					localctx.(*Binary_exprContext).op = _m
				}
				{
					p.SetState(332)
					p.binary_expr(2)
				}

			}

		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 58, NitroParserRULE_unary_expr)

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

	p.SetState(345)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)

			var _m = p.Match(NitroParserNOT)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(339)
			p.Unary_expr()
		}

	case NitroParserADD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(340)

			var _m = p.Match(NitroParserADD)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(341)
			p.Unary_expr()
		}

	case NitroParserSUB:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(342)

			var _m = p.Match(NitroParserSUB)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(343)
			p.Unary_expr()
		}

	case NitroParserFALSE, NitroParserFUNC, NitroParserTRUE, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserNUMBER, NitroParserID, NitroParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(344)
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
	_startState := 60
	p.EnterRecursionRule(localctx, 60, NitroParserRULE_primary_expr, _p)
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
	p.SetState(357)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserID:
		localctx = NewPrimary_expr_simple_refContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(348)
			p.Match(NitroParserID)
		}

	case NitroParserFUNC:
		localctx = NewPrimary_expr_lambdaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(349)
			p.Lambda_expr()
		}

	case NitroParserOCURLY:
		localctx = NewPrimary_expr_objectContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(350)
			p.Object_literal()
		}

	case NitroParserOBRACKET:
		localctx = NewPrimary_expr_arrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(351)
			p.Array_literal()
		}

	case NitroParserFALSE, NitroParserTRUE, NitroParserNUMBER, NitroParserSTRING:
		localctx = NewPrimary_expr_literalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(352)
			p.Simple_literal()
		}

	case NitroParserOPAREN:
		localctx = NewPrimary_expr_parenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(353)
			p.Match(NitroParserOPAREN)
		}
		{
			p.SetState(354)
			p.Expr()
		}
		{
			p.SetState(355)
			p.Match(NitroParserCPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(383)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimary_expr_member_accessContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(359)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(360)
					p.Match(NitroParserPERIOD)
				}
				{
					p.SetState(361)
					p.Match(NitroParserID)
				}

			case 2:
				localctx = NewPrimary_expr_indexContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(362)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(363)
					p.Match(NitroParserOBRACKET)
				}
				{
					p.SetState(364)
					p.Expr()
				}
				{
					p.SetState(365)
					p.Match(NitroParserCBRACKET)
				}

			case 3:
				localctx = NewPrimary_expr_sliceContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(367)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(368)
					p.Match(NitroParserOBRACKET)
				}
				p.SetState(370)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserLAMBDA-35))|(1<<(NitroParserPIPE-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
					{
						p.SetState(369)

						var _x = p.Expr()

						localctx.(*Primary_expr_sliceContext).b = _x
					}

				}
				{
					p.SetState(372)
					p.Match(NitroParserCOLON)
				}
				p.SetState(374)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserLAMBDA-35))|(1<<(NitroParserPIPE-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
					{
						p.SetState(373)

						var _x = p.Expr()

						localctx.(*Primary_expr_sliceContext).e = _x
					}

				}
				{
					p.SetState(376)
					p.Match(NitroParserCBRACKET)
				}

			case 4:
				localctx = NewPrimary_expr_callContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(377)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(378)
					p.Match(NitroParserOPAREN)
				}
				p.SetState(380)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserLAMBDA-35))|(1<<(NitroParserPIPE-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
					{
						p.SetState(379)
						p.Arg_list()
					}

				}
				{
					p.SetState(382)
					p.Match(NitroParserCPAREN)
				}

			}

		}
		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 62, NitroParserRULE_simple_literal)
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
		p.SetState(388)

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
	p.EnterRule(localctx, 64, NitroParserRULE_arg_list)
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
		p.SetState(390)
		p.Expr()
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(391)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(392)
			p.Expr()
		}

		p.SetState(397)
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
	p.EnterRule(localctx, 66, NitroParserRULE_lvalue_expr)

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

	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLvalue_expr_simple_refContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.Match(NitroParserID)
		}

	case 2:
		localctx = NewLvalue_expr_member_accessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(399)
			p.primary_expr(0)
		}
		{
			p.SetState(400)
			p.Match(NitroParserPERIOD)
		}
		{
			p.SetState(401)
			p.Match(NitroParserID)
		}

	case 3:
		localctx = NewLvalue_expr_indexContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(403)
			p.primary_expr(0)
		}
		{
			p.SetState(404)
			p.Match(NitroParserOBRACKET)
		}
		{
			p.SetState(405)
			p.Expr()
		}
		{
			p.SetState(406)
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
	p.EnterRule(localctx, 68, NitroParserRULE_lambda_expr)
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
		p.SetState(410)
		p.Match(NitroParserFUNC)
	}
	{
		p.SetState(411)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(412)
			p.Param_list()
		}

	}
	{
		p.SetState(415)
		p.Match(NitroParserCPAREN)
	}
	{
		p.SetState(416)
		p.Stmts()
	}
	{
		p.SetState(417)
		p.Match(NitroParserEND)
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
	p.EnterRule(localctx, 70, NitroParserRULE_short_lambda_expr)
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
		p.SetState(419)
		p.Match(NitroParserLAMBDA)
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(420)
			p.Param_list()
		}

	}
	{
		p.SetState(423)
		p.Match(NitroParserARROW)
	}
	{
		p.SetState(424)
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
	p.EnterRule(localctx, 72, NitroParserRULE_object_literal)

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
		p.SetState(426)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(427)
		p.Object_fields()
	}
	{
		p.SetState(428)
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
	p.EnterRule(localctx, 74, NitroParserRULE_object_fields)
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
	p.SetState(444)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(430)
			p.Object_field()
		}
		p.SetState(435)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(431)
					_la = p.GetTokenStream().LA(1)

					if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(432)
					p.Object_field()
				}

			}
			p.SetState(437)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
			{
				p.SetState(438)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(443)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 76, NitroParserRULE_object_field)

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

	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		localctx = NewObject_field_id_keyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(446)
			p.Id_or_keyword()
		}
		{
			p.SetState(447)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(448)
			p.Expr()
		}

	case 2:
		localctx = NewObject_field_expr_keyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(450)
			p.Match(NitroParserOBRACKET)
		}
		{
			p.SetState(451)
			p.Expr()
		}
		{
			p.SetState(452)
			p.Match(NitroParserCBRACKET)
		}
		{
			p.SetState(453)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(454)
			p.Expr()
		}

	case 3:
		localctx = NewObject_field_expansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(456)
			p.primary_expr(0)
		}
		{
			p.SetState(457)
			p.Match(NitroParserEXPAND)
		}

	case 4:
		localctx = NewObject_field_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(459)
			p.Object_if()
		}

	case 5:
		localctx = NewObject_field_forContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(460)
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

func (s *Object_ifContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Object_ifContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_ifContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
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
	p.EnterRule(localctx, 78, NitroParserRULE_object_if)
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
		p.SetState(463)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(464)
		p.Expr()
	}
	{
		p.SetState(465)
		p.Match(NitroParserSEMICOLON)
	}
	{
		p.SetState(466)
		p.Object_fields()
	}
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(467)
				p.Object_elif()
			}

		}
		p.SetState(472)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(473)
			p.Object_else()
		}

	}
	{
		p.SetState(476)
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

func (s *Object_elifContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
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
	p.EnterRule(localctx, 80, NitroParserRULE_object_elif)

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
		p.SetState(478)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(479)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(480)
		p.Expr()
	}
	{
		p.SetState(481)
		p.Match(NitroParserSEMICOLON)
	}
	{
		p.SetState(482)
		p.Object_fields()
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
	p.EnterRule(localctx, 82, NitroParserRULE_object_else)

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
		p.SetState(484)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(485)
		p.Object_fields()
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

func (s *Object_forContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Object_forContext) Object_fields() IObject_fieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_fieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_fieldsContext)
}

func (s *Object_forContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
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
	p.EnterRule(localctx, 84, NitroParserRULE_object_for)

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
		p.SetState(487)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(488)
		p.For_vars()
	}
	{
		p.SetState(489)
		p.Match(NitroParserIN)
	}
	{
		p.SetState(490)
		p.Expr()
	}
	{
		p.SetState(491)
		p.Match(NitroParserSEMICOLON)
	}
	{
		p.SetState(492)
		p.Object_fields()
	}
	{
		p.SetState(493)
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
	p.EnterRule(localctx, 86, NitroParserRULE_array_literal)

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
		p.Match(NitroParserOBRACKET)
	}
	{
		p.SetState(496)
		p.Array_elems()
	}
	{
		p.SetState(497)
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
	p.EnterRule(localctx, 88, NitroParserRULE_array_elems)
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
	p.SetState(513)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserNOT)|(1<<NitroParserTRUE)|(1<<NitroParserADD)|(1<<NitroParserSUB))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(NitroParserOPAREN-35))|(1<<(NitroParserOBRACKET-35))|(1<<(NitroParserOCURLY-35))|(1<<(NitroParserLAMBDA-35))|(1<<(NitroParserPIPE-35))|(1<<(NitroParserNUMBER-35))|(1<<(NitroParserID-35))|(1<<(NitroParserSTRING-35)))) != 0) {
		{
			p.SetState(499)
			p.Array_elem()
		}
		p.SetState(504)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(500)
					_la = p.GetTokenStream().LA(1)

					if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(501)
					p.Array_elem()
				}

			}
			p.SetState(506)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
		}
		p.SetState(510)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
			{
				p.SetState(507)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(512)
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
	p.EnterRule(localctx, 90, NitroParserRULE_array_elem)

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

	p.SetState(518)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserFALSE, NitroParserFUNC, NitroParserNOT, NitroParserTRUE, NitroParserADD, NitroParserSUB, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserLAMBDA, NitroParserPIPE, NitroParserNUMBER, NitroParserID, NitroParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(515)
			p.Expr()
		}

	case NitroParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(516)
			p.Array_if()
		}

	case NitroParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(517)
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

func (s *Array_ifContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Array_ifContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_ifContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
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
	p.EnterRule(localctx, 92, NitroParserRULE_array_if)
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
		p.SetState(520)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(521)
		p.Expr()
	}
	{
		p.SetState(522)
		p.Match(NitroParserSEMICOLON)
	}
	{
		p.SetState(523)
		p.Array_elems()
	}
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(524)
				p.Array_elif()
			}

		}
		p.SetState(529)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(530)
			p.Array_else()
		}

	}
	{
		p.SetState(533)
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

func (s *Array_elifContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
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
	p.EnterRule(localctx, 94, NitroParserRULE_array_elif)

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
		p.SetState(535)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(536)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(537)
		p.Expr()
	}
	{
		p.SetState(538)
		p.Match(NitroParserSEMICOLON)
	}
	{
		p.SetState(539)
		p.Array_elems()
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
	p.EnterRule(localctx, 96, NitroParserRULE_array_else)

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
		p.SetState(541)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(542)
		p.Array_elems()
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

func (s *Array_forContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(NitroParserSEMICOLON, 0)
}

func (s *Array_forContext) Array_elems() IArray_elemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elemsContext)
}

func (s *Array_forContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
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
	p.EnterRule(localctx, 98, NitroParserRULE_array_for)

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
		p.SetState(544)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(545)
		p.For_vars()
	}
	{
		p.SetState(546)
		p.Match(NitroParserIN)
	}
	{
		p.SetState(547)
		p.Expr()
	}
	{
		p.SetState(548)
		p.Match(NitroParserSEMICOLON)
	}
	{
		p.SetState(549)
		p.Array_elems()
	}
	{
		p.SetState(550)
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

func (s *Id_or_keywordContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NitroParserELSE, 0)
}

func (s *Id_or_keywordContext) END() antlr.TerminalNode {
	return s.GetToken(NitroParserEND, 0)
}

func (s *Id_or_keywordContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NitroParserFALSE, 0)
}

func (s *Id_or_keywordContext) FUNC() antlr.TerminalNode {
	return s.GetToken(NitroParserFUNC, 0)
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

func (s *Id_or_keywordContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NitroParserTRUE, 0)
}

func (s *Id_or_keywordContext) VAR() antlr.TerminalNode {
	return s.GetToken(NitroParserVAR, 0)
}

func (s *Id_or_keywordContext) WHILE() antlr.TerminalNode {
	return s.GetToken(NitroParserWHILE, 0)
}

func (s *Id_or_keywordContext) DEFER() antlr.TerminalNode {
	return s.GetToken(NitroParserDEFER, 0)
}

func (s *Id_or_keywordContext) TRY() antlr.TerminalNode {
	return s.GetToken(NitroParserTRY, 0)
}

func (s *Id_or_keywordContext) CATCH() antlr.TerminalNode {
	return s.GetToken(NitroParserCATCH, 0)
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
	p.EnterRule(localctx, 100, NitroParserRULE_id_or_keyword)
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
		p.SetState(552)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Id_or_keywordContext).t = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserAND)|(1<<NitroParserELSE)|(1<<NitroParserEND)|(1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserFOR)|(1<<NitroParserIF)|(1<<NitroParserIN)|(1<<NitroParserMETA)|(1<<NitroParserNOT)|(1<<NitroParserOR)|(1<<NitroParserRETURN)|(1<<NitroParserTRUE)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserDEFER)|(1<<NitroParserTRY)|(1<<NitroParserCATCH))) != 0) || _la == NitroParserID) {
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
	case 28:
		var t *Binary_exprContext = nil
		if localctx != nil {
			t = localctx.(*Binary_exprContext)
		}
		return p.Binary_expr_Sempred(t, predIndex)

	case 30:
		var t *Primary_exprContext = nil
		if localctx != nil {
			t = localctx.(*Primary_exprContext)
		}
		return p.Primary_expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NitroParser) Binary_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
