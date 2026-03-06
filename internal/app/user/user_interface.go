package user

type UserHttpHandler interface{}

type UserUseCase interface{}

type UserRepository interface {
	Creating(user *UserEntity) (*UserEntity, error)
	FindById(userId string) (*UserEntity, error)
}
