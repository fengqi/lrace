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
