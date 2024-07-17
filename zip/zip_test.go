package zip

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const zipData = "../data/zipData"

func TestZip(t *testing.T) {

	tempDir, err := os.MkdirTemp("", "zip_")
	if err != nil {
		t.Error(err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println("RemoveAll failed", err)
		}
	}(tempDir)

	zipFile := filepath.Join(tempDir, "test.zip")
	t.Log("ZIP file:", zipFile)
	srcFiles, err := pathList(zipData)
	if err != nil {
		t.Error(err)
	}
	for _, f := range srcFiles {
		_, e := os.Lstat(f)
		if e != nil {
			t.Errorf("Unable to stat: %s, %s", filepath.Base(f), e)
		}
	}

	zipper, err := NewZip(zipFile)
	if err != nil {
		t.Error(err)
	}
	err = zipper.Compress(zipData, srcFiles, func(s string) {
		t.Log("adding:", s)
	})
	if err != nil {
		t.Error(err)
	}

	unZipper, err := NewUnZip(zipFile)
	err = unZipper.Extract(tempDir, func(s string) {
		t.Log("extracting:", s)
	})

	// check the results
	dstFiles, err := pathList(tempDir)
	if err != nil {
		t.Error(err)
	}
	if len(srcFiles) != len(dstFiles)-1 {
		t.Error(errors.New("mismatch of files count"))
	}

}
