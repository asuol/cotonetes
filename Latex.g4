grammar Latex;

latex
    : note_title note_url note_created note_updated line_break NEWLINE* note_text EOF
    ;

note_title
    : '\\textbf{Title:}' WS* word+ '\\\\' NEWLINE?
    ;

note_url
    : '\\textbf{URL:} \\url{' word+ '}' '\\\\' NEWLINE?
    ;

note_created
    : '\\textbf{Created:}' WS* word+ '\\\\' NEWLINE?
    ;

note_updated
    : '\\textbf{Last Updated:}' WS* word+ '\\\\' NEWLINE?
    ;

note_text
    : (text | block | line_break | empty_line)* 
    ;

text
    : (tag | word)+ NEWLINE?
    ;

line_break
    : '\\\\' WS*
    ;

empty_line
    : NEWLINE+
    ;

escaped_word
    : '\\' content=(LETTER | SYMBOL)+ space=WS* NEWLINE?
    ;

tag
    : '\\' name=LETTER+ '{' word+ '}'
    ;

word
    : escaped_word    #escaped
    | LETTER          #letter
    | PUNCTUATION     #punctuation
    | NUMBER          #number
    | WS              #ws
    ;

verbatim_content
    : word            #verbatim_word
    | SYMBOL          #verbatim_symbol
    | line_break      #verbatim_linebreak
    ;

verbatim_line
    : verbatim_content* NEWLINE
    ;

block_line
    : word+ NEWLINE+
    ; 

block_item
    : '\\item' WS* block_line
    ;

block
    : '\\begin{itemize}' NEWLINE* block_item* '\\end{itemize}' NEWLINE?       #itemize
    | '\\begin{enumerate}' NEWLINE* block_item* '\\end{enumerate}' NEWLINE?   #enumerate
    | '\\begin{verbatim}' NEWLINE? verbatim_line* '\\end{verbatim}' NEWLINE?     #verbatim
    ;

LETTER
    : [\p{L}]+
    ;

PUNCTUATION
    // all the punctuation that does not require escaping in latex
    : (
          '.'
        | ','
        | ';'
        | ':'
        | '!'
        | '?'
        | '-'
        | '\''
        | '"'
        | '('
        | ')'
        | '['
        | ']'
        | '/'
        | '@'
        | '*'
        | '+'
        | '='
        )+
    ;

SYMBOL
    : '#'
    | '$'
    | '%'
    | '&'
    | '_'
    | '^'
    | '>'
    | '<'
    ;

NUMBER
    : '-'? INT ('.' [0-9]+)?
    ;

fragment INT
    // integer part forbids leading 0s (e.g. `01`)
    : '0'
    | [1-9] [0-9]*
    ;

// Required to detect empty lines, which counts as a paragraph in latex
NEWLINE
    : '\n'
    ;

WS
    : [ \t]+
    ;

CR
    : '\r' -> skip
    ;
