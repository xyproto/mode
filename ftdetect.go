package editormode

import (
	"path/filepath"
	"strings"
)

// Mode is a per-filetype mode, like for Markdown
type Mode int

const (
	// Mode "enum" values
	ModeBlank          = iota
	ModeGit            // for git commits and interactive rebases
	ModeMarkdown       // for Markdown (and asciidoctor and rst files)
	ModeMakefile       // for Makefiles
	ModeShell          // for shell scripts and PKGBUILD files
	ModeConfig         // for yml, toml, and ini files etc
	ModeAssembly       // for Assembly
	ModeGoAssembly     // for Go-style Assembly
	ModeGo             // for Go
	ModeHaskell        // for Haskell
	ModeOCaml          // for OCaml
	ModeStandardML     // for Standard ML
	ModePython         // for Python
	ModeText           // for plain text documents
	ModeCMake          // for CMake files
	ModeVim            // for Vim or NeoVim configuration, or .vim scripts
	ModeV              // the V programming language
	ModeClojure        // for Clojure
	ModeLisp           // for Common Lisp and Emacs Lisp
	ModeZig            // for Zig
	ModeKotlin         // for Kotlin
	ModeJava           // for Java
	ModeHIDL           // for the Android-related Hardware Abstraction Layer Interface Definition Language
	ModeSQL            // for Structured Query Language
	ModeOak            // for Oak
	ModeRust           // for Rust
	ModeLua            // for Lua
	ModeCrystal        // for Crystal
	ModeNim            // for Nim
	ModeObjectPascal   // for Object Pascal and Delphi
	ModeBat            // for DOS batch files
	ModeCpp            // for C++
	ModeC              // for C
	ModeAda            // for Ada
	ModeHTML           // for HTML
	ModeOdin           // for Odin
	ModeXML            // for XML
	ModePolicyLanguage // for SE Linux configuration files
	ModeNroff          // for editing man pages
	ModeScala          // for Scala
	ModeJSON           // for JSON and iPython notebooks
	ModeBattlestar     // for Battlestar
	ModeCS             // for C#
	ModeJavaScript     // for JavaScript
	ModeTypeScript     // for TypeScript
	ModeManPage        // for viewing man pages
	ModeAmber          // for Amber templates
	ModeBazel          // for Bazel and Starlark
	ModeD              // for D
	ModePerl           // for Perl
	ModeM4             // for M4 macros
)

// Detect looks at the filename and tries to guess what could be an appropriate editor mode.
// This mainly affects syntax highlighting (which can be toggled with ctrl-t) and indentation.
func Detect(filename string) (Mode, bool) {

	// A list of the most common configuration filenames that does not have an extension
	var (
		configFilenames = []string{"fstab", "config", "BUILD", "WORKSPACE", "passwd", "group", "environment", "shadow", "gshadow", "hostname", "hosts", "issue", "mirrorlist"}
		mode            Mode
	)

	baseFilename := filepath.Base(filename)
	ext := filepath.Ext(baseFilename)

	// Check if we should be in a particular mode for a particular type of file
	switch {
	case baseFilename == "COMMIT_EDITMSG" ||
		baseFilename == "MERGE_MSG" ||
		(strings.HasPrefix(baseFilename, "git-") &&
			!strings.Contains(baseFilename, ".") &&
			strings.Count(baseFilename, "-") >= 2):
		// Git mode
		mode = ModeGit
	case ext == ".vimrc" || ext == ".vim" || ext == ".nvim":
		mode = ModeVim
	case strings.HasPrefix(baseFilename, "Makefile") || strings.HasPrefix(baseFilename, "makefile") || baseFilename == "GNUmakefile":
		// NOTE: This one MUST come before the ext == "" check below!
		mode = ModeMakefile
	case strings.HasSuffix(filename, ".git/config") || ext == ".ini" || ext == ".cfg" || ext == ".conf" || ext == ".service" || ext == ".target" || ext == ".socket" || strings.HasPrefix(ext, "rc"):
		fallthrough
	case ext == ".yml" || ext == ".toml" || ext == ".ini" || ext == ".bp" || strings.HasSuffix(filename, ".git/config") || (ext == "" && (strings.HasSuffix(baseFilename, "file") || strings.HasSuffix(baseFilename, "rc") || hasS(configFilenames, baseFilename))):
		mode = ModeConfig
	case ext == ".sh" || ext == ".ksh" || ext == ".tcsh" || ext == ".bash" || ext == ".zsh" || ext == ".local" || ext == ".profile" || baseFilename == "PKGBUILD" || (strings.HasPrefix(baseFilename, ".") && strings.Contains(baseFilename, "sh")): // This last part covers .bashrc, .zshrc etc
		mode = ModeShell
	case ext == ".bzl" || baseFilename == "BUILD" || baseFilename == "WORKSPACE":
		mode = ModeBazel
	case baseFilename == "CMakeLists.txt" || ext == ".cmake":
		mode = ModeCMake
	default:
		switch ext {
		case ".s", ".S", ".asm", ".inc":
			// Go-style assembly (modeGoAssembly) is enabled if a mid-dot is discovered
			mode = ModeAssembly
		//case ".s":
		//mode = ModeGoAssembly
		case ".amber":
			mode = ModeAmber
		case ".go":
			mode = ModeGo
		case ".odin":
			mode = ModeOdin
		case ".hs":
			mode = ModeHaskell
		case ".sml":
			mode = ModeStandardML
		case ".m4":
			mode = ModeM4
		case ".ml":
			mode = ModeOCaml // or standard ML, if the file does not contain ";;"
		case ".py":
			mode = ModePython
		case ".pl":
			mode = ModePerl
		case ".md":
			// Markdown mode
			mode = ModeMarkdown
		case ".bts":
			mode = ModeBattlestar
		case ".cpp", ".cc", ".c++", ".cxx", ".hpp", ".h":
			// C++ mode
			// TODO: Find a way to discover is a .h file is most likely to be C or C++
			mode = ModeCpp
		case ".c":
			// C mode
			mode = ModeC
		case ".d":
			// D mode
			mode = ModeD
		case ".cs":
			// C# mode
			mode = ModeCS
		case ".adoc", ".rst", ".scdoc", ".scd":
			// Markdown-like syntax highlighting
			// TODO: Introduce a separate mode for these.
			mode = ModeMarkdown
		case ".txt", ".text", ".nfo", ".diz":
			mode = ModeText
		case ".clj", ".clojure", "cljs":
			mode = ModeClojure
		case ".lsp", ".emacs", ".el", ".elisp", ".lisp", ".cl", ".l":
			mode = ModeLisp
		case ".zig", ".zir":
			mode = ModeZig
		case ".v":
			mode = ModeV
		case ".kt", ".kts":
			mode = ModeKotlin
		case ".java", ".gradle":
			mode = ModeJava
		case ".hal":
			mode = ModeHIDL
		case ".sql":
			mode = ModeSQL
		case ".ok":
			mode = ModeOak
		case ".rs":
			mode = ModeRust
		case ".lua":
			mode = ModeLua
		case ".cr":
			mode = ModeCrystal
		case ".nim":
			mode = ModeNim
		case ".pas", ".pp", ".lpr":
			mode = ModeObjectPascal
		case ".bat":
			mode = ModeBat
		case ".adb", ".gpr", ".ads", ".ada":
			mode = ModeAda
		case ".htm", ".html":
			mode = ModeHTML
		case ".xml":
			mode = ModeXML
		case ".te":
			mode = ModePolicyLanguage
		case ".1", ".2", ".3", ".4", ".5", ".6", ".7", ".8":
			mode = ModeNroff
		case ".scala":
			mode = ModeScala
		case ".json", ".ipynb":
			mode = ModeJSON
		case ".js":
			mode = ModeJavaScript
		case ".ts":
			mode = ModeTypeScript
		default:
			mode = ModeBlank
		}
	}

	if mode == ModeText {
		mode = ModeMarkdown
	}

	// If the mode is not set and the filename is all uppercase and no ".", use modeMarkdown
	if mode == ModeBlank && !strings.Contains(baseFilename, ".") && baseFilename == strings.ToUpper(baseFilename) {
		mode = ModeMarkdown
	}

	// Check if we should enable syntax highlighting by default
	syntaxHighlightingEnabled := (mode != ModeBlank || ext != "") && mode != ModeText

	return mode, syntaxHighlightingEnabled
}

// String will return a short lowercase string representing the given editor mode
func (mode Mode) String() string {
	switch mode {
	case ModeBlank:
		return "-"
	case ModeGit:
		return "Git"
	case ModeMarkdown:
		return "Markdown"
	case ModeMakefile:
		return "Make"
	case ModeShell:
		return "Shell"
	case ModeConfig:
		return "Configuration"
	case ModeAssembly:
		return "Assembly"
	case ModeGoAssembly:
		return "Go-style Assembly"
	case ModeGo:
		return "Go"
	case ModeHaskell:
		return "Haskell"
	case ModeOCaml:
		return "Ocaml"
	case ModeStandardML:
		return "Standard ML"
	case ModePython:
		return "Python"
	case ModeText:
		return "Text"
	case ModeCMake:
		return "Cmake"
	case ModeVim:
		return "ViM"
	case ModeClojure:
		return "Clojure"
	case ModeLisp:
		return "Lisp"
	case ModeZig:
		return "Zig"
	case ModeKotlin:
		return "Kotlin"
	case ModeJava:
		return "Java"
	case ModeHIDL:
		return "HIDL"
	case ModeSQL:
		return "SQL"
	case ModeOak:
		return "Oak"
	case ModeRust:
		return "Rust"
	case ModeLua:
		return "Lua"
	case ModeCrystal:
		return "Crystal"
	case ModeNim:
		return "Nim"
	case ModeObjectPascal:
		return "Pas"
	case ModeBat:
		return "Bat"
	case ModeCpp:
		return "C++"
	case ModeC:
		return "C"
	case ModeAda:
		return "Ada"
	case ModeHTML:
		return "HTML"
	case ModeOdin:
		return "Odin"
	case ModePerl:
		return "Perl"
	case ModeXML:
		return "XML"
	case ModePolicyLanguage:
		return "SELinux"
	case ModeNroff:
		return "Nroff"
	case ModeScala:
		return "Scala"
	case ModeJSON:
		return "JSON"
	case ModeBattlestar:
		return "Battlestar"
	case ModeCS:
		return "C#"
	case ModeTypeScript:
		return "TypeScript"
	case ModeJavaScript:
		return "JavaScript"
	case ModeManPage:
		return "Man"
	case ModeAmber:
		return "Amber"
	case ModeBazel:
		return "Bazel"
	case ModeD:
		return "D"
	case ModeV:
		return "V"
	case ModeM4:
		return "M4"
	default:
		return "?"
	}
}
