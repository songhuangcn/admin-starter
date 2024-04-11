package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/songhuangcn/admin-template/internal/config"
	"github.com/songhuangcn/admin-template/internal/global"
	_ "github.com/songhuangcn/admin-template/internal/init" // 初始化一些库设置
	"github.com/songhuangcn/admin-template/internal/router"
	"github.com/songhuangcn/admin-template/internal/validator"
)

func main() {
	log.Debugf("环境信息：%#v", config.Env)
	log.Debugf("配置信息：%#v", config.Config)

	// 设置运行环境
	gin.SetMode(config.App.GinMode)

	// 配置验证器
	validator.Config()

	// 配置路由(在全局包中创建 engine，方便其他包中使用，比如请求处理中获取路由信息)
	router.Config(global.Engine)

	// 运行服务
	global.Engine.Run(":" + config.App.Port)
}
