package user

import (
	"errors"
	"log/slog"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/database"
)

type UserRepositoryImpl struct {
	logger *slog.Logger
	db     database.Database
}

func NewUserRepository(logger *slog.Logger, conf *config.Config) UserRepository {
	return &UserRepositoryImpl{
		logger: logger,
		db:     conf.GetDb("main"),
	}
}

func (r *UserRepositoryImpl) Creating(user *UserEntity) (*UserEntity, error) {
	newUser := new(UserEntity)
	if err := r.db.Connect().Create(user).Scan(newUser).Error; err != nil {
		r.logger.Error("failed to create user", "error", err)
		return nil, errors.New("failed to create user")
	}

	return newUser, nil
}

func (r *UserRepositoryImpl) FindById(userId string) (*UserEntity, error) {
	user := new(UserEntity)
	if err := r.db.Connect().Where("UserId = ?", userId).First(user).Error; err != nil {
		r.logger.Error("failed to find user by id", "error", err)
		return nil, errors.New("user not found")
	}

	return user, nil
}
