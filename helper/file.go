package helper

import (
	"encoding/base64"
	"path/filepath"
	"strings"
)

type File struct {
	Name    string
	Content string
}

func (file *File) Decode() ([]byte, error) {
	return base64.StdEncoding.DecodeString(file.Content[strings.IndexByte(file.Content, ',')+1:])
}

func (file *File) GetMime() string {
	parts := strings.Split(file.Content, ";")
	return parts[0][5:]
}

func (file *File) GetExt() string {
	return filepath.Ext(file.Name)
}

func NewFile(name string, content string) *File {
	return &File{
		Name:    filepath.Base(name),
		Content: content,
	}
}
