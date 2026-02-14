package mylogger

import (
	"go/ast"
	"go/token"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "selectellinter",
	Doc:  "Test exercise for Selectel",
	Run:  run,
}

var englishLetters = regexp.MustCompile(`^[a-zA-Z]+$`)

/*
type Pass struct {
	Fset       *token.FileSet // file position information
	Files      []*ast.File    // the abstract syntax tree of each file
	OtherFiles []string       // names of non-Go files of this package
	Pkg        *types.Package // type information about the package
	TypesInfo  *types.Info    // type information about the syntax trees
	TypesSizes types.Sizes    // function for computing sizes of types
	...
}
*/

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			selection, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// ident -- var
			ident, ok := selection.X.(*ast.Ident)
			if !ok || (ident.Name != "slog" && ident.Name != "log") {
				return true
			}

			switch selection.Sel.Name {
			case "Info", "Error", "Debug", "Warn":
				if len(call.Args) == 0 {
					return true
				}

				literal, ok := call.Args[0].(*ast.BasicLit)
				if !ok || literal.Kind != token.STRING {
					return true
				}

				// 1. Лог-сообщения должны начинаться со строчной буквы
				msg := strings.Trim(literal.Value, "\"")
				if len(msg) == 0 {
					return true
				}

				first := []rune(msg)[0]
				if unicode.IsUpper(first) {
					pass.Reportf(literal.Pos(), "log messages must be capitalized")
				}

				// 2. Только английские буквы без юникод символов и эмоджи
				if !isEnglishOnly(msg) {
					pass.Reportf(literal.Pos(), "log messages must be on English language")
				}
			}

			// 3. безопасность данных
			switch selection.Sel.Name {
			case "Info", "Error", "Debug", "Warn":
				if len(call.Args) == 0 {
					return true
				}
				be, ok := n.(*ast.BinaryExpr)
				if !ok {
					return true
				}

				if be.Op == token.ADD {
					if ident, ok := be.Y.(*ast.Ident); ok {
						substrings := []string{"password", "key", "api", "pass"}
						for _, sub := range substrings {
							if strings.Contains(ident.Name, sub) {
								pass.Reportf(ident.Pos(), "cannot log sensitive data directly")
							}
						}
						return false
					}
					return false
				}
			}

			return true
		})
	}
	return nil, nil

}

func isEnglishOnly(s string) bool {
	return englishLetters.MatchString(s)
}
