package remote

import "os"

type IFileTransProtocol interface {
	Connect() error
	Upload(file *os.File, filePath, fileName string) error
	Mkdir(path string) error
	Disconnect() error
	Delete(filePath string) error
	Ls(path string) ([]string, error)
}
