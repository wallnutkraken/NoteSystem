package NoteSystem

import (
	"encoding/json"
	"os"

	"github.com/wallnutkraken/NoteSystem/Constant"
)

type metadataFile struct {
	UsersDir string
}

func writeStruct(filePath string, content interface{}) error {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		logger.Println("[E]:", err)
		return err
	}
	encoder := json.NewEncoder(file)
	return encoder.Encode(content)
}

func readMetadataFile() (*metadataFile, error) {
	file, err := os.Open(Constant.DataPath + "/" + Constant.MetadataFile)
	if err != nil {
		logger.Println("[E]:", err)
		return nil, err
	}
	decoder := json.NewDecoder(file)
	meta := &metadataFile{}
	err = decoder.Decode(meta)
	file.Close()
	return meta, err
}
