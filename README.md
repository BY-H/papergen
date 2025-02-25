# papergen 试卷生成系统

## 架构

采用 gin + gorm + mysql 作为项目的基本架构

## 部署

### 直接部署

#### Windows

在 Windows 下，可以通过安装Go环境，然后直接编译项目：

`go build -o _output/papergen.exe cmd/papergen/main.go`

运行对应的 `_output/papergen.exe` 即可完成部署

#### Linux

在 Linux 下，可以直接通过 Makefile 进行构建，直接执行 `make` 即可完成构建

执行 `make clean` 即可清除构建产物

### Docker 部署

采用 docker 进行项目的部署

`docker build -t papergen .`

`docker run -d -p 1020:1020 papergen`
