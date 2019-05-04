package file

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// 判断所给路径文件/文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

// 会递归创建文件夹
func CreateDir(dirPath string) error {
	if !IsExist(dirPath) {
		return os.MkdirAll(dirPath, os.ModePerm)
	}

	return nil
}

// 保存文件
func SaveFile(file *multipart.FileHeader, filePath, fileName string) error {
	if err := CreateDir(filePath); err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(path.Join(filePath, fileName))
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, src); err != nil {
		return err
	}

	return nil
}

// ReadFile 读取文件内容
func ReadFile(filePath string) (string, error) {
	fmt.Println(os.Getwd())
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

// ReadTemplateToString 读取模板并转换为 string
func ReadTemplateToString(tplName string, tplPath string, tplData map[string]interface{}) (string, error) {
	t, err := template.New(tplName).ParseFiles(tplPath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, tplData); err != nil {
		return "", err
	}

	return buf.String(), nil
}
