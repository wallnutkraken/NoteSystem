package NoteSystem

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/wallnutkraken/NoteSystem/Constant"
)

var (
	logfile *os.File
	logger  *log.Logger
)

func initLog() {
	logname := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_logfile.txt"
	if !fileExists(Constant.LogDir) {
		err := os.Mkdir(Constant.LogDir, os.ModeDir)
		if err != nil {
			panic(fmt.Sprintln("Could not create log directory:", err))
		}
	}
	logfile, err := os.Create(Constant.LogDir + "/" + logname)
	if err != nil {
		panic(fmt.Sprintln("Could not create file:", "\""+logname+"\"", err))
	}
	logger = log.New(logfile, "[NoteSystem]", log.Lmicroseconds)
}

func init() {
	initLog()
	/* Check if a NoteSystem data dir exists */
	if _, err := os.Stat(Constant.DataPath); err != nil {
		if os.IsNotExist(err) {
			err = createFilesystem()
			if err != nil {
				panic(fmt.Sprintln("Could not create filesystem:", err))
			}
		}
	}
}

func createFilesystem() error {
	if err := os.Mkdir(Constant.DataPath, os.ModeDir); err != nil {
		return err
	}

	return nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
