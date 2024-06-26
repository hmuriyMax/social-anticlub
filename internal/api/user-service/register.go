package user_service

import (
	"context"
	"fmt"
	"github.com/hmuriyMax/social-anticlub/internal/api/user-service/converters"
	"github.com/hmuriyMax/social-anticlub/internal/helpers"
	proto "github.com/hmuriyMax/social-anticlub/internal/pb/user_service"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users/model"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"unicode/utf8"
)

func (i *Implementation) Register(ctx context.Context, req *proto.RegRequest) (*proto.RegResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	registerReq := &model.RegisterRequest{
		UserInfo: &model.UserInfo{
			Nickname:   req.Info.Nickname,
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

	res, err := i.authProvider.Register(ctx, registerReq)
	if err != nil {
		return nil, converters.ToRPCErr(err)
	}

	return &proto.RegResponse{
		Status: proto.RegResponse_Success,
		UserID: res.Login.String(),
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
	case utf8.RuneCountInString(req.Info.Nickname) < 3 || utf8.RuneCountInString(req.Info.Nickname) > 15:
		return errors.New("invalid nickname")
	case req.Info.Birthday == nil:
		return errors.New("birthday is required")
	case validateDate(req.Info.Birthday) != nil:
		return fmt.Errorf("invalid birthday: %w", validateDate(req.Info.Birthday))
	}
	return nil
}

func validateDate(val *date.Date) error {
	switch {
	case val == nil:
		return errors.New("date is nil")
	case val.GetYear() < 1800 || val.GetYear() > int32(time.Now().Year()):
		return errors.New("invalid year")
	case val.GetMonth() < 1 || val.GetMonth() > 12:
		return errors.New("invalid month")
	case val.GetDay() < 1 || val.GetDay() > 31:
		return errors.New("invalid day")
	}
	return nil
}
