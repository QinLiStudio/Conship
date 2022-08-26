# Conship

> 配置文件托管平台 - 年轻人的第一个开源项目
> 
> [Github](https://github.com/QinLiStudio/Conship) •
> [DockerHub](https://hub.docker.com/r/qlstudio/conship) •
> [前端地址](https://github.com/QinLiStudio/Conship-web) •
> [接口文档](https://www.apifox.cn/apidoc/project-1111612/api-23876845) •
> [关于我们](https://ql.sylu.edu.cn)

## 介绍

如你所见，Conship 是一个配置文件托管平台，可以将你的配置文件分享给其他人。

类似于 Gist，但更加简洁优雅，并且在国内也可以部署并访问。

每个配置文件被称为一个 Meta，上传配置文件后，会提供一个唯一的链接（URL），可以被其他人访问。

除此之外，还提供一个密钥（Secret），你可以通过密钥来访问修改、删除配置文件。

选用了 Apifox 作为接口文档管理工具，可在浏览器进行模拟请求。

后端提供了增、删、改、查四个接口：[接口文档](https://www.apifox.cn/apidoc/project-1111612/api-23876845)

已经部署了一份在我们的集群上：[后端地址](https://qlapi.sylu.edu.cn/conship/meta/UOUjAag)

前端正在开发中：[前端地址](https://github.com/QinLiStudio/Conship-web)

## 特性

使用 Golang + Gin + Gorm 作为后端框架。

使用 PostgreSQL 作为主数据库，Redis 作为缓存数据库。
  
使用 Github Action 自动打包发布 Docker 镜像。

## 部署

> 使用 docker-compose 部署

1. 创建并进入路径

```bash
mkdir -p conship/config && cd conship
```

2. 创建 `config/config.toml`

```toml
[Mode]
# 运行模式
RunMode = "release" # debug:调试 ｜ test:测试 ｜ release:正式

[Http]
# 服务运行配置
Url = "https://qlapi.sylu.edu.cn/conship"
Host = "127.0.0.1" # 监听主机
Port = 8080 # 监听配置

[Redis]
# redis 配置
Addr = "redis:6379" # 地址
Password = "" # 密码

[Postgres]
# postgres 配置
Host = "postgres" # 地址
Port = 5432 # 端口
User = "conship" # 用户名
Password = "conship" # 密码
DBName = "conship" # 数据库

[Limit]
# 访问限制
Content = 2 # 文件大小限制 MB
Request = 1000 # 每小时 1000 次

[Cors]
# 跨域配置
Enable = true # 是否开启跨域限制

AllowOrigins = ["*"] # 允许列表：* 表示全部允许
AllowMethods = ["GET", "POST", "PUT", "DELETE"] # 允许请求方式
AllowHeaders = [] # 允许特殊请求头
```

3. 创建 `docker-compose.yml`

```yml
version: '3'

services:
  conship:
    image: qlstudio/conship:latest
    restart: unless-stopped
    volumes:
      - ./log:/app/log
      - ./config:/app/config
    ports:
      - '8080:8080'

  postgres:
    image: postgres:14-alpine
    environment:
      POSTGRES_DB: conship
      POSTGRES_USER: conship
      POSTGRES_PASSWORD: conship
    volumes:
      - ./postgres:/var/lib/postgresql/data
    # ports:
    #   - '5432:5432'

  redis:
    image: redis:6-alpine
    # ports:
    #   - '6379:6379'
```

4. 运行

```bash
docker pull qlstudio/conship:latest
docker-compose up -d
```
