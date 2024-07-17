package zip

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"
)

const gzipData = "../data/gzipData/generic.tar"

func TestGzip(t *testing.T) {

	tempDir, err := os.MkdirTemp("", "gzip_")
	if err != nil {
		t.Error(err)
	}
	defer func(path string) {
		if err := os.RemoveAll(path); err != nil {
			log.Println("RemoveAll failed", err)
		}
	}(tempDir)

	gzipFile := filepath.Join(tempDir, filepath.Base(gzipData)+".gzip")
	t.Log("GZIP file:", gzipFile)

	gzipControl, err := NewGZipper(gzipFile)
	if err != nil {
		t.Error(err)
	}
	err = gzipControl.Compress(gzipData)
	if err != nil {
		t.Error(err)
	}

	unGzip, err := NewUnGZipper(gzipFile)
	err = unGzip.Extract(tempDir)

	dstFiles, err := pathList(tempDir)
	if err != nil {
		t.Error(err)
	}
	if len(dstFiles) != 2 {
		t.Error(errors.New("mismatch of files count"))
	}

}
