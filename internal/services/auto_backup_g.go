package services

import (
	"AutoBuckup/internal/adapter/remote"
	"AutoBuckup/internal/config"
	"AutoBuckup/internal/enum"
	"AutoBuckup/internal/log"
	"AutoBuckup/internal/util"
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"strings"
	"time"
)

func AutoBackup(cfg *config.Config) {
	if cfg.Debug {
		backupAndUpload(cfg)
		return
	}
	c := cron.New()
	_, err := c.AddFunc(cfg.Cron, func() {
		backupAndUpload(cfg)
	})
	if err != nil {
		return
	}
	c.Start()
}

func backupAndUpload(cfg *config.Config) {
	if len(cfg.Directory) == 0 {
		log.Logger.Error("No auto_backup directory specified")
		return
	}

	remoteProtocolClient, err := SelectRemoteProtocolClient(cfg)
	if err != nil {
		log.Logger.Error("No remote protocol client selected")
		return
	}

	archiveClient := SelectArchive(cfg)

	archive, err := archiveClient.BatchArchive(cfg.Archive, cfg.Directory)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	if len(archive) == 0 {
		log.Logger.Warn("No archive to upload")
		return
	}

	err = remoteProtocolClient.Connect()
	if err != nil {
		log.Logger.Error(err)
		return
	}
	defer remoteProtocolClient.Disconnect()

	currentDate := time.Now()
	ls, err := remoteProtocolClient.Ls(cfg.Remote.ArchivePath)
	if err != nil {
		log.Logger.Error(err)
	}
	for _, s := range archive {
		_, fileName := util.SeparatePath(s)
		err := remoteProtocolClient.Mkdir(cfg.Remote.ArchivePath)
		if err != nil {
			log.Logger.Error(err)
		}
		file, err := os.Open(s)
		if err != nil {
			log.Logger.Error(err)
			continue
		}

		if !cfg.Archive.SavePrevious {
			removePreviousFiles(cfg, ls, fileName, remoteProtocolClient, currentDate)
		}

		err = remoteProtocolClient.Upload(file, cfg.Remote.ArchivePath, fileName)
		if err != nil {
			log.Logger.Error(err)
		}
		file.Close()
		err = os.Remove(s)
		if err != nil {
			log.Logger.Error(err)
		}

	}

}

func removePreviousFiles(cfg *config.Config, ls []string, fileName string, client remote.IFileTransProtocol, now time.Time) {
	fileNameWithoutExt := strings.TrimSuffix(fileName, util.GetExt(cfg.Archive.Type))
	fileNameWithoutExtAndDate := fileNameWithoutExt
	if strings.Contains(fileNameWithoutExt, enum.FolderSeparateFlag) {
		fileNameWithoutExtAndDate = fileNameWithoutExtAndDate[:strings.LastIndex(fileNameWithoutExtAndDate, enum.FolderSeparateFlag)]
	}
	compareDate := now
	if cfg.Archive.StoreExpired > 0 {
		compareDate = now.AddDate(0, 0, -cfg.Archive.StoreExpired)
	}

	pendingRemove := make([]string, 0)

	for _, remoteFileName := range ls {
		if strings.LastIndex(remoteFileName, util.GetExt(cfg.Archive.Type)) != len(remoteFileName)-len(util.GetExt(cfg.Archive.Type)) {
			continue
		}
		remoteFileNameWithoutExt := strings.TrimSuffix(remoteFileName, util.GetExt(cfg.Archive.Type))

		if strings.Contains(remoteFileNameWithoutExt, enum.FolderSeparateFlag) {
			remoteFileDate := remoteFileNameWithoutExt[strings.LastIndex(remoteFileNameWithoutExt, enum.FolderSeparateFlag)+1:]
			if remoteFileDate != "" {
				remoteCreateAt, err := time.Parse(util.NameFormat2DateFormat(cfg.Archive.NameFormat), remoteFileDate)
				if err != nil {
					log.Logger.Debug(err)
				} else {
					if compareDate.After(remoteCreateAt) {
						remoteFileNameWithoutExt = remoteFileNameWithoutExt[:strings.LastIndex(remoteFileNameWithoutExt, enum.FolderSeparateFlag+remoteFileDate)]
					}
				}
			}

		}

		if remoteFileNameWithoutExt == fileNameWithoutExtAndDate {
			pendingRemove = append(pendingRemove, fmt.Sprintf("%s/%s", cfg.Remote.ArchivePath, remoteFileName))
		}
	}

	if len(pendingRemove) == 0 {
		log.Logger.Info("No removed files")
	}

	for _, remoteFileName := range pendingRemove {
		err := client.Delete(remoteFileName)
		if err != nil {
			log.Logger.Error("Delete remote file failed ", remoteFileName)
		}
	}
}
