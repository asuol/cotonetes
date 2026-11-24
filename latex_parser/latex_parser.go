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
    "'\\textbf{Created:}'", "'\\textbf{Last Updated:}'", "'{'", "'}{'", 
    "'\\url{'", "'\\'", "'\\item'", "'\\begin{itemize}'", "'\\end{itemize}'", 
    "'\\begin{enumerate}'", "'\\end{enumerate}'", "'\\begin{verbatim}'", 
    "'\\end{verbatim}'", "", "", "", "", "'\\n'", "", "'\\r'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "LETTER", "SPECIAL", "SYMBOL", "NUMBER", "NEWLINE", "WS", "CR", 
    "PUNCTUATION", "EXTENDED_UNICODE",
  }
  staticData.RuleNames = []string{
    "latex", "note_title", "note_url", "note_created", "note_updated", "note_text", 
    "text", "line_break", "empty_line", "escaped_word", "composite_tag_separator", 
    "tag", "word", "verbatim_content", "curly_parens", "verbatim_block_content", 
    "verbatim_line", "block_content", "block_line", "block_item", "block",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 26, 322, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1, 
	0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 49, 8, 0, 10, 0, 12, 0, 52, 9, 0, 
	1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 59, 8, 1, 10, 1, 12, 1, 62, 9, 1, 1, 
	1, 4, 1, 65, 8, 1, 11, 1, 12, 1, 66, 1, 1, 1, 1, 3, 1, 71, 8, 1, 1, 2, 
	1, 2, 4, 2, 75, 8, 2, 11, 2, 12, 2, 76, 1, 2, 1, 2, 1, 2, 3, 2, 82, 8, 
	2, 1, 3, 1, 3, 5, 3, 86, 8, 3, 10, 3, 12, 3, 89, 9, 3, 1, 3, 4, 3, 92, 
	8, 3, 11, 3, 12, 3, 93, 1, 3, 1, 3, 3, 3, 98, 8, 3, 1, 4, 1, 4, 5, 4, 102, 
	8, 4, 10, 4, 12, 4, 105, 9, 4, 1, 4, 4, 4, 108, 8, 4, 11, 4, 12, 4, 109, 
	1, 4, 1, 4, 3, 4, 114, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 120, 8, 5, 10, 
	5, 12, 5, 123, 9, 5, 1, 6, 1, 6, 4, 6, 127, 8, 6, 11, 6, 12, 6, 128, 1, 
	6, 3, 6, 132, 8, 6, 1, 7, 1, 7, 5, 7, 136, 8, 7, 10, 7, 12, 7, 139, 9, 
	7, 1, 8, 4, 8, 142, 8, 8, 11, 8, 12, 8, 143, 1, 9, 4, 9, 147, 8, 9, 11, 
	9, 12, 9, 148, 1, 9, 5, 9, 152, 8, 9, 10, 9, 12, 9, 155, 9, 9, 1, 9, 1, 
	9, 5, 9, 159, 8, 9, 10, 9, 12, 9, 162, 9, 9, 3, 9, 164, 8, 9, 1, 10, 1, 
	10, 1, 11, 1, 11, 4, 11, 170, 8, 11, 11, 11, 12, 11, 171, 1, 11, 1, 11, 
	1, 11, 1, 11, 4, 11, 178, 8, 11, 11, 11, 12, 11, 179, 1, 11, 1, 11, 4, 
	11, 184, 8, 11, 11, 11, 12, 11, 185, 1, 11, 1, 11, 1, 11, 1, 11, 4, 11, 
	192, 8, 11, 11, 11, 12, 11, 193, 1, 11, 1, 11, 4, 11, 198, 8, 11, 11, 11, 
	12, 11, 199, 1, 11, 1, 11, 4, 11, 204, 8, 11, 11, 11, 12, 11, 205, 1, 11, 
	1, 11, 3, 11, 210, 8, 11, 1, 12, 1, 12, 3, 12, 214, 8, 12, 1, 12, 1, 12, 
	1, 12, 1, 12, 3, 12, 220, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 226, 
	8, 13, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15, 232, 8, 15, 1, 16, 5, 16, 235, 
	8, 16, 10, 16, 12, 16, 238, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 3, 
	17, 245, 8, 17, 1, 18, 4, 18, 248, 8, 18, 11, 18, 12, 18, 249, 1, 18, 4, 
	18, 253, 8, 18, 11, 18, 12, 18, 254, 1, 19, 5, 19, 258, 8, 19, 10, 19, 
	12, 19, 261, 9, 19, 1, 19, 1, 19, 5, 19, 265, 8, 19, 10, 19, 12, 19, 268, 
	9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 5, 20, 274, 8, 20, 10, 20, 12, 20, 277, 
	9, 20, 1, 20, 5, 20, 280, 8, 20, 10, 20, 12, 20, 283, 9, 20, 1, 20, 1, 
	20, 3, 20, 287, 8, 20, 1, 20, 1, 20, 5, 20, 291, 8, 20, 10, 20, 12, 20, 
	294, 9, 20, 1, 20, 5, 20, 297, 8, 20, 10, 20, 12, 20, 300, 9, 20, 1, 20, 
	1, 20, 3, 20, 304, 8, 20, 1, 20, 1, 20, 3, 20, 308, 8, 20, 1, 20, 5, 20, 
	311, 8, 20, 10, 20, 12, 20, 314, 9, 20, 1, 20, 1, 20, 3, 20, 318, 8, 20, 
	3, 20, 320, 8, 20, 1, 20, 0, 0, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 3, 3, 0, 4, 4, 7, 7, 20, 
	20, 1, 0, 25, 26, 2, 0, 4, 4, 7, 7, 360, 0, 42, 1, 0, 0, 0, 2, 56, 1, 0, 
	0, 0, 4, 72, 1, 0, 0, 0, 6, 83, 1, 0, 0, 0, 8, 99, 1, 0, 0, 0, 10, 121, 
	1, 0, 0, 0, 12, 126, 1, 0, 0, 0, 14, 133, 1, 0, 0, 0, 16, 141, 1, 0, 0, 
	0, 18, 163, 1, 0, 0, 0, 20, 165, 1, 0, 0, 0, 22, 209, 1, 0, 0, 0, 24, 219, 
	1, 0, 0, 0, 26, 225, 1, 0, 0, 0, 28, 227, 1, 0, 0, 0, 30, 231, 1, 0, 0, 
	0, 32, 236, 1, 0, 0, 0, 34, 244, 1, 0, 0, 0, 36, 247, 1, 0, 0, 0, 38, 259, 
	1, 0, 0, 0, 40, 319, 1, 0, 0, 0, 42, 43, 3, 2, 1, 0, 43, 44, 3, 4, 2, 0, 
	44, 45, 3, 6, 3, 0, 45, 46, 3, 8, 4, 0, 46, 50, 3, 14, 7, 0, 47, 49, 5, 
	22, 0, 0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 
	51, 1, 0, 0, 0, 51, 53, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 54, 3, 10, 
	5, 0, 54, 55, 5, 0, 0, 1, 55, 1, 1, 0, 0, 0, 56, 60, 5, 1, 0, 0, 57, 59, 
	5, 23, 0, 0, 58, 57, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 
	60, 61, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 65, 3, 
	24, 12, 0, 64, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 
	66, 67, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 5, 2, 0, 0, 69, 71, 5, 
	22, 0, 0, 70, 69, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 3, 1, 0, 0, 0, 72, 
	74, 5, 3, 0, 0, 73, 75, 3, 26, 13, 0, 74, 73, 1, 0, 0, 0, 75, 76, 1, 0, 
	0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 
	5, 4, 0, 0, 79, 81, 5, 2, 0, 0, 80, 82, 5, 22, 0, 0, 81, 80, 1, 0, 0, 0, 
	81, 82, 1, 0, 0, 0, 82, 5, 1, 0, 0, 0, 83, 87, 5, 5, 0, 0, 84, 86, 5, 23, 
	0, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 
	1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 92, 3, 24, 12, 
	0, 91, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 
	1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 97, 5, 2, 0, 0, 96, 98, 5, 22, 0, 0, 
	97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 7, 1, 0, 0, 0, 99, 103, 5, 
	6, 0, 0, 100, 102, 5, 23, 0, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 
	0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 
	103, 1, 0, 0, 0, 106, 108, 3, 24, 12, 0, 107, 106, 1, 0, 0, 0, 108, 109, 
	1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 
	0, 0, 111, 113, 5, 2, 0, 0, 112, 114, 5, 22, 0, 0, 113, 112, 1, 0, 0, 0, 
	113, 114, 1, 0, 0, 0, 114, 9, 1, 0, 0, 0, 115, 120, 3, 12, 6, 0, 116, 120, 
	3, 40, 20, 0, 117, 120, 3, 14, 7, 0, 118, 120, 3, 16, 8, 0, 119, 115, 1, 
	0, 0, 0, 119, 116, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 
	0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 
	11, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124, 127, 3, 22, 11, 0, 125, 127, 
	3, 24, 12, 0, 126, 124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 
	0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 1, 0, 0, 
	0, 130, 132, 5, 22, 0, 0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 
	13, 1, 0, 0, 0, 133, 137, 5, 2, 0, 0, 134, 136, 5, 23, 0, 0, 135, 134, 
	1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 
	0, 0, 138, 15, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 142, 5, 22, 0, 0, 
	141, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 
	144, 1, 0, 0, 0, 144, 17, 1, 0, 0, 0, 145, 147, 5, 18, 0, 0, 146, 145, 
	1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 
	0, 0, 149, 153, 1, 0, 0, 0, 150, 152, 5, 23, 0, 0, 151, 150, 1, 0, 0, 0, 
	152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 
	164, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 160, 7, 0, 0, 0, 157, 159, 
	5, 23, 0, 0, 158, 157, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1, 0, 
	0, 0, 160, 161, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 
	163, 146, 1, 0, 0, 0, 163, 156, 1, 0, 0, 0, 164, 19, 1, 0, 0, 0, 165, 166, 
	5, 8, 0, 0, 166, 21, 1, 0, 0, 0, 167, 169, 5, 9, 0, 0, 168, 170, 3, 26, 
	13, 0, 169, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 
	171, 172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 5, 4, 0, 0, 174, 
	210, 1, 0, 0, 0, 175, 177, 5, 10, 0, 0, 176, 178, 5, 18, 0, 0, 177, 176, 
	1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 
	0, 0, 180, 181, 1, 0, 0, 0, 181, 183, 5, 7, 0, 0, 182, 184, 3, 24, 12, 
	0, 183, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 
	186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 5, 4, 0, 0, 188, 210, 
	1, 0, 0, 0, 189, 191, 5, 10, 0, 0, 190, 192, 5, 18, 0, 0, 191, 190, 1, 
	0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 
	0, 194, 195, 1, 0, 0, 0, 195, 197, 5, 7, 0, 0, 196, 198, 3, 24, 12, 0, 
	197, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 
	200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 203, 3, 20, 10, 0, 202, 204, 
	3, 24, 12, 0, 203, 202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 203, 1, 
	0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 208, 5, 4, 0, 
	0, 208, 210, 1, 0, 0, 0, 209, 167, 1, 0, 0, 0, 209, 175, 1, 0, 0, 0, 209, 
	189, 1, 0, 0, 0, 210, 23, 1, 0, 0, 0, 211, 213, 5, 10, 0, 0, 212, 214, 
	3, 18, 9, 0, 213, 212, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 220, 1, 0, 
	0, 0, 215, 220, 5, 18, 0, 0, 216, 220, 7, 1, 0, 0, 217, 220, 5, 21, 0, 
	0, 218, 220, 5, 23, 0, 0, 219, 211, 1, 0, 0, 0, 219, 215, 1, 0, 0, 0, 219, 
	216, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 218, 1, 0, 0, 0, 220, 25, 1, 
	0, 0, 0, 221, 226, 3, 24, 12, 0, 222, 226, 5, 20, 0, 0, 223, 226, 5, 19, 
	0, 0, 224, 226, 3, 14, 7, 0, 225, 221, 1, 0, 0, 0, 225, 222, 1, 0, 0, 0, 
	225, 223, 1, 0, 0, 0, 225, 224, 1, 0, 0, 0, 226, 27, 1, 0, 0, 0, 227, 228, 
	7, 2, 0, 0, 228, 29, 1, 0, 0, 0, 229, 232, 3, 26, 13, 0, 230, 232, 3, 28, 
	14, 0, 231, 229, 1, 0, 0, 0, 231, 230, 1, 0, 0, 0, 232, 31, 1, 0, 0, 0, 
	233, 235, 3, 30, 15, 0, 234, 233, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 
	234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 239, 1, 0, 0, 0, 238, 236, 
	1, 0, 0, 0, 239, 240, 5, 22, 0, 0, 240, 33, 1, 0, 0, 0, 241, 245, 3, 22, 
	11, 0, 242, 245, 3, 24, 12, 0, 243, 245, 3, 28, 14, 0, 244, 241, 1, 0, 
	0, 0, 244, 242, 1, 0, 0, 0, 244, 243, 1, 0, 0, 0, 245, 35, 1, 0, 0, 0, 
	246, 248, 3, 34, 17, 0, 247, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 
	247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 252, 1, 0, 0, 0, 251, 253, 
	5, 22, 0, 0, 252, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 252, 1, 0, 
	0, 0, 254, 255, 1, 0, 0, 0, 255, 37, 1, 0, 0, 0, 256, 258, 5, 23, 0, 0, 
	257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 
	260, 1, 0, 0, 0, 260, 262, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 266, 
	5, 11, 0, 0, 263, 265, 5, 23, 0, 0, 264, 263, 1, 0, 0, 0, 265, 268, 1, 
	0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 269, 1, 0, 0, 
	0, 268, 266, 1, 0, 0, 0, 269, 270, 3, 36, 18, 0, 270, 39, 1, 0, 0, 0, 271, 
	275, 5, 12, 0, 0, 272, 274, 5, 22, 0, 0, 273, 272, 1, 0, 0, 0, 274, 277, 
	1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 281, 1, 0, 
	0, 0, 277, 275, 1, 0, 0, 0, 278, 280, 3, 38, 19, 0, 279, 278, 1, 0, 0, 
	0, 280, 283, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 
	284, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 284, 286, 5, 13, 0, 0, 285, 287, 
	5, 22, 0, 0, 286, 285, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 320, 1, 0, 
	0, 0, 288, 292, 5, 14, 0, 0, 289, 291, 5, 22, 0, 0, 290, 289, 1, 0, 0, 
	0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 
	298, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 297, 3, 38, 19, 0, 296, 295, 
	1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 
	0, 0, 299, 301, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301, 303, 5, 15, 0, 0, 
	302, 304, 5, 22, 0, 0, 303, 302, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 
	320, 1, 0, 0, 0, 305, 307, 5, 16, 0, 0, 306, 308, 5, 22, 0, 0, 307, 306, 
	1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 312, 1, 0, 0, 0, 309, 311, 3, 32, 
	16, 0, 310, 309, 1, 0, 0, 0, 311, 314, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 
	312, 313, 1, 0, 0, 0, 313, 315, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 315, 
	317, 5, 17, 0, 0, 316, 318, 5, 22, 0, 0, 317, 316, 1, 0, 0, 0, 317, 318, 
	1, 0, 0, 0, 318, 320, 1, 0, 0, 0, 319, 271, 1, 0, 0, 0, 319, 288, 1, 0, 
	0, 0, 319, 305, 1, 0, 0, 0, 320, 41, 1, 0, 0, 0, 50, 50, 60, 66, 70, 76, 
	81, 87, 93, 97, 103, 109, 113, 119, 121, 126, 128, 131, 137, 143, 148, 
	153, 160, 163, 171, 179, 185, 193, 199, 205, 209, 213, 219, 225, 231, 236, 
	244, 249, 254, 259, 266, 275, 281, 286, 292, 298, 303, 307, 312, 317, 319,
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
	LatexParserT__16 = 17
	LatexParserLETTER = 18
	LatexParserSPECIAL = 19
	LatexParserSYMBOL = 20
	LatexParserNUMBER = 21
	LatexParserNEWLINE = 22
	LatexParserWS = 23
	LatexParserCR = 24
	LatexParserPUNCTUATION = 25
	LatexParserEXTENDED_UNICODE = 26
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
	LatexParserRULE_curly_parens = 14
	LatexParserRULE_verbatim_block_content = 15
	LatexParserRULE_verbatim_line = 16
	LatexParserRULE_block_content = 17
	LatexParserRULE_block_line = 18
	LatexParserRULE_block_item = 19
	LatexParserRULE_block = 20
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
		p.SetState(42)
		p.Note_title()
	}
	{
		p.SetState(43)
		p.Note_url()
	}
	{
		p.SetState(44)
		p.Note_created()
	}
	{
		p.SetState(45)
		p.Note_updated()
	}
	{
		p.SetState(46)
		p.Line_break()
	}
	p.SetState(50)
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
				p.SetState(47)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(52)
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
		p.SetState(53)
		p.Note_text()
	}
	{
		p.SetState(54)
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
		p.SetState(56)
		p.Match(LatexParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(60)
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
				p.SetState(57)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 111412224) != 0) {
		{
			p.SetState(63)
			p.Word()
		}


		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(69)
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
	AllVerbatim_content() []IVerbatim_contentContext
	Verbatim_content(i int) IVerbatim_contentContext
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

func (s *Note_urlContext) AllVerbatim_content() []IVerbatim_contentContext {
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

func (s *Note_urlContext) Verbatim_content(i int) IVerbatim_contentContext {
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
		p.SetState(72)
		p.Match(LatexParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 112985092) != 0) {
		{
			p.SetState(73)
			p.Verbatim_content()
		}


		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Match(LatexParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(LatexParserT__1)
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
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(80)
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
		p.SetState(83)
		p.Match(LatexParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(87)
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
				p.SetState(84)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 111412224) != 0) {
		{
			p.SetState(90)
			p.Word()
		}


		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(LatexParserT__1)
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
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(96)
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
		p.SetState(99)
		p.Match(LatexParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(103)
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
				p.SetState(100)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 111412224) != 0) {
		{
			p.SetState(106)
			p.Word()
		}


		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(111)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(112)
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
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 115693060) != 0) {
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LatexParserT__8, LatexParserT__9, LatexParserLETTER, LatexParserNUMBER, LatexParserWS, LatexParserPUNCTUATION, LatexParserEXTENDED_UNICODE:
			{
				p.SetState(115)
				p.Text()
			}


		case LatexParserT__11, LatexParserT__13, LatexParserT__15:
			{
				p.SetState(116)
				p.Block()
			}


		case LatexParserT__1:
			{
				p.SetState(117)
				p.Line_break()
			}


		case LatexParserNEWLINE:
			{
				p.SetState(118)
				p.Empty_line()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(123)
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
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				p.SetState(126)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(124)
						p.Tag()
					}


				case 2:
					{
						p.SetState(125)
						p.Word()
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(130)
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
		p.SetState(133)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(137)
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
				p.SetState(134)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(139)
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

		p.SetState(143)
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
	AllLETTER() []antlr.TerminalNode
	LETTER(i int) antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	SYMBOL() antlr.TerminalNode

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


func (s *Escaped_wordContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(LatexParserLETTER)
}

func (s *Escaped_wordContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserLETTER, i)
}

func (s *Escaped_wordContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(LatexParserWS)
}

func (s *Escaped_wordContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserWS, i)
}

func (s *Escaped_wordContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(LatexParserSYMBOL, 0)
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

	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserLETTER:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(145)

						var _m = p.Match(LatexParserLETTER)

						localctx.(*Escaped_wordContext).content = _m
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(153)
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
					p.SetState(150)

					var _m = p.Match(LatexParserWS)

					localctx.(*Escaped_wordContext).space = _m
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			}
			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}


	case LatexParserT__3, LatexParserT__6, LatexParserSYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Escaped_wordContext).content = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 1048720) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Escaped_wordContext).content = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(157)

					var _m = p.Match(LatexParserWS)

					localctx.(*Escaped_wordContext).space = _m
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			}
			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
			if p.HasError() {
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
		p.SetState(165)
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


type Url_tagContext struct {
	TagContext
}

func NewUrl_tagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Url_tagContext {
	var p = new(Url_tagContext)

	InitEmptyTagContext(&p.TagContext)
	p.parser = parser
	p.CopyAll(ctx.(*TagContext))

	return p
}

func (s *Url_tagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Url_tagContext) AllVerbatim_content() []IVerbatim_contentContext {
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

func (s *Url_tagContext) Verbatim_content(i int) IVerbatim_contentContext {
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


func (s *Url_tagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterUrl_tag(s)
	}
}

func (s *Url_tagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitUrl_tag(s)
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

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUrl_tagContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(LatexParserT__8)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 112985092) != 0) {
			{
				p.SetState(168)
				p.Verbatim_content()
			}


			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(173)
			p.Match(LatexParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		localctx = NewSimple_tagContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(LatexParserT__9)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == LatexParserLETTER {
			{
				p.SetState(176)

				var _m = p.Match(LatexParserLETTER)

				localctx.(*Simple_tagContext).name = _m
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(181)
			p.Match(LatexParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 111412224) != 0) {
			{
				p.SetState(182)
				p.Word()
			}


			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(187)
			p.Match(LatexParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 3:
		localctx = NewComposite_tagContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(189)
			p.Match(LatexParserT__9)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == LatexParserLETTER {
			{
				p.SetState(190)

				var _m = p.Match(LatexParserLETTER)

				localctx.(*Composite_tagContext).name = _m
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(195)
			p.Match(LatexParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 111412224) != 0) {
			{
				p.SetState(196)
				p.Word()
			}


			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(201)
			p.Composite_tag_separator()
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 111412224) != 0) {
			{
				p.SetState(202)
				p.Word()
			}


			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(207)
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

func (s *PunctuationContext) EXTENDED_UNICODE() antlr.TerminalNode {
	return s.GetToken(LatexParserEXTENDED_UNICODE, 0)
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
	var _la int

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserT__9:
		localctx = NewEscapedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(LatexParserT__9)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(212)
				p.Escaped_word()
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}


	case LatexParserLETTER:
		localctx = NewLetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(LatexParserLETTER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LatexParserPUNCTUATION, LatexParserEXTENDED_UNICODE:
		localctx = NewPunctuationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LatexParserPUNCTUATION || _la == LatexParserEXTENDED_UNICODE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	case LatexParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
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
			p.SetState(218)
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


type Verbatim_specialContext struct {
	Verbatim_contentContext
}

func NewVerbatim_specialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Verbatim_specialContext {
	var p = new(Verbatim_specialContext)

	InitEmptyVerbatim_contentContext(&p.Verbatim_contentContext)
	p.parser = parser
	p.CopyAll(ctx.(*Verbatim_contentContext))

	return p
}

func (s *Verbatim_specialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Verbatim_specialContext) SPECIAL() antlr.TerminalNode {
	return s.GetToken(LatexParserSPECIAL, 0)
}


func (s *Verbatim_specialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterVerbatim_special(s)
	}
}

func (s *Verbatim_specialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitVerbatim_special(s)
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
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserT__9, LatexParserLETTER, LatexParserNUMBER, LatexParserWS, LatexParserPUNCTUATION, LatexParserEXTENDED_UNICODE:
		localctx = NewVerbatim_wordContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Word()
		}


	case LatexParserSYMBOL:
		localctx = NewVerbatim_symbolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Match(LatexParserSYMBOL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LatexParserSPECIAL:
		localctx = NewVerbatim_specialContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.Match(LatexParserSPECIAL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LatexParserT__1:
		localctx = NewVerbatim_linebreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(224)
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


// ICurly_parensContext is an interface to support dynamic dispatch.
type ICurly_parensContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCurly_parensContext differentiates from other interfaces.
	IsCurly_parensContext()
}

type Curly_parensContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurly_parensContext() *Curly_parensContext {
	var p = new(Curly_parensContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_curly_parens
	return p
}

func InitEmptyCurly_parensContext(p *Curly_parensContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_curly_parens
}

func (*Curly_parensContext) IsCurly_parensContext() {}

func NewCurly_parensContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Curly_parensContext {
	var p = new(Curly_parensContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_curly_parens

	return p
}

func (s *Curly_parensContext) GetParser() antlr.Parser { return s.parser }
func (s *Curly_parensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Curly_parensContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Curly_parensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterCurly_parens(s)
	}
}

func (s *Curly_parensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitCurly_parens(s)
	}
}




func (p *LatexParser) Curly_parens() (localctx ICurly_parensContext) {
	localctx = NewCurly_parensContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LatexParserRULE_curly_parens)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LatexParserT__3 || _la == LatexParserT__6) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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


// IVerbatim_block_contentContext is an interface to support dynamic dispatch.
type IVerbatim_block_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Verbatim_content() IVerbatim_contentContext
	Curly_parens() ICurly_parensContext

	// IsVerbatim_block_contentContext differentiates from other interfaces.
	IsVerbatim_block_contentContext()
}

type Verbatim_block_contentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerbatim_block_contentContext() *Verbatim_block_contentContext {
	var p = new(Verbatim_block_contentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_verbatim_block_content
	return p
}

func InitEmptyVerbatim_block_contentContext(p *Verbatim_block_contentContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_verbatim_block_content
}

func (*Verbatim_block_contentContext) IsVerbatim_block_contentContext() {}

func NewVerbatim_block_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Verbatim_block_contentContext {
	var p = new(Verbatim_block_contentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_verbatim_block_content

	return p
}

func (s *Verbatim_block_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Verbatim_block_contentContext) Verbatim_content() IVerbatim_contentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerbatim_contentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerbatim_contentContext)
}

func (s *Verbatim_block_contentContext) Curly_parens() ICurly_parensContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICurly_parensContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICurly_parensContext)
}

func (s *Verbatim_block_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Verbatim_block_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Verbatim_block_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterVerbatim_block_content(s)
	}
}

func (s *Verbatim_block_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitVerbatim_block_content(s)
	}
}




func (p *LatexParser) Verbatim_block_content() (localctx IVerbatim_block_contentContext) {
	localctx = NewVerbatim_block_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LatexParserRULE_verbatim_block_content)
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserT__1, LatexParserT__9, LatexParserLETTER, LatexParserSPECIAL, LatexParserSYMBOL, LatexParserNUMBER, LatexParserWS, LatexParserPUNCTUATION, LatexParserEXTENDED_UNICODE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Verbatim_content()
		}


	case LatexParserT__3, LatexParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Curly_parens()
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
	AllVerbatim_block_content() []IVerbatim_block_contentContext
	Verbatim_block_content(i int) IVerbatim_block_contentContext

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

func (s *Verbatim_lineContext) AllVerbatim_block_content() []IVerbatim_block_contentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVerbatim_block_contentContext); ok {
			len++
		}
	}

	tst := make([]IVerbatim_block_contentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVerbatim_block_contentContext); ok {
			tst[i] = t.(IVerbatim_block_contentContext)
			i++
		}
	}

	return tst
}

func (s *Verbatim_lineContext) Verbatim_block_content(i int) IVerbatim_block_contentContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerbatim_block_contentContext); ok {
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

	return t.(IVerbatim_block_contentContext)
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
	p.EnterRule(localctx, 32, LatexParserRULE_verbatim_line)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 112985236) != 0) {
		{
			p.SetState(233)
			p.Verbatim_block_content()
		}


		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(239)
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


// IBlock_contentContext is an interface to support dynamic dispatch.
type IBlock_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tag() ITagContext
	Word() IWordContext
	Curly_parens() ICurly_parensContext

	// IsBlock_contentContext differentiates from other interfaces.
	IsBlock_contentContext()
}

type Block_contentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_contentContext() *Block_contentContext {
	var p = new(Block_contentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_block_content
	return p
}

func InitEmptyBlock_contentContext(p *Block_contentContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LatexParserRULE_block_content
}

func (*Block_contentContext) IsBlock_contentContext() {}

func NewBlock_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_contentContext {
	var p = new(Block_contentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatexParserRULE_block_content

	return p
}

func (s *Block_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_contentContext) Tag() ITagContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Block_contentContext) Word() IWordContext {
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

func (s *Block_contentContext) Curly_parens() ICurly_parensContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICurly_parensContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICurly_parensContext)
}

func (s *Block_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Block_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterBlock_content(s)
	}
}

func (s *Block_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitBlock_content(s)
	}
}




func (p *LatexParser) Block_content() (localctx IBlock_contentContext) {
	localctx = NewBlock_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LatexParserRULE_block_content)
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.Tag()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Word()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(243)
			p.Curly_parens()
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


// IBlock_lineContext is an interface to support dynamic dispatch.
type IBlock_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBlock_content() []IBlock_contentContext
	Block_content(i int) IBlock_contentContext
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

func (s *Block_lineContext) AllBlock_content() []IBlock_contentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_contentContext); ok {
			len++
		}
	}

	tst := make([]IBlock_contentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_contentContext); ok {
			tst[i] = t.(IBlock_contentContext)
			i++
		}
	}

	return tst
}

func (s *Block_lineContext) Block_content(i int) IBlock_contentContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_contentContext); ok {
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

	return t.(IBlock_contentContext)
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
	p.EnterRule(localctx, 36, LatexParserRULE_block_line)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 111412880) != 0) {
		{
			p.SetState(246)
			p.Block_content()
		}


		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LatexParserNEWLINE {
		{
			p.SetState(251)
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
	p.EnterRule(localctx, 38, LatexParserRULE_block_item)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == LatexParserWS {
		{
			p.SetState(256)
			p.Match(LatexParserWS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(262)
		p.Match(LatexParserT__10)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(263)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(269)
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
	p.EnterRule(localctx, 40, LatexParserRULE_block)
	var _la int

	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserT__11:
		localctx = NewItemizeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Match(LatexParserT__11)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserNEWLINE {
			{
				p.SetState(272)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserT__10 || _la == LatexParserWS {
			{
				p.SetState(278)
				p.Block_item()
			}


			p.SetState(283)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(284)
			p.Match(LatexParserT__12)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(285)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}


	case LatexParserT__13:
		localctx = NewEnumerateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(288)
			p.Match(LatexParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserNEWLINE {
			{
				p.SetState(289)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserT__10 || _la == LatexParserWS {
			{
				p.SetState(295)
				p.Block_item()
			}


			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(301)
			p.Match(LatexParserT__14)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(302)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}


	case LatexParserT__15:
		localctx = NewVerbatimContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(305)
			p.Match(LatexParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(306)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 117179540) != 0) {
			{
				p.SetState(309)
				p.Verbatim_line()
			}


			p.SetState(314)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(315)
			p.Match(LatexParserT__16)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(316)
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


