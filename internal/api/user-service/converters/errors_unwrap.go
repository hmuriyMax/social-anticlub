package converters

import (
	"context"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users/model"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToRPCErr(err error) error {
	switch {
	case errors.Is(err, context.Canceled):
		return status.Error(codes.Canceled, err.Error())
	case errors.Is(err, context.DeadlineExceeded):
		return status.Error(codes.DeadlineExceeded, err.Error())
	case errors.Is(err, model.ErrNicknameTaken):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, model.ErrNoUser):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, model.ErrWrongPassword):
		return status.Error(codes.Unauthenticated, err.Error())
	case status.Code(err) == codes.Unknown:
		return status.Error(codes.Internal, err.Error())
	default:
		return err
	}
}
