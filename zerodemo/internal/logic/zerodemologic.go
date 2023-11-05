package logic

import (
	"context"

	"zerodemo/internal/svc"
	"zerodemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZerodemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZerodemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZerodemoLogic {
	return &ZerodemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZerodemoLogic) Zerodemo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
