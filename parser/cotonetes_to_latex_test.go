/*

These tests ensure that the latex produced by the system will produce the same markdown that originated that latex when ingested by the system

The otherway around is not tested as the latex ingestion step needs to be flexible to accept any valid type of latex (e.g.: Introducing a blank line somewhere in the input latex will not affect the produced markdown, and therefore the blank line in the input latex will not affect the output latex. The latex ingestion can also fix some latex issues such as escaping special chars, so these are better tested as part of the latex parser tests, instead of a full-lifecycle test

*/

package parser

import (
	"testing"
	"cotonetes/utils"
)

func mdLxParserTest(t *testing.T, test_input utils.TestInput) {
	_, latex_folder_path := LatexToCotonetes(t, test_input.Markdown)

	LatexParserTest(t, latex_folder_path, test_input.Markdown)
}

func TestMdLxTextOnly(t *testing.T) {
	mdLxParserTest(t, utils.TdTextOnly)
}

func TestMdLxSpecialCharsString(t *testing.T) {
	mdLxParserTest(t, utils.TdTitleSpecialChars)
}

func TestMdLxBoldString(t *testing.T) {
	mdLxParserTest(t, utils.TdNoteBoldText)
}

func TestMdLxUrl(t *testing.T) {
	mdLxParserTest(t, utils.TdNoteUrl)
}

func TestMdLxNewline(t *testing.T) {
	mdLxParserTest(t, utils.TdNoteNewline)
}

func TestMdLxItemize(t *testing.T) {
	mdLxParserTest(t, utils.TdNoteItemize)
}
