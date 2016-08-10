package NoteSystem

import (
	"os"
	"time"

	"github.com/wallnutkraken/NoteSystem/Constant"
	"github.com/wallnutkraken/NoteSystem/Helpers"
)

type internalUser struct {
	username string
}

func (user internalUser) CreateNote(parent Note, noteName string) (Note, error) {
	var prefixPath string
	if parent == nil {
		prefixPath = "."
	} else {
		prefixPath = parent.PathTo()
	}

	/* TODO: Check if noteName is a legal name */

	pathToNewNote := Helpers.PathBuilder(prefixPath, noteName)
	/* Create dir */
	if err := os.Mkdir(pathToNewNote, os.ModeDir); err != nil {
		return nil, err
	}
	/* And all relevant files */
	contentFile, err := os.Create(Helpers.PathBuilder(pathToNewNote, Constant.DataFilename))
	if err != nil {
		return nil, err
	}
	metadataFile, err := os.Create(Helpers.PathBuilder(pathToNewNote, Constant.MetadataFile))
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()
	meta := internalNote{metadata: metadataNote{owner: user.username, created: currentTime, lastModified: currentTime},
		file: contentFile, path: pathToNewNote}
	meta.MetadataJSON(metadataFile)

	panic("Unfinished. Todo: Make 'Note' more usable in regards to the file. Concurrency and whatnot.")

	return meta, nil
}

// User is an interface for dealing with a specific user's notes/metadata
type User interface {
	FindNote(parent Note, noteName string) (Note, error)
	CreateNote(parent Note, noteName string) (Note, error)
}

func createUser(name string) error {
	userPath := Helpers.PathBuilder(Constant.DataPath, Constant.UsersDir, name)
	if err := os.Mkdir(userPath, os.ModeDir); err != nil {
		return err
	}
	if err := os.Mkdir(Helpers.PathBuilder(userPath, Constant.UsersNotesDir), os.ModeDir); err != nil {
		return err
	}
	return nil
}
