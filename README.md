# AutoBackup G

## Introduction

AutoBackup is a scheduled automatic backup tool that can archive specified local directories and upload them to a remote server via the SFTP protocol, enabling automated data backup and archive management.

---

## Initialization

Automatically generates `config/config.yaml` in the current directory.

```shell
./abg --init
```

## Configuration Example

```yaml
appName: AutoBackup
directory: [ "" ]
cron: "*/1 * * * *"
remote:
  protocol: sftp
  host: "example.com"
  port: "22"
  username: "xxx"
  password: ""
  sshPublicKey: "ed25519"
  archivePath: "/home/xxx/pal_backup"
archive:
  nameFormat: '%Y%m%D%H%M'
  SortByDate: true
```

---

## Configuration Parameters

| Parameter                    | Description                                              |
|------------------------------|---------------------------------------------------------|
| appName                      | Application name                                        |
| directory                    | List of local directories to back up (supports multiple)|
| cron                         | Cron expression to control backup frequency (e.g., every minute) |
| remote                       | Remote server configuration                             |
| remote.protocol              | Remote transfer protocol (sftp)                         |
| remote.host                  | Remote server address                                   |
| remote.port                  | Remote server port                                      |
| remote.username              | Remote server username                                  |
| remote.password              | Remote server password (optional, SSH key authentication recommended) |
| remote.sshPublicKey          | SSH private key identifier (for key authentication)     |
| remote.archivePath           | Remote server archive storage path                      |
| archive                      | Archive-related configuration                           |
| archive.nameFormat           | Archive file naming format (supports time variables, e.g., `%Y%m%D%H%M`) |
| archive.SortByDate           | Whether to sort archive files by date (true/false)      |

---

## Workflow

1. Scheduled Trigger: Executes backup tasks according to the `cron` expression.
2. Directory Packaging: Packages the local directories specified in `directory` into a tar.gz archive file.
3. Archive Naming: Archive file names are automatically generated based on `archive.nameFormat`.
4. Upload Archive: Uploads the archive file to the `archivePath` directory on the remote server via SFTP.

---

## Use Cases

- Scheduled automatic data backup for individuals or teams
- Automatic server file archiving and remote synchronization
- Offsite disaster recovery backup for important data

---