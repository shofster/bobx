package misc

import (
	"log"
	"testing"
)

func TestSampleScanner(t *testing.T) {
	file := "../data/555Sample.ged"
	scanner, err := NewScanner(file, UTF8)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		t.Log(scanner.Text())
	}
	if err := scanner.Close(); err != nil {
		t.Errorf("FAIL: Unable to close file handle after scan")
	}
}
func TestSample16LEScanner(t *testing.T) {
	file := "../data/555Sample16LE.ged"
	scanner, err := NewScanner(file, UTF8)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		t.Log(scanner.Text())
	}
	if err := scanner.Close(); err != nil {
		t.Errorf("FAIL: Unable to close file handle after scan")
	}
}
func TestSample16BEScanner(t *testing.T) {
	file := "../data/555Sample16BE.ged"
	scanner, err := NewScanner(file, UTF8)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		t.Log(scanner.Text())
	}
	if err := scanner.Close(); err != nil {
		t.Errorf("FAIL: Unable to close file handle after scan")
	}
}
