package NoteSystem

import (
	"encoding/json"
	"time"

	"io/ioutil"
	"os"

	"github.com/wallnutkraken/NoteSystem/Constant"
	"github.com/wallnutkraken/NoteSystem/Helpers"
)

type metadataNote struct {
	owner        string
	lastModified time.Time
	created      time.Time
}

func readNoteMetadata(path string) (metadataNote, error) {
	meta := metadataNote{}
	file, err := os.Open(path)
	if err != nil {
		return meta, err
	}
	json.NewDecoder(file).Decode(meta)
	return meta, nil
}

type internalNote struct {
	metadata metadataNote
	content  []rune
	file     *os.File
	path     string
}

// Note represents a note in the filesystem, it also contains all the relevant metadata for a file
type Note interface {
	File() *os.File
	Content() []rune
	Owner() string
	LastModified() time.Time
	Created() time.Time
	MetadataJSON(file *os.File) error
	Children() ([]string, error)
	PathTo() string
	GetChild(name string) (Note, error)
}

func (m internalNote) File() *os.File {
	return m.file
}

func (m internalNote) Content() []rune {
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

// GetChild returns the Note object for the first child with the given name
func (m internalNote) GetChild(name string) (Note, error) {
	newNote := internalNote{path: Helpers.PathBuilder(m.path, name)}
	fileBytes, err := ioutil.ReadFile(Helpers.PathBuilder(m.PathTo(), name, Constant.DataFilename))
	if err != nil {
		return nil, err
	}
	newNote.content = []rune(string(fileBytes))
	newNote.file, err = os.Open(Helpers.PathBuilder(newNote.path, Constant.DataFilename))
	if err != nil {
		return nil, err
	}
	metafile, err := readNoteMetadata(Helpers.PathBuilder(newNote.path, Constant.MetadataFile))
	if err != nil {
		return nil, err
	}
	newNote.metadata = metafile
	return newNote, nil
}
