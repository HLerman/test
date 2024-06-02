package user_test

import (
	"context"
	"testing"

	api "github.com/HLerman/test"
	"github.com/HLerman/test/internal/business/models"
	"github.com/HLerman/test/internal/business/user"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestService_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := user.NewMockDb(ctrl)
	svc := user.NewService(db)

	t.Run("ok", func(t *testing.T) {
		expected := api.UserArray{
			api.User{
				Balance:   100,
				FirstName: "John",
				LastName:  "Doeuf",
				UserId:    123,
			},
			api.User{
				Balance:   200,
				FirstName: "Alain",
				LastName:  "Proviste",
				UserId:    124,
			},
		}

		db.EXPECT().GetUsers(gomock.Any()).Return([]models.User{
			{
				Balance:   100,
				FirstName: "John",
				LastName:  "Doeuf",
				UserId:    123,
			},
			{
				Balance:   200,
				FirstName: "Alain",
				LastName:  "Proviste",
				UserId:    124,
			},
		}, nil)

		users, err := svc.GetUsers(context.Background())
		require.NoError(t, err)
		require.Equal(t, expected, users)
	})
}
