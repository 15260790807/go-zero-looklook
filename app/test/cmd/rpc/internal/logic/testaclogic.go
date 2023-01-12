package logic

import (
	"context"

	"looklook/app/test/cmd/rpc/internal/svc"
	"looklook/app/test/cmd/rpc/pb"
	"looklook/app/test/model"

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
	// if err := l.svcCtx.JqTestModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {

	// }); err != nil {

	// }
	user := new(model.JqTest)
	_, err := l.svcCtx.JqTestModel.Insert(l.ctx, user)
	resp := pb.TsResp{
		Msg: "rpc添加成功1112223334",
	}
	// if err != nil {
	// 	resp = pb.TsResp{
	// 		Msg: "121212",
	// 	}
	// }

	return &resp, nil
}
