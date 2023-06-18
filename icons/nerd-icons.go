package icons

import "fmt"

type iconInfo2 struct {
	Id    string `json:"id" yaml:"id" toml:"id" xml:"id" ini:"id" csv:"id"`
	Font  string `json:"font" yaml:"font" toml:"font" xml:"font" ini:"font" csv:"font"`
	Name  string `json:"name" yaml:"name" toml:"name" xml:"name" ini:"name" csv:"name"`
	Color color  `json:"color" yaml:"color" toml:"color" xml:"color" ini:"color" csv:"color"`
	Glyph string `json:"glyph" yaml:"glyph" toml:"glyph" xml:"glyph" ini:"glyph" csv:"glyph"`
}

func (i iconInfo2) GetColor() string {
	c := i.Color.RGB(Dark)
	return fmt.Sprintf("\033[38;2;%03d;%03d;%03dm", c[0], c[1], c[2])
}

var iconsByExt = []iconInfo2{
	// Shell
	{Id: "fish", Font: "nerd-icons-devicon", Name: "nf-dev-terminal", Color: nerdIconsLpink},
	{Id: "zsh", Font: "nerd-icons-devicon", Name: "nf-dev-terminal", Color: nerdIconsLcyan},
	{Id: "sh", Font: "nerd-icons-devicon", Name: "nf-dev-terminal", Color: nerdIconsPurple},
	{Id: "bat", Font: "nerd-icons-codicon", Name: "nf-cod-terminal_cmd", Color: nerdIconsLsilver},
	{Id: "cmd", Font: "nerd-icons-codicon", Name: "nf-cod-terminal_cmd", Color: nerdIconsLsilver},
	{Id: "exe", Font: "nerd-icons-codicon", Name: "nf-cod-terminal_cmd", Color: nerdIconsLsilver},
	// Meta
	{Id: "tags", Font: "nerd-icons-octicon", Name: "nf-oct-tag", Color: nerdIconsBlue},
	{Id: "log", Font: "nerd-icons-octicon", Name: "nf-oct-log", Color: nerdIconsMaroon},
	// Config
	{Id: "node", Font: "nerd-icons-devicon", Name: "nf-dev-nodejs_small", Color: nerdIconsGreen},
	{Id: "babelrc", Font: "nerd-icons-mdicon", Name: "nf-seti-babel", Color: nerdIconsYellow},
	{Id: "bashrc", Font: "nerd-icons-mdicon", Name: "nf-dev-terminal", Color: nerdIconsDpink},
	{Id: "bowerrc", Font: "nerd-icons-devicon", Name: "nf-dev-bower", Color: nerdIconsSilver},
	{Id: "cr", Font: "nerd-icons-sucicon", Name: "nf-seti-crystal", Color: nerdIconsYellow},
	{Id: "ecr", Font: "nerd-icons-sucicon", Name: "nf-seti-crystal", Color: nerdIconsYellow},
	{Id: "ini", Font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsYellow},
	{Id: "eslintignore", Font: "nerd-icons-mdicon", Name: "nf-seti-eslint", Color: nerdIconsPurple},
	{Id: "eslint", Font: "nerd-icons-mdicon", Name: "nf-seti-eslint", Color: nerdIconsLpurple},
	{Id: "git", Font: "nerd-icons-devicon", Name: "nf-dev-git", Color: nerdIconsLred},
	{Id: "gitattributes", Font: "nerd-icons-devicon", Name: "nf-dev-git", Color: nerdIconsLred},
	{Id: "gitignore", Font: "nerd-icons-devicon", Name: "nf-dev-git", Color: nerdIconsLred},
	{Id: "gitmodules", Font: "nerd-icons-devicon", Name: "nf-dev-git", Color: nerdIconsLred},
	{Id: "mk", Font: "nerd-icons-devicon", Name: "nf-dev-gnu", Color: nerdIconsDorange},
	// {"cmake"    , font: "nerd-icons-devicon", name: "cmake}, TODO: cmake
	{Id: "dockerignore", Font: "nerd-icons-devicon", Name: "nf-dev-docker", Color: nerdIconsDblue},
	{Id: "xml", Font: "nerd-icons-faicon", Name: "nf-fa-file_code_o", Color: nerdIconsLorange},
	{Id: "json", Font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsYellow},
	{Id: "cson", Font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsYellow},
	{Id: "yml", Font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsDyellow},
	{Id: "yaml", Font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsDyellow},
	{Id: "toml", Font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsOrange},
	{Id: "conf", Font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsDorange},
	{Id: "editorconfig", Font: "nerd-icons-sucicon", Name: "nf-seti-editorconfig", Color: nerdIconsSilver},
	// ?
	{Id: "pkg", Font: "nerd-icons-octicon", Name: "nf-oct-package", Color: nerdIconsDsilver},
	{Id: "rpm", Font: "nerd-icons-octicon", Name: "nf-oct-package", Color: nerdIconsDsilver},
	{Id: "pkgbuild", Font: "nerd-icons-octicon", Name: "nf-oct-package", Color: nerdIconsDsilver},
	{Id: "elc", Font: "nerd-icons-octicon", Name: "nf-oct-file_binary", Color: nerdIconsDsilver},
	{Id: "eln", Font: "nerd-icons-octicon", Name: "nf-oct-file_binary", Color: nerdIconsDsilver},
	{Id: "gz", Font: "nerd-icons-octicon", Name: "nf-oct-file_binary", Color: nerdIconsLmaroon},
	{Id: "zip", Font: "nerd-icons-octicon", Name: "nf-oct-file_zip", Color: nerdIconsLmaroon},
	{Id: "7z", Font: "nerd-icons-octicon", Name: "nf-oct-file_zip", Color: nerdIconsLmaroon},
	{Id: "dat", Font: "nerd-icons-faicon", Name: "nf-fa-bar_chart", Color: nerdIconsCyan},
	{Id: "dmg", Font: "nerd-icons-octicon", Name: "nf-oct-tools", Color: nerdIconsLsilver},
	{Id: "dll", Font: "nerd-icons-faicon", Name: "nf-fa-cogs", Color: nerdIconsSilver},
	{Id: "ds_store", Font: "nerd-icons-faicon", Name: "nf-fa-cogs", Color: nerdIconsSilver},
	{Id: "exe", Font: "nerd-icons-octicon", Name: "nf-oct-file_binary", Color: nerdIconsDsilver},
	// Source Codes
	{Id: "scpt", Font: "nerd-icons-devicon", Name: "nf-dev-apple", Color: nerdIconsPink},
	// {"aup"         , font: "nerd-icons-fileicon", name: "audacity}, TODO: audacity
	{Id: "elm", Font: "nerd-icons-sucicon", Name: "nf-seti-elm", Color: nerdIconsBlue},
	{Id: "erl", Font: "nerd-icons-devicon", Name: "nf-dev-erlang", Color: nerdIconsRed},
	{Id: "hrl", Font: "nerd-icons-devicon", Name: "nf-dev-erlang", Color: nerdIconsDred},
	{Id: "eex", Font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLorange},
	{Id: "leex", Font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLorange},
	{Id: "heex", Font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLorange},
	{Id: "ex", Font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLpurple},
	{Id: "exs", Font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLred},
	{Id: "java", Font: "nerd-icons-devicon", Name: "nf-dev-java", Color: nerdIconsPurple},
	{Id: "gradle", Font: "nerd-icons-sucicon", Name: "nf-seti-gradle", Color: nerdIconsSilver},
	{Id: "ebuild", Font: "nerd-icons-mdicon", Name: "nf-linux-gentoo", Color: nerdIconsCyan},
	{Id: "eclass", Font: "nerd-icons-mdicon", Name: "nf-linux-gentoo", Color: nerdIconsBlue},
	{Id: "go", Font: "nerd-icons-devicon", Name: "nf-seti-go", Color: nerdIconsBlue},
	{Id: "jl", Font: "nerd-icons-sucicon", Name: "nf-seti-julia", Color: nerdIconsPurple},
	{Id: "magik", Font: "nerd-icons-faicon", Name: "nf-fa-magic", Color: nerdIconsBlue},
	// {"matlab"      , font: "nerd-icons-devicon", name: "matlab}, TODO: matlab
	{Id: "nix", Font: "nerd-icons-mdicon", Name: "nf-linux-nixos", Color: nerdIconsBlue},
	{Id: "pl", Font: "nerd-icons-sucicon", Name: "nf-seti-perl", Color: nerdIconsLorange},
	{Id: "pm", Font: "nerd-icons-sucicon", Name: "nf-seti-perl", Color: nerdIconsLorange},
	// {"pl6"         , font: "nerd-icons-devicon", name:"raku}, TODO: raku
	// {"pm6"         , font: "nerd-icons-devicon", name: "raku}, TODO: raku
	{Id: "pod", Font: "nerd-icons-devicon", Name: "nf-dev-perl", Color: nerdIconsLgreen},
	{Id: "php", Font: "nerd-icons-devicon", Name: "nf-dev-php", Color: nerdIconsLsilver},
	// {"pony"        , font: "nerd-icons-devicon", name: "pony}, TODO: pony
	{Id: "ps1", Font: "nerd-icons-mdicon", Name: "nf-seti-powershell", Color: nerdIconsBlue},
	{Id: "pro", Font: "nerd-icons-sucicon", Name: "nf-seti-prolog", Color: nerdIconsLmaroon},
	{Id: "proog", Font: "nerd-icons-sucicon", Name: "nf-seti-prolog", Color: nerdIconsLmaroon},
	{Id: "py", Font: "nerd-icons-devicon", Name: "nf-dev-python", Color: nerdIconsDblue},
	// {"idr"         , font: "nerd-icons-devicon", name: "idris}, TODO: idris
	// {"ipynb"       , font: "nerd-icons-devicon", name: "jupyter}, TODO: jupyter
	{Id: "gem", Font: "nerd-icons-devicon", Name: "nf-dev-ruby", Color: nerdIconsRed},
	// {"raku"        , font: "nerd-icons-devicon", name: "raku}, TODO: raku
	// {"rakumod"     , font: "nerd-icons-devicon", name: "raku}, TODO: raku
	{Id: "rb", Font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsLred},
	{Id: "rs", Font: "nerd-icons-devicon", Name: "nf-dev-rust", Color: nerdIconsMaroon},
	{Id: "rlib", Font: "nerd-icons-devicon", Name: "nf-dev-rust", Color: nerdIconsDmaroon},
	{Id: "r", Font: "nerd-icons-sucicon", Name: "nf-seti-r", Color: nerdIconsLblue},
	{Id: "rd", Font: "nerd-icons-sucicon", Name: "nf-seti-r", Color: nerdIconsLblue},
	{Id: "rdx", Font: "nerd-icons-sucicon", Name: "nf-seti-r", Color: nerdIconsLblue},
	{Id: "rsx", Font: "nerd-icons-sucicon", Name: "nf-seti-r", Color: nerdIconsLblue},
	{Id: "svelte", Font: "nerd-icons-sucicon", Name: "nf-seti-svelte", Color: nerdIconsRed},
	{Id: "gql", Font: "nerd-icons-mdicon", Name: "nf-seti-graphql", Color: nerdIconsDpink},
	{Id: "graphql", Font: "nerd-icons-mdicon", Name: "nf-seti-graphql", Color: nerdIconsDpink},
	// There seems to be a a bug with this font icon which does not
	// let you propertise it without it reverting to being a lower
	// case phi
	{Id: "c", Font: "nerd-icons-sucicon", Name: "nf-custom-c", Color: nerdIconsBlue},
	{Id: "h", Font: "nerd-icons-faicon", Name: "nf-fa-h_square", Color: nerdIconsPurple},
	{Id: "m", Font: "nerd-icons-devicon", Name: "nf-dev-apple", Color: nerdIconsLblue},
	{Id: "mm", Font: "nerd-icons-devicon", Name: "nf-dev-apple", Color: nerdIconsLblue},
	//
	{Id: "cc", Font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsBlue},
	{Id: "cpp", Font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsBlue},
	{Id: "cxx", Font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsBlue},
	{Id: "hh", Font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsPurple},
	{Id: "hpp", Font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsPurple},
	{Id: "hxx", Font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsPurple},
	// Lisps
	{Id: "cl", Font: "nerd-icons-mdicon", Name: "nf-md-yin_yang", Color: nerdIconsRed},
	{Id: "l", Font: "nerd-icons-mdicon", Name: "nf-md-yin_yang", Color: nerdIconsRed},
	{Id: "lisp", Font: "nerd-icons-mdicon", Name: "nf-md-yin_yang", Color: nerdIconsRed},
	{Id: "hy", Font: "nerd-icons-mdicon", Name: "nf-md-yin_yang", Color: nerdIconsRed},
	{Id: "el", Font: "nerd-icons-sucicon", Name: "nf-custom-emacs", Color: nerdIconsPurple},
	{Id: "clj", Font: "nerd-icons-devicon", Name: "nf-dev-clojure", Color: nerdIconsBlue},
	{Id: "cljc", Font: "nerd-icons-devicon", Name: "nf-dev-clojure", Color: nerdIconsBlue},
	// {"cljs"        , font: "nerd-icons-devicon", name: "cljs}, TODO: cljs
	{Id: "coffee", Font: "nerd-icons-devicon", Name: "nf-dev-coffeescript", Color: nerdIconsMaroon},
	{Id: "iced", Font: "nerd-icons-devicon", Name: "nf-dev-coffeescript", Color: nerdIconsLmaroon},
	{Id: "dart", Font: "nerd-icons-devicon", Name: "nf-dev-dart", Color: nerdIconsBlue},
	// {"rkt"         , font: "nerd-icons-devicon", name: "racket}, TODO: racket
	// {"scrbl"       , font: "nerd-icons-devicon", name: "racket}, TODO: racket
	// Stylesheeting
	{Id: "css", Font: "nerd-icons-devicon", Name: "nf-dev-css3", Color: nerdIconsYellow},
	{Id: "scss", Font: "nerd-icons-mdicon", Name: "nf-seti-sass", Color: nerdIconsPink},
	{Id: "sass", Font: "nerd-icons-mdicon", Name: "nf-seti-sass", Color: nerdIconsDpink},
	{Id: "less", Font: "nerd-icons-devicon", Name: "nf-dev-less", Color: nerdIconsDyellow},
	// {"postcss"     , font: "nerd-icons-devicon", name: "postcss}, TODO: postcss
	// {"sss"         , font: "nerd-icons-devicon", name: "postcss}, TODO: postcss
	{Id: "styl", Font: "nerd-icons-devicon", Name: "nf-dev-stylus", Color: nerdIconsLgreen},
	{Id: "csv", Font: "nerd-icons-octicon", Name: "nf-oct-graph", Color: nerdIconsDblue},
	// haskell
	{Id: "hs", Font: "nerd-icons-devicon", Name: "nf-dev-haskell", Color: nerdIconsRed},
	{Id: "chs", Font: "nerd-icons-devicon", Name: "nf-dev-haskell", Color: nerdIconsRed},
	{Id: "lhs", Font: "nerd-icons-devicon", Name: "nf-dev-haskell", Color: nerdIconsRed},
	{Id: "hsc", Font: "nerd-icons-devicon", Name: "nf-dev-haskell", Color: nerdIconsRed},
	// Web modes
	{Id: "inky-haml", Font: "nerd-icons-sucicon", Name: "nf-seti-haml", Color: nerdIconsLyellow},
	{Id: "haml", Font: "nerd-icons-sucicon", Name: "nf-seti-haml", Color: nerdIconsLyellow},
	{Id: "htm", Font: "nerd-icons-devicon", Name: "nf-dev-html5", Color: nerdIconsOrange},
	{Id: "html", Font: "nerd-icons-devicon", Name: "nf-dev-html5", Color: nerdIconsOrange},
	{Id: "inky-er", Font: "nerd-icons-devicon", Name: "nf-dev-html5", Color: nerdIconsLred},
	{Id: "inky-erb", Font: "nerd-icons-devicon", Name: "nf-dev-html5", Color: nerdIconsLred},
	{Id: "erb", Font: "nerd-icons-devicon", Name: "nf-dev-html5", Color: nerdIconsLred},
	// {"hbs"         , font: "nerd-icons-fileicon", name: "moustache}, TODO: moustache
	{Id: "inky-slim", Font: "nerd-icons-codicon", Name: "nf-cod-dashboard", Color: nerdIconsYellow},
	{Id: "slim", Font: "nerd-icons-codicon", Name: "nf-cod-dashboard", Color: nerdIconsYellow},
	{Id: "jade", Font: "nerd-icons-sucicon", Name: "nf-seti-jade", Color: nerdIconsRed},
	{Id: "pug", Font: "nerd-icons-sucicon", Name: "nf-seti-pug", Color: nerdIconsRed},
	// Javascript
	// {"d3js"        , font: "nerd-icons-devicon", name: "d3}, TODO: d3
	{Id: "re", Font: "nerd-icons-sucicon", Name: "nf-seti-reasonml", Color: nerdIconsRedAlt},
	{Id: "rei", Font: "nerd-icons-sucicon", Name: "nf-seti-reasonml", Color: nerdIconsDred},
	{Id: "ml", Font: "nerd-icons-sucicon", Name: "nf-seti-ocaml", Color: nerdIconsLpink},
	{Id: "mli", Font: "nerd-icons-sucicon", Name: "nf-seti-ocaml", Color: nerdIconsDpink},
	{Id: "react", Font: "nerd-icons-devicon", Name: "nf-dev-react", Color: nerdIconsLblue},
	{Id: "ts", Font: "nerd-icons-sucicon", Name: "nf-seti-typescript", Color: nerdIconsBlueAlt},
	{Id: "js", Font: "nerd-icons-devicon", Name: "nf-dev-javascript", Color: nerdIconsYellow},
	{Id: "es", Font: "nerd-icons-devicon", Name: "nf-dev-javascript", Color: nerdIconsYellow},
	{Id: "jsx", Font: "nerd-icons-devicon", Name: "nf-dev-javascript", Color: nerdIconsCyanAlt},
	{Id: "tsx", Font: "nerd-icons-sucicon", Name: "nf-seti-typescript", Color: nerdIconsBlueAlt},
	{Id: "njs", Font: "nerd-icons-mdicon", Name: "nf-dev-nodejs_small", Color: nerdIconsLgreen},
	{Id: "vue", Font: "nerd-icons-sucicon", Name: "nf-seti-vue", Color: nerdIconsLgreen},

	{Id: "sbt", Font: "nerd-icons-sucicon", Name: "nf-seti-sbt", Color: nerdIconsRed},
	{Id: "scala", Font: "nerd-icons-devicon", Name: "nf-dev-scala", Color: nerdIconsRed},
	{Id: "scm", Font: "nerd-icons-mdicon", Name: "nf-md-lambda", Color: nerdIconsRed},
	{Id: "swift", Font: "nerd-icons-devicon", Name: "nf-dev-swift", Color: nerdIconsGreen},

	{Id: "tcl", Font: "nerd-icons-mdicon", Name: "nf-md-feather", Color: nerdIconsDred},

	{Id: "tf", Font: "nerd-icons-mdicon", Name: "nf-seti-terraform", Color: nerdIconsPurpleAlt},
	{Id: "tfvars", Font: "nerd-icons-mdicon", Name: "nf-seti-terraform", Color: nerdIconsPurpleAlt},
	{Id: "tfstate", Font: "nerd-icons-mdicon", Name: "nf-seti-terraform", Color: nerdIconsPurpleAlt},

	{Id: "asm", Font: "nerd-icons-sucicon", Name: "nf-seti-asm", Color: nerdIconsBlue},
	// Verilog(-AMS) and SystemVerilog(-AMS     // Verilog(-AMS) and SystemVerilog(-AMS)
	{Id: "v", Font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	{Id: "vams", Font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	{Id: "sv", Font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	{Id: "sva", Font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	{Id: "svh", Font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	{Id: "svams", Font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	// VHDL(-AMS     // VHDL(-AMS)
	{Id: "vhd", Font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsBlue},
	{Id: "vhdl", Font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsBlue},
	{Id: "vhms", Font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsBlue},
	// Cabal
	// {"cabal"       , font: "nerd-icons-devicon", name: "cabal}, TODO: cabal
	// Kotlin
	{Id: "kt", Font: "nerd-icons-sucicon", Name: "nf-seti-kotlin", Color: nerdIconsOrange},
	{Id: "kts", Font: "nerd-icons-sucicon", Name: "nf-seti-kotlin", Color: nerdIconsOrange},
	// Nimrod
	{Id: "nim", Font: "nerd-icons-sucicon", Name: "nf-seti-nim", Color: nerdIconsYellow},
	{Id: "nims", Font: "nerd-icons-sucicon", Name: "nf-seti-nim", Color: nerdIconsYellow},
	// SQL
	{Id: "sql", Font: "nerd-icons-octicon", Name: "nf-oct-database", Color: nerdIconsSilver},
	// Styles
	// {"styles"      , font: "nerd-icons-devicon", name: "style}, TODO: style
	// Lua
	{Id: "lua", Font: "nerd-icons-sucicon", Name: "nf-seti-lua", Color: nerdIconsDblue},
	// ASCII doc
	// {"adoc"        , font: "nerd-icons-devicon", name: "asciidoc}, TODO: asciidoc
	// {"asciidoc"    , font: "nerd-icons-devicon", name: "asciidoc}, TODO: asciidoc
	// Puppet
	{Id: "pp", Font: "nerd-icons-sucicon", Name: "nf-seti-puppet", Color: nerdIconsYellow},
	// Jinja
	{Id: "j2", Font: "nerd-icons-sucicon", Name: "nf-seti-jinja", Color: nerdIconsSilver},
	{Id: "jinja2", Font: "nerd-icons-sucicon", Name: "nf-seti-jinja", Color: nerdIconsSilver},
	// Docker
	{Id: "dockerfile", Font: "nerd-icons-sucicon", Name: "nf-seti-docker", Color: nerdIconsCyan},
	// Vagrant
	// {"vagrantfile" , font: "nerd-icons-fileicon", name: "vagrant}, TODO: vagrant
	// GLSL
	{Id: "glsl", Font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsBlue},
	{Id: "vert", Font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsBlue},
	{Id: "tesc", Font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsPurple},
	{Id: "tese", Font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsDpurple},
	{Id: "geom", Font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsGreen},
	{Id: "frag", Font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsRed},
	{Id: "comp", Font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsDblue},
	// CUDA
	{Id: "cu", Font: "nerd-icons-sucicon", Name: "nf-custom-c", Color: nerdIconsGreen},
	{Id: "cuh", Font: "nerd-icons-faicon", Name: "nf-fa-h_square", Color: nerdIconsGreen},
	// Fortran
	{Id: "f90", Font: "nerd-icons-mdicon", Name: "F", Color: nerdIconsPurple},
	// C#
	{Id: "cs", Font: "nerd-icons-mdicon", Name: "nf-md-language_csharp", Color: nerdIconsDblue},
	{Id: "csx", Font: "nerd-icons-mdicon", Name: "nf-md-language_csharp", Color: nerdIconsDblue},
	// F#
	{Id: "fs", Font: "nerd-icons-devicon", Name: "nf-dev-fsharp", Color: nerdIconsBlueAlt},
	{Id: "fsi", Font: "nerd-icons-devicon", Name: "nf-dev-fsharp", Color: nerdIconsBlueAlt},
	{Id: "fsx", Font: "nerd-icons-devicon", Name: "nf-dev-fsharp", Color: nerdIconsBlueAlt},
	{Id: "fsscript", Font: "nerd-icons-devicon", Name: "nf-dev-fsharp", Color: nerdIconsBlueAlt},
	// zig
	{Id: "zig", Font: "nerd-icons-sucicon", Name: "nf-seti-zig", Color: nerdIconsOrange},
	// odin
	// {"odin"        , font: "nerd-icons-fileicon", name: "odin}, TODO: odin
	// File Types
	{Id: "ico", Font: "nerd-icons-octicon", Name: "nf-oct-file_media", Color: nerdIconsBlue},
	{Id: "png", Font: "nerd-icons-mdicon", Name: "nf-md-image", Color: nerdIconsOrange},
	{Id: "gif", Font: "nerd-icons-mdicon", Name: "nf-md-image", Color: nerdIconsGreen},
	{Id: "jpeg", Font: "nerd-icons-mdicon", Name: "nf-md-image", Color: nerdIconsDblue},
	{Id: "jpg", Font: "nerd-icons-mdicon", Name: "nf-md-image", Color: nerdIconsDblue},
	{Id: "webp", Font: "nerd-icons-octicon", Name: "nf-oct-file_media", Color: nerdIconsDblue},
	{Id: "xpm", Font: "nerd-icons-octicon", Name: "nf-oct-file_media", Color: nerdIconsDgreen},
	// Audio
	{Id: "mp3", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "wav", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "m4a", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "ogg", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "flac", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "opus", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "au", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "aif", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "aifc", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "aiff", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "svg", Font: "nerd-icons-sucicon", Name: "nf-seti-svg", Color: nerdIconsLgreen},
	// Video
	{Id: "mov", Font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "mp4", Font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "ogv", Font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsDblue},
	{Id: "mpg", Font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "mpeg", Font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "flv", Font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "ogv", Font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsDblue},
	{Id: "mkv", Font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "webm", Font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	// Fonts
	{Id: "ttf", Font: "nerd-icons-faicon", Name: "nf-fa-font", Color: nerdIconsDcyan},
	{Id: "woff", Font: "nerd-icons-faicon", Name: "nf-fa-font", Color: nerdIconsCyan},
	{Id: "woff2", Font: "nerd-icons-faicon", Name: "nf-fa-font", Color: nerdIconsCyan},
	// Archives
	{Id: "tar", Font: "nerd-icons-mdicon", Name: "nf-seti-zip", Color: nerdIconsOrange},
	{Id: "rar", Font: "nerd-icons-mdicon", Name: "nf-seti-zip", Color: nerdIconsOrange},
	{Id: "tgz", Font: "nerd-icons-mdicon", Name: "nf-seti-zip", Color: nerdIconsOrange},
	{Id: "jar", Font: "nerd-icons-devicon", Name: "nf-dev-java", Color: nerdIconsDpurple},
	// Doc
	{Id: "pdf", Font: "nerd-icons-codicon", Name: "nf-cod-file_pdf", Color: nerdIconsDred},
	{Id: "text", Font: "nerd-icons-faicon", Name: "nf-fa-file_text", Color: nerdIconsCyan},
	{Id: "txt", Font: "nerd-icons-faicon", Name: "nf-fa-file_text", Color: nerdIconsCyan},
	{Id: "doc", Font: "nerd-icons-mdicon", Name: "nf-seti-word", Color: nerdIconsBlue},
	{Id: "docx", Font: "nerd-icons-mdicon", Name: "nf-seti-word", Color: nerdIconsBlue},
	{Id: "docm", Font: "nerd-icons-mdicon", Name: "nf-seti-word", Color: nerdIconsBlue},
	{Id: "texi", Font: "nerd-icons-sucicon", Name: "nf-seti-tex", Color: nerdIconsLred},
	{Id: "tex", Font: "nerd-icons-sucicon", Name: "nf-seti-tex", Color: nerdIconsLred},
	{Id: "ltx", Font: "nerd-icons-sucicon", Name: "nf-seti-tex", Color: nerdIconsLred},
	{Id: "dtx", Font: "nerd-icons-sucicon", Name: "nf-seti-tex", Color: nerdIconsLred},
	{Id: "sty", Font: "nerd-icons-sucicon", Name: "nf-seti-tex", Color: nerdIconsLred},
	{Id: "md", Font: "nerd-icons-octicon", Name: "nf-oct-markdown", Color: nerdIconsLblue},
	{Id: "bib", Font: "nerd-icons-mdicon", Name: "nf-fa-book", Color: nerdIconsLblue},
	{Id: "org", Font: "nerd-icons-sucicon", Name: "nf-custom-orgmode", Color: nerdIconsLgreen},
	{Id: "pps", Font: "nerd-icons-mdicon", Name: "nf-fa-file_powerpoint_o", Color: nerdIconsOrange},
	{Id: "ppt", Font: "nerd-icons-mdicon", Name: "nf-fa-file_powerpoint_o", Color: nerdIconsOrange},
	{Id: "pptsx", Font: "nerd-icons-mdicon", Name: "nf-fa-file_powerpoint_o", Color: nerdIconsOrange},
	{Id: "ppttx", Font: "nerd-icons-mdicon", Name: "nf-fa-file_powerpoint_o", Color: nerdIconsOrange},
	{Id: "knt", Font: "nerd-icons-mdicon", Name: "nf-fa-file_powerpoint_o", Color: nerdIconsCyan},
	{Id: "xlsx", Font: "nerd-icons-mdicon", Name: "nf-seti-xls", Color: nerdIconsDgreen},
	{Id: "xlsm", Font: "nerd-icons-mdicon", Name: "nf-seti-xls", Color: nerdIconsDgreen},
	{Id: "xlsb", Font: "nerd-icons-mdicon", Name: "nf-seti-xls", Color: nerdIconsDgreen},
	{Id: "xltx", Font: "nerd-icons-mdicon", Name: "nf-seti-xls", Color: nerdIconsDgreen},
	{Id: "xltm", Font: "nerd-icons-mdicon", Name: "nf-seti-xls", Color: nerdIconsDgreen},
	{Id: "epub", Font: "nerd-icons-mdicon", Name: "nf-fa-book", Color: nerdIconsGreen},
	{Id: "ly", Font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsGreen},
	//
	{Id: "key", Font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsLblue},
	{Id: "pem", Font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsOrange},
	{Id: "p12", Font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsDorange},
	{Id: "crt", Font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsLblue},
	{Id: "pub", Font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsBlue},
	{Id: "gpg", Font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsLblue},
	{Id: "cache", Font: "nerd-icons-octicon", Name: "nf-oct-database", Color: nerdIconsGreen},
}

var regexes = []iconInfo2{
	{Id: "^TAGS$", Font: "nerd-icons-octicon", Name: "nf-oct-tag", Color: nerdIconsBlue},
	{Id: "^TODO$", Font: "nerd-icons-octicon", Name: "nf-oct-checklist", Color: nerdIconsLyellow},
	{Id: "^LICENSE$", Font: "nerd-icons-octicon", Name: "nf-oct-book", Color: nerdIconsBlue},
	{Id: "^readme", Font: "nerd-icons-octicon", Name: "nf-oct-book", Color: nerdIconsLcyan},
	// Config
	{Id: "nginx$", Font: "nerd-icons-devicon", Name: "nf-dev-nginx", Color: nerdIconsDgreen},
	// {"apache$"              , font: "nerd-icons-alltheicon", name: "apache}, TODO: apache
	// C
	{Id: "^Makefile$", Font: "nerd-icons-sucicon", Name: "nf-seti-makefile", Color: nerdIconsDorange},
	{Id: "^CMakeLists.txt$", Font: "nerd-icons-sucicon", Name: "nf-seti-makefile", Color: nerdIconsRed},       // TODO: cmake
	{Id: "^CMakeCache.txt$", Font: "nerd-icons-sucicon", Name: "nf-seti-makefile", Color: nerdIconsBlue},      // TODO: cmakecache
	{Id: "^meson.build$", Font: "nerd-icons-sucicon", Name: "nf-seti-makefile", Color: nerdIconsPurple},       // TODO: meson
	{Id: "^meson_options.txt$", Font: "nerd-icons-sucicon", Name: "nf-seti-makefile", Color: nerdIconsPurple}, // TODO: meson
	// Docker
	{Id: "^\\.?Dockerfile", Font: "nerd-icons-sucicon", Name: "nf-seti-docker", Color: nerdIconsBlue},
	// Homebrew
	{Id: "^Brewfile$", Font: "nerd-icons-faicon", Name: "nf-fa-beer", Color: nerdIconsLsilver},
	// // AWS
	{Id: "^stack.*.json$", Font: "nerd-icons-devicon", Name: "nf-dev-aws", Color: nerdIconsOrange},
	{Id: "^serverless\\.yml$", Font: "nerd-icons-faicon", Name: "nf-fa-bolt", Color: nerdIconsYellow},
	// lock files
	{Id: "~$", Font: "nerd-icons-octicon", Name: "nf-oct-lock", Color: nerdIconsMaroon},
	// Source Codes
	{Id: "^mix.lock$", Font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLyellow},
	// Ruby
	{Id: "^Gemfile\\(\\.lock\\)?$", Font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsRed},
	{Id: "_?test\\.rb$", Font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsRed},
	{Id: "_?test_helper\\.rb$", Font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsDred},
	{Id: "_?spec\\.rb$", Font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsRed},
	{Id: "_?spec_helper\\.rb$", Font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsDred},

	{Id: "-?spec\\.ts$", Font: "nerd-icons-mdicon", Name: "nf-md-language_typescript", Color: nerdIconsBlue},
	{Id: "-?test\\.ts$", Font: "nerd-icons-mdicon", Name: "nf-md-language_typescript", Color: nerdIconsBlue},
	{Id: "-?spec\\.js$", Font: "nerd-icons-mdicon", Name: "nf-md-language_javascript", Color: nerdIconsLpurple},
	{Id: "-?test\\.js$", Font: "nerd-icons-mdicon", Name: "nf-md-language_javascript", Color: nerdIconsLpurple},
	{Id: "-?spec\\.jsx$", Font: "nerd-icons-mdicon", Name: "nf-md-react", Color: nerdIconsBlueAlt},
	{Id: "-?test\\.jsx$", Font: "nerd-icons-mdicon", Name: "nf-md-react", Color: nerdIconsBlueAlt},
	// Git
	{Id: "^MERGE_", Font: "nerd-icons-octicon", Name: "nf-oct-git_merge", Color: nerdIconsRed},
	{Id: "^COMMIT_EDITMSG", Font: "nerd-icons-octicon", Name: "nf-oct-git_commit", Color: nerdIconsRed},
	// Stylesheeting
	{Id: "stylelint", Font: "nerd-icons-sucicon", Name: "nf-seti-stylelint", Color: nerdIconsLyellow},
	// JavaScript
	{Id: "^package.json$", Font: "nerd-icons-devicon", Name: "nf-dev-npm", Color: nerdIconsRed},
	{Id: "^package.lock.json$", Font: "nerd-icons-devicon", Name: "nf-dev-npm", Color: nerdIconsDred},
	{Id: "^yarn\\.lock", Font: "nerd-icons-sucicon", Name: "nf-seti-yarn", Color: nerdIconsBlueAlt},
	{Id: "\\.npmignore$", Font: "nerd-icons-devicon", Name: "nf-dev-npm", Color: nerdIconsDred},
	{Id: "^bower.json$", Font: "nerd-icons-devicon", Name: "nf-dev-bower", Color: nerdIconsLorange},
	{Id: "^gulpfile", Font: "nerd-icons-devicon", Name: "nf-dev-gulp", Color: nerdIconsLred},
	{Id: "^gruntfile", Font: "nerd-icons-devicon", Name: "nf-dev-grunt", Color: nerdIconsLyellow},
	{Id: "^webpack", Font: "nerd-icons-mdicon", Name: "nf-md-webpack", Color: nerdIconsLblue},
	// Go
	{Id: "^go.mod$", Font: "nerd-icons-sucicon", Name: "nf-seti-config", Color: nerdIconsBlueAlt},
	{Id: "^go.work$", Font: "nerd-icons-sucicon", Name: "nf-seti-config", Color: nerdIconsBlueAlt},
	// Emacs
	{Id: "bookmark", Font: "nerd-icons-octicon", Name: "nf-oct-bookmark", Color: nerdIconsLpink},
	{Id: "^\\*scratch\\*$", Font: "nerd-icons-faicon", Name: "nf-fa-sticky_note", Color: nerdIconsLyellow},
	{Id: "^\\*scratch.*", Font: "nerd-icons-faicon", Name: "nf-fa-sticky_note", Color: nerdIconsYellow},
	{Id: "^\\*new-tab\\*$", Font: "nerd-icons-mdicon", Name: "nf-md-star", Color: nerdIconsCyan},
	{Id: "^\\.", Font: "nerd-icons-octicon", Name: "nf-oct-gear"},
}

var defFolderIcons = []iconInfo2{
	{Id: "trash", Font: "nerd-icons-faicon", Name: "nf-fa-trash_o"},
	{Id: "dropbox", Font: "nerd-icons-faicon", Name: "nf-fa-dropbox"},
	{Id: "google[ _-]drive", Font: "nerd-icons-mdicon", Name: "nf-md-folder_google_drive"},
	{Id: "github", Font: "nerd-icons-sucicon", Name: "nf-custom-folder_github"},
	{Id: "^atom$", Font: "nerd-icons-devicon", Name: "nf-dev-atom"},
	{Id: "documents", Font: "nerd-icons-mdicon", Name: "nf-md-folder_file"},
	{Id: "download", Font: "nerd-icons-mdicon", Name: "nf-md-folder_download"},
	{Id: "desktop", Font: "nerd-icons-octicon", Name: "nf-oct-device_desktop"},
	{Id: "pictures", Font: "nerd-icons-mdicon", Name: "nf-md-folder_image"},
	{Id: "photos", Font: "nerd-icons-faicon", Name: "nf-fa-camera_retro"},
	{Id: "music", Font: "nerd-icons-mdicon", Name: "nf-md-folder_music"},
	{Id: "movies", Font: "nerd-icons-faicon", Name: "nf-fa-film"},
	{Id: "code", Font: "nerd-icons-octicon", Name: "nf-oct-code"},
	{Id: "workspace", Font: "nerd-icons-octicon", Name: "nf-oct-code"},
	// {"test"             , font: "nerd-icons-devicon", name: "test-dir"},
	{Id: "\\.git", Font: "nerd-icons-sucicon", Name: "nf-custom-folder_git"},
	{Id: "\\.config", Font: "nerd-icons-sucicon", Name: "nf-custom-folder_config"},
	{Id: ".?", Font: "nerd-icons-sucicon", Name: "nf-custom-folder_oct"},
}
