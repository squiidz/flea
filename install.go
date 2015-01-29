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
	place, _ := os.Getwd()
	repo := strings.Split(link, "/")

	cmd := exec.Command("git", "clone", "-b", branch, link, place+"/libs/"+repo[len(repo)-1])
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
}
