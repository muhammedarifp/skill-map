package managers

type UserManager struct {
	// Db Client
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (um *UserManager) CreateUser() {}
