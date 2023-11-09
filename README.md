# Yearning前后端环境配置

## 调试

**直接clone本仓库文件，部分内容已经修改**

**前端配置**

```
npm run dev
```

**后端配置**

```
go run main.go debug
```

- [x] 前后端分离调试
- [x] 与Postgres数据库建立连接
- [ ] 与其他数据库建立连接
- [x] 去除语法检查引擎（部分）
- [ ] 简化工单申请执行流程
- [ ] SQL语句执行与回显

## 部署

**前端配置**

首先下载前端项目文件

```
git clone https://github.com/cookieY/gemini-next
```

安装node.js，配置npm

随后运行命令build前端文件

```
npm install
npm run dev # 测试能否正常启动
npm build
```

此时在前端根目录下生成打包文件目录/dist

**后端配置**

首先下载后端项目文件

```
https://github.com/cookieY/Yearning
```

安装go语言环境，v1.20以上都可以（大概吧，至少1.20和1.24可以）

配置mysql数据库，版本至少为5.7

配置mysql数据库账户，创建Yearning数据库备用，修改其字符集

```
create database Yearning
alter database Yearning character set utf8mb4
```

**修改软件根目录下的conf.toml.template文件**

```
[Mysql]
Db = "Yearning"
Host = "127.0.0.1"
Port = "3306" # 没改过端口就不用动
Password = "your pwd"
User = "your account"
```

**修改完毕后，重命名该文件，把.template后缀去掉，得到conf.toml文件**

随后把***前端生成的/dist文件***，***拷贝至Yearning软件目录下/src/service/dist***

运行如下命令进行模块安装

```
go mod tidy
# 如果github卡了换国内镜像源
echo "export GO111MODULE=on" >> ~/.profile
echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
source ~/.profile
```

随后运行命令测试配置是否正常

```
go run main.go
```

运行命令初始化

```
go run main.go install
```

询问是否设置字符集选yes

启动，访问本地8000端口，默认账户admin，密码Yearning_admin

```
go run main.go run
```

**由于原项目开发者没有写readme，不敢保证上述流程为对应部署方式，但反正能跑就行**

