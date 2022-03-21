package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/common/cryptx"
	"mall/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/user"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	_,err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	//如果数据库里已经有这个号码，说明用户已存在
	if err == nil {
		return nil,status.Error(100,"该用户已存在")
	}
	//否则用户未存在，新建用户并把信息加入数据库
	if err == model.ErrNotFound{
		newUser := model.User{
			Name: in.Name,
			Gender: in.Gender,
			Mobile: in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt,in.Password),
		}

		res,err := l.svcCtx.UserModel.Insert(&newUser)
		if err != nil {
			return nil,status.Error(500,err.Error())
		}
		//用户ID设置为数据库最后一个插入数据的编号
		newUser.Id,err = res.LastInsertId()
		if err != nil {
			return nil,status.Error(500,err.Error())
		}

		return &user.RegisterResponse{
			Id: newUser.Id,
			Name: newUser.Name,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
		},nil
	}

	return nil, status.Error(500, err.Error())
}
