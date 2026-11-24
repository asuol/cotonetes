grammar Markdown;

markdown
    // block must be before text, to ensure that a line starting with a * ou number is recognized as a block and not text
    : (block | text | empty_line)* EOF
    ;

text
    // star and url must be after bold and url to ensure a line starting with * or [ is recognized as bold or url instead of star or word
    : (bold | url | star | word)+ NEWLINE?
    ;

empty_line
    : NEWLINE+
    ;

bold
    : ASTERISK ASTERISK word+ ASTERISK ASTERISK
    ;

url
    : '[' url_title ']' '(' url_value ')'
    ;

star
    // Cannot be part of the word rule to ensure that e.g. the word rule within bold does not match the bold end (**), in the case of two consecutive bold blocks (**bold 1** non bold **bold2**) avoids matching the whole string as a single bold instead of two
    : ASTERISK
    ;

word
    : SYMBOL          #symbol
    | SPECIAL         #special
    | (LETTER | DOT )         #letter
    | (PUNCTUATION | EXTENDED_UNICODE )    #punctuation
    | BACKSLASH       #backslash
    | ('[' | ']')     #square
    | ( '(' | ')')    #round
    | NUMBER          #number
    | WS              #ws
    ;

url_title
    : (
          SYMBOL
        | SPECIAL
        | LETTER
        | ASTERISK
        | DOT
        | PUNCTUATION
        | EXTENDED_UNICODE
        | ( '(' | ')' )
        | NUMBER
        | WS
    )+
    ;

url_value
    : (
          SYMBOL
        | SPECIAL
        | LETTER
        | ASTERISK
        | DOT
        | PUNCTUATION
        | EXTENDED_UNICODE
        | ( '[' | ']')
        | NUMBER
        | WS
    )+
    ;

item_line
    : WS* ASTERISK WS+ (bold | url | star | word)+ NEWLINE*
    ;

number_line
    : WS* NUMBER DOT WS* (bold | url | star | word)+ NEWLINE*
    ;

verbatim_line
    : (star | word)+ NEWLINE
    | NEWLINE
    ;

block
    : item_line+                                                #itemize
    | number_line+                                              #enumerate
    | FENCED_BLOCK NEWLINE? verbatim_line+ FENCED_BLOCK NEWLINE?  #verbatim
    ;

LETTER
    : [\p{L}]+
    ;

ASTERISK
    : '*'
    ;

DOT
    : '.'
    ;

BACKSLASH
    : '\\'
    ;

FENCED_BLOCK
    : BACKTICK BACKTICK BACKTICK
    ;

fragment BACKTICK
    : '`'
    ;

SPECIAL
    // all chars that require a dedicated latex command
    : '~'
    | '^'
    | '|'
    ;

SYMBOL
    // all chars that can be directly escaped in latex
    : '#'
    | '$'
    | '%'
    | '&'
    | '>'
    | '<'
    | '€'
    | '_'
    | '{'
    | '}'
    ;

NUMBER
    : '-'? INT ('.' [0-9]+)?
    ;

fragment INT
    // integer part forbids leading 0s (e.g. `01`)
    : '0'
    | [1-9] [0-9]*
    ;

// Required to detect end of blocks (bullet list or enumerations), as well as to preserve lines
NEWLINE
    : '\n'
    ;

WS
    : [ \t]+
    ;

CR
    : '\r' -> skip
    ;

PUNCTUATION
    // all the punctuation that does not require escaping in latex
    // except the parens, which interfere with the markdown URL syntax
    : [\p{Pc}]
    | [\p{Pd}]
    | [\p{Pi}]
    | [\p{Pf}]
    | [\p{Po}]
    | ( '(' | ')' | '[' | ']' )
    | [\p{S}]
    | [\p{Z}]
    ;

EXTENDED_UNICODE
    : [\u0001-\uFFFF]
    ;
