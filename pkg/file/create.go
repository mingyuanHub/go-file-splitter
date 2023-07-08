package file

import (
	"fmt"
	"log"
	"os"
)

// 将字符串, 写入文本文件
func WriteToText(fold, fileName, content string) {
	wd, _ := os.Getwd()
	filePath :=  fmt.Sprintf("%s\\%s\\%s.txt", wd, fold, fileName)   // 存放小说的TXT文件路径
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)  // 以追加方式写入
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	_, err = fmt.Fprintln(f, content)
	if err != nil {
		log.Fatal(err)
	}
}