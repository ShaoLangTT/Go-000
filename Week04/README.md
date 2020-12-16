# 1.作业
```
按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。
```

# 2.目录结构

```batch
├── api
│   └── article
│       ├── article.pb.go
│       └── article.proto
├── internal
│   ├── dao
│   │    ├── tag.go
│   │    └── dao.go
│   │         
│   ├── middleware
│   │       └── jwt.go
│   ├── models
│   │   ├── article
│   │   │   ├── article.go
│   │   │   └── crud.go
│   │   ├── model.go
│   │   └── user
│   │       ├── crud.go
│   │       ├── hooks.go
│   │       └── user.go
│   └── routers
│       └── api
│            ├── v1
│            │    └── article.go
│            └── router.go
├── configs
│   ├── config.yaml
│   └── config.go
├── go.mod
├── go.sum
├── main.go
├── pkg
│   └── logger
│        └── logger.go
├── storage
│   └── logs
│        └── app.log 
├── tests
│    └── pages_test.go
│── global
│    └── db.go
└── tmp

```