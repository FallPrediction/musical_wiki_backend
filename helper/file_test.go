package helper_test

import (
	"github.com/stretchr/testify/assert"
	. "musical_wiki/helper"
	"testing"
)

func TestGetExt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name      string
		fileName  string
		dataUri   string
		expectExt string
	}{
		{
			"png",
			"png",
			"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAEnQAABJ0Ad5mH3gAAAANSURBVBhXY3gro/IfAAVUAi3GPZKdAAAAAElFTkSuQmCC",
			".png",
		},
		{
			"gif1",
			"gif1",
			"data:image/gif;base64,R0lGODlhAQABAPAAAO0cJAAAACH/C05FVFNDQVBFMi4wAwEAAAAh/wtJbWFnZU1hZ2ljaw1nYW1tYT0wLjQ1NDU1ACH5BAAUAP8ALAAAAAABAAEAAAICRAEAOw==",
			".gif",
		},
		{
			"gif2",
			"gif2",
			"data:image/gif;base64,R0lGODdhAQABAPAAAO0cJAAAACH/C05FVFNDQVBFMi4wAwEAAAAh/wtJbWFnZU1hZ2ljaw1nYW1tYT0wLjQ1NDU1ACH5BAAUAP8ALAAAAAABAAEAAAICRAEAOw==",
			".gif",
		},
		{
			"jpg",
			"jpg",
			"data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAeAB4AAD/4QAiRXhpZgAATU0AKgAAAAgAAQESAAMAAAABAAEAAAAAAAD/2wBDAAIBAQIBAQICAgICAgICAwUDAwMDAwYEBAMFBwYHBwcGBwcICQsJCAgKCAcHCg0KCgsMDAwMBwkODw0MDgsMDAz/2wBDAQICAgMDAwYDAwYMCAcIDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAz/wAARCAABAAEDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD+f+iiigD/2Q==",
			".jpg",
		},
		{
			"not acceptable file type",
			"html",
			"data:text/html;base64,PGh0bWw+PC9odG1sPg==",
			"",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			file := NewFile(test.fileName, test.dataUri)
			assert.Equal(file.GetExt(), test.expectExt)
		})
	}
}

func TestGetMime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name       string
		fileName   string
		dataUri    string
		expectMime string
	}{
		{
			"png",
			"png",
			"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAEnQAABJ0Ad5mH3gAAAANSURBVBhXY3gro/IfAAVUAi3GPZKdAAAAAElFTkSuQmCC",
			"image/png",
		},
		{
			"gif1",
			"gif1",
			"data:image/gif;base64,R0lGODlhAQABAPAAAO0cJAAAACH/C05FVFNDQVBFMi4wAwEAAAAh/wtJbWFnZU1hZ2ljaw1nYW1tYT0wLjQ1NDU1ACH5BAAUAP8ALAAAAAABAAEAAAICRAEAOw==",
			"image/gif",
		},
		{
			"gif2",
			"gif2",
			"data:image/gif;base64,R0lGODdhAQABAPAAAO0cJAAAACH/C05FVFNDQVBFMi4wAwEAAAAh/wtJbWFnZU1hZ2ljaw1nYW1tYT0wLjQ1NDU1ACH5BAAUAP8ALAAAAAABAAEAAAICRAEAOw==",
			"image/gif",
		},
		{
			"jpg",
			"jpg",
			"data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAeAB4AAD/4QAiRXhpZgAATU0AKgAAAAgAAQESAAMAAAABAAEAAAAAAAD/2wBDAAIBAQIBAQICAgICAgICAwUDAwMDAwYEBAMFBwYHBwcGBwcICQsJCAgKCAcHCg0KCgsMDAwMBwkODw0MDgsMDAz/2wBDAQICAgMDAwYDAwYMCAcIDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAz/wAARCAABAAEDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD+f+iiigD/2Q==",
			"image/jpeg",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			file := NewFile(test.fileName, test.dataUri)
			assert.Equal(file.GetMime(), test.expectMime)
		})
	}
}

func TestDecode(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name     string
		fileName string
		dataUri  string
		expect   []byte
	}{
		{
			"png",
			"png",
			"data:image/png;base64,iVBORw0KGgoAAA==",
			[]byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0},
		},
		{
			"gif1",
			"gif1",
			"data:image/gif;base64,R0lGODlhAQABAA==",
			[]byte{71, 73, 70, 56, 57, 97, 1, 0, 1, 0},
		},
		{
			"gif2",
			"gif2",
			"data:image/gif;base64,R0lGODdhAQABAA==",
			[]byte{71, 73, 70, 56, 55, 97, 1, 0, 1, 0},
		},
		{
			"jpg",
			"jpg",
			"data:image/jpeg;base64,/9j/4AAQSkZJRg==",
			[]byte{255, 216, 255, 224, 0, 16, 74, 70, 73, 70},
		},
		{
			"html",
			"html",
			"data:text/html;base64,PGh0bWw+PC9odA==",
			[]byte{60, 104, 116, 109, 108, 62, 60, 47, 104, 116},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			file := NewFile(test.fileName, test.dataUri)
			decoded, _ := file.Decode()
			assert.Equal(decoded, test.expect)
		})
	}
}
