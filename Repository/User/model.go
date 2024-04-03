package User

import (
	"context"
	"usr-service/Controller/Dto/Response"
)

type UserRepository interface {
	LoginUser(ctx context.Context, username, password string) (resp Response.UserResp, err error)
}

type user struct{}

func NewRepository() UserRepository {
	return &user{}
}
