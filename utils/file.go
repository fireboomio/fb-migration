package utils

import (
	"bufio"
	"encoding/json"
	"fireboom-migrate/types/origin"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ReadFile(path string) (content []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if content != nil {
			content = append(content, '\n')
		}
		content = append(content, fileScanner.Bytes()...)
	}
	return
}

// CopyFile 高效地拷贝文件，使用底层操作系统的零拷贝特性，不需要将整个文件的内容加载到内存中。
func CopyFile(srcPath, dstPath string) (err error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return
	}
	defer func() { _ = srcFile.Close() }()

	dstFile, err := CreateFile(dstPath)
	if err != nil {
		return
	}
	defer func() { _ = dstFile.Close() }()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return
	}

	err = dstFile.Sync()
	return
}

func CopyDir(srcDir, dstDir string) error {
	err := os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("failed to read source directory: %w", err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(dstDir, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return fmt.Errorf("failed to copy subdirectory: %w", err)
			}
		} else {
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return fmt.Errorf("failed to copy file: %w", err)
			}
		}
	}

	return nil
}

func CreateFile(path string) (file *os.File, err error) {
	if err = MkdirAll(filepath.Dir(path)); err != nil {
		return
	}

	return os.Create(path)
}

func MkdirAll(dirname string) error {
	return os.MkdirAll(dirname, os.ModePerm)
}

func WriteFile(path string, dataBytes []byte) error {
	if err := MkdirAll(filepath.Dir(path)); err != nil {
		return err
	}

	return os.WriteFile(path, dataBytes, 0644)
}

func NotExistFile(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// ReadDir 获取path下面所有的文件名称
func ReadDir(path string) (fileList []string, err error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	fileList = make([]string, 0)
	for _, entry := range dir {
		fileList = append(fileList, entry.Name())
	}

	return fileList, nil
}

func GetHookConfigEnable(path string) bool {
	content, err := ReadFile(path)
	if err != nil {
		panic(err)
	}

	hook := origin.HookStruct{}
	err = json.Unmarshal(content, &hook)
	if err != nil {
		panic(err)
	}

	if hook.Enabled == nil {
		return false
	}
	return *hook.Enabled
}
