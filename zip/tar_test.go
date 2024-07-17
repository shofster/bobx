package zip

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"
)

const tarData = "../data/tarData"

func TestTar(t *testing.T) {

	tempDir, err := os.MkdirTemp("", "tar_")
	if err != nil {
		t.Error(err)
	}
	defer func(path string) {
		if err := os.RemoveAll(path); err != nil {
			log.Println("RemoveAll failed", err)
		}
	}(tempDir)

	tarFile := filepath.Join(tempDir, "test.tar")
	t.Log("TAR file:", tarFile)
	srcFiles, err := pathList(tarData)
	if err != nil {
		t.Error(err)
	}
	for _, f := range srcFiles {
		_, e := os.Lstat(f)
		if e != nil {
			t.Errorf("Unable to stat: %s, %s", filepath.Base(f), e)
		}
	}

	tarControl, err := NewTar(tarFile)
	if err != nil {
		t.Error(err)
	}
	err = tarControl.Compress(tarData, srcFiles, func(s string) {
		t.Log("adding:", s)
	})
	if err != nil {
		t.Error(err)
	}

	unTar, err := NewUnTar(tarFile)
	_, err = unTar.Extract(tempDir, func(s string) {
		t.Log("extracting:", s)
	})

	dstFiles, err := pathList(tempDir)
	if err != nil {
		t.Error(err)
	}
	if len(srcFiles) != len(dstFiles)-1 {
		t.Error(errors.New("mismatch of files count"))
	}

}
