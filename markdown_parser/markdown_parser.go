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
    "", "'**'", "'['", "']('", "')'", "'*'", "'```'", "", "", "", "'\\'", 
    "", "", "", "", "'\\n'", "", "'\\r'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "LETTER", "NUMBER_ITEM", "PUNCTUATION", 
    "BACKSLASH", "SQUARE_PARENS", "ROUND_PARENS", "SYMBOL", "NUMBER", "NEWLINE", 
    "WS", "CR",
  }
  staticData.RuleNames = []string{
    "markdown", "text", "empty_line", "bold", "url", "word", "url_title", 
    "url_value", "item_line", "number_line", "verbatim_line", "block",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 17, 163, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0, 31, 9, 
	0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 4, 1, 38, 8, 1, 11, 1, 12, 1, 39, 1, 1, 
	3, 1, 43, 8, 1, 1, 2, 4, 2, 46, 8, 2, 11, 2, 12, 2, 47, 1, 3, 1, 3, 4, 
	3, 52, 8, 3, 11, 3, 12, 3, 53, 1, 3, 1, 3, 1, 4, 1, 4, 4, 4, 60, 8, 4, 
	11, 4, 12, 4, 61, 1, 4, 1, 4, 4, 4, 66, 8, 4, 11, 4, 12, 4, 67, 1, 4, 1, 
	4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 80, 8, 5, 1, 6, 
	4, 6, 83, 8, 6, 11, 6, 12, 6, 84, 1, 7, 4, 7, 88, 8, 7, 11, 7, 12, 7, 89, 
	1, 8, 1, 8, 5, 8, 94, 8, 8, 10, 8, 12, 8, 97, 9, 8, 1, 8, 4, 8, 100, 8, 
	8, 11, 8, 12, 8, 101, 1, 8, 5, 8, 105, 8, 8, 10, 8, 12, 8, 108, 9, 8, 1, 
	9, 1, 9, 5, 9, 112, 8, 9, 10, 9, 12, 9, 115, 9, 9, 1, 9, 4, 9, 118, 8, 
	9, 11, 9, 12, 9, 119, 1, 9, 5, 9, 123, 8, 9, 10, 9, 12, 9, 126, 9, 9, 1, 
	10, 4, 10, 129, 8, 10, 11, 10, 12, 10, 130, 1, 10, 1, 10, 1, 10, 3, 10, 
	136, 8, 10, 1, 11, 4, 11, 139, 8, 11, 11, 11, 12, 11, 140, 1, 11, 4, 11, 
	144, 8, 11, 11, 11, 12, 11, 145, 1, 11, 1, 11, 3, 11, 150, 8, 11, 1, 11, 
	4, 11, 153, 8, 11, 11, 11, 12, 11, 154, 1, 11, 1, 11, 3, 11, 159, 8, 11, 
	3, 11, 161, 8, 11, 1, 11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 
	20, 22, 0, 2, 4, 0, 7, 7, 9, 9, 12, 14, 16, 16, 5, 0, 7, 7, 9, 9, 11, 11, 
	13, 14, 16, 16, 185, 0, 29, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 45, 1, 0, 
	0, 0, 6, 49, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 79, 1, 0, 0, 0, 12, 82, 
	1, 0, 0, 0, 14, 87, 1, 0, 0, 0, 16, 91, 1, 0, 0, 0, 18, 109, 1, 0, 0, 0, 
	20, 135, 1, 0, 0, 0, 22, 160, 1, 0, 0, 0, 24, 28, 3, 2, 1, 0, 25, 28, 3, 
	22, 11, 0, 26, 28, 3, 4, 2, 0, 27, 24, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 
	27, 26, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 
	0, 0, 0, 30, 32, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 33, 5, 0, 0, 1, 33, 
	1, 1, 0, 0, 0, 34, 38, 3, 10, 5, 0, 35, 38, 3, 6, 3, 0, 36, 38, 3, 8, 4, 
	0, 37, 34, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 36, 1, 0, 0, 0, 38, 39, 
	1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 42, 1, 0, 0, 0, 
	41, 43, 5, 15, 0, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 3, 1, 
	0, 0, 0, 44, 46, 5, 15, 0, 0, 45, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 
	45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 5, 1, 0, 0, 0, 49, 51, 5, 1, 0, 
	0, 50, 52, 3, 10, 5, 0, 51, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 
	1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 5, 1, 0, 0, 
	56, 7, 1, 0, 0, 0, 57, 59, 5, 2, 0, 0, 58, 60, 3, 12, 6, 0, 59, 58, 1, 
	0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 
	63, 1, 0, 0, 0, 63, 65, 5, 3, 0, 0, 64, 66, 3, 14, 7, 0, 65, 64, 1, 0, 
	0, 0, 66, 67, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 
	1, 0, 0, 0, 69, 70, 5, 4, 0, 0, 70, 9, 1, 0, 0, 0, 71, 80, 5, 13, 0, 0, 
	72, 80, 5, 7, 0, 0, 73, 80, 5, 9, 0, 0, 74, 80, 5, 10, 0, 0, 75, 80, 5, 
	11, 0, 0, 76, 80, 5, 12, 0, 0, 77, 80, 5, 14, 0, 0, 78, 80, 5, 16, 0, 0, 
	79, 71, 1, 0, 0, 0, 79, 72, 1, 0, 0, 0, 79, 73, 1, 0, 0, 0, 79, 74, 1, 
	0, 0, 0, 79, 75, 1, 0, 0, 0, 79, 76, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 
	78, 1, 0, 0, 0, 80, 11, 1, 0, 0, 0, 81, 83, 7, 0, 0, 0, 82, 81, 1, 0, 0, 
	0, 83, 84, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 13, 
	1, 0, 0, 0, 86, 88, 7, 1, 0, 0, 87, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 
	89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 15, 1, 0, 0, 0, 91, 95, 5, 
	5, 0, 0, 92, 94, 5, 16, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 
	93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 
	0, 98, 100, 3, 10, 5, 0, 99, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 
	99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 106, 1, 0, 0, 0, 103, 105, 5, 
	15, 0, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 
	0, 106, 107, 1, 0, 0, 0, 107, 17, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 
	113, 5, 8, 0, 0, 110, 112, 5, 16, 0, 0, 111, 110, 1, 0, 0, 0, 112, 115, 
	1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 117, 1, 0, 
	0, 0, 115, 113, 1, 0, 0, 0, 116, 118, 3, 10, 5, 0, 117, 116, 1, 0, 0, 0, 
	118, 119, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 
	124, 1, 0, 0, 0, 121, 123, 5, 15, 0, 0, 122, 121, 1, 0, 0, 0, 123, 126, 
	1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 19, 1, 0, 
	0, 0, 126, 124, 1, 0, 0, 0, 127, 129, 3, 10, 5, 0, 128, 127, 1, 0, 0, 0, 
	129, 130, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 
	132, 1, 0, 0, 0, 132, 133, 5, 15, 0, 0, 133, 136, 1, 0, 0, 0, 134, 136, 
	5, 15, 0, 0, 135, 128, 1, 0, 0, 0, 135, 134, 1, 0, 0, 0, 136, 21, 1, 0, 
	0, 0, 137, 139, 3, 16, 8, 0, 138, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 
	140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 161, 1, 0, 0, 0, 142, 
	144, 3, 18, 9, 0, 143, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 143, 
	1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 161, 1, 0, 0, 0, 147, 149, 5, 6, 
	0, 0, 148, 150, 5, 15, 0, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 
	150, 152, 1, 0, 0, 0, 151, 153, 3, 20, 10, 0, 152, 151, 1, 0, 0, 0, 153, 
	154, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 
	1, 0, 0, 0, 156, 158, 5, 6, 0, 0, 157, 159, 5, 15, 0, 0, 158, 157, 1, 0, 
	0, 0, 158, 159, 1, 0, 0, 0, 159, 161, 1, 0, 0, 0, 160, 138, 1, 0, 0, 0, 
	160, 143, 1, 0, 0, 0, 160, 147, 1, 0, 0, 0, 161, 23, 1, 0, 0, 0, 26, 27, 
	29, 37, 39, 42, 47, 53, 61, 67, 79, 84, 89, 95, 101, 106, 113, 119, 124, 
	130, 135, 140, 145, 149, 154, 158, 160,
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
	MarkdownParserT__4 = 5
	MarkdownParserT__5 = 6
	MarkdownParserLETTER = 7
	MarkdownParserNUMBER_ITEM = 8
	MarkdownParserPUNCTUATION = 9
	MarkdownParserBACKSLASH = 10
	MarkdownParserSQUARE_PARENS = 11
	MarkdownParserROUND_PARENS = 12
	MarkdownParserSYMBOL = 13
	MarkdownParserNUMBER = 14
	MarkdownParserNEWLINE = 15
	MarkdownParserWS = 16
	MarkdownParserCR = 17
)

// MarkdownParser rules.
const (
	MarkdownParserRULE_markdown = 0
	MarkdownParserRULE_text = 1
	MarkdownParserRULE_empty_line = 2
	MarkdownParserRULE_bold = 3
	MarkdownParserRULE_url = 4
	MarkdownParserRULE_word = 5
	MarkdownParserRULE_url_title = 6
	MarkdownParserRULE_url_value = 7
	MarkdownParserRULE_item_line = 8
	MarkdownParserRULE_number_line = 9
	MarkdownParserRULE_verbatim_line = 10
	MarkdownParserRULE_block = 11
)

// IMarkdownContext is an interface to support dynamic dispatch.
type IMarkdownContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllText() []ITextContext
	Text(i int) ITextContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
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
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 131046) != 0) {
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MarkdownParserT__0, MarkdownParserT__1, MarkdownParserLETTER, MarkdownParserPUNCTUATION, MarkdownParserBACKSLASH, MarkdownParserSQUARE_PARENS, MarkdownParserROUND_PARENS, MarkdownParserSYMBOL, MarkdownParserNUMBER, MarkdownParserWS:
			{
				p.SetState(24)
				p.Text()
			}


		case MarkdownParserT__4, MarkdownParserT__5, MarkdownParserNUMBER_ITEM:
			{
				p.SetState(25)
				p.Block()
			}


		case MarkdownParserNEWLINE:
			{
				p.SetState(26)
				p.Empty_line()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
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
	AllWord() []IWordContext
	Word(i int) IWordContext
	AllBold() []IBoldContext
	Bold(i int) IBoldContext
	AllUrl() []IUrlContext
	Url(i int) IUrlContext
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				p.SetState(37)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case MarkdownParserLETTER, MarkdownParserPUNCTUATION, MarkdownParserBACKSLASH, MarkdownParserSQUARE_PARENS, MarkdownParserROUND_PARENS, MarkdownParserSYMBOL, MarkdownParserNUMBER, MarkdownParserWS:
					{
						p.SetState(34)
						p.Word()
					}


				case MarkdownParserT__0:
					{
						p.SetState(35)
						p.Bold()
					}


				case MarkdownParserT__1:
					{
						p.SetState(36)
						p.Url()
					}



				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(41)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(44)
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

		p.SetState(47)
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
		p.SetState(49)
		p.Match(MarkdownParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 97920) != 0) {
		{
			p.SetState(50)
			p.Word()
		}


		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Match(MarkdownParserT__0)
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
	AllUrl_title() []IUrl_titleContext
	Url_title(i int) IUrl_titleContext
	AllUrl_value() []IUrl_valueContext
	Url_value(i int) IUrl_valueContext

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

func (s *UrlContext) AllUrl_title() []IUrl_titleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUrl_titleContext); ok {
			len++
		}
	}

	tst := make([]IUrl_titleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUrl_titleContext); ok {
			tst[i] = t.(IUrl_titleContext)
			i++
		}
	}

	return tst
}

func (s *UrlContext) Url_title(i int) IUrl_titleContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrl_titleContext); ok {
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

	return t.(IUrl_titleContext)
}

func (s *UrlContext) AllUrl_value() []IUrl_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUrl_valueContext); ok {
			len++
		}
	}

	tst := make([]IUrl_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUrl_valueContext); ok {
			tst[i] = t.(IUrl_valueContext)
			i++
		}
	}

	return tst
}

func (s *UrlContext) Url_value(i int) IUrl_valueContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrl_valueContext); ok {
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(MarkdownParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 94848) != 0) {
		{
			p.SetState(58)
			p.Url_title()
		}


		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.Match(MarkdownParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 92800) != 0) {
		{
			p.SetState(64)
			p.Url_value()
		}


		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(69)
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

func (s *SquareContext) SQUARE_PARENS() antlr.TerminalNode {
	return s.GetToken(MarkdownParserSQUARE_PARENS, 0)
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

func (s *RoundContext) ROUND_PARENS() antlr.TerminalNode {
	return s.GetToken(MarkdownParserROUND_PARENS, 0)
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
	p.EnterRule(localctx, 10, MarkdownParserRULE_word)
	p.SetState(79)
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


	case MarkdownParserLETTER:
		localctx = NewLetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(MarkdownParserLETTER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserPUNCTUATION:
		localctx = NewPunctuationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Match(MarkdownParserPUNCTUATION)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserBACKSLASH:
		localctx = NewBackslashContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)
			p.Match(MarkdownParserBACKSLASH)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserSQUARE_PARENS:
		localctx = NewSquareContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)
			p.Match(MarkdownParserSQUARE_PARENS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserROUND_PARENS:
		localctx = NewRoundContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(76)
			p.Match(MarkdownParserROUND_PARENS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(77)
			p.Match(MarkdownParserNUMBER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserWS:
		localctx = NewWsContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(78)
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
	AllLETTER() []antlr.TerminalNode
	LETTER(i int) antlr.TerminalNode
	AllPUNCTUATION() []antlr.TerminalNode
	PUNCTUATION(i int) antlr.TerminalNode
	AllROUND_PARENS() []antlr.TerminalNode
	ROUND_PARENS(i int) antlr.TerminalNode
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

func (s *Url_titleContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserLETTER)
}

func (s *Url_titleContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserLETTER, i)
}

func (s *Url_titleContext) AllPUNCTUATION() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserPUNCTUATION)
}

func (s *Url_titleContext) PUNCTUATION(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserPUNCTUATION, i)
}

func (s *Url_titleContext) AllROUND_PARENS() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserROUND_PARENS)
}

func (s *Url_titleContext) ROUND_PARENS(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserROUND_PARENS, i)
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
	p.EnterRule(localctx, 12, MarkdownParserRULE_url_title)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(81)
					_la = p.GetTokenStream().LA(1)

					if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 94848) != 0)) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}




		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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


// IUrl_valueContext is an interface to support dynamic dispatch.
type IUrl_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSYMBOL() []antlr.TerminalNode
	SYMBOL(i int) antlr.TerminalNode
	AllLETTER() []antlr.TerminalNode
	LETTER(i int) antlr.TerminalNode
	AllPUNCTUATION() []antlr.TerminalNode
	PUNCTUATION(i int) antlr.TerminalNode
	AllSQUARE_PARENS() []antlr.TerminalNode
	SQUARE_PARENS(i int) antlr.TerminalNode
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

func (s *Url_valueContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserLETTER)
}

func (s *Url_valueContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserLETTER, i)
}

func (s *Url_valueContext) AllPUNCTUATION() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserPUNCTUATION)
}

func (s *Url_valueContext) PUNCTUATION(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserPUNCTUATION, i)
}

func (s *Url_valueContext) AllSQUARE_PARENS() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserSQUARE_PARENS)
}

func (s *Url_valueContext) SQUARE_PARENS(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserSQUARE_PARENS, i)
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
	p.EnterRule(localctx, 14, MarkdownParserRULE_url_value)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(86)
					_la = p.GetTokenStream().LA(1)

					if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 92800) != 0)) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}




		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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


// IItem_lineContext is an interface to support dynamic dispatch.
type IItem_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
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

func (s *Item_lineContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWS)
}

func (s *Item_lineContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWS, i)
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
	p.EnterRule(localctx, 16, MarkdownParserRULE_item_line)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(MarkdownParserT__4)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(92)
				p.Match(MarkdownParserWS)
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
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(98)
					p.Word()
				}




		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(103)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	NUMBER_ITEM() antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
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

func (s *Number_lineContext) NUMBER_ITEM() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNUMBER_ITEM, 0)
}

func (s *Number_lineContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWS)
}

func (s *Number_lineContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWS, i)
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
	p.EnterRule(localctx, 18, MarkdownParserRULE_number_line)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(MarkdownParserNUMBER_ITEM)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(110)
				p.Match(MarkdownParserWS)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(116)
					p.Word()
				}




		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(124)
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
				p.SetState(121)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(126)
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


// IVerbatim_lineContext is an interface to support dynamic dispatch.
type IVerbatim_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEWLINE() antlr.TerminalNode
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
	p.EnterRule(localctx, 20, MarkdownParserRULE_verbatim_line)
	var _la int

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserLETTER, MarkdownParserPUNCTUATION, MarkdownParserBACKSLASH, MarkdownParserSQUARE_PARENS, MarkdownParserROUND_PARENS, MarkdownParserSYMBOL, MarkdownParserNUMBER, MarkdownParserWS:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 97920) != 0) {
			{
				p.SetState(127)
				p.Word()
			}


			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(132)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MarkdownParserNEWLINE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(134)
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
	p.EnterRule(localctx, 22, MarkdownParserRULE_block)
	var _la int

	var _alt int

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserT__4:
		localctx = NewItemizeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(137)
						p.Item_line()
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(140)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}


	case MarkdownParserNUMBER_ITEM:
		localctx = NewEnumerateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(142)
						p.Number_line()
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}


	case MarkdownParserT__5:
		localctx = NewVerbatimContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(147)
			p.Match(MarkdownParserT__5)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(148)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 130688) != 0) {
			{
				p.SetState(151)
				p.Verbatim_line()
			}


			p.SetState(154)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(156)
			p.Match(MarkdownParserT__5)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(157)
				p.Match(MarkdownParserNEWLINE)
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


