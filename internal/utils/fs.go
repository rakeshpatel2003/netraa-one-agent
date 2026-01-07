package utils

import (
	"io"
	"os"
	"path/filepath"
)

func CopyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)

		if info.IsDir() {
			return os.MkdirAll(target, info.Mode())
		}

		from, err := os.Open(path)
		if err != nil {
			return err
		}
		defer from.Close()

		to, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer to.Close()

		_, err = io.Copy(to, from)
		return err
	})
}

