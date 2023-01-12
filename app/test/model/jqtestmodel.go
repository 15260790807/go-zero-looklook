package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ JqTestModel = (*customJqTestModel)(nil)

type (
	// JqTestModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJqTestModel.
	JqTestModel interface {
		jqTestModel
	}

	customJqTestModel struct {
		*defaultJqTestModel
	}
)

// NewJqTestModel returns a model for the database table.
func NewJqTestModel(conn sqlx.SqlConn) JqTestModel {
	return &customJqTestModel{
		defaultJqTestModel: newJqTestModel(conn),
	}
}
