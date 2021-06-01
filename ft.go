package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// var MODULE = []string{"cont", "name"}

var dir, oldstr, newstr, file_path, include, exclude, module string
var includeList []string
var excludeList []string
var hidden bool
var dictionary bool
var mtime int

func main() {
	// defer golog.Sync()
	flag.StringVar(&module, "m", "", "是修改内容还是文件名 content | name | delete")
	flag.StringVar(&file_path, "f", "", "修改内容才支持文件路径")
	flag.StringVar(&dir, "d", "", "文件路径")
	flag.StringVar(&oldstr, "o", "", "原字符串")
	flag.StringVar(&newstr, "n", "", "新字符串")
	flag.BoolVar(&hidden, "H", true, "忽略隐藏文件")
	flag.BoolVar(&dictionary, "R", false, "是否删除目录(-m 为 delete 才生效)")
	flag.IntVar(&mtime, "t", 0, "选择删除多少天以前的")
	flag.StringVar(&include, "i", "", "指定包含字符串的文件名，逗号分隔多个, 修改内容才有效")
	flag.StringVar(&exclude, "e", "", "跳过指定包含字符串的文件名，逗号分隔多个， 修改内容才有效")
	flag.Parse()
	if include != "" {
		includeList = strings.Split(include, ",")
	}
	if exclude != "" {
		excludeList = strings.Split(exclude, ",")
	}

	if dir == "" && file_path == "" {
		log.Fatal("not specify file_path or path")
	}
	if dir != "" {
		if !filepath.IsAbs(dir) {
			pre_path, _ := filepath.Abs(".")
			dir = filepath.Join(pre_path, dir)
		}
	}
	switch module {
	case "content":
		// 修改内容
		if file_path != "" {
			replace(file_path)
		}
		if dir != "" {
			fileDir(dir)
		}
	case "name":
		walkdir(dir, oldstr, newstr)
	case "delete":
		walkDirDelete(dir)
	default:
		fmt.Println("module must be content or name")
	}

}

func fileDir(dirpath string) {
	infos, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		if hidden && info.Name()[0:1] == "." {
			continue
		}
		if info.IsDir() {
			fileDir(filepath.Join(dirpath, info.Name()))
		} else {
			if include == "" && exclude == "" {
				replace(filepath.Join(dirpath, info.Name()))
				continue
			}
			if include != "" && strInArray(info.Name(), includeList) {
				// 名字包含在内， 包含优先不包含
				fmt.Println(info.Name())
				replace(filepath.Join(dirpath, info.Name()))
				continue
			}
			if exclude != "" && !strInArray(info.Name(), excludeList) {
				// 名字包含在内， 包含优先不包含
				replace(filepath.Join(dirpath, info.Name()))
				continue
			}

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

func walkdir(thisdir string, old, new string) {
	fl, err := ioutil.ReadDir(thisdir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dir:", thisdir)
	for _, fi := range fl {
		if hidden && fi.Name()[0:1] == "." {
			continue
		}
		if fi.IsDir() {
			walkdir(filepath.Join(thisdir, fi.Name()), old, new)
		} else {
			if strings.Contains(fi.Name(), old) {
				os.Rename(filepath.Join(thisdir, fi.Name()),
					filepath.Join(thisdir, strings.Replace(fi.Name(), old, new, -1)))
			}
		}
	}
}

func walkDirDelete(thisdir string) {
	fl, err := ioutil.ReadDir(thisdir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fl {
		if fi.IsDir() {
			fmt.Println(dictionary)
			if dictionary {
				// 如果删除目录的话，那么不用递归，直接删除目录
				if len(include) > 0 && strInArray(fi.Name(), includeList) {
					if mtime == 0 || (mtime > 0 && time.Since(fi.ModTime()) >= time.Hour*24*time.Duration(mtime)) {
						os.RemoveAll(filepath.Join(thisdir, fi.Name()))
					}
					continue
				}
				os.RemoveAll(filepath.Join(thisdir, fi.Name()))
				continue
			}
			walkDirDelete(filepath.Join(thisdir, fi.Name()))

		} else {

			if len(include) > 0 && strInArray(fi.Name(), includeList) {
				if mtime == 0 || (mtime > 0 && time.Since(fi.ModTime()) >= time.Hour*24*time.Duration(mtime)) {
					fmt.Println("delete: ", filepath.Join(thisdir, fi.Name()))
					os.RemoveAll(filepath.Join(thisdir, fi.Name()))
				}
				continue
			}
			os.RemoveAll(filepath.Join(thisdir, fi.Name()))
		}
	}
}

func strInArray(str string, arr []string) bool {
	str = strings.Trim(str, " ")
	for _, v := range arr {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}
