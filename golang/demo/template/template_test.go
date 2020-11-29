/// 所谓模板 text/template实现基于数据的文本化输出
package template

import (
	"log"
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestTemplate1(t *testing.T) {
	// 首先创建一个函数字典用于注册函数
	funcMap := template.FuncMap{
		// 注册函数title, strings.Title会将单词首字母大写
		"title": strings.Title,
	}

	// A simple template definition to test our function.
	// 打印输出文本
	// - 原样
	// - 执行title后的
	// - 执行title和printf
	// - 执行printf和title
	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
`

	// 创建模板, 添加模板函数,添加解析模板文本.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// 运行模板，出入数据参数
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}

func TestTemplate2(t *testing.T) {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}
