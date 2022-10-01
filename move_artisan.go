package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func (a Artisan) MoveArtisan(project_name string) {
	osType := runtime.GOOS

	vueArtisanBinary := ""
	switch osType {
	case "windows":
		vueArtisanBinary = "vue_artisan.exe"
	case "linux":
		vueArtisanBinary = "vue_artisan"
	}

	input, err := ioutil.ReadFile(vueArtisanBinary)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = ioutil.WriteFile(path.Join(project_name, vueArtisanBinary), input, 0644)
	if err != nil {
		fmt.Println("Error creating destination file: vue_artisan")
		fmt.Println(err.Error())
		return
	}

	e := os.Remove(vueArtisanBinary)

	if e != nil {
		log.Fatalln("Error removing ", vueArtisanBinary)
	}

	cmd := exec.Command("chmod", "+x", vueArtisanBinary)
	cmd.Dir = path.Join(project_name)

	_, err = cmd.Output()

	if err != nil {
		log.Fatalln("Error making vue_artisan executable: ", err.Error())
	}
}
