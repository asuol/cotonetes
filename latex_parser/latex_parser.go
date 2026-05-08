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
    "'\\textbf{Created:}'", "'\\textbf{Last Updated:}'", "'\\'", "'{'", 
    "'\\item'", "'\\begin{itemize}'", "'\\end{itemize}'", "'\\begin{enumerate}'", 
    "'\\end{enumerate}'", "'\\begin{verbatim}'", "'\\end{verbatim}'", "", 
    "", "", "", "'\\n'", "", "'\\r'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LETTER", 
    "PUNCTUATION", "SYMBOL", "NUMBER", "NEWLINE", "WS", "CR",
  }
  staticData.RuleNames = []string{
    "latex", "note_title", "note_url", "note_created", "note_updated", "note_text", 
    "text", "line_break", "empty_line", "escaped_word", "tag", "word", "verbatim_content", 
    "verbatim_line", "block_line", "block_item", "block",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 22, 256, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 41, 8, 0, 10, 0, 
	12, 0, 44, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 51, 8, 1, 10, 1, 12, 
	1, 54, 9, 1, 1, 1, 4, 1, 57, 8, 1, 11, 1, 12, 1, 58, 1, 1, 1, 1, 3, 1, 
	63, 8, 1, 1, 2, 1, 2, 4, 2, 67, 8, 2, 11, 2, 12, 2, 68, 1, 2, 1, 2, 1, 
	2, 3, 2, 74, 8, 2, 1, 3, 1, 3, 5, 3, 78, 8, 3, 10, 3, 12, 3, 81, 9, 3, 
	1, 3, 4, 3, 84, 8, 3, 11, 3, 12, 3, 85, 1, 3, 1, 3, 3, 3, 90, 8, 3, 1, 
	4, 1, 4, 5, 4, 94, 8, 4, 10, 4, 12, 4, 97, 9, 4, 1, 4, 4, 4, 100, 8, 4, 
	11, 4, 12, 4, 101, 1, 4, 1, 4, 3, 4, 106, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 
	5, 5, 112, 8, 5, 10, 5, 12, 5, 115, 9, 5, 1, 6, 1, 6, 4, 6, 119, 8, 6, 
	11, 6, 12, 6, 120, 1, 6, 3, 6, 124, 8, 6, 1, 7, 1, 7, 5, 7, 128, 8, 7, 
	10, 7, 12, 7, 131, 9, 7, 1, 8, 4, 8, 134, 8, 8, 11, 8, 12, 8, 135, 1, 9, 
	1, 9, 4, 9, 140, 8, 9, 11, 9, 12, 9, 141, 1, 9, 5, 9, 145, 8, 9, 10, 9, 
	12, 9, 148, 9, 9, 1, 9, 3, 9, 151, 8, 9, 1, 10, 1, 10, 4, 10, 155, 8, 10, 
	11, 10, 12, 10, 156, 1, 10, 1, 10, 4, 10, 161, 8, 10, 11, 10, 12, 10, 162, 
	1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 172, 8, 11, 1, 
	12, 1, 12, 1, 12, 3, 12, 177, 8, 12, 1, 13, 5, 13, 180, 8, 13, 10, 13, 
	12, 13, 183, 9, 13, 1, 13, 1, 13, 1, 14, 4, 14, 188, 8, 14, 11, 14, 12, 
	14, 189, 1, 14, 4, 14, 193, 8, 14, 11, 14, 12, 14, 194, 1, 15, 1, 15, 5, 
	15, 199, 8, 15, 10, 15, 12, 15, 202, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 
	5, 16, 208, 8, 16, 10, 16, 12, 16, 211, 9, 16, 1, 16, 5, 16, 214, 8, 16, 
	10, 16, 12, 16, 217, 9, 16, 1, 16, 1, 16, 3, 16, 221, 8, 16, 1, 16, 1, 
	16, 5, 16, 225, 8, 16, 10, 16, 12, 16, 228, 9, 16, 1, 16, 5, 16, 231, 8, 
	16, 10, 16, 12, 16, 234, 9, 16, 1, 16, 1, 16, 3, 16, 238, 8, 16, 1, 16, 
	1, 16, 3, 16, 242, 8, 16, 1, 16, 5, 16, 245, 8, 16, 10, 16, 12, 16, 248, 
	9, 16, 1, 16, 1, 16, 3, 16, 252, 8, 16, 3, 16, 254, 8, 16, 1, 16, 0, 0, 
	17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 1, 
	2, 0, 16, 16, 18, 18, 285, 0, 34, 1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 4, 64, 
	1, 0, 0, 0, 6, 75, 1, 0, 0, 0, 8, 91, 1, 0, 0, 0, 10, 113, 1, 0, 0, 0, 
	12, 118, 1, 0, 0, 0, 14, 125, 1, 0, 0, 0, 16, 133, 1, 0, 0, 0, 18, 137, 
	1, 0, 0, 0, 20, 152, 1, 0, 0, 0, 22, 171, 1, 0, 0, 0, 24, 176, 1, 0, 0, 
	0, 26, 181, 1, 0, 0, 0, 28, 187, 1, 0, 0, 0, 30, 196, 1, 0, 0, 0, 32, 253, 
	1, 0, 0, 0, 34, 35, 3, 2, 1, 0, 35, 36, 3, 4, 2, 0, 36, 37, 3, 6, 3, 0, 
	37, 38, 3, 8, 4, 0, 38, 42, 3, 14, 7, 0, 39, 41, 5, 20, 0, 0, 40, 39, 1, 
	0, 0, 0, 41, 44, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 
	45, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 45, 46, 3, 10, 5, 0, 46, 47, 5, 0, 
	0, 1, 47, 1, 1, 0, 0, 0, 48, 52, 5, 1, 0, 0, 49, 51, 5, 21, 0, 0, 50, 49, 
	1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 
	53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 57, 3, 22, 11, 0, 56, 55, 1, 
	0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 
	60, 1, 0, 0, 0, 60, 62, 5, 2, 0, 0, 61, 63, 5, 20, 0, 0, 62, 61, 1, 0, 
	0, 0, 62, 63, 1, 0, 0, 0, 63, 3, 1, 0, 0, 0, 64, 66, 5, 3, 0, 0, 65, 67, 
	3, 22, 11, 0, 66, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0, 
	0, 68, 69, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 5, 4, 0, 0, 71, 73, 
	5, 2, 0, 0, 72, 74, 5, 20, 0, 0, 73, 72, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 
	74, 5, 1, 0, 0, 0, 75, 79, 5, 5, 0, 0, 76, 78, 5, 21, 0, 0, 77, 76, 1, 
	0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 
	83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 84, 3, 22, 11, 0, 83, 82, 1, 0, 
	0, 0, 84, 85, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 
	1, 0, 0, 0, 87, 89, 5, 2, 0, 0, 88, 90, 5, 20, 0, 0, 89, 88, 1, 0, 0, 0, 
	89, 90, 1, 0, 0, 0, 90, 7, 1, 0, 0, 0, 91, 95, 5, 6, 0, 0, 92, 94, 5, 21, 
	0, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 
	1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 100, 3, 22, 11, 
	0, 99, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 
	102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 105, 5, 2, 0, 0, 104, 106, 
	5, 20, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 9, 1, 0, 
	0, 0, 107, 112, 3, 12, 6, 0, 108, 112, 3, 32, 16, 0, 109, 112, 3, 14, 7, 
	0, 110, 112, 3, 16, 8, 0, 111, 107, 1, 0, 0, 0, 111, 108, 1, 0, 0, 0, 111, 
	109, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 
	1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 11, 1, 0, 0, 0, 115, 113, 1, 0, 
	0, 0, 116, 119, 3, 20, 10, 0, 117, 119, 3, 22, 11, 0, 118, 116, 1, 0, 0, 
	0, 118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 
	121, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 124, 5, 20, 0, 0, 123, 122, 
	1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 13, 1, 0, 0, 0, 125, 129, 5, 2, 
	0, 0, 126, 128, 5, 21, 0, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 
	129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 15, 1, 0, 0, 0, 131, 129, 
	1, 0, 0, 0, 132, 134, 5, 20, 0, 0, 133, 132, 1, 0, 0, 0, 134, 135, 1, 0, 
	0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 17, 1, 0, 0, 0, 
	137, 139, 5, 7, 0, 0, 138, 140, 7, 0, 0, 0, 139, 138, 1, 0, 0, 0, 140, 
	141, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 146, 
	1, 0, 0, 0, 143, 145, 5, 21, 0, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0, 
	0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 
	148, 146, 1, 0, 0, 0, 149, 151, 5, 20, 0, 0, 150, 149, 1, 0, 0, 0, 150, 
	151, 1, 0, 0, 0, 151, 19, 1, 0, 0, 0, 152, 154, 5, 7, 0, 0, 153, 155, 5, 
	16, 0, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 154, 1, 0, 0, 
	0, 156, 157, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 160, 5, 8, 0, 0, 159, 
	161, 3, 22, 11, 0, 160, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 160, 
	1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 5, 4, 
	0, 0, 165, 21, 1, 0, 0, 0, 166, 172, 3, 18, 9, 0, 167, 172, 5, 16, 0, 0, 
	168, 172, 5, 17, 0, 0, 169, 172, 5, 19, 0, 0, 170, 172, 5, 21, 0, 0, 171, 
	166, 1, 0, 0, 0, 171, 167, 1, 0, 0, 0, 171, 168, 1, 0, 0, 0, 171, 169, 
	1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 23, 1, 0, 0, 0, 173, 177, 3, 22, 
	11, 0, 174, 177, 5, 18, 0, 0, 175, 177, 3, 14, 7, 0, 176, 173, 1, 0, 0, 
	0, 176, 174, 1, 0, 0, 0, 176, 175, 1, 0, 0, 0, 177, 25, 1, 0, 0, 0, 178, 
	180, 3, 24, 12, 0, 179, 178, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 
	1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183, 181, 1, 0, 
	0, 0, 184, 185, 5, 20, 0, 0, 185, 27, 1, 0, 0, 0, 186, 188, 3, 22, 11, 
	0, 187, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 
	190, 1, 0, 0, 0, 190, 192, 1, 0, 0, 0, 191, 193, 5, 20, 0, 0, 192, 191, 
	1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 
	0, 0, 195, 29, 1, 0, 0, 0, 196, 200, 5, 9, 0, 0, 197, 199, 5, 21, 0, 0, 
	198, 197, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 
	201, 1, 0, 0, 0, 201, 203, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 204, 
	3, 28, 14, 0, 204, 31, 1, 0, 0, 0, 205, 209, 5, 10, 0, 0, 206, 208, 5, 
	20, 0, 0, 207, 206, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 
	0, 209, 210, 1, 0, 0, 0, 210, 215, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212, 
	214, 3, 30, 15, 0, 213, 212, 1, 0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213, 
	1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 218, 1, 0, 0, 0, 217, 215, 1, 0, 
	0, 0, 218, 220, 5, 11, 0, 0, 219, 221, 5, 20, 0, 0, 220, 219, 1, 0, 0, 
	0, 220, 221, 1, 0, 0, 0, 221, 254, 1, 0, 0, 0, 222, 226, 5, 12, 0, 0, 223, 
	225, 5, 20, 0, 0, 224, 223, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 
	1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 232, 1, 0, 0, 0, 228, 226, 1, 0, 
	0, 0, 229, 231, 3, 30, 15, 0, 230, 229, 1, 0, 0, 0, 231, 234, 1, 0, 0, 
	0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 235, 1, 0, 0, 0, 234, 
	232, 1, 0, 0, 0, 235, 237, 5, 13, 0, 0, 236, 238, 5, 20, 0, 0, 237, 236, 
	1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 254, 1, 0, 0, 0, 239, 241, 5, 14, 
	0, 0, 240, 242, 5, 20, 0, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 
	242, 246, 1, 0, 0, 0, 243, 245, 3, 26, 13, 0, 244, 243, 1, 0, 0, 0, 245, 
	248, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 249, 
	1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 249, 251, 5, 15, 0, 0, 250, 252, 5, 20, 
	0, 0, 251, 250, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 254, 1, 0, 0, 0, 
	253, 205, 1, 0, 0, 0, 253, 222, 1, 0, 0, 0, 253, 239, 1, 0, 0, 0, 254, 
	33, 1, 0, 0, 0, 40, 42, 52, 58, 62, 68, 73, 79, 85, 89, 95, 101, 105, 111, 
	113, 118, 120, 123, 129, 135, 141, 146, 150, 156, 162, 171, 176, 181, 189, 
	194, 200, 209, 215, 220, 226, 232, 237, 241, 246, 251, 253,
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
	LatexParserLETTER = 16
	LatexParserPUNCTUATION = 17
	LatexParserSYMBOL = 18
	LatexParserNUMBER = 19
	LatexParserNEWLINE = 20
	LatexParserWS = 21
	LatexParserCR = 22
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
	LatexParserRULE_tag = 10
	LatexParserRULE_word = 11
	LatexParserRULE_verbatim_content = 12
	LatexParserRULE_verbatim_line = 13
	LatexParserRULE_block_line = 14
	LatexParserRULE_block_item = 15
	LatexParserRULE_block = 16
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
		p.SetState(34)
		p.Note_title()
	}
	{
		p.SetState(35)
		p.Note_url()
	}
	{
		p.SetState(36)
		p.Note_created()
	}
	{
		p.SetState(37)
		p.Note_updated()
	}
	{
		p.SetState(38)
		p.Line_break()
	}
	p.SetState(42)
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
				p.SetState(39)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


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
	}
	{
		p.SetState(45)
		p.Note_text()
	}
	{
		p.SetState(46)
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
		p.SetState(48)
		p.Match(LatexParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(52)
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
				p.SetState(49)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
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
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2818176) != 0) {
		{
			p.SetState(55)
			p.Word()
		}


		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(61)
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
		p.SetState(64)
		p.Match(LatexParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2818176) != 0) {
		{
			p.SetState(65)
			p.Word()
		}


		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(LatexParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(72)
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
		p.SetState(75)
		p.Match(LatexParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(79)
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
				p.SetState(76)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
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
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2818176) != 0) {
		{
			p.SetState(82)
			p.Word()
		}


		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(88)
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
		p.SetState(91)
		p.Match(LatexParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(95)
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
				p.SetState(92)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
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
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2818176) != 0) {
		{
			p.SetState(98)
			p.Word()
		}


		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LatexParserNEWLINE {
		{
			p.SetState(104)
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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3888260) != 0) {
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LatexParserT__6, LatexParserLETTER, LatexParserPUNCTUATION, LatexParserNUMBER, LatexParserWS:
			{
				p.SetState(107)
				p.Text()
			}


		case LatexParserT__9, LatexParserT__11, LatexParserT__13:
			{
				p.SetState(108)
				p.Block()
			}


		case LatexParserT__1:
			{
				p.SetState(109)
				p.Line_break()
			}


		case LatexParserNEWLINE:
			{
				p.SetState(110)
				p.Empty_line()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(115)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				p.SetState(118)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(116)
						p.Tag()
					}


				case 2:
					{
						p.SetState(117)
						p.Word()
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(122)
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
		p.SetState(125)
		p.Match(LatexParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(129)
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
				p.SetState(126)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(132)
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

		p.SetState(135)
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
		p.SetState(137)
		p.Match(LatexParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(138)

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

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(146)
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
				p.SetState(143)

				var _m = p.Match(LatexParserWS)

				localctx.(*Escaped_wordContext).space = _m
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
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
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(149)
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


// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	AllWord() []IWordContext
	Word(i int) IWordContext
	AllLETTER() []antlr.TerminalNode
	LETTER(i int) antlr.TerminalNode

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
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

func (s *TagContext) GetName() antlr.Token { return s.name }


func (s *TagContext) SetName(v antlr.Token) { s.name = v }


func (s *TagContext) AllWord() []IWordContext {
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

func (s *TagContext) Word(i int) IWordContext {
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

func (s *TagContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(LatexParserLETTER)
}

func (s *TagContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(LatexParserLETTER, i)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatexListener); ok {
		listenerT.ExitTag(s)
	}
}




func (p *LatexParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LatexParserRULE_tag)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(LatexParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LatexParserLETTER {
		{
			p.SetState(153)

			var _m = p.Match(LatexParserLETTER)

			localctx.(*TagContext).name = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(158)
		p.Match(LatexParserT__7)
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


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2818176) != 0) {
		{
			p.SetState(159)
			p.Word()
		}


		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(164)
		p.Match(LatexParserT__3)
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
	p.EnterRule(localctx, 22, LatexParserRULE_word)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserT__6:
		localctx = NewEscapedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Escaped_word()
		}


	case LatexParserLETTER:
		localctx = NewLetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
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
			p.SetState(168)
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
			p.SetState(169)
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
			p.SetState(170)
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
	p.EnterRule(localctx, 24, LatexParserRULE_verbatim_content)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserT__6, LatexParserLETTER, LatexParserPUNCTUATION, LatexParserNUMBER, LatexParserWS:
		localctx = NewVerbatim_wordContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Word()
		}


	case LatexParserSYMBOL:
		localctx = NewVerbatim_symbolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
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
			p.SetState(175)
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
	p.EnterRule(localctx, 26, LatexParserRULE_verbatim_line)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3080324) != 0) {
		{
			p.SetState(178)
			p.Verbatim_content()
		}


		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
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
	p.EnterRule(localctx, 28, LatexParserRULE_block_line)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2818176) != 0) {
		{
			p.SetState(186)
			p.Word()
		}


		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LatexParserNEWLINE {
		{
			p.SetState(191)
			p.Match(LatexParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(194)
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
	p.EnterRule(localctx, 30, LatexParserRULE_block_item)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(LatexParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(197)
				p.Match(LatexParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(203)
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
	p.EnterRule(localctx, 32, LatexParserRULE_block)
	var _la int

	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LatexParserT__9:
		localctx = NewItemizeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(LatexParserT__9)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserNEWLINE {
			{
				p.SetState(206)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(211)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserT__8 {
			{
				p.SetState(212)
				p.Block_item()
			}


			p.SetState(217)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(218)
			p.Match(LatexParserT__10)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(219)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}


	case LatexParserT__11:
		localctx = NewEnumerateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Match(LatexParserT__11)
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
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserNEWLINE {
			{
				p.SetState(223)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(228)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LatexParserT__8 {
			{
				p.SetState(229)
				p.Block_item()
			}


			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(235)
			p.Match(LatexParserT__12)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(236)
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
		localctx = NewVerbatimContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)
			p.Match(LatexParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(240)
				p.Match(LatexParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 4128900) != 0) {
			{
				p.SetState(243)
				p.Verbatim_line()
			}


			p.SetState(248)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(249)
			p.Match(LatexParserT__14)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(250)
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


