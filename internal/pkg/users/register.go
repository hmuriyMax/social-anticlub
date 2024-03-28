package users

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"socialanticlub/internal/pkg/passwork"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Service) Register(ctx context.Context, request *model.RegisterRequest) (*model.RegisterResponse, error) {
	var (
		userUUID uuid.UUID
	)

	execErr := s.repo.ExecTx(ctx, func(ctx context.Context) error {
		var err error

		userUUID, err = s.repo.UserInfoInsert(ctx, request.UserInfo)
		if err != nil {
			return fmt.Errorf("repo.UserInfoInsert: %w", err)
		}

		err = s.repo.UserAuthInsert(ctx, &model.Login{
			UserUUID: userUUID,
			PassHash: passwork.GetHash(request.Password),
		})
		if err != nil {
			return fmt.Errorf("repo.UserAuthInsert: %w", err)
		}

		return nil
	})

	return &model.RegisterResponse{
		Login: userUUID,
	}, execErr
}
