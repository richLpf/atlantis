
## 项目搭建

> go创建go.md  project_dir 项目所在根目录

```
go mod init atlantis
```

## 项目结构搭建

```
|-- atlantis
	|-- api api层
	|-- application 应用启动层
	|-- services  数据库交互层
	|-- docs swagger文档生成
	|-- common  公用函数
	|-- conf    读取配置文件
	|-- dbs     数据库文件
	|-- middleware  中间件
	|-- model   数据模型
	|-- router   路由层
```

## 可能遇到的问题解决

一、加载依赖，网络问题, 解决问题

https://goproxy.io/zh/

二、swagger常用命令

swag init

## dockerfile 构建


