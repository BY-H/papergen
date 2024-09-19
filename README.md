# cyclopropane 刷课后台看板管理系统

## 架构

采用 gin + gorm + mysql 作为项目的基本架构

## 部署

项目的对应端口为 `1020`, 这是我的生日 

### 直接部署

采用直接编译的方式直接部署

`go build -o server ./cmd/cyclopropane/server.go`

然后运行生成的对应 `server` 即可

### Docker 部署

采用 docker 进行项目的部署

`docker build -t cyclopropane .`

`docker run -d -p 1020:1020 cyclopropane`

