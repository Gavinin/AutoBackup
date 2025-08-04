package services

import (
	"AutoBuckup/internal/adapter/remote"
	"AutoBuckup/internal/config"
	"strings"
)

func SelectRemoteProtocolClient(cfg *config.Config) (remote.IFileTransProtocol, error) {
	switch strings.ToLower(cfg.Remote.Protocol) {
	case "sftp":
		return remote.NewSftp(cfg)

	default:
		return remote.NewSftp(cfg)
	}
}
