package main

import (
	"encoding/json"
	"io/ioutil"
	"os/exec"
)

func SaveProject(note string) {
	p := &Project{}
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, p)

	add := exec.Command("git", "add", "-A")
	add.Run()

	commit := exec.Command("git", "commit", "-m", note)
	commit.Run()

	push := exec.Command("git", "push", p.Git, "master")
	push.Run()
}
