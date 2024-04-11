COMPOSE  = docker compose -f docker-compose.yml
BACKEND  = $(COMPOSE) run --rm backend
FRONTEND = $(COMPOSE) run --rm frontend

.DEFAULT_GOAL := help

## setup: 初始化环境
.PHONY: setup
setup: image install verify

## status: 服务状态
.PHONY: status
status:
	$(COMPOSE) ps

## run: 前台运行服务
.PHONY: run
run:
	$(COMPOSE) up --remove-orphans
	
## run.<SERVICE>: 前台运行指定服务
.PHONY: run.%
run.%:
	$(COMPOSE) up --remove-orphans $*

## start: 后台运行服务
.PHONY: start
start:
	$(COMPOSE) up --remove-orphans
	
## start.<SERVICE>: 后台运行指定服务
.PHONY: start.%
start.%:
	$(COMPOSE) up --remove-orphans $*

## stop: 停止服务
.PHONY: stop
stop:
	$(COMPOSE) down --remove-orphans
	
## stop.<SERVICE>: 停止指定服务
.PHONY: stop.%
stop.%:
	$(COMPOSE) down --remove-orphans $*

## image: 构建镜像
.PHONY: image
image:
	$(COMPOSE) build
	
## image.<SERVICE>: 构建指定镜像
.PHONY: image.%
image.%:
	$(COMPOSE) build $*
	
## install: 安装服务包
.PHONY: install
install: install.backend install.frontend
	
## install.<SERVICE>: 安装指定服务包
.PHONY: install.backend
install.backend:
	$(COMPOSE) run --rm backend go mod tidy

.PHONY: install.frontend
install.frontend:
	$(COMPOSE) run --rm frontend pnpm install

## verify: 验证和修复代码问题
.PHONY: verify
verify: verify.backend verify.frontend
	
## verify.<SERVICE>: 验证和修复指定服务代码问题
.PHONY: verify.backend
verify.backend: backend.locale backend.lint
	
.PHONY: verify.frontend
verify.frontend:
	$(FRONTEND) pnpm lint --fix

## clean: 清空所有数据
.PHONY: clean
clean:
	$(COMPOSE) down --volumes --remove-orphans

## sh.<SERVICE>: 进入指定服务 Shell
.PHONY: sh.%
sh.%:
	$(COMPOSE) exec -it $* sh

## backend.locale: 更新后端翻译文件
.PHONY: backend.locale
backend.locale:
	$(BACKEND) xspreak --directory internal --output-dir locales --default-domain app

## backend.lint: 检测并修复后端代码规范
.PHONY: backend.lint
backend.lint:
	$(BACKEND) goimports -w .

BACKEND_TAG  = latest
FRONTEND_TAG = latest

## prod.build: 生产环境构建
.PHONY: prod.build
prod.build:
#	@docker build --target prod -t registry.gitlab.com/songhuangcn/admin-starter/backend:$(BACKEND_TAG) ./backend/
	@docker build --target prod -t registry.gitlab.com/songhuangcn/admin-starter/frontend:$(FRONTEND_TAG) --platform linux/amd64 ./frontend/
#	@docker push registry.gitlab.com/songhuangcn/admin-starter/backend:$(BACKEND_TAG)
	@docker push registry.gitlab.com/songhuangcn/admin-starter/frontend:$(FRONTEND_TAG)

## prod.deploy: 生产环境部署
.PHONY: prod.deploy
prod.deploy:
	@kubectl apply -f deployment/kubernetes/admin-starter.yml

## help: 命令帮助
.PHONY: help
help: Makefile
	@printf "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
