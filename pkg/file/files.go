package file

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// Pwd get the current directory
func Pwd() (string, error) {
	return os.Getwd()
}

// MakeSureFile create if the file does not exist
func MakeSureFile(f string) error {
	err := FileIsOK(f)
	if err != nil {
		if err != nil {
			err := os.MkdirAll(filepath.Dir(f), fs.ModePerm)
			log.Printf("mkdir %s ,%v", filepath.Dir(f), err)
			if err != nil {
				return err
			}
		}
		_, err = os.Create(f)
		if err != nil {
			log.Printf("create file err: %v", err)
			return err
		}
	}
	return nil
}

// FileIsOK exists and is a file
func FileIsOK(f string) error {
	if f == "" {
		return fmt.Errorf("file path must be not empty")
	}
	s, err := os.Stat(f)
	if err != nil {
		log.Printf("file %s not exists", f)
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("%s is a directory", f)
	}
	return nil
}

// DirIsOK exists and is a folder
func DirIsOK(d string) error {
	if d == "" {
		return fmt.Errorf("dir path must be not empty")
	}
	s, err := os.Stat(d)
	if err != nil {
		log.Printf("dir %s not exists", d)
		return err
	}
	if !s.IsDir() {
		return fmt.Errorf("%s is not a directory", d)
	}
	return nil
}
