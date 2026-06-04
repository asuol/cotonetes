// Code generated from Latex.g4 by ANTLR 4.13.2. DO NOT EDIT.

package latex_parser // Latex
import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type LatexParser struct {
	*antlr.BaseParser
}

var LatexParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func latexParserInit() {
  staticData := &LatexParserStaticData
  staticData.LiteralNames = []string{
    "", "'\\textbf{Title:}'", "'\\\\'", "'\\textbf{URL:} \\url{'", "'}'", 
    "'\\textbf{Created:}'", "'\\textbf{Last Updated:}'", "'\\'", "'}{'", 
    "'{'", "'\\item'", "'\\begin{itemize}'", "'\\end{itemize}'", "'\\begin{enumerate}'", 
    "'\\end{enumerate}'", "'\\begin{verbatim}'", "'\\end{verbatim}'", "", 
    "", "", "", "'\\n'", "", "'\\r'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "LETTER", "PUNCTUATION", "SYMBOL", "NUMBER", "NEWLINE", "WS", "CR",
  }
  staticData.RuleNames = []string{
    "latex", "note_title", "note_url", "note_created", "note_updated", "note_text", 
    "text", "line_break", "empty_line", "escaped_word", "composite_tag_separator", 
    "tag", "word", "verbatim_content", "verbatim_line", "block_line", "block_item", 
    "block",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 23, 282, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 43, 
	8, 0, 10, 0, 12, 0, 46, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 53, 8, 
	1, 10, 1, 12, 1, 56, 9, 1, 1, 1, 4, 1, 59, 8, 1, 11, 1, 12, 1, 60, 1, 1, 
	1, 1, 3, 1, 65, 8, 1, 1, 2, 1, 2, 4, 2, 69, 8, 2, 11, 2, 12, 2, 70, 1, 
	2, 1, 2, 1, 2, 3, 2, 76, 8, 2, 1, 3, 1, 3, 5, 3, 80, 8, 3, 10, 3, 12, 3, 
	83, 9, 3, 1, 3, 4, 3, 86, 8, 3, 11, 3, 12, 3, 87, 1, 3, 1, 3, 3, 3, 92, 
	8, 3, 1, 4, 1, 4, 5, 4, 96, 8, 4, 10, 4, 12, 4, 99, 9, 4, 1, 4, 4, 4, 102, 
	8, 4, 11, 4, 12, 4, 103, 1, 4, 1, 4, 3, 4, 108, 8, 4, 1, 5, 1, 5, 1, 5, 
	1, 5, 5, 5, 114, 8, 5, 10, 5, 12, 5, 117, 9, 5, 1, 6, 1, 6, 4, 6, 121, 
	8, 6, 11, 6, 12, 6, 122, 1, 6, 3, 6, 126, 8, 6, 1, 7, 1, 7, 5, 7, 130, 
	8, 7, 10, 7, 12, 7, 133, 9, 7, 1, 8, 4, 8, 136, 8, 8, 11, 8, 12, 8, 137, 
	1, 9, 1, 9, 4, 9, 142, 8, 9, 11, 9, 12, 9, 143, 1, 9, 5, 9, 147, 8, 9, 
	10, 9, 12, 9, 150, 9, 9, 1, 9, 3, 9, 153, 8, 9, 1, 10, 1, 10, 1, 11, 1, 
	11, 4, 11, 159, 8, 11, 11, 11, 12, 11, 160, 1, 11, 1, 11, 4, 11, 165, 8, 
	11, 11, 11, 12, 11, 166, 1, 11, 1, 11, 1, 11, 1, 11, 4, 11, 173, 8, 11, 
	11, 11, 12, 11, 174, 1, 11, 1, 11, 4, 11, 179, 8, 11, 11, 11, 12, 11, 180, 
	1, 11, 1, 11, 4, 11, 185, 8, 11, 11, 11, 12, 11, 186, 1, 11, 1, 11, 3, 
	11, 191, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 198, 8, 12, 1, 
	13, 1, 13, 1, 13, 3, 13, 203, 8, 13, 1, 14, 5, 14, 206, 8, 14, 10, 14, 
	12, 14, 209, 9, 14, 1, 14, 1, 14, 1, 15, 4, 15, 214, 8, 15, 11, 15, 12, 
	15, 215, 1, 15, 4, 15, 219, 8, 15, 11, 15, 12, 15, 220, 1, 16, 1, 16, 5, 
	16, 225, 8, 16, 10, 16, 12, 16, 228, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 
	5, 17, 234, 8, 17, 10, 17, 12, 17, 237, 9, 17, 1, 17, 5, 17, 240, 8, 17, 
	10, 17, 12, 17, 243, 9, 17, 1, 17, 1, 17, 3, 17, 247, 8, 17, 1, 17, 1, 
	17, 5, 17, 251, 8, 17, 10, 17, 12, 17, 254, 9, 17, 1, 17, 5, 17, 257, 8, 
	17, 10, 17, 12, 17, 260, 9, 17, 1, 17, 1, 17, 3, 17, 264, 8, 17, 1, 17, 
	1, 17, 3, 17, 268, 8, 17, 1, 17, 5, 17, 271, 8, 17, 10, 17, 12, 17, 274, 
	9, 17, 1, 17, 1, 17, 3, 17, 278, 8, 17, 3, 17, 280, 8, 17, 1, 17, 0, 0, 
	18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 
	0, 1, 2, 0, 17, 17, 19, 19, 314, 0, 36, 1, 0, 0, 0, 2, 50, 1, 0, 0, 0, 
	4, 66, 1, 0, 0, 0, 6, 77, 1, 0, 0, 0, 8, 93, 1, 0, 0, 0, 10, 115, 1, 0, 
	0, 0, 12, 120, 1, 0, 0, 0, 14, 127, 1, 0, 0, 0, 16, 135, 1, 0, 0, 0, 18, 
	139, 1, 0, 0, 0, 20, 154, 1, 0, 0, 0, 22, 190, 1, 0, 0, 0, 24, 197, 1, 
	0, 0, 0, 26, 202, 1, 0, 0, 0, 28, 207, 1, 0, 0, 0, 30, 213, 1, 0, 0, 0, 
	32, 222, 1, 0, 0, 0, 34, 279, 1, 0, 0, 0, 36, 37, 3, 2, 1, 0, 37, 38, 3, 
	4, 2, 0, 38, 39, 3, 6, 3, 0, 39, 40, 3, 8, 4, 0, 40, 44, 3, 14, 7, 0, 41, 
	43, 5, 21, 0, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 
	0, 0, 44, 45, 1, 0, 0, 0, 45, 47, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 
	3, 10, 5, 0, 48, 49, 5, 0, 0, 1, 49, 1, 1, 0, 0, 0, 50, 54, 5, 1, 0, 0, 
	51, 53, 5, 22, 0, 0, 52, 51, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 
	0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 
	59, 3, 24, 12, 0, 58, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 58, 1, 0, 
	0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 64, 5, 2, 0, 0, 63, 65, 
	5, 21, 0, 0, 64, 63, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 3, 1, 0, 0, 0, 
	66, 68, 5, 3, 0, 0, 67, 69, 3, 24, 12, 0, 68, 67, 1, 0, 0, 0, 69, 70, 1, 
	0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 
	73, 5, 4, 0, 0, 73, 75, 5, 2, 0, 0, 74, 76, 5, 21, 0, 0, 75, 74, 1, 0, 
	0, 0, 75, 76, 1, 0, 0, 0, 76, 5, 1, 0, 0, 0, 77, 81, 5, 5, 0, 0, 78, 80, 
	5, 22, 0, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 
	81, 82, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 86, 3, 
	24, 12, 0, 85, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 
	87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 5, 2, 0, 0, 90, 92, 5, 
	21, 0, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 7, 1, 0, 0, 0, 93, 
	97, 5, 6, 0, 0, 94, 96, 5, 22, 0, 0, 95, 94, 1, 0, 0, 0, 96, 99, 1, 0, 
	0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 
	97, 1, 0, 0, 0, 100, 102, 3, 24, 12, 0, 101, 100, 1, 0, 0, 0, 102, 103, 
	1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0, 
	0, 0, 105, 107, 5, 2, 0, 0, 106, 108, 5, 21, 0, 0, 107, 106, 1, 0, 0, 0, 
	107, 108, 1, 0, 0, 0, 108, 9, 1, 0, 0, 0, 109, 114, 3, 12, 6, 0, 110, 114, 
	3, 34, 17, 0, 111, 114, 3, 14, 7, 0, 112, 114, 3, 16, 8, 0, 113, 109, 1, 
	0, 0, 0, 113, 110, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 112, 1, 0, 0, 
	0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 
	11, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 121, 3, 22, 11, 0, 119, 121, 
	3, 24, 12, 0, 120, 118, 1, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 122, 1, 
	0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 125, 1, 0, 0, 
	0, 124, 126, 5, 21, 0, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 
	13, 1, 0, 0, 0, 127, 131, 5, 2, 0, 0, 128, 130, 5, 22, 0, 0, 129, 128, 
	1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 
	0, 0, 132, 15, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 136, 5, 21, 0, 0, 
	135, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 
	138, 1, 0, 0, 0, 138, 17, 1, 0, 0, 0, 139, 141, 5, 7, 0, 0, 140, 142, 7, 
	0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0, 
	0, 143, 144, 1, 0, 0, 0, 144, 148, 1, 0, 0, 0, 145, 147, 5, 22, 0, 0, 146, 
	145, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 
	1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 153, 5, 21, 
	0, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 19, 1, 0, 0, 0, 
	154, 155, 5, 8, 0, 0, 155, 21, 1, 0, 0, 0, 156, 158, 5, 7, 0, 0, 157, 159, 
	5, 17, 0, 0, 158, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 158, 1, 0, 
	0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 164, 5, 9, 0, 0, 
	163, 165, 3, 24, 12, 0, 164, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 
	164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 
	5, 4, 0, 0, 169, 191, 1, 0, 0, 0, 170, 172, 5, 7, 0, 0, 171, 173, 5, 17, 
	0, 0, 172, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 
	174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 5, 9, 0, 0, 177, 
	179, 3, 24, 12, 0, 178, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 178, 
	1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184, 3, 20, 
	10, 0, 183, 185, 3, 24, 12, 0, 184, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 
	0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 
	189, 5, 4, 0, 0, 189, 191, 1, 0, 0, 0, 190, 156, 1, 0, 0, 0, 190, 170, 
	1, 0, 0, 0, 191, 23, 1, 0, 0, 0, 192, 198, 3, 18, 9, 0, 193, 198, 5, 17, 
	0, 0, 194, 198, 5, 18, 0, 0, 195, 198, 5, 20, 0, 0, 196, 198, 5, 22, 0, 
	0, 197, 192, 1, 0, 0, 0, 197, 193, 1, 0, 0, 0, 197, 194, 1, 0, 0, 0, 197, 
	195, 1, 0, 0, 0, 197, 196, 1, 0, 0, 0, 198, 25, 1, 0, 0, 0, 199, 203, 3, 
	24, 12, 0, 200, 203, 5, 19, 0, 0, 201, 203, 3, 14, 7, 0, 202, 199, 1, 0, 
	0, 0, 202, 200, 1, 0, 0, 0, 202, 201, 1, 0, 0, 0, 203, 27, 1, 0, 0, 0, 
	204, 206, 3, 26, 13, 0, 205, 204, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 
	205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 210, 1, 0, 0, 0, 209, 207, 
	1, 0, 0, 0, 210, 211, 5, 21, 0, 0, 211, 29, 1, 0, 0, 0, 212, 214, 3, 24, 
	12, 0, 213, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 
	215, 216, 1, 0, 0, 0, 216, 218, 1, 0, 0, 0, 217, 219, 5, 21, 0, 0, 218, 
	217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221, 
	1, 0, 0, 0, 221, 31, 1, 0, 0, 0, 222, 226, 5, 10, 0, 0, 223, 225, 5, 22, 
	0, 0, 224, 223, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 
	226, 227, 1, 0, 0, 0, 227, 229, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 
	230, 3, 30, 15, 0, 230, 33, 1, 0, 0, 0, 231, 235, 5, 11, 0, 0, 232, 234, 
	5, 21, 0, 0, 233, 232, 1, 0, 0, 0, 234, 237, 1, 0, 0, 0, 235, 233, 1, 0, 
	0, 0, 235, 236, 1, 0, 0, 0, 236, 241, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 
	238, 240, 3, 32, 16, 0, 239, 238, 1, 0, 0, 0, 240, 243, 1, 0, 0, 0, 241, 
	239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 244, 1, 0, 0, 0, 243, 241, 
	1, 0, 0, 0, 244, 246, 5, 12, 0, 0, 245, 247, 5, 21, 0, 0, 246, 245, 1, 
	0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 280, 1, 0, 0, 0, 248, 252, 5, 13, 0, 
	0, 249, 251, 5, 21, 0, 0, 250, 249, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 
	250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 258, 1, 0, 0, 0, 254, 252, 
	1, 0, 0, 0, 255, 257, 3, 32, 16, 0, 256, 255, 1, 0, 0, 0, 257, 260, 1, 
	0, 0, 0, 258, 256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 261, 1, 0, 0, 
	0, 260, 258, 1, 0, 0, 0, 261, 263, 5, 14, 0, 0, 262, 264, 5, 21, 0, 0, 
	263, 262, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 280, 1, 0, 0, 0, 265, 
	267, 5, 15, 0, 0, 266, 268, 5, 21, 0, 0, 267, 266, 1, 0, 0, 0, 267, 268, 
	1, 0, 0, 0, 268, 272, 1, 0, 0, 0, 269, 271, 3, 28, 14, 0, 270, 269, 1, 
	0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 273, 1, 0, 0, 
	0, 273, 275, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 275, 277, 5, 16, 0, 0, 276, 
	278, 5, 21, 0, 0, 277, 276, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 280, 
	1, 0, 0, 0, 279, 231, 1, 0, 0, 0, 279, 248, 1, 0, 0, 0, 279, 265, 1, 0, 
	0, 0, 280, 35, 1, 0, 0, 0, 44, 44, 54, 60, 64, 70, 75, 81, 87, 91, 97, 
	103, 107, 113, 115, 120, 122, 125, 131, 137, 143, 148, 152, 160, 166, 174, 
	180, 186, 190, 197, 202, 207, 215, 220, 226, 235, 241, 246, 252, 258, 263, 
	267, 272, 277, 279,
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

// LatexParserInit initializes any static state used to implement LatexParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLatexParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LatexParserInit() {
  staticData := &LatexParserStaticData
  staticData.once.Do(latexParserInit)
}

// NewLatexParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLatexParser(input antlr.TokenStream) *LatexParser {
	LatexParserInit()
	this := new(LatexParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &LatexParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Latex.g4"

	return this
}


// LatexParser tokens.
const (
	LatexParserEOF = antlr.TokenEOF
	LatexParserT__0 = 1
	LatexParserT__1 = 2
	LatexParserT__2 = 3
	LatexParserT__3 = 4
	LatexParserT__4 = 5
	LatexParserT__5 = 6
	LatexParserT__6 = 7
	LatexParserT__7 = 8
	LatexParserT__8 = 9
	LatexParserT__9 = 10
	LatexParserT__10 = 11
	LatexParserT__11 = 12
	LatexParserT__12 = 13
	LatexParserT__13 = 14
	LatexParserT__14 = 15
	LatexParserT__15 = 16
	LatexParserLETTER = 17
	LatexParserPUNCTUATION = 18
	LatexParserSYMBOL = 19
	LatexParserNUMBER = 20
	LatexParserNEWLINE = 21
	LatexParserWS = 22
	LatexParserCR = 23
)

// LatexParser rules.
const (
	LatexParserRULE_latex = 0
	LatexParserRULE_note_title = 1
	LatexParserRULE_note_url = 2
	LatexParserRULE_note_created = 3
	LatexParserRULE_note_updated = 4
	LatexParserRULE_note_text = 5
	LatexParserRULE_text = 6
	LatexParserRULE_line_break = 7
	LatexParserRULE_empty_line = 8
	LatexParserRULE_escaped_word = 9
	LatexParserRULE_composite_tag_separator = 10
	LatexParserRULE_tag = 11
	LatexParserRULE_word = 12
	LatexParserRULE_verbatim_content = 13
	LatexParserRULE_verbatim_line = 14
	LatexParserRULE_block_line = 15
	LatexParserRULE_block_item = 16
	LatexParserRULE_block = 17
)

// ILatexContext is an interface to support dynamic dispatch.
type ILatexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Note_title() INote_titleContext
	Note_url() INote_urlContext
	Note_created() INote_createdContext
	Note_updated() INote_updatedContext
	Line_break() ILine_breakContext
	Note_text() INote_textContext
	EOF() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsLatexContext differentiates from other interfaces.
	IsLatexContext()
}

type LatexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLatexContext() *LatexContext {
	var p = new(LatexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_latex
	return p
}

func InitEmptyLatexContext(p *LatexContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_latex
}

func (*LatexContext) IsLatexContext() {}

func NewLatexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LatexContext {
	var p = new(LatexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_latex

	return p
}

func (s *LatexContext) GetParser() antlr.Parser { return s.parser }

func (s *LatexContext) Note_title() INote_titleContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INote_titleContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INote_titleContext)
}

func (s *LatexContext) Note_url() INote_urlContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INote_urlContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INote_urlContext)
}

func (s *LatexContext) Note_created() INote_createdContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INote_createdContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INote_createdContext)
}

func (s *LatexContext) Note_updated() INote_updatedContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INote_updatedContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INote_updatedContext)
}

func (s *LatexContext) Line_break() ILine_breakContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_breakContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_breakContext)
}

func (s *LatexContext) Note_text() INote_textContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INote_textContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INote_textContext)
}

func (s *LatexContext) EOF() antlr.TerminalNode {
	return s.GetToken(LatexParserEOF, 0)
}

func (s *LatexContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(LatexParserNEWLINE)
}

func (s *LatexContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, i)
}

func (s *LatexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LatexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LatexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterLatex(s)
	}
}

func (s *LatexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitLatex(s)
	}
}




func (p *LatexParser) Latex() (localctx ILatexContext) {
	localctx = NewLatexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LatexParserRULE_latex)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Note_title()
	}
	{
		p.SetState(37)
		p.Note_url()
	}
	{
		p.SetState(38)
		p.Note_created()
	}
	{
		p.SetState(39)
		p.Note_updated()
	}
	{
		p.SetState(40)
		p.Line_break()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(41)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Note_text()
	}
	{
		p.SetState(48)
		p.Match(LatexParserEOF)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// INote_titleContext is an interface to support dynamic dispatch.
type INote_titleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllWord() []IWordContext
	Word(i int) IWordContext
	NEWLINE() antlr.TerminalNode

	// IsNote_titleContext differentiates from other interfaces.
	IsNote_titleContext()
}

type Note_titleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNote_titleContext() *Note_titleContext {
	var p = new(Note_titleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_note_title
	return p
}

func InitEmptyNote_titleContext(p *Note_titleContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_note_title
}

func (*Note_titleContext) IsNote_titleContext() {}

func NewNote_titleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Note_titleContext {
	var p = new(Note_titleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_note_title

	return p
}

func (s *Note_titleContext) GetParser() antlr.Parser { return s.parser }

func (s *Note_titleContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(LatexParserWS)
}

func (s *Note_titleContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserWS, i)
}

func (s *Note_titleContext) AllWord() []IWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordContext); ok {
			len++
		}
	}

	tst := make([]IWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordContext); ok {
			tst[i] = t.(IWordContext)
			i++
		}
	}

	return tst
}

func (s *Note_titleContext) Word(i int) IWordContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *Note_titleContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, 0)
}

func (s *Note_titleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Note_titleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Note_titleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterNote_title(s)
	}
}

func (s *Note_titleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitNote_title(s)
	}
}




func (p *LatexParser) Note_title() (localctx INote_titleContext) {
	localctx = NewNote_titleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LatexParserRULE_note_title)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(LatexParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(51)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 5636224) != 0) {
		{
			p.SetState(57)
			p.Word()
		}


		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(63)
			p.Match(LatexParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// INote_urlContext is an interface to support dynamic dispatch.
type INote_urlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWord() []IWordContext
	Word(i int) IWordContext
	NEWLINE() antlr.TerminalNode

	// IsNote_urlContext differentiates from other interfaces.
	IsNote_urlContext()
}

type Note_urlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNote_urlContext() *Note_urlContext {
	var p = new(Note_urlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_note_url
	return p
}

func InitEmptyNote_urlContext(p *Note_urlContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_note_url
}

func (*Note_urlContext) IsNote_urlContext() {}

func NewNote_urlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Note_urlContext {
	var p = new(Note_urlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_note_url

	return p
}

func (s *Note_urlContext) GetParser() antlr.Parser { return s.parser }

func (s *Note_urlContext) AllWord() []IWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordContext); ok {
			len++
		}
	}

	tst := make([]IWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordContext); ok {
			tst[i] = t.(IWordContext)
			i++
		}
	}

	return tst
}

func (s *Note_urlContext) Word(i int) IWordContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *Note_urlContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, 0)
}

func (s *Note_urlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Note_urlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Note_urlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterNote_url(s)
	}
}

func (s *Note_urlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitNote_url(s)
	}
}




func (p *LatexParser) Note_url() (localctx INote_urlContext) {
	localctx = NewNote_urlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LatexParserRULE_note_url)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(LatexParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 5636224) != 0) {
		{
			p.SetState(67)
			p.Word()
		}


		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(LatexParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(74)
			p.Match(LatexParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// INote_createdContext is an interface to support dynamic dispatch.
type INote_createdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllWord() []IWordContext
	Word(i int) IWordContext
	NEWLINE() antlr.TerminalNode

	// IsNote_createdContext differentiates from other interfaces.
	IsNote_createdContext()
}

type Note_createdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNote_createdContext() *Note_createdContext {
	var p = new(Note_createdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_note_created
	return p
}

func InitEmptyNote_createdContext(p *Note_createdContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_note_created
}

func (*Note_createdContext) IsNote_createdContext() {}

func NewNote_createdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Note_createdContext {
	var p = new(Note_createdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_note_created

	return p
}

func (s *Note_createdContext) GetParser() antlr.Parser { return s.parser }

func (s *Note_createdContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(LatexParserWS)
}

func (s *Note_createdContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserWS, i)
}

func (s *Note_createdContext) AllWord() []IWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordContext); ok {
			len++
		}
	}

	tst := make([]IWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordContext); ok {
			tst[i] = t.(IWordContext)
			i++
		}
	}

	return tst
}

func (s *Note_createdContext) Word(i int) IWordContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *Note_createdContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, 0)
}

func (s *Note_createdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Note_createdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Note_createdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterNote_created(s)
	}
}

func (s *Note_createdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitNote_created(s)
	}
}




func (p *LatexParser) Note_created() (localctx INote_createdContext) {
	localctx = NewNote_createdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LatexParserRULE_note_created)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(LatexParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(78)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 5636224) != 0) {
		{
			p.SetState(84)
			p.Word()
		}


		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(90)
			p.Match(LatexParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// INote_updatedContext is an interface to support dynamic dispatch.
type INote_updatedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllWord() []IWordContext
	Word(i int) IWordContext
	NEWLINE() antlr.TerminalNode

	// IsNote_updatedContext differentiates from other interfaces.
	IsNote_updatedContext()
}

type Note_updatedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNote_updatedContext() *Note_updatedContext {
	var p = new(Note_updatedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_note_updated
	return p
}

func InitEmptyNote_updatedContext(p *Note_updatedContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_note_updated
}

func (*Note_updatedContext) IsNote_updatedContext() {}

func NewNote_updatedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Note_updatedContext {
	var p = new(Note_updatedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_note_updated

	return p
}

func (s *Note_updatedContext) GetParser() antlr.Parser { return s.parser }

func (s *Note_updatedContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(LatexParserWS)
}

func (s *Note_updatedContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserWS, i)
}

func (s *Note_updatedContext) AllWord() []IWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordContext); ok {
			len++
		}
	}

	tst := make([]IWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordContext); ok {
			tst[i] = t.(IWordContext)
			i++
		}
	}

	return tst
}

func (s *Note_updatedContext) Word(i int) IWordContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *Note_updatedContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, 0)
}

func (s *Note_updatedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Note_updatedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Note_updatedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterNote_updated(s)
	}
}

func (s *Note_updatedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitNote_updated(s)
	}
}




func (p *LatexParser) Note_updated() (localctx INote_updatedContext) {
	localctx = NewNote_updatedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LatexParserRULE_note_updated)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(LatexParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(94)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 5636224) != 0) {
		{
			p.SetState(100)
			p.Word()
		}


		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(106)
			p.Match(LatexParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// INote_textContext is an interface to support dynamic dispatch.
type INote_textContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllText() []ITextContext
	Text(i int) ITextContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllLine_break() []ILine_breakContext
	Line_break(i int) ILine_breakContext
	AllEmpty_line() []IEmpty_lineContext
	Empty_line(i int) IEmpty_lineContext

	// IsNote_textContext differentiates from other interfaces.
	IsNote_textContext()
}

type Note_textContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNote_textContext() *Note_textContext {
	var p = new(Note_textContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_note_text
	return p
}

func InitEmptyNote_textContext(p *Note_textContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_note_text
}

func (*Note_textContext) IsNote_textContext() {}

func NewNote_textContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Note_textContext {
	var p = new(Note_textContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_note_text

	return p
}

func (s *Note_textContext) GetParser() antlr.Parser { return s.parser }

func (s *Note_textContext) AllText() []ITextContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITextContext); ok {
			len++
		}
	}

	tst := make([]ITextContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITextContext); ok {
			tst[i] = t.(ITextContext)
			i++
		}
	}

	return tst
}

func (s *Note_textContext) Text(i int) ITextContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *Note_textContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *Note_textContext) Block(i int) IBlockContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Note_textContext) AllLine_break() []ILine_breakContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILine_breakContext); ok {
			len++
		}
	}

	tst := make([]ILine_breakContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILine_breakContext); ok {
			tst[i] = t.(ILine_breakContext)
			i++
		}
	}

	return tst
}

func (s *Note_textContext) Line_break(i int) ILine_breakContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_breakContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_breakContext)
}

func (s *Note_textContext) AllEmpty_line() []IEmpty_lineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmpty_lineContext); ok {
			len++
		}
	}

	tst := make([]IEmpty_lineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmpty_lineContext); ok {
			tst[i] = t.(IEmpty_lineContext)
			i++
		}
	}

	return tst
}

func (s *Note_textContext) Empty_line(i int) IEmpty_lineContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmpty_lineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmpty_lineContext)
}

func (s *Note_textContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Note_textContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Note_textContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterNote_text(s)
	}
}

func (s *Note_textContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitNote_text(s)
	}
}




func (p *LatexParser) Note_text() (localctx INote_textContext) {
	localctx = NewNote_textContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LatexParserRULE_note_text)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 7776388) != 0) {
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LatexParserT__6, LatexParserLETTER, LatexParserPUNCTUATION, LatexParserNUMBER, LatexParserWS:
			{
				p.SetState(109)
				p.Text()
			}


		case LatexParserT__10, LatexParserT__12, LatexParserT__14:
			{
				p.SetState(110)
				p.Block()
			}


		case LatexParserT__1:
			{
				p.SetState(111)
				p.Line_break()
			}


		case LatexParserNEWLINE:
			{
				p.SetState(112)
				p.Empty_line()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTag() []ITagContext
	Tag(i int) ITagContext
	AllWord() []IWordContext
	Word(i int) IWordContext
	NEWLINE() antlr.TerminalNode

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_text
	return p
}

func InitEmptyTextContext(p *TextContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_text
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) AllTag() []ITagContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITagContext); ok {
			len++
		}
	}

	tst := make([]ITagContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITagContext); ok {
			tst[i] = t.(ITagContext)
			i++
		}
	}

	return tst
}

func (s *TextContext) Tag(i int) ITagContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *TextContext) AllWord() []IWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordContext); ok {
			len++
		}
	}

	tst := make([]IWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordContext); ok {
			tst[i] = t.(IWordContext)
			i++
		}
	}

	return tst
}

func (s *TextContext) Word(i int) IWordContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *TextContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, 0)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitText(s)
	}
}




func (p *LatexParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LatexParserRULE_text)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				p.SetState(120)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(118)
						p.Tag()
					}


				case 2:
					{
						p.SetState(119)
						p.Word()
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(124)
			p.Match(LatexParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILine_breakContext is an interface to support dynamic dispatch.
type ILine_breakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsLine_breakContext differentiates from other interfaces.
	IsLine_breakContext()
}

type Line_breakContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLine_breakContext() *Line_breakContext {
	var p = new(Line_breakContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_line_break
	return p
}

func InitEmptyLine_breakContext(p *Line_breakContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_line_break
}

func (*Line_breakContext) IsLine_breakContext() {}

func NewLine_breakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Line_breakContext {
	var p = new(Line_breakContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_line_break

	return p
}

func (s *Line_breakContext) GetParser() antlr.Parser { return s.parser }

func (s *Line_breakContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(LatexParserWS)
}

func (s *Line_breakContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserWS, i)
}

func (s *Line_breakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_breakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Line_breakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterLine_break(s)
	}
}

func (s *Line_breakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitLine_break(s)
	}
}




func (p *LatexParser) Line_break() (localctx ILine_breakContext) {
	localctx = NewLine_breakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LatexParserRULE_line_break)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(128)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEmpty_lineContext is an interface to support dynamic dispatch.
type IEmpty_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsEmpty_lineContext differentiates from other interfaces.
	IsEmpty_lineContext()
}

type Empty_lineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmpty_lineContext() *Empty_lineContext {
	var p = new(Empty_lineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_empty_line
	return p
}

func InitEmptyEmpty_lineContext(p *Empty_lineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_empty_line
}

func (*Empty_lineContext) IsEmpty_lineContext() {}

func NewEmpty_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Empty_lineContext {
	var p = new(Empty_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_empty_line

	return p
}

func (s *Empty_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Empty_lineContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(LatexParserNEWLINE)
}

func (s *Empty_lineContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, i)
}

func (s *Empty_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Empty_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Empty_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterEmpty_line(s)
	}
}

func (s *Empty_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitEmpty_line(s)
	}
}




func (p *LatexParser) Empty_line() (localctx IEmpty_lineContext) {
	localctx = NewEmpty_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LatexParserRULE_empty_line)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(134)
					p.Match(LatexParserNEWLINE)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}




		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEscaped_wordContext is an interface to support dynamic dispatch.
type IEscaped_wordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetContent returns the content token.
	GetContent() antlr.Token 

	// GetSpace returns the space token.
	GetSpace() antlr.Token 


	// SetContent sets the content token.
	SetContent(antlr.Token) 

	// SetSpace sets the space token.
	SetSpace(antlr.Token) 


	// Getter signatures
	NEWLINE() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllLETTER() []antlr.TerminalNode
	LETTER(i int) antlr.TerminalNode
	AllSYMBOL() []antlr.TerminalNode
	SYMBOL(i int) antlr.TerminalNode

	// IsEscaped_wordContext differentiates from other interfaces.
	IsEscaped_wordContext()
}

type Escaped_wordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	content antlr.Token
	space antlr.Token
}

func NewEmptyEscaped_wordContext() *Escaped_wordContext {
	var p = new(Escaped_wordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_escaped_word
	return p
}

func InitEmptyEscaped_wordContext(p *Escaped_wordContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_escaped_word
}

func (*Escaped_wordContext) IsEscaped_wordContext() {}

func NewEscaped_wordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Escaped_wordContext {
	var p = new(Escaped_wordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_escaped_word

	return p
}

func (s *Escaped_wordContext) GetParser() antlr.Parser { return s.parser }

func (s *Escaped_wordContext) GetContent() antlr.Token { return s.content }

func (s *Escaped_wordContext) GetSpace() antlr.Token { return s.space }


func (s *Escaped_wordContext) SetContent(v antlr.Token) { s.content = v }

func (s *Escaped_wordContext) SetSpace(v antlr.Token) { s.space = v }


func (s *Escaped_wordContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, 0)
}

func (s *Escaped_wordContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(LatexParserWS)
}

func (s *Escaped_wordContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserWS, i)
}

func (s *Escaped_wordContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(LatexParserLETTER)
}

func (s *Escaped_wordContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserLETTER, i)
}

func (s *Escaped_wordContext) AllSYMBOL() []antlr.TerminalNode {
	return s.GetTokens(LatexParserSYMBOL)
}

func (s *Escaped_wordContext) SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserSYMBOL, i)
}

func (s *Escaped_wordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Escaped_wordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Escaped_wordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterEscaped_word(s)
	}
}

func (s *Escaped_wordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitEscaped_word(s)
	}
}




func (p *LatexParser) Escaped_word() (localctx IEscaped_wordContext) {
	localctx = NewEscaped_wordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LatexParserRULE_escaped_word)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(LatexParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(140)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Escaped_wordContext).content = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LatexParserLETTER || _la == LatexParserSYMBOL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Escaped_wordContext).content = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}




		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(145)

				var _m = p.Match(LatexParserWS)

				localctx.(*Escaped_wordContext).space = _m
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(151)
			p.Match(LatexParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IComposite_tag_separatorContext is an interface to support dynamic dispatch.
type IComposite_tag_separatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComposite_tag_separatorContext differentiates from other interfaces.
	IsComposite_tag_separatorContext()
}

type Composite_tag_separatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComposite_tag_separatorContext() *Composite_tag_separatorContext {
	var p = new(Composite_tag_separatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_composite_tag_separator
	return p
}

func InitEmptyComposite_tag_separatorContext(p *Composite_tag_separatorContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_composite_tag_separator
}

func (*Composite_tag_separatorContext) IsComposite_tag_separatorContext() {}

func NewComposite_tag_separatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Composite_tag_separatorContext {
	var p = new(Composite_tag_separatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_composite_tag_separator

	return p
}

func (s *Composite_tag_separatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Composite_tag_separatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Composite_tag_separatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Composite_tag_separatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterComposite_tag_separator(s)
	}
}

func (s *Composite_tag_separatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitComposite_tag_separator(s)
	}
}




func (p *LatexParser) Composite_tag_separator() (localctx IComposite_tag_separatorContext) {
	localctx = NewComposite_tag_separatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LatexParserRULE_composite_tag_separator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(LatexParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_tag
	return p
}

func InitEmptyTagContext(p *TagContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_tag
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) CopyAll(ctx *TagContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type Composite_tagContext struct {
	TagContext
	name antlr.Token
}

func NewComposite_tagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Composite_tagContext {
	var p = new(Composite_tagContext)

	InitEmptyTagContext(&p.TagContext)
	p.parser = parser
	p.CopyAll(ctx.(*TagContext))

	return p
}


func (s *Composite_tagContext) GetName() antlr.Token { return s.name }


func (s *Composite_tagContext) SetName(v antlr.Token) { s.name = v }

func (s *Composite_tagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Composite_tagContext) Composite_tag_separator() IComposite_tag_separatorContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComposite_tag_separatorContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComposite_tag_separatorContext)
}

func (s *Composite_tagContext) AllWord() []IWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordContext); ok {
			len++
		}
	}

	tst := make([]IWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordContext); ok {
			tst[i] = t.(IWordContext)
			i++
		}
	}

	return tst
}

func (s *Composite_tagContext) Word(i int) IWordContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *Composite_tagContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(LatexParserLETTER)
}

func (s *Composite_tagContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserLETTER, i)
}


func (s *Composite_tagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterComposite_tag(s)
	}
}

func (s *Composite_tagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitComposite_tag(s)
	}
}


type Simple_tagContext struct {
	TagContext
	name antlr.Token
}

func NewSimple_tagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Simple_tagContext {
	var p = new(Simple_tagContext)

	InitEmptyTagContext(&p.TagContext)
	p.parser = parser
	p.CopyAll(ctx.(*TagContext))

	return p
}


func (s *Simple_tagContext) GetName() antlr.Token { return s.name }


func (s *Simple_tagContext) SetName(v antlr.Token) { s.name = v }

func (s *Simple_tagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_tagContext) AllWord() []IWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordContext); ok {
			len++
		}
	}

	tst := make([]IWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordContext); ok {
			tst[i] = t.(IWordContext)
			i++
		}
	}

	return tst
}

func (s *Simple_tagContext) Word(i int) IWordContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *Simple_tagContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(LatexParserLETTER)
}

func (s *Simple_tagContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserLETTER, i)
}


func (s *Simple_tagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterSimple_tag(s)
	}
}

func (s *Simple_tagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitSimple_tag(s)
	}
}



func (p *LatexParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LatexParserRULE_tag)
	var _la int

	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimple_tagContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Match(LatexParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == LatexParserLETTER {
			{
				p.SetState(157)

				var _m = p.Match(LatexParserLETTER)

				localctx.(*Simple_tagContext).name = _m
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(160)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(162)
			p.Match(LatexParserT__8)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 5636224) != 0) {
			{
				p.SetState(163)
				p.Word()
			}


			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(168)
			p.Match(LatexParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		localctx = NewComposite_tagContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.Match(LatexParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == LatexParserLETTER {
			{
				p.SetState(171)

				var _m = p.Match(LatexParserLETTER)

				localctx.(*Composite_tagContext).name = _m
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(176)
			p.Match(LatexParserT__8)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 5636224) != 0) {
			{
				p.SetState(177)
				p.Word()
			}


			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(182)
			p.Composite_tag_separator()
		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 5636224) != 0) {
			{
				p.SetState(183)
				p.Word()
			}


			p.SetState(186)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(188)
			p.Match(LatexParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IWordContext is an interface to support dynamic dispatch.
type IWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWordContext differentiates from other interfaces.
	IsWordContext()
}

type WordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordContext() *WordContext {
	var p = new(WordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_word
	return p
}

func InitEmptyWordContext(p *WordContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_word
}

func (*WordContext) IsWordContext() {}

func NewWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordContext {
	var p = new(WordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_word

	return p
}

func (s *WordContext) GetParser() antlr.Parser { return s.parser }

func (s *WordContext) CopyAll(ctx *WordContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *WordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type EscapedContext struct {
	WordContext
}

func NewEscapedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EscapedContext {
	var p = new(EscapedContext)

	InitEmptyWordContext(&p.WordContext)
	p.parser = parser
	p.CopyAll(ctx.(*WordContext))

	return p
}

func (s *EscapedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EscapedContext) Escaped_word() IEscaped_wordContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEscaped_wordContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEscaped_wordContext)
}


func (s *EscapedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterEscaped(s)
	}
}

func (s *EscapedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitEscaped(s)
	}
}


type NumberContext struct {
	WordContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	InitEmptyWordContext(&p.WordContext)
	p.parser = parser
	p.CopyAll(ctx.(*WordContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LatexParserNUMBER, 0)
}


func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitNumber(s)
	}
}


type LetterContext struct {
	WordContext
}

func NewLetterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetterContext {
	var p = new(LetterContext)

	InitEmptyWordContext(&p.WordContext)
	p.parser = parser
	p.CopyAll(ctx.(*WordContext))

	return p
}

func (s *LetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetterContext) LETTER() antlr.TerminalNode {
	return s.GetToken(LatexParserLETTER, 0)
}


func (s *LetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterLetter(s)
	}
}

func (s *LetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitLetter(s)
	}
}


type PunctuationContext struct {
	WordContext
}

func NewPunctuationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PunctuationContext {
	var p = new(PunctuationContext)

	InitEmptyWordContext(&p.WordContext)
	p.parser = parser
	p.CopyAll(ctx.(*WordContext))

	return p
}

func (s *PunctuationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PunctuationContext) PUNCTUATION() antlr.TerminalNode {
	return s.GetToken(LatexParserPUNCTUATION, 0)
}


func (s *PunctuationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterPunctuation(s)
	}
}

func (s *PunctuationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitPunctuation(s)
	}
}


type WsContext struct {
	WordContext
}

func NewWsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WsContext {
	var p = new(WsContext)

	InitEmptyWordContext(&p.WordContext)
	p.parser = parser
	p.CopyAll(ctx.(*WordContext))

	return p
}

func (s *WsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WsContext) WS() antlr.TerminalNode {
	return s.GetToken(LatexParserWS, 0)
}


func (s *WsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterWs(s)
	}
}

func (s *WsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitWs(s)
	}
}



func (p *LatexParser) Word() (localctx IWordContext) {
	localctx = NewWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LatexParserRULE_word)
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserT__6:
		localctx = NewEscapedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Escaped_word()
		}


	case LatexParserLETTER:
		localctx = NewLetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Match(LatexParserLETTER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LatexParserPUNCTUATION:
		localctx = NewPunctuationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.Match(LatexParserPUNCTUATION)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LatexParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(195)
			p.Match(LatexParserNUMBER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LatexParserWS:
		localctx = NewWsContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(196)
			p.Match(LatexParserWS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IVerbatim_contentContext is an interface to support dynamic dispatch.
type IVerbatim_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVerbatim_contentContext differentiates from other interfaces.
	IsVerbatim_contentContext()
}

type Verbatim_contentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerbatim_contentContext() *Verbatim_contentContext {
	var p = new(Verbatim_contentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_verbatim_content
	return p
}

func InitEmptyVerbatim_contentContext(p *Verbatim_contentContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_verbatim_content
}

func (*Verbatim_contentContext) IsVerbatim_contentContext() {}

func NewVerbatim_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Verbatim_contentContext {
	var p = new(Verbatim_contentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_verbatim_content

	return p
}

func (s *Verbatim_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Verbatim_contentContext) CopyAll(ctx *Verbatim_contentContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Verbatim_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Verbatim_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type Verbatim_wordContext struct {
	Verbatim_contentContext
}

func NewVerbatim_wordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Verbatim_wordContext {
	var p = new(Verbatim_wordContext)

	InitEmptyVerbatim_contentContext(&p.Verbatim_contentContext)
	p.parser = parser
	p.CopyAll(ctx.(*Verbatim_contentContext))

	return p
}

func (s *Verbatim_wordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Verbatim_wordContext) Word() IWordContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}


func (s *Verbatim_wordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterVerbatim_word(s)
	}
}

func (s *Verbatim_wordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitVerbatim_word(s)
	}
}


type Verbatim_symbolContext struct {
	Verbatim_contentContext
}

func NewVerbatim_symbolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Verbatim_symbolContext {
	var p = new(Verbatim_symbolContext)

	InitEmptyVerbatim_contentContext(&p.Verbatim_contentContext)
	p.parser = parser
	p.CopyAll(ctx.(*Verbatim_contentContext))

	return p
}

func (s *Verbatim_symbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Verbatim_symbolContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(LatexParserSYMBOL, 0)
}


func (s *Verbatim_symbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterVerbatim_symbol(s)
	}
}

func (s *Verbatim_symbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitVerbatim_symbol(s)
	}
}


type Verbatim_linebreakContext struct {
	Verbatim_contentContext
}

func NewVerbatim_linebreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Verbatim_linebreakContext {
	var p = new(Verbatim_linebreakContext)

	InitEmptyVerbatim_contentContext(&p.Verbatim_contentContext)
	p.parser = parser
	p.CopyAll(ctx.(*Verbatim_contentContext))

	return p
}

func (s *Verbatim_linebreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Verbatim_linebreakContext) Line_break() ILine_breakContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_breakContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_breakContext)
}


func (s *Verbatim_linebreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterVerbatim_linebreak(s)
	}
}

func (s *Verbatim_linebreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitVerbatim_linebreak(s)
	}
}



func (p *LatexParser) Verbatim_content() (localctx IVerbatim_contentContext) {
	localctx = NewVerbatim_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LatexParserRULE_verbatim_content)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserT__6, LatexParserLETTER, LatexParserPUNCTUATION, LatexParserNUMBER, LatexParserWS:
		localctx = NewVerbatim_wordContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Word()
		}


	case LatexParserSYMBOL:
		localctx = NewVerbatim_symbolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
			p.Match(LatexParserSYMBOL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LatexParserT__1:
		localctx = NewVerbatim_linebreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(201)
			p.Line_break()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IVerbatim_lineContext is an interface to support dynamic dispatch.
type IVerbatim_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEWLINE() antlr.TerminalNode
	AllVerbatim_content() []IVerbatim_contentContext
	Verbatim_content(i int) IVerbatim_contentContext

	// IsVerbatim_lineContext differentiates from other interfaces.
	IsVerbatim_lineContext()
}

type Verbatim_lineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerbatim_lineContext() *Verbatim_lineContext {
	var p = new(Verbatim_lineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_verbatim_line
	return p
}

func InitEmptyVerbatim_lineContext(p *Verbatim_lineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_verbatim_line
}

func (*Verbatim_lineContext) IsVerbatim_lineContext() {}

func NewVerbatim_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Verbatim_lineContext {
	var p = new(Verbatim_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_verbatim_line

	return p
}

func (s *Verbatim_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Verbatim_lineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, 0)
}

func (s *Verbatim_lineContext) AllVerbatim_content() []IVerbatim_contentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVerbatim_contentContext); ok {
			len++
		}
	}

	tst := make([]IVerbatim_contentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVerbatim_contentContext); ok {
			tst[i] = t.(IVerbatim_contentContext)
			i++
		}
	}

	return tst
}

func (s *Verbatim_lineContext) Verbatim_content(i int) IVerbatim_contentContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerbatim_contentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerbatim_contentContext)
}

func (s *Verbatim_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Verbatim_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Verbatim_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterVerbatim_line(s)
	}
}

func (s *Verbatim_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitVerbatim_line(s)
	}
}




func (p *LatexParser) Verbatim_line() (localctx IVerbatim_lineContext) {
	localctx = NewVerbatim_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LatexParserRULE_verbatim_line)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 6160516) != 0) {
		{
			p.SetState(204)
			p.Verbatim_content()
		}


		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(210)
		p.Match(LatexParserNEWLINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IBlock_lineContext is an interface to support dynamic dispatch.
type IBlock_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWord() []IWordContext
	Word(i int) IWordContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsBlock_lineContext differentiates from other interfaces.
	IsBlock_lineContext()
}

type Block_lineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_lineContext() *Block_lineContext {
	var p = new(Block_lineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_block_line
	return p
}

func InitEmptyBlock_lineContext(p *Block_lineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_block_line
}

func (*Block_lineContext) IsBlock_lineContext() {}

func NewBlock_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_lineContext {
	var p = new(Block_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_block_line

	return p
}

func (s *Block_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_lineContext) AllWord() []IWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordContext); ok {
			len++
		}
	}

	tst := make([]IWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordContext); ok {
			tst[i] = t.(IWordContext)
			i++
		}
	}

	return tst
}

func (s *Block_lineContext) Word(i int) IWordContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *Block_lineContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(LatexParserNEWLINE)
}

func (s *Block_lineContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, i)
}

func (s *Block_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Block_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterBlock_line(s)
	}
}

func (s *Block_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitBlock_line(s)
	}
}




func (p *LatexParser) Block_line() (localctx IBlock_lineContext) {
	localctx = NewBlock_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LatexParserRULE_block_line)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 5636224) != 0) {
		{
			p.SetState(212)
			p.Word()
		}


		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LatexParserNEWLINE {
		{
			p.SetState(217)
			p.Match(LatexParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IBlock_itemContext is an interface to support dynamic dispatch.
type IBlock_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block_line() IBlock_lineContext
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsBlock_itemContext differentiates from other interfaces.
	IsBlock_itemContext()
}

type Block_itemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_itemContext() *Block_itemContext {
	var p = new(Block_itemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_block_item
	return p
}

func InitEmptyBlock_itemContext(p *Block_itemContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_block_item
}

func (*Block_itemContext) IsBlock_itemContext() {}

func NewBlock_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_itemContext {
	var p = new(Block_itemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_block_item

	return p
}

func (s *Block_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_itemContext) Block_line() IBlock_lineContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_lineContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_lineContext)
}

func (s *Block_itemContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(LatexParserWS)
}

func (s *Block_itemContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserWS, i)
}

func (s *Block_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Block_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterBlock_item(s)
	}
}

func (s *Block_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitBlock_item(s)
	}
}




func (p *LatexParser) Block_item() (localctx IBlock_itemContext) {
	localctx = NewBlock_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LatexParserRULE_block_item)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(LatexParserT__9)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(223)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Block_line()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) CopyAll(ctx *BlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type ItemizeContext struct {
	BlockContext
}

func NewItemizeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ItemizeContext {
	var p = new(ItemizeContext)

	InitEmptyBlockContext(&p.BlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*BlockContext))

	return p
}

func (s *ItemizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemizeContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(LatexParserNEWLINE)
}

func (s *ItemizeContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, i)
}

func (s *ItemizeContext) AllBlock_item() []IBlock_itemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_itemContext); ok {
			len++
		}
	}

	tst := make([]IBlock_itemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_itemContext); ok {
			tst[i] = t.(IBlock_itemContext)
			i++
		}
	}

	return tst
}

func (s *ItemizeContext) Block_item(i int) IBlock_itemContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_itemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_itemContext)
}


func (s *ItemizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterItemize(s)
	}
}

func (s *ItemizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitItemize(s)
	}
}


type EnumerateContext struct {
	BlockContext
}

func NewEnumerateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumerateContext {
	var p = new(EnumerateContext)

	InitEmptyBlockContext(&p.BlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*BlockContext))

	return p
}

func (s *EnumerateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumerateContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(LatexParserNEWLINE)
}

func (s *EnumerateContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, i)
}

func (s *EnumerateContext) AllBlock_item() []IBlock_itemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_itemContext); ok {
			len++
		}
	}

	tst := make([]IBlock_itemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_itemContext); ok {
			tst[i] = t.(IBlock_itemContext)
			i++
		}
	}

	return tst
}

func (s *EnumerateContext) Block_item(i int) IBlock_itemContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_itemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_itemContext)
}


func (s *EnumerateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterEnumerate(s)
	}
}

func (s *EnumerateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitEnumerate(s)
	}
}


type VerbatimContext struct {
	BlockContext
}

func NewVerbatimContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VerbatimContext {
	var p = new(VerbatimContext)

	InitEmptyBlockContext(&p.BlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*BlockContext))

	return p
}

func (s *VerbatimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerbatimContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(LatexParserNEWLINE)
}

func (s *VerbatimContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserNEWLINE, i)
}

func (s *VerbatimContext) AllVerbatim_line() []IVerbatim_lineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVerbatim_lineContext); ok {
			len++
		}
	}

	tst := make([]IVerbatim_lineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVerbatim_lineContext); ok {
			tst[i] = t.(IVerbatim_lineContext)
			i++
		}
	}

	return tst
}

func (s *VerbatimContext) Verbatim_line(i int) IVerbatim_lineContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerbatim_lineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerbatim_lineContext)
}


func (s *VerbatimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterVerbatim(s)
	}
}

func (s *VerbatimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitVerbatim(s)
	}
}



func (p *LatexParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LatexParserRULE_block)
	var _la int

	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserT__10:
		localctx = NewItemizeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)
			p.Match(LatexParserT__10)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserNEWLINE {
			{
				p.SetState(232)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserT__9 {
			{
				p.SetState(238)
				p.Block_item()
			}


			p.SetState(243)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(244)
			p.Match(LatexParserT__11)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(245)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}


	case LatexParserT__12:
		localctx = NewEnumerateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(248)
			p.Match(LatexParserT__12)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserNEWLINE {
			{
				p.SetState(249)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserT__9 {
			{
				p.SetState(255)
				p.Block_item()
			}


			p.SetState(260)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(261)
			p.Match(LatexParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(262)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}


	case LatexParserT__14:
		localctx = NewVerbatimContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(265)
			p.Match(LatexParserT__14)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(266)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 8257668) != 0) {
			{
				p.SetState(269)
				p.Verbatim_line()
			}


			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(275)
			p.Match(LatexParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(276)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


