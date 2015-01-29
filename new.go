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
	err := os.Mkdir(p.Name, 0600)
	if err != nil {
		panic(err)
		return
	}
	os.Mkdir(fmt.Sprintf("%s/public", p.Name), 0600)
	os.Mkdir(fmt.Sprintf("%s/public/js", p.Name), 0600)
	os.Mkdir(fmt.Sprintf("%s/public/css", p.Name), 0600)
	os.Mkdir(fmt.Sprintf("%s/public/fonts", p.Name), 0600)

	os.Mkdir(fmt.Sprintf("%s/libs", p.Name), 0600)
	os.Mkdir(fmt.Sprintf("%s/template", p.Name), 0600)

	file, _ := os.Create(fmt.Sprintf("%s/app.go", p.Name))
	file.WriteString(APPCONTENT)
	file.Close()
}
