// Package main реализует multichecker.
//
// # Сборка и запуск
//
// Для сборки линтера выполните:
//
//	go build -o staticlint ./cmd/staticlint
//
// Затем запустите его в корне вашего проекта:
//
//	./staticlint ./...
//
// Можно передавать стандартные флаги анализаторов, например -tests для включения
// тестовых файлов или -help для просмотра всех доступных анализаторов и их флагов.
//
// # Включённые анализаторы
//
// Мультичекер состоит из следующих групп:
//
//  1. Включены стандартные анализаторы из пакета golang.org/x/tools/go/analysis/passes.
//     - printf    : проверка согласованности форматных строк
//     - shadow    : проверка на затенение переменных
//     - structtag : проверка корректности тегов полей структур
//
//  2. Все анализаторы staticcheck класса SA (staticcheck.io).
//
//  3. Один анализатор из другого класса staticcheck:
//     - ST1003: проверяет что идентификаторы, имена пакетов и переменные соответствует правилам.
//
//  4. Два публичных анализатора:
//     - errcheck   (github.com/kisielk/errcheck) – проверяет, что все возвращаемые ошибки обрабатываются.
//     - bodyclose  (github.com/timakin/bodyclose) – проверяет, что тела HTTP-ответов закрываются.
//
//  5. Пользовательский анализатор:
//     - exitinmain – запрещает прямой вызов os.Exit внутри функции main пакета main.
package main

import (
	"go/ast"
	"strings"

	"github.com/kisielk/errcheck/errcheck"
	"github.com/timakin/bodyclose/passes/bodyclose"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/stylecheck/st1003"
)

func main() {
	var sachecks []*analysis.Analyzer
	sachecks = append(sachecks, printf.Analyzer, shadow.Analyzer, structtag.Analyzer)
	for _, v := range staticcheck.Analyzers {
		if strings.HasPrefix(v.Analyzer.Name, "SA") {
			sachecks = append(sachecks, v.Analyzer)
		}
	}

	sachecks = append(sachecks, st1003.Analyzer, errcheck.Analyzer, bodyclose.Analyzer)

	checkMain := &analysis.Analyzer{
		Name: "exitinmain",
		Doc:  "проверка вызова os.Exit в main функции main пакета",
		Run:  runExitInMain,
	}

	sachecks = append(sachecks, checkMain)

	multichecker.Main(
		sachecks...,
	)
}

func runExitInMain(pass *analysis.Pass) (any, error) {
	if pass.Pkg.Name() != "main" {
		return nil, nil
	}

	for _, file := range pass.Files {
		fname := pass.Fset.File(file.Pos()).Name()

		if !strings.Contains(fname, "/cmd/shortener") { // более грубый, можно точнее
			continue
		}
		ast.Inspect(file, func(funcNode ast.Node) bool {
			funcDecl, ok := funcNode.(*ast.FuncDecl)
			if !ok {
				return true
			}
			if funcDecl.Name.Name == "main" && funcDecl.Recv == nil {
				ast.Inspect(funcDecl.Body, func(expNode ast.Node) bool {
					call, ok := expNode.(*ast.CallExpr)
					if !ok {
						return true
					}
					sel, ok := call.Fun.(*ast.SelectorExpr)
					if !ok {
						return true
					}
					if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "os" && sel.Sel.Name == "Exit" {
						pass.Reportf(call.Pos(), "direct call to os.Exit in main function is forbidden")
					}
					return true
				})
			}
			return true
		})
	}
	return nil, nil
}
