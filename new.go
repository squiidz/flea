package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func NewProject(name string) {
	p := New(name)
	p.buildTree()
	p.buildConfig()
}

func (p *Project) buildConfig() {
	file, err := os.Create(fmt.Sprintf("%s/config.json", p.Name))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}

	file.Write(data)
}

func (p *Project) buildTree() {
	err := os.Mkdir(p.Name, 0777)
	if err != nil {
		panic(err)
		return
	}
	os.Mkdir(fmt.Sprintf("%s/public", p.Name), 0777)
	os.Mkdir(fmt.Sprintf("%s/public/js", p.Name), 0777)
	os.Mkdir(fmt.Sprintf("%s/public/css", p.Name), 0777)
	os.Mkdir(fmt.Sprintf("%s/public/fonts", p.Name), 0777)

	os.Mkdir(fmt.Sprintf("%s/lib", p.Name), 0777)

	os.Mkdir(fmt.Sprintf("%s/template", p.Name), 0777)
	file1, _ := os.Create(fmt.Sprintf("%s/template/index.html", p.Name))
	file1.WriteString(INDEX)
	file1.Close()

	file2, _ := os.Create(fmt.Sprintf("%s/app.go", p.Name))
	file2.WriteString(APPCONTENT)
	file2.Close()
}
