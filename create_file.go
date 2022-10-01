package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func (a Artisan) CreateFile(p string, filename string, fileContents string) *os.File {
	createFile, err := os.Create(path.Join(p, filename))

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer createFile.Close()

	fmt.Fprint(createFile, fileContents)

	return createFile
}
