# Conship

> 配置文件托管平台
>
> 年轻人的第一个开源项目

## 快速开始

> 使用 docker-compose 部署

1. 创建文件

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
