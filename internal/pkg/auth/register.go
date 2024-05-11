package auth

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/passwork"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users/model"
)

func (s *Service) Register(ctx context.Context, request *model.RegisterRequest) (*model.RegisterResponse, error) {
	var userUUID uuid.UUID

	execErr := s.repo.ExecTx(ctx, func(ctx context.Context) error {
		var err error

		user, err := s.repo.UserInfoSelect(ctx, nil, &request.UserInfo.Nickname)
		if err != nil {
			return fmt.Errorf("repo.UserInfoSelect: %w", err)
		}

		if user != nil {
			return model.ErrNicknameTaken
		}

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
