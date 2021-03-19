//nolint:dupl Similar Methods between entities
package clients

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
)

// FileClient is client for operate files
type FileClient struct {
	Config
}

// GetGeneric returns the generic info
func (q *FileClient) GetFile(nameFile string) (*File, error) {
	log.Debug().Msgf("[DEBUG] custom::GetFile - file name %s ",nameFile)
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s%s", q.GetConfigCustom().LocationFiles, nameFile))
	if err != nil {
		return nil, err
	}

	return &File{
		Name:     nameFile,
		Content:  string(dat),
	}, nil

}

// CreateFile returns the generic info
func (q *FileClient) CreateFile(fileInput *File) (*File, error) {
	log.Debug().Msgf("[DEBUG] custom::CreateFile - file name %+v ",&fileInput)
	d1 := []byte(fileInput.Content)
	err := ioutil.WriteFile(fmt.Sprintf("%s%s", q.GetConfigCustom().LocationFiles, fileInput.Name), d1, 0777)
	if err != nil {
		log.Debug().Msgf("[DEBUG] error %+v ",err)
		return nil, err
	}
	return fileInput, nil
}

// DeleteGeneric returns the generic info
func (q *FileClient) DeleteFile(fileName string) (bool, error) {
	err := os.Remove(fmt.Sprintf("%s%s", q.GetConfigCustom().LocationFiles, fileName))
	if err != nil {
		return false, err
	}
	return true, nil
}

type File struct {
	Name          string `json:"name" `
	Content       string `json:"content"`
}
