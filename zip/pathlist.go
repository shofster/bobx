package zip

import (
	"os"
	"path/filepath"
)

/*

  File:    pathlist.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*
  Description: get a list of all files and sub-directories and files
*/

func pathList(root string) (fileList []string, err error) {

	entries, err := os.ReadDir(root)
	if err != nil {
		return
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			fileList = append(fileList, filepath.Join(root, entry.Name()))
		}
	}
	for _, entry := range entries {
		if entry.IsDir() {
			fileList = append(fileList, filepath.Join(root, entry.Name()))
			fileList, err = dirList(filepath.Join(root, entry.Name()), fileList)
		}
	}
	return
}

func dirList(dir string, fileList []string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err == nil {
		for _, entry := range entries {
			if !entry.IsDir() {
				fileList = append(fileList, filepath.Join(dir, entry.Name()))
			}
		}
		for _, entry := range entries {
			if entry.IsDir() {
				fileList = append(fileList, filepath.Join(dir, entry.Name()))
				fileList, err = dirList(filepath.Join(dir, entry.Name()), fileList)
			}
		}
	}
	return fileList, err
}
