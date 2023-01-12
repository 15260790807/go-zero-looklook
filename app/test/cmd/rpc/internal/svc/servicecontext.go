package svc

import (
	"looklook/app/test/cmd/rpc/internal/config"
	"looklook/app/test/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// ServiceContext 上下文
type ServiceContext struct {
	Config      config.Config
	JqTestModel model.JqTestModel
}

// NewServiceContext 创建
func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		JqTestModel: model.NewJqTestModel(sqlConn),
	}
}
