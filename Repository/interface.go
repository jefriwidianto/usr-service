package Repository

import "usr-service/Repository/User"

type Repository struct {
	User User.UserRepository
}

var ApplicationRepository = Repository{
	User: User.NewRepository(),
}
