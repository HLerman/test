package user

import (
	"context"

	api "github.com/HLerman/test"
	"github.com/rs/zerolog/log"
)

type Service struct {
	db Db
}

func NewService(db Db) Service {
	return Service{
		db: db,
	}
}

func (s Service) GetUsers(ctx context.Context) (api.UserArray, error) {
	var res api.UserArray
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("failed to retrieve users")
		return res, err
	}

	for _, v := range users {
		res = append(res, api.User{
			UserId:    v.UserId,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Balance:   v.Balance,
		})
	}

	return res, err
}
