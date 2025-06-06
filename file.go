package lrace

import (
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (int64, error) {
	src, err := os.Open(dstName)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.OpenFile(srcName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// FileExist 判断文件存在
func FileExist(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

func IsDir(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func IsFile(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}

	return !info.IsDir()
}
