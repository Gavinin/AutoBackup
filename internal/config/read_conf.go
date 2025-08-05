package config

import (
	"AutoBuckup/internal/log"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Config struct {
	AppName    string   `yaml:"appName"`
	Directory  []string `yaml:"directory"`
	Cron       string   `yaml:"cron"`
	Docker     bool     `yaml:"docker"`
	Debug      bool     `yaml:"debug,omitempty"`
	Remote     Remote   `yaml:"remote"`
	Archive    Archive  `yaml:"archive"`
	HideFolder bool     `yaml:"hideFolder"`
}

type Remote struct {
	Protocol    string `yaml:"protocol"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	SSHKeyPath  string `yaml:"sshPublicKey"`
	ArchivePath string `yaml:"archivePath"`
}

type Archive struct {
	Type         string `yaml:"type"`
	SavePrevious bool   `yaml:"savePreviousArchive"`
	NameFormat   string `yaml:"nameFormat"`
	SortByDate   bool   `yaml:"SortByDate"`
	TmpFilePath  string `yaml:"tmpFilePath,omitempty"`
	StoreExpired int    `yaml:"storeExpired"`
}

func ReadConfig() *Config {
	filePath := fmt.Sprintf("%s/%s", dirPath, confFileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Logger.Errorf("Read file fail: %v", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Logger.Errorf("Analysis YAML fail: %v", err)
	}

	err = ParseSSHKey(&cfg)
	if err != nil {
		log.Logger.Errorf("ParseSSHKey fail: %v", err)
		return nil
	}

	if cfg.Archive.TmpFilePath == "" {
		cfg.Archive.TmpFilePath = "."
	}
	if cfg.Archive.TmpFilePath[len(cfg.Archive.TmpFilePath)-1] == '/' {
		cfg.Archive.TmpFilePath = cfg.Archive.TmpFilePath[:len(cfg.Archive.TmpFilePath)-1]
	}

	for i, s := range cfg.Directory {
		if s == "" || s == "/" {
			continue
		}
		if s[len(s)-1:] == "/" {
			cfg.Directory[i] = s[:len(s)-1]
		}
	}

	if cfg.Docker {
		cfg.Directory = append(cfg.Directory, "/data")
	}

	return &cfg
}

func ParseSSHKey(cfg *Config) error {
	sshKeyPath := ""
	if strings.Contains(cfg.Remote.SSHKeyPath, "/") {
		sshKeyPath = cfg.Remote.SSHKeyPath
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Logger.Errorf("Get home dir fail: %v", err)
			return err
		}
		sshKeyPath = fmt.Sprintf("%s/.ssh/%s", home, cfg.Remote.SSHKeyPath)
	}

	cfg.Remote.SSHKeyPath = sshKeyPath

	return nil
}
