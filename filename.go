package mode

import (
	"path/filepath"
	"strconv"
	"strings"
)

// configFilenames is a list of the most common configuration filenames that does not have an extension
var configFilenames = []string{"BUILD", "WORKSPACE", "config", "environment", "fstab", "group", "gshadow", "hostname", "hosts", "issue", "mirrorlist", "passwd", "shadow"}

// Detect looks at the filename and tries to guess what could be an appropriate editor mode.
func Detect(filename string) Mode {
	var mode Mode

	baseFilename := filepath.Base(filename)
	ext := filepath.Ext(baseFilename)

	// Check if we should be in a particular mode for a particular type of file
	// TODO: Create a hash map to look up many of the extensions
	switch {
	case baseFilename == "COMMIT_EDITMSG" ||
		baseFilename == "MERGE_MSG" ||
		(strings.HasPrefix(baseFilename, "git-") &&
			!strings.Contains(baseFilename, ".") &&
			strings.Count(baseFilename, "-") >= 2):
		// Git mode
		mode = Git
	case baseFilename == "svn-commit.tmp":
		mode = Subversion
	case ext == ".vimrc" || ext == ".vim" || ext == ".nvim":
		mode = Vim
	case ext == ".mk" || strings.HasPrefix(baseFilename, "Make") || strings.HasPrefix(baseFilename, "makefile") || baseFilename == "GNUmakefile":
		// NOTE: This one MUST come before the ext == "" check below!
		mode = Make
	case ext == ".just" || ext == ".justfile" || baseFilename == "justfile":
		// NOTE: This one MUST come before the ext == "" check below!
		mode = Just
	case strings.HasSuffix(filename, ".git/config") || ext == ".ini" || ext == ".cfg" || ext == ".conf" || ext == ".service" || ext == ".target" || ext == ".socket" || strings.HasPrefix(ext, "rc"):
		fallthrough
	case ext == ".yml" || ext == ".toml" || ext == ".ini" || ext == ".bp" || ext == ".rule" || strings.HasSuffix(filename, ".git/config") || (ext == "" && (strings.HasSuffix(baseFilename, "file") || strings.HasSuffix(baseFilename, "rc") || hasS(configFilenames, baseFilename))):
		mode = Config
	case ext == ".sh" || ext == ".install" || ext == ".ksh" || ext == ".tcsh" || ext == ".bash" || ext == ".zsh" || ext == ".local" || ext == ".profile" || baseFilename == "PKGBUILD" || baseFilename == "APKBUILD" || (strings.HasPrefix(baseFilename, ".") && strings.Contains(baseFilename, "sh")): // This last part covers .bashrc, .zshrc etc
		mode = Shell
	case ext == ".bzl" || baseFilename == "BUILD" || baseFilename == "WORKSPACE":
		mode = Bazel
	case baseFilename == "CMakeLists.txt" || ext == ".cmake":
		mode = CMake
	case strings.HasPrefix(baseFilename, "man.") && len(ext) > 4: // ie.: /tmp/man.0asdfadf
		mode = ManPage
	case strings.HasPrefix(baseFilename, "mutt-"): // ie.: /tmp/mutt-hostname-0000-0000-00000000000000000
		mode = Email
	case strings.HasSuffix(baseFilename, "Log.txt"): // ie. MinecraftLog.txt
		mode = Log
	default:
		switch ext {
		case ".1", ".2", ".3", ".4", ".5", ".6", ".7", ".8": // not .9
			mode = Nroff
		case ".adb", ".gpr", ".ads", ".ada":
			mode = Ada
		case ".adoc", ".scdoc", ".scd":
			mode = Doc
		case ".aidl":
			mode = AIDL
		case ".agda":
			mode = Agda
		case ".amber":
			mode = Amber
		case ".bas", ".module", ".frm", ".cls", ".ctl", ".vbp", ".vbg", ".form", ".gambas":
			mode = Basic
		case ".bat":
			mode = Bat
		case ".bts":
			mode = Battlestar
		case ".c":
			// C mode
			mode = C
		case ".cm":
			// Standard ML project file
			mode = StandardML
		case ".cpp", ".cc", ".c++", ".cxx", ".hpp", ".h": // C++ mode
			// TODO: Find a way to discover is a .h file is most likely to be C or C++
			mode = Cpp
		case ".clj", ".clojure", "cljs":
			mode = Clojure
		case ".cs": // C#
			mode = CS
		case ".cl", ".el", ".elisp", ".emacs", ".l", ".lisp", ".lsp":
			mode = Lisp
		case ".cr":
			mode = Crystal
		case ".d":
			mode = D
		case ".dart":
			mode = Dart
		case ".elm":
			mode = Elm
		case ".eml":
			mode = Email
		case ".erl":
			mode = Erlang
		case ".f", ".f77":
			mode = Fortran77
		case ".f90":
			mode = Fortran90
		case ".fs", ".fsx":
			mode = FSharp
		case ".gd":
			mode = GDScript
		case ".gt":
			mode = Garnet
		case ".go":
			mode = Go
		case ".glsl":
			mode = Shader
		case ".gradle":
			mode = Gradle
		case ".ha":
			mode = Hare
		case ".hal":
			mode = HIDL
		case ".hs", ".hts", ".cabal":
			mode = Haskell
		case ".htm", ".html":
			mode = HTML
		case ".hx", ".hxml":
			mode = Haxe
		case ".ino":
			mode = Arduino
		case ".ivy":
			mode = Ivy
		case ".jakt":
			mode = Jakt
		case ".java":
			mode = Java
		case ".js":
			mode = JavaScript
		case ".json", ".ipynb":
			mode = JSON
		case ".kk":
			mode = Koka
		case ".kt", ".kts":
			mode = Kotlin
		case ".log":
			mode = Log
		case ".lua":
			mode = Lua
		case ".m4":
			mode = M4
		case ".md":
			// Markdown mode
			mode = Markdown
		case ".ml":
			mode = OCaml // or standard ML, if the file does not contain ";;"
		case ".nim":
			mode = Nim
		case ".odin":
			mode = Odin
		case ".ok":
			mode = Oak
		case ".pas", ".pp", ".lpr":
			mode = ObjectPascal
		case ".pl", ".pro":
			mode = Prolog
		case ".py":
			mode = Python
		case ".mojo", "." + fireEmoji:
			mode = Mojo
		case ".r":
			mode = R
		case ".rs":
			mode = Rust
		case ".rst":
			mode = ReStructured // reStructuredText
		case ".s", ".S", ".asm", ".inc":
			// Go-style assembly (modeGoAssembly) is enabled if a mid-dot is discovered
			mode = Assembly
		case ".scala":
			mode = Scala
		case ".fun", ".sml":
			mode = StandardML
		case ".sql":
			mode = SQL
		case ".t":
			mode = Terra
		case ".te":
			mode = PolicyLanguage
		case ".tl":
			mode = Teal
		case ".ts":
			mode = TypeScript
		case ".txt", ".text", ".nfo", ".diz":
			mode = Text
		case ".v":
			mode = V
		case ".xml":
			mode = XML
		case ".zig", ".zir":
			mode = Zig
		default:
			mode = Blank
		}
	}

	// If the mode is not set, and there is no extensions
	if mode == Blank && !strings.Contains(baseFilename, ".") {
		if baseFilename == strings.ToUpper(baseFilename) {
			// If the filename is all uppercase and no ".", use mode.Markdown
			mode = Markdown
		} else if len(baseFilename) > 2 && baseFilename[2] == '-' {
			// Could it be a rule-file, that starts with ie. "90-" ?
			if _, err := strconv.Atoi(baseFilename[:2]); err == nil { // success
				// Yes, assume this is a shell-like configuration file
				mode = Config
			}
		}
	}

	return mode
}

// Exts returns a slice of glob strings for files that might be examined for this file mode,
// For example, "*.fs" and "*.fsx" for F#. Or BUILD, WORKSPACE and other filenames for mode.Config.
func (mode Mode) Globs() string {
	// TODO: Add a test that makes sure every mode has at least one ext
	switch mode {
	case Ada:
		return ".ada"
	case Agda:
		return ".agda"
	case AIDL:
		return "AIDL"
	case Amber:
		return "Amber"
	case Arduino:
		return "Arduino"
	case Assembly:
		return "Assembly"
	case Basic:
		return "Basic"
	case Bat:
		return "Bat"
	case Battlestar:
		return "Battlestar"
	case Bazel:
		return "Bazel"
	case Blank:
		return "-"
	case Clojure:
		return "Clojure"
	case CMake:
		return "CMake"
	case Config:
		return append(configFilenames, "config", "*.ini", "*.cfg", "*.conf", "*.service", "*.target", "*.socket", "*rc", "*.ini", "*.cfg")
	case Cpp:
		return "C++"
	case C:
		return "C"
	case Crystal:
		return "Crystal"
	case CS:
		return "C#"
	case Doc:
		return "Document"
	case D:
		return "D"
	case Dart:
		return "Dart"
	case Elm:
		return "Elm"
	case Email:
		return "E-mail"
	case Erlang:
		return "Erlang"
	case Fortran77:
		return "Fortran 77"
	case Fortran90:
		return "Fortran 90"
	case FSharp:
		return "F#"
	case Garnet:
		return "Garnet"
	case GDScript:
		return "Godot Script"
	case Git:
		return []string{"COMMIT_EDITMSG", "MERGE_MSG", "git-*"}
	case GoAssembly:
		return "Go-style Assembly"
	case Go:
		return "Go"
	case Gradle:
		return "Gradle"
	case Hare:
		return "Hare"
	case Haskell:
		return "Haskell"
	case Haxe:
		return "Haxe"
	case HIDL:
		return "HIDL"
	case HTML:
		return "HTML"
	case Ivy:
		return "Ivy"
	case Jakt:
		return "Jakt"
	case Java:
		return "Java"
	case JavaScript:
		return "JavaScript"
	case JSON:
		return "JSON"
	case Just:
		return []string{"*.just", "*.justfile", "justfile"}
	case Koka:
		return "Koka"
	case Kotlin:
		return "Kotlin"
	case Lisp:
		return "Lisp"
	case Log:
		return "Log"
	case Lua:
		return "Lua"
	case M4:
		return "M4"
	case Make:
		return []string{"makefile", "*.mk", "GNUmakefile", "Make*"}
	case ManPage:
		return "Man"
	case Markdown:
		return "Markdown"
	case Mojo:
		return "Mojo"
	case Nim:
		return "Nim"
	case Nroff:
		return "Nroff"
	case Oak:
		return "Oak"
	case ObjectPascal:
		return "Pas"
	case OCaml:
		return "Ocaml"
	case Odin:
		return "Odin"
	case Perl:
		return "Perl"
	case PolicyLanguage:
		return "SELinux"
	case Prolog:
		return "Prolog"
	case Python:
		return "Python"
	case R:
		return "R"
	case ReStructured:
		return "reStructuredText"
	case Rust:
		return "Rust"
	case Scala:
		return "Scala"
	case Shader:
		return "Shader"
	case Shell:
		return "Shell"
	case SQL:
		return "SQL"
	case StandardML:
		return "Standard ML"
	case Subversion:
		return []string{"svn-commit.tmp"}
	case Teal:
		return "Teal"
	case Terra:
		return "Terra"
	case Text:
		return "Text"
	case TypeScript:
		return "TypeScript"
	case Vim:
		return []string{".vimrc", "*.vimrc", "*.vim", "*.nvim"}
	case V:
		return "V"
	case XML:
		return "XML"
	case Zig:
		return "Zig"
	default:
		return "?"
	}
}
