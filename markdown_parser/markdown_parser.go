// Code generated from Markdown.g4 by ANTLR 4.13.2. DO NOT EDIT.

package markdown_parser // Markdown
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


type MarkdownParser struct {
	*antlr.BaseParser
}

var MarkdownParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func markdownParserInit() {
  staticData := &MarkdownParserStaticData
  staticData.LiteralNames = []string{
    "", "'['", "']'", "'('", "')'", "", "'*'", "'.'", "'\\'", "", "", "", 
    "", "'\\n'", "", "'\\r'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "LETTER", "ASTERISK", "DOT", "BACKSLASH", "FENCED_BLOCK", 
    "SPECIAL", "SYMBOL", "NUMBER", "NEWLINE", "WS", "CR", "PUNCTUATION", 
    "EXTENDED_UNICODE",
  }
  staticData.RuleNames = []string{
    "markdown", "text", "empty_line", "bold", "url", "star", "word", "url_title", 
    "url_value", "item_line", "number_line", "verbatim_line", "block",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 17, 201, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 5, 0, 30, 8, 0, 10, 0, 
	12, 0, 33, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 41, 8, 1, 11, 
	1, 12, 1, 42, 1, 1, 3, 1, 46, 8, 1, 1, 2, 4, 2, 49, 8, 2, 11, 2, 12, 2, 
	50, 1, 3, 1, 3, 1, 3, 4, 3, 56, 8, 3, 11, 3, 12, 3, 57, 1, 3, 1, 3, 1, 
	3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 
	6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 81, 8, 6, 1, 7, 1, 7, 1, 7, 
	1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 93, 8, 7, 11, 7, 12, 7, 
	94, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 107, 
	8, 8, 11, 8, 12, 8, 108, 1, 9, 5, 9, 112, 8, 9, 10, 9, 12, 9, 115, 9, 9, 
	1, 9, 1, 9, 4, 9, 119, 8, 9, 11, 9, 12, 9, 120, 1, 9, 1, 9, 1, 9, 1, 9, 
	4, 9, 127, 8, 9, 11, 9, 12, 9, 128, 1, 9, 5, 9, 132, 8, 9, 10, 9, 12, 9, 
	135, 9, 9, 1, 10, 5, 10, 138, 8, 10, 10, 10, 12, 10, 141, 9, 10, 1, 10, 
	1, 10, 1, 10, 5, 10, 146, 8, 10, 10, 10, 12, 10, 149, 9, 10, 1, 10, 1, 
	10, 1, 10, 1, 10, 4, 10, 155, 8, 10, 11, 10, 12, 10, 156, 1, 10, 5, 10, 
	160, 8, 10, 10, 10, 12, 10, 163, 9, 10, 1, 11, 1, 11, 4, 11, 167, 8, 11, 
	11, 11, 12, 11, 168, 1, 11, 1, 11, 1, 11, 3, 11, 174, 8, 11, 1, 12, 4, 
	12, 177, 8, 12, 11, 12, 12, 12, 178, 1, 12, 4, 12, 182, 8, 12, 11, 12, 
	12, 12, 183, 1, 12, 1, 12, 3, 12, 188, 8, 12, 1, 12, 4, 12, 191, 8, 12, 
	11, 12, 12, 12, 192, 1, 12, 1, 12, 3, 12, 197, 8, 12, 3, 12, 199, 8, 12, 
	1, 12, 0, 0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 4, 2, 
	0, 5, 5, 7, 7, 1, 0, 16, 17, 1, 0, 1, 2, 1, 0, 3, 4, 249, 0, 31, 1, 0, 
	0, 0, 2, 40, 1, 0, 0, 0, 4, 48, 1, 0, 0, 0, 6, 52, 1, 0, 0, 0, 8, 62, 1, 
	0, 0, 0, 10, 69, 1, 0, 0, 0, 12, 80, 1, 0, 0, 0, 14, 92, 1, 0, 0, 0, 16, 
	106, 1, 0, 0, 0, 18, 113, 1, 0, 0, 0, 20, 139, 1, 0, 0, 0, 22, 173, 1, 
	0, 0, 0, 24, 198, 1, 0, 0, 0, 26, 30, 3, 24, 12, 0, 27, 30, 3, 2, 1, 0, 
	28, 30, 3, 4, 2, 0, 29, 26, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 28, 1, 
	0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 
	34, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 35, 5, 0, 0, 1, 35, 1, 1, 0, 0, 
	0, 36, 41, 3, 6, 3, 0, 37, 41, 3, 8, 4, 0, 38, 41, 3, 10, 5, 0, 39, 41, 
	3, 12, 6, 0, 40, 36, 1, 0, 0, 0, 40, 37, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 
	40, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 
	0, 0, 0, 43, 45, 1, 0, 0, 0, 44, 46, 5, 13, 0, 0, 45, 44, 1, 0, 0, 0, 45, 
	46, 1, 0, 0, 0, 46, 3, 1, 0, 0, 0, 47, 49, 5, 13, 0, 0, 48, 47, 1, 0, 0, 
	0, 49, 50, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 5, 1, 
	0, 0, 0, 52, 53, 5, 6, 0, 0, 53, 55, 5, 6, 0, 0, 54, 56, 3, 12, 6, 0, 55, 
	54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 
	0, 58, 59, 1, 0, 0, 0, 59, 60, 5, 6, 0, 0, 60, 61, 5, 6, 0, 0, 61, 7, 1, 
	0, 0, 0, 62, 63, 5, 1, 0, 0, 63, 64, 3, 14, 7, 0, 64, 65, 5, 2, 0, 0, 65, 
	66, 5, 3, 0, 0, 66, 67, 3, 16, 8, 0, 67, 68, 5, 4, 0, 0, 68, 9, 1, 0, 0, 
	0, 69, 70, 5, 6, 0, 0, 70, 11, 1, 0, 0, 0, 71, 81, 5, 11, 0, 0, 72, 81, 
	5, 10, 0, 0, 73, 81, 7, 0, 0, 0, 74, 81, 7, 1, 0, 0, 75, 81, 5, 8, 0, 0, 
	76, 81, 7, 2, 0, 0, 77, 81, 7, 3, 0, 0, 78, 81, 5, 12, 0, 0, 79, 81, 5, 
	14, 0, 0, 80, 71, 1, 0, 0, 0, 80, 72, 1, 0, 0, 0, 80, 73, 1, 0, 0, 0, 80, 
	74, 1, 0, 0, 0, 80, 75, 1, 0, 0, 0, 80, 76, 1, 0, 0, 0, 80, 77, 1, 0, 0, 
	0, 80, 78, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 13, 1, 0, 0, 0, 82, 93, 
	5, 11, 0, 0, 83, 93, 5, 10, 0, 0, 84, 93, 5, 5, 0, 0, 85, 93, 5, 6, 0, 
	0, 86, 93, 5, 7, 0, 0, 87, 93, 5, 16, 0, 0, 88, 93, 5, 17, 0, 0, 89, 93, 
	7, 3, 0, 0, 90, 93, 5, 12, 0, 0, 91, 93, 5, 14, 0, 0, 92, 82, 1, 0, 0, 
	0, 92, 83, 1, 0, 0, 0, 92, 84, 1, 0, 0, 0, 92, 85, 1, 0, 0, 0, 92, 86, 
	1, 0, 0, 0, 92, 87, 1, 0, 0, 0, 92, 88, 1, 0, 0, 0, 92, 89, 1, 0, 0, 0, 
	92, 90, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 
	0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 15, 1, 0, 0, 0, 96, 107, 5, 11, 0, 0, 
	97, 107, 5, 10, 0, 0, 98, 107, 5, 5, 0, 0, 99, 107, 5, 6, 0, 0, 100, 107, 
	5, 7, 0, 0, 101, 107, 5, 16, 0, 0, 102, 107, 5, 17, 0, 0, 103, 107, 7, 
	2, 0, 0, 104, 107, 5, 12, 0, 0, 105, 107, 5, 14, 0, 0, 106, 96, 1, 0, 0, 
	0, 106, 97, 1, 0, 0, 0, 106, 98, 1, 0, 0, 0, 106, 99, 1, 0, 0, 0, 106, 
	100, 1, 0, 0, 0, 106, 101, 1, 0, 0, 0, 106, 102, 1, 0, 0, 0, 106, 103, 
	1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 105, 1, 0, 0, 0, 107, 108, 1, 0, 
	0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 17, 1, 0, 0, 0, 
	110, 112, 5, 14, 0, 0, 111, 110, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 
	111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115, 113, 
	1, 0, 0, 0, 116, 118, 5, 6, 0, 0, 117, 119, 5, 14, 0, 0, 118, 117, 1, 0, 
	0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 
	121, 126, 1, 0, 0, 0, 122, 127, 3, 6, 3, 0, 123, 127, 3, 8, 4, 0, 124, 
	127, 3, 10, 5, 0, 125, 127, 3, 12, 6, 0, 126, 122, 1, 0, 0, 0, 126, 123, 
	1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 
	0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 133, 1, 0, 0, 0, 
	130, 132, 5, 13, 0, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 
	131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 19, 1, 0, 0, 0, 135, 133, 1, 
	0, 0, 0, 136, 138, 5, 14, 0, 0, 137, 136, 1, 0, 0, 0, 138, 141, 1, 0, 0, 
	0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 142, 1, 0, 0, 0, 141, 
	139, 1, 0, 0, 0, 142, 143, 5, 12, 0, 0, 143, 147, 5, 7, 0, 0, 144, 146, 
	5, 14, 0, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 
	0, 0, 147, 148, 1, 0, 0, 0, 148, 154, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 
	150, 155, 3, 6, 3, 0, 151, 155, 3, 8, 4, 0, 152, 155, 3, 10, 5, 0, 153, 
	155, 3, 12, 6, 0, 154, 150, 1, 0, 0, 0, 154, 151, 1, 0, 0, 0, 154, 152, 
	1, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 154, 1, 0, 
	0, 0, 156, 157, 1, 0, 0, 0, 157, 161, 1, 0, 0, 0, 158, 160, 5, 13, 0, 0, 
	159, 158, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 
	162, 1, 0, 0, 0, 162, 21, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 167, 3, 
	10, 5, 0, 165, 167, 3, 12, 6, 0, 166, 164, 1, 0, 0, 0, 166, 165, 1, 0, 
	0, 0, 167, 168, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 
	169, 170, 1, 0, 0, 0, 170, 171, 5, 13, 0, 0, 171, 174, 1, 0, 0, 0, 172, 
	174, 5, 13, 0, 0, 173, 166, 1, 0, 0, 0, 173, 172, 1, 0, 0, 0, 174, 23, 
	1, 0, 0, 0, 175, 177, 3, 18, 9, 0, 176, 175, 1, 0, 0, 0, 177, 178, 1, 0, 
	0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 199, 1, 0, 0, 0, 
	180, 182, 3, 20, 10, 0, 181, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 
	181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 199, 1, 0, 0, 0, 185, 187, 
	5, 9, 0, 0, 186, 188, 5, 13, 0, 0, 187, 186, 1, 0, 0, 0, 187, 188, 1, 0, 
	0, 0, 188, 190, 1, 0, 0, 0, 189, 191, 3, 22, 11, 0, 190, 189, 1, 0, 0, 
	0, 191, 192, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 
	194, 1, 0, 0, 0, 194, 196, 5, 9, 0, 0, 195, 197, 5, 13, 0, 0, 196, 195, 
	1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 1, 0, 0, 0, 198, 176, 1, 0, 
	0, 0, 198, 181, 1, 0, 0, 0, 198, 185, 1, 0, 0, 0, 199, 25, 1, 0, 0, 0, 
	31, 29, 31, 40, 42, 45, 50, 57, 80, 92, 94, 106, 108, 113, 120, 126, 128, 
	133, 139, 147, 154, 156, 161, 166, 168, 173, 178, 183, 187, 192, 196, 198,
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

// MarkdownParserInit initializes any static state used to implement MarkdownParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMarkdownParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MarkdownParserInit() {
  staticData := &MarkdownParserStaticData
  staticData.once.Do(markdownParserInit)
}

// NewMarkdownParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMarkdownParser(input antlr.TokenStream) *MarkdownParser {
	MarkdownParserInit()
	this := new(MarkdownParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &MarkdownParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Markdown.g4"

	return this
}


// MarkdownParser tokens.
const (
	MarkdownParserEOF = antlr.TokenEOF
	MarkdownParserT__0 = 1
	MarkdownParserT__1 = 2
	MarkdownParserT__2 = 3
	MarkdownParserT__3 = 4
	MarkdownParserLETTER = 5
	MarkdownParserASTERISK = 6
	MarkdownParserDOT = 7
	MarkdownParserBACKSLASH = 8
	MarkdownParserFENCED_BLOCK = 9
	MarkdownParserSPECIAL = 10
	MarkdownParserSYMBOL = 11
	MarkdownParserNUMBER = 12
	MarkdownParserNEWLINE = 13
	MarkdownParserWS = 14
	MarkdownParserCR = 15
	MarkdownParserPUNCTUATION = 16
	MarkdownParserEXTENDED_UNICODE = 17
)

// MarkdownParser rules.
const (
	MarkdownParserRULE_markdown = 0
	MarkdownParserRULE_text = 1
	MarkdownParserRULE_empty_line = 2
	MarkdownParserRULE_bold = 3
	MarkdownParserRULE_url = 4
	MarkdownParserRULE_star = 5
	MarkdownParserRULE_word = 6
	MarkdownParserRULE_url_title = 7
	MarkdownParserRULE_url_value = 8
	MarkdownParserRULE_item_line = 9
	MarkdownParserRULE_number_line = 10
	MarkdownParserRULE_verbatim_line = 11
	MarkdownParserRULE_block = 12
)

// IMarkdownContext is an interface to support dynamic dispatch.
type IMarkdownContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllText() []ITextContext
	Text(i int) ITextContext
	AllEmpty_line() []IEmpty_lineContext
	Empty_line(i int) IEmpty_lineContext

	// IsMarkdownContext differentiates from other interfaces.
	IsMarkdownContext()
}

type MarkdownContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMarkdownContext() *MarkdownContext {
	var p = new(MarkdownContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_markdown
	return p
}

func InitEmptyMarkdownContext(p *MarkdownContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_markdown
}

func (*MarkdownContext) IsMarkdownContext() {}

func NewMarkdownContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MarkdownContext {
	var p = new(MarkdownContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_markdown

	return p
}

func (s *MarkdownContext) GetParser() antlr.Parser { return s.parser }

func (s *MarkdownContext) EOF() antlr.TerminalNode {
	return s.GetToken(MarkdownParserEOF, 0)
}

func (s *MarkdownContext) AllBlock() []IBlockContext {
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

func (s *MarkdownContext) Block(i int) IBlockContext {
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

func (s *MarkdownContext) AllText() []ITextContext {
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

func (s *MarkdownContext) Text(i int) ITextContext {
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

func (s *MarkdownContext) AllEmpty_line() []IEmpty_lineContext {
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

func (s *MarkdownContext) Empty_line(i int) IEmpty_lineContext {
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

func (s *MarkdownContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarkdownContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MarkdownContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterMarkdown(s)
	}
}

func (s *MarkdownContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitMarkdown(s)
	}
}




func (p *MarkdownParser) Markdown() (localctx IMarkdownContext) {
	localctx = NewMarkdownContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MarkdownParserRULE_markdown)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 229374) != 0) {
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(26)
				p.Block()
			}


		case 2:
			{
				p.SetState(27)
				p.Text()
			}


		case 3:
			{
				p.SetState(28)
				p.Empty_line()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
		p.Match(MarkdownParserEOF)
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


// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBold() []IBoldContext
	Bold(i int) IBoldContext
	AllUrl() []IUrlContext
	Url(i int) IUrlContext
	AllStar() []IStarContext
	Star(i int) IStarContext
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
	p.RuleIndex = MarkdownParserRULE_text
	return p
}

func InitEmptyTextContext(p *TextContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_text
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) AllBold() []IBoldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoldContext); ok {
			len++
		}
	}

	tst := make([]IBoldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoldContext); ok {
			tst[i] = t.(IBoldContext)
			i++
		}
	}

	return tst
}

func (s *TextContext) Bold(i int) IBoldContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoldContext); ok {
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

	return t.(IBoldContext)
}

func (s *TextContext) AllUrl() []IUrlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUrlContext); ok {
			len++
		}
	}

	tst := make([]IUrlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUrlContext); ok {
			tst[i] = t.(IUrlContext)
			i++
		}
	}

	return tst
}

func (s *TextContext) Url(i int) IUrlContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrlContext); ok {
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

	return t.(IUrlContext)
}

func (s *TextContext) AllStar() []IStarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStarContext); ok {
			len++
		}
	}

	tst := make([]IStarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStarContext); ok {
			tst[i] = t.(IStarContext)
			i++
		}
	}

	return tst
}

func (s *TextContext) Star(i int) IStarContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
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

	return t.(IStarContext)
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
	return s.GetToken(MarkdownParserNEWLINE, 0)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitText(s)
	}
}




func (p *MarkdownParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MarkdownParserRULE_text)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				p.SetState(40)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(36)
						p.Bold()
					}


				case 2:
					{
						p.SetState(37)
						p.Url()
					}


				case 3:
					{
						p.SetState(38)
						p.Star()
					}


				case 4:
					{
						p.SetState(39)
						p.Word()
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(44)
			p.Match(MarkdownParserNEWLINE)
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
	p.RuleIndex = MarkdownParserRULE_empty_line
	return p
}

func InitEmptyEmpty_lineContext(p *Empty_lineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_empty_line
}

func (*Empty_lineContext) IsEmpty_lineContext() {}

func NewEmpty_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Empty_lineContext {
	var p = new(Empty_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_empty_line

	return p
}

func (s *Empty_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Empty_lineContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *Empty_lineContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *Empty_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Empty_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Empty_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterEmpty_line(s)
	}
}

func (s *Empty_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitEmpty_line(s)
	}
}




func (p *MarkdownParser) Empty_line() (localctx IEmpty_lineContext) {
	localctx = NewEmpty_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MarkdownParserRULE_empty_line)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(47)
					p.Match(MarkdownParserNEWLINE)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}




		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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


// IBoldContext is an interface to support dynamic dispatch.
type IBoldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllASTERISK() []antlr.TerminalNode
	ASTERISK(i int) antlr.TerminalNode
	AllWord() []IWordContext
	Word(i int) IWordContext

	// IsBoldContext differentiates from other interfaces.
	IsBoldContext()
}

type BoldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoldContext() *BoldContext {
	var p = new(BoldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_bold
	return p
}

func InitEmptyBoldContext(p *BoldContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_bold
}

func (*BoldContext) IsBoldContext() {}

func NewBoldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoldContext {
	var p = new(BoldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_bold

	return p
}

func (s *BoldContext) GetParser() antlr.Parser { return s.parser }

func (s *BoldContext) AllASTERISK() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserASTERISK)
}

func (s *BoldContext) ASTERISK(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserASTERISK, i)
}

func (s *BoldContext) AllWord() []IWordContext {
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

func (s *BoldContext) Word(i int) IWordContext {
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

func (s *BoldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BoldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterBold(s)
	}
}

func (s *BoldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitBold(s)
	}
}




func (p *MarkdownParser) Bold() (localctx IBoldContext) {
	localctx = NewBoldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MarkdownParserRULE_bold)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(MarkdownParserASTERISK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(MarkdownParserASTERISK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 220606) != 0) {
		{
			p.SetState(54)
			p.Word()
		}


		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)
		p.Match(MarkdownParserASTERISK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(MarkdownParserASTERISK)
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


// IUrlContext is an interface to support dynamic dispatch.
type IUrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Url_title() IUrl_titleContext
	Url_value() IUrl_valueContext

	// IsUrlContext differentiates from other interfaces.
	IsUrlContext()
}

type UrlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrlContext() *UrlContext {
	var p = new(UrlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_url
	return p
}

func InitEmptyUrlContext(p *UrlContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_url
}

func (*UrlContext) IsUrlContext() {}

func NewUrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UrlContext {
	var p = new(UrlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_url

	return p
}

func (s *UrlContext) GetParser() antlr.Parser { return s.parser }

func (s *UrlContext) Url_title() IUrl_titleContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrl_titleContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrl_titleContext)
}

func (s *UrlContext) Url_value() IUrl_valueContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrl_valueContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrl_valueContext)
}

func (s *UrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterUrl(s)
	}
}

func (s *UrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitUrl(s)
	}
}




func (p *MarkdownParser) Url() (localctx IUrlContext) {
	localctx = NewUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MarkdownParserRULE_url)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(MarkdownParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Url_title()
	}
	{
		p.SetState(64)
		p.Match(MarkdownParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(MarkdownParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Url_value()
	}
	{
		p.SetState(67)
		p.Match(MarkdownParserT__3)
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


// IStarContext is an interface to support dynamic dispatch.
type IStarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASTERISK() antlr.TerminalNode

	// IsStarContext differentiates from other interfaces.
	IsStarContext()
}

type StarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStarContext() *StarContext {
	var p = new(StarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_star
	return p
}

func InitEmptyStarContext(p *StarContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_star
}

func (*StarContext) IsStarContext() {}

func NewStarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StarContext {
	var p = new(StarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_star

	return p
}

func (s *StarContext) GetParser() antlr.Parser { return s.parser }

func (s *StarContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserASTERISK, 0)
}

func (s *StarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterStar(s)
	}
}

func (s *StarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitStar(s)
	}
}




func (p *MarkdownParser) Star() (localctx IStarContext) {
	localctx = NewStarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MarkdownParserRULE_star)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(MarkdownParserASTERISK)
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
	p.RuleIndex = MarkdownParserRULE_word
	return p
}

func InitEmptyWordContext(p *WordContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_word
}

func (*WordContext) IsWordContext() {}

func NewWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordContext {
	var p = new(WordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_word

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




type SpecialContext struct {
	WordContext
}

func NewSpecialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SpecialContext {
	var p = new(SpecialContext)

	InitEmptyWordContext(&p.WordContext)
	p.parser = parser
	p.CopyAll(ctx.(*WordContext))

	return p
}

func (s *SpecialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecialContext) SPECIAL() antlr.TerminalNode {
	return s.GetToken(MarkdownParserSPECIAL, 0)
}


func (s *SpecialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterSpecial(s)
	}
}

func (s *SpecialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitSpecial(s)
	}
}


type SymbolContext struct {
	WordContext
}

func NewSymbolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SymbolContext {
	var p = new(SymbolContext)

	InitEmptyWordContext(&p.WordContext)
	p.parser = parser
	p.CopyAll(ctx.(*WordContext))

	return p
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(MarkdownParserSYMBOL, 0)
}


func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitSymbol(s)
	}
}


type SquareContext struct {
	WordContext
}

func NewSquareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SquareContext {
	var p = new(SquareContext)

	InitEmptyWordContext(&p.WordContext)
	p.parser = parser
	p.CopyAll(ctx.(*WordContext))

	return p
}

func (s *SquareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SquareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterSquare(s)
	}
}

func (s *SquareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitSquare(s)
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
	return s.GetToken(MarkdownParserNUMBER, 0)
}


func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitNumber(s)
	}
}


type RoundContext struct {
	WordContext
}

func NewRoundContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoundContext {
	var p = new(RoundContext)

	InitEmptyWordContext(&p.WordContext)
	p.parser = parser
	p.CopyAll(ctx.(*WordContext))

	return p
}

func (s *RoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterRound(s)
	}
}

func (s *RoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitRound(s)
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
	return s.GetToken(MarkdownParserLETTER, 0)
}

func (s *LetterContext) DOT() antlr.TerminalNode {
	return s.GetToken(MarkdownParserDOT, 0)
}


func (s *LetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterLetter(s)
	}
}

func (s *LetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
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
	return s.GetToken(MarkdownParserPUNCTUATION, 0)
}

func (s *PunctuationContext) EXTENDED_UNICODE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserEXTENDED_UNICODE, 0)
}


func (s *PunctuationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterPunctuation(s)
	}
}

func (s *PunctuationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
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
	return s.GetToken(MarkdownParserWS, 0)
}


func (s *WsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterWs(s)
	}
}

func (s *WsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitWs(s)
	}
}


type BackslashContext struct {
	WordContext
}

func NewBackslashContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BackslashContext {
	var p = new(BackslashContext)

	InitEmptyWordContext(&p.WordContext)
	p.parser = parser
	p.CopyAll(ctx.(*WordContext))

	return p
}

func (s *BackslashContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BackslashContext) BACKSLASH() antlr.TerminalNode {
	return s.GetToken(MarkdownParserBACKSLASH, 0)
}


func (s *BackslashContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterBackslash(s)
	}
}

func (s *BackslashContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitBackslash(s)
	}
}



func (p *MarkdownParser) Word() (localctx IWordContext) {
	localctx = NewWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MarkdownParserRULE_word)
	var _la int

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserSYMBOL:
		localctx = NewSymbolContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Match(MarkdownParserSYMBOL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserSPECIAL:
		localctx = NewSpecialContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(MarkdownParserSPECIAL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserLETTER, MarkdownParserDOT:
		localctx = NewLetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MarkdownParserLETTER || _la == MarkdownParserDOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	case MarkdownParserPUNCTUATION, MarkdownParserEXTENDED_UNICODE:
		localctx = NewPunctuationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MarkdownParserPUNCTUATION || _la == MarkdownParserEXTENDED_UNICODE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	case MarkdownParserBACKSLASH:
		localctx = NewBackslashContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)
			p.Match(MarkdownParserBACKSLASH)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserT__0, MarkdownParserT__1:
		localctx = NewSquareContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(76)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MarkdownParserT__0 || _la == MarkdownParserT__1) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	case MarkdownParserT__2, MarkdownParserT__3:
		localctx = NewRoundContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(77)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MarkdownParserT__2 || _la == MarkdownParserT__3) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	case MarkdownParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(78)
			p.Match(MarkdownParserNUMBER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserWS:
		localctx = NewWsContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(79)
			p.Match(MarkdownParserWS)
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


// IUrl_titleContext is an interface to support dynamic dispatch.
type IUrl_titleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSYMBOL() []antlr.TerminalNode
	SYMBOL(i int) antlr.TerminalNode
	AllSPECIAL() []antlr.TerminalNode
	SPECIAL(i int) antlr.TerminalNode
	AllLETTER() []antlr.TerminalNode
	LETTER(i int) antlr.TerminalNode
	AllASTERISK() []antlr.TerminalNode
	ASTERISK(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllPUNCTUATION() []antlr.TerminalNode
	PUNCTUATION(i int) antlr.TerminalNode
	AllEXTENDED_UNICODE() []antlr.TerminalNode
	EXTENDED_UNICODE(i int) antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsUrl_titleContext differentiates from other interfaces.
	IsUrl_titleContext()
}

type Url_titleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrl_titleContext() *Url_titleContext {
	var p = new(Url_titleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_url_title
	return p
}

func InitEmptyUrl_titleContext(p *Url_titleContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_url_title
}

func (*Url_titleContext) IsUrl_titleContext() {}

func NewUrl_titleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Url_titleContext {
	var p = new(Url_titleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_url_title

	return p
}

func (s *Url_titleContext) GetParser() antlr.Parser { return s.parser }

func (s *Url_titleContext) AllSYMBOL() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserSYMBOL)
}

func (s *Url_titleContext) SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserSYMBOL, i)
}

func (s *Url_titleContext) AllSPECIAL() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserSPECIAL)
}

func (s *Url_titleContext) SPECIAL(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserSPECIAL, i)
}

func (s *Url_titleContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserLETTER)
}

func (s *Url_titleContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserLETTER, i)
}

func (s *Url_titleContext) AllASTERISK() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserASTERISK)
}

func (s *Url_titleContext) ASTERISK(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserASTERISK, i)
}

func (s *Url_titleContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserDOT)
}

func (s *Url_titleContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserDOT, i)
}

func (s *Url_titleContext) AllPUNCTUATION() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserPUNCTUATION)
}

func (s *Url_titleContext) PUNCTUATION(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserPUNCTUATION, i)
}

func (s *Url_titleContext) AllEXTENDED_UNICODE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserEXTENDED_UNICODE)
}

func (s *Url_titleContext) EXTENDED_UNICODE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserEXTENDED_UNICODE, i)
}

func (s *Url_titleContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNUMBER)
}

func (s *Url_titleContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNUMBER, i)
}

func (s *Url_titleContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWS)
}

func (s *Url_titleContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWS, i)
}

func (s *Url_titleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Url_titleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Url_titleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterUrl_title(s)
	}
}

func (s *Url_titleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitUrl_title(s)
	}
}




func (p *MarkdownParser) Url_title() (localctx IUrl_titleContext) {
	localctx = NewUrl_titleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MarkdownParserRULE_url_title)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 220408) != 0) {
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MarkdownParserSYMBOL:
			{
				p.SetState(82)
				p.Match(MarkdownParserSYMBOL)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserSPECIAL:
			{
				p.SetState(83)
				p.Match(MarkdownParserSPECIAL)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserLETTER:
			{
				p.SetState(84)
				p.Match(MarkdownParserLETTER)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserASTERISK:
			{
				p.SetState(85)
				p.Match(MarkdownParserASTERISK)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserDOT:
			{
				p.SetState(86)
				p.Match(MarkdownParserDOT)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserPUNCTUATION:
			{
				p.SetState(87)
				p.Match(MarkdownParserPUNCTUATION)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserEXTENDED_UNICODE:
			{
				p.SetState(88)
				p.Match(MarkdownParserEXTENDED_UNICODE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserT__2, MarkdownParserT__3:
			{
				p.SetState(89)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MarkdownParserT__2 || _la == MarkdownParserT__3) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}


		case MarkdownParserNUMBER:
			{
				p.SetState(90)
				p.Match(MarkdownParserNUMBER)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserWS:
			{
				p.SetState(91)
				p.Match(MarkdownParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(94)
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


// IUrl_valueContext is an interface to support dynamic dispatch.
type IUrl_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSYMBOL() []antlr.TerminalNode
	SYMBOL(i int) antlr.TerminalNode
	AllSPECIAL() []antlr.TerminalNode
	SPECIAL(i int) antlr.TerminalNode
	AllLETTER() []antlr.TerminalNode
	LETTER(i int) antlr.TerminalNode
	AllASTERISK() []antlr.TerminalNode
	ASTERISK(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllPUNCTUATION() []antlr.TerminalNode
	PUNCTUATION(i int) antlr.TerminalNode
	AllEXTENDED_UNICODE() []antlr.TerminalNode
	EXTENDED_UNICODE(i int) antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsUrl_valueContext differentiates from other interfaces.
	IsUrl_valueContext()
}

type Url_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrl_valueContext() *Url_valueContext {
	var p = new(Url_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_url_value
	return p
}

func InitEmptyUrl_valueContext(p *Url_valueContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_url_value
}

func (*Url_valueContext) IsUrl_valueContext() {}

func NewUrl_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Url_valueContext {
	var p = new(Url_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_url_value

	return p
}

func (s *Url_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Url_valueContext) AllSYMBOL() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserSYMBOL)
}

func (s *Url_valueContext) SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserSYMBOL, i)
}

func (s *Url_valueContext) AllSPECIAL() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserSPECIAL)
}

func (s *Url_valueContext) SPECIAL(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserSPECIAL, i)
}

func (s *Url_valueContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserLETTER)
}

func (s *Url_valueContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserLETTER, i)
}

func (s *Url_valueContext) AllASTERISK() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserASTERISK)
}

func (s *Url_valueContext) ASTERISK(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserASTERISK, i)
}

func (s *Url_valueContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserDOT)
}

func (s *Url_valueContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserDOT, i)
}

func (s *Url_valueContext) AllPUNCTUATION() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserPUNCTUATION)
}

func (s *Url_valueContext) PUNCTUATION(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserPUNCTUATION, i)
}

func (s *Url_valueContext) AllEXTENDED_UNICODE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserEXTENDED_UNICODE)
}

func (s *Url_valueContext) EXTENDED_UNICODE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserEXTENDED_UNICODE, i)
}

func (s *Url_valueContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNUMBER)
}

func (s *Url_valueContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNUMBER, i)
}

func (s *Url_valueContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWS)
}

func (s *Url_valueContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWS, i)
}

func (s *Url_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Url_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Url_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterUrl_value(s)
	}
}

func (s *Url_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitUrl_value(s)
	}
}




func (p *MarkdownParser) Url_value() (localctx IUrl_valueContext) {
	localctx = NewUrl_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MarkdownParserRULE_url_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 220390) != 0) {
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MarkdownParserSYMBOL:
			{
				p.SetState(96)
				p.Match(MarkdownParserSYMBOL)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserSPECIAL:
			{
				p.SetState(97)
				p.Match(MarkdownParserSPECIAL)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserLETTER:
			{
				p.SetState(98)
				p.Match(MarkdownParserLETTER)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserASTERISK:
			{
				p.SetState(99)
				p.Match(MarkdownParserASTERISK)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserDOT:
			{
				p.SetState(100)
				p.Match(MarkdownParserDOT)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserPUNCTUATION:
			{
				p.SetState(101)
				p.Match(MarkdownParserPUNCTUATION)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserEXTENDED_UNICODE:
			{
				p.SetState(102)
				p.Match(MarkdownParserEXTENDED_UNICODE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserT__0, MarkdownParserT__1:
			{
				p.SetState(103)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MarkdownParserT__0 || _la == MarkdownParserT__1) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}


		case MarkdownParserNUMBER:
			{
				p.SetState(104)
				p.Match(MarkdownParserNUMBER)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		case MarkdownParserWS:
			{
				p.SetState(105)
				p.Match(MarkdownParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(108)
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


// IItem_lineContext is an interface to support dynamic dispatch.
type IItem_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASTERISK() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllBold() []IBoldContext
	Bold(i int) IBoldContext
	AllUrl() []IUrlContext
	Url(i int) IUrlContext
	AllStar() []IStarContext
	Star(i int) IStarContext
	AllWord() []IWordContext
	Word(i int) IWordContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsItem_lineContext differentiates from other interfaces.
	IsItem_lineContext()
}

type Item_lineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItem_lineContext() *Item_lineContext {
	var p = new(Item_lineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_item_line
	return p
}

func InitEmptyItem_lineContext(p *Item_lineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_item_line
}

func (*Item_lineContext) IsItem_lineContext() {}

func NewItem_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Item_lineContext {
	var p = new(Item_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_item_line

	return p
}

func (s *Item_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Item_lineContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserASTERISK, 0)
}

func (s *Item_lineContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWS)
}

func (s *Item_lineContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWS, i)
}

func (s *Item_lineContext) AllBold() []IBoldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoldContext); ok {
			len++
		}
	}

	tst := make([]IBoldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoldContext); ok {
			tst[i] = t.(IBoldContext)
			i++
		}
	}

	return tst
}

func (s *Item_lineContext) Bold(i int) IBoldContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoldContext); ok {
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

	return t.(IBoldContext)
}

func (s *Item_lineContext) AllUrl() []IUrlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUrlContext); ok {
			len++
		}
	}

	tst := make([]IUrlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUrlContext); ok {
			tst[i] = t.(IUrlContext)
			i++
		}
	}

	return tst
}

func (s *Item_lineContext) Url(i int) IUrlContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrlContext); ok {
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

	return t.(IUrlContext)
}

func (s *Item_lineContext) AllStar() []IStarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStarContext); ok {
			len++
		}
	}

	tst := make([]IStarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStarContext); ok {
			tst[i] = t.(IStarContext)
			i++
		}
	}

	return tst
}

func (s *Item_lineContext) Star(i int) IStarContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
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

	return t.(IStarContext)
}

func (s *Item_lineContext) AllWord() []IWordContext {
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

func (s *Item_lineContext) Word(i int) IWordContext {
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

func (s *Item_lineContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *Item_lineContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *Item_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Item_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Item_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterItem_line(s)
	}
}

func (s *Item_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitItem_line(s)
	}
}




func (p *MarkdownParser) Item_line() (localctx IItem_lineContext) {
	localctx = NewItem_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MarkdownParserRULE_item_line)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == MarkdownParserWS {
		{
			p.SetState(110)
			p.Match(MarkdownParserWS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(116)
		p.Match(MarkdownParserASTERISK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(117)
					p.Match(MarkdownParserWS)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}




		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
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
						p.SetState(122)
						p.Bold()
					}


				case 2:
					{
						p.SetState(123)
						p.Url()
					}


				case 3:
					{
						p.SetState(124)
						p.Star()
					}


				case 4:
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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(130)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
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


// INumber_lineContext is an interface to support dynamic dispatch.
type INumber_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	DOT() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllBold() []IBoldContext
	Bold(i int) IBoldContext
	AllUrl() []IUrlContext
	Url(i int) IUrlContext
	AllStar() []IStarContext
	Star(i int) IStarContext
	AllWord() []IWordContext
	Word(i int) IWordContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsNumber_lineContext differentiates from other interfaces.
	IsNumber_lineContext()
}

type Number_lineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_lineContext() *Number_lineContext {
	var p = new(Number_lineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_number_line
	return p
}

func InitEmptyNumber_lineContext(p *Number_lineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_number_line
}

func (*Number_lineContext) IsNumber_lineContext() {}

func NewNumber_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_lineContext {
	var p = new(Number_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_number_line

	return p
}

func (s *Number_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_lineContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNUMBER, 0)
}

func (s *Number_lineContext) DOT() antlr.TerminalNode {
	return s.GetToken(MarkdownParserDOT, 0)
}

func (s *Number_lineContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWS)
}

func (s *Number_lineContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWS, i)
}

func (s *Number_lineContext) AllBold() []IBoldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoldContext); ok {
			len++
		}
	}

	tst := make([]IBoldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoldContext); ok {
			tst[i] = t.(IBoldContext)
			i++
		}
	}

	return tst
}

func (s *Number_lineContext) Bold(i int) IBoldContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoldContext); ok {
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

	return t.(IBoldContext)
}

func (s *Number_lineContext) AllUrl() []IUrlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUrlContext); ok {
			len++
		}
	}

	tst := make([]IUrlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUrlContext); ok {
			tst[i] = t.(IUrlContext)
			i++
		}
	}

	return tst
}

func (s *Number_lineContext) Url(i int) IUrlContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrlContext); ok {
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

	return t.(IUrlContext)
}

func (s *Number_lineContext) AllStar() []IStarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStarContext); ok {
			len++
		}
	}

	tst := make([]IStarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStarContext); ok {
			tst[i] = t.(IStarContext)
			i++
		}
	}

	return tst
}

func (s *Number_lineContext) Star(i int) IStarContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
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

	return t.(IStarContext)
}

func (s *Number_lineContext) AllWord() []IWordContext {
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

func (s *Number_lineContext) Word(i int) IWordContext {
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

func (s *Number_lineContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *Number_lineContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *Number_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Number_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterNumber_line(s)
	}
}

func (s *Number_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitNumber_line(s)
	}
}




func (p *MarkdownParser) Number_line() (localctx INumber_lineContext) {
	localctx = NewNumber_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MarkdownParserRULE_number_line)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == MarkdownParserWS {
		{
			p.SetState(136)
			p.Match(MarkdownParserWS)
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
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(142)
		p.Match(MarkdownParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(MarkdownParserDOT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(144)
				p.Match(MarkdownParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				p.SetState(154)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(150)
						p.Bold()
					}


				case 2:
					{
						p.SetState(151)
						p.Url()
					}


				case 3:
					{
						p.SetState(152)
						p.Star()
					}


				case 4:
					{
						p.SetState(153)
						p.Word()
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(161)
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
				p.SetState(158)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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


// IVerbatim_lineContext is an interface to support dynamic dispatch.
type IVerbatim_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEWLINE() antlr.TerminalNode
	AllStar() []IStarContext
	Star(i int) IStarContext
	AllWord() []IWordContext
	Word(i int) IWordContext

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
	p.RuleIndex = MarkdownParserRULE_verbatim_line
	return p
}

func InitEmptyVerbatim_lineContext(p *Verbatim_lineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_verbatim_line
}

func (*Verbatim_lineContext) IsVerbatim_lineContext() {}

func NewVerbatim_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Verbatim_lineContext {
	var p = new(Verbatim_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_verbatim_line

	return p
}

func (s *Verbatim_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Verbatim_lineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, 0)
}

func (s *Verbatim_lineContext) AllStar() []IStarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStarContext); ok {
			len++
		}
	}

	tst := make([]IStarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStarContext); ok {
			tst[i] = t.(IStarContext)
			i++
		}
	}

	return tst
}

func (s *Verbatim_lineContext) Star(i int) IStarContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
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

	return t.(IStarContext)
}

func (s *Verbatim_lineContext) AllWord() []IWordContext {
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

func (s *Verbatim_lineContext) Word(i int) IWordContext {
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

func (s *Verbatim_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Verbatim_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Verbatim_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterVerbatim_line(s)
	}
}

func (s *Verbatim_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitVerbatim_line(s)
	}
}




func (p *MarkdownParser) Verbatim_line() (localctx IVerbatim_lineContext) {
	localctx = NewVerbatim_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MarkdownParserRULE_verbatim_line)
	var _la int

	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserT__0, MarkdownParserT__1, MarkdownParserT__2, MarkdownParserT__3, MarkdownParserLETTER, MarkdownParserASTERISK, MarkdownParserDOT, MarkdownParserBACKSLASH, MarkdownParserSPECIAL, MarkdownParserSYMBOL, MarkdownParserNUMBER, MarkdownParserWS, MarkdownParserPUNCTUATION, MarkdownParserEXTENDED_UNICODE:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 220670) != 0) {
			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case MarkdownParserASTERISK:
				{
					p.SetState(164)
					p.Star()
				}


			case MarkdownParserT__0, MarkdownParserT__1, MarkdownParserT__2, MarkdownParserT__3, MarkdownParserLETTER, MarkdownParserDOT, MarkdownParserBACKSLASH, MarkdownParserSPECIAL, MarkdownParserSYMBOL, MarkdownParserNUMBER, MarkdownParserWS, MarkdownParserPUNCTUATION, MarkdownParserEXTENDED_UNICODE:
				{
					p.SetState(165)
					p.Word()
				}



			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(168)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(170)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserNEWLINE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Match(MarkdownParserNEWLINE)
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
	p.RuleIndex = MarkdownParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_block

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

func (s *ItemizeContext) AllItem_line() []IItem_lineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IItem_lineContext); ok {
			len++
		}
	}

	tst := make([]IItem_lineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IItem_lineContext); ok {
			tst[i] = t.(IItem_lineContext)
			i++
		}
	}

	return tst
}

func (s *ItemizeContext) Item_line(i int) IItem_lineContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_lineContext); ok {
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

	return t.(IItem_lineContext)
}


func (s *ItemizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterItemize(s)
	}
}

func (s *ItemizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
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

func (s *EnumerateContext) AllNumber_line() []INumber_lineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumber_lineContext); ok {
			len++
		}
	}

	tst := make([]INumber_lineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumber_lineContext); ok {
			tst[i] = t.(INumber_lineContext)
			i++
		}
	}

	return tst
}

func (s *EnumerateContext) Number_line(i int) INumber_lineContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_lineContext); ok {
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

	return t.(INumber_lineContext)
}


func (s *EnumerateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterEnumerate(s)
	}
}

func (s *EnumerateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
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

func (s *VerbatimContext) AllFENCED_BLOCK() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserFENCED_BLOCK)
}

func (s *VerbatimContext) FENCED_BLOCK(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserFENCED_BLOCK, i)
}

func (s *VerbatimContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *VerbatimContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
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
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.EnterVerbatim(s)
	}
}

func (s *VerbatimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownListener); ok {
		listenerT.ExitVerbatim(s)
	}
}



func (p *MarkdownParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MarkdownParserRULE_block)
	var _la int

	var _alt int

	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewItemizeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(175)
						p.Item_line()
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}


	case 2:
		localctx = NewEnumerateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(180)
						p.Number_line()
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}


	case 3:
		localctx = NewVerbatimContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(185)
			p.Match(MarkdownParserFENCED_BLOCK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(186)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 228862) != 0) {
			{
				p.SetState(189)
				p.Verbatim_line()
			}


			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(194)
			p.Match(MarkdownParserFENCED_BLOCK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(195)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
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


