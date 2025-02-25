# papergen 试卷生成系统

## 架构

采用 gin + gorm + mysql 作为项目的基本架构

## 部署

### 直接部署

Windows下，可以通过直接编译项目来进行部署

### Docker 部署

采用 docker 进行项目的部署

`docker build -t papergen .`

`docker run -d -p 1020:1020 papergen`
