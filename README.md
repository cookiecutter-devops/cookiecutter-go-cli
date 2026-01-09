# cookiecutter-go-cli

一个使用 cookiecutter 生成的 Go 命令行工具项目的模板。

使用三方库：

- cobra
- viper
- zerolog


## 使用

```sh
cookiecutter gh:cookiecutter-devops/cookiecutter-go-cli
  [1/9] app_name (gocoo): git-demo
  [2/9] env_prefix (git-demo):
  [3/9] bin_name (git-demo):
  [4/9] config_type (yaml):
  [5/9] module_name (github.com/cookiecutter-devops/git-demo):
  [6/9] go_version (Enter Go version (default: 1.23.3):): 1.21.8
  [7/9] git (y): y
  [8/9] subcommands (n): show
  [9/9] init_version (0.0.1):
```


```sh
go env -w GOSUMDB=off
go mod tidy

# 需要先将代码上传到仓库，否则 Git ref: HEAD-DIRTY查询不到数值
make build
```


## git-demo为例

### 配置

conifg.yml 文件为项目的配置文件，默认内容如下：

```yaml
log:
  level: "info" # 日志级别: debug, info, warn, error, fatal, panic
  file: "" # 日志输出文件路径，留空则不输出到文件
  caller: false # 是否在日志中包含调用者信息（文件和行号）
  json: false # 是否以JSON格式输出日志
```

### 测试

```sh
$ ./bin/git-demo
A longer description

Usage:
  git-demo [flags]
  git-demo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  logtest     Test logging functionality with different log levels
  version     Print the version details

Flags:
  -c, --config string      yaml config file
  -h, --help               help for git-demo
      --log-caller         Include the caller file and line number
      --log-file string    Write logs in json format to this file
      --log-json           Log as json messages
      --log-level string   Set the log level (debug, info, warn, error, fatal, panic) (default "info")

Use "git-demo [command] --help" for more information about a command.



# 查看版本
$ ./bin/git-demo version
Version: 0.0.1
Build Date: 2026-01-09-14:32:16
Git ref: a342433d9a7770cf63cee6a8cb3993e53e219311-DIRTY
sha256: 287ac8eeee5908de92d2edfcb62256863f2d1039e834253b3bdd3f0fb7fa755d
OS: linux
Arch: amd64


# 使用默认日志级别 (info)
$ ./bin/git-demo logtest
2:35PM INF This is an info message - shown when log level is info or lower
2:35PM WRN This is a warning message - shown when log level is warn or lower
2:35PM ERR This is an error message - shown when log level is error or lower


# 设置日志级别为 debug (会显示所有日志)
$ ./bin/git-demo logtest --log-level debug

# 设置日志级别为 warn (只会显示 warn 和 error 级别的日志)
$ ./bin/git-demo logtest --log-level warn

# 使用配置文件
$ ./bin/git-demo logtest -c bin/config.yml

# 显示调用者信息
$ ./bin/git-demo logtest -c bin/config.yml --log-caller

# JSON格式日志
$ ./bin/git-demo logtest --log-json

# 日志文件输出
$ ./bin/git-demo logtest -c bin/config.yml --log-file bin/log.log

# 使用 JSON 格式输出日志
$ ./bin/git-demo logtest --log-json

# 结合其他日志选项
$ ./bin/git-demo logtest --log-level debug --log-json --log-caller
$ ./bin/git-demo logtest -c bin/config.yml --log-json --log-file bin/log.json
```
