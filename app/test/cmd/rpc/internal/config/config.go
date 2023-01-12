package config

import "github.com/zeromicro/go-zero/zrpc"

// Config 配置项
type Config struct {
	zrpc.RpcServerConf
	// DB 数据库
	DB struct {
		DataSource string
	}
}
