# AutoBackup 自动备份程序

## 简介

AutoBackup 是一个定时自动备份工具，可以将本地指定目录打包归档，并通过 SFTP 协议上传到远程服务器，实现数据的自动化备份和归档管理。

---

## 初始化

自动在当前目录下生成config/config.yaml

```shell
./abg --init
```

## 配置示例

```yaml
appName: AutoBackup
directory: [ "" ]
cron: "*/1 * * * *"
docker: true
hideFolder: false
remote:
  protocol: sftp
  host: "example.com"
  port: "22"
  username: "xxx"
  password: ""
  sshPublicKey: "ed25519"
  archivePath: "/home/xxx/pal_backup"
archive:
  type: tar.gz
  savePreviousArchive: false
  nameFormat: '%Y%m%D%H%M'
  SortByDate: true
  storeExpired: 1
```

---

## 配置参数说明

| 参数                          | 说明                                                 |
|-----------------------------|----------------------------------------------------|
| appName                     | 应用名称                                               |
| directory                   | 需要备份的本地目录列表（支持多个目录）                                |
| cron                        | 定时任务表达式，控制备份频率（如每分钟执行一次）                           |
| docker                      | docker 模式. 增加读取/data, 归档指定目录的下一级目录,当前目录不会被归档       |
| hideFolder                  | 是否归档隐藏目录如: .hide/                                  |
| remote                      | 远程服务器相关配置                                          |
| remote.protocol             | 远程传输协议（sftp）                                       |
| remote.host                 | 远程服务器地址                                            |
| remote.port                 | 远程服务器端口                                            |
| remote.username             | 远程服务器用户名                                           |
| remote.password             | 远程服务器密码（可为空，建议使用SSHKey认证）                          |
| remote.sshPublicKey         | SSH 私钥标识（用于密钥认证）                                   |
| remote.archivePath          | 远程服务器归档文件存储路径                                      |
| archive                     | 归档相关配置                                             |
| archive.type                | 支持 zip 和  tar.gz 格式, 默认tar.gz                      |
| archive.savePreviousArchive | 保存之前的归档                                            |
| archive.SortByDate          | 是否按日期排序归档文件（true/false）,此功能搭配 `archive.nameFormat` |
| archive.nameFormat          | 归档文件命名格式（支持时间变量，如 `%Y%m%D%H%M`）                    |
| archive.storeExpired        | N天之前的归档将被删除                                        |

---

## 工作流程

1. 定时触发：根据 `cron` 表达式定时执行备份任务。
2. 目录打包：将 `directory` 指定的本地目录打包为 tar.gz 归档文件。
3. 归档命名：归档文件名根据 `archive.nameFormat` 自动生成。
4. 上传归档：通过 SFTP 协议将归档文件上传到远程服务器的 `archivePath` 目录。

---

## 适用场景

- 个人或团队的数据定时自动备份
- 服务器文件自动归档与远程同步
- 重要数据的异地容灾备份

---