package user

type UserRepository interface {
	Creating(user *UserEntity) (*UserEntity, error)
	FindById(userId string) (*UserEntity, error)
}
