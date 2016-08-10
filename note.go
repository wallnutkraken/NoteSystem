package NoteSystem

import (
	"encoding/json"
	"time"

	"io/ioutil"
	"os"
)

type metadataNote struct {
	owner        string
	lastModified time.Time
	created      time.Time
}

type internalNote struct {
	metadata metadataNote
	content  string
	file     *os.File
	path     string
}

// Note represents a note in the filesystem, it also contains all the relevant metadata for a file
type Note interface {
	File() *os.File
	Content() string
	Owner() string
	LastModified() time.Time
	Created() time.Time
	MetadataJSON(file *os.File) error
	Children() ([]string, error)
	PathTo() string
}

func (m internalNote) File() *os.File {
	return m.file
}

func (m internalNote) Content() string {
	return m.content
}

func (m internalNote) Owner() string {
	return m.metadata.owner
}

func (m internalNote) LastModified() time.Time {
	return m.metadata.lastModified
}

func (m internalNote) Created() time.Time {
	return m.metadata.created
}

func (m internalNote) MetadataJSON(file *os.File) error {
	return json.NewEncoder(file).Encode(m.metadata)
}

func (m internalNote) Children() ([]string, error) {
	fileInfo, err := ioutil.ReadDir(m.path)
	if err != nil {
		return nil, err
	}
	var children = make([]string, 0)
	for _, file := range fileInfo {
		if file.IsDir() && []rune(file.Name())[0] == '.' {
			children = append(children, file.Name())
		}
	}
	return children, nil
}

func (m internalNote) PathTo() string {
	return m.path
}
