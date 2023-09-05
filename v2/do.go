package v2

import (
	"fmt"
	"go-file-splitter/pkg/common"
	"go-file-splitter/pkg/file"
	"os"
	"path/filepath"
	"strings"
)

var (
	ASSETS = "./v2/assets/"

	SPLIT_KEY = "文件"

	oldPdfDir = "外文数据"
	newPdfDir = "外文数据-切割"

	DirArr = []string{
		"阿根廷",
		"埃及",
		"巴基斯坦",
		"巴西",
		"俄罗斯",
		"法媒",
		"韩国",
		"美媒",
		"秘鲁",
		"南非",
		"尼日利亚",
		"日媒",
		"沙特",
		"乌克兰",
		"希腊",
		"新加坡",
		"伊朗",
		"印度",
		"英媒",
		"越南",
	}
)

func Do() {

	var (
		files []string
	)

	err := filepath.Walk(ASSETS + oldPdfDir, func(path string, info os.FileInfo, err error) error {
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

		if strings.Index(f, "日媒") != -1 || strings.Index(f, "韩国") != -1 {
			go wr(f)
		}
	}
}

func wr(f string) {
	var (
		dir      string
		fileName = strings.Split(f, "\\")[4]
	)

	for _, key := range DirArr {
		if strings.Contains(f, key) {
			dir = ASSETS + newPdfDir + "/" + key
		}
	}

	if len(dir) == 0 {
		fmt.Println(11111111, f, dir, fileName)
	}

	content, err := file.ReadFileUtf8(f)
	if err != nil {
		fmt.Println("file.ReadFile err=", err.Error())
	}

	spans := strings.Split(content, SPLIT_KEY)

	for idx, span := range spans {
		if idx > 0 &&  len(span) > 27 {
			span = span[27:]
		}
		if idx+1 == len(spans) {
			continue
		}
		fmt.Println(dir, fileName)
		file.WriteToText(dir, fmt.Sprintf("%s-%d-%d", fileName, idx+1, common.GetNowMillisecond()), span)
	}

	//fmt.Println(f)

}