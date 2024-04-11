package init

import (
	"github.com/sirupsen/logrus"
	"github.com/songhuangcn/admin-template/internal/config"
)

func init() {
	level, err := logrus.ParseLevel(config.App.LogLevel)
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(level)
}
