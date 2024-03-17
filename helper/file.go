package helper

import (
	"encoding/base64"
	"path/filepath"
	"strings"
)

var extensions = map[string]string{
	"R0lGODdh":    ".gif",
	"R0lGODlh":    ".gif",
	"iVBORw0KGgo": ".png",
	"/9j/":        ".jpg",
}

type File struct {
	Name      string
	DataUri   string
	PlainText string
}

func (file *File) Decode() ([]byte, error) {
	return base64.StdEncoding.DecodeString(file.PlainText)
}

func (file *File) GetMime() string {
	parts := strings.Split(file.DataUri, ";")
	return parts[0][5:]
}

func (file *File) GetExt() string {
	for key, value := range extensions {
		if strings.Index(file.PlainText, key) == 0 {
			return value
		}
	}
	return ""
}

func NewFile(name string, dataUri string) *File {
	return &File{
		Name:      filepath.Base(name),
		DataUri:   dataUri,
		PlainText: dataUri[strings.IndexByte(dataUri, ',')+1:],
	}
}
