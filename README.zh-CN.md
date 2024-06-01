[English](README.md) | 简体中文

# Admin Starter

Golang 管理系统启动器。相关链接：
- 博文：[如何做好一个软件项目](https://hdgcs.com/posts/28-how-to-make-better-project)
- Ruby 版：[Admin Template](https://github.com/songhuangcn/admin-template)

## 功能

- Makefile 项目管理
- 登录认证
- RBAC 权限管理，权限由路由表自动生成
- I18n 多语言支持

## 技术栈

- Gin
- Vue.js
- MySQL
- Docker
- GitLab CI
- Kubernetes

## 快速开始

1. 拉取项目后，直接本地执行 `make setup` 即可
1. 可以执行 `make` 或 `make help`，了解其他命令：
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

## 代办

- [ ] 项目部署
- [ ] 文档完善
