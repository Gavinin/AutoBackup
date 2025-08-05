package archive

import (
	"AutoBuckup/internal/config"
	"AutoBuckup/internal/enum"
	"AutoBuckup/internal/log"
	"AutoBuckup/internal/util"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ZipStore struct{}

func NewZipStore() *ZipStore {
	return &ZipStore{}
}

func (z *ZipStore) BatchArchive(archive config.Archive, paths []string) ([]string, error) {
	results := make([]string, 0)

	for _, path := range paths {

		zipPath, err := ZipStoreFolder(archive, path)
		if err != nil {
			log.Logger.Error(err)
		} else {
			results = append(results, zipPath)
		}
	}

	return results, nil
}

func ZipStoreFolder(archive config.Archive, path string) (string, error) {
	_, folderName := util.SeparatePath(path)
	if archive.SortByDate {
		folderName = util.GetFolderName(archive.NameFormat, folderName)
	}
	filePathStr := fmt.Sprintf("%s/%s%s", archive.TmpFilePath, folderName, util.GetExt(enum.TypeZipStore))
	filePath := filepath.Join(filePathStr)
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	zw := zip.NewWriter(f)
	defer zw.Close()

	err = filepath.Walk(path, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 目录不需要写入
		if fi.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(filepath.Dir(path), fileName)
		if err != nil {
			return err
		}
		relPath = strings.ReplaceAll(relPath, "\\", "/") // 兼容windows

		header, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}
		header.Name = relPath
		header.Method = zip.Store // 仅存储不压缩

		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}

		fr, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer fr.Close()

		n, err := io.Copy(writer, fr)
		if err != nil {
			return err
		}

		log.Logger.Debug("Success save %s ，Write: %d bytes\n", fileName, n)
		return nil
	})

	return filePath, err
}
