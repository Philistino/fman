// MIT License

// Copyright (c) 2020 Yash Handa

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// assets contains all the Icon glyphs info
package icons

var iconSet = map[string]Icon{
	"html":           {glyph: "\uf13b", name: "nf-fa-html5", rgb: [3]uint8{228, 79, 57}},             // html
	"markdown":       {glyph: "\uf853", rgb: [3]uint8{66, 165, 245}},                                 // markdown
	"css":            {glyph: "\uf81b", rgb: [3]uint8{66, 165, 245}},                                 // css
	"css-map":        {glyph: "\ue749", rgb: [3]uint8{66, 165, 245}},                                 // css-map
	"sass":           {glyph: "\ue603", name: "nf-seti-sass", rgb: [3]uint8{237, 80, 122}},           // sass
	"less":           {glyph: "\ue60b", name: "nf-seti-json", rgb: [3]uint8{2, 119, 189}},            // less
	"json":           {glyph: "\ue60b", name: "nf-seti-json", rgb: [3]uint8{251, 193, 60}},           // json
	"yaml":           {glyph: "\ue6a8", name: "nf-seti-yml", rgb: [3]uint8{160, 116, 196}},           // yaml
	"xml":            {glyph: "\uf72d", rgb: [3]uint8{64, 153, 69}},                                  // xml
	"image":          {glyph: "\uf71e", rgb: [3]uint8{48, 166, 154}},                                 // image
	"javascript":     {glyph: "\ue74e", rgb: [3]uint8{255, 202, 61}},                                 // javascript
	"javascript-map": {glyph: "\ue781", rgb: [3]uint8{255, 202, 61}},                                 // javascript-map
	"test-jsx":       {glyph: "\uf595", rgb: [3]uint8{35, 188, 212}},                                 // test-jsx
	"test-js":        {glyph: "\uf595", rgb: [3]uint8{255, 202, 61}},                                 // test-js
	"react":          {glyph: "\ue7ba", rgb: [3]uint8{35, 188, 212}},                                 // react
	"react_ts":       {glyph: "\ue7ba", rgb: [3]uint8{36, 142, 211}},                                 // react_ts
	"settings":       {glyph: "\uf013", rgb: [3]uint8{66, 165, 245}},                                 // settings
	"typescript":     {glyph: "\ue628", rgb: [3]uint8{3, 136, 209}},                                  // typescript
	"typescript-def": {glyph: "\ufbe4", rgb: [3]uint8{3, 136, 209}},                                  // typescript-def
	"test-ts":        {glyph: "\uf595", rgb: [3]uint8{3, 136, 209}},                                  // test-ts
	"pdf":            {glyph: "\uf724", rgb: [3]uint8{244, 68, 62}},                                  // pdf
	"table":          {glyph: "\uf71a", rgb: [3]uint8{139, 195, 74}},                                 // table
	"visualstudio":   {glyph: "\ue70c", rgb: [3]uint8{173, 99, 188}},                                 // visualstudio
	"database":       {glyph: "\ue706", rgb: [3]uint8{255, 202, 61}},                                 // database
	"mysql":          {glyph: "\ue704", rgb: [3]uint8{1, 94, 134}},                                   // mysql
	"postgresql":     {glyph: "\ue76e", rgb: [3]uint8{49, 99, 140}},                                  // postgresql
	"sqlite":         {glyph: "\ue7c4", rgb: [3]uint8{1, 57, 84}},                                    // sqlite
	"csharp":         {glyph: "\uf81a", rgb: [3]uint8{2, 119, 189}},                                  // csharp
	"zip":            {glyph: "\uf410", rgb: [3]uint8{175, 180, 43}},                                 // zip
	"exe":            {glyph: "\uf2d0", rgb: [3]uint8{229, 77, 58}},                                  // exe
	"java":           {glyph: "\uf675", rgb: [3]uint8{244, 68, 62}},                                  // java
	"c":              {glyph: "\ufb70", rgb: [3]uint8{2, 119, 189}},                                  // c
	"cpp":            {glyph: "\ue646", name: "nf-seti-cpp", rgb: [3]uint8{2, 119, 189}},             // cpp
	"go":             {glyph: "\ue627", name: "nf-seti-go", rgb: [3]uint8{32, 173, 194}},             // go
	"go-mod":         {glyph: "\ue627", name: "nf-seti-go", rgb: [3]uint8{237, 80, 122}},             // go-mod
	"go-test":        {glyph: "\ue627", name: "nf-seti-go", rgb: [3]uint8{255, 213, 79}},             // go-test
	"python":         {glyph: "\ue606", name: "nf-seti-python", rgb: [3]uint8{52, 102, 143}},         // python
	"python-misc":    {glyph: "\uf820", rgb: [3]uint8{130, 61, 28}},                                  // python-misc
	"url":            {glyph: "\uf836", rgb: [3]uint8{66, 165, 245}},                                 // url
	"console":        {glyph: "\uf68c", rgb: [3]uint8{250, 111, 66}},                                 // console
	"word":           {glyph: "\ue6a5", name: "nf-seti-word", rgb: [3]uint8{1, 87, 155}},             // word
	"certificate":    {glyph: "\uf623", rgb: [3]uint8{249, 89, 63}},                                  // certificate
	"key":            {glyph: "\uf805", rgb: [3]uint8{48, 166, 154}},                                 // key
	"font":           {glyph: "\uf031", rgb: [3]uint8{244, 68, 62}},                                  // font
	"lib":            {glyph: "\uf831", rgb: [3]uint8{139, 195, 74}},                                 // lib
	"ruby":           {glyph: "\ue739", rgb: [3]uint8{229, 61, 58}},                                  // ruby
	"gemfile":        {glyph: "\ue21e", rgb: [3]uint8{229, 61, 58}},                                  // gemfile
	"fsharp":         {glyph: "\ue7a7", rgb: [3]uint8{55, 139, 186}},                                 // fsharp
	"swift":          {glyph: "\ufbe3", rgb: [3]uint8{249, 95, 63}},                                  // swift
	"docker":         {glyph: "\uf308", rgb: [3]uint8{1, 135, 201}},                                  // docker
	"powerpoint":     {glyph: "\uf1c4", name: "nf-fa-file_powerpoint_o", rgb: [3]uint8{209, 71, 51}}, // powerpoint
	"video":          {glyph: "\ue69f", name: "nf-seti-video", rgb: [3]uint8{253, 154, 62}},          // video
	"virtual":        {glyph: "\uf822", rgb: [3]uint8{3, 155, 229}},                                  // virtual
	"email":          {glyph: "\uf6ed", rgb: [3]uint8{66, 165, 245}},                                 // email
	"audio":          {glyph: "\ufb75", rgb: [3]uint8{239, 83, 80}},                                  // audio
	"coffee":         {glyph: "\uf675", rgb: [3]uint8{66, 165, 245}},                                 // coffee
	"document":       {glyph: "\uf718", rgb: [3]uint8{66, 165, 245}},                                 // document
	"rust":           {glyph: "\ue7a8", rgb: [3]uint8{250, 111, 66}},                                 // rust
	"raml":           {glyph: "\ue60b", rgb: [3]uint8{66, 165, 245}},                                 // raml
	"xaml":           {glyph: "\ufb72", rgb: [3]uint8{66, 165, 245}},                                 // xaml
	"haskell":        {glyph: "\ue61f", rgb: [3]uint8{254, 168, 62}},                                 // haskell
	"git":            {glyph: "\ue65d", rgb: [3]uint8{229, 77, 58}},                                  // git
	"lua":            {glyph: "\ue620", rgb: [3]uint8{66, 165, 245}},                                 // lua
	"clojure":        {glyph: "\ue76a", rgb: [3]uint8{100, 221, 23}},                                 // clojure
	"groovy":         {glyph: "\uf2a6", rgb: [3]uint8{41, 198, 218}},                                 // groovy
	"r":              {glyph: "\ufcd2", rgb: [3]uint8{25, 118, 210}},                                 // r
	"dart":           {glyph: "\ue798", rgb: [3]uint8{87, 182, 240}},                                 // dart
	"mxml":           {glyph: "\uf72d", rgb: [3]uint8{254, 168, 62}},                                 // mxml
	"assembly":       {glyph: "\uf471", rgb: [3]uint8{250, 109, 63}},                                 // assembly
	"vue":            {glyph: "\ufd42", rgb: [3]uint8{65, 184, 131}},                                 // vue
	"vue-config":     {glyph: "\ufd42", rgb: [3]uint8{58, 121, 110}},                                 // vue-config
	"lock":           {glyph: "\uf83d", rgb: [3]uint8{255, 213, 79}},                                 // lock
	"handlebars":     {glyph: "\ue60f", rgb: [3]uint8{250, 111, 66}},                                 // handlebars
	"perl":           {glyph: "\ue769", rgb: [3]uint8{149, 117, 205}},                                // perl
	"elixir":         {glyph: "\ue62d", rgb: [3]uint8{149, 117, 205}},                                // elixir
	"erlang":         {glyph: "\ue7b1", rgb: [3]uint8{244, 68, 62}},                                  // erlang
	"twig":           {glyph: "\ue61c", rgb: [3]uint8{155, 185, 47}},                                 // twig
	"julia":          {glyph: "\ue624", rgb: [3]uint8{134, 82, 159}},                                 // julia
	"elm":            {glyph: "\ue62c", rgb: [3]uint8{96, 181, 204}},                                 // elm
	"smarty":         {glyph: "\uf834", rgb: [3]uint8{255, 207, 60}},                                 // smarty
	"stylus":         {glyph: "\ue600", rgb: [3]uint8{192, 202, 51}},                                 // stylus
	"verilog":        {glyph: "\ufb19", rgb: [3]uint8{250, 111, 66}},                                 // verilog
	"robot":          {glyph: "\ufba7", rgb: [3]uint8{249, 89, 63}},                                  // robot
	"solidity":       {glyph: "\ufcb9", rgb: [3]uint8{3, 136, 209}},                                  // solidity
	"yang":           {glyph: "\ufb7e", rgb: [3]uint8{66, 165, 245}},                                 // yang
	"vercel":         {glyph: "\uf47e", rgb: [3]uint8{207, 216, 220}},                                // vercel
	"applescript":    {glyph: "\uf302", rgb: [3]uint8{120, 144, 156}},                                // applescript
	"cake":           {glyph: "\uf5ea", rgb: [3]uint8{250, 111, 66}},                                 // cake
	"nim":            {glyph: "\uf6a4", rgb: [3]uint8{255, 202, 61}},                                 // nim
	"todo":           {glyph: "\uf058", rgb: [3]uint8{124, 179, 66}},                                 // todo
	"nix":            {glyph: "\uf313", rgb: [3]uint8{80, 117, 193}},                                 // nix
	"http":           {glyph: "\uf484", rgb: [3]uint8{66, 165, 245}},                                 // http
	"webpack":        {glyph: "\ufc29", rgb: [3]uint8{142, 214, 251}},                                // webpack
	"ionic":          {glyph: "\ue7a9", rgb: [3]uint8{79, 143, 247}},                                 // ionic
	"gulp":           {glyph: "\ue763", rgb: [3]uint8{229, 61, 58}},                                  // gulp
	"nodejs":         {glyph: "\uf898", rgb: [3]uint8{139, 195, 74}},                                 // nodejs
	"npm":            {glyph: "\ue71e", rgb: [3]uint8{203, 56, 55}},                                  // npm
	"yarn":           {glyph: "\uf61a", rgb: [3]uint8{44, 142, 187}},                                 // yarn
	"android":        {glyph: "\uf531", rgb: [3]uint8{139, 195, 74}},                                 // android
	"tune":           {glyph: "\ufb69", rgb: [3]uint8{251, 193, 60}},                                 // tune
	"contributing":   {glyph: "\uf64d", rgb: [3]uint8{255, 202, 61}},                                 // contributing
	// "readme":           {glyph: "\uf7fb", rgb: [3]uint8{66, 165, 245}},                    // readme
	"readme":           {glyph: "\ue66a", name: "nf-seti-info", rgb: [3]uint8{66, 165, 245}}, // readme
	"changelog":        {glyph: "\ufba6", rgb: [3]uint8{139, 195, 74}},                       // changelog
	"credits":          {glyph: "\uf75f", rgb: [3]uint8{156, 204, 101}},                      // credits
	"authors":          {glyph: "\uf0c0", rgb: [3]uint8{244, 68, 62}},                        // authors
	"favicon":          {glyph: "\ue623", rgb: [3]uint8{255, 213, 79}},                       // favicon
	"karma":            {glyph: "\ue622", rgb: [3]uint8{60, 190, 174}},                       // karma
	"travis":           {glyph: "\ue77e", rgb: [3]uint8{203, 58, 73}},                        // travis
	"heroku":           {glyph: "\ue607", rgb: [3]uint8{105, 99, 185}},                       // heroku
	"gitlab":           {glyph: "\uf296", rgb: [3]uint8{226, 69, 57}},                        // gitlab
	"bower":            {glyph: "\ue61a", rgb: [3]uint8{239, 88, 60}},                        // bower
	"conduct":          {glyph: "\uf64b", rgb: [3]uint8{205, 220, 57}},                       // conduct
	"jenkins":          {glyph: "\ue767", rgb: [3]uint8{240, 214, 183}},                      // jenkins
	"code-climate":     {glyph: "\uf7f4", rgb: [3]uint8{238, 238, 238}},                      // code-climate
	"log":              {glyph: "\uf719", rgb: [3]uint8{175, 180, 43}},                       // log
	"ejs":              {glyph: "\ue618", rgb: [3]uint8{255, 202, 61}},                       // ejs
	"grunt":            {glyph: "\ue611", rgb: [3]uint8{251, 170, 61}},                       // grunt
	"django":           {glyph: "\ue71d", rgb: [3]uint8{67, 160, 71}},                        // django
	"makefile":         {glyph: "\uf728", rgb: [3]uint8{239, 83, 80}},                        // makefile
	"bitbucket":        {glyph: "\uf171", rgb: [3]uint8{31, 136, 229}},                       // bitbucket
	"d":                {glyph: "\ue7af", rgb: [3]uint8{244, 68, 62}},                        // d
	"mdx":              {glyph: "\uf853", rgb: [3]uint8{255, 202, 61}},                       // mdx
	"azure-pipelines":  {glyph: "\uf427", rgb: [3]uint8{20, 101, 192}},                       // azure-pipelines
	"azure":            {glyph: "\ufd03", rgb: [3]uint8{31, 136, 229}},                       // azure
	"razor":            {glyph: "\uf564", rgb: [3]uint8{66, 165, 245}},                       // razor
	"asciidoc":         {glyph: "\uf718", rgb: [3]uint8{244, 68, 62}},                        // asciidoc
	"edge":             {glyph: "\uf564", rgb: [3]uint8{239, 111, 60}},                       // edge
	"scheme":           {glyph: "\ufb26", rgb: [3]uint8{244, 68, 62}},                        // scheme
	"3d":               {glyph: "\ue79b", rgb: [3]uint8{40, 182, 246}},                       // 3d
	"svg":              {glyph: "\ue698", name: "nf-seti-svg", rgb: [3]uint8{255, 181, 62}},  // svg
	"vim":              {glyph: "\ue62b", rgb: [3]uint8{67, 160, 71}},                        // vim
	"moonscript":       {glyph: "\uf186", rgb: [3]uint8{251, 193, 60}},                       // moonscript
	"codeowners":       {glyph: "\uf507", rgb: [3]uint8{175, 180, 43}},                       // codeowners
	"disc":             {glyph: "\ue271", rgb: [3]uint8{176, 190, 197}},                      // disc
	"fortran":          {glyph: "F", rgb: [3]uint8{250, 111, 66}},                            // fortran
	"tcl":              {glyph: "\ufbd1", rgb: [3]uint8{239, 83, 80}},                        // tcl
	"liquid":           {glyph: "\ue275", rgb: [3]uint8{40, 182, 246}},                       // liquid
	"prolog":           {glyph: "\ue7a1", rgb: [3]uint8{239, 83, 80}},                        // prolog
	"husky":            {glyph: "\uf8e8", rgb: [3]uint8{229, 229, 229}},                      // husky
	"coconut":          {glyph: "\uf5d2", rgb: [3]uint8{141, 110, 99}},                       // coconut
	"sketch":           {glyph: "\uf6c7", rgb: [3]uint8{255, 194, 61}},                       // sketch
	"pawn":             {glyph: "\ue261", rgb: [3]uint8{239, 111, 60}},                       // pawn
	"commitlint":       {glyph: "\ufc16", rgb: [3]uint8{43, 150, 137}},                       // commitlint
	"dhall":            {glyph: "\uf448", rgb: [3]uint8{120, 144, 156}},                      // dhall
	"dune":             {glyph: "\uf7f4", rgb: [3]uint8{244, 127, 61}},                       // dune
	"shaderlab":        {glyph: "\ufbad", rgb: [3]uint8{25, 118, 210}},                       // shaderlab
	"command":          {glyph: "\ufb32", rgb: [3]uint8{175, 188, 194}},                      // command
	"stryker":          {glyph: "\uf05b", rgb: [3]uint8{239, 83, 80}},                        // stryker
	"modernizr":        {glyph: "\ue720", rgb: [3]uint8{234, 72, 99}},                        // modernizr
	"roadmap":          {glyph: "\ufb6d", rgb: [3]uint8{48, 166, 154}},                       // roadmap
	"debian":           {glyph: "\uf306", rgb: [3]uint8{211, 61, 76}},                        // debian
	"ubuntu":           {glyph: "\uf31c", rgb: [3]uint8{214, 73, 53}},                        // ubuntu
	"arch":             {glyph: "\uf303", rgb: [3]uint8{33, 142, 202}},                       // arch
	"redhat":           {glyph: "\uf316", rgb: [3]uint8{231, 61, 58}},                        // redhat
	"gentoo":           {glyph: "\uf30d", rgb: [3]uint8{148, 141, 211}},                      // gentoo
	"linux":            {glyph: "\ue712", rgb: [3]uint8{238, 207, 55}},                       // linux
	"raspberry-pi":     {glyph: "\uf315", rgb: [3]uint8{208, 60, 76}},                        // raspberry-pi
	"manjaro":          {glyph: "\uf312", rgb: [3]uint8{73, 185, 90}},                        // manjaro
	"opensuse":         {glyph: "\uf314", rgb: [3]uint8{111, 180, 36}},                       // opensuse
	"fedora":           {glyph: "\uf30a", rgb: [3]uint8{52, 103, 172}},                       // fedora
	"freebsd":          {glyph: "\uf30c", rgb: [3]uint8{175, 44, 42}},                        // freebsd
	"centOS":           {glyph: "\uf304", rgb: [3]uint8{157, 83, 135}},                       // centOS
	"alpine":           {glyph: "\uf300", rgb: [3]uint8{14, 87, 123}},                        // alpine
	"mint":             {glyph: "\uf30f", rgb: [3]uint8{125, 190, 58}},                       // mint
	"routing":          {glyph: "\ufb40", rgb: [3]uint8{67, 160, 71}},                        // routing
	"laravel":          {glyph: "\ue73f", rgb: [3]uint8{248, 80, 81}},                        // laravel
	"pug":              {glyph: "\ue60e", rgb: [3]uint8{239, 204, 163}},                      // pug (Not supported by nerdFont)
	"blink":            {glyph: "\uf72a", rgb: [3]uint8{249, 169, 60}},                       // blink (The Foundry Nuke) (Not supported by nerdFont)
	"postcss":          {glyph: "\uf81b", rgb: [3]uint8{244, 68, 62}},                        // postcss (Not supported by nerdFont)
	"jinja":            {glyph: "\ue000", rgb: [3]uint8{174, 44, 42}},                        // jinja (Not supported by nerdFont)
	"sublime":          {glyph: "\ue7aa", rgb: [3]uint8{239, 148, 58}},                       // sublime (Not supported by nerdFont)
	"markojs":          {glyph: "\uf13b", rgb: [3]uint8{2, 119, 189}},                        // markojs (Not supported by nerdFont)
	"vscode":           {glyph: "\ue70c", rgb: [3]uint8{33, 150, 243}},                       // vscode (Not supported by nerdFont)
	"qsharp":           {glyph: "\uf292", rgb: [3]uint8{251, 193, 60}},                       // qsharp (Not supported by nerdFont)
	"vala":             {glyph: "\uf7ab", rgb: [3]uint8{149, 117, 205}},                      // vala (Not supported by nerdFont)
	"zig":              {glyph: "Z", rgb: [3]uint8{249, 169, 60}},                            // zig (Not supported by nerdFont)
	"h":                {glyph: "h", rgb: [3]uint8{2, 119, 189}},                             // h (Not supported by nerdFont)
	"hpp":              {glyph: "h", rgb: [3]uint8{2, 119, 189}},                             // hpp (Not supported by nerdFont)
	"powershell":       {glyph: "\ufcb5", rgb: [3]uint8{5, 169, 244}},                        // powershell (Not supported by nerdFont)
	"gradle":           {glyph: "\ufcc4", rgb: [3]uint8{29, 151, 167}},                       // gradle (Not supported by nerdFont)
	"arduino":          {glyph: "\ue255", rgb: [3]uint8{35, 151, 156}},                       // arduino (Not supported by nerdFont)
	"tex":              {glyph: "\uf783", rgb: [3]uint8{66, 165, 245}},                       // tex (Not supported by nerdFont)
	"graphql":          {glyph: "\ue284", rgb: [3]uint8{237, 80, 122}},                       // graphql (Not supported by nerdFont)
	"kotlin":           {glyph: "\ue70e", rgb: [3]uint8{139, 195, 74}},                       // kotlin (Not supported by nerdFont)
	"actionscript":     {glyph: "\ufb25", rgb: [3]uint8{244, 68, 62}},                        // actionscript (Not supported by nerdFont)
	"autohotkey":       {glyph: "\uf812", rgb: [3]uint8{76, 175, 80}},                        // autohotkey (Not supported by nerdFont)
	"flash":            {glyph: "\uf740", rgb: [3]uint8{198, 52, 54}},                        // flash (Not supported by nerdFont)
	"swc":              {glyph: "\ufbd3", rgb: [3]uint8{198, 52, 54}},                        // swc (Not supported by nerdFont)
	"cmake":            {glyph: "\uf425", rgb: [3]uint8{178, 178, 179}},                      // cmake (Not supported by nerdFont)
	"nuxt":             {glyph: "\ue2a6", rgb: [3]uint8{65, 184, 131}},                       // nuxt (Not supported by nerdFont)
	"ocaml":            {glyph: "\uf1ce", rgb: [3]uint8{253, 154, 62}},                       // ocaml (Not supported by nerdFont)
	"haxe":             {glyph: "\uf425", rgb: [3]uint8{246, 137, 61}},                       // haxe (Not supported by nerdFont)
	"puppet":           {glyph: "\uf595", rgb: [3]uint8{251, 193, 60}},                       // puppet (Not supported by nerdFont)
	"purescript":       {glyph: "\uf670", rgb: [3]uint8{66, 165, 245}},                       // purescript (Not supported by nerdFont)
	"merlin":           {glyph: "\uf136", rgb: [3]uint8{66, 165, 245}},                       // merlin (Not supported by nerdFont)
	"mjml":             {glyph: "\ue714", rgb: [3]uint8{249, 89, 63}},                        // mjml (Not supported by nerdFont)
	"terraform":        {glyph: "\ue20f", rgb: [3]uint8{92, 107, 192}},                       // terraform (Not supported by nerdFont)
	"apiblueprint":     {glyph: "\uf031", rgb: [3]uint8{66, 165, 245}},                       // apiblueprint (Not supported by nerdFont)
	"slim":             {glyph: "\uf24e", rgb: [3]uint8{245, 129, 61}},                       // slim (Not supported by nerdFont)
	"babel":            {glyph: "\uf5a0", rgb: [3]uint8{253, 217, 59}},                       // babel (Not supported by nerdFont)
	"codecov":          {glyph: "\ue37c", rgb: [3]uint8{237, 80, 122}},                       // codecov (Not supported by nerdFont)
	"protractor":       {glyph: "\uf288", rgb: [3]uint8{229, 61, 58}},                        // protractor (Not supported by nerdFont)
	"eslint":           {glyph: "\ufbf6", rgb: [3]uint8{121, 134, 203}},                      // eslint (Not supported by nerdFont)
	"mocha":            {glyph: "\uf6a9", rgb: [3]uint8{161, 136, 127}},                      // mocha (Not supported by nerdFont)
	"firebase":         {glyph: "\ue787", rgb: [3]uint8{251, 193, 60}},                       // firebase (Not supported by nerdFont)
	"stylelint":        {glyph: "\ufb76", rgb: [3]uint8{207, 216, 220}},                      // stylelint (Not supported by nerdFont)
	"prettier":         {glyph: "\uf8e2", rgb: [3]uint8{86, 179, 180}},                       // prettier (Not supported by nerdFont)
	"jest":             {glyph: "J", rgb: [3]uint8{244, 85, 62}},                             // jest (Not supported by nerdFont)
	"storybook":        {glyph: "\ufd2c", rgb: [3]uint8{237, 80, 122}},                       // storybook (Not supported by nerdFont)
	"fastlane":         {glyph: "\ufbff", rgb: [3]uint8{149, 119, 232}},                      // fastlane (Not supported by nerdFont)
	"helm":             {glyph: "\ufd31", rgb: [3]uint8{32, 173, 194}},                       // helm (Not supported by nerdFont)
	"i18n":             {glyph: "\uf7be", rgb: [3]uint8{121, 134, 203}},                      // i18n (Not supported by nerdFont)
	"semantic-release": {glyph: "\uf70f", rgb: [3]uint8{245, 245, 245}},                      // semantic-release (Not supported by nerdFont)
	"godot":            {glyph: "\ufba7", rgb: [3]uint8{79, 195, 247}},                       // godot (Not supported by nerdFont)
	"godot-assets":     {glyph: "\ufba7", rgb: [3]uint8{129, 199, 132}},                      // godot-assets (Not supported by nerdFont)
	"vagrant":          {glyph: "\uf27d", rgb: [3]uint8{20, 101, 192}},                       // vagrant (Not supported by nerdFont)
	"tailwindcss":      {glyph: "\ufc8b", rgb: [3]uint8{77, 182, 172}},                       // tailwindcss (Not supported by nerdFont)
	"gcp":              {glyph: "\uf662", rgb: [3]uint8{70, 136, 250}},                       // gcp (Not supported by nerdFont)
	"opam":             {glyph: "\uf1ce", rgb: [3]uint8{255, 213, 79}},                       // opam (Not supported by nerdFont)
	"pascal":           {glyph: "\uf8da", rgb: [3]uint8{3, 136, 209}},                        // pascal (Not supported by nerdFont)
	"nuget":            {glyph: "\ue77f", rgb: [3]uint8{3, 136, 209}},                        // nuget (Not supported by nerdFont)
	"denizenscript":    {glyph: "D", rgb: [3]uint8{255, 213, 79}},                            // denizenscript (Not supported by nerdFont)
	"tags":             {glyph: "\uf412", rgb: [3]uint8{106, 159, 181}},                      // tags TODO: Standardize colors
	// "riot":             {i:"\u", c:[3]uint8{255, 255, 255}},       // riot
	// "autoit":           {i:"\u", c:[3]uint8{255, 255, 255}},       // autoit
	// "livescript":       {i:"\u", c:[3]uint8{255, 255, 255}},       // livescript
	// "reason":           {i:"\u", c:[3]uint8{255, 255, 255}},       // reason
	// "bucklescript":     {i:"\u", c:[3]uint8{255, 255, 255}},       // bucklescript
	// "mathematica":      {i:"\u", c:[3]uint8{255, 255, 255}},       // mathematica
	// "wolframlanguage":  {i:"\u", c:[3]uint8{255, 255, 255}},       // wolframlanguage
	// "nunjucks":         {i:"\u", c:[3]uint8{255, 255, 255}},       // nunjucks
	// "haml":             {i:"\u", c:[3]uint8{255, 255, 255}},       // haml
	// "cucumber":         {i:"\u", c:[3]uint8{255, 255, 255}},       // cucumber
	// "vfl":              {i:"\u", c:[3]uint8{255, 255, 255}},       // vfl
	// "kl":               {i:"\u", c:[3]uint8{255, 255, 255}},       // kl
	// "coldfusion":       {i:"\u", c:[3]uint8{255, 255, 255}},       // coldfusion
	// "cabal":            {i:"\u", c:[3]uint8{255, 255, 255}},       // cabal
	// "restql":           {i:"\u", c:[3]uint8{255, 255, 255}},       // restql
	// "kivy":             {i:"\u", c:[3]uint8{255, 255, 255}},       // kivy
	// "graphcool":        {i:"\u", c:[3]uint8{255, 255, 255}},       // graphcool
	// "sbt":              {i:"\u", c:[3]uint8{255, 255, 255}},       // sbt
	// "flow":             {i:"\u", c:[3]uint8{255, 255, 255}},       // flow
	// "bithound":         {i:"\u", c:[3]uint8{255, 255, 255}},       // bithound
	// "appveyor":         {i:"\u", c:[3]uint8{255, 255, 255}},       // appveyor
	// "fusebox":          {i:"\u", c:[3]uint8{255, 255, 255}},       // fusebox
	// "editorconfig":     {i:"\u", c:[3]uint8{255, 255, 255}},       // editorconfig
	// "watchman":         {i:"\u", c:[3]uint8{255, 255, 255}},       // watchman
	// "aurelia":          {i:"\u", c:[3]uint8{255, 255, 255}},       // aurelia
	// "rollup":           {i:"\u", c:[3]uint8{255, 255, 255}},       // rollup
	// "hack":             {i:"\u", c:[3]uint8{255, 255, 255}},       // hack
	// "apollo":           {i:"\u", c:[3]uint8{255, 255, 255}},       // apollo
	// "nodemon":          {i:"\u", c:[3]uint8{255, 255, 255}},       // nodemon
	// "webhint":          {i:"\u", c:[3]uint8{255, 255, 255}},       // webhint
	// "browserlist":      {i:"\u", c:[3]uint8{255, 255, 255}},       // browserlist
	// "crystal":          {i:"\u", c:[3]uint8{255, 255, 255}},       // crystal
	// "snyk":             {i:"\u", c:[3]uint8{255, 255, 255}},       // snyk
	// "drone":            {i:"\u", c:[3]uint8{255, 255, 255}},       // drone
	// "cuda":             {i:"\u", c:[3]uint8{255, 255, 255}},       // cuda
	// "dotjs":            {i:"\u", c:[3]uint8{255, 255, 255}},       // dotjs
	// "sequelize":        {i:"\u", c:[3]uint8{255, 255, 255}},       // sequelize
	// "gatsby":           {i:"\u", c:[3]uint8{255, 255, 255}},       // gatsby
	// "wakatime":         {i:"\u", c:[3]uint8{255, 255, 255}},       // wakatime
	// "circleci":         {i:"\u", c:[3]uint8{255, 255, 255}},       // circleci
	// "cloudfoundry":     {i:"\u", c:[3]uint8{255, 255, 255}},       // cloudfoundry
	// "processing":       {i:"\u", c:[3]uint8{255, 255, 255}},       // processing
	// "wepy":             {i:"\u", c:[3]uint8{255, 255, 255}},       // wepy
	// "hcl":              {i:"\u", c:[3]uint8{255, 255, 255}},       // hcl
	// "san":              {i:"\u", c:[3]uint8{255, 255, 255}},       // san
	// "wallaby":          {i:"\u", c:[3]uint8{255, 255, 255}},       // wallaby
	// "stencil":          {i:"\u", c:[3]uint8{255, 255, 255}},       // stencil
	// "red":              {i:"\u", c:[3]uint8{255, 255, 255}},       // red
	// "webassembly":      {i:"\u", c:[3]uint8{255, 255, 255}},       // webassembly
	// "foxpro":           {i:"\u", c:[3]uint8{255, 255, 255}},       // foxpro
	// "jupyter":          {i:"\u", c:[3]uint8{255, 255, 255}},       // jupyter
	// "ballerina":        {i:"\u", c:[3]uint8{255, 255, 255}},       // ballerina
	// "racket":           {i:"\u", c:[3]uint8{255, 255, 255}},       // racket
	// "bazel":            {i:"\u", c:[3]uint8{255, 255, 255}},       // bazel
	// "mint":             {i:"\u", c:[3]uint8{255, 255, 255}},       // mint
	// "velocity":         {i:"\u", c:[3]uint8{255, 255, 255}},       // velocity
	// "prisma":           {i:"\u", c:[3]uint8{255, 255, 255}},       // prisma
	// "abc":              {i:"\u", c:[3]uint8{255, 255, 255}},       // abc
	// "istanbul":         {i:"\u", c:[3]uint8{255, 255, 255}},       // istanbul
	// "lisp":             {i:"\u", c:[3]uint8{255, 255, 255}},       // lisp
	// "buildkite":        {i:"\u", c:[3]uint8{255, 255, 255}},       // buildkite
	// "netlify":          {i:"\u", c:[3]uint8{255, 255, 255}},       // netlify
	// "svelte":           {i:"\u", c:[3]uint8{255, 255, 255}},       // svelte
	// "nest":             {i:"\u", c:[3]uint8{255, 255, 255}},       // nest
	// "percy":            {i:"\u", c:[3]uint8{255, 255, 255}},       // percy
	// "gitpod":           {i:"\u", c:[3]uint8{255, 255, 255}},       // gitpod
	// "advpl_prw":        {i:"\u", c:[3]uint8{255, 255, 255}},       // advpl_prw
	// "advpl_ptm":        {i:"\u", c:[3]uint8{255, 255, 255}},       // advpl_ptm
	// "advpl_tlpp":       {i:"\u", c:[3]uint8{255, 255, 255}},       // advpl_tlpp
	// "advpl_include":    {i:"\u", c:[3]uint8{255, 255, 255}},       // advpl_include
	// "tilt":             {i:"\u", c:[3]uint8{255, 255, 255}},       // tilt
	// "capacitor":        {i:"\u", c:[3]uint8{255, 255, 255}},       // capacitor
	// "adonis":           {i:"\u", c:[3]uint8{255, 255, 255}},       // adonis
	// "forth":            {i:"\u", c:[3]uint8{255, 255, 255}},       // forth
	// "uml":              {i:"\u", c:[3]uint8{255, 255, 255}},       // uml
	// "meson":            {i:"\u", c:[3]uint8{255, 255, 255}},       // meson
	// "buck":             {i:"\u", c:[3]uint8{255, 255, 255}},       // buck
	// "sml":              {i:"\u", c:[3]uint8{255, 255, 255}},       // sml
	// "nrwl":             {i:"\u", c:[3]uint8{255, 255, 255}},       // nrwl
	// "imba":             {i:"\u", c:[3]uint8{255, 255, 255}},       // imba
	// "drawio":           {i:"\u", c:[3]uint8{255, 255, 255}},       // drawio
	// "sas":              {i:"\u", c:[3]uint8{255, 255, 255}},       // sas
	// "slug":             {i:"\u", c:[3]uint8{255, 255, 255}},       // slug

	"dir-config":      {glyph: "\ue5fc", name: "nf-custom-folder_config", rgb: [3]uint8{32, 173, 194}}, // dir-config
	"dir-controller":  {glyph: "\ue5fc", name: "nf-custom-folder_config", rgb: [3]uint8{255, 194, 61}}, // dir-controller
	"dir-git":         {glyph: "\ue5fb", name: "nf-custom-folder_git", rgb: [3]uint8{250, 111, 66}},    // dir-git
	"dir-github":      {glyph: "\ue5fd", name: "nf-custom-folder_github", rgb: [3]uint8{84, 110, 122}}, // dir-github
	"dir-npm":         {glyph: "\ue5fa", name: "nf-custom-folder_npm", rgb: [3]uint8{203, 56, 55}},     // dir-npm
	"dir-include":     {glyph: "\uf756", name: "nf-mdi-folder_plus", rgb: [3]uint8{3, 155, 229}},       // dir-include
	"dir-import":      {glyph: "\uf756", name: "nf-mdi-folder_plus", rgb: [3]uint8{175, 180, 43}},      // dir-import
	"dir-upload":      {glyph: "\uf758", name: "nf-mdi-folder_upload", rgb: [3]uint8{250, 111, 66}},    // dir-upload
	"dir-download":    {glyph: "\uf74c", name: "nf-mdi-folder_download", rgb: [3]uint8{76, 175, 80}},   // dir-download
	"dir-secure":      {glyph: "\uf74f", name: "nf-mdi-folder_lock", rgb: [3]uint8{249, 169, 60}},      // dir-secure
	"dir-images":      {glyph: "\uf74e", name: "nf-mdi-folder_image", rgb: [3]uint8{43, 150, 137}},     // dir-images
	"dir-environment": {glyph: "\uf74e", name: "nf-mdi-folder_image", rgb: [3]uint8{102, 187, 106}},    // dir-environment
}

// default icons in case nothing can be found
var iconDefault = map[string]Icon{
	"dir":       {glyph: "\uf07b", name: "nf-fa-folder", rgb: [3]uint8{224, 177, 77}},
	"diropen":   {glyph: "\uf07c", name: "nf-fa-folder_open", rgb: [3]uint8{224, 177, 77}},
	"hiddendir": {glyph: "\uf114", name: "nf-fa-folder_o", rgb: [3]uint8{224, 177, 77}},
	"exe":       {glyph: "\uf15b", name: "nf-fa-file", rgb: [3]uint8{76, 175, 80}},
	"file":      {glyph: "\uf4a5", name: "nf-oct-file", rgb: [3]uint8{65, 129, 190}},
	// "hiddenfile": {glyph: "\ufb12", name: "nf-mdi-file_hidden", rgb: [3]uint8{65, 129, 190}}, // dotted outline folder
	"hiddenfile": {glyph: "\ue668", name: "nf-seti-ignored", rgb: [3]uint8{65, 129, 190}}, // eye with line through it
}
