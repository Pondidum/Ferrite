package zmk

import (
	"io"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var parser = participle.MustBuild[File](
	participle.UseLookahead(2),
	participle.Unquote("String"),
)

type File struct {
	pos lexer.Position

	Includes []*Include `parser:"@@+"`
	Defines  []*Define  `parser:"@@*"`
	Configs  []*Config  `parser:"@@*"`
	Device   *Device    `parser:"'/' '{' @@ '}'';'"`
}

type Include struct {
	pos lexer.Position

	Value string `parser:"'#'Ident'<'@((Ident ('-' Ident)? '/'?)* ('.' Ident))'>'"`
}

type Define struct {
	pos lexer.Position

	Value string `parser:"'#'Ident @Ident (Ident|Int)"`
}

type Config struct {
	pos lexer.Position

	Behavior string     `parser:"'&'@Ident '{'"`
	Values   []*Options `parser:"@@* '}'';'"`
}

type Options struct {
	Key   *string `parser:"@Ident(@'-' @Ident)* '='"`
	Value *Value  `parser:"@@ ';'"`
}

type Value struct {
	String *string `parser:"  @String"`
	Number *int    `parser:"| '<'@Int'>'"`
}

type Device struct {
	pos lexer.Position

	Combos *Combos `parser:"('combos' '{' @@)?"`
	Keymap *Keymap `parser:"'keymap' '{' @@"`
}

type Combos struct {
	pos lexer.Position

	Compatible string   `parser:"'compatible' '=' @String';'"`
	Combos     []*Combo `parser:"@@* '}'';'"`
}

type Combo struct {
	pos lexer.Position

	Name         string      `parser:"@Ident '{'"`
	Timeout      int32       `parser:"'timeout''-''ms' '=' '<'@Int'>'';'"`
	KeyPositions []*List     `parser:"'key''-''positions' '=' '<'@@+'>'';'"`
	Layers       []*List     `parser:"('layers' '=' '<'@@+'>'';')?"`
	Bindings     []*Behavior `parser:"'bindings' '=' '<'@@+'>'';' '}'';'"`
}

type Keymap struct {
	pos lexer.Position

	Compatible string   `parser:"'compatible' '=' @String';'"`
	Layers     []*Layer `parser:"@@+ '}'';'"`
}

type Layer struct {
	pos lexer.Position

	Name           string      `parser:"@Ident '{'"`
	Bindings       []*Behavior `parser:"'bindings' '=' '<'@@+'>'';'"`
	SensorBindings []*Behavior `parser:"('sensor''-''bindings' '=' '<'@@+'>'';')?"`
	EndBrace       string      `parser:" '}'';'"`
}

type List struct {
	Number  *int32  `parser:"@Int"`
	KeyCode *string `parser:"| @(Ident('('Ident('('Ident')')?')')?)"`
}

type Behavior struct {
	pos lexer.Position

	Action string  `parser:"'&'@Ident"`
	Params []*List `parser:"@@*"`
}

func Parse(r io.Reader) (*File, error) {
	return parser.Parse("", r)
}

func Enbf() string {
	return parser.String()
}
