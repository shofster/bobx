package zip

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

/*

  File:    gzip.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*
  Description: compress / expand a single file using gzip algorithm.
*/

type GZipper struct {
	gzFile string
	reader *gzip.Reader
	target *os.File
}

// NewUnGZipper - expand .gzip file into a local file
func NewUnGZipper(gzFile string) (*GZipper, error) {
	f, err := os.Open(gzFile)
	if err != nil {
		return nil, err
	}
	r, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	z := GZipper{
		gzFile: gzFile,
		reader: r,
	}
	return &z, nil
}

// Extract - un-compress the contents of a gzipped file to a file
func (z *GZipper) Extract(dest string) error {

	if e := os.MkdirAll(dest, os.ModePerm); e != nil {
		return e
	}

	defer func() {
		_ = z.reader.Close()
	}()

	modTime := z.reader.Header.ModTime
	name := z.reader.Header.Name

	destination := filepath.Join(dest, name)
	out, e := os.Create(destination)
	if e != nil {
		log.Println("GZIP Extract create file error:", e)
		return e
	}
	defer func() {
		_ = out.Close()
		if err := z.reader.Close(); err != nil {
			log.Println("GZIP Extract reader close:", err)
		}
		if err := os.Chtimes(destination, modTime, modTime); err != nil {
			log.Println("GZIP Extract change times error:", err)
		}
	}()
	b := make([]byte, 4096)
	n := 0
	for {
		nb, ee := z.reader.Read(b)
		if ee == io.EOF {
			break
		}
		n += nb
		nw, _ := out.Write(b[:nb])
		if nb != nw {
			break
		}
	}
	return nil
}

// NewGZipper generates struct to create a .GZIP file
func NewGZipper(gzFile string) (*GZipper, error) {
	target, err := os.Create(gzFile)
	if err != nil {
		return nil, err
	}
	z := GZipper{
		gzFile: gzFile,
		target: target,
	}
	return &z, nil
}

// Compress adds 1 file to the gzip file
func (z *GZipper) Compress(path string) error {
	archive := gzip.NewWriter(z.target)
	defer func() {
		if e := archive.Close(); e != nil {
			log.Println("GZIP Compress archive close:", e)
		}
		if e := z.target.Close(); e != nil {
			log.Println("GZIP Compress target close:", e)
		}
	}()

	archive.Header.Name = filepath.Base(path)
	f, err := os.Open(path)
	if err != nil {
		log.Println("GZIP Compress error:", err)
		return err
	}
	archive.Header.Comment = strings.ReplaceAll(z.gzFile, "\\", "/")
	switch runtime.GOOS {
	case "windows":
		archive.Header.OS = 0
	case "linux", "darwin", "freebsd", "netbsd", "openbsd", "solaris":
		archive.Header.OS = 3
	default:
		archive.Header.OS = 11
	}
	loc, _ := time.LoadLocation("Local")
	if info, err := f.Stat(); err == nil {
		archive.Header.ModTime = info.ModTime().In(loc)
	}

	b := make([]byte, 4096)
	for {
		nb, err := f.Read(b)
		if err != nil {
			return err
		}
		_, err = archive.Write(b[:nb])
		if err != nil {
			return err
		}
		if nb < 4096 {
			break
		}
	}
	return err
}
