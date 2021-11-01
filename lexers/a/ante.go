package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ante lexer.
var Ante = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "Ante",
		Aliases:   []string{"ante", "an"},
		Filenames: []string{"*.an"},
		MimeTypes: []string{"text/x-ante"},
	},
	anteRules,
))

func anteRules() Rules {
	return Rules{
		"root": {
			{`\s+`, Text, nil},
			{`//(?![!#$%&*+./<=>?@^|_~:\\]).*?$`, CommentSingle, nil},
			{`\{-`, CommentMultiline, Push("comment")},
			{`\berror\b`, NameException, nil},
			{`\b(if|else|import|with|in|do|try|given|as|can|effect|handler|handle|continue|return|fn|opaque|impl|match|trait|module|recur|type|and|loop|do|then|not|or|_)(?!\')\b`, KeywordReserved, nil},
			{`'[^\\]'`, LiteralStringChar, nil},
			{`^[_\p{Ll}][\w\']*`, NameFunction, nil},
			{`'?[_\p{Ll}][\w']*`, Name, nil},
			{`('')?[\p{Lu}][\w\']*`, KeywordType, nil},
			{`(')[\p{Lu}][\w\']*`, KeywordType, nil},
			{`(')\[[^\]]*\]`, KeywordType, nil},
			{`(')\([^)]*\)`, KeywordType, nil},
			{`\\(?![:!#$%&*+.\\/<=>?@^|~-]+)`, NameFunction, nil},
			{`(<-|::|->|=>|=)(?![:!#$%&*+.\\/<=>?@^|~-]+)`, OperatorWord, nil},
			{`:[:!#$%&*+.\\/<=>?@^|~-]*`, KeywordType, nil},
			{`[:!#$%&*+.\\/<=>?@^|~-]+`, Operator, nil},
			{`\d+[eE][+-]?\d+`, LiteralNumberFloat, nil},
			{`\d+\.\d+([eE][+-]?\d+)?`, LiteralNumberFloat, nil},
			{`0[oO][0-7]+`, LiteralNumberOct, nil},
			{`0[xX][\da-fA-F]+`, LiteralNumberHex, nil},
			{`\d+`, LiteralNumberInteger, nil},
			{`'`, LiteralStringChar, Push("character")},
			{`"`, LiteralString, Push("string")},
			{`\[\]`, KeywordType, nil},
			{`\(\)`, NameBuiltin, nil},
			{"[][(),;`{}]", Punctuation, nil},
		},
		"funclist": {
			{`\s+`, Text, nil},
			{`[\p{Lu}]\w*`, KeywordType, nil},
			{`(_[\w\']+|[\p{Ll}][\w\']*)`, NameFunction, nil},
			{`--(?![!#$%&*+./<=>?@^|_~:\\]).*?$`, CommentSingle, nil},
			{`\{-`, CommentMultiline, Push("comment")},
			{`,`, Punctuation, nil},
			{`[:!#$%&*+.\\/<=>?@^|~-]+`, Operator, nil},
			{`\(`, Punctuation, Push("funclist", "funclist")},
			{`\)`, Punctuation, Pop(2)},
		},
		"comment": {
			{`/\*`, CommentMultiline, Push()},
			{`\*/`, CommentMultiline, Pop(1)},
		},
		"character": {
			{`[^\\']'`, LiteralStringChar, Pop(1)},
			{`\\`, LiteralStringEscape, Push("escape")},
			{`'`, LiteralStringChar, Pop(1)},
		},
		"string": {
			{`[^\\"]+`, LiteralString, nil},
			{`\\`, LiteralStringEscape, Push("escape")},
			{`"`, LiteralString, Pop(1)},
		},
		"escape": {
			{`[abfnrtv"\'&\\]`, LiteralStringEscape, Pop(1)},
			{`\^[][\p{Lu}@^_]`, LiteralStringEscape, Pop(1)},
			{`NUL|SOH|[SE]TX|EOT|ENQ|ACK|BEL|BS|HT|LF|VT|FF|CR|S[OI]|DLE|DC[1-4]|NAK|SYN|ETB|CAN|EM|SUB|ESC|[FGRU]S|SP|DEL`, LiteralStringEscape, Pop(1)},
			{`o[0-7]+`, LiteralStringEscape, Pop(1)},
			{`x[\da-fA-F]+`, LiteralStringEscape, Pop(1)},
			{`\d+`, LiteralStringEscape, Pop(1)},
			{`\s+\\`, LiteralStringEscape, Pop(1)},
		},
	}
}
