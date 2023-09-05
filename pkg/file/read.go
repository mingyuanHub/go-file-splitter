package file

import (
	"bytes"
	"github.com/ledongthuc/pdf"
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

func ReadFileUtf8(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	utf8Reader := bytes.NewReader(content)
	utf8Bytes, err := ioutil.ReadAll(utf8Reader)

	return string(utf8Bytes), nil
}

// ReadPdf 获取pdf文字内容
func ReadPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}