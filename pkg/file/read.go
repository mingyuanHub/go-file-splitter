package file

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
)

func ReadFile(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	gbkReader := bytes.NewReader(content)
	utf8Reader := transform.NewReader(gbkReader, simplifiedchinese.GBK.NewDecoder())
	utf8Bytes, err := ioutil.ReadAll(utf8Reader)

	return string(utf8Bytes), nil
}
