package testgroup

import (
	"context"

	"looklook/app/test/cmd/api/internal/svc"
	"looklook/app/test/cmd/api/internal/types"
	"looklook/app/test/cmd/rpc/test"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/pkg/errors"
	"github.com/jinzhu/copier"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test(req *types.TsReq) (*types.TsResp,error) {
	// todo: add your logic here and delete this line
	rpcresp, err := l.svcCtx.TestRpc.Testac(l.ctx, &test.TsReq{
		OrderSn:   "argea",
		ServiceType: "serviceType", 
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", rpcresp)
	}
	var resp types.TsResp
	_ = copier.Copy(&resp, rpcresp)
	return &resp, nil 
}
