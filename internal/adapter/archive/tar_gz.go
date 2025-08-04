package archive

import (
	"AutoBuckupG/internal/config"
	"AutoBuckupG/internal/log"
	"AutoBuckupG/internal/util"
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	time "time"
)

type TagGz struct {
}

func NewTagGz() *TagGz {
	return &TagGz{}
}

func (t TagGz) BatchArchive(archive config.Archive, paths []string) ([]string, error) {
	results := make([]string, 0)

	for _, path := range paths {
		if path == "" {
			continue
		}
		path, err := TagGzFolder(archive, path)
		if err != nil {
			log.Logger.Error(err)
		} else {
			results = append(results, path)
		}
	}

	return results, nil
}

func TagGzFolder(archive config.Archive, path string) (string, error) {
	_, folderName := util.SeparatePath(path)
	if archive.SortByDate {
		timeFormatStr := time.Now().Format(util.NameFormat2DateFormat(archive.NameFormat))
		folderName = fmt.Sprintf("%s_%s", folderName, timeFormatStr)
	}
	filePathStr := fmt.Sprintf("%s/%s.tar.gz", archive.TmpFilePath, folderName)
	filePath := filepath.Join(filePathStr)
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// create gzip writer
	gw := gzip.NewWriter(f)
	defer gw.Close()

	// create tar writer
	tw := tar.NewWriter(gw)
	defer tw.Close()

	err = filepath.Walk(path, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}

		hdr.Name = strings.TrimPrefix(fileName, string(filepath.Separator))

		// 写入文件信息
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		// open file
		fr, err := os.Open(fileName)
		defer fr.Close()
		if err != nil {
			return err
		}

		n, err := io.Copy(tw, fr)
		if err != nil {
			return err
		}

		log.Logger.Debug("Success save %s ，Write: %d bytes\n", fileName, n)

		return nil
	})

	return filePath, err
}
