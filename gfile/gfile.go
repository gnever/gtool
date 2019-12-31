package gfile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/util/gconv"
)

func Mkdir(path string) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	return
}

//Create 创建文件。若文件存在则覆盖
func Create(fileName string) (err error) {
	dir := filepath.Dir(fileName)
	if !Exists(dir) {
		if err = Mkdir(dir); err != nil {
			return
		}
	}
	_, err = os.Create(fileName)
	return
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return os.IsExist(err)
	}
	return true
}

func IsFile(path string) bool {
	if s, err := os.Stat(path); err == nil {
		return !s.IsDir()
	}

	return false
}

func IsDir(path string) bool {
	if s, err := os.Stat(path); err == nil {
		return s.IsDir()
	}

	return false
}

func Move(path string, dst string) error {
	return os.Rename(path, dst)
}

func Copy(path string, dst string) error {
	if !Exists(path) {
		return fmt.Errorf("%s not exists", path)
	}

	if IsFile(path) {
		return copyFile(path, dst)
	} else {
		return copyDir(path, dst)
	}
}

func copyFile(path string, dst string) error {
	if !IsFile(path) {
		return fmt.Errorf("%s not a file", path)
	}

	if Exists(dst) {
		return fmt.Errorf("%s alerady exists", dst)
	}

	src, err := os.Open(path)
	if err != nil {
		return err
	}
	defer src.Close()

	if err := Create(dst); err != nil {
		return err
	}

	dstIn, err := os.OpenFile(dst, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer dstIn.Close()
	io.Copy(dstIn, src)
	if err := dstIn.Sync(); err != nil {
		return err
	}

	//保持权限不变
	srcStat, err := os.Stat(path)
	if err != nil {
		return err
	}

	return os.Chmod(dst, srcStat.Mode())
}

func copyDir(path string, dst string) (err error) {
	path = filepath.Clean(path)
	dst = filepath.Clean(dst)
	fmt.Println()

	if !IsDir(path) {
		return fmt.Errorf("%s not dir", path)
	}

	if !Exists(path) {
		return fmt.Errorf("%s not exists", path)
	}

	if Exists(dst) {
		return fmt.Errorf("%s alerady exists", dst)
	}

	if err = Mkdir(dst); err != nil {
		return
	}

	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	for _, fileInfo := range fileInfos {
		_path := filepath.Join(path, fileInfo.Name())
		_dst := filepath.Join(dst, fileInfo.Name())
		if fileInfo.IsDir() {
			err = copyDir(_path, _dst)
		} else {
			err = copyFile(_path, _dst)
		}
		if err != nil {
			return
		}
	}

	return nil
}

func Remove(path string) error {
	return os.RemoveAll(path)
}

//Basename 返回 path 中最后一个分隔符之后的部分(不包含分隔符)
func Basename(path string) string {
	return filepath.Base(path)
}

// Name path中最后一个分隔符之后的部分(不包含分隔符),若文件包含 .后缀 的扩展。则去掉
func Name(path string) string {
	basename := Basename(path)

	if i := strings.LastIndexByte(basename, '.'); i != -1 {
		return basename[0:i]
	}
	return basename
}

//Ext 包含 . ,比如 .txt
func Ext(fileName string) string {
	return filepath.Ext(fileName)
}

//ExtName 不包含 . 的扩展名
func ExtName(fileName string) string {
	return strings.TrimLeft(Ext(fileName), ".")
}

//全部读取
func GetContents(file string) (string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

//ReadLines 以 string 的格式按行读取
func ReadLines(file string, callback func(text string)) error {
	cb := func(bytes []byte) {
		callback(gconv.UnsafeBytesToStr(bytes))
	}
	return ReadByteLines(file, cb)
}

//ReadByteLines 以 []byte 的格式按行读取
func ReadByteLines(file string, callback func(bytes []byte)) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		callback(scanner.Bytes())
	}
	return nil
}

func PutContents(file string, contents string) error {
	return putContents(file, []byte(contents), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0666), true)
}

func PutBytes(file string, contents []byte) error {
	return putContents(file, contents, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0666), true)
}

func PutContentsAppend(file string, contents string) error {
	return putContents(file, []byte(contents), os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.FileMode(0666), true)
}

func PutBytesAppend(file string, contents []byte) error {
	return putContents(file, contents, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.FileMode(0666), true)
}

func putContents(file string, contents []byte, flag int, perm os.FileMode, autoMkdir bool) error {

	if autoMkdir {
		dir := filepath.Dir(file)
		if !IsDir(dir) {
			if err := Mkdir(dir); err != nil {
				return err
			}
		}
	}

	f, err := os.OpenFile(file, flag, perm)
	if err != nil {
		return err
	}
	defer f.Close()

	if n, err := f.Write(contents); err != nil {
		return err
	} else if n != len(contents) {
		return io.ErrShortWrite
	}

	return nil
}
