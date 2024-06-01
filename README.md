English | [简体中文](README.zh-CN.md)

# Admin Starter

A admin starter by Golang. Related links:
- Blog post：[如何做好一个软件项目](https://hdgcs.com/posts/28-how-to-make-better-project)
- Ruby version: [Admin Template](https://github.com/songhuangcn/admin-template)

## Features

- Project management by Makefile
- Login authentication
- RBAC permission management, permissions automatically generated from router
- I18n multi-language support

## Tech Stack

- Gin
- Vue.js
- MySQL
- Docker
- GitLab CI
- Kubernetes

## Quick Start

1. After pulling the repo, just execute the command `make setup` locally
1. You can execute `make` or `make help` to learn other commands:
    ```
    Usage: make <TARGETS> <OPTIONS> ...

    Targets:
      setup               初始化环境
      status              服务状态
      run                 前台运行服务
      run.<SERVICE>       前台运行指定服务
      start               后台运行服务
      start.<SERVICE>     后台运行指定服务
      stop                停止服务
      stop.<SERVICE>      停止指定服务
      image               构建镜像
      image.<SERVICE>     构建指定镜像
      install             安装服务包
      install.<SERVICE>   安装指定服务包
      verify              验证和修复代码问题
      verify.<SERVICE>    验证和修复指定服务代码问题
      clean               清空所有数据
      sh.<SERVICE>        进入指定服务 Shell
      backend.locale      更新后端翻译文件
      backend.lint        检测并修复后端代码规范
      prod.build          生产环境构建
      prod.deploy         生产环境部署
      help                命令帮助
    ```

## To-do

- [ ] Project deployment
- [ ] Documentation improvement
