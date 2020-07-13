package usecase

import (
	"errors"
	"testing"

	"github.com/confus1on/authez/internal/model"
	"github.com/confus1on/authez/internal/service/auth/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAuthUsecase_AuthenticationValidation(t *testing.T) {
	controller := gomock.NewController(t)

	t.Run("test authentication validation", func(t *testing.T) {
		input := model.InputAuth{
			Username:  "risky",
			Password:  "risky123",
			TableName: "users",
		}

		resultRepo := map[string]interface{}{
			"username": "risky",
			"password": "risky123",
			"fullname": "risky feryansyah",
		}

		authRepo := mock.NewMockRepositoryAuth(controller)
		authRepo.EXPECT().FindUser(input).Return(resultRepo, nil).Times(1)

		authUC := NewAuthUsecase(authRepo)
		_, err := authUC.AuthenticationValidation(input)

		require.NoError(t, err)
	})

	t.Run("test authentication validation failed empty username", func(t *testing.T) {
		input := model.InputAuth{
			Username:  "",
			Password:  "risky123",
			TableName: "users",
		}

		resultError := errors.New("invalid username or password")

		authRepo := mock.NewMockRepositoryAuth(controller)
		authRepo.EXPECT().FindUser(input).Return(nil, resultError).Times(1)

		authUC := NewAuthUsecase(authRepo)
		res, err := authUC.AuthenticationValidation(input)

		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("test authentication validation failed empty username", func(t *testing.T) {
		input := model.InputAuth{
			Username:  "risky",
			Password:  "",
			TableName: "users",
		}

		resultError := errors.New("invalid username or password")

		authRepo := mock.NewMockRepositoryAuth(controller)
		authRepo.EXPECT().FindUser(input).Return(nil, resultError).Times(1)

		authUC := NewAuthUsecase(authRepo)
		_, err := authUC.AuthenticationValidation(input)

		require.Error(t, err)
	})

	t.Run("test authentication validation failed empty username", func(t *testing.T) {
		input := model.InputAuth{
			Username:  "risky",
			Password:  "risky123",
			TableName: "",
		}

		resultError := errors.New("invalid username or password")

		authRepo := mock.NewMockRepositoryAuth(controller)
		authRepo.EXPECT().FindUser(input).Return(nil, resultError).Times(1)

		authUC := NewAuthUsecase(authRepo)
		_, err := authUC.AuthenticationValidation(input)

		require.Error(t, err)
	})

	t.Run("test authentication validation failed empty username", func(t *testing.T) {
		input := model.InputAuth{
			Username:  "risky",
			Password:  "risky",
			TableName: "users",
		}

		resultError := errors.New("invalid username or password")

		authRepo := mock.NewMockRepositoryAuth(controller)
		authRepo.EXPECT().FindUser(input).Return(nil, resultError).Times(1)

		authUC := NewAuthUsecase(authRepo)
		_, err := authUC.AuthenticationValidation(input)

		require.Error(t, err)
	})
}
