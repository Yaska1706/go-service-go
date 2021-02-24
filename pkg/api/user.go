package api

// UserService contains the methods of the user service
type UserService interface {
	NewUser(name string) (string, error)
}

// User repository is what lets our service do db operations without knowing anything about the implementation
type UserRepository interface {
	CreateUser(name string) (string, error)
}

type userService struct {
	storage UserRepository
}

func NewUserService(usrRepo UserRepository) UserService {
	return &userService{usrRepo}
}

func (u *userService) NewUser(name string) (string, error) {
	return u.storage.CreateUser("")
}
