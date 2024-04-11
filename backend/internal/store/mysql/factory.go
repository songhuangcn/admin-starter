package mysql

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/config"
	"github.com/songhuangcn/admin-template/internal/store"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type factory struct {
	db *gorm.DB
}

func NewStore() store.Factory {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database,
		config.Mysql.Options,
	)
	config := &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Info),
	}

	log.Debugf("MySQL DSN：%#v", dsn)
	db, err := gorm.Open(mysql.Open(dsn), config)
	HandleError(err)

	return &factory{
		db: db,
	}
}

func (m *factory) User() store.UserStore {
	return &userStore{
		db: m.db,
	}
}

func (m *factory) Role() store.RoleStore {
	return &roleStore{
		db: m.db,
	}
}

func (m *factory) RolesPermission() store.RolesPermissionStore {
	return &rolesPermissionStore{
		db: m.db,
	}
}
