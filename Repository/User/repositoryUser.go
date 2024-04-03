package User

import (
	"context"
	"usr-service/Config"
	"usr-service/Controller/Dto/Response"
)

func (u *user) LoginUser(ctx context.Context, username, password string) (resp Response.UserResp, err error) {
	query := `SELECT id, username, email, phone_number FROM t_user WHERE (LOWER(email) = LOWER(?) OR phone_number = ?) AND password = ?`
	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, username, username, password).Scan(&resp.Id, &resp.Username, &resp.Email, &resp.PhoneNumber)
	return
}
