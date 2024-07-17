package zip

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*

  File:    tar.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*
  Description: collect / extract a .TAR file from list of files
*/

type Tar struct {
	tarFile string
	reader  *tar.Reader
	target  *os.File
	writer  *tar.Writer
}

// NewUnTar - expand .tar file into individual files
func NewUnTar(tarFile string) (*Tar, error) {
	f, err := os.Open(tarFile)
	if err != nil {
		return nil, err
	}
	r := tar.NewReader(f)
	z := Tar{
		tarFile: tarFile,
		target:  f,
		reader:  r,
	}
	return &z, nil
}

// Extract - copies the contents of a zip file to a directory
func (z *Tar) Extract(dest string, processing func(string)) (int, error) {

	if e := os.MkdirAll(dest, os.ModePerm); e != nil {
		return 0, e
	}

	var out *os.File
	var closeout = func() {
		if out != nil {
			if e := out.Close(); e != nil {
				log.Println("TAR Extract out close:", e)
			}
			out = nil
		}
	}
	defer func() {
		closeout()
		z.reader = nil
		_ = z.target.Close()
	}()

	count := 0
	for {
		header, err := z.reader.Next()
		if err == io.EOF {
			break
		}
		processing(header.Name)
		if err != nil {
			return count, err
		}
		destination := filepath.Join(dest, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			err = os.MkdirAll(destination, os.ModePerm)
			if err != nil {
				log.Println("TAR Extract create path error:", err)
				return count, err
			}
		case tar.TypeReg:
			err = os.MkdirAll(filepath.Dir(destination), os.ModePerm)
			out, err = os.Create(destination)
			if err != nil {
				log.Println("TAR Extract create file error:", err)
				out = nil
				return count, err
			}
			_, err = io.Copy(out, z.reader)
			if err != nil {
				log.Println("TAR Extract copy error:", err)
				return count, err
			}
			if err = os.Chtimes(destination, header.ModTime, header.ModTime); err != nil {
				log.Println("TAR Extract change times error:", err)
			}
			closeout()
			count++
		}
	}
	return count, nil
}

// NewTar generates struct to create a .TAR file
func NewTar(tarFile string) (*Tar, error) {
	target, err := os.Create(tarFile)
	if err != nil {
		return nil, err
	}
	archive := tar.NewWriter(target)
	z := Tar{
		tarFile: tarFile,
		target:  target,
		writer:  archive,
	}
	return &z, nil
}

// Compress adds all directories & files to the tar file
func (z *Tar) Compress(parent string, files []string, processing func(string)) error {
	if len(parent) > 1 {
		parent += "/"
	}
	defer func() {
		if e := z.target.Close(); e != nil {
			log.Println("TAR Compress target close:", e)
		}
	}()

	for _, file := range files {
		if processing != nil {
			processing(filepath.Base(file))
		}
		if err := addTarFile(z.writer, parent, file); err != nil {
			log.Println("TAR Compress error:", err)
			return err
		}
	}
	return nil
}

// addTarFile adds a single (directory or file) to the tar file
func addTarFile(tarWriter *tar.Writer, parent string, file string) error {
	// input file
	fileToTar, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func() {
		_ = fileToTar.Close()
	}()
	info, err := fileToTar.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	header := &tar.Header{
		Size:     info.Size(),
		Mode:     int64(info.Mode()),
		ModTime:  info.ModTime(),
		Typeflag: tar.TypeReg,
	}
	if info.IsDir() {
		header.Typeflag = tar.TypeDir
	}
	path := file[len(parent):]
	header.Name = strings.ReplaceAll(path, "\\", "/")
	loc, _ := time.LoadLocation("Local")
	header.ModTime = header.ModTime.In(loc)
	err = tarWriter.WriteHeader(header)
	if err == nil {
		_, err = io.Copy(tarWriter, fileToTar)
	}
	return nil
}
