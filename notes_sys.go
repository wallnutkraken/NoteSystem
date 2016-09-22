package NoteSystem

import (
	"log"
	"os"
)

type noteSys struct {
	absPath string
	logger  *log.Logger
}

func (n *noteSys) SetLogger(logger *log.Logger) {
	n.logger = logger
}

func (n *noteSys) Create() error {
	panic("to be implemented")
}

// NoteSys represents an interface to the underlying filesystem of notes.
type NoteSys interface {
	SetLogger(*log.Logger)
	Create() error
}

// New interface to the NoteSystem. Must be instanciated first even if a NoteSystem already exists
// on disk. To create a new one, call Create() \ on it.
func New(logpath, nsRoot string) (NoteSys, error) {
	notes := new(noteSys)
	var logFile *os.File
	var err error

	if _, err = os.Stat(logpath); err.Error() != os.ErrNotExist.Error() {
		/* Exists */
		logFile, err = os.Open(logpath)
	} else {
		logFile, err = os.Create(logpath)
	}
	if err != nil {
		return nil, err
	}
	notes.logger = log.New(logFile, "[NoteSys]", log.Ltime)
	notes.logger.Println("Started log")

	return notes, nil
}
