package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func (p *Project) FetchLibs() {
	var wg sync.WaitGroup
	for _, dep := range p.Depends {
		wg.Add(1)
		go p.GitCall(dep.Name, dep.Version, &wg)
	}
	wg.Wait()
}

func (p *Project) GitCall(link string, branch string, wg *sync.WaitGroup) error {
	repo := strings.Split(link, "/")

	cmd := exec.Command("git", "clone", "-b", branch, link, os.Getenv("GOPATH")+"src/"+strings.Join(repo[2:], "/"))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		wg.Done()
		return err
	}
	wg.Done()

	return nil
}

func GitInit() {
	init := exec.Command("git", "init")
	init.Run()
}

func BuildProject() {
	p := &Project{}
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Not able to read config.json")
		return
	}
	err = json.Unmarshal(data, p)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.FetchLibs()
	GitInit()
}
