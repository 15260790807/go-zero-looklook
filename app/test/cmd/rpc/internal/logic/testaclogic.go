package logic

import (
	"context"

	"looklook/app/test/cmd/rpc/internal/svc"
	"looklook/app/test/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestacLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestacLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestacLogic {
	return &TestacLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestacLogic) Testac(in *pb.TsReq) (*pb.TsResp, error) {
	// todo: add your logic here and delete this line
    resp := pb.TsResp{
		Msg: "rpc添加成功111222333",
	}
	return &resp, nil
}
