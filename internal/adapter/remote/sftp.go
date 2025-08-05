package remote

import (
	"AutoBuckup/internal/config"
	"AutoBuckup/internal/log"
	"fmt"
	"io"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Sftp struct {
	Host        string
	Port        string
	Username    string
	Password    string
	SSHKeyPath  string
	ArchivePath string
	sftpClient  *sftp.Client
	sshClient   *ssh.Client
}

func (s *Sftp) Ls(path string) ([]string, error) {
	dirs := make([]string, 0)

	dir, err := s.sftpClient.ReadDir(path)
	if err != nil {
		return dirs, err
	}
	for _, info := range dir {
		if !info.IsDir() {
			dirs = append(dirs, info.Name())
		}
	}

	return dirs, nil
}

func (s *Sftp) Delete(filePath string) error {
	return s.sftpClient.Remove(filePath)
}

func (s *Sftp) Connect() error {
	authMethods := make([]ssh.AuthMethod, 0)
	if s.SSHKeyPath != "" {
		key, err := os.ReadFile(s.SSHKeyPath)
		if err != nil {
			log.Logger.Errorf("Unable to read key file: %v", err)
		}

		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			log.Logger.Errorf("Can't parse key: %v", err)
		}

		authMethods = append(authMethods, ssh.PublicKeys(signer))

	}
	if s.Password != "" {
		authMethods = append(authMethods, ssh.Password(s.Password))
	}
	sshConfig := &ssh.ClientConfig{
		User:            s.Username,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	log.Logger.Debugln("connect to ", addr)
	sshClient, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		log.Logger.Errorf("Unable to connect SSH: %v", err)
		if sshClient != nil {
			sshClient.Close()
		}
		return err
	}
	s.sshClient = sshClient

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Logger.Errorf("Unable to connect SFTP client: %v", err)
		if sftpClient != nil {
			sftpClient.Close()
		}
	}
	s.sftpClient = sftpClient
	return nil
}

func (s *Sftp) Upload(file *os.File, filePath, fileName string) error {
	remoteFile, err := s.sftpClient.Create(fmt.Sprintf("%s/%s", filePath, fileName))
	if err != nil {
		log.Logger.Errorf("Unable create file: %v", err)
		return err
	}
	defer remoteFile.Close()
	_, err = io.Copy(remoteFile, file)
	if err != nil {
		log.Logger.Errorf("update %s fail", fileName)
		return err
	}

	log.Logger.Infof("update %s successfully", fileName)
	return nil
}
func (s *Sftp) Mkdir(path string) error {
	err := s.sftpClient.MkdirAll(path)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sftp) Disconnect() error {
	err := s.sftpClient.Close()
	err = s.sshClient.Close()
	return err
}

func NewSftp(conf *config.Config) (*Sftp, error) {
	if conf == nil {
		return nil, fmt.Errorf("NewSftp: conf is nil")
	}
	return &Sftp{
		Host:        conf.Remote.Host,
		Port:        conf.Remote.Port,
		Username:    conf.Remote.Username,
		Password:    conf.Remote.Password,
		SSHKeyPath:  conf.Remote.SSHKeyPath,
		ArchivePath: conf.Remote.ArchivePath,
	}, nil
}
