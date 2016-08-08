package NoteSystem

import (
	"os"

	"github.com/wallnutkraken/NoteSystem/Constant"
	"github.com/wallnutkraken/NoteSystem/Helpers"
)

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
