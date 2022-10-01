package main

import (
	"fmt"
	"os/exec"
)

func (a Artisan) CreateDirectory(path string, dirname string) bool {
	cmd := exec.Command("mkdir", dirname)
	cmd.Dir = path

	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
