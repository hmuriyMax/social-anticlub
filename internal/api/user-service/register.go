package user_service

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"socialanticlub/internal/helpers"
	proto "socialanticlub/internal/pb/user_service"
	"socialanticlub/internal/pkg/users/model"
	"time"
)

func (i *Implementation) Register(ctx context.Context, req *proto.RegRequest) (*proto.RegResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	registerReq := &model.RegisterRequest{
		UserInfo: &model.UserInfo{
			FirstName:  req.Info.Name,
			SecondName: req.Info.Surname,
			Birthday: time.Date(
				int(req.Info.Birthday.GetYear()),
				time.Month(req.Info.Birthday.GetMonth()),
				int(req.Info.Birthday.GetDay()),
				0, 0, 0, 0,
				time.UTC,
			),
			HomeTown: req.Info.Hometown,
			About:    req.Info.Hobbies,
		},
		Password: req.Pass,
	}

	if req.Info.Gender != nil {
		registerReq.UserInfo.Gender = helpers.Ptr(model.Gender(*req.Info.Gender))
	}

	res, err := i.usersProvider.Register(ctx, registerReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.RegResponse{
		Status: proto.RegResponse_Success,
		Login:  res.Login.String(),
		UserID: res.UserID,
	}, nil
}

func validateRegister(req *proto.RegRequest) error {
	switch {
	case req.Pass == "":
		return errors.New("password is required")
	case req.Info == nil:
		return errors.New("info is required")
	case req.Info.Name == "":
		return errors.New("name is required")
	case req.Info.Birthday == nil:
		return errors.New("birthday is required")
	}
	return nil
}
