package main

import (
	"fmt"
	"go-file-splitter/pkg/file"
	"go-file-splitter/pkg/common"
	"os"
	"path/filepath"
	"strings"
)

var (
	ASSETS = "./assets"

	SPLIT_KEY = "文件"

	DirMap = map[string]string{
		"德语" : ASSETS + "/deyu",
		"法语" : ASSETS + "/fayu",
		"日语" : ASSETS + "/riyu",
		"英语" : ASSETS + "/yingyu",
		"意语" : ASSETS + "/yiyu",
	}
)

func main() {

	var (
		files []string
	)

	err := filepath.Walk(ASSETS, func(path string, info os.FileInfo, err error) error {
		if !strings.Contains(path, ".txt") {
			return nil
		}
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, f := range files {

		var (
			l        string
			dir      string
			fileName = strings.Split(f, "\\")[2]
		)

		for key, _ := range DirMap {
			if strings.Contains(f, key) {
				l = key
				dir, _ = DirMap[key]
			}
		}

		content, err := file.ReadFile(f)
		if err != nil {
			fmt.Println("file.ReadFile err=", err.Error())
		}

		spans := strings.Split(content, SPLIT_KEY)

		//fmt.Println("111111111", content[0:500], len(spans) )

		for idx, span := range spans {
			if idx > 0 {
				span = span[27:]
			}
			if idx+1 == len(spans) {
				continue
			}
			file.WriteToText(dir, fmt.Sprintf("%s-%s-%d-%d", l, fileName, idx+1, common.GetNowMillisecond()), span)
		}

		fmt.Println(f)
	}
}