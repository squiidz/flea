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

func ReadConfig() *Project {
	p := &Project{}
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, p)
	return p
}

func (p *Project) WriteConfig() {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("config.json", data, 0600)
	if err != nil {
		panic(err)
	}
}

func (p *Project) FetchLibs() {
	var wg sync.WaitGroup
	for _, dep := range p.Depends {
		wg.Add(1)
		go GitCall(dep.Name, dep.Version, &wg)
	}
	wg.Wait()
}

func GitCall(link string, branch string, wg *sync.WaitGroup) error {
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

func LoadDepen(src string) {
	var wg sync.WaitGroup
	p := ReadConfig()
	wg.Add(1)
	err := GitCall(src, "master", &wg)
	if err != nil {
		fmt.Printf("[!] Error at loading %s\n", src)
		return
	}
	p.Depends = append(p.Depends, Dependance{Name: src, Version: "master"})
	fmt.Printf("[+] Successfuly install %s\n", src)
	p.WriteConfig()
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
