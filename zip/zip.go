package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*

  File:    zip.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*
  Description: compress / extract a .ZIP file from list of files
*/

type Zipper struct {
	zipFile string
	reader  *zip.ReadCloser
	target  *os.File
}

// NewUnZipper generates struct to extract entries from  a .ZIP file
func NewUnZip(zipFile string) (*Zipper, error) {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil, err
	}
	z := Zipper{
		zipFile: zipFile,
		reader:  r,
	}
	return &z, nil
}

// Extract - copies the contents of a zip file to a directory
func (z *Zipper) Extract(dest string, processing func(string)) error {

	if e := os.MkdirAll(dest, os.ModePerm); e != nil {
		return e
	}

	defer func() {
		if e := z.reader.Close(); e != nil {
			log.Println("ZIP Extract reader close:", e)
		}
		z.reader = nil
	}()

	for _, file := range z.reader.File {
		if processing != nil {
			processing(file.Name)
		}
		isDir := file.FileInfo().IsDir()
		if isDir {
			dir := filepath.Dir(file.Name)
			path := filepath.Join(dest, dir)
			e := os.MkdirAll(path, os.ModePerm)
			if e != nil {
				return e
			}
			continue
		}
		// parent of this file
		parent := filepath.Dir(file.Name)
		if parent != "" && parent != "." {
			path := filepath.Join(dest, parent)
			em := os.MkdirAll(path, os.ModePerm)
			if em != nil {
				return em
			}
		}
		destination := filepath.Join(dest, file.Name)
		fmt.Println("dest", destination)
		out, err := os.Create(destination)
		if err != nil {
			return err
		}
		fr, err := file.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(out, fr)
		_ = out.Close()
		if err != nil {
			log.Println("ZIP Extract copy error:", err)
			return err
		}
		if err = os.Chtimes(destination, file.Modified, file.Modified); err != nil {
			log.Println("ZIP Extract change times error:", err)
		}
	}
	return nil
}

// NewZipper generates struct to create a .ZIP file
func NewZip(zipFile string) (*Zipper, error) {
	target, err := os.Create(zipFile)
	if err != nil {
		return nil, err
	}
	z := Zipper{
		zipFile: zipFile,
		target:  target,
	}
	return &z, nil
}

// Compress adds all directories & files to the zip file
func (z *Zipper) Compress(parent string, files []string, processing func(string)) error {
	if len(parent) > 1 {
		parent += "/"
	}
	archive := zip.NewWriter(z.target)
	defer func() {
		if e := archive.Close(); e != nil {
			log.Println("ZIP Compress archive close:", e)
		}
		if e := z.target.Close(); e != nil {
			log.Println("ZIP Compress target close:", e)
		}
	}()
	for _, file := range files {
		err := addZipFile(archive, parent, file, processing)
		if err != nil {
			log.Println("ZIP Compress error:", err)
			return err
		}
	}
	return nil
}

// addZipFile adds a single (directory or file) to the zip file
func addZipFile(zipWriter *zip.Writer, parent string, file string, processing func(string)) error {
	if processing != nil {
		processing(filepath.Base(file))
	}
	fileToZip, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func() {
		_ = fileToZip.Close()
	}()
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}
	header, _ := zip.FileInfoHeader(info)
	path := file[len(parent):]
	if info.IsDir() {
		path += "/"
	}
	header.Name = strings.ReplaceAll(path, "\\", "/")
	loc, _ := time.LoadLocation("Local")
	header.Modified = header.Modified.In(loc)
	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
