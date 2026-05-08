// Code generated from Latex.g4 by ANTLR 4.13.2. DO NOT EDIT.

package latex_parser // Latex
import "github.com/antlr4-go/antlr/v4"

// BaseLatexListener is a complete listener for a parse tree produced by LatexParser.
type BaseLatexListener struct{}

var _ LatexListener = &BaseLatexListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLatexListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLatexListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLatexListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLatexListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLatex is called when production latex is entered.
func (s *BaseLatexListener) EnterLatex(ctx *LatexContext) {}

// ExitLatex is called when production latex is exited.
func (s *BaseLatexListener) ExitLatex(ctx *LatexContext) {}

// EnterNote_title is called when production note_title is entered.
func (s *BaseLatexListener) EnterNote_title(ctx *Note_titleContext) {}

// ExitNote_title is called when production note_title is exited.
func (s *BaseLatexListener) ExitNote_title(ctx *Note_titleContext) {}

// EnterNote_url is called when production note_url is entered.
func (s *BaseLatexListener) EnterNote_url(ctx *Note_urlContext) {}

// ExitNote_url is called when production note_url is exited.
func (s *BaseLatexListener) ExitNote_url(ctx *Note_urlContext) {}

// EnterNote_created is called when production note_created is entered.
func (s *BaseLatexListener) EnterNote_created(ctx *Note_createdContext) {}

// ExitNote_created is called when production note_created is exited.
func (s *BaseLatexListener) ExitNote_created(ctx *Note_createdContext) {}

// EnterNote_updated is called when production note_updated is entered.
func (s *BaseLatexListener) EnterNote_updated(ctx *Note_updatedContext) {}

// ExitNote_updated is called when production note_updated is exited.
func (s *BaseLatexListener) ExitNote_updated(ctx *Note_updatedContext) {}

// EnterNote_text is called when production note_text is entered.
func (s *BaseLatexListener) EnterNote_text(ctx *Note_textContext) {}

// ExitNote_text is called when production note_text is exited.
func (s *BaseLatexListener) ExitNote_text(ctx *Note_textContext) {}

// EnterText is called when production text is entered.
func (s *BaseLatexListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseLatexListener) ExitText(ctx *TextContext) {}

// EnterLine_break is called when production line_break is entered.
func (s *BaseLatexListener) EnterLine_break(ctx *Line_breakContext) {}

// ExitLine_break is called when production line_break is exited.
func (s *BaseLatexListener) ExitLine_break(ctx *Line_breakContext) {}

// EnterEmpty_line is called when production empty_line is entered.
func (s *BaseLatexListener) EnterEmpty_line(ctx *Empty_lineContext) {}

// ExitEmpty_line is called when production empty_line is exited.
func (s *BaseLatexListener) ExitEmpty_line(ctx *Empty_lineContext) {}

// EnterEscaped_word is called when production escaped_word is entered.
func (s *BaseLatexListener) EnterEscaped_word(ctx *Escaped_wordContext) {}

// ExitEscaped_word is called when production escaped_word is exited.
func (s *BaseLatexListener) ExitEscaped_word(ctx *Escaped_wordContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseLatexListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseLatexListener) ExitTag(ctx *TagContext) {}

// EnterEscaped is called when production escaped is entered.
func (s *BaseLatexListener) EnterEscaped(ctx *EscapedContext) {}

// ExitEscaped is called when production escaped is exited.
func (s *BaseLatexListener) ExitEscaped(ctx *EscapedContext) {}

// EnterLetter is called when production letter is entered.
func (s *BaseLatexListener) EnterLetter(ctx *LetterContext) {}

// ExitLetter is called when production letter is exited.
func (s *BaseLatexListener) ExitLetter(ctx *LetterContext) {}

// EnterPunctuation is called when production punctuation is entered.
func (s *BaseLatexListener) EnterPunctuation(ctx *PunctuationContext) {}

// ExitPunctuation is called when production punctuation is exited.
func (s *BaseLatexListener) ExitPunctuation(ctx *PunctuationContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseLatexListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseLatexListener) ExitNumber(ctx *NumberContext) {}

// EnterWs is called when production ws is entered.
func (s *BaseLatexListener) EnterWs(ctx *WsContext) {}

// ExitWs is called when production ws is exited.
func (s *BaseLatexListener) ExitWs(ctx *WsContext) {}

// EnterVerbatim_word is called when production verbatim_word is entered.
func (s *BaseLatexListener) EnterVerbatim_word(ctx *Verbatim_wordContext) {}

// ExitVerbatim_word is called when production verbatim_word is exited.
func (s *BaseLatexListener) ExitVerbatim_word(ctx *Verbatim_wordContext) {}

// EnterVerbatim_symbol is called when production verbatim_symbol is entered.
func (s *BaseLatexListener) EnterVerbatim_symbol(ctx *Verbatim_symbolContext) {}

// ExitVerbatim_symbol is called when production verbatim_symbol is exited.
func (s *BaseLatexListener) ExitVerbatim_symbol(ctx *Verbatim_symbolContext) {}

// EnterVerbatim_linebreak is called when production verbatim_linebreak is entered.
func (s *BaseLatexListener) EnterVerbatim_linebreak(ctx *Verbatim_linebreakContext) {}

// ExitVerbatim_linebreak is called when production verbatim_linebreak is exited.
func (s *BaseLatexListener) ExitVerbatim_linebreak(ctx *Verbatim_linebreakContext) {}

// EnterVerbatim_line is called when production verbatim_line is entered.
func (s *BaseLatexListener) EnterVerbatim_line(ctx *Verbatim_lineContext) {}

// ExitVerbatim_line is called when production verbatim_line is exited.
func (s *BaseLatexListener) ExitVerbatim_line(ctx *Verbatim_lineContext) {}

// EnterBlock_line is called when production block_line is entered.
func (s *BaseLatexListener) EnterBlock_line(ctx *Block_lineContext) {}

// ExitBlock_line is called when production block_line is exited.
func (s *BaseLatexListener) ExitBlock_line(ctx *Block_lineContext) {}

// EnterBlock_item is called when production block_item is entered.
func (s *BaseLatexListener) EnterBlock_item(ctx *Block_itemContext) {}

// ExitBlock_item is called when production block_item is exited.
func (s *BaseLatexListener) ExitBlock_item(ctx *Block_itemContext) {}

// EnterItemize is called when production itemize is entered.
func (s *BaseLatexListener) EnterItemize(ctx *ItemizeContext) {}

// ExitItemize is called when production itemize is exited.
func (s *BaseLatexListener) ExitItemize(ctx *ItemizeContext) {}

// EnterEnumerate is called when production enumerate is entered.
func (s *BaseLatexListener) EnterEnumerate(ctx *EnumerateContext) {}

// ExitEnumerate is called when production enumerate is exited.
func (s *BaseLatexListener) ExitEnumerate(ctx *EnumerateContext) {}

// EnterVerbatim is called when production verbatim is entered.
func (s *BaseLatexListener) EnterVerbatim(ctx *VerbatimContext) {}

// ExitVerbatim is called when production verbatim is exited.
func (s *BaseLatexListener) ExitVerbatim(ctx *VerbatimContext) {}
