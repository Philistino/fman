package icons

import "fmt"

type iconInfo2 struct {
	Id    string `json:"id" yaml:"id" toml:"id" xml:"id" ini:"id" csv:"id"`
	font  string `json:"font" yaml:"font" toml:"font" xml:"font" ini:"font" csv:"font"`
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
	{Id: "fish", font: "nerd-icons-devicon", Name: "nf-dev-terminal", Color: nerdIconsLpink},
	{Id: "zsh", font: "nerd-icons-devicon", Name: "nf-dev-terminal", Color: nerdIconsLcyan},
	{Id: "sh", font: "nerd-icons-devicon", Name: "nf-dev-terminal", Color: nerdIconsPurple},
	{Id: "bat", font: "nerd-icons-codicon", Name: "nf-cod-terminal_cmd", Color: nerdIconsLsilver},
	{Id: "cmd", font: "nerd-icons-codicon", Name: "nf-cod-terminal_cmd", Color: nerdIconsLsilver},
	{Id: "exe", font: "nerd-icons-codicon", Name: "nf-cod-terminal_cmd", Color: nerdIconsLsilver},
	// Meta
	{Id: "tags", font: "nerd-icons-octicon", Name: "nf-oct-tag", Color: nerdIconsBlue},
	{Id: "log", font: "nerd-icons-octicon", Name: "nf-oct-log", Color: nerdIconsMaroon},
	// Config
	{Id: "node", font: "nerd-icons-devicon", Name: "nf-dev-nodejs_small", Color: nerdIconsGreen},
	{Id: "babelrc", font: "nerd-icons-mdicon", Name: "nf-seti-babel", Color: nerdIconsYellow},
	{Id: "bashrc", font: "nerd-icons-mdicon", Name: "nf-dev-terminal", Color: nerdIconsDpink},
	{Id: "bowerrc", font: "nerd-icons-devicon", Name: "nf-dev-bower", Color: nerdIconsSilver},
	{Id: "cr", font: "nerd-icons-sucicon", Name: "nf-seti-crystal", Color: nerdIconsYellow},
	{Id: "ecr", font: "nerd-icons-sucicon", Name: "nf-seti-crystal", Color: nerdIconsYellow},
	{Id: "ini", font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsYellow},
	{Id: "eslintignore", font: "nerd-icons-mdicon", Name: "nf-seti-eslint", Color: nerdIconsPurple},
	{Id: "eslint", font: "nerd-icons-mdicon", Name: "nf-seti-eslint", Color: nerdIconsLpurple},
	{Id: "git", font: "nerd-icons-devicon", Name: "nf-dev-git", Color: nerdIconsLred},
	{Id: "gitattributes", font: "nerd-icons-devicon", Name: "nf-dev-git", Color: nerdIconsLred},
	{Id: "gitignore", font: "nerd-icons-devicon", Name: "nf-dev-git", Color: nerdIconsLred},
	{Id: "gitmodules", font: "nerd-icons-devicon", Name: "nf-dev-git", Color: nerdIconsLred},
	{Id: "mk", font: "nerd-icons-devicon", Name: "nf-dev-gnu", Color: nerdIconsDorange},
	// {"cmake"    , font: "nerd-icons-devicon", name: "cmake}, TODO: cmake
	{Id: "dockerignore", font: "nerd-icons-devicon", Name: "nf-dev-docker", Color: nerdIconsDblue},
	{Id: "xml", font: "nerd-icons-faicon", Name: "nf-fa-file_code_o", Color: nerdIconsLorange},
	{Id: "json", font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsYellow},
	{Id: "cson", font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsYellow},
	{Id: "yml", font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsDyellow},
	{Id: "yaml", font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsDyellow},
	{Id: "toml", font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsOrange},
	{Id: "conf", font: "nerd-icons-codicon", Name: "nf-cod-settings", Color: nerdIconsDorange},
	{Id: "editorconfig", font: "nerd-icons-sucicon", Name: "nf-seti-editorconfig", Color: nerdIconsSilver},
	// ?
	{Id: "pkg", font: "nerd-icons-octicon", Name: "nf-oct-package", Color: nerdIconsDsilver},
	{Id: "rpm", font: "nerd-icons-octicon", Name: "nf-oct-package", Color: nerdIconsDsilver},
	{Id: "pkgbuild", font: "nerd-icons-octicon", Name: "nf-oct-package", Color: nerdIconsDsilver},
	{Id: "elc", font: "nerd-icons-octicon", Name: "nf-oct-file_binary", Color: nerdIconsDsilver},
	{Id: "eln", font: "nerd-icons-octicon", Name: "nf-oct-file_binary", Color: nerdIconsDsilver},
	{Id: "gz", font: "nerd-icons-octicon", Name: "nf-oct-file_binary", Color: nerdIconsLmaroon},
	{Id: "zip", font: "nerd-icons-octicon", Name: "nf-oct-file_zip", Color: nerdIconsLmaroon},
	{Id: "7z", font: "nerd-icons-octicon", Name: "nf-oct-file_zip", Color: nerdIconsLmaroon},
	{Id: "dat", font: "nerd-icons-faicon", Name: "nf-fa-bar_chart", Color: nerdIconsCyan},
	{Id: "dmg", font: "nerd-icons-octicon", Name: "nf-oct-tools", Color: nerdIconsLsilver},
	{Id: "dll", font: "nerd-icons-faicon", Name: "nf-fa-cogs", Color: nerdIconsSilver},
	{Id: "ds_store", font: "nerd-icons-faicon", Name: "nf-fa-cogs", Color: nerdIconsSilver},
	{Id: "exe", font: "nerd-icons-octicon", Name: "nf-oct-file_binary", Color: nerdIconsDsilver},
	// Source Codes
	{Id: "scpt", font: "nerd-icons-devicon", Name: "nf-dev-apple", Color: nerdIconsPink},
	// {"aup"         , font: "nerd-icons-fileicon", name: "audacity}, TODO: audacity
	{Id: "elm", font: "nerd-icons-sucicon", Name: "nf-seti-elm", Color: nerdIconsBlue},
	{Id: "erl", font: "nerd-icons-devicon", Name: "nf-dev-erlang", Color: nerdIconsRed},
	{Id: "hrl", font: "nerd-icons-devicon", Name: "nf-dev-erlang", Color: nerdIconsDred},
	{Id: "eex", font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLorange},
	{Id: "leex", font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLorange},
	{Id: "heex", font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLorange},
	{Id: "ex", font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLpurple},
	{Id: "exs", font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLred},
	{Id: "java", font: "nerd-icons-devicon", Name: "nf-dev-java", Color: nerdIconsPurple},
	{Id: "gradle", font: "nerd-icons-sucicon", Name: "nf-seti-gradle", Color: nerdIconsSilver},
	{Id: "ebuild", font: "nerd-icons-mdicon", Name: "nf-linux-gentoo", Color: nerdIconsCyan},
	{Id: "eclass", font: "nerd-icons-mdicon", Name: "nf-linux-gentoo", Color: nerdIconsBlue},
	{Id: "go", font: "nerd-icons-devicon", Name: "nf-seti-go", Color: nerdIconsBlue},
	{Id: "jl", font: "nerd-icons-sucicon", Name: "nf-seti-julia", Color: nerdIconsPurple},
	{Id: "magik", font: "nerd-icons-faicon", Name: "nf-fa-magic", Color: nerdIconsBlue},
	// {"matlab"      , font: "nerd-icons-devicon", name: "matlab}, TODO: matlab
	{Id: "nix", font: "nerd-icons-mdicon", Name: "nf-linux-nixos", Color: nerdIconsBlue},
	{Id: "pl", font: "nerd-icons-sucicon", Name: "nf-seti-perl", Color: nerdIconsLorange},
	{Id: "pm", font: "nerd-icons-sucicon", Name: "nf-seti-perl", Color: nerdIconsLorange},
	// {"pl6"         , font: "nerd-icons-devicon", name:"raku}, TODO: raku
	// {"pm6"         , font: "nerd-icons-devicon", name: "raku}, TODO: raku
	{Id: "pod", font: "nerd-icons-devicon", Name: "nf-dev-perl", Color: nerdIconsLgreen},
	{Id: "php", font: "nerd-icons-devicon", Name: "nf-dev-php", Color: nerdIconsLsilver},
	// {"pony"        , font: "nerd-icons-devicon", name: "pony}, TODO: pony
	{Id: "ps1", font: "nerd-icons-mdicon", Name: "nf-seti-powershell", Color: nerdIconsBlue},
	{Id: "pro", font: "nerd-icons-sucicon", Name: "nf-seti-prolog", Color: nerdIconsLmaroon},
	{Id: "proog", font: "nerd-icons-sucicon", Name: "nf-seti-prolog", Color: nerdIconsLmaroon},
	{Id: "py", font: "nerd-icons-devicon", Name: "nf-dev-python", Color: nerdIconsDblue},
	// {"idr"         , font: "nerd-icons-devicon", name: "idris}, TODO: idris
	// {"ipynb"       , font: "nerd-icons-devicon", name: "jupyter}, TODO: jupyter
	{Id: "gem", font: "nerd-icons-devicon", Name: "nf-dev-ruby", Color: nerdIconsRed},
	// {"raku"        , font: "nerd-icons-devicon", name: "raku}, TODO: raku
	// {"rakumod"     , font: "nerd-icons-devicon", name: "raku}, TODO: raku
	{Id: "rb", font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsLred},
	{Id: "rs", font: "nerd-icons-devicon", Name: "nf-dev-rust", Color: nerdIconsMaroon},
	{Id: "rlib", font: "nerd-icons-devicon", Name: "nf-dev-rust", Color: nerdIconsDmaroon},
	{Id: "r", font: "nerd-icons-sucicon", Name: "nf-seti-r", Color: nerdIconsLblue},
	{Id: "rd", font: "nerd-icons-sucicon", Name: "nf-seti-r", Color: nerdIconsLblue},
	{Id: "rdx", font: "nerd-icons-sucicon", Name: "nf-seti-r", Color: nerdIconsLblue},
	{Id: "rsx", font: "nerd-icons-sucicon", Name: "nf-seti-r", Color: nerdIconsLblue},
	{Id: "svelte", font: "nerd-icons-sucicon", Name: "nf-seti-svelte", Color: nerdIconsRed},
	{Id: "gql", font: "nerd-icons-mdicon", Name: "nf-seti-graphql", Color: nerdIconsDpink},
	{Id: "graphql", font: "nerd-icons-mdicon", Name: "nf-seti-graphql", Color: nerdIconsDpink},
	// There seems to be a a bug with this font icon which does not
	// let you propertise it without it reverting to being a lower
	// case phi
	{Id: "c", font: "nerd-icons-sucicon", Name: "nf-custom-c", Color: nerdIconsBlue},
	{Id: "h", font: "nerd-icons-faicon", Name: "nf-fa-h_square", Color: nerdIconsPurple},
	{Id: "m", font: "nerd-icons-devicon", Name: "nf-dev-apple", Color: nerdIconsLblue},
	{Id: "mm", font: "nerd-icons-devicon", Name: "nf-dev-apple", Color: nerdIconsLblue},
	//
	{Id: "cc", font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsBlue},
	{Id: "cpp", font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsBlue},
	{Id: "cxx", font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsBlue},
	{Id: "hh", font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsPurple},
	{Id: "hpp", font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsPurple},
	{Id: "hxx", font: "nerd-icons-sucicon", Name: "nf-custom-cpp", Color: nerdIconsPurple},
	// Lisps
	{Id: "cl", font: "nerd-icons-mdicon", Name: "nf-md-yin_yang", Color: nerdIconsRed},
	{Id: "l", font: "nerd-icons-mdicon", Name: "nf-md-yin_yang", Color: nerdIconsRed},
	{Id: "lisp", font: "nerd-icons-mdicon", Name: "nf-md-yin_yang", Color: nerdIconsRed},
	{Id: "hy", font: "nerd-icons-mdicon", Name: "nf-md-yin_yang", Color: nerdIconsRed},
	{Id: "el", font: "nerd-icons-sucicon", Name: "nf-custom-emacs", Color: nerdIconsPurple},
	{Id: "clj", font: "nerd-icons-devicon", Name: "nf-dev-clojure", Color: nerdIconsBlue},
	{Id: "cljc", font: "nerd-icons-devicon", Name: "nf-dev-clojure", Color: nerdIconsBlue},
	// {"cljs"        , font: "nerd-icons-devicon", name: "cljs}, TODO: cljs
	{Id: "coffee", font: "nerd-icons-devicon", Name: "nf-dev-coffeescript", Color: nerdIconsMaroon},
	{Id: "iced", font: "nerd-icons-devicon", Name: "nf-dev-coffeescript", Color: nerdIconsLmaroon},
	{Id: "dart", font: "nerd-icons-devicon", Name: "nf-dev-dart", Color: nerdIconsBlue},
	// {"rkt"         , font: "nerd-icons-devicon", name: "racket}, TODO: racket
	// {"scrbl"       , font: "nerd-icons-devicon", name: "racket}, TODO: racket
	// Stylesheeting
	{Id: "css", font: "nerd-icons-devicon", Name: "nf-dev-css3", Color: nerdIconsYellow},
	{Id: "scss", font: "nerd-icons-mdicon", Name: "nf-seti-sass", Color: nerdIconsPink},
	{Id: "sass", font: "nerd-icons-mdicon", Name: "nf-seti-sass", Color: nerdIconsDpink},
	{Id: "less", font: "nerd-icons-devicon", Name: "nf-dev-less", Color: nerdIconsDyellow},
	// {"postcss"     , font: "nerd-icons-devicon", name: "postcss}, TODO: postcss
	// {"sss"         , font: "nerd-icons-devicon", name: "postcss}, TODO: postcss
	{Id: "styl", font: "nerd-icons-devicon", Name: "nf-dev-stylus", Color: nerdIconsLgreen},
	{Id: "csv", font: "nerd-icons-octicon", Name: "nf-oct-graph", Color: nerdIconsDblue},
	// haskell
	{Id: "hs", font: "nerd-icons-devicon", Name: "nf-dev-haskell", Color: nerdIconsRed},
	{Id: "chs", font: "nerd-icons-devicon", Name: "nf-dev-haskell", Color: nerdIconsRed},
	{Id: "lhs", font: "nerd-icons-devicon", Name: "nf-dev-haskell", Color: nerdIconsRed},
	{Id: "hsc", font: "nerd-icons-devicon", Name: "nf-dev-haskell", Color: nerdIconsRed},
	// Web modes
	{Id: "inky-haml", font: "nerd-icons-sucicon", Name: "nf-seti-haml", Color: nerdIconsLyellow},
	{Id: "haml", font: "nerd-icons-sucicon", Name: "nf-seti-haml", Color: nerdIconsLyellow},
	{Id: "htm", font: "nerd-icons-devicon", Name: "nf-dev-html5", Color: nerdIconsOrange},
	{Id: "html", font: "nerd-icons-devicon", Name: "nf-dev-html5", Color: nerdIconsOrange},
	{Id: "inky-er", font: "nerd-icons-devicon", Name: "nf-dev-html5", Color: nerdIconsLred},
	{Id: "inky-erb", font: "nerd-icons-devicon", Name: "nf-dev-html5", Color: nerdIconsLred},
	{Id: "erb", font: "nerd-icons-devicon", Name: "nf-dev-html5", Color: nerdIconsLred},
	// {"hbs"         , font: "nerd-icons-fileicon", name: "moustache}, TODO: moustache
	{Id: "inky-slim", font: "nerd-icons-codicon", Name: "nf-cod-dashboard", Color: nerdIconsYellow},
	{Id: "slim", font: "nerd-icons-codicon", Name: "nf-cod-dashboard", Color: nerdIconsYellow},
	{Id: "jade", font: "nerd-icons-sucicon", Name: "nf-seti-jade", Color: nerdIconsRed},
	{Id: "pug", font: "nerd-icons-sucicon", Name: "nf-seti-pug", Color: nerdIconsRed},
	// Javascript
	// {"d3js"        , font: "nerd-icons-devicon", name: "d3}, TODO: d3
	{Id: "re", font: "nerd-icons-sucicon", Name: "nf-seti-reasonml", Color: nerdIconsRedAlt},
	{Id: "rei", font: "nerd-icons-sucicon", Name: "nf-seti-reasonml", Color: nerdIconsDred},
	{Id: "ml", font: "nerd-icons-sucicon", Name: "nf-seti-ocaml", Color: nerdIconsLpink},
	{Id: "mli", font: "nerd-icons-sucicon", Name: "nf-seti-ocaml", Color: nerdIconsDpink},
	{Id: "react", font: "nerd-icons-devicon", Name: "nf-dev-react", Color: nerdIconsLblue},
	{Id: "ts", font: "nerd-icons-sucicon", Name: "nf-seti-typescript", Color: nerdIconsBlueAlt},
	{Id: "js", font: "nerd-icons-devicon", Name: "nf-dev-javascript", Color: nerdIconsYellow},
	{Id: "es", font: "nerd-icons-devicon", Name: "nf-dev-javascript", Color: nerdIconsYellow},
	{Id: "jsx", font: "nerd-icons-devicon", Name: "nf-dev-javascript", Color: nerdIconsCyanAlt},
	{Id: "tsx", font: "nerd-icons-sucicon", Name: "nf-seti-typescript", Color: nerdIconsBlueAlt},
	{Id: "njs", font: "nerd-icons-mdicon", Name: "nf-dev-nodejs_small", Color: nerdIconsLgreen},
	{Id: "vue", font: "nerd-icons-sucicon", Name: "nf-seti-vue", Color: nerdIconsLgreen},

	{Id: "sbt", font: "nerd-icons-sucicon", Name: "nf-seti-sbt", Color: nerdIconsRed},
	{Id: "scala", font: "nerd-icons-devicon", Name: "nf-dev-scala", Color: nerdIconsRed},
	{Id: "scm", font: "nerd-icons-mdicon", Name: "nf-md-lambda", Color: nerdIconsRed},
	{Id: "swift", font: "nerd-icons-devicon", Name: "nf-dev-swift", Color: nerdIconsGreen},

	{Id: "tcl", font: "nerd-icons-mdicon", Name: "nf-md-feather", Color: nerdIconsDred},

	{Id: "tf", font: "nerd-icons-mdicon", Name: "nf-seti-terraform", Color: nerdIconsPurpleAlt},
	{Id: "tfvars", font: "nerd-icons-mdicon", Name: "nf-seti-terraform", Color: nerdIconsPurpleAlt},
	{Id: "tfstate", font: "nerd-icons-mdicon", Name: "nf-seti-terraform", Color: nerdIconsPurpleAlt},

	{Id: "asm", font: "nerd-icons-sucicon", Name: "nf-seti-asm", Color: nerdIconsBlue},
	// Verilog(-AMS) and SystemVerilog(-AMS     // Verilog(-AMS) and SystemVerilog(-AMS)
	{Id: "v", font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	{Id: "vams", font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	{Id: "sv", font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	{Id: "sva", font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	{Id: "svh", font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	{Id: "svams", font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsSilver},
	// VHDL(-AMS     // VHDL(-AMS)
	{Id: "vhd", font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsBlue},
	{Id: "vhdl", font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsBlue},
	{Id: "vhms", font: "nerd-icons-faicon", Name: "nf-fa-microchip", Color: nerdIconsBlue},
	// Cabal
	// {"cabal"       , font: "nerd-icons-devicon", name: "cabal}, TODO: cabal
	// Kotlin
	{Id: "kt", font: "nerd-icons-sucicon", Name: "nf-seti-kotlin", Color: nerdIconsOrange},
	{Id: "kts", font: "nerd-icons-sucicon", Name: "nf-seti-kotlin", Color: nerdIconsOrange},
	// Nimrod
	{Id: "nim", font: "nerd-icons-sucicon", Name: "nf-seti-nim", Color: nerdIconsYellow},
	{Id: "nims", font: "nerd-icons-sucicon", Name: "nf-seti-nim", Color: nerdIconsYellow},
	// SQL
	{Id: "sql", font: "nerd-icons-octicon", Name: "nf-oct-database", Color: nerdIconsSilver},
	// Styles
	// {"styles"      , font: "nerd-icons-devicon", name: "style}, TODO: style
	// Lua
	{Id: "lua", font: "nerd-icons-sucicon", Name: "nf-seti-lua", Color: nerdIconsDblue},
	// ASCII doc
	// {"adoc"        , font: "nerd-icons-devicon", name: "asciidoc}, TODO: asciidoc
	// {"asciidoc"    , font: "nerd-icons-devicon", name: "asciidoc}, TODO: asciidoc
	// Puppet
	{Id: "pp", font: "nerd-icons-sucicon", Name: "nf-seti-puppet", Color: nerdIconsYellow},
	// Jinja
	{Id: "j2", font: "nerd-icons-sucicon", Name: "nf-seti-jinja", Color: nerdIconsSilver},
	{Id: "jinja2", font: "nerd-icons-sucicon", Name: "nf-seti-jinja", Color: nerdIconsSilver},
	// Docker
	{Id: "dockerfile", font: "nerd-icons-sucicon", Name: "nf-seti-docker", Color: nerdIconsCyan},
	// Vagrant
	// {"vagrantfile" , font: "nerd-icons-fileicon", name: "vagrant}, TODO: vagrant
	// GLSL
	{Id: "glsl", font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsBlue},
	{Id: "vert", font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsBlue},
	{Id: "tesc", font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsPurple},
	{Id: "tese", font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsDpurple},
	{Id: "geom", font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsGreen},
	{Id: "frag", font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsRed},
	{Id: "comp", font: "nerd-icons-faicon", Name: "nf-fa-paint_brush", Color: nerdIconsDblue},
	// CUDA
	{Id: "cu", font: "nerd-icons-sucicon", Name: "nf-custom-c", Color: nerdIconsGreen},
	{Id: "cuh", font: "nerd-icons-faicon", Name: "nf-fa-h_square", Color: nerdIconsGreen},
	// Fortran
	{Id: "f90", font: "nerd-icons-mdicon", Name: "F", Color: nerdIconsPurple},
	// C#
	{Id: "cs", font: "nerd-icons-mdicon", Name: "nf-md-language_csharp", Color: nerdIconsDblue},
	{Id: "csx", font: "nerd-icons-mdicon", Name: "nf-md-language_csharp", Color: nerdIconsDblue},
	// F#
	{Id: "fs", font: "nerd-icons-devicon", Name: "nf-dev-fsharp", Color: nerdIconsBlueAlt},
	{Id: "fsi", font: "nerd-icons-devicon", Name: "nf-dev-fsharp", Color: nerdIconsBlueAlt},
	{Id: "fsx", font: "nerd-icons-devicon", Name: "nf-dev-fsharp", Color: nerdIconsBlueAlt},
	{Id: "fsscript", font: "nerd-icons-devicon", Name: "nf-dev-fsharp", Color: nerdIconsBlueAlt},
	// zig
	{Id: "zig", font: "nerd-icons-sucicon", Name: "nf-seti-zig", Color: nerdIconsOrange},
	// odin
	// {"odin"        , font: "nerd-icons-fileicon", name: "odin}, TODO: odin
	// File Types
	{Id: "ico", font: "nerd-icons-octicon", Name: "nf-oct-file_media", Color: nerdIconsBlue},
	{Id: "png", font: "nerd-icons-mdicon", Name: "nf-md-image", Color: nerdIconsOrange},
	{Id: "gif", font: "nerd-icons-mdicon", Name: "nf-md-image", Color: nerdIconsGreen},
	{Id: "jpeg", font: "nerd-icons-mdicon", Name: "nf-md-image", Color: nerdIconsDblue},
	{Id: "jpg", font: "nerd-icons-mdicon", Name: "nf-md-image", Color: nerdIconsDblue},
	{Id: "webp", font: "nerd-icons-octicon", Name: "nf-oct-file_media", Color: nerdIconsDblue},
	{Id: "xpm", font: "nerd-icons-octicon", Name: "nf-oct-file_media", Color: nerdIconsDgreen},
	// Audio
	{Id: "mp3", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "wav", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "m4a", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "ogg", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "flac", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "opus", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "au", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "aif", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "aifc", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "aiff", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsDred},
	{Id: "svg", font: "nerd-icons-sucicon", Name: "nf-seti-svg", Color: nerdIconsLgreen},
	// Video
	{Id: "mov", font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "mp4", font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "ogv", font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsDblue},
	{Id: "mpg", font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "mpeg", font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "flv", font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "ogv", font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsDblue},
	{Id: "mkv", font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	{Id: "webm", font: "nerd-icons-faicon", Name: "nf-fa-film", Color: nerdIconsBlue},
	// Fonts
	{Id: "ttf", font: "nerd-icons-faicon", Name: "nf-fa-font", Color: nerdIconsDcyan},
	{Id: "woff", font: "nerd-icons-faicon", Name: "nf-fa-font", Color: nerdIconsCyan},
	{Id: "woff2", font: "nerd-icons-faicon", Name: "nf-fa-font", Color: nerdIconsCyan},
	// Archives
	{Id: "tar", font: "nerd-icons-mdicon", Name: "nf-seti-zip", Color: nerdIconsOrange},
	{Id: "rar", font: "nerd-icons-mdicon", Name: "nf-seti-zip", Color: nerdIconsOrange},
	{Id: "tgz", font: "nerd-icons-mdicon", Name: "nf-seti-zip", Color: nerdIconsOrange},
	{Id: "jar", font: "nerd-icons-devicon", Name: "nf-dev-java", Color: nerdIconsDpurple},
	// Doc
	{Id: "pdf", font: "nerd-icons-codicon", Name: "nf-cod-file_pdf", Color: nerdIconsDred},
	{Id: "text", font: "nerd-icons-faicon", Name: "nf-fa-file_text", Color: nerdIconsCyan},
	{Id: "txt", font: "nerd-icons-faicon", Name: "nf-fa-file_text", Color: nerdIconsCyan},
	{Id: "doc", font: "nerd-icons-mdicon", Name: "nf-seti-word", Color: nerdIconsBlue},
	{Id: "docx", font: "nerd-icons-mdicon", Name: "nf-seti-word", Color: nerdIconsBlue},
	{Id: "docm", font: "nerd-icons-mdicon", Name: "nf-seti-word", Color: nerdIconsBlue},
	{Id: "texi", font: "nerd-icons-sucicon", Name: "nf-seti-tex", Color: nerdIconsLred},
	{Id: "tex", font: "nerd-icons-sucicon", Name: "nf-seti-tex", Color: nerdIconsLred},
	{Id: "ltx", font: "nerd-icons-sucicon", Name: "nf-seti-tex", Color: nerdIconsLred},
	{Id: "dtx", font: "nerd-icons-sucicon", Name: "nf-seti-tex", Color: nerdIconsLred},
	{Id: "sty", font: "nerd-icons-sucicon", Name: "nf-seti-tex", Color: nerdIconsLred},
	{Id: "md", font: "nerd-icons-octicon", Name: "nf-oct-markdown", Color: nerdIconsLblue},
	{Id: "bib", font: "nerd-icons-mdicon", Name: "nf-fa-book", Color: nerdIconsLblue},
	{Id: "org", font: "nerd-icons-sucicon", Name: "nf-custom-orgmode", Color: nerdIconsLgreen},
	{Id: "pps", font: "nerd-icons-mdicon", Name: "nf-fa-file_powerpoint_o", Color: nerdIconsOrange},
	{Id: "ppt", font: "nerd-icons-mdicon", Name: "nf-fa-file_powerpoint_o", Color: nerdIconsOrange},
	{Id: "pptsx", font: "nerd-icons-mdicon", Name: "nf-fa-file_powerpoint_o", Color: nerdIconsOrange},
	{Id: "ppttx", font: "nerd-icons-mdicon", Name: "nf-fa-file_powerpoint_o", Color: nerdIconsOrange},
	{Id: "knt", font: "nerd-icons-mdicon", Name: "nf-fa-file_powerpoint_o", Color: nerdIconsCyan},
	{Id: "xlsx", font: "nerd-icons-mdicon", Name: "nf-seti-xls", Color: nerdIconsDgreen},
	{Id: "xlsm", font: "nerd-icons-mdicon", Name: "nf-seti-xls", Color: nerdIconsDgreen},
	{Id: "xlsb", font: "nerd-icons-mdicon", Name: "nf-seti-xls", Color: nerdIconsDgreen},
	{Id: "xltx", font: "nerd-icons-mdicon", Name: "nf-seti-xls", Color: nerdIconsDgreen},
	{Id: "xltm", font: "nerd-icons-mdicon", Name: "nf-seti-xls", Color: nerdIconsDgreen},
	{Id: "epub", font: "nerd-icons-mdicon", Name: "nf-fa-book", Color: nerdIconsGreen},
	{Id: "ly", font: "nerd-icons-faicon", Name: "nf-fa-music", Color: nerdIconsGreen},
	//
	{Id: "key", font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsLblue},
	{Id: "pem", font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsOrange},
	{Id: "p12", font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsDorange},
	{Id: "crt", font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsLblue},
	{Id: "pub", font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsBlue},
	{Id: "gpg", font: "nerd-icons-octicon", Name: "nf-oct-key", Color: nerdIconsLblue},
	{Id: "cache", font: "nerd-icons-octicon", Name: "nf-oct-database", Color: nerdIconsGreen},
}

var regexes = []iconInfo2{
	{Id: "^TAGS$", font: "nerd-icons-octicon", Name: "nf-oct-tag", Color: nerdIconsBlue},
	{Id: "^TODO$", font: "nerd-icons-octicon", Name: "nf-oct-checklist", Color: nerdIconsLyellow},
	{Id: "^LICENSE$", font: "nerd-icons-octicon", Name: "nf-oct-book", Color: nerdIconsBlue},
	{Id: "^readme", font: "nerd-icons-octicon", Name: "nf-oct-book", Color: nerdIconsLcyan},
	// Config
	{Id: "nginx$", font: "nerd-icons-devicon", Name: "nf-dev-nginx", Color: nerdIconsDgreen},
	// {"apache$"              , font: "nerd-icons-alltheicon", name: "apache}, TODO: apache
	// C
	{Id: "^Makefile$", font: "nerd-icons-sucicon", Name: "nf-seti-makefile", Color: nerdIconsDorange},
	{Id: "^CMakeLists.txt$", font: "nerd-icons-sucicon", Name: "nf-seti-makefile", Color: nerdIconsRed},       // TODO: cmake
	{Id: "^CMakeCache.txt$", font: "nerd-icons-sucicon", Name: "nf-seti-makefile", Color: nerdIconsBlue},      // TODO: cmakecache
	{Id: "^meson.build$", font: "nerd-icons-sucicon", Name: "nf-seti-makefile", Color: nerdIconsPurple},       // TODO: meson
	{Id: "^meson_options.txt$", font: "nerd-icons-sucicon", Name: "nf-seti-makefile", Color: nerdIconsPurple}, // TODO: meson
	// Docker
	{Id: "^\\.?Dockerfile", font: "nerd-icons-sucicon", Name: "nf-seti-docker", Color: nerdIconsBlue},
	// Homebrew
	{Id: "^Brewfile$", font: "nerd-icons-faicon", Name: "nf-fa-beer", Color: nerdIconsLsilver},
	// // AWS
	{Id: "^stack.*.json$", font: "nerd-icons-devicon", Name: "nf-dev-aws", Color: nerdIconsOrange},
	{Id: "^serverless\\.yml$", font: "nerd-icons-faicon", Name: "nf-fa-bolt", Color: nerdIconsYellow},
	// lock files
	{Id: "~$", font: "nerd-icons-octicon", Name: "nf-oct-lock", Color: nerdIconsMaroon},
	// Source Codes
	{Id: "^mix.lock$", font: "nerd-icons-sucicon", Name: "nf-seti-elixir", Color: nerdIconsLyellow},
	// Ruby
	{Id: "^Gemfile\\(\\.lock\\)?$", font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsRed},
	{Id: "_?test\\.rb$", font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsRed},
	{Id: "_?test_helper\\.rb$", font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsDred},
	{Id: "_?spec\\.rb$", font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsRed},
	{Id: "_?spec_helper\\.rb$", font: "nerd-icons-octicon", Name: "nf-dev-ruby", Color: nerdIconsDred},

	{Id: "-?spec\\.ts$", font: "nerd-icons-mdicon", Name: "nf-md-language_typescript", Color: nerdIconsBlue},
	{Id: "-?test\\.ts$", font: "nerd-icons-mdicon", Name: "nf-md-language_typescript", Color: nerdIconsBlue},
	{Id: "-?spec\\.js$", font: "nerd-icons-mdicon", Name: "nf-md-language_javascript", Color: nerdIconsLpurple},
	{Id: "-?test\\.js$", font: "nerd-icons-mdicon", Name: "nf-md-language_javascript", Color: nerdIconsLpurple},
	{Id: "-?spec\\.jsx$", font: "nerd-icons-mdicon", Name: "nf-md-react", Color: nerdIconsBlueAlt},
	{Id: "-?test\\.jsx$", font: "nerd-icons-mdicon", Name: "nf-md-react", Color: nerdIconsBlueAlt},
	// Git
	{Id: "^MERGE_", font: "nerd-icons-octicon", Name: "nf-oct-git_merge", Color: nerdIconsRed},
	{Id: "^COMMIT_EDITMSG", font: "nerd-icons-octicon", Name: "nf-oct-git_commit", Color: nerdIconsRed},
	// Stylesheeting
	{Id: "stylelint", font: "nerd-icons-sucicon", Name: "nf-seti-stylelint", Color: nerdIconsLyellow},
	// JavaScript
	{Id: "^package.json$", font: "nerd-icons-devicon", Name: "nf-dev-npm", Color: nerdIconsRed},
	{Id: "^package.lock.json$", font: "nerd-icons-devicon", Name: "nf-dev-npm", Color: nerdIconsDred},
	{Id: "^yarn\\.lock", font: "nerd-icons-sucicon", Name: "nf-seti-yarn", Color: nerdIconsBlueAlt},
	{Id: "\\.npmignore$", font: "nerd-icons-devicon", Name: "nf-dev-npm", Color: nerdIconsDred},
	{Id: "^bower.json$", font: "nerd-icons-devicon", Name: "nf-dev-bower", Color: nerdIconsLorange},
	{Id: "^gulpfile", font: "nerd-icons-devicon", Name: "nf-dev-gulp", Color: nerdIconsLred},
	{Id: "^gruntfile", font: "nerd-icons-devicon", Name: "nf-dev-grunt", Color: nerdIconsLyellow},
	{Id: "^webpack", font: "nerd-icons-mdicon", Name: "nf-md-webpack", Color: nerdIconsLblue},
	// Go
	{Id: "^go.mod$", font: "nerd-icons-sucicon", Name: "nf-seti-config", Color: nerdIconsBlueAlt},
	{Id: "^go.work$", font: "nerd-icons-sucicon", Name: "nf-seti-config", Color: nerdIconsBlueAlt},
	// Emacs
	{Id: "bookmark", font: "nerd-icons-octicon", Name: "nf-oct-bookmark", Color: nerdIconsLpink},
	{Id: "^\\*scratch\\*$", font: "nerd-icons-faicon", Name: "nf-fa-sticky_note", Color: nerdIconsLyellow},
	{Id: "^\\*scratch.*", font: "nerd-icons-faicon", Name: "nf-fa-sticky_note", Color: nerdIconsYellow},
	{Id: "^\\*new-tab\\*$", font: "nerd-icons-mdicon", Name: "nf-md-star", Color: nerdIconsCyan},
	{Id: "^\\.", font: "nerd-icons-octicon", Name: "nf-oct-gear"},
}

var defFolderIcons = []iconInfo2{
	{Id: "trash", font: "nerd-icons-faicon", Name: "nf-fa-trash_o"},
	{Id: "dropbox", font: "nerd-icons-faicon", Name: "nf-fa-dropbox"},
	{Id: "google[ _-]drive", font: "nerd-icons-mdicon", Name: "nf-md-folder_google_drive"},
	{Id: "github", font: "nerd-icons-sucicon", Name: "nf-custom-folder_github"},
	{Id: "^atom$", font: "nerd-icons-devicon", Name: "nf-dev-atom"},
	{Id: "documents", font: "nerd-icons-mdicon", Name: "nf-md-folder_file"},
	{Id: "download", font: "nerd-icons-mdicon", Name: "nf-md-folder_download"},
	{Id: "desktop", font: "nerd-icons-octicon", Name: "nf-oct-device_desktop"},
	{Id: "pictures", font: "nerd-icons-mdicon", Name: "nf-md-folder_image"},
	{Id: "photos", font: "nerd-icons-faicon", Name: "nf-fa-camera_retro"},
	{Id: "music", font: "nerd-icons-mdicon", Name: "nf-md-folder_music"},
	{Id: "movies", font: "nerd-icons-faicon", Name: "nf-fa-film"},
	{Id: "code", font: "nerd-icons-octicon", Name: "nf-oct-code"},
	{Id: "workspace", font: "nerd-icons-octicon", Name: "nf-oct-code"},
	// {"test"             , font: "nerd-icons-devicon", name: "test-dir"},
	{Id: "\\.git", font: "nerd-icons-sucicon", Name: "nf-custom-folder_git"},
	{Id: "\\.config", font: "nerd-icons-sucicon", Name: "nf-custom-folder_config"},
	{Id: ".?", font: "nerd-icons-sucicon", Name: "nf-custom-folder_oct"},
}
