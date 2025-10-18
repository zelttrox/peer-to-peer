package net

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// File type
type File struct {
	Name string
	Path string
	Size string
	Ext string
}

// Turn a filepath into a File type variable
func GetFile(path string) File {
	stat, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	name := strings.Split(path, "/")
	size := float64(stat.Size()) / 1000000.0
	return File {
		name[len(name)-1], 
		path, 
		fmt.Sprintf("%.2f Mo", size), 
		filepath.Ext(path),
	}
}

// Process all bytes of a file
func ProcessFile(file File) {

}
