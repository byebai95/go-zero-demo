package logic

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
	"time"
	"user/model"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) InsertUser(ctx context.Context, w http.ResponseWriter) (resp *types.Response, err error) {
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Name:     "insert-02",
		Type:     1,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})
	if err != nil {
		httpx.ErrorCtx(ctx, w, err)
	}
	id, err := res.LastInsertId()
	return &types.Response{Message: strconv.FormatInt(id, 10)}, nil
}

func (l *UserLogic) ListUser(ctx context.Context, w http.ResponseWriter) (resp any, err error) {
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, 1)
	if err != nil {
		httpx.ErrorCtx(ctx, w, err)
	}
	//jsonStr, err := json.Marshal(res)
	//return &types.Response{Message: res}, nil
	return res, nil
}
