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
		var sensitiveRanges [][2]token.Pos

		inSensitiveRange := func(p token.Pos) bool {
			for _, r := range sensitiveRanges {
				if p >= r[0] && p <= r[1] {
					return true
				}
			}
			return false
		}

		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			selection, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			ident, ok := selection.X.(*ast.Ident) // ident -- var
			if !ok || (ident.Name != "slog" && ident.Name != "log") {
				return true
			}

			switch selection.Sel.Name {
			case "Info", "Error", "Debug", "Warn":
				if len(call.Args) == 0 {
					return true
				}

				// 3. Чувствительные данные
				// Если в аргументах логов встречаются подозрительные идентификаторы,
				// отмечаем это как отдельную ошибку и пропускаем остальные этапы
				if containsSensitiveInArgs(call.Args) {
					pass.Reportf(call.Pos(), "log messages must not include potentially sensitive data (passwords, tokens, API keys)")
					sensitiveRanges = append(sensitiveRanges, [2]token.Pos{call.Pos(), call.End()})
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
					if !inSensitiveRange(literal.Pos()) {
						pass.Reportf(literal.Pos(), "log messages must be capitalized")
					}
				}

				// 2. Только английские буквы без юникод символов и эмоджи
				if !isEnglishOnly(msg) {
					if !inSensitiveRange(literal.Pos()) {
						pass.Reportf(literal.Pos(), "log messages must be in English and should not contain any Unicode characters or emojis.")
					}
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

var sensitiveKeys = []string{
	"password",
	"passwd",
	"pwd",
	"token",
	"api_key",
	"apikey",
	"api-key",
	"secret",
	"access_token",
	"auth_token",
}

func isSensitiveName(name string) bool {
	n := strings.ToLower(name)
	for _, k := range sensitiveKeys {
		if strings.Contains(n, k) {
			return true
		}
	}
	return false
}

func containsSensitiveInArgs(args []ast.Expr) bool {
	for _, a := range args {
		if exprContainsSensitive(a) {
			return true
		}
	}
	return false
}

func exprContainsSensitive(n ast.Node) bool {
	found := false
	ast.Inspect(n, func(x ast.Node) bool {
		if found {
			return false
		}
		switch v := x.(type) {
		case *ast.Ident:
			if isSensitiveName(v.Name) {
				found = true
				return false
			}
		case *ast.BasicLit:
			if v.Kind == token.STRING {
				txt := strings.Trim(v.Value, "\"")
				if isSensitiveName(txt) {
					found = true
					return false
				}
			}
		case *ast.SelectorExpr:
			if isSensitiveName(v.Sel.Name) {
				found = true
				return false
			}
		}
		return true
	})
	return found
}
