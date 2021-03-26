package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

//go:embed markdown.tpl
var markdownTPL string

func main() {
	var importfile = flag.String("i", "", "需要转换的postman导出文件")
	var exportfile = flag.String("o", "./export.md", "输出文件名")
	flag.Parse()
	pJson, err := ioutil.ReadFile(*importfile)
	if err != nil {
		fmt.Printf("读取文件错误:%v\n", err)
		os.Exit(-1)
	}
	var postmanData *PostManStructV2Point1
	err = json.Unmarshal(pJson, &postmanData)
	if err != nil {
		fmt.Printf("json解析错误:%v\n", err)
		os.Exit(-1)
	}
	postmanData.Item = FlattenFolder(postmanData.Item)
	tpl, err := template.New("markdown").Parse(string(markdownTPL))
	if err != nil {
		fmt.Printf("模版错误%v\n", err)
		os.Exit(-1)
	}
	o, err := os.OpenFile(*exportfile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("输出路径错误%v\n", err)
		os.Exit(-1)
	}
	defer o.Close()

	err = tpl.ExecuteTemplate(o, "markdown", postmanData)
	if err != nil {
		fmt.Printf("模版错误%v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("转换完成")
}

func FlattenFolder(data []ItemStructV2Point1) (item []ItemStructV2Point1) {
	item = make([]ItemStructV2Point1, 0)
	for i := range data {
		if len(data[i].Item) == 0 {
			item = append(item, data[i])
			continue
		}
		item = append(item, FlattenFolder(data[i].Item)...)
	}
	return
}
