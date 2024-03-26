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
		userID   int64
		userUUID uuid.UUID
	)

	execErr := s.repo.ExecTx(ctx, func(ctx context.Context) error {
		var err error

		userID, err = s.repo.UserInfoCreate(ctx, request.UserInfo)
		if err != nil {
			return fmt.Errorf("repo.UserInfoCreate: %w", err)
		}

		userUUID, err = uuid.NewRandom()
		if err != nil {
			return fmt.Errorf("failed to create user uuid: %w", err)
		}

		err = s.repo.UserAuthInsert(ctx, &model.Login{
			ID:       userID,
			Login:    userUUID,
			PassHash: passwork.GetHash(request.Password),
		})
		if err != nil {
			return fmt.Errorf("repo.UserAuthInsert: %w", err)
		}

		return nil
	})

	return &model.RegisterResponse{
		Login:  userUUID,
		UserID: userID,
	}, execErr
}
