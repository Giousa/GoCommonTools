/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/15
 */
package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ZipDir(dir, zipFile string) error{

	fz, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer fz.Close()

	w := zip.NewWriter(fz)
	defer w.Close()

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			//fDest, err := w.Create(path[len(dir)+1:])
			fDest, err := w.Create(path[len(dir):])
			if err != nil {
				return err
			}
			fSrc, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fSrc.Close()
			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				return err
			}
		}
		return nil
	})

	fmt.Println("压缩成功")
	return nil
}

