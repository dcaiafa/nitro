// Code generated from NitroParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 67, 656,
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
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 127, 10, 3, 12, 3, 14,
	3, 130, 11, 3, 3, 3, 7, 3, 133, 10, 3, 12, 3, 14, 3, 136, 11, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 143, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	5, 5, 150, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 158, 10, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 164, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 5, 7, 172, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 178, 10, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 185, 10, 8, 12, 8, 14, 8, 188, 11, 8, 3,
	8, 5, 8, 191, 10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 196, 10, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 5, 11, 202, 10, 11, 3, 11, 3, 11, 3, 11, 3, 12, 5, 12, 208,
	10, 12, 3, 12, 7, 12, 211, 10, 12, 12, 12, 14, 12, 214, 11, 12, 3, 13,
	3, 13, 6, 13, 218, 10, 13, 13, 13, 14, 13, 219, 3, 13, 7, 13, 223, 10,
	13, 12, 13, 14, 13, 226, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	243, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 7, 16, 252,
	10, 16, 12, 16, 14, 16, 255, 11, 16, 3, 17, 3, 17, 3, 17, 7, 17, 260, 10,
	17, 12, 17, 14, 17, 263, 11, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 3, 19, 5, 19, 273, 10, 19, 3, 20, 3, 20, 3, 20, 7, 20, 278,
	10, 20, 12, 20, 14, 20, 281, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 7, 22, 294, 10, 22, 12, 22, 14,
	22, 297, 11, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 311, 10, 24, 12, 24, 14, 24, 314, 11,
	24, 3, 24, 5, 24, 317, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 5,
	27, 335, 10, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28,
	7, 28, 345, 10, 28, 12, 28, 14, 28, 348, 11, 28, 3, 29, 3, 29, 5, 29, 352,
	10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 360, 10, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33,
	3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 7, 36, 385, 10, 36, 12, 36, 14, 36, 388, 11, 36, 3, 37, 3, 37,
	5, 37, 392, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 38, 7, 38, 403, 10, 38, 12, 38, 14, 38, 406, 11, 38, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 7, 39, 426, 10, 39, 12, 39, 14,
	39, 429, 11, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40,
	438, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 41, 5, 41, 451, 10, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 464, 10, 41, 3, 41, 3,
	41, 5, 41, 468, 10, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 475,
	10, 41, 5, 41, 477, 10, 41, 3, 41, 7, 41, 480, 10, 41, 12, 41, 14, 41,
	483, 11, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 7, 43, 490, 10, 43, 12,
	43, 14, 43, 493, 11, 43, 3, 43, 5, 43, 496, 10, 43, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 508, 10, 44, 3,
	45, 3, 45, 3, 45, 5, 45, 513, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 46, 3, 46, 5, 46, 522, 10, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	47, 3, 47, 3, 48, 3, 48, 3, 48, 7, 48, 534, 10, 48, 12, 48, 14, 48, 537,
	11, 48, 3, 48, 5, 48, 540, 10, 48, 5, 48, 542, 10, 48, 3, 49, 3, 49, 3,
	49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 5, 49, 559, 10, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3,
	50, 7, 50, 567, 10, 50, 12, 50, 14, 50, 570, 11, 50, 3, 50, 5, 50, 573,
	10, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52,
	3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3,
	53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 7, 55, 602, 10, 55,
	12, 55, 14, 55, 605, 11, 55, 3, 55, 7, 55, 608, 10, 55, 12, 55, 14, 55,
	611, 11, 55, 5, 55, 613, 10, 55, 3, 56, 3, 56, 3, 56, 5, 56, 618, 10, 56,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 7, 57, 626, 10, 57, 12, 57, 14,
	57, 629, 11, 57, 3, 57, 5, 57, 632, 10, 57, 3, 58, 3, 58, 3, 58, 3, 58,
	3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3,
	60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 2, 6, 70, 74,
	76, 80, 62, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
	70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
	106, 108, 110, 112, 114, 116, 118, 120, 2, 9, 3, 2, 45, 46, 8, 2, 12, 12,
	17, 17, 22, 22, 59, 59, 63, 63, 66, 66, 3, 2, 28, 32, 3, 2, 41, 43, 3,
	2, 39, 40, 3, 2, 33, 38, 4, 2, 6, 26, 60, 60, 2, 684, 2, 122, 3, 2, 2,
	2, 4, 128, 3, 2, 2, 2, 6, 142, 3, 2, 2, 2, 8, 144, 3, 2, 2, 2, 10, 153,
	3, 2, 2, 2, 12, 167, 3, 2, 2, 2, 14, 181, 3, 2, 2, 2, 16, 192, 3, 2, 2,
	2, 18, 197, 3, 2, 2, 2, 20, 199, 3, 2, 2, 2, 22, 207, 3, 2, 2, 2, 24, 215,
	3, 2, 2, 2, 26, 242, 3, 2, 2, 2, 28, 244, 3, 2, 2, 2, 30, 248, 3, 2, 2,
	2, 32, 256, 3, 2, 2, 2, 34, 264, 3, 2, 2, 2, 36, 268, 3, 2, 2, 2, 38, 274,
	3, 2, 2, 2, 40, 282, 3, 2, 2, 2, 42, 290, 3, 2, 2, 2, 44, 298, 3, 2, 2,
	2, 46, 304, 3, 2, 2, 2, 48, 318, 3, 2, 2, 2, 50, 325, 3, 2, 2, 2, 52, 330,
	3, 2, 2, 2, 54, 341, 3, 2, 2, 2, 56, 349, 3, 2, 2, 2, 58, 353, 3, 2, 2,
	2, 60, 365, 3, 2, 2, 2, 62, 368, 3, 2, 2, 2, 64, 371, 3, 2, 2, 2, 66, 374,
	3, 2, 2, 2, 68, 376, 3, 2, 2, 2, 70, 378, 3, 2, 2, 2, 72, 391, 3, 2, 2,
	2, 74, 393, 3, 2, 2, 2, 76, 407, 3, 2, 2, 2, 78, 437, 3, 2, 2, 2, 80, 450,
	3, 2, 2, 2, 82, 484, 3, 2, 2, 2, 84, 486, 3, 2, 2, 2, 86, 507, 3, 2, 2,
	2, 88, 509, 3, 2, 2, 2, 90, 519, 3, 2, 2, 2, 92, 526, 3, 2, 2, 2, 94, 541,
	3, 2, 2, 2, 96, 558, 3, 2, 2, 2, 98, 560, 3, 2, 2, 2, 100, 574, 3, 2, 2,
	2, 102, 581, 3, 2, 2, 2, 104, 586, 3, 2, 2, 2, 106, 594, 3, 2, 2, 2, 108,
	612, 3, 2, 2, 2, 110, 617, 3, 2, 2, 2, 112, 619, 3, 2, 2, 2, 114, 633,
	3, 2, 2, 2, 116, 640, 3, 2, 2, 2, 118, 645, 3, 2, 2, 2, 120, 653, 3, 2,
	2, 2, 122, 123, 5, 4, 3, 2, 123, 124, 7, 2, 2, 3, 124, 3, 3, 2, 2, 2, 125,
	127, 5, 6, 4, 2, 126, 125, 3, 2, 2, 2, 127, 130, 3, 2, 2, 2, 128, 126,
	3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 134, 3, 2, 2, 2, 130, 128, 3, 2,
	2, 2, 131, 133, 5, 20, 11, 2, 132, 131, 3, 2, 2, 2, 133, 136, 3, 2, 2,
	2, 134, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2, 136,
	134, 3, 2, 2, 2, 137, 138, 5, 22, 12, 2, 138, 5, 3, 2, 2, 2, 139, 143,
	5, 8, 5, 2, 140, 143, 5, 10, 6, 2, 141, 143, 5, 12, 7, 2, 142, 139, 3,
	2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 141, 3, 2, 2, 2, 143, 7, 3, 2, 2, 2,
	144, 149, 7, 3, 2, 2, 145, 146, 7, 53, 2, 2, 146, 147, 5, 14, 8, 2, 147,
	148, 7, 54, 2, 2, 148, 150, 3, 2, 2, 2, 149, 145, 3, 2, 2, 2, 149, 150,
	3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 152, 7, 45, 2, 2, 152, 9, 3, 2,
	2, 2, 153, 154, 7, 4, 2, 2, 154, 157, 7, 60, 2, 2, 155, 156, 7, 27, 2,
	2, 156, 158, 5, 70, 36, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2,
	158, 163, 3, 2, 2, 2, 159, 160, 7, 53, 2, 2, 160, 161, 5, 14, 8, 2, 161,
	162, 7, 54, 2, 2, 162, 164, 3, 2, 2, 2, 163, 159, 3, 2, 2, 2, 163, 164,
	3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 7, 45, 2, 2, 166, 11, 3, 2,
	2, 2, 167, 168, 7, 5, 2, 2, 168, 171, 7, 60, 2, 2, 169, 170, 7, 27, 2,
	2, 170, 172, 5, 70, 36, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2,
	172, 177, 3, 2, 2, 2, 173, 174, 7, 53, 2, 2, 174, 175, 5, 14, 8, 2, 175,
	176, 7, 54, 2, 2, 176, 178, 3, 2, 2, 2, 177, 173, 3, 2, 2, 2, 177, 178,
	3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 180, 7, 45, 2, 2, 180, 13, 3, 2,
	2, 2, 181, 186, 5, 16, 9, 2, 182, 183, 9, 2, 2, 2, 183, 185, 5, 16, 9,
	2, 184, 182, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186,
	187, 3, 2, 2, 2, 187, 190, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189, 191,
	9, 2, 2, 2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 15, 3, 2,
	2, 2, 192, 195, 5, 120, 61, 2, 193, 194, 7, 47, 2, 2, 194, 196, 5, 18,
	10, 2, 195, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 17, 3, 2, 2, 2,
	197, 198, 9, 3, 2, 2, 198, 19, 3, 2, 2, 2, 199, 201, 7, 16, 2, 2, 200,
	202, 7, 60, 2, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203,
	3, 2, 2, 2, 203, 204, 7, 66, 2, 2, 204, 205, 7, 45, 2, 2, 205, 21, 3, 2,
	2, 2, 206, 208, 5, 24, 13, 2, 207, 206, 3, 2, 2, 2, 207, 208, 3, 2, 2,
	2, 208, 212, 3, 2, 2, 2, 209, 211, 7, 45, 2, 2, 210, 209, 3, 2, 2, 2, 211,
	214, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 23, 3,
	2, 2, 2, 214, 212, 3, 2, 2, 2, 215, 224, 5, 26, 14, 2, 216, 218, 7, 45,
	2, 2, 217, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 217, 3, 2, 2, 2,
	219, 220, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 223, 5, 26, 14, 2, 222,
	217, 3, 2, 2, 2, 223, 226, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 224, 225,
	3, 2, 2, 2, 225, 25, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 227, 243, 5, 28,
	15, 2, 228, 243, 5, 34, 18, 2, 229, 243, 5, 36, 19, 2, 230, 243, 5, 40,
	21, 2, 231, 243, 5, 44, 23, 2, 232, 243, 5, 46, 24, 2, 233, 243, 5, 52,
	27, 2, 234, 243, 5, 56, 29, 2, 235, 243, 5, 70, 36, 2, 236, 243, 5, 58,
	30, 2, 237, 243, 5, 60, 31, 2, 238, 243, 5, 62, 32, 2, 239, 243, 5, 64,
	33, 2, 240, 243, 5, 66, 34, 2, 241, 243, 5, 68, 35, 2, 242, 227, 3, 2,
	2, 2, 242, 228, 3, 2, 2, 2, 242, 229, 3, 2, 2, 2, 242, 230, 3, 2, 2, 2,
	242, 231, 3, 2, 2, 2, 242, 232, 3, 2, 2, 2, 242, 233, 3, 2, 2, 2, 242,
	234, 3, 2, 2, 2, 242, 235, 3, 2, 2, 2, 242, 236, 3, 2, 2, 2, 242, 237,
	3, 2, 2, 2, 242, 238, 3, 2, 2, 2, 242, 239, 3, 2, 2, 2, 242, 240, 3, 2,
	2, 2, 242, 241, 3, 2, 2, 2, 243, 27, 3, 2, 2, 2, 244, 245, 5, 30, 16, 2,
	245, 246, 7, 27, 2, 2, 246, 247, 5, 32, 17, 2, 247, 29, 3, 2, 2, 2, 248,
	253, 5, 86, 44, 2, 249, 250, 7, 46, 2, 2, 250, 252, 5, 86, 44, 2, 251,
	249, 3, 2, 2, 2, 252, 255, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254,
	3, 2, 2, 2, 254, 31, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 256, 261, 5, 70,
	36, 2, 257, 258, 7, 46, 2, 2, 258, 260, 5, 70, 36, 2, 259, 257, 3, 2, 2,
	2, 260, 263, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262,
	33, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 264, 265, 5, 86, 44, 2, 265, 266,
	9, 4, 2, 2, 266, 267, 5, 70, 36, 2, 267, 35, 3, 2, 2, 2, 268, 269, 7, 24,
	2, 2, 269, 272, 5, 38, 20, 2, 270, 271, 7, 27, 2, 2, 271, 273, 5, 32, 17,
	2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 37, 3, 2, 2, 2, 274,
	279, 7, 60, 2, 2, 275, 276, 7, 46, 2, 2, 276, 278, 7, 60, 2, 2, 277, 275,
	3, 2, 2, 2, 278, 281, 3, 2, 2, 2, 279, 277, 3, 2, 2, 2, 279, 280, 3, 2,
	2, 2, 280, 39, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 282, 283, 7, 13, 2, 2,
	283, 284, 5, 42, 22, 2, 284, 285, 7, 60, 2, 2, 285, 286, 5, 70, 36, 2,
	286, 287, 7, 53, 2, 2, 287, 288, 5, 22, 12, 2, 288, 289, 7, 54, 2, 2, 289,
	41, 3, 2, 2, 2, 290, 295, 7, 60, 2, 2, 291, 292, 7, 46, 2, 2, 292, 294,
	7, 60, 2, 2, 293, 291, 3, 2, 2, 2, 294, 297, 3, 2, 2, 2, 295, 293, 3, 2,
	2, 2, 295, 296, 3, 2, 2, 2, 296, 43, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2,
	298, 299, 7, 25, 2, 2, 299, 300, 5, 70, 36, 2, 300, 301, 7, 53, 2, 2, 301,
	302, 5, 22, 12, 2, 302, 303, 7, 54, 2, 2, 303, 45, 3, 2, 2, 2, 304, 305,
	7, 15, 2, 2, 305, 306, 5, 70, 36, 2, 306, 307, 7, 53, 2, 2, 307, 308, 5,
	22, 12, 2, 308, 312, 7, 54, 2, 2, 309, 311, 5, 48, 25, 2, 310, 309, 3,
	2, 2, 2, 311, 314, 3, 2, 2, 2, 312, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2,
	2, 313, 316, 3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 315, 317, 5, 50, 26, 2,
	316, 315, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 47, 3, 2, 2, 2, 318, 319,
	7, 11, 2, 2, 319, 320, 7, 15, 2, 2, 320, 321, 5, 70, 36, 2, 321, 322, 7,
	53, 2, 2, 322, 323, 5, 22, 12, 2, 323, 324, 7, 54, 2, 2, 324, 49, 3, 2,
	2, 2, 325, 326, 7, 11, 2, 2, 326, 327, 7, 53, 2, 2, 327, 328, 5, 22, 12,
	2, 328, 329, 7, 54, 2, 2, 329, 51, 3, 2, 2, 2, 330, 331, 7, 14, 2, 2, 331,
	332, 7, 60, 2, 2, 332, 334, 7, 49, 2, 2, 333, 335, 5, 54, 28, 2, 334, 333,
	3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 336, 3, 2, 2, 2, 336, 337, 7, 50,
	2, 2, 337, 338, 7, 53, 2, 2, 338, 339, 5, 22, 12, 2, 339, 340, 7, 54, 2,
	2, 340, 53, 3, 2, 2, 2, 341, 346, 7, 60, 2, 2, 342, 343, 7, 46, 2, 2, 343,
	345, 7, 60, 2, 2, 344, 342, 3, 2, 2, 2, 345, 348, 3, 2, 2, 2, 346, 344,
	3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 55, 3, 2, 2, 2, 348, 346, 3, 2,
	2, 2, 349, 351, 7, 20, 2, 2, 350, 352, 5, 32, 17, 2, 351, 350, 3, 2, 2,
	2, 351, 352, 3, 2, 2, 2, 352, 57, 3, 2, 2, 2, 353, 354, 7, 23, 2, 2, 354,
	355, 7, 53, 2, 2, 355, 356, 5, 22, 12, 2, 356, 357, 7, 54, 2, 2, 357, 359,
	7, 8, 2, 2, 358, 360, 7, 60, 2, 2, 359, 358, 3, 2, 2, 2, 359, 360, 3, 2,
	2, 2, 360, 361, 3, 2, 2, 2, 361, 362, 7, 53, 2, 2, 362, 363, 5, 22, 12,
	2, 363, 364, 7, 54, 2, 2, 364, 59, 3, 2, 2, 2, 365, 366, 7, 21, 2, 2, 366,
	367, 5, 70, 36, 2, 367, 61, 3, 2, 2, 2, 368, 369, 7, 10, 2, 2, 369, 370,
	5, 80, 41, 2, 370, 63, 3, 2, 2, 2, 371, 372, 7, 26, 2, 2, 372, 373, 5,
	32, 17, 2, 373, 65, 3, 2, 2, 2, 374, 375, 7, 7, 2, 2, 375, 67, 3, 2, 2,
	2, 376, 377, 7, 9, 2, 2, 377, 69, 3, 2, 2, 2, 378, 379, 8, 36, 1, 2, 379,
	380, 5, 72, 37, 2, 380, 386, 3, 2, 2, 2, 381, 382, 12, 4, 2, 2, 382, 383,
	7, 57, 2, 2, 383, 385, 5, 70, 36, 5, 384, 381, 3, 2, 2, 2, 385, 388, 3,
	2, 2, 2, 386, 384, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 71, 3, 2, 2,
	2, 388, 386, 3, 2, 2, 2, 389, 392, 5, 90, 46, 2, 390, 392, 5, 74, 38, 2,
	391, 389, 3, 2, 2, 2, 391, 390, 3, 2, 2, 2, 392, 73, 3, 2, 2, 2, 393, 394,
	8, 38, 1, 2, 394, 395, 5, 76, 39, 2, 395, 404, 3, 2, 2, 2, 396, 397, 12,
	4, 2, 2, 397, 398, 7, 44, 2, 2, 398, 399, 5, 74, 38, 2, 399, 400, 7, 47,
	2, 2, 400, 401, 5, 74, 38, 4, 401, 403, 3, 2, 2, 2, 402, 396, 3, 2, 2,
	2, 403, 406, 3, 2, 2, 2, 404, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405,
	75, 3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 407, 408, 8, 39, 1, 2, 408, 409,
	5, 78, 40, 2, 409, 427, 3, 2, 2, 2, 410, 411, 12, 7, 2, 2, 411, 412, 9,
	5, 2, 2, 412, 426, 5, 76, 39, 8, 413, 414, 12, 6, 2, 2, 414, 415, 9, 6,
	2, 2, 415, 426, 5, 76, 39, 7, 416, 417, 12, 5, 2, 2, 417, 418, 9, 7, 2,
	2, 418, 426, 5, 76, 39, 6, 419, 420, 12, 4, 2, 2, 420, 421, 7, 6, 2, 2,
	421, 426, 5, 76, 39, 5, 422, 423, 12, 3, 2, 2, 423, 424, 7, 19, 2, 2, 424,
	426, 5, 76, 39, 4, 425, 410, 3, 2, 2, 2, 425, 413, 3, 2, 2, 2, 425, 416,
	3, 2, 2, 2, 425, 419, 3, 2, 2, 2, 425, 422, 3, 2, 2, 2, 426, 429, 3, 2,
	2, 2, 427, 425, 3, 2, 2, 2, 427, 428, 3, 2, 2, 2, 428, 77, 3, 2, 2, 2,
	429, 427, 3, 2, 2, 2, 430, 431, 7, 18, 2, 2, 431, 438, 5, 80, 41, 2, 432,
	433, 7, 39, 2, 2, 433, 438, 5, 80, 41, 2, 434, 435, 7, 40, 2, 2, 435, 438,
	5, 80, 41, 2, 436, 438, 5, 80, 41, 2, 437, 430, 3, 2, 2, 2, 437, 432, 3,
	2, 2, 2, 437, 434, 3, 2, 2, 2, 437, 436, 3, 2, 2, 2, 438, 79, 3, 2, 2,
	2, 439, 440, 8, 41, 1, 2, 440, 451, 7, 60, 2, 2, 441, 451, 5, 88, 45, 2,
	442, 451, 5, 92, 47, 2, 443, 451, 5, 106, 54, 2, 444, 451, 5, 82, 42, 2,
	445, 451, 7, 61, 2, 2, 446, 447, 7, 49, 2, 2, 447, 448, 5, 70, 36, 2, 448,
	449, 7, 50, 2, 2, 449, 451, 3, 2, 2, 2, 450, 439, 3, 2, 2, 2, 450, 441,
	3, 2, 2, 2, 450, 442, 3, 2, 2, 2, 450, 443, 3, 2, 2, 2, 450, 444, 3, 2,
	2, 2, 450, 445, 3, 2, 2, 2, 450, 446, 3, 2, 2, 2, 451, 481, 3, 2, 2, 2,
	452, 453, 12, 12, 2, 2, 453, 454, 7, 48, 2, 2, 454, 480, 7, 60, 2, 2, 455,
	456, 12, 11, 2, 2, 456, 457, 7, 51, 2, 2, 457, 458, 5, 70, 36, 2, 458,
	459, 7, 52, 2, 2, 459, 480, 3, 2, 2, 2, 460, 461, 12, 10, 2, 2, 461, 463,
	7, 51, 2, 2, 462, 464, 5, 70, 36, 2, 463, 462, 3, 2, 2, 2, 463, 464, 3,
	2, 2, 2, 464, 465, 3, 2, 2, 2, 465, 467, 7, 47, 2, 2, 466, 468, 5, 70,
	36, 2, 467, 466, 3, 2, 2, 2, 467, 468, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2,
	469, 480, 7, 52, 2, 2, 470, 471, 12, 9, 2, 2, 471, 476, 7, 49, 2, 2, 472,
	474, 5, 84, 43, 2, 473, 475, 7, 58, 2, 2, 474, 473, 3, 2, 2, 2, 474, 475,
	3, 2, 2, 2, 475, 477, 3, 2, 2, 2, 476, 472, 3, 2, 2, 2, 476, 477, 3, 2,
	2, 2, 477, 478, 3, 2, 2, 2, 478, 480, 7, 50, 2, 2, 479, 452, 3, 2, 2, 2,
	479, 455, 3, 2, 2, 2, 479, 460, 3, 2, 2, 2, 479, 470, 3, 2, 2, 2, 480,
	483, 3, 2, 2, 2, 481, 479, 3, 2, 2, 2, 481, 482, 3, 2, 2, 2, 482, 81, 3,
	2, 2, 2, 483, 481, 3, 2, 2, 2, 484, 485, 9, 3, 2, 2, 485, 83, 3, 2, 2,
	2, 486, 491, 5, 70, 36, 2, 487, 488, 7, 46, 2, 2, 488, 490, 5, 70, 36,
	2, 489, 487, 3, 2, 2, 2, 490, 493, 3, 2, 2, 2, 491, 489, 3, 2, 2, 2, 491,
	492, 3, 2, 2, 2, 492, 495, 3, 2, 2, 2, 493, 491, 3, 2, 2, 2, 494, 496,
	9, 2, 2, 2, 495, 494, 3, 2, 2, 2, 495, 496, 3, 2, 2, 2, 496, 85, 3, 2,
	2, 2, 497, 508, 7, 60, 2, 2, 498, 499, 5, 80, 41, 2, 499, 500, 7, 48, 2,
	2, 500, 501, 7, 60, 2, 2, 501, 508, 3, 2, 2, 2, 502, 503, 5, 80, 41, 2,
	503, 504, 7, 51, 2, 2, 504, 505, 5, 70, 36, 2, 505, 506, 7, 52, 2, 2, 506,
	508, 3, 2, 2, 2, 507, 497, 3, 2, 2, 2, 507, 498, 3, 2, 2, 2, 507, 502,
	3, 2, 2, 2, 508, 87, 3, 2, 2, 2, 509, 510, 7, 14, 2, 2, 510, 512, 7, 49,
	2, 2, 511, 513, 5, 54, 28, 2, 512, 511, 3, 2, 2, 2, 512, 513, 3, 2, 2,
	2, 513, 514, 3, 2, 2, 2, 514, 515, 7, 50, 2, 2, 515, 516, 7, 53, 2, 2,
	516, 517, 5, 22, 12, 2, 517, 518, 7, 54, 2, 2, 518, 89, 3, 2, 2, 2, 519,
	521, 7, 56, 2, 2, 520, 522, 5, 54, 28, 2, 521, 520, 3, 2, 2, 2, 521, 522,
	3, 2, 2, 2, 522, 523, 3, 2, 2, 2, 523, 524, 7, 55, 2, 2, 524, 525, 5, 74,
	38, 2, 525, 91, 3, 2, 2, 2, 526, 527, 7, 53, 2, 2, 527, 528, 5, 94, 48,
	2, 528, 529, 7, 54, 2, 2, 529, 93, 3, 2, 2, 2, 530, 535, 5, 96, 49, 2,
	531, 532, 9, 2, 2, 2, 532, 534, 5, 96, 49, 2, 533, 531, 3, 2, 2, 2, 534,
	537, 3, 2, 2, 2, 535, 533, 3, 2, 2, 2, 535, 536, 3, 2, 2, 2, 536, 539,
	3, 2, 2, 2, 537, 535, 3, 2, 2, 2, 538, 540, 9, 2, 2, 2, 539, 538, 3, 2,
	2, 2, 539, 540, 3, 2, 2, 2, 540, 542, 3, 2, 2, 2, 541, 530, 3, 2, 2, 2,
	541, 542, 3, 2, 2, 2, 542, 95, 3, 2, 2, 2, 543, 544, 5, 120, 61, 2, 544,
	545, 7, 47, 2, 2, 545, 546, 5, 70, 36, 2, 546, 559, 3, 2, 2, 2, 547, 548,
	7, 51, 2, 2, 548, 549, 5, 70, 36, 2, 549, 550, 7, 52, 2, 2, 550, 551, 7,
	47, 2, 2, 551, 552, 5, 70, 36, 2, 552, 559, 3, 2, 2, 2, 553, 554, 5, 80,
	41, 2, 554, 555, 7, 58, 2, 2, 555, 559, 3, 2, 2, 2, 556, 559, 5, 98, 50,
	2, 557, 559, 5, 104, 53, 2, 558, 543, 3, 2, 2, 2, 558, 547, 3, 2, 2, 2,
	558, 553, 3, 2, 2, 2, 558, 556, 3, 2, 2, 2, 558, 557, 3, 2, 2, 2, 559,
	97, 3, 2, 2, 2, 560, 561, 7, 15, 2, 2, 561, 562, 5, 70, 36, 2, 562, 563,
	7, 53, 2, 2, 563, 564, 5, 94, 48, 2, 564, 568, 7, 54, 2, 2, 565, 567, 5,
	100, 51, 2, 566, 565, 3, 2, 2, 2, 567, 570, 3, 2, 2, 2, 568, 566, 3, 2,
	2, 2, 568, 569, 3, 2, 2, 2, 569, 572, 3, 2, 2, 2, 570, 568, 3, 2, 2, 2,
	571, 573, 5, 102, 52, 2, 572, 571, 3, 2, 2, 2, 572, 573, 3, 2, 2, 2, 573,
	99, 3, 2, 2, 2, 574, 575, 7, 11, 2, 2, 575, 576, 7, 15, 2, 2, 576, 577,
	5, 70, 36, 2, 577, 578, 7, 53, 2, 2, 578, 579, 5, 94, 48, 2, 579, 580,
	7, 54, 2, 2, 580, 101, 3, 2, 2, 2, 581, 582, 7, 11, 2, 2, 582, 583, 7,
	53, 2, 2, 583, 584, 5, 94, 48, 2, 584, 585, 7, 54, 2, 2, 585, 103, 3, 2,
	2, 2, 586, 587, 7, 13, 2, 2, 587, 588, 5, 42, 22, 2, 588, 589, 7, 60, 2,
	2, 589, 590, 5, 70, 36, 2, 590, 591, 7, 53, 2, 2, 591, 592, 5, 94, 48,
	2, 592, 593, 7, 54, 2, 2, 593, 105, 3, 2, 2, 2, 594, 595, 7, 51, 2, 2,
	595, 596, 5, 108, 55, 2, 596, 597, 7, 52, 2, 2, 597, 107, 3, 2, 2, 2, 598,
	603, 5, 110, 56, 2, 599, 600, 9, 2, 2, 2, 600, 602, 5, 110, 56, 2, 601,
	599, 3, 2, 2, 2, 602, 605, 3, 2, 2, 2, 603, 601, 3, 2, 2, 2, 603, 604,
	3, 2, 2, 2, 604, 609, 3, 2, 2, 2, 605, 603, 3, 2, 2, 2, 606, 608, 9, 2,
	2, 2, 607, 606, 3, 2, 2, 2, 608, 611, 3, 2, 2, 2, 609, 607, 3, 2, 2, 2,
	609, 610, 3, 2, 2, 2, 610, 613, 3, 2, 2, 2, 611, 609, 3, 2, 2, 2, 612,
	598, 3, 2, 2, 2, 612, 613, 3, 2, 2, 2, 613, 109, 3, 2, 2, 2, 614, 618,
	5, 70, 36, 2, 615, 618, 5, 112, 57, 2, 616, 618, 5, 118, 60, 2, 617, 614,
	3, 2, 2, 2, 617, 615, 3, 2, 2, 2, 617, 616, 3, 2, 2, 2, 618, 111, 3, 2,
	2, 2, 619, 620, 7, 15, 2, 2, 620, 621, 5, 70, 36, 2, 621, 622, 7, 53, 2,
	2, 622, 623, 5, 108, 55, 2, 623, 627, 7, 54, 2, 2, 624, 626, 5, 114, 58,
	2, 625, 624, 3, 2, 2, 2, 626, 629, 3, 2, 2, 2, 627, 625, 3, 2, 2, 2, 627,
	628, 3, 2, 2, 2, 628, 631, 3, 2, 2, 2, 629, 627, 3, 2, 2, 2, 630, 632,
	5, 116, 59, 2, 631, 630, 3, 2, 2, 2, 631, 632, 3, 2, 2, 2, 632, 113, 3,
	2, 2, 2, 633, 634, 7, 11, 2, 2, 634, 635, 7, 15, 2, 2, 635, 636, 5, 70,
	36, 2, 636, 637, 7, 53, 2, 2, 637, 638, 5, 108, 55, 2, 638, 639, 7, 54,
	2, 2, 639, 115, 3, 2, 2, 2, 640, 641, 7, 11, 2, 2, 641, 642, 7, 53, 2,
	2, 642, 643, 5, 108, 55, 2, 643, 644, 7, 54, 2, 2, 644, 117, 3, 2, 2, 2,
	645, 646, 7, 13, 2, 2, 646, 647, 5, 42, 22, 2, 647, 648, 7, 60, 2, 2, 648,
	649, 5, 70, 36, 2, 649, 650, 7, 53, 2, 2, 650, 651, 5, 108, 55, 2, 651,
	652, 7, 54, 2, 2, 652, 119, 3, 2, 2, 2, 653, 654, 9, 8, 2, 2, 654, 121,
	3, 2, 2, 2, 60, 128, 134, 142, 149, 157, 163, 171, 177, 186, 190, 195,
	201, 207, 212, 219, 224, 242, 253, 261, 272, 279, 295, 312, 316, 334, 346,
	351, 359, 386, 391, 404, 425, 427, 437, 450, 463, 467, 474, 476, 479, 481,
	491, 495, 507, 512, 521, 535, 539, 541, 558, 568, 572, 603, 609, 612, 617,
	627, 631,
}
var literalNames = []string{
	"", "'!info'", "'!param'", "'!flag'", "'and'", "'break'", "'catch'", "'continue'",
	"'defer'", "'else'", "'false'", "'for'", "'func'", "'if'", "'import'",
	"'nil'", "'not'", "'or'", "'return'", "'throw'", "'true'", "'try'", "'var'",
	"'while'", "'yield'", "'='", "'+='", "'-='", "'*='", "'/='", "'%='", "'=='",
	"'!='", "'<'", "'<='", "'>'", "'>='", "'+'", "'-'", "'*'", "'/'", "'%'",
	"'?'", "';'", "','", "':'", "'.'", "'('", "')'", "'['", "']'", "'{'", "'}'",
	"'->'", "'&'", "'|'", "'...'",
}
var symbolicNames = []string{
	"", "M_INFO", "M_PARAM", "M_FLAG", "AND", "BREAK", "CATCH", "CONTINUE",
	"DEFER", "ELSE", "FALSE", "FOR", "FUNC", "IF", "IMPORT", "NIL", "NOT",
	"OR", "RETURN", "THROW", "TRUE", "TRY", "VAR", "WHILE", "YIELD", "ASSIGN",
	"ASSIGN_ADD", "ASSIGN_SUB", "ASSIGN_MUL", "ASSIGN_DIV", "ASSIGN_MOD", "EQ",
	"NE", "LT", "LE", "GT", "GE", "ADD", "SUB", "MUL", "DIV", "MOD", "QUESTION_MARK",
	"SEMICOLON", "COMMA", "COLON", "PERIOD", "OPAREN", "CPAREN", "OBRACKET",
	"CBRACKET", "OCURLY", "CCURLY", "ARROW", "LAMBDA", "PIPE", "EXPAND", "NUMBER",
	"ID", "REGEX", "NEWLINE", "CHAR", "COMMENT", "WS", "STRING", "LDQUOTE",
}

var ruleNames = []string{
	"start", "unit", "meta_directive", "meta_info", "meta_param", "meta_flag",
	"meta_attribs", "meta_attrib", "meta_literal", "import_stmt", "stmts",
	"stmt_list", "stmt", "assignment_stmt", "assignment_lvalues", "rvalues",
	"assignment_op_stmt", "var_decl_stmt", "var_decl_vars", "for_stmt", "for_vars",
	"while_stmt", "if_stmt", "if_elif", "if_else", "func_stmt", "param_list",
	"return_stmt", "try_catch_stmt", "throw_stmt", "defer_stmt", "yield_stmt",
	"break_stmt", "continue_stmt", "expr", "expr2", "expr3", "binary_expr",
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
	NitroParserEOF           = antlr.TokenEOF
	NitroParserM_INFO        = 1
	NitroParserM_PARAM       = 2
	NitroParserM_FLAG        = 3
	NitroParserAND           = 4
	NitroParserBREAK         = 5
	NitroParserCATCH         = 6
	NitroParserCONTINUE      = 7
	NitroParserDEFER         = 8
	NitroParserELSE          = 9
	NitroParserFALSE         = 10
	NitroParserFOR           = 11
	NitroParserFUNC          = 12
	NitroParserIF            = 13
	NitroParserIMPORT        = 14
	NitroParserNIL           = 15
	NitroParserNOT           = 16
	NitroParserOR            = 17
	NitroParserRETURN        = 18
	NitroParserTHROW         = 19
	NitroParserTRUE          = 20
	NitroParserTRY           = 21
	NitroParserVAR           = 22
	NitroParserWHILE         = 23
	NitroParserYIELD         = 24
	NitroParserASSIGN        = 25
	NitroParserASSIGN_ADD    = 26
	NitroParserASSIGN_SUB    = 27
	NitroParserASSIGN_MUL    = 28
	NitroParserASSIGN_DIV    = 29
	NitroParserASSIGN_MOD    = 30
	NitroParserEQ            = 31
	NitroParserNE            = 32
	NitroParserLT            = 33
	NitroParserLE            = 34
	NitroParserGT            = 35
	NitroParserGE            = 36
	NitroParserADD           = 37
	NitroParserSUB           = 38
	NitroParserMUL           = 39
	NitroParserDIV           = 40
	NitroParserMOD           = 41
	NitroParserQUESTION_MARK = 42
	NitroParserSEMICOLON     = 43
	NitroParserCOMMA         = 44
	NitroParserCOLON         = 45
	NitroParserPERIOD        = 46
	NitroParserOPAREN        = 47
	NitroParserCPAREN        = 48
	NitroParserOBRACKET      = 49
	NitroParserCBRACKET      = 50
	NitroParserOCURLY        = 51
	NitroParserCCURLY        = 52
	NitroParserARROW         = 53
	NitroParserLAMBDA        = 54
	NitroParserPIPE          = 55
	NitroParserEXPAND        = 56
	NitroParserNUMBER        = 57
	NitroParserID            = 58
	NitroParserREGEX         = 59
	NitroParserNEWLINE       = 60
	NitroParserCHAR          = 61
	NitroParserCOMMENT       = 62
	NitroParserWS            = 63
	NitroParserSTRING        = 64
	NitroParserLDQUOTE       = 65
)

// NitroParser rules.
const (
	NitroParserRULE_start              = 0
	NitroParserRULE_unit               = 1
	NitroParserRULE_meta_directive     = 2
	NitroParserRULE_meta_info          = 3
	NitroParserRULE_meta_param         = 4
	NitroParserRULE_meta_flag          = 5
	NitroParserRULE_meta_attribs       = 6
	NitroParserRULE_meta_attrib        = 7
	NitroParserRULE_meta_literal       = 8
	NitroParserRULE_import_stmt        = 9
	NitroParserRULE_stmts              = 10
	NitroParserRULE_stmt_list          = 11
	NitroParserRULE_stmt               = 12
	NitroParserRULE_assignment_stmt    = 13
	NitroParserRULE_assignment_lvalues = 14
	NitroParserRULE_rvalues            = 15
	NitroParserRULE_assignment_op_stmt = 16
	NitroParserRULE_var_decl_stmt      = 17
	NitroParserRULE_var_decl_vars      = 18
	NitroParserRULE_for_stmt           = 19
	NitroParserRULE_for_vars           = 20
	NitroParserRULE_while_stmt         = 21
	NitroParserRULE_if_stmt            = 22
	NitroParserRULE_if_elif            = 23
	NitroParserRULE_if_else            = 24
	NitroParserRULE_func_stmt          = 25
	NitroParserRULE_param_list         = 26
	NitroParserRULE_return_stmt        = 27
	NitroParserRULE_try_catch_stmt     = 28
	NitroParserRULE_throw_stmt         = 29
	NitroParserRULE_defer_stmt         = 30
	NitroParserRULE_yield_stmt         = 31
	NitroParserRULE_break_stmt         = 32
	NitroParserRULE_continue_stmt      = 33
	NitroParserRULE_expr               = 34
	NitroParserRULE_expr2              = 35
	NitroParserRULE_expr3              = 36
	NitroParserRULE_binary_expr        = 37
	NitroParserRULE_unary_expr         = 38
	NitroParserRULE_primary_expr       = 39
	NitroParserRULE_simple_literal     = 40
	NitroParserRULE_arg_list           = 41
	NitroParserRULE_lvalue_expr        = 42
	NitroParserRULE_lambda_expr        = 43
	NitroParserRULE_short_lambda_expr  = 44
	NitroParserRULE_object_literal     = 45
	NitroParserRULE_object_fields      = 46
	NitroParserRULE_object_field       = 47
	NitroParserRULE_object_if          = 48
	NitroParserRULE_object_elif        = 49
	NitroParserRULE_object_else        = 50
	NitroParserRULE_object_for         = 51
	NitroParserRULE_array_literal      = 52
	NitroParserRULE_array_elems        = 53
	NitroParserRULE_array_elem         = 54
	NitroParserRULE_array_if           = 55
	NitroParserRULE_array_elif         = 56
	NitroParserRULE_array_else         = 57
	NitroParserRULE_array_for          = 58
	NitroParserRULE_id_or_keyword      = 59
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnitContext)(nil)).Elem(), 0)

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
		p.SetState(120)
		p.Unit()
	}
	{
		p.SetState(121)
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

func (s *UnitContext) Stmts() IStmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *UnitContext) AllMeta_directive() []IMeta_directiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMeta_directiveContext)(nil)).Elem())
	var tst = make([]IMeta_directiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMeta_directiveContext)
		}
	}

	return tst
}

func (s *UnitContext) Meta_directive(i int) IMeta_directiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_directiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMeta_directiveContext)
}

func (s *UnitContext) AllImport_stmt() []IImport_stmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImport_stmtContext)(nil)).Elem())
	var tst = make([]IImport_stmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImport_stmtContext)
		}
	}

	return tst
}

func (s *UnitContext) Import_stmt(i int) IImport_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImport_stmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImport_stmtContext)
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
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserM_INFO)|(1<<NitroParserM_PARAM)|(1<<NitroParserM_FLAG))) != 0 {
		{
			p.SetState(123)
			p.Meta_directive()
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserIMPORT {
		{
			p.SetState(129)
			p.Import_stmt()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(135)
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

func (s *Meta_directiveContext) Meta_info() IMeta_infoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_infoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeta_infoContext)
}

func (s *Meta_directiveContext) Meta_param() IMeta_paramContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_paramContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeta_paramContext)
}

func (s *Meta_directiveContext) Meta_flag() IMeta_flagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_flagContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 4, NitroParserRULE_meta_directive)

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

	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserM_INFO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Meta_info()
		}

	case NitroParserM_PARAM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Meta_param()
		}

	case NitroParserM_FLAG:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_attribsContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 6, NitroParserRULE_meta_info)
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
		p.Match(NitroParserM_INFO)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserOCURLY {
		{
			p.SetState(143)
			p.Match(NitroParserOCURLY)
		}
		{
			p.SetState(144)
			p.Meta_attribs()
		}
		{
			p.SetState(145)
			p.Match(NitroParserCCURLY)
		}

	}
	{
		p.SetState(149)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Meta_paramContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Meta_paramContext) Meta_attribs() IMeta_attribsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_attribsContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 8, NitroParserRULE_meta_param)
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
		p.SetState(151)
		p.Match(NitroParserM_PARAM)
	}
	{
		p.SetState(152)
		p.Match(NitroParserID)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserASSIGN {
		{
			p.SetState(153)
			p.Match(NitroParserASSIGN)
		}
		{
			p.SetState(154)
			p.expr(0)
		}

	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserOCURLY {
		{
			p.SetState(157)
			p.Match(NitroParserOCURLY)
		}
		{
			p.SetState(158)
			p.Meta_attribs()
		}
		{
			p.SetState(159)
			p.Match(NitroParserCCURLY)
		}

	}
	{
		p.SetState(163)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Meta_flagContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(NitroParserOCURLY, 0)
}

func (s *Meta_flagContext) Meta_attribs() IMeta_attribsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_attribsContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 10, NitroParserRULE_meta_flag)
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
		p.SetState(165)
		p.Match(NitroParserM_FLAG)
	}
	{
		p.SetState(166)
		p.Match(NitroParserID)
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserASSIGN {
		{
			p.SetState(167)
			p.Match(NitroParserASSIGN)
		}
		{
			p.SetState(168)
			p.expr(0)
		}

	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserOCURLY {
		{
			p.SetState(171)
			p.Match(NitroParserOCURLY)
		}
		{
			p.SetState(172)
			p.Meta_attribs()
		}
		{
			p.SetState(173)
			p.Match(NitroParserCCURLY)
		}

	}
	{
		p.SetState(177)
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
	this := p
	_ = this

	localctx = NewMeta_attribsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NitroParserRULE_meta_attribs)
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
		p.SetState(179)
		p.Meta_attrib()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(180)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(181)
				p.Meta_attrib()
			}

		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
		{
			p.SetState(187)
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
	this := p
	_ = this

	localctx = NewMeta_attribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NitroParserRULE_meta_attrib)
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
		p.Id_or_keyword()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserCOLON {
		{
			p.SetState(191)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(192)
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
	p.EnterRule(localctx, 16, NitroParserRULE_meta_literal)
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
		p.SetState(195)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Meta_literalContext).val = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserNIL)|(1<<NitroParserTRUE))) != 0) || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(NitroParserNUMBER-57))|(1<<(NitroParserCHAR-57))|(1<<(NitroParserSTRING-57)))) != 0)) {
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
	p.EnterRule(localctx, 18, NitroParserRULE_import_stmt)
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
		p.SetState(197)
		p.Match(NitroParserIMPORT)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(198)
			p.Match(NitroParserID)
		}

	}
	{
		p.SetState(201)
		p.Match(NitroParserSTRING)
	}
	{
		p.SetState(202)
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
	this := p
	_ = this

	localctx = NewStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NitroParserRULE_stmts)
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
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserBREAK)|(1<<NitroParserCONTINUE)|(1<<NitroParserDEFER)|(1<<NitroParserFALSE)|(1<<NitroParserFOR)|(1<<NitroParserFUNC)|(1<<NitroParserIF)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserRETURN)|(1<<NitroParserTHROW)|(1<<NitroParserTRUE)|(1<<NitroParserTRY)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserYIELD))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37)))) != 0) {
		{
			p.SetState(204)
			p.Stmt_list()
		}

	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserSEMICOLON {
		{
			p.SetState(207)
			p.Match(NitroParserSEMICOLON)
		}

		p.SetState(212)
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
	this := p
	_ = this

	localctx = NewStmt_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NitroParserRULE_stmt_list)
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
		p.SetState(213)
		p.Stmt()
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(215)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == NitroParserSEMICOLON {
				{
					p.SetState(214)
					p.Match(NitroParserSEMICOLON)
				}

				p.SetState(217)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(219)
				p.Stmt()
			}

		}
		p.SetState(224)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignment_op_stmtContext)(nil)).Elem(), 0)

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
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NitroParserRULE_stmt)

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

	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStmt_assignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.Assignment_stmt()
		}

	case 2:
		localctx = NewStmt_op_assignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.Assignment_op_stmt()
		}

	case 3:
		localctx = NewStmt_var_decContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(227)
			p.Var_decl_stmt()
		}

	case 4:
		localctx = NewStmt_forContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(228)
			p.For_stmt()
		}

	case 5:
		localctx = NewStmt_whileContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(229)
			p.While_stmt()
		}

	case 6:
		localctx = NewStmt_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(230)
			p.If_stmt()
		}

	case 7:
		localctx = NewStmt_funcContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(231)
			p.Func_stmt()
		}

	case 8:
		localctx = NewStmt_returnContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(232)
			p.Return_stmt()
		}

	case 9:
		localctx = NewStmt_exprContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(233)
			p.expr(0)
		}

	case 10:
		localctx = NewStmt_try_catchContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(234)
			p.Try_catch_stmt()
		}

	case 11:
		localctx = NewStmt_throwContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(235)
			p.Throw_stmt()
		}

	case 12:
		localctx = NewStmt_deferContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(236)
			p.Defer_stmt()
		}

	case 13:
		localctx = NewStmt_yieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(237)
			p.Yield_stmt()
		}

	case 14:
		localctx = NewStmt_breakContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(238)
			p.Break_stmt()
		}

	case 15:
		localctx = NewStmt_continueContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(239)
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
	this := p
	_ = this

	localctx = NewAssignment_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NitroParserRULE_assignment_stmt)

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
		p.SetState(242)
		p.Assignment_lvalues()
	}
	{
		p.SetState(243)
		p.Match(NitroParserASSIGN)
	}
	{
		p.SetState(244)
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
	this := p
	_ = this

	localctx = NewAssignment_lvaluesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NitroParserRULE_assignment_lvalues)
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
		p.Lvalue_expr()
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
			p.Lvalue_expr()
		}

		p.SetState(253)
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
	this := p
	_ = this

	localctx = NewRvaluesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NitroParserRULE_rvalues)
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
		p.expr(0)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(255)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(256)
			p.expr(0)
		}

		p.SetState(261)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILvalue_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILvalue_exprContext)
}

func (s *Assignment_op_stmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 32, NitroParserRULE_assignment_op_stmt)
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
		p.SetState(262)
		p.Lvalue_expr()
	}
	{
		p.SetState(263)

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
		p.SetState(264)
		p.expr(0)
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
	this := p
	_ = this

	localctx = NewVar_decl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NitroParserRULE_var_decl_stmt)
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
		p.Match(NitroParserVAR)
	}
	{
		p.SetState(267)
		p.Var_decl_vars()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserASSIGN {
		{
			p.SetState(268)
			p.Match(NitroParserASSIGN)
		}
		{
			p.SetState(269)
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
	p.EnterRule(localctx, 36, NitroParserRULE_var_decl_vars)
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
		p.SetState(272)
		p.Match(NitroParserID)
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(273)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(274)
			p.Match(NitroParserID)
		}

		p.SetState(279)
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
	this := p
	_ = this

	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, NitroParserRULE_for_stmt)

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
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(281)
		p.For_vars()
	}
	{
		p.SetState(282)
		p.Match(NitroParserID)
	}
	{
		p.SetState(283)
		p.expr(0)
	}
	{
		p.SetState(284)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(285)
		p.Stmts()
	}
	{
		p.SetState(286)
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
	p.EnterRule(localctx, 40, NitroParserRULE_for_vars)
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
		p.Match(NitroParserID)
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(289)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(290)
			p.Match(NitroParserID)
		}

		p.SetState(295)
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
	this := p
	_ = this

	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, NitroParserRULE_while_stmt)

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
		p.SetState(296)
		p.Match(NitroParserWHILE)
	}
	{
		p.SetState(297)
		p.expr(0)
	}
	{
		p.SetState(298)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(299)
		p.Stmts()
	}
	{
		p.SetState(300)
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
	this := p
	_ = this

	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, NitroParserRULE_if_stmt)
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
		p.SetState(302)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(303)
		p.expr(0)
	}
	{
		p.SetState(304)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(305)
		p.Stmts()
	}
	{
		p.SetState(306)
		p.Match(NitroParserCCURLY)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(307)
				p.If_elif()
			}

		}
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(313)
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
	this := p
	_ = this

	localctx = NewIf_elifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, NitroParserRULE_if_elif)

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
		p.SetState(316)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(317)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(318)
		p.expr(0)
	}
	{
		p.SetState(319)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(320)
		p.Stmts()
	}
	{
		p.SetState(321)
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
	this := p
	_ = this

	localctx = NewIf_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, NitroParserRULE_if_else)

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
		p.SetState(323)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(324)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(325)
		p.Stmts()
	}
	{
		p.SetState(326)
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
	this := p
	_ = this

	localctx = NewFunc_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, NitroParserRULE_func_stmt)
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
		p.SetState(328)
		p.Match(NitroParserFUNC)
	}
	{
		p.SetState(329)
		p.Match(NitroParserID)
	}
	{
		p.SetState(330)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(331)
			p.Param_list()
		}

	}
	{
		p.SetState(334)
		p.Match(NitroParserCPAREN)
	}
	{
		p.SetState(335)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(336)
		p.Stmts()
	}
	{
		p.SetState(337)
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
	p.EnterRule(localctx, 52, NitroParserRULE_param_list)
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
		p.Match(NitroParserID)
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NitroParserCOMMA {
		{
			p.SetState(340)
			p.Match(NitroParserCOMMA)
		}
		{
			p.SetState(341)
			p.Match(NitroParserID)
		}

		p.SetState(346)
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
	this := p
	_ = this

	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, NitroParserRULE_return_stmt)
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
		p.SetState(347)
		p.Match(NitroParserRETURN)
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37)))) != 0) {
		{
			p.SetState(348)
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
	this := p
	_ = this

	localctx = NewTry_catch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, NitroParserRULE_try_catch_stmt)
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
		p.SetState(351)
		p.Match(NitroParserTRY)
	}
	{
		p.SetState(352)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(353)
		p.Stmts()
	}
	{
		p.SetState(354)
		p.Match(NitroParserCCURLY)
	}
	{
		p.SetState(355)
		p.Match(NitroParserCATCH)
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(356)
			p.Match(NitroParserID)
		}

	}
	{
		p.SetState(359)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(360)
		p.Stmts()
	}
	{
		p.SetState(361)
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
	this := p
	_ = this

	localctx = NewThrow_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, NitroParserRULE_throw_stmt)

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
		p.SetState(363)
		p.Match(NitroParserTHROW)
	}
	{
		p.SetState(364)
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
	this := p
	_ = this

	localctx = NewDefer_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, NitroParserRULE_defer_stmt)

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
		p.Match(NitroParserDEFER)
	}
	{
		p.SetState(367)
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
	this := p
	_ = this

	localctx = NewYield_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, NitroParserRULE_yield_stmt)

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
		p.SetState(369)
		p.Match(NitroParserYIELD)
	}
	{
		p.SetState(370)
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
	p.EnterRule(localctx, 64, NitroParserRULE_break_stmt)

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
		p.SetState(372)
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
	p.EnterRule(localctx, 66, NitroParserRULE_continue_stmt)

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
		p.SetState(374)
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

func (s *ExprContext) Expr2() IExpr2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr2Context)
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
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, NitroParserRULE_expr, _p)

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
		p.SetState(377)
		p.Expr2()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_expr)
			p.SetState(379)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(380)
				p.Match(NitroParserPIPE)
			}
			{
				p.SetState(381)
				p.expr(3)
			}

		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

// IExpr2Context is an interface to support dynamic dispatch.
type IExpr2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr2Context differentiates from other interfaces.
	IsExpr2Context()
}

type Expr2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr2Context() *Expr2Context {
	var p = new(Expr2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NitroParserRULE_expr2
	return p
}

func (*Expr2Context) IsExpr2Context() {}

func NewExpr2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr2Context {
	var p = new(Expr2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NitroParserRULE_expr2

	return p
}

func (s *Expr2Context) GetParser() antlr.Parser { return s.parser }

func (s *Expr2Context) Short_lambda_expr() IShort_lambda_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShort_lambda_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShort_lambda_exprContext)
}

func (s *Expr2Context) Expr3() IExpr3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr3Context)
}

func (s *Expr2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.EnterExpr2(s)
	}
}

func (s *Expr2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NitroParserListener); ok {
		listenerT.ExitExpr2(s)
	}
}

func (p *NitroParser) Expr2() (localctx IExpr2Context) {
	this := p
	_ = this

	localctx = NewExpr2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, NitroParserRULE_expr2)

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

	p.SetState(389)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserLAMBDA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(387)
			p.Short_lambda_expr()
		}

	case NitroParserFALSE, NitroParserFUNC, NitroParserNIL, NitroParserNOT, NitroParserTRUE, NitroParserADD, NitroParserSUB, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserNUMBER, NitroParserID, NitroParserREGEX, NitroParserCHAR, NitroParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(388)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exprContext)
}

func (s *Expr3Context) AllExpr3() []IExpr3Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr3Context)(nil)).Elem())
	var tst = make([]IExpr3Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr3Context)
		}
	}

	return tst
}

func (s *Expr3Context) Expr3(i int) IExpr3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr3Context)(nil)).Elem(), i)

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
	_startState := 72
	p.EnterRecursionRule(localctx, 72, NitroParserRULE_expr3, _p)

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
		p.SetState(392)
		p.binary_expr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr3Context(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_expr3)
			p.SetState(394)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(395)
				p.Match(NitroParserQUESTION_MARK)
			}
			{
				p.SetState(396)
				p.expr3(0)
			}
			{
				p.SetState(397)
				p.Match(NitroParserCOLON)
			}
			{
				p.SetState(398)
				p.expr3(2)
			}

		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
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
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBinary_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBinary_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 74
	p.EnterRecursionRule(localctx, 74, NitroParserRULE_binary_expr, _p)
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
		p.SetState(406)
		p.Unary_expr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(423)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(408)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(409)

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
					p.SetState(410)
					p.binary_expr(6)
				}

			case 2:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(411)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(412)

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
					p.SetState(413)
					p.binary_expr(5)
				}

			case 3:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(414)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(415)

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
					p.SetState(416)
					p.binary_expr(4)
				}

			case 4:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(417)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(418)

					var _m = p.Match(NitroParserAND)

					localctx.(*Binary_exprContext).op = _m
				}
				{
					p.SetState(419)
					p.binary_expr(3)
				}

			case 5:
				localctx = NewBinary_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_binary_expr)
				p.SetState(420)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(421)

					var _m = p.Match(NitroParserOR)

					localctx.(*Binary_exprContext).op = _m
				}
				{
					p.SetState(422)
					p.binary_expr(2)
				}

			}

		}
		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 76, NitroParserRULE_unary_expr)

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

	p.SetState(435)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(428)

			var _m = p.Match(NitroParserNOT)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(429)
			p.primary_expr(0)
		}

	case NitroParserADD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(430)

			var _m = p.Match(NitroParserADD)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(431)
			p.primary_expr(0)
		}

	case NitroParserSUB:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(432)

			var _m = p.Match(NitroParserSUB)

			localctx.(*Unary_exprContext).op = _m
		}
		{
			p.SetState(433)
			p.primary_expr(0)
		}

	case NitroParserFALSE, NitroParserFUNC, NitroParserNIL, NitroParserTRUE, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserNUMBER, NitroParserID, NitroParserREGEX, NitroParserCHAR, NitroParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(434)
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
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimary_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimary_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 78
	p.EnterRecursionRule(localctx, 78, NitroParserRULE_primary_expr, _p)
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
	p.SetState(448)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserID:
		localctx = NewPrimary_expr_simple_refContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(438)
			p.Match(NitroParserID)
		}

	case NitroParserFUNC:
		localctx = NewPrimary_expr_lambdaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(439)
			p.Lambda_expr()
		}

	case NitroParserOCURLY:
		localctx = NewPrimary_expr_objectContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(440)
			p.Object_literal()
		}

	case NitroParserOBRACKET:
		localctx = NewPrimary_expr_arrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(441)
			p.Array_literal()
		}

	case NitroParserFALSE, NitroParserNIL, NitroParserTRUE, NitroParserNUMBER, NitroParserCHAR, NitroParserSTRING:
		localctx = NewPrimary_expr_literalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(442)
			p.Simple_literal()
		}

	case NitroParserREGEX:
		localctx = NewPrimary_expr_regexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(443)
			p.Match(NitroParserREGEX)
		}

	case NitroParserOPAREN:
		localctx = NewPrimary_expr_parenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(444)
			p.Match(NitroParserOPAREN)
		}
		{
			p.SetState(445)
			p.expr(0)
		}
		{
			p.SetState(446)
			p.Match(NitroParserCPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(477)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimary_expr_member_accessContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(450)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(451)
					p.Match(NitroParserPERIOD)
				}
				{
					p.SetState(452)
					p.Match(NitroParserID)
				}

			case 2:
				localctx = NewPrimary_expr_indexContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(453)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(454)
					p.Match(NitroParserOBRACKET)
				}
				{
					p.SetState(455)
					p.expr(0)
				}
				{
					p.SetState(456)
					p.Match(NitroParserCBRACKET)
				}

			case 3:
				localctx = NewPrimary_expr_sliceContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(458)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(459)
					p.Match(NitroParserOBRACKET)
				}
				p.SetState(461)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37)))) != 0) {
					{
						p.SetState(460)

						var _x = p.expr(0)

						localctx.(*Primary_expr_sliceContext).b = _x
					}

				}
				{
					p.SetState(463)
					p.Match(NitroParserCOLON)
				}
				p.SetState(465)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37)))) != 0) {
					{
						p.SetState(464)

						var _x = p.expr(0)

						localctx.(*Primary_expr_sliceContext).e = _x
					}

				}
				{
					p.SetState(467)
					p.Match(NitroParserCBRACKET)
				}

			case 4:
				localctx = NewPrimary_expr_callContext(p, NewPrimary_exprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, NitroParserRULE_primary_expr)
				p.SetState(468)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(469)
					p.Match(NitroParserOPAREN)
				}
				p.SetState(474)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFUNC)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37)))) != 0) {
					{
						p.SetState(470)
						p.Arg_list()
					}
					p.SetState(472)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if _la == NitroParserEXPAND {
						{
							p.SetState(471)
							p.Match(NitroParserEXPAND)
						}

					}

				}
				{
					p.SetState(476)
					p.Match(NitroParserCPAREN)
				}

			}

		}
		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 80, NitroParserRULE_simple_literal)
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
		p.SetState(482)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Simple_literalContext).val = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserNIL)|(1<<NitroParserTRUE))) != 0) || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(NitroParserNUMBER-57))|(1<<(NitroParserCHAR-57))|(1<<(NitroParserSTRING-57)))) != 0)) {
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
	this := p
	_ = this

	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, NitroParserRULE_arg_list)
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
		p.SetState(484)
		p.expr(0)
	}
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(485)
				p.Match(NitroParserCOMMA)
			}
			{
				p.SetState(486)
				p.expr(0)
			}

		}
		p.SetState(491)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}
	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
		{
			p.SetState(492)
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
	this := p
	_ = this

	localctx = NewLvalue_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, NitroParserRULE_lvalue_expr)

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

	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLvalue_expr_simple_refContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(495)
			p.Match(NitroParserID)
		}

	case 2:
		localctx = NewLvalue_expr_member_accessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(496)
			p.primary_expr(0)
		}
		{
			p.SetState(497)
			p.Match(NitroParserPERIOD)
		}
		{
			p.SetState(498)
			p.Match(NitroParserID)
		}

	case 3:
		localctx = NewLvalue_expr_indexContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(500)
			p.primary_expr(0)
		}
		{
			p.SetState(501)
			p.Match(NitroParserOBRACKET)
		}
		{
			p.SetState(502)
			p.expr(0)
		}
		{
			p.SetState(503)
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
	this := p
	_ = this

	localctx = NewLambda_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, NitroParserRULE_lambda_expr)
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
		p.Match(NitroParserFUNC)
	}
	{
		p.SetState(508)
		p.Match(NitroParserOPAREN)
	}
	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(509)
			p.Param_list()
		}

	}
	{
		p.SetState(512)
		p.Match(NitroParserCPAREN)
	}
	{
		p.SetState(513)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(514)
		p.Stmts()
	}
	{
		p.SetState(515)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr3Context)
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
	this := p
	_ = this

	localctx = NewShort_lambda_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, NitroParserRULE_short_lambda_expr)
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
		p.Match(NitroParserLAMBDA)
	}
	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserID {
		{
			p.SetState(518)
			p.Param_list()
		}

	}
	{
		p.SetState(521)
		p.Match(NitroParserARROW)
	}
	{
		p.SetState(522)
		p.expr3(0)
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
	this := p
	_ = this

	localctx = NewObject_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, NitroParserRULE_object_literal)

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
		p.SetState(524)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(525)
		p.Object_fields()
	}
	{
		p.SetState(526)
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
	this := p
	_ = this

	localctx = NewObject_fieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, NitroParserRULE_object_fields)
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
	p.SetState(539)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserAND)|(1<<NitroParserBREAK)|(1<<NitroParserCATCH)|(1<<NitroParserCONTINUE)|(1<<NitroParserDEFER)|(1<<NitroParserELSE)|(1<<NitroParserFALSE)|(1<<NitroParserFOR)|(1<<NitroParserFUNC)|(1<<NitroParserIF)|(1<<NitroParserIMPORT)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserOR)|(1<<NitroParserRETURN)|(1<<NitroParserTHROW)|(1<<NitroParserTRUE)|(1<<NitroParserTRY)|(1<<NitroParserVAR)|(1<<NitroParserWHILE)|(1<<NitroParserYIELD))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(NitroParserOPAREN-47))|(1<<(NitroParserOBRACKET-47))|(1<<(NitroParserOCURLY-47))|(1<<(NitroParserNUMBER-47))|(1<<(NitroParserID-47))|(1<<(NitroParserREGEX-47))|(1<<(NitroParserCHAR-47))|(1<<(NitroParserSTRING-47)))) != 0) {
		{
			p.SetState(528)
			p.Object_field()
		}
		p.SetState(533)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(529)
					_la = p.GetTokenStream().LA(1)

					if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(530)
					p.Object_field()
				}

			}
			p.SetState(535)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
		}
		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
			{
				p.SetState(536)
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
	this := p
	_ = this

	localctx = NewObject_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, NitroParserRULE_object_field)

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

	p.SetState(556)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		localctx = NewObject_field_id_keyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(541)
			p.Id_or_keyword()
		}
		{
			p.SetState(542)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(543)
			p.expr(0)
		}

	case 2:
		localctx = NewObject_field_expr_keyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(545)
			p.Match(NitroParserOBRACKET)
		}
		{
			p.SetState(546)
			p.expr(0)
		}
		{
			p.SetState(547)
			p.Match(NitroParserCBRACKET)
		}
		{
			p.SetState(548)
			p.Match(NitroParserCOLON)
		}
		{
			p.SetState(549)
			p.expr(0)
		}

	case 3:
		localctx = NewObject_field_expansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(551)
			p.primary_expr(0)
		}
		{
			p.SetState(552)
			p.Match(NitroParserEXPAND)
		}

	case 4:
		localctx = NewObject_field_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(554)
			p.Object_if()
		}

	case 5:
		localctx = NewObject_field_forContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(555)
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
	this := p
	_ = this

	localctx = NewObject_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, NitroParserRULE_object_if)
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
		p.SetState(558)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(559)
		p.expr(0)
	}
	{
		p.SetState(560)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(561)
		p.Object_fields()
	}
	{
		p.SetState(562)
		p.Match(NitroParserCCURLY)
	}
	p.SetState(566)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(563)
				p.Object_elif()
			}

		}
		p.SetState(568)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())
	}
	p.SetState(570)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(569)
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
	this := p
	_ = this

	localctx = NewObject_elifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, NitroParserRULE_object_elif)

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
		p.SetState(572)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(573)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(574)
		p.expr(0)
	}
	{
		p.SetState(575)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(576)
		p.Object_fields()
	}
	{
		p.SetState(577)
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
	this := p
	_ = this

	localctx = NewObject_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, NitroParserRULE_object_else)

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
		p.SetState(579)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(580)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(581)
		p.Object_fields()
	}
	{
		p.SetState(582)
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
	this := p
	_ = this

	localctx = NewObject_forContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, NitroParserRULE_object_for)

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
		p.SetState(584)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(585)
		p.For_vars()
	}
	{
		p.SetState(586)
		p.Match(NitroParserID)
	}
	{
		p.SetState(587)
		p.expr(0)
	}
	{
		p.SetState(588)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(589)
		p.Object_fields()
	}
	{
		p.SetState(590)
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
	this := p
	_ = this

	localctx = NewArray_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, NitroParserRULE_array_literal)

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
		p.SetState(592)
		p.Match(NitroParserOBRACKET)
	}
	{
		p.SetState(593)
		p.Array_elems()
	}
	{
		p.SetState(594)
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
	this := p
	_ = this

	localctx = NewArray_elemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, NitroParserRULE_array_elems)
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
	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NitroParserFALSE)|(1<<NitroParserFOR)|(1<<NitroParserFUNC)|(1<<NitroParserIF)|(1<<NitroParserNIL)|(1<<NitroParserNOT)|(1<<NitroParserTRUE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(NitroParserADD-37))|(1<<(NitroParserSUB-37))|(1<<(NitroParserOPAREN-37))|(1<<(NitroParserOBRACKET-37))|(1<<(NitroParserOCURLY-37))|(1<<(NitroParserLAMBDA-37))|(1<<(NitroParserNUMBER-37))|(1<<(NitroParserID-37))|(1<<(NitroParserREGEX-37))|(1<<(NitroParserCHAR-37))|(1<<(NitroParserSTRING-37)))) != 0) {
		{
			p.SetState(596)
			p.Array_elem()
		}
		p.SetState(601)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(597)
					_la = p.GetTokenStream().LA(1)

					if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(598)
					p.Array_elem()
				}

			}
			p.SetState(603)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())
		}
		p.SetState(607)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NitroParserSEMICOLON || _la == NitroParserCOMMA {
			{
				p.SetState(604)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NitroParserSEMICOLON || _la == NitroParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(609)
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
	this := p
	_ = this

	localctx = NewArray_elemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, NitroParserRULE_array_elem)

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

	p.SetState(615)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NitroParserFALSE, NitroParserFUNC, NitroParserNIL, NitroParserNOT, NitroParserTRUE, NitroParserADD, NitroParserSUB, NitroParserOPAREN, NitroParserOBRACKET, NitroParserOCURLY, NitroParserLAMBDA, NitroParserNUMBER, NitroParserID, NitroParserREGEX, NitroParserCHAR, NitroParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(612)
			p.expr(0)
		}

	case NitroParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(613)
			p.Array_if()
		}

	case NitroParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(614)
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
	this := p
	_ = this

	localctx = NewArray_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, NitroParserRULE_array_if)
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
		p.SetState(617)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(618)
		p.expr(0)
	}
	{
		p.SetState(619)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(620)
		p.Array_elems()
	}
	{
		p.SetState(621)
		p.Match(NitroParserCCURLY)
	}
	p.SetState(625)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(622)
				p.Array_elif()
			}

		}
		p.SetState(627)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())
	}
	p.SetState(629)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NitroParserELSE {
		{
			p.SetState(628)
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
	this := p
	_ = this

	localctx = NewArray_elifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, NitroParserRULE_array_elif)

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
		p.SetState(631)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(632)
		p.Match(NitroParserIF)
	}
	{
		p.SetState(633)
		p.expr(0)
	}
	{
		p.SetState(634)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(635)
		p.Array_elems()
	}
	{
		p.SetState(636)
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
	this := p
	_ = this

	localctx = NewArray_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, NitroParserRULE_array_else)

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
		p.SetState(638)
		p.Match(NitroParserELSE)
	}
	{
		p.SetState(639)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(640)
		p.Array_elems()
	}
	{
		p.SetState(641)
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
	this := p
	_ = this

	localctx = NewArray_forContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, NitroParserRULE_array_for)

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
		p.SetState(643)
		p.Match(NitroParserFOR)
	}
	{
		p.SetState(644)
		p.For_vars()
	}
	{
		p.SetState(645)
		p.Match(NitroParserID)
	}
	{
		p.SetState(646)
		p.expr(0)
	}
	{
		p.SetState(647)
		p.Match(NitroParserOCURLY)
	}
	{
		p.SetState(648)
		p.Array_elems()
	}
	{
		p.SetState(649)
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
	p.EnterRule(localctx, 118, NitroParserRULE_id_or_keyword)
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
		p.SetState(651)

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
	case 34:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 36:
		var t *Expr3Context = nil
		if localctx != nil {
			t = localctx.(*Expr3Context)
		}
		return p.Expr3_Sempred(t, predIndex)

	case 37:
		var t *Binary_exprContext = nil
		if localctx != nil {
			t = localctx.(*Binary_exprContext)
		}
		return p.Binary_expr_Sempred(t, predIndex)

	case 39:
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
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *NitroParser) Expr3_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *NitroParser) Binary_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

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
	this := p
	_ = this

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
