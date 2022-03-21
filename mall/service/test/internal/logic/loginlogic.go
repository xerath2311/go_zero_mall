package logic

import (
	"context"

	"mall/service/test/internal/svc"
	"mall/service/test/remote"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  登录
func (l *LoginLogic) Login(in *remote.Request) (*remote.Response, error) {
	// todo: add your logic here and delete this line

	return &remote.Response{}, nil
}
