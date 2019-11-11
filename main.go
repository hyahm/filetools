package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var path, oldstr, newstr string

func main() {

	flag.StringVar(&path, "p", ".", "文件路径")
	flag.StringVar(&oldstr, "o", "", "原字符串")
	flag.StringVar(&newstr, "n", "", "新字符串")
	flag.Parse()

	dir(path)
}

func dir(dirpath string) {

	infos, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		if info.IsDir() {
			dir(filepath.Join(dirpath, info.Name()))
		} else {
			replace(filepath.Join(dirpath, info.Name()))
		}
	}
}

func replace(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	newdata := bytes.Replace(data, []byte(oldstr), []byte(newstr), -1)
	os.Remove(filename)
	ioutil.WriteFile(filename, newdata, 0644)
}
