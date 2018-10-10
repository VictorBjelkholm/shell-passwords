package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Password struct {
	UUID    string `json:"uuid"`
	Details struct {
		Fields []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"fields"`
	} `json:"details"`
}

func get(key string) string {
	cmd := exec.Command("sh", "-c", "op get item "+key)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(stdoutStderr))
		log.Fatal(err)
	}
	var pw Password
	err = json.Unmarshal(stdoutStderr, &pw)
	if err != nil {
		panic(err)
	}
	for _, p := range pw.Details.Fields {
		if p.Type == "P" {
			return p.Value
		}
	}
	return ""
}

func filter(key string) {
	password := get(key)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, password) {
			pwLen := len(password)
			pwReplace := ""
			for len(pwReplace) != pwLen {
				pwReplace = pwReplace + "*"
			}
			fmt.Println(strings.Replace(text, password, pwReplace, -1))
		} else {
			fmt.Println(text)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Not enough args. Usage: <sp [g/f] $key> <sp get/filter $key>")
	}

	cmdToRun := os.Args[1]
	pwToGrab := os.Args[2]

	if cmdToRun == "g" {
		pw := get(pwToGrab)
		fmt.Print(pw)
	}

	if cmdToRun == "f" {
		filter(pwToGrab)
	}

	if cmdToRun != "g" && cmdToRun != "f" {
		log.Fatal("Only g[et] and f[ilter] is supported")
	}

}
