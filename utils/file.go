package utils

import (
	"bufio"
	"encoding/json"
	"fireboom-migrate/consts"
	"fireboom-migrate/types/origin"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func DelDir(src string) {
	os.RemoveAll(src)
}

func CutDir(src string, dst string) {
	CopyDir(src, dst)
	DelDir(src)
}

func MoveDir() {

	// 移动store目录到old目录
	CutDir(consts.StorePath, filepath.Join(consts.BackendPath, consts.StorePath))

	// 移动upload目录到old目录
	CutDir(consts.UploadPath, filepath.Join(consts.BackendPath, consts.UploadPath))

	// 移动exported目录到old目录
	CutDir(consts.ExportedPath, filepath.Join(consts.BackendPath, consts.ExportedPath))

	os.Rename(consts.StoreCloudPath, consts.StorePath)
	os.Rename(consts.UploadCloudPath, consts.UploadPath)

}

func RollBack() {

	DelDir(consts.StorePath)
	DelDir(consts.UploadPath)

	CutDir(filepath.Join(consts.BackendPath, consts.StorePath), consts.StorePath)

	CutDir(filepath.Join(consts.BackendPath, consts.UploadPath), consts.UploadPath)

	CutDir(filepath.Join(consts.BackendPath, consts.ExportedPath), consts.ExportedPath)

	os.Remove(".env")
	CopyFile(consts.BackendPath+consts.PathSep+".env", ".env")

	DelDir(consts.BackendPath)
}

func RenameCustom() {
	fmt.Println("输入需要修改的目录名称(例如：custom-go 回车换行可输入多个，两次回车结束输入；若无需修改，直接回车)")
	reader := bufio.NewReader(os.Stdin)
	var dirs []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		dirs = append(dirs, line)
	}

	for _, dir := range dirs {
		dir = strings.Trim(dir, "\n")
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			fmt.Println(dir + " is not exist")
		}
		if err == nil {
			_ = os.Rename(dir+consts.PathSep+consts.Auth, dir+consts.PathSep+consts.Authentication)
			_ = os.Rename(dir+consts.PathSep+consts.Proxys, dir+consts.PathSep+consts.Proxy)
			_ = os.Rename(dir+consts.PathSep+consts.Hooks, dir+consts.PathSep+consts.Operation)
			fmt.Println("renaming " + dir)
		}
	}
	fmt.Println("查看本次修改内容：https://github.com/fireboomio/fb-migration")
}
