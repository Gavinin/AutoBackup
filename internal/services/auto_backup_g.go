package services

import (
	"AutoBuckup/internal/config"
	"AutoBuckup/internal/log"
	"AutoBuckup/internal/util"
	"github.com/robfig/cron/v3"
	"os"
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

	err = remoteProtocolClient.Connect()
	if err != nil {
		log.Logger.Error(err)
		return
	}
	defer remoteProtocolClient.Disconnect()

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
