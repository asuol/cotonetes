// Code generated from Markdown.g4 by ANTLR 4.13.2. DO NOT EDIT.

package markdown_parser // Markdown
import "github.com/antlr4-go/antlr/v4"

// BaseMarkdownListener is a complete listener for a parse tree produced by MarkdownParser.
type BaseMarkdownListener struct{}

var _ MarkdownListener = &BaseMarkdownListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMarkdownListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMarkdownListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMarkdownListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMarkdownListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMarkdown is called when production markdown is entered.
func (s *BaseMarkdownListener) EnterMarkdown(ctx *MarkdownContext) {}

// ExitMarkdown is called when production markdown is exited.
func (s *BaseMarkdownListener) ExitMarkdown(ctx *MarkdownContext) {}

// EnterText is called when production text is entered.
func (s *BaseMarkdownListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseMarkdownListener) ExitText(ctx *TextContext) {}

// EnterEmpty_line is called when production empty_line is entered.
func (s *BaseMarkdownListener) EnterEmpty_line(ctx *Empty_lineContext) {}

// ExitEmpty_line is called when production empty_line is exited.
func (s *BaseMarkdownListener) ExitEmpty_line(ctx *Empty_lineContext) {}

// EnterBold is called when production bold is entered.
func (s *BaseMarkdownListener) EnterBold(ctx *BoldContext) {}

// ExitBold is called when production bold is exited.
func (s *BaseMarkdownListener) ExitBold(ctx *BoldContext) {}

// EnterUrl is called when production url is entered.
func (s *BaseMarkdownListener) EnterUrl(ctx *UrlContext) {}

// ExitUrl is called when production url is exited.
func (s *BaseMarkdownListener) ExitUrl(ctx *UrlContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseMarkdownListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseMarkdownListener) ExitSymbol(ctx *SymbolContext) {}

// EnterLetter is called when production letter is entered.
func (s *BaseMarkdownListener) EnterLetter(ctx *LetterContext) {}

// ExitLetter is called when production letter is exited.
func (s *BaseMarkdownListener) ExitLetter(ctx *LetterContext) {}

// EnterPunctuation is called when production punctuation is entered.
func (s *BaseMarkdownListener) EnterPunctuation(ctx *PunctuationContext) {}

// ExitPunctuation is called when production punctuation is exited.
func (s *BaseMarkdownListener) ExitPunctuation(ctx *PunctuationContext) {}

// EnterBackslash is called when production backslash is entered.
func (s *BaseMarkdownListener) EnterBackslash(ctx *BackslashContext) {}

// ExitBackslash is called when production backslash is exited.
func (s *BaseMarkdownListener) ExitBackslash(ctx *BackslashContext) {}

// EnterSquare is called when production square is entered.
func (s *BaseMarkdownListener) EnterSquare(ctx *SquareContext) {}

// ExitSquare is called when production square is exited.
func (s *BaseMarkdownListener) ExitSquare(ctx *SquareContext) {}

// EnterRound is called when production round is entered.
func (s *BaseMarkdownListener) EnterRound(ctx *RoundContext) {}

// ExitRound is called when production round is exited.
func (s *BaseMarkdownListener) ExitRound(ctx *RoundContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseMarkdownListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseMarkdownListener) ExitNumber(ctx *NumberContext) {}

// EnterWs is called when production ws is entered.
func (s *BaseMarkdownListener) EnterWs(ctx *WsContext) {}

// ExitWs is called when production ws is exited.
func (s *BaseMarkdownListener) ExitWs(ctx *WsContext) {}

// EnterUrl_title is called when production url_title is entered.
func (s *BaseMarkdownListener) EnterUrl_title(ctx *Url_titleContext) {}

// ExitUrl_title is called when production url_title is exited.
func (s *BaseMarkdownListener) ExitUrl_title(ctx *Url_titleContext) {}

// EnterUrl_value is called when production url_value is entered.
func (s *BaseMarkdownListener) EnterUrl_value(ctx *Url_valueContext) {}

// ExitUrl_value is called when production url_value is exited.
func (s *BaseMarkdownListener) ExitUrl_value(ctx *Url_valueContext) {}

// EnterItem_line is called when production item_line is entered.
func (s *BaseMarkdownListener) EnterItem_line(ctx *Item_lineContext) {}

// ExitItem_line is called when production item_line is exited.
func (s *BaseMarkdownListener) ExitItem_line(ctx *Item_lineContext) {}

// EnterNumber_line is called when production number_line is entered.
func (s *BaseMarkdownListener) EnterNumber_line(ctx *Number_lineContext) {}

// ExitNumber_line is called when production number_line is exited.
func (s *BaseMarkdownListener) ExitNumber_line(ctx *Number_lineContext) {}

// EnterVerbatim_line is called when production verbatim_line is entered.
func (s *BaseMarkdownListener) EnterVerbatim_line(ctx *Verbatim_lineContext) {}

// ExitVerbatim_line is called when production verbatim_line is exited.
func (s *BaseMarkdownListener) ExitVerbatim_line(ctx *Verbatim_lineContext) {}

// EnterItemize is called when production itemize is entered.
func (s *BaseMarkdownListener) EnterItemize(ctx *ItemizeContext) {}

// ExitItemize is called when production itemize is exited.
func (s *BaseMarkdownListener) ExitItemize(ctx *ItemizeContext) {}

// EnterEnumerate is called when production enumerate is entered.
func (s *BaseMarkdownListener) EnterEnumerate(ctx *EnumerateContext) {}

// ExitEnumerate is called when production enumerate is exited.
func (s *BaseMarkdownListener) ExitEnumerate(ctx *EnumerateContext) {}

// EnterVerbatim is called when production verbatim is entered.
func (s *BaseMarkdownListener) EnterVerbatim(ctx *VerbatimContext) {}

// ExitVerbatim is called when production verbatim is exited.
func (s *BaseMarkdownListener) ExitVerbatim(ctx *VerbatimContext) {}
